package main

import (
	"errors"
	"fmt"
	"net/http"
)

func main() {
	urls := []string{"https://google.com", "https://facebook.com", "https://zhihu.com"}
	ch := make(chan contentType)

	for i, url := range urls {
		go func(url string, i int) {
			ch <- getContentType(url)
			if i == len(urls)-1 {
				close(ch)
			}
		}(url, i)

	}

	for content := range ch {
		if content.err != nil {
			fmt.Println(content.err)
		} else {
			fmt.Println(content.contentType)
		}
	}

}

func getContentType(url string) contentType {
	resp, err := http.Get(url)

	if err != nil {
		return contentType{"", err}
	}

	defer resp.Body.Close()

	header := resp.Header.Get("Content-Type")
	if header == "" {
		return contentType{"", errors.New("empty content type")}
	}

	return contentType{header, nil}

}

type contentType struct {
	contentType string
	err         error
}
