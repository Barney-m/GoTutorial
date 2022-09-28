package main

import (
	"fmt"
	"io"
	"net/http"
)

type logWriter struct{}

func main() {
	res, err := http.Get("https://google.com")

	if err != nil {
		fmt.Printf("%v", err)
	}

	// Method 1
	// Initialize Byte Slice with 32 Bytes Capacity
	// byteSlice := make([]byte, 32*1024)

	// Read Response Body (Probably HTML) into byteSlice
	// res.Body.Read(byteSlice)

	// Print Out parsed byteSlice
	// fmt.Println(string(byteSlice))

	// Method 2
	// io.Copy(os.Stdout, res.Body)

	lw := logWriter{}
	io.Copy(lw, res.Body)
}

func (logWriter) Write(p []byte) (int, error) {
	fmt.Println(string(p))
	return len(p), nil
}
