package model

// Transfer asset transfer data structure
type Transfer struct {
	Wallet   string `json:"wallet_id"`
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	Amount   uint64 `json:"amount"`
}
