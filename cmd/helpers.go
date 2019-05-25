package cmd

import (
	"encoding/json"
	"fmt"
)

// PrettyPrint print pretty output
func PrettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
	return
}
