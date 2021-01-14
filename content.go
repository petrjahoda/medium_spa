package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
)

type ContentDataInput struct {
	Content string
}

type ContentDataOutput struct {
	Content string
	Result  string
}

func getContent(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	fmt.Println("Parsing incoming data")
	var data ContentDataInput
	err := json.NewDecoder(request.Body).Decode(&data)
	if err != nil {
		fmt.Println("Error parsing data: " + err.Error())
		var responseData ContentDataOutput
		responseData.Result = "problem parsing data:" + err.Error()
		writer.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(writer).Encode(responseData)
		return
	}
	fmt.Println("Data parsed, requested content: " + data.Content)
	file, err := ioutil.ReadFile("html/" + data.Content + ".html")
	if err != nil {
		fmt.Println("Error reading file: " + err.Error())
		var responseData ContentDataOutput
		responseData.Result = "problem reading file:" + err.Error()
		writer.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(writer).Encode(responseData)
		return
	}
	fmt.Println("File successfully read: " + data.Content)
	var responseData ContentDataOutput
	responseData.Content = string(file)
	responseData.Result = "ok"
	writer.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(writer).Encode(responseData)
	return
}
