package client

import (
	"fmt"
	"testing"
)

type exampleWallet struct {
	address string
	name    string
}
type exampleBlock struct {
	ID                int    `json:"id"`
	Hash              string `json:"hash"`
	Size              int    `json:"size"`
	TransactionsCount int    `json:"transactions_count"`
	Header            *BlockHeader
}

type exampleTXN struct {
	ID          int    `json:"id"`
	BlockID     int    `json:"block_id"`
	Attributes  string `json:"attributes"`
	Fee         int    `json:"fee"`
	Hash        string `json:"hash"`
	Nonce       string `json:"nonce"`
	TxType      string `json:"txType"`
	BlockHeight int    `json:"block_height"`
	CreatedAt   string `json:"created_at"`
	Payload     *TransactionPayload
}

var (
	wallet exampleWallet
	block  exampleBlock
	txn    exampleTXN
	c      Client
)

func getClient() Client {
	c := New()
	c.SetAddress("https://openapi.nkn.org/api/v1/")
	return c
}

func TestGetRegisteredNames(t *testing.T) {
	c := getClient()
	resp, err := c.GetRegisteredNames()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("[TestGetRegisteredNames] Address: %s - Name: %s\n", resp.Data[0].Address, resp.Data[0].Name)
	if len(resp.Data) > 0 {
		wallet.address = resp.Data[0].Address
		wallet.name = resp.Data[0].Name
	}
	if !resp.HasMore() {
		return
	}
	// now fetch next page
	err = c.Next(resp)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("[TestGetRegisteredNames] Address: %s - Name: %s\n", resp.Data[0].Address, resp.Data[0].Name)
	if len(resp.Data) > 0 {
		wallet.address = resp.Data[0].Address
		wallet.name = resp.Data[0].Name
	}
}

func TestGetRegisteredNamesByAddress(t *testing.T) {
	c := getClient()
	resp, err := c.GetRegisteredNamesByAddress(wallet.address)
	if err != nil {
		t.Fatal(err)
	}
	for _, r := range resp {
		fmt.Printf("[TestGetRegisteredNamesByAddress] Address: %s - Name: %s\n", r.Address, r.Name)
	}
}

func TestGetRegisteredNamesByName(t *testing.T) {
	c := getClient()
	resp, err := c.GetAddressByRegisteredName(wallet.name)
	if err != nil {
		t.Fatal(err)
	}
	if resp != nil {
		fmt.Printf("[TestGetRegisteredNamesByName] Address: %s - Name: %s\n", resp.Address, wallet.name)
	}
}

func TestGetAddresses(t *testing.T) {
	c := getClient()
	resp, err := c.GetAddresses()
	if err != nil {
		t.Error(err)
	}
	if resp == nil {
		t.Fatalf("Response has no data")
	}
	if len(resp.Addresses.Data) == 0 {
		t.Fatalf("Response has no data")
	}
	fmt.Printf("[TestGetAddresses] %#v\n", resp.Addresses.Data[0])

	if !resp.Addresses.HasMore() {
		return
	}
	// now fetch next page
	err = c.Next(resp)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("[TestGetAddresses] %#v\n", resp.Addresses.Data[0])
}

func TestGetSingleAddress(t *testing.T) {
	c := getClient()
	resp, err := c.GetSingleAddress(wallet.address)
	if err != nil {
		t.Error(err)
	}
	if resp == nil {
		t.Fatalf("Response has no data")
	}
	fmt.Printf("[TestGetAddresses] %#v\n", resp)
}

func TestGetAddressTransactions(t *testing.T) {
	c := getClient()
	resp, err := c.GetAddressTransactions(wallet.address)
	if err != nil {
		t.Error(err)
	}
	if resp == nil {
		t.Fatalf("Response has no data")
	}
	fmt.Printf("[TestGetAddressesTransactions] %#v\n", resp)
}

func TestGetAllBlocks(t *testing.T) {
	c := getClient()
	resp, err := c.GetAllBlocks()
	if err != nil {
		t.Error(err)
	}
	if resp == nil {
		t.Fatalf("Response has no data")
	}

	fmt.Println(len(resp.Blocks.Data))
	fmt.Printf("[TestGetAllBlocks] %#v\n", resp.Blocks.Data[0])
	block = resp.Blocks.Data[0]

	if !resp.Blocks.HasMore() {
		return
	}
	// now fetch next page
	err = c.Next(resp)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("[TestGetAllBlocks] %#v\n", resp.Blocks.Data[0])
	block = resp.Blocks.Data[0]
}

func TestGetBlockByHeight(t *testing.T) {
	c := getClient()
	resp, err := c.GetBlockByHeight(block.ID)
	if err != nil {
		t.Error(err)
	}
	if resp == nil {
		t.Fatalf("Response has no data")
	}
	fmt.Printf("[TestGetBlockByHeight] %#v\n", resp)
}

func TestGetBlockByHash(t *testing.T) {
	c := getClient()
	resp, err := c.GetBlockByHash(block.Hash)
	if err != nil {
		t.Error(err)
	}
	if resp == nil {
		t.Fatalf("Response has no data")
	}
	fmt.Printf("[TestGetBlockByHash] %#v\n", resp)
}

func TestGetTransactionsByBlockHeight(t *testing.T) {
	c := getClient()
	resp, err := c.GetTransactionsByBlockHeight(block.ID)
	if err != nil {
		t.Error(err)
	}
	if resp == nil {
		t.Fatalf("Response has no data")
	}
	fmt.Printf("[TestGetTransactionsByBlockHeight] %#v\n", resp)
}

// according to docs querying txn by block hash is possible. but it isn't.
// reported bug to rule110 team on discord https://discord.com/channels/443413382737952778/724296220222160946
// func TestGetTransactionsByBlockHash(t *testing.T) {
// c := getClient()
// 	resp, err := c.GetTransactionsByBlockHash(block.Hash)
// 	if err != nil {
// 		t.Error(err)
// 	}
//
// 	if resp == nil {
// 		t.Fatalf("Response has no data")
// 	}
//
// 	fmt.Printf("[TestGetTransactionsByBlockHash] %#v\n", resp)
// }

func TestGetAllSigchains(t *testing.T) {
	c := getClient()
	resp, err := c.GetAllSigchains()
	if err != nil {
		t.Error(err)
	}
	if resp == nil {
		t.Fatalf("Response has no data")
	}
	fmt.Printf("[TestGetAllSigchains] %#v\n", resp.Data[0])

	if !resp.MetaData.HasMore() {
		return
	}
	// now fetch next page
	err = c.Next(resp)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("[TestGetAllSigchains] %#v\n", resp.Data[0])
}

func TestGetAllTransactions(t *testing.T) {
	c := getClient()
	resp, err := c.GetAllTransactions()
	if err != nil {
		t.Error(err)
	}
	if resp == nil {
		t.Fatalf("Response has no data")
	}
	fmt.Printf("[TestGetAllTransactions] %#v\n", resp.Transactions.Data[0])
	txn = resp.Transactions.Data[0]

	// test c.Next() implementation
	// save a txn of first page
	tmp := resp.Transactions.Data[0].ID
	// check if next page is available
	if !resp.Transactions.HasMore() {
		return
	}
	// now fetch next page
	err = c.Next(resp)
	if err != nil {
		t.Fatal(err)
	}
	// check if this txn is the same as txn above
	// if NOT, c.Next() works
	if tmp == resp.Transactions.Data[0].ID {
		t.Fatal("EQUAL")
	}
	fmt.Printf("[TestGetAllTransactions] %#v\n", resp.Transactions.Data[0])
	txn = resp.Transactions.Data[0]
}

func TestGetTransactionByHash(t *testing.T) {
	c := getClient()
	resp, err := c.GetTransactionByHash(txn.Hash)
	if err != nil {
		t.Error(err)
	}
	if resp == nil {
		t.Fatalf("Response has no data")
	}
	fmt.Printf("[TestGetTransactionByHash] %#v\n", resp)
}

func TestStatsBlocksPerDay(t *testing.T) {
	c := getClient()
	resp, err := c.StatsBlocksPerDay()
	if err != nil {
		t.Error(err)
	}
	if resp == nil {
		t.Fatalf("Response has no data")
	}
	fmt.Printf("[TestStatsBlocksPerDay] %#v\n", resp)
}

func TestStatsTransactionsPerDay(t *testing.T) {
	c := getClient()
	resp, err := c.StatsTransactionsPerDay()
	if err != nil {
		t.Error(err)
	}
	if resp == nil {
		t.Fatalf("Response has no data")
	}
	fmt.Printf("[TestStatsTransactionsPerDay] %#v\n", resp)
}

func TestStatsSupplies(t *testing.T) {
	c := getClient()
	resp, err := c.StatsSupplies()
	if err != nil {
		t.Error(err)
	}

	if resp == nil {
		t.Fatalf("Response has no data")
	}
	fmt.Printf("[TestStatsSupplies] %#v\n", resp)
}
