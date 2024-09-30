// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	err := http.ListenAndServe(`:8080`, http.HandlerFunc(WriteHandle))
	if err != nil {
		panic(err)
	}
}

func WriteHandle(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "1")
	fmt.Fprint(w, "2")
	w.Write([]byte("3"))
}
