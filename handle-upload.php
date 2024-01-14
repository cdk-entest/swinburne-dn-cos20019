<html>

<head>

    <style>
        :root {
            box-sizing: border-box;
        }

        *,
        ::before,
        ::after {
            box-sizing: inherit;
        }

        .body {
            background-color: antiquewhite;
        }

        .container {
            max-width: 800px;
            margin-left: auto;
            margin-right: auto;
        }
    </style>

</head>

<body class="body">
    <div class="container">

        <?php

        $servername = "localhost";
        $username = "dev";
        $password = "Admin2024";
        $dbname = "demo";

        // Create connection
        $conn = new mysqli($servername, $username, $password, $dbname);
        // Check connection
        if ($conn->connect_error) {
            die("Connection failed: " . $conn->connect_error);
        }

        // uploaded file name 
        $filename = $_FILES['myFile']['name'];

        // save uploaded file 
        move_uploaded_file($_FILES['myFile']['tmp_name'], "./static/" . basename($filename));

        // create a record in database 
        $sql = "INSERT INTO book(author, title, amazon, image) VALUES ('Hai Tran', 'Deep Learning', '', '$filename')";
        $result = $conn->query($sql);

        // response to browser 
        echo "<h1> Sucessfully upload file $filename </h1>"
        ?>

    </div>
</body>

</html>