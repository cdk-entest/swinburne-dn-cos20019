<html>
  <head>
    <style>
      :root {
        box-sizing: border-box;
      }

      body {
        background-color: antiquewhite;
      }

      .container {
        max-width: 800px;
        margin-left: auto;
        margin-right: auto;
      }

      .input-upload {
        width: 350px;
        max-width: 100%;
        color: #444;
        padding: 5px;
        background: #fff;
        border-radius: 10px;
        border: 1px solid #555;
      }

      .input-upload::-webkit-file-upload-button {
        margin-right: 20px;
        border: none;
        background: #084cdf;
        padding: 10px 20px;
        border-radius: 10px;
        color: #fff;
        cursor: pointer;
        transition: background 0.2s ease-in-out;
      }

      .input-submit {
        background-color: #084cdf;
        color: white;
        border: none;
        padding: 10px 30px;
        border-radius: 10px;
        cursor: pointer;
        margin-top: 10px;
      }

      .drop-container {
        position: relative;
        display: flex;
        gap: 10px;
        flex-direction: column;
        justify-content: center;
        align-items: center;
        padding: 20px;
        height: 200px;
        border: 2px dashed #555;
        color: #444;
        cursor: pointer;
        transition: background 0.2s ease-in-out, border 0.2s ease-in-out;
      }
      .drop-container:hover {
        background: #eee;
        border-color: #111;
      }

      .drop-container:hover .drop-title {
        color: #222;
      }

      .drop-title {
        color: #444;
        font-size: 20px;
        font-weight: bold;
        text-align: center;
        transition: color 0.2s ease-in-out;
      }
    </style>
  </head>
  <body>
    <div class="container">
      <div>
        <form enctype="multipart/form-data" action="./handle-upload.php" method="post">
          <label for="file" class="drop-container" id="dropcontainer">
            <input type="file" class="input-upload" id="file" name="myFile"/>
          </label>
          <input type="submit" class="input-submit" />
        </form>
      </div>
    </div>
  </body>
  <script></script>
</html>
