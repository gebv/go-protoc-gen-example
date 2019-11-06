package main

import (
	"log"

	_ "github.com/gebv/go-protoc-gen-example/api/bar"
	_ "github.com/gebv/go-protoc-gen-example/api/foo"
)

func main() {
	log.Println("ok")
}
