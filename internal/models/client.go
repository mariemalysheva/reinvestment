package models

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/itsahedge/go-cowswap/util/signature-scheme/eip712"
	"math/big"
	"time"
)

const (
	DomainName    = "Reinvestment"
	DomainVersion = "0"
	RateDenom     = 100
)

type Client struct {
	ID        int64     `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Savings   int64     `json:"savings"`
	Salt      string    `json:"salt"`
	CreatedAt time.Time `json:"created_at"`
}

type Reinvestment struct {
	ID    int64
	Asset common.Address
	Rate  float64
	Price int64
	Start time.Time
	End   time.Time
}

type ReinvestmentRecord struct {
	Reinvestment Reinvestment
	Savings      int64
	Amount       int64
}

func (u Client) GetTypedData(verifyingContract common.Address) apitypes.TypedData {

	var typesMap = apitypes.Types{
		"Client": []apitypes.Type{
			{Name: "clientId", Type: "uint256"},
			{Name: "firstName", Type: "string"},
			{Name: "lastName", Type: "string"},
		},
		"EIP712Domain": []apitypes.Type{
			{Name: "name", Type: "string"},
			{Name: "version", Type: "string"},
			{Name: "chainId", Type: "uint256"},
			{Name: "verifyingContract", Type: "address"},
		},
	}

	var typedDomain = apitypes.TypedDataDomain{
		Name:              DomainName,
		Version:           DomainVersion,
		ChainId:           math.NewHexOrDecimal256(0),
		VerifyingContract: verifyingContract.Hex(),
	}

	var typedDataMessage = apitypes.TypedDataMessage{
		"clientId":  big.NewInt(u.ID),
		"firstName": u.FirstName,
		"lastName":  u.LastName,
	}

	return apitypes.TypedData{
		Types:       typesMap,
		Domain:      typedDomain,
		Message:     typedDataMessage,
		PrimaryType: "Client",
	}
}

func (u Client) Sign(priv *ecdsa.PrivateKey, reinvestmentAddr common.Address) (string, error) {
	typedData := u.GetTypedData(reinvestmentAddr)

	sig, err := eip712.SignTypedData(typedData, priv)
	if err != nil {
		return "", err
	}

	return hexutil.Encode(sig), nil
}
