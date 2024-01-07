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
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime/types"
)

func main()  {

	mux := http.NewServeMux()

	// home page
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/bedrock.html")
	})

	// book page
	mux.HandleFunc("/book", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/book.html")
	})

	// upload page
	mux.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			http.ServeFile(w, r, "./static/upload.html")
		case "POST":
			uploadFile(w, r)
		}
	})

	// bedrock page
	mux.HandleFunc("/bedrock-stream", bedrock)

	// create web server 
	server := &http.Server{
		Addr: ":3000",
		Handler: mux,
		ReadTimeout: 30* time.Second,
		WriteTimeout: 30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// enable logging
	log.Fatal(server.ListenAndServe())
	
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
		Prompt: prompt,
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
			Body: payloadBytes,
			ModelId: aws.String("anthropic.claude-v2"),
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
				fmt.Println("Damn, no flush");
			 }

		case *types.UnknownUnionMember:
			fmt.Println("unknown tag:", v.Tag)

		default:
			fmt.Println("union is nil or unknown type")
		}
	}
}

// stream bedorck response to web client 
func SendStream(message string) (string, error)  {
	prompt := "" + fmt.Sprintf(claudePromptFormat, message)

	payload := Request{
		Prompt: prompt,
		MaxTokensToSample: 2048,
	}

	payloadBytes, error := json.Marshal(payload)

	if error != nil {
		return "", error
	}

	output, error := brc.InvokeModelWithResponseStream(
		context.Background(),
		&bedrockruntime.InvokeModelWithResponseStreamInput{
			Body: payloadBytes,
			ModelId: aws.String("anthropic.claude-v2"),
			ContentType: aws.String("application/json"),
		},
	)

	if error != nil {
		return "", error
	}

	for event := range output.GetStream().Events() {
		switch v := event.(type) {
		case *types.ResponseStreamMemberChunk:

			//fmt.Println("payload", string(v.Value.Bytes))

			var resp Response
			err := json.NewDecoder(bytes.NewReader(v.Value.Bytes)).Decode(&resp)
			if err != nil {
				return "", err
			}

			fmt.Println(resp.Completion)

		case *types.UnknownUnionMember:
			fmt.Println("unknown tag:", v.Tag)

		default:
			fmt.Println("union is nil or unknown type")
		}
	}
	
	return "", error
}

// server handle uploaded file 
func uploadFile(w http.ResponseWriter, r *http.Request)  {

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

	// Create file 
	dest, error := os.Create(handler.Filename)
	if error != nil {
		return 
	}
	defer dest.Close()

	// Copy uploaded file to dest 
	if _, err := io.Copy(dest, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Successfully Uploaded File\n")

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

type HelloHandler struct {}

type Query struct {
	Topic string `json:"topic"`
}