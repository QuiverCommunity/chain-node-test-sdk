package story

import (
	"encoding/json"
	"fmt"

	"github.com/QuiverCommunity/chain-node-test-sdk/test/utils"
)

type SendBalanceParam struct {
	FromKey string `json:"from_key"`
	TxFile  string `json:"tx_file"`
}

func SendBalance(paramf string) (string, error) {
	log := "started SendBalance function\n"
	paramBytes, readFileErr := utils.ReadFile(paramf)
	if readFileErr != nil {
		log += string(paramBytes)
		log += fmt.Sprintf("failed for reading file %s\n", paramf)
		return log, readFileErr
	}
	param := SendBalanceParam{}
	paramErr := json.Unmarshal(paramBytes, &param)
	if paramErr != nil {
		log += fmt.Sprintf("failed parsing transaction for %s, paramErr %+v\n", paramf, paramErr)
		return log, paramErr
	}
	txResponse, txLog, txErr := utils.SendTxFromSignerKey(param.TxFile, param.FromKey)
	txResponseBytes, _ := json.Marshal(txResponse)
	log += string(txResponseBytes)
	log += txLog
	if txErr != nil {
		log += fmt.Sprintf("\nfailed sending transaction from %s txLog %s txErr %+v\n", param.FromKey, txLog, txErr)
		return log, txErr
	}
	log += fmt.Sprintf("sent transaction without any issue from %s", param.FromKey)
	return log, nil
}
