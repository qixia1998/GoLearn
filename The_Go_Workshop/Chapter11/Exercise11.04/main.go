package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"
	"log"
)

type UserClient struct {
	ID   string
	Name string
}

type TxClient struct {
	ID          string
	User        *UserClient
	AccountFrom string
	AccountTo   string
	Amount      float64
}

type UserServer struct {
	ID string
}

type TxServer struct {
	ID          string
	User        UserServer
	AccountFrom string
	AccountTo   string
	Amount      *float32
}

func main() {
	
	var net bytes.Buffer
	
	clientTx := &TxClient{
		ID: "123456789",
		User: &UserClient{
			ID:   "ABCDEF",
			Name: "James",
		},
		AccountFrom: "Bob",
		AccountTo:   "Jane",
		Amount:      9.99,
	}
	
	enc := gob.NewEncoder(&net)
	if err := enc.Encode(clientTx); err != nil {
		log.Fatal("error encoding: ", err)
	}
	serverTx, err := sendToServer(&net)
	if err != nil {
		log.Fatal("server error: ", err)
	}
	
	fmt.Printf("%#v\n", serverTx)
}

func sendToServer(net io.Reader) (*TxServer, error) {
	
	tx := &TxServer{}
	
	dec := gob.NewDecoder(net)
	err := dec.Decode(tx)
	return tx, err
}