package main

import (
	"fmt"
	"os"

	tml "github.com/pelletier/go-toml"
)

type Config struct {
	Login struct {
		User     string
		Password string
	}
}

func main() {
	file, err := os.Open("example.config")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	cfg := Config{}
	dec := tml.NewDecoder(file)

	if err := dec.Decode(&cfg); err != nil {
		fmt.Printf("error: cannot decode file - %s\n", err)
	}

	fmt.Println(cfg)
}
