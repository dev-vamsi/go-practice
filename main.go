package main

import (
	"encoding/json"
	"fmt"
)

type JSONObject struct {
	Name   string   `json:"name"`
	Age    int      `json:"age"`
	Skills []string `json:"skills"`
}

func main() {
	data := ReadFileContent("sample.json")
	var jsonData JSONObject
	json.Unmarshal(data, &jsonData)
	fmt.Println(jsonData)
}
