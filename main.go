package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"text/template"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime/types"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// aws rds
// const HOST = "database-1.c9y4mg20eppz.ap-southeast-1.rds.amazonaws.com"
// const USER = "postgresql"
// const PASS = "Admin2024"
// const DBNAME = "demo"

// local db
const HOST = "localhost"
const USER = "postgres"
const DBNAME = "dvdrental"
const PASS = "Mike@865525"

type Book struct {
	ID          uint
	Title       string
	Author      string
	Amazon      string
	Image       string
	Description string
}

func main() {

	// db init
	dns := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v", HOST, "5432", USER, PASS, DBNAME)
	db, _ := gorm.Open(postgres.Open(dns), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			NoLowerCase:   false,
			SingularTable: true,
		},
	})

	mux := http.NewServeMux()

	// home page
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/bedrock.html")
	})

	// book page
	mux.HandleFunc("/book", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/book.html")
	})

	mux.HandleFunc("/postgresql", func(w http.ResponseWriter, r *http.Request) {

		// query a list of book []Book
		books := getBooks(db)

		// load template
		tmpl, error := template.ParseFiles("./static/book-template.html")

		if error != nil {
			fmt.Println(error)
		}

		// pass data to template and write to writer
		tmpl.Execute(w, books)
	})

	// upload page
	mux.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			http.ServeFile(w, r, "./static/upload.html")
		case "POST":
			uploadFile(w, r, db)
		}
	})

	// bedrock page
	mux.HandleFunc("/bedrock-stream", bedrock)

	// create web server
	server := &http.Server{
		Addr:           ":3000",
		Handler:        mux,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// static files
	mux.Handle("/demo/", http.StripPrefix("/demo/", http.FileServer(http.Dir("./static"))))

	// enable logging
	log.Fatal(server.ListenAndServe())

}

func getBooks(db *gorm.DB) []Book {
	var books []Book

	db.Limit(10).Find(&books)

	for _, book := range books {
		fmt.Println(book.Title)
	}

	return books
}

func uploadFile(w http.ResponseWriter, r *http.Request, db *gorm.DB) {

	// maximum upload file of 10 MB files
	r.ParseMultipartForm(10 << 20)

	// Get handler for filename, size and heanders
	file, handler, error := r.FormFile("myFile")
	if error != nil {
		fmt.Println("Error")
		fmt.Println(error)
		return
	}

	defer file.Close()
	fmt.Printf("upload file %v\n", handler.Filename)
	fmt.Printf("file size %v\n", handler.Size)
	fmt.Printf("MIME header %v\n", handler.Header)

	// upload file to s3
	// _, error = s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
	// 	Bucket: aws.String("cdk-entest-videos"),
	// 	Key:    aws.String("golang/" + handler.Filename),
	// 	Body:   file,
	// })

	// if error != nil {
	// 	fmt.Println("error upload s3")
	// }

	// Create file
	dest, error := os.Create("./static/" + handler.Filename)
	if error != nil {
		return
	}
	defer dest.Close()

	// Copy uploaded file to dest
	if _, err := io.Copy(dest, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// create a record in database
	db.Create(&Book{
		Title:       "Database Internals",
		Author:      "Hai Tran",
		Description: "Hello",
		Image:       handler.Filename,
	})

	fmt.Fprintf(w, "Successfully Uploaded File\n")

}

// promt format
const claudePromptFormat = "\n\nHuman: %s\n\nAssistant:"

// bedrock runtime client
var brc *bedrockruntime.Client

// init bedorck credentials connecting to aws
func init() {

	region := os.Getenv("AWS_REGION")
	if region == "" {
		region = "us-east-1"
	}

	cfg, err := config.LoadDefaultConfig(context.Background(), config.WithRegion(region))
	if err != nil {
		log.Fatal(err)
	}

	brc = bedrockruntime.NewFromConfig(cfg)
}

// bedrock handler request
func bedrock(w http.ResponseWriter, r *http.Request) {

	var query Query
	var message string

	// parse mesage from request
	error := json.NewDecoder(r.Body).Decode(&query)

	if error != nil {
		message = "how to learn japanese as quick as possible?"
		panic(error)
	}

	message = query.Topic

	fmt.Println(message)

	prompt := "" + fmt.Sprintf(claudePromptFormat, message)

	payload := Request{
		Prompt:            prompt,
		MaxTokensToSample: 2048,
	}

	payloadBytes, error := json.Marshal(payload)

	if error != nil {
		fmt.Fprintf(w, "ERROR")
		// return "", error
	}

	output, error := brc.InvokeModelWithResponseStream(
		context.Background(),
		&bedrockruntime.InvokeModelWithResponseStreamInput{
			Body:        payloadBytes,
			ModelId:     aws.String("anthropic.claude-v2"),
			ContentType: aws.String("application/json"),
		},
	)

	if error != nil {
		fmt.Fprintf(w, "ERROR")
		// return "", error
	}

	for event := range output.GetStream().Events() {
		switch v := event.(type) {
		case *types.ResponseStreamMemberChunk:

			//fmt.Println("payload", string(v.Value.Bytes))

			var resp Response
			err := json.NewDecoder(bytes.NewReader(v.Value.Bytes)).Decode(&resp)
			if err != nil {
				fmt.Fprintf(w, "ERROR")
				// return "", err
			}

			fmt.Println(resp.Completion)

			fmt.Fprintf(w, resp.Completion)
			if f, ok := w.(http.Flusher); ok {
				f.Flush()
			} else {
				fmt.Println("Damn, no flush")
			}

		case *types.UnknownUnionMember:
			fmt.Println("unknown tag:", v.Tag)

		default:
			fmt.Println("union is nil or unknown type")
		}
	}
}

type Request struct {
	Prompt            string   `json:"prompt"`
	MaxTokensToSample int      `json:"max_tokens_to_sample"`
	Temperature       float64  `json:"temperature,omitempty"`
	TopP              float64  `json:"top_p,omitempty"`
	TopK              int      `json:"top_k,omitempty"`
	StopSequences     []string `json:"stop_sequences,omitempty"`
}

type Response struct {
	Completion string `json:"completion"`
}

type HelloHandler struct{}

type Query struct {
	Topic string `json:"topic"`
}
