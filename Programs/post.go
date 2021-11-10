package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
)

func main() {
	
	fmt.Println("Enter email: ")
	fmt.Scanf("%s", &email)
	
    resp, err := http.PostForm("https://httpbin.org/post", email)

    if err != nil {
        log.Fatal(err)
    }

    var res map[string]interface{}

    json.NewDecoder(resp.Body).Decode(&res)

    fmt.Println(res["form"])
}