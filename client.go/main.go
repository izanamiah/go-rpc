package main

import (
	"fmt"
	"log"
	"net/rpc"
)


type Item struct {
	Title string
	Body string
}

func main() {
	var reply Item
	var db []Item

	client, err := rpc.DialHTTP("tcp", "localhost:4040")
	if err != nil {
		log.Fatal("Connection error:", err)
	}

	a := Item{"First", "A first item"}
	b := Item{"Second", "A second item"}
	c := Item{"Third", "A third item"}

	client.Call("API.CreateItem", a, &reply)
	client.Call("API.CreateItem", b, &reply)
	client.Call("API.CreateItem", c, &reply)

	fmt.Println("Database: ", db)
}