package main

import (
	"fmt"
	"os"

	"log"

	"github.com/pkg/errors"
)

func main() {
	setupLogging()

	cfg, err := readConfigFile("path\\to\\config")

	if err != nil {
		log.Printf("Error: %+v", err)
	}

	fmt.Println(cfg)

	defer handlePanic()

	panic("PANIC!!")

}

type config struct {
}

func readConfigFile(path string) (*config, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, errors.Wrap(err, "cannot open configuration file")
	}

	defer file.Close()

	cfg := &config{}

	return cfg, nil

}

func setupLogging() {
	out, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_RDONLY, 0644)
	if err != nil {
		return
	}

	log.SetOutput(out)
}

func handlePanic() {
	err := recover()
	if err != nil {
		fmt.Printf("Panic has been handled: %s", err)
	}
}
