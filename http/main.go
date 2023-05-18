package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

// The Body is of type io.ReadCloser
// ReadCloser is an intarface containing interfaces Reader and Closer

// Writer interface is used to transform []byte -> output
// Writer interface can read the Body

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	lw := logWriter{}

	// Copy needs a Writer (logWriter) and a Reader (Body)
	io.Copy(lw, resp.Body)

	// os.Stdout is File and implements Write, so it's a Writer
	// io.Copy(os.Stdout, resp.Body)
}

// To be Writer need to implement Write -> logWriter is a Writer
func (logWriter) Write(bs []byte) (int, error) {
	//Custom implementation
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
