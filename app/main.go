package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/algorand/go-algorand-sdk/client/kmd"
	"github.com/algorand/go-algorand-sdk/client/v2/algod"
	"github.com/algorand/go-algorand-sdk/client/v2/indexer"
	"github.com/yudai/pp"
)

var (
	_algodAddress = "http://localhost:4001"
	_algodToken   = strings.Repeat("a", 64)

	_kmdAddress = "http://localhost:4002"
	_kmdToken   = strings.Repeat("a", 64)

	_indexerAddress = "http://localhost:8980"
	_indexerToken   = strings.Repeat("a", 64)
)

func main() {
	algodClient, err := algod.MakeClient(_algodAddress, _algodToken)
	if err != nil {
		panic(err)
	}

	kmdClient, err := kmd.MakeClient(_kmdAddress, _kmdToken)
	if err != nil {
		panic(err)
	}

	fmt.Printf("algod: %T, kmd: %T\n", algodClient, kmdClient)

	s, err := algodClient.Status().Do(context.Background())
	if err != nil {
		panic(err)
	}
	pp.Println(s)

	tot, txs, err := algodClient.PendingTransactions().Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println("total txs:", tot)
	pp.Println(txs)

	wallets, err := kmdClient.ListWallets()
	if err != nil {
		panic(err)
	}
	pp.Println(wallets)

	indxer, _ := indexer.MakeClient(_indexerAddress, _indexerToken)
	acc, err := indxer.SearchAccounts().Do(context.Background())
	if err != nil {
		panic(err)
	}
	pp.Println(acc)

}
