package main

import (
	"fmt"
	"net/http"
)

func (bs *BlockchainServer) HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", "Hello world!")
}
