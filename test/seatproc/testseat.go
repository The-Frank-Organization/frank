package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	tools := []string{"submit", "project", "read"}
	data, err := json.Marshal(tools)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(string(data))
}
