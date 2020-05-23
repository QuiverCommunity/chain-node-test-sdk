package utils

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
)

func GetLocalKey(key string) (string, string, error) {
	addressBytes, cmdLog, err := RunCli([]string{"keys", "show", key, "-a"}, "")
	address := strings.Trim(string(addressBytes), "\n ")

	return address, cmdLog, err
}

func QueryAccountByAddress(address string) (auth.BaseAccount, string, error) {
	var account auth.BaseAccount

	accJSONBytes, cmdLog, queryErr := RunCli([]string{"query", "account", address}, "")
	if queryErr != nil {
		return account, cmdLog, queryErr
	}
	codecErr := MakeCodec().UnmarshalJSON(accJSONBytes, &account)
	return account, cmdLog, codecErr
}

func QueryAccountByKey(key string) (auth.BaseAccount, string, error) {
	var account auth.BaseAccount
	address, cmdLog, err := GetLocalKey(key)
	if err != nil {
		return account, cmdLog, err
	}
	return QueryAccountByAddress(address)
}

func QueryNodeStatus() (*ctypes.ResultStatus, error) {
	var nodeStatus ctypes.ResultStatus
	nodeStatusBytes, _, queryErr := RunCli([]string{"status"}, "")

	if queryErr != nil {
		return nil, queryErr
	}

	codecErr := MakeCodec().UnmarshalJSON(nodeStatusBytes, &nodeStatus)
	return &nodeStatus, codecErr
}

func QueryRawTxResponse(txhash string) (sdk.TxResponse, error) {
	var txResponse sdk.TxResponse
	txResponseBytes, _, queryErr := RunCli([]string{"query", "tx", txhash}, "")

	if queryErr != nil {
		return txResponse, queryErr
	}

	codecErr := MakeCodec().UnmarshalJSON([]byte(txResponseBytes), &txResponse)
	return txResponse, codecErr
}
