package main

import (
	"log"

	ed "github.com/coincircle/go-etherdelta"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var contract *HorseDex

func main() {
	conn, err := ethclient.Dial("wss://kovan.infura.io/ws")
	if err != nil {
		log.Fatalf("Failed to init node: %v", err)
	}
	// Instantiate the contract and display its name
	contract, err = NewHorseDex(common.HexToAddress("0x6926aeca67f038fffa5b3fb0209cb9e4e2b23f10"), conn)
	if err != nil {
		log.Fatalf("Failed to init contract: %v", err)
	}
	watchNewOrders()

}

func watchNewOrders() {
	logs := make(chan *HorseDexOrderPlaced)
	var buyer []common.Address
	sub, err := contract.HorseDexFilterer.WatchOrderPlaced(nil, logs, buyer)
	if err != nil {
		log.Fatalf("Failed to subscribe to event: %v", err)
	}

	fetchEtherdeltaBook()

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case order := <-logs:
			log.Println("New order detected! %v", order.Buyer)
		}
	}
}

func fetchEtherdeltaBook() {
	service := ed.New(&ed.Options{
		ProviderURI: "wss://mainnet.infura.io/ws",
	})

	orders, err := service.GetOrderBook(&ed.GetOrderBookOpts{
		TokenAddress: "0x5B0751713b2527d7f002c0c4e2a37e1219610A6B",
	})

	if err != nil {
		panic(err)
	}

	log.Println(orders)
}
