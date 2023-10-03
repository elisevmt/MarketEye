package utils

import (
	"encoding/json"
	"fmt"
)

func PrintStructJSON(value interface{}) {
	result, err := json.MarshalIndent(value, "", "\t")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("\n" + string(result) + "\n")
}

func GetStructJSON(value interface{}) string {
	result, err := json.MarshalIndent(value, "", "\t")
	if err != nil {
		fmt.Println(err)
	}
	return string(result)
}