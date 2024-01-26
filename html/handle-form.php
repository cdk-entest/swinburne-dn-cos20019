<html>

<head>
  <title>Handle Form</title>
  <style>
    :root {
      box-sizing: border-box;
    }

    *,
    ::before,
    ::after {
      box-sizing: inherit;
    }

    body {
      background-color: antiquewhite;
    }
  </style>
</head>

<body>
  <div>
    <?php

    // database connection 
    $servername = "localhost";
    $username = "dev";
    $password = "Admin2024";
    $dbname = "demo";

    // create connection
    $conn = new mysqli($servername, $username, $password, $dbname);
    // check connection
    if ($conn->connect_error) {
      die("Connection failed: " . $conn->connect_error);
    }

    // extract form data 
    $filename = $_FILES["file"]["name"];
    $title = $_POST["title"];
    $description = $_POST["description"];
    $date = $_POST["date"];
    $keyword = $_POST["keyword"];

    // save file to server 
    move_uploaded_file($_FILES["file"]["tmp_name"], "./../static/" . basename($filename));

    // create a record in database 
    $sql = "INSERT INTO photo(title, description, date, image, keyword) VALUES ('$title', '$description', '$date', '$filename', '$keyword')";
    $result = $conn->query($sql);

    // resposne to browser 
    echo "<h1>" . $_POST["title"] . "</h1>";
    echo "<h1>" . $_FILES["file"]["name"] . "</h1>";
    echo "<h1>" . $_POST["description"] . "</h1>";
    echo "<h1>" . $_POST["date"] . "</h1>";
    echo "<h1>" . $_POST["keyword"] . "</h1>";
    echo "<a href='./photolookup.php'>Photo Lookup</a>";

    // close conn 
    $conn->close();
    ?>
  </div>
</body>

</html>