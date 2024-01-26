<html>

<head>
  <title>Photo Lookup Page</title>
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
      cursor: pointer;
    }
  </style>
</head>

<body>
  <div class="container">
    <h1>Photo Lookup</h1>
    <form class="form-upload" method="post" action="./handle-lookup.php">
      <div class="form-grid">
        <div>
          <label for="title">Photo Title</label>
          <input type="text" id="title" name="title" class="input-title" />
        </div>
        <div>
          <label for="keyword">Keywords</label>
          <input type="text" id="keyword" name="keyword" class="input-keyword" />
        </div>
        <div>
          <label for="from-date">From Date</label>
          <input type="date" id="from-date" name="from-date" class="input-date" />
        </div>
        <div>
          <label for="to-date">To Date</label>
          <input type="date" id="to-date" name="to-date" class="input-date" />
        </div>
        <button type="submit" class="button-submit">Search</button>
      </div>
    </form>
    <a href="/photouploader.php">Photo Uploader</a>
  </div>
</body>

</html>