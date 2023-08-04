package commands

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/jedib0t/go-pretty/table"
	"github.com/jedib0t/go-pretty/text"
	"github.com/spf13/cobra"
)

// transactionCmd represents the wallet command
var blocksCmd = &cobra.Command{
	Use:   "blocks",
	Short: "list blocks of NKN blockchain",
	Long:  "List blocks of NKN blockchain",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runBlocks()
	},
}

var (
	height int
	hash   string
)

func init() {
	rootCmd.AddCommand(blocksCmd)

	blocksCmd.Flags().IntVarP(&height, "height", "e", 0, "Fetch block with given block height")
	blocksCmd.Flags().StringVarP(&hash, "hash", "a", "", "Fetch block with given block hash. Special case is 'latest' which will return the lastest mined block.")
}

func runBlocks() error {
	if height == 0 && len(hash) == 0 {
		return errors.New("Need either height or hash parameter.")
	}
	t := table.NewWriter()
	// t.SetStyle(table.StyleRounded)
	mw = io.MultiWriter(os.Stdout)
	t.SetOutputMirror(mw)
	t.SetAlign([]text.Align{text.AlignCenter, text.AlignCenter})
	t.AppendHeader(table.Row{"created at", "block height", "block hash", "size", "#txn", "miner wallet", "benificiary"})

	if height > 0 {
		resp, err := c.GetBlockByHeight(height)
		if err != nil {
			return err
		}
		if resp == nil {
			return nil
		}
		t.AppendRow(table.Row{resp.Header.CreatedAt, resp.Header.Height, resp.Hash, resp.Size, resp.TransactionsCount, resp.Header.Wallet, resp.Header.BenificiaryWallet})
		t.Render()
	}
	if len(hash) > 0 {
		if hash == "latest" {
			fmt.Println("Fetching latest block...")
			resp, err := c.GetAllBlocks()
			if err != nil {
				return err
			}
			hash = resp.Blocks.Data[0].Hash
		}
		resp, err := c.GetBlockByHash(hash)
		if err != nil {
			return err
		}
		if resp == nil {
			return nil
		}
		t.AppendRow(table.Row{resp.Header.CreatedAt, resp.Header.Height, resp.Hash, resp.Size, resp.TransactionsCount, resp.Header.Wallet, resp.Header.BenificiaryWallet})
		t.Render()
	}
	return nil
}
