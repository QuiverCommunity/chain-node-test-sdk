package utils

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TxSign(txFilePath string, signerKey string) (string, string, error) {
	log := ""
	args := []string{
		"tx", "sign", txFilePath,
		"--from", signerKey,
		"--chain-id", Config.Chain,
	}

	signedTxBytes, signLog, signErr := RunCli(args)
	log += signLog
	if signErr != nil {
		log += string(signedTxBytes)
		return "", log, signErr
	}

	signedTxFileName := fmt.Sprintf("signed_tx_%d.json", time.Now().Unix())
	tmpDir, err := ioutil.TempDir("", Config.Chain)
	if err != nil {
		return "", log, err
	}
	signedTxFilePath := filepath.Join(tmpDir, signedTxFileName)
	writeErr := WriteFile(signedTxFilePath, signedTxBytes)
	return signedTxFilePath, signLog, writeErr
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

func SendTxFromSignerKey(txFilePath string, signerKey string) (sdk.TxResponse, string, error) {
	signedTxFilePath, signLog, signErr := TxSign(txFilePath, signerKey)
	if signErr != nil {
		return sdk.TxResponse{}, signLog, signErr
	}
	txResponse, err := TxBroadcast(signedTxFilePath)
	return txResponse, signLog, err
}
