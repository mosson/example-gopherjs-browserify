package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/index.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "serve/index.js")
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<!DOCTYPE html>
      <html lang="ja">
      <head>
      <meta charset="utf-8">
      </head>
      <body>
			<script type="text/javascript" src="/index.js"></script>
      </body>
      </html>
      `))
	})

	log.Printf("HTTP Server starting at %s", ":8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
