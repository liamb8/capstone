package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
)

func main() {
	
	fmt.Println("Enter email address")

    resp, err := http.PostForm("https://httpbin.org/post", data)

    if err != nil {
        log.Fatal(err)
    }

    var res map[string]interface{}

    json.NewDecoder(resp.Body).Decode(&res)

    fmt.Println(res["form"])
}