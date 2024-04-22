package main

import (
	"fmt"
	"time"
)

type Block struct {
	data map[string]interface{}
	hash string
	prevHash string
	timestamp time.Time 
	nonce int  
}

//Blockchain containing Blocks

func main() {
	
}