package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func main() {

	err := killServer("path//to//file")

	if err != nil {
		fmt.Printf("Error: %+v", err)
	}
}

func killServer(pidfile string) error {
	content, err := ioutil.ReadFile(pidfile)

	if err != nil {
		return errors.Wrap(err, "cannot read from pidfile")
	}

	pid, err := strconv.Atoi(strings.TrimSpace(string(content)))

	if err != nil {
		return errors.Wrap(err, "file contain non-numerical value")
	}

	fmt.Printf("%d process has been killed", pid)

	return nil

}
