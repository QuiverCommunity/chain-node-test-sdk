package story

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/QuiverCommunity/chain-node-test-sdk/test/utils"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type CheckBalanceParam struct {
	Key     string `json:"key"`
	Denom   string `json:"denom"`
	Balance int64  `json:"balance"`
}

func CheckBalance(paramf string) (string, error) {
	log := "started CheckBalance function\n"
	paramBytes, readFileErr := utils.ReadFile(paramf)
	// log += "\n"
	// log += string(paramBytes)
	if readFileErr != nil {
		log += fmt.Sprintf("failed reading file for %s", paramf)
		return log, readFileErr
	}
	param := CheckBalanceParam{}
	paramErr := json.Unmarshal(paramBytes, &param)
	if paramErr != nil {
		log += fmt.Sprintf("failed decoding parameter for %+v", paramErr)
		return log, paramErr
	}
	account, cmdLog, queryErr := utils.QueryAccountByKey(param.Key)
	log += cmdLog
	if queryErr != nil {
		log += fmt.Sprintf("failed querying account for %s", param.Key)
		return log, queryErr
	}
	denomAmount := account.Coins.AmountOf(param.Denom)
	if !denomAmount.Equal(sdk.NewInt(param.Balance)) {
		checkErrText := fmt.Sprintf("account %s does not have balance of %d for %s denom, but have %d", param.Key, param.Balance, param.Denom, denomAmount.Int64())
		log += "\n"
		log += checkErrText
		log += fmt.Sprintf("\n%+v", account)
		return log, errors.New(checkErrText)
	}
	log += fmt.Sprintf("checked that balance is ok for %s key", param.Key)
	return log, nil
}
