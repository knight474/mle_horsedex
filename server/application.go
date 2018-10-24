package main

import (
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var contract *HorseDex

func main() {
	conn, err := ethclient.Dial("https://mainnet.infura.io/76d846153845432cb5760b832c6bd0f0")
	if err != nil {
		log.Fatalf("Failed to init node: %v", err)
	}
	// Instantiate the contract and display its name
	contract, err = NewHorseDex(common.HexToAddress("0x6926aeca67f038fffa5b3fb0209cb9e4e2b23f10"), conn)
	if err != nil {
		log.Fatalf("Failed to init contract: %v", err)
	}
	go watchNewOrders()

}

func watchNewOrders() {
	logs := make(chan *HorseDexOrderPlaced)
	var buyer []common.Address
	_, err := contract.HorseDexFilterer.WatchOrderPlaced(nil, logs, buyer)
	if err != nil {
		log.Fatalf("Failed to subscribe to event: %v", err)
	}
	for {
		order := <-logs
		log.Println("New order detected! %v", order.Buyer)

	}
}
