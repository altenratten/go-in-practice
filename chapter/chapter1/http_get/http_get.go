package main

import (
	"fmt"
	"net/http"
	"io"

)

func main() {
	resp, _ := http.Get("http://example.com/")
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
	resp.Body.Close()
}
