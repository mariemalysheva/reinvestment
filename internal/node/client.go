package node

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	contractreinvestment "github.com/mariemalysheva/tokenized-reinvestment/contracts/reinvestment"
	"github.com/mariemalysheva/tokenized-reinvestment/internal/models"
	"math/big"
	"time"
)

type Client struct {
	chainID              *big.Int
	nodeClient           *ethclient.Client
	reinvestmentAddress  common.Address
	reinvestmentContract *contractreinvestment.ContractReinvestment
}

func New(
	nodeClient *ethclient.Client,
	reinvestmentAddress common.Address,
	chainID *big.Int,
) *Client {
	return &Client{
		nodeClient:          nodeClient,
		reinvestmentAddress: reinvestmentAddress,
		chainID:             chainID,
	}
}

func (c *Client) GetNode() *ethclient.Client { return c.nodeClient }

func (c *Client) GetReinvestmentAddress() common.Address { return c.reinvestmentAddress }

func (c *Client) GetReinvestmentContract() (contr *contractreinvestment.ContractReinvestment, err error) {
	if c.reinvestmentContract != nil {
		return c.reinvestmentContract, nil
	}
	c.reinvestmentContract, err = contractreinvestment.NewContractReinvestment(c.reinvestmentAddress, c.nodeClient)
	return c.reinvestmentContract, nil
}

func (c *Client) GetChainID() *big.Int {
	return c.chainID
}

func (c *Client) TxOpts(priv *ecdsa.PrivateKey) (txOpts *bind.TransactOpts, err error) {
	return bind.NewKeyedTransactorWithChainID(priv, c.chainID)
}

func (c *Client) GetCurrentReinvestment() (models.Reinvestment, error) {
	reinv, err := c.GetReinvestmentContract()
	if err != nil {
		return models.Reinvestment{}, err
	}

	curReinv, err := reinv.ReinvestmentPeriod(&bind.CallOpts{})
	if err != nil {
		return models.Reinvestment{}, err
	}
	return models.Reinvestment{
		ID:    curReinv.ID.Int64(),
		Asset: curReinv.CurrentAsset,
		Rate:  float64(curReinv.Rate.Int64()),
		Price: curReinv.AssetPrice.Int64(),
		Start: time.Unix(curReinv.Start.Int64(), 0),
		End:   time.Unix(0, 0),
	}, nil
}
