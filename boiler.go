package main

import (
	"log"
	"os"
)

func main() {

	html := `<!DOCTYPE html>
<html lang="en">
  <head>
  <title></title>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <link href="css/style.css" rel="stylesheet">
  <script src="js/script.js" defer></script>
  </head>
  <body>

  </body>
</html>`

	css := `html,body {
	  background-color: #0f0f0f;
	  color: #e3e3e3;
	  box-sizing: border-box;
	  margin: 0;
	  padding: 0;
	}`

	js := `console.log("Hello World!");`

	err := os.WriteFile("index.html", []byte(html), 0644)

	err = os.Mkdir("css", 0755)
	err = os.WriteFile("css/style.css", []byte(css), 0644)

	err = os.Mkdir("js", 0755)
	err = os.WriteFile("js/script.js", []byte(js), 0644)

	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}

}
