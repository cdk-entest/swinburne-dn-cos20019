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

    .grid {
      display: grid;
      row-gap: 10px;
      column-gap: 10px;
      grid-template-columns: repeat(1, minmax(0, 1fr));
    }

    .card {
      margin-left: 4px;
      margin-right: 4px;
      padding: 0.5em;
      background-color: white;
      width: 100%;
    }

    @media (min-width: 800px) {
      .grid {
        grid-template-columns: repeat(2, minmax(0, 1fr));
      }
    }

    .image {
      float: left;
      height: auto;
      width: 128px;
      margin-right: 6px;
    }

    .title {
        font: bold;
        margin-bottom: 8px;
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

    $sql = "SELECT id, author, title, amazon, image FROM book";
    $result = $conn->query($sql);

    if ($result->num_rows > 0) {
      echo "<div class='grid'>";
      // output data of each row
      while ($row = $result->fetch_assoc()) {
        echo "<div class='card'>"
        . "<h4 class='title'>" .$row["title"] ."</h4>"
        . "<h4 class='title'>" .$row["author"] ."</h4>"
        . "<img src='https://d2cvlmmg8c0xrp.cloudfront.net/web-css/singapore.jpg' class='image' />"
        . "<p>" . "Lorem ipsum, dolor sit amet consectetur adipisicing elit. Officia officiis voluptates ab eum totam atque deleniti accusantium nulla illo provident et nesciunt, nisi laudantium iusto animi rem repudiandae, asperiores consequuntur Lorem ipsum dolor sit amet, consectetur adipisicing elit. Doloremque ipsam deserunt quaerat corrupti nihil error amet libero. Dignissimos, dolorem laudantium optio id, blanditiis eveniet repellendus pariatur neque facilis reprehenderit excepturi! Lorem ipsum dolor sit amet consectetur, adipisicing elit. Non repellendus, praesentium quasi quidem itaque numquam qui ex ducimus harum, perferendis officia deserunt libero magni assumenda mollitia aut ratione ipsam illo! Lorem ipsum dolor sit amet consectetur adipisicing elit. Maxime, corporis suscipit, natus odio nobis vel totam atque vitae porro animi in, cupiditate mollitia pariatur minus quos! Maiores assumenda explicabo expedita?" . "</p>"
        . "</div>";
      }
      echo "</div>";
    } else {
      echo "0 results";
    }
    $conn->close();
    ?>
    <div>

</body>

</html>
