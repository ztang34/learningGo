package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

var data = `
{
	"user": "Zhenguan",
	"type": "deposit",
	"amount": 10000.4
}`

type Request struct {
	Login  string  `json:"user"`
	Type   string  `json:"type"`
	Amount float64 `json:"amount"`
}

func main() {
	decoder := json.NewDecoder(strings.NewReader(data))
	out := &Request{}

	err := decoder.Decode(out)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Parsed json is : %+v\n", *out)
	}

	remainingFunds := 15000.5
	resp := map[string]interface{}{
		"completed":  true,
		"totalFunds": remainingFunds + out.Amount,
	}
	encoder := json.NewEncoder(os.Stdout)
	err = encoder.Encode(&resp)

	if err != nil {
		fmt.Println(err)
	}

}
