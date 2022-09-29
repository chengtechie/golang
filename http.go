package main

import (
	"fmt"
	"io"
	"net/http"
)

func getUrl(url string) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	// default implementation
	//io.Copy(os.Stdout, response.Body)

	// custom implementation
	lw := logWriter{}
	length, err := io.Copy(lw, response.Body)
	fmt.Println("Length is ", length)
}
