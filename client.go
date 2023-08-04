package client

import (
	"errors"
	"fmt"
	"time"

	resty "github.com/go-resty/resty/v2"
)

// Client is an NKN OpenAPI client
type Client interface {
	// Addressbook
	GetRegisteredNames() (*ResponseGetRegisteredName, error)
	GetRegisteredNamesByAddress(string) ([]*ResponseGetRegisteredNameByAddress, error)
	GetAddressByRegisteredName(string) (*ResponseGetRegisteredNameByAddress, error)

	// Addresses
	GetAddresses() (*ResponseGetAddresses, error)
	GetSingleAddress(string) (*ResponseGetSingleAddress, error)
	GetAddressTransactions(string) (*ResponseGetAddressTransaction, error)

	// Blocks
	GetAllBlocks() (*ResponseGetBlock, error)
	GetBlockByHeight(int) (*ResponseGetBlockByHeight, error)
	GetBlockByHash(string) (*ResponseGetBlockByHash, error)
	GetTransactionsByBlockHeight(int) (*ResponseGetTransaction, error)
	GetTransactionsByBlockHash(string) (*ResponseGetTransaction, error)

	// Sigchain
	GetAllSigchains() (*ResponseGetSigchain, error)

	// Transactions
	GetAllTransactions() (*ResponseGetAllTransactions, error)
	GetTransactionByHash(string) (*ResponseGetTransaction, error)

	// Statistics
	StatsBlocksPerDay() ([]*ResponseStatsBlocksPerDay, error)
	StatsTransactionsPerDay() ([]*ResponseStatsTransactionsPerDay, error)
	StatsSupplies() (*ResponseStatsSupply, error)

	SetDebug(debug bool)
	SetAddress(address string)

	Next(interface{}) error
}

// New returns a new Client
func New() Client {
	c := resty.New()
	c.EnableTrace()
	c.SetHeader("Content-Type", "application/json")
	c.SetHeader("Accept", "application/json")
	c.SetTimeout(time.Second * 10)
	c.SetRetryCount(3)

	return &client{
		rest: c,
	}
}

type client struct {
	rest *resty.Client
}

func (c *client) SetDebug(debug bool) {
	c.rest.Debug = debug
}

func (c *client) SetAddress(address string) {
	c.rest.SetHostURL(address)
}

// Helper function
func (c *client) Next(in interface{}) error {
	nextpage := ""

	switch x := in.(type) {
	case *ResponseGetRegisteredName:
		nextpage = x.NextPageUrl
		in = x
	case *ResponseGetAddresses:
		nextpage = x.Addresses.NextPageUrl
		in = x
	case *ResponseGetAddressTransaction:
		nextpage = x.NextPageUrl
		in = x
	case *ResponseGetBlock:
		nextpage = x.Blocks.NextPageUrl
		in = x
	case *ResponseGetSigchain:
		nextpage = x.MetaData.NextPageUrl
		in = x
	case *ResponseGetAllTransactions:
		nextpage = x.Transactions.NextPageUrl
		in = x
	default:
		return errors.New("Unknown method")
	}
	c.rest.SetHostURL(nextpage)
	_, err := c.rest.R().
		SetQueryParam("per_page", "1000").
		SetResult(in).
		Get("")
	if err != nil {
		return err
	}
	return nil
}
func (c *client) do(method string, out interface{}) error {
	_, err := c.rest.R().
		SetQueryParam("per_page", "1000").
		SetResult(out).
		Get(method)
	return err
}

// METHODS
// Addressbook
func (c *client) GetRegisteredNames() (out *ResponseGetRegisteredName, err error) {
	err = c.do("address-book", &out)
	return
}

func (c *client) GetRegisteredNamesByAddress(address string) (out []*ResponseGetRegisteredNameByAddress, err error) {
	err = c.do(fmt.Sprintf("address-book/address/%s", address), &out)
	return
}

func (c *client) GetAddressByRegisteredName(name string) (out *ResponseGetRegisteredNameByAddress, err error) {
	err = c.do(fmt.Sprintf("address-book/name/%s", name), &out)
	return
}

// Addresses
func (c *client) GetAddresses() (out *ResponseGetAddresses, err error) {
	err = c.do("addresses", &out)
	return
}

func (c *client) GetSingleAddress(address string) (out *ResponseGetSingleAddress, err error) {
	err = c.do(fmt.Sprintf("addresses/%s", address), &out)
	return
}

func (c *client) GetAddressTransactions(address string) (out *ResponseGetAddressTransaction, err error) {
	err = c.do(fmt.Sprintf("addresses/%s/transactions", address), &out)
	return
}

// Blocks
func (c *client) GetAllBlocks() (out *ResponseGetBlock, err error) {
	err = c.do("blocks", &out)
	return
}

func (c *client) GetBlockByHeight(height int) (out *ResponseGetBlockByHeight, err error) {
	err = c.do(fmt.Sprintf("blocks/%d", height), &out)
	return
}

func (c *client) GetBlockByHash(hash string) (out *ResponseGetBlockByHash, err error) {
	err = c.do(fmt.Sprintf("blocks/%s", hash), &out)
	return
}

func (c *client) GetTransactionsByBlockHeight(height int) (out *ResponseGetTransaction, err error) {
	err = c.do(fmt.Sprintf("blocks/%d/transactions", height), &out)
	return
}

func (c *client) GetTransactionsByBlockHash(hash string) (out *ResponseGetTransaction, err error) {
	fmt.Println(fmt.Sprintf("blocks/%s/transactions", hash))
	err = c.do(fmt.Sprintf("blocks/%s/transactions", hash), &out)
	return
}

// Sigchain
func (c *client) GetAllSigchains() (out *ResponseGetSigchain, err error) {
	err = c.do("sigchains", &out)
	return
}

// Transactions
func (c *client) GetAllTransactions() (out *ResponseGetAllTransactions, err error) {
	err = c.do("transactions", &out)
	return
}

func (c *client) GetTransactionByHash(hash string) (out *ResponseGetTransaction, err error) {
	err = c.do(fmt.Sprintf("transactions/%s", hash), &out)
	return
}

// Statistics
func (c *client) StatsBlocksPerDay() (out []*ResponseStatsBlocksPerDay, err error) {
	err = c.do("statistics/daily/blocks", &out)
	return
}

func (c *client) StatsTransactionsPerDay() (out []*ResponseStatsTransactionsPerDay, err error) {
	err = c.do("statistics/daily/transactions", &out)
	return
}

func (c *client) StatsSupplies() (out *ResponseStatsSupply, err error) {
	err = c.do("statistics/supply", &out)
	return
}
