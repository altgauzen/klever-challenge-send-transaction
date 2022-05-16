package models

type Detail struct {
	// Id int           `json:"id"`
	Address string   `json:"address"`
	Balance string   `json:"balance"`
	TotalTx int64    `json:"totalTx"`
	// Balance          `json:"balance"`
	// Total            `json:"total"`
}

var Details []Detail


type Balance struct {
	Confirmed string   `json:"confirmed"`
	Unconfirmed string   `json:"unconfirmed"`
}

var Balances []Balance


type Utxo struct {
	Txid          string  `json:"txid"`
	Amount        string  `json:"amount"`
}


type Transaction struct {
	Address        string             `json:"address"`
	Block          int                `json:"block"`
	Txid           string             `json:"txid"`
	Value          string             `json:"value"`
}

var Transactions []Transaction
