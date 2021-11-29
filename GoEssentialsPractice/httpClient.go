package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Request struct {
	Name string  `json:"username"`
	GPA  float64 `json:"GPA"`
}

type User struct {
	Name        string `json:"login"`
	PublicRepos int    `json:"public_repos"`
}

func main() {
	httpGet()
	httpPost()

	user, err := userInfo("ztang34")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("The user is: %+v", *user)
	}
}

func httpGet() {
	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		fmt.Println("Cannot GET from the webpage")
	}

	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)

}

func httpPost() {
	payload := Request{"Zhenguan", 3.75}
	buf := bytes.NewBuffer([]byte{})
	encoder := json.NewEncoder(buf)
	if err := encoder.Encode(payload); err != nil {
		fmt.Println(err)
	}

	resp, err := http.Post("https://httpbin.org/post", "application/json", buf)
	if err != nil {
		fmt.Println()
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}

func userInfo(login string) (*User, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.github.com/users/%s", login))

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	user := &User{}

	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(user); err != nil {
		return nil, err
	}

	return user, nil
}
