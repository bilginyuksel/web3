package main

import (
	"app/internal/handler"
	"app/internal/service"
	"context"
	"strings"

	"github.com/algorand/go-algorand-sdk/client/kmd"
	"github.com/algorand/go-algorand-sdk/client/v2/algod"
	"github.com/algorand/go-algorand-sdk/client/v2/indexer"
	"github.com/algorand/go-algorand-sdk/future"
	"github.com/labstack/echo/v4"
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

	indexerClient, _ := indexer.MakeClient(_indexerAddress, _indexerToken)
	if err != nil {
		panic(err)
	}

	e := echo.New()

	transactionService := service.NewTransaction(algodClient, kmdClient, indexerClient)
	transactionHandler := handler.NewTransaction(transactionService)

	transactionHandler.RegisterRoutes(e)

	if err := e.Start(":8080"); err != nil {
		panic(err)
	}
}

var (
	receiver = "FJKB3ZN467HCFDIPD5S3DFXWFBXZ2EPJVOQLY3E5VUSCFFC4VNMIZI7PEE"
	sender   = "UGRUK2CXCHPXTMWPB4PEWCQFEVRSFBATBAEWR2ZURFC4HYVA2TSHV2MOZU"
	walletID = "7eec4313bd3d78352bf9a1d8234a0a3a"
)

// StartTransaction ...
func StartTransaction(algoClient *algod.Client, kmdClient kmd.Client) {
	walletResponse, err := kmdClient.ListWallets()
	if err != nil {
		panic(err)
	}

	defaultWallet := walletResponse.Wallets[0]
	pp.Println(defaultWallet)

	walletHandle, err := kmdClient.InitWalletHandle(defaultWallet.ID, "")
	if err != nil {
		panic(err)
	}

	walletHandleToken := walletHandle.WalletHandleToken

	txParams, err := algoClient.SuggestedParams().Do(context.Background())
	if err != nil {
		panic(err)
	}

	transaction, err := future.MakePaymentTxn(sender, receiver, 10, nil, "", txParams)
	if err != nil {
		panic(err)
	}

	signedTransaction, err := kmdClient.SignTransaction(walletHandleToken, "", transaction)
	if err != nil {
		panic(err)
	}

	rawTransaction, err := algoClient.SendRawTransaction(signedTransaction.SignedTransaction).Do(context.Background())
	if err != nil {
		panic(err)
	}

	pp.Println(rawTransaction)
}
