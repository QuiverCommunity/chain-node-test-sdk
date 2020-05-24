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
	paramBytes, readFileErr := utils.ReadFile(paramf)
	if readFileErr != nil {
		return "", readFileErr
	}
	param := CheckBalanceParam{}
	paramErr := json.Unmarshal(paramBytes, &param)
	if paramErr != nil {
		return "", paramErr
	}
	account, cmdLog, queryErr := utils.QueryAccountByKey(param.Key)
	if queryErr != nil {
		return cmdLog, queryErr
	}
	denomAmount := account.Coins.AmountOf(param.Denom)
	if !denomAmount.Equal(sdk.NewInt(param.Balance)) {
		return "", errors.New(fmt.Sprintf("account %s does not have balance of %d for %s denom, but have %d", param.Key, param.Balance, param.Denom, denomAmount))
	}
	return cmdLog, nil
}
