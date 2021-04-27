package main

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func main() {
	logrus.Info("Hello world")

	port := os.Getenv("PORT")

	if len(port) == 0 {
		logrus.Fatal("Порт не указан")
	}

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("This is a simple page"))
	})

	http.ListenAndServe(":" + port, nil)
}