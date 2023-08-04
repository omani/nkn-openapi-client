package client

import "github.com/nknorg/nkn/v2/common"

type MetaData struct {
	From         int         `json:"from,omitempty"`
	To           int         `json:"to,omitempty"`
	PerPage      interface{} `json:"per_page,omitempty"`
	CurrentPage  int         `json:"current_page,omitempty"`
	FirstPageURL string      `json:"first_page_url,omitempty"`
	PrevPageUrl  string      `json:"prev_page_url,omitempty"`
	NextPageUrl  string      `json:"next_page_url,omitempty"`
	Path         string      `json:"path,omitempty"`
}

func (m *MetaData) HasMore() bool {
	if len(m.NextPageUrl) == 0 {
		return false
	}
	return true
}

type ResponseGetRegisteredNameByAddress struct {
	Name      string `json:"name"`
	PublicKey string `json:"public_key"`
	Address   string `json:"address"`
	ExpiresAt string `json:"expires_at"`
}

// Addressbook
type ResponseGetRegisteredName struct {
	*MetaData
	Data []struct {
		Name      string `json:"name"`
		PublicKey string `json:"public_key"`
		Address   string `json:"address"`
		ExpiresAt string `json:"expires_at"`
	} `json:"data"`
}

type ResponseGetAddresses struct {
	Addresses struct {
		*MetaData
		Data []struct {
			Address            string `json:"address"`
			TransactionCount   int    `json:"count_transactions"`
			FirstTransactionAt string `json:"first_transaction"`
			LastTransactionAt  string `json:"last_transaction"`
			Balance            int    `json:"balance"`
		} `json:"data"`
	} `json:"addresses"`
	SumAddresses int `json:"sumAddresses"`
}

// Addresses
type ResponseGetSingleAddress struct {
	Address            string `json:"address"`
	TransactionCount   int    `json:"count_transactions"`
	FirstTransactionAt string `json:"first_transaction"`
	LastTransactionAt  string `json:"last_transaction"`
	Balance            int    `json:"balance"`
}

// Transactions

type ResponseGetAddressTransaction struct {
	*MetaData
	Data []struct {
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
	} `json:"data"`
}

type ResponseGetTransaction struct {
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

type ResponseGetAllTransactions struct {
	Transactions struct {
		*MetaData
		Data []struct {
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
		} `json:"data"`
	} `json:"transactions"`
	AvgSize         float64 `json:"avgSize"`
	SumTransactions int     `json:"sumTransactions"`
}

type TransactionPayload struct {
	ID                int            `json:"id"`
	TransactionID     int            `json:"transaction_id"`
	PayloadType       string         `json:"payloadType"`
	Sender            string         `json:"sender"`
	SenderWallet      string         `json:"senderWallet"`
	Recipient         string         `json:"recipient"`
	RecipientWallet   string         `json:"recipientWallet"`
	Amount            common.Fixed64 `json:"amount"`
	Submitter         string         `json:"submitter"`
	Registrant        string         `json:"registrant"`
	RegistrantWallet  string         `json:"registrantWallet"`
	Name              string         `json:"name"`
	Subscriber        string         `json:"subscriber"`
	Identifier        string         `json:"identifier"`
	Topic             string         `json:"topic"`
	Bucket            int            `json:"bucket"`
	Duration          int            `json:"duration"`
	Meta              string         `json:"meta"`
	PublicKey         string         `json:"public_key"`
	RegistrationFee   int            `json:"registration_fee"`
	Nonce             string         `json:"nonce"`
	TXNExpiration     int            `json:"txn_expiration"`
	Symbol            string         `json:"symbol"`
	TotalSupply       int            `json:"total_supply"`
	Precision         int            `json:"precision"`
	NanoPayExpiration int            `json:"nano_pay_expiration"`
	SignerPK          string         `json:"signerPk"`
	AddedAt           string         `json:"added_at"`
	CreatedAt         string         `json:"created_at"`
	GenerateWallet    string         `json:"generateWallet"`
	SubscriberWallet  string         `json:"subscriberWallet"`
	Sigchain          SigChain       `json:"sigchain"`
}

type ResponseGetBlockByHeight struct {
	ID                int    `json:"id"`
	Hash              string `json:"hash"`
	Size              int    `json:"size"`
	TransactionsCount int    `json:"transactions_count"`
	Header            *BlockHeader
}

type ResponseGetBlockByHash struct {
	ID                int    `json:"id"`
	Hash              string `json:"hash"`
	Size              int    `json:"size"`
	TransactionsCount int    `json:"transactions_count"`
	Header            *BlockHeader
}

// Blocks
type ResponseGetBlock struct {
	Blocks struct {
		*MetaData
		Data []struct {
			ID                int    `json:"id"`
			Hash              string `json:"hash"`
			Size              int    `json:"size"`
			TransactionsCount int    `json:"transactions_count"`
			Header            *BlockHeader
		} `json:"data"`
	} `json:"blocks"`
	AvgSize   string `json:"avgSize"`
	SumBlocks int    `json:"sumBlocks"`
}

type BlockHeader struct {
	Height            int    `json:"height"`
	SignerPK          string `json:"signerPk"`
	Wallet            string `json:"wallet"`
	BenificiaryWallet string `json:"benificiaryWallet"`
	CreatedAt         string `json:"created_at"`
}

type ResponseGetSigchain struct {
	*MetaData
	Data []struct {
		ID            int    `json:"id"`
		PayloadID     int    `json:"payload_id"`
		Nonce         int    `json:"nonce"`
		DataSize      int    `json:"dataSize"`
		BlockHash     string `json:"blockHash"`
		SrcID         string `json:"srcId"`
		SrcPubkey     string `json:"srcPubkey"`
		DestID        string `json:"destId"`
		DestPubkey    string `json:"destPubkey"`
		AddedAt       string `json:"added_at"`
		CreatedAt     string `json:"created_at"`
		SigchainElems []*SigchainElems
	} `json:"data"`
}

type SigChain struct {
	ID            int    `json:"id"`
	PayloadID     int    `json:"payload_id"`
	Nonce         int    `json:"nonce"`
	DataSize      int    `json:"dataSize"`
	BlockHash     string `json:"blockHash"`
	SrcID         string `json:"srcId"`
	SrcPubkey     string `json:"srcPubkey"`
	DestID        string `json:"destId"`
	DestPubkey    string `json:"destPubkey"`
	AddedAt       string `json:"added_at"`
	CreatedAt     string `json:"created_at"`
	SigchainElems []*SigchainElems
}

type SigchainElems struct {
	ID         int    `json:"id"`
	SigchainID int    `json:"sigchain_id"`
	ID2        string `json:"id2"`
	Pubkey     string `json:"pubkey"`
	Wallet     string `json:"wallet"`
	NextPubkey string `json:"nextPubkey"`
	Mining     bool   `json:"mining"`
	SignAlgo   string `json:"signAlgo"`
	Signature  string `json:"signature"`
	VRF        string `json:"vrf"`
	Proof      string `json:"proof"`
	AddedAt    string `json:"added_at"`
	CreatedAt  string `json:"created_at"`
}

// Stats
type Stats struct {
	Date  string `json:"date"`
	Count int    `json:"count"`
}
type ResponseStatsBlocksPerDay struct {
	Stats
}
type ResponseStatsTransactionsPerDay struct {
	Stats
}
type ResponseStatsSupply struct {
	MaxSupply         float64 `json:"max_supply"`
	TotalSupply       float64 `json:"total_supply"`
	CirculatingSupply float64 `json:"circulating_supply"`
}
