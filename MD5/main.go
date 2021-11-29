package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	getMD5("nasa-00.log")

	filenames, err := getFileName("md5sum.txt")

	if err != nil {
		fmt.Println(err)
	}

	ch := make(chan cresult)
	for filename, md5 := range *filenames {
		go checkSum(filename, md5, ch)
	}

	for range *filenames {
		result := <-ch
		if result.correct {
			fmt.Printf("The signature of %s is correct\n", result.filename)
		} else {
			fmt.Printf("The signature of %s is incorrect\n", result.filename)
		}
	}

}

func getMD5(filename string) string {
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	hash := md5.New()
	_, err = io.Copy(hash, file)

	if err != nil {
		fmt.Println(err)
	}

	return fmt.Sprintf("%x", hash.Sum(nil))

}

func getFileName(filename string) (*map[string]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	result := map[string]string{}
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Fields(line)
		if len(tokens) != 2 {
			return nil, fmt.Errorf("this is a bad line")
		}
		result[tokens[1]] = tokens[0]
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return &result, nil

}

func checkSum(filename, md5 string, channel chan cresult) {
	result := cresult{filename, getMD5(filename) == md5}
	channel <- result
}

type cresult struct {
	filename string
	correct  bool
}
