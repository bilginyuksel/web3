package service

import (
	"app/internal/model"
	"context"
	"log"

	"github.com/algorand/go-algorand-sdk/client/kmd"
	"github.com/algorand/go-algorand-sdk/client/v2/algod"
	"github.com/algorand/go-algorand-sdk/client/v2/indexer"
	"github.com/algorand/go-algorand-sdk/future"
)

// Transaction service implementation
type Transaction struct {
	algodClient   *algod.Client
	indexerClient *indexer.Client
	kmdClient     kmd.Client
}

// NewTransaction ...
func NewTransaction(algodClient *algod.Client, kmdClient kmd.Client, indexerClient *indexer.Client) *Transaction {
	return &Transaction{
		algodClient:   algodClient,
		kmdClient:     kmdClient,
		indexerClient: indexerClient,
	}
}

// Transfer data structure to transfer assets
// Transfer assets from an account to another account
func (t *Transaction) Transfer(ctx context.Context, transfer model.Transfer) (string, error) {
	wHandle, err := t.kmdClient.InitWalletHandle(transfer.Wallet, "" /*pass*/)
	if err != nil {
		return "", err
	}

	walletToken := wHandle.WalletHandleToken

	txParams, err := t.algodClient.SuggestedParams().Do(ctx)
	if err != nil {
		return "", err
	}
	log.Println(txParams)

	txNote := []byte("test transaction")
	tx, err := future.MakePaymentTxn(transfer.Sender, transfer.Receiver, transfer.Amount, txNote, "" /*closeremainderto*/, txParams)
	if err != nil {
		return "", err
	}

	signedTx, err := t.kmdClient.SignTransaction(walletToken, "" /*pass*/, tx)
	if err != nil {
		return "", err
	}

	rawTxID, err := t.algodClient.SendRawTransaction(signedTx.SignedTransaction).Do(ctx)
	return rawTxID, err
}
