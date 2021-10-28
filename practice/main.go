package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Maps")

	states := make(map[string]string)
	states["WA"] = "Washington"
	states["OR"] = "Oregan"
	states["CA"] = "California"
	fmt.Println(states)
	fmt.Println(states["WA"])

	delete(states, "OR")
	states["NY"] = "New York"
	fmt.Println(states)

	for k, v := range states {
		fmt.Printf("%v: %v\n", k, v)
	}

	var keys []string

	for k := range states {
		keys = append(keys, k)
	}

	fmt.Println(keys)
	sort.Strings(keys)

	for k := range keys {
		fmt.Println(states[keys[k]])
	}

}
