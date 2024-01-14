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
        <form enctype="multipart/form-data" action="./handle-upload.php" method="post">
            <input type="file" name="myFile" />
            <input type="submit" value="upload" />
        </form>
    </div>
</body>

</html>