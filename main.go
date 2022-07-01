package main

import (
	"fmt"
	"os"
)

//Demo ...
type Demo struct {
}

func main() {

	var demo *Demo
	demo.Initialize(
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB_NAME"))

	fmt.Println("Connected!")
}
