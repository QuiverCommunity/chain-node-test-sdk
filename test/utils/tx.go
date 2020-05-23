package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"sync"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/utils"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authtxb "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"
)

func TxSign(txFilePath string, signerKey string) (string, error) {
	args := []string{
		"tx", "sign", txFilePath,
		"--from", signerKey,
		"--chain-id", cf.Config.Chain,
	}

	signedTxBytes, cmdLog, signErr = RunPylonsCli(args, "")
	if signErr != nil {
		return "", signErr
	}

	signedTxFileName := fmt.Sprintf("signed_tx_%d.json", time.Now().Unix())
	signedTxFilePath := filepath.Join(ioutil.TempDir("", cf.Config.Chain), signedTxFileName)
	writeErr := WriteFile(signedTxFilePath, signedTxBytes)
	return signedTxFilePath, writeErr
}

func TxBroadcast(signedTxFilePath string, msg sdk.Msg) (sdk.TxResponse, error) {
	txResponse := sdk.TxResponse{}
	args := []string{"tx", "broadcast", signedTxFilePath}
	castOutputBytes, castErr := RunCli(args, "")

	if castErr != nil {
		return txResponse, castErr
	}

	codecErr = MakeCodec().UnmarshalJSON(castOutputBytes, &txResponse)
	return txResponse, codecErr
}

func SendTxFromSignerKey(txFilePath string, signerKey string) (sdk.TxResponse, error) {
	signedTxFilePath, signErr := TxSign(txFilePath, signerKey)
	if signErr != nil {
		return sdk.TxResponse{}, signErr
	}
	return TxBroadcast(t, execType.Sender.String(), true)
}
