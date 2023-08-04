package commands

import (
	"errors"
	"io"
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/table"
	"github.com/jedib0t/go-pretty/text"
	"github.com/nknorg/nkn/v2/common"
	"github.com/spf13/cobra"
)

// transactionCmd represents the wallet command
var transactionCmd = &cobra.Command{
	Use:   "transactions",
	Short: "list transactions for an address",
	Long:  "List transactions for an address",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runTransactions()
	},
}

var (
	nknaddress string
	txntype    string
	mw         io.Writer
)

func init() {
	rootCmd.AddCommand(transactionCmd)

	transactionCmd.Flags().StringVarP(&nknaddress, "address", "a", "", "Fetch transactions for given NKN address")
	transactionCmd.Flags().StringVar(&hash, "hash", "", "Fetch transaction for given tx hash")
	transactionCmd.Flags().StringVarP(&txntype, "type", "t", "transfer", "Filter transaction type [asset, reward, subscription]")
}

func runTransactions() error {
	if len(hash) == 0 && len(nknaddress) == 0 {
		return errors.New("Need either address or hash flag set.")
	}
	if len(nknaddress) > 0 {
		_, err := common.ToScriptHash(nknaddress)
		if err != nil {
			return err
		}
	}
	t := table.NewWriter()
	// t.SetStyle(table.StyleRounded)
	mw = io.MultiWriter(os.Stdout)
	t.SetOutputMirror(mw)
	t.SetAlign([]text.Align{text.AlignCenter, text.AlignCenter})

	if len(nknaddress) > 0 {
		resp, err := c.GetAddressTransactions(nknaddress)
		if err != nil {
			return err
		}
		if resp == nil {
			return nil
		}
		switch strings.ToLower(txntype) {
		case "transfer":
			t.AppendHeader(table.Row{"created at", "block height", "txn hash", "sender", "recipient", "amount"})
			loop := func() {
				for _, tx := range resp.Data {
					if tx.TxType != "TRANSFER_ASSET_TYPE" {
						continue
					}
					t.AppendRow(table.Row{tx.CreatedAt, tx.BlockHeight, tx.Hash, tx.Payload.SenderWallet, tx.Payload.RecipientWallet, tx.Payload.Amount})
				}
			}
			loop()
			cnt := 0
			for resp.HasMore() {
				err = c.Next(resp)
				if err != nil {
					return err
				}
				if resp == nil {
					return nil
				}
				loop()
				cnt++
				if cnt == 3 {
					break
				}
			}
		case "reward":
			t.AppendHeader(table.Row{"created at", "block height", "txn hash", "recipient", "reward"})
			loop := func() {
				for _, tx := range resp.Data {
					if tx.TxType != "COINBASE_TYPE" {
						continue
					}
					t.AppendRow(table.Row{tx.CreatedAt, tx.BlockHeight, tx.Hash, tx.Payload.RecipientWallet, tx.Payload.Amount})
				}
			}
			loop()
			cnt := 0
			for resp.HasMore() {
				err = c.Next(resp)
				if err != nil {
					return err
				}
				if resp == nil {
					return nil
				}
				loop()
				cnt++
				if cnt == 3 {
					break
				}
			}
		}
		t.Render()
	}
	if len(hash) > 0 {
		tx, err := c.GetTransactionByHash(hash)
		if err != nil {
			return err
		}
		if tx == nil {
			return nil
		}
		t.AppendHeader(table.Row{"created at", "block height", "txn hash", "sender", "recipient", "amount"})
		t.AppendRow(table.Row{tx.CreatedAt, tx.BlockHeight, tx.Hash, tx.Payload.RecipientWallet, tx.Payload.Recipient, tx.Payload.Amount})
		t.Render()
	}

	return nil
}
