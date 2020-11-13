package main

import (
	"net/http"

	"server-master/services"

	"github.com/sirupsen/logrus"
)

func main() {
	http.HandleFunc("/example/get", services.ExampleGet)
	http.HandleFunc("/example/post", services.ExamplePost)

	logrus.Infof("server start succeeded , listen port :9000")
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		panic(err)
	}
}
