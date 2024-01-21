<html>

<head>
    <title>Upload Page</title>
    <style>

        .container {
            max-width: 800px;
            margin-left: auto;
            margin-right: auto;
        }

        .form-grid {
            display: grid;
            row-gap: 10px;
            grid-template-columns: repeat(1, minxmax(0, 1fr));
        }

        .form-upload {
            padding: 2em 2em;
            border: solid 1px black;
        }

        .button-submit {
            max-width: 150px;
        }

    </style>
</head>

<body>
    <div class="container">
        <h1>Photo Uploader</h1>
        <form class="form-upload">
            <div class="form-grid">
                <div>
                    <label for="title">Photo Title</label>
                    <input type="text" id="title" class="input-title" />
                </div>
                <div>
                    <label for="file">Select a photo</label>
                    <input type="file" id="file" class="input-file" />
                </div>
                <div>
                    <label for="description">Description</label>
                    <input type="text" id="description" class="input-desc" />
                </div>
                <div>
                    <label for="date">Date</label>
                    <input type="date" id="date" class="input-date" />
                </div>
                <div>
                    <label for="keyword">Keywords(comman-delimited, e.g. keyword1, keyword2, ...)</label>
                    <input type="text" id="keyword" class="input-keyword" />
                </div>
                <button type="submit" class="button-submit">Upload</button>
            </div>
        </form>
        <a href="/photolookup.php">Photo Lookup</a>
    </div>
</body>

</html>