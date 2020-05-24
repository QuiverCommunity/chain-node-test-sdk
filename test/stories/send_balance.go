package story

import (
	"encoding/json"

	"github.com/QuiverCommunity/chain-node-test-sdk/test/utils"
)

type SendBalanceParam struct {
	FromKey string `json:"from_key"`
	TxFile  string `json:"tx_file"`
}

func SendBalance(paramf string) (string, error) {
	paramBytes, readFileErr := utils.ReadFile(paramf)
	if readFileErr != nil {
		return "", readFileErr
	}
	param := SendBalanceParam{}
	paramErr := json.Unmarshal(paramBytes, &param)
	if paramErr != nil {
		return "", paramErr
	}
	txResponse, txErr := utils.SendTxFromSignerKey(param.TxFile, param.FromKey)
	logBytes, _ := json.Marshal(txResponse)
	logText := string(logBytes)
	if txErr != nil {
		return logText, txErr
	}
	return logText, nil
}
