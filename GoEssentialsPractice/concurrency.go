package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	/** goroutine
	urls := []string{
		"https://golang.org",
		"https://api.github.com",
		"https://httpbin.org/xml",
	}

	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			returnType(url)
			wg.Done()

		}(url)
	}

	wg.Wait()**/

	/**channel
	ch := make(chan int)
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("Sending: ", i)
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch)

	}()

	for i := range ch {
		fmt.Println("Got: ", i)
	}**/

	ch1, ch2 := make(chan int), make(chan int)
	go func() {
		ch1 <- 42
	}()

	select {
	case val := <-ch1:
		fmt.Printf("The number is %d", val)
	case val := <-ch2:
		fmt.Printf("The number is %d from ch2", val)
	}

	out := make(chan float64)
	go func() {
		time.Sleep(1 * time.Second)
		out <- 3.14
	}()

	select {
	case val2 := <-out:
		fmt.Println("Got result from out channel", val2)
	case <-time.After(2 * time.Second):
		fmt.Println("time out!")
	}

}

func returnType(url string) {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	ctype := resp.Header.Get("content-type")

	fmt.Printf("The content type of %s is %s", url, ctype)
}
