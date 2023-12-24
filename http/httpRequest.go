package main

import (
	"io"
	"net/http"
)

func main() {
	call, err := http.Get("https://www.google.com")
	if err != nil {
		panic(err)
	}

	result, err := io.ReadAll(call.Body)
	if err != nil {
		panic(err)
	}

	print(string(result))

	call.Body.Close()
}
