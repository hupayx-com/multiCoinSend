package cli

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hupayx-com/multiCoinSend/x/multicoinsend/types"
	"github.com/spf13/cobra"
)

const (
	flagMultiFile = "multi"
)

var _ = strconv.Itoa(0)

func CmdCoinSend() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "coin-send",
		Short: "Broadcast message coinSend",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			reciverFile, _ := cmd.Flags().GetString(flagMultiFile)
			if reciverFile == "" {
				return fmt.Errorf("must specify at least one of %s", flagMultiFile)
			}

			var receivers []types.Receiver

			if reciverFile != "" {
				receivers, err = ReadTxFile(reciverFile)
				if err != nil {
					return err
				}
			}

			msg := types.NewMsgCoinSend(
				clientCtx.GetFromAddress().String(),
				receivers,
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(flagMultiFile, "", "path to file containg receivers")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

type ReceiverData struct {
	Receivers []Receiver `json:"Receiver"`
}

type Receiver struct {
	ToAddr string `json:"string"`
	Amount string `json:"coins"`
}

func ReadTxFile(path string) ([]types.Receiver, error) {
	contents, err := ioutil.ReadFile(filepath.Clean(path))

	if err != nil {
		return nil, err
	}

	var data ReceiverData

	if err = json.Unmarshal(contents, &data); err != nil {
		return nil, err
	}

	var txList []types.Receiver

	for i, t := range data.Receivers {

		fmt.Println("-------------------")

		amount, err := sdk.ParseCoinsNormalized(t.Amount)
		if err != nil {
			return nil, fmt.Errorf("invalid period length of %s in period %d, length must be greater than 0", t.Amount, i)
		}

		tx := types.Receiver{To: t.ToAddr, Amount: amount}
		txList = append(txList, tx)
	}

	return txList, nil
}
