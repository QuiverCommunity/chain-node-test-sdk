package utils

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TxSign(txFilePath string, signerKey string) (string, error) {
	args := []string{
		"tx", "sign", txFilePath,
		"--from", signerKey,
		"--chain-id", Config.Chain,
	}

	signedTxBytes, _, signErr := RunCli(args)
	if signErr != nil {
		return "", signErr
	}

	signedTxFileName := fmt.Sprintf("signed_tx_%d.json", time.Now().Unix())
	tmpDir, err := ioutil.TempDir("", Config.Chain)
	if err != nil {
		panic(err.Error())
	}
	signedTxFilePath := filepath.Join(tmpDir, signedTxFileName)
	writeErr := WriteFile(signedTxFilePath, signedTxBytes)
	return signedTxFilePath, writeErr
}

func TxBroadcast(signedTxFilePath string) (sdk.TxResponse, error) {
	txResponse := sdk.TxResponse{}
	args := []string{"tx", "broadcast", signedTxFilePath}
	castOutputBytes, _, castErr := RunCli(args)

	if castErr != nil {
		return txResponse, castErr
	}

	codecErr := MakeCodec().UnmarshalJSON(castOutputBytes, &txResponse)
	return txResponse, codecErr
}

func SendTxFromSignerKey(txFilePath string, signerKey string) (sdk.TxResponse, error) {
	signedTxFilePath, signErr := TxSign(txFilePath, signerKey)
	if signErr != nil {
		return sdk.TxResponse{}, signErr
	}
	return TxBroadcast(signedTxFilePath)
}
