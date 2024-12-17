package container

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/mariemalysheva/tokenized-reinvestment/config"
	"github.com/mariemalysheva/tokenized-reinvestment/internal/node"
	"log"
)

func (c *Container) GetNode(ctx context.Context) (*node.Client, error) {
	if c.node != nil {
		return c.node, nil
	}

	nodeCli, err := ethclient.DialContext(ctx, config.RPC)
	if err != nil {
		return nil, err
	}

	chainID, err := nodeCli.ChainID(context.Background())
	if err != nil {
		log.Fatalln("error getting chain id from node client", err.Error())
	}

	c.node = node.New(nodeCli, common.HexToAddress(config.ReinvestmentAddr), chainID)

	return c.node, nil
}

func (c *Container) GetOwnerKey() (priv *ecdsa.PrivateKey, err error) {
	if c.key != nil {
		return c.key, nil
	}

	c.key, err = crypto.HexToECDSA(config.OwnerPriv)
	if err != nil {
		return nil, fmt.Errorf("can not get executor private key: %w", err)
	}

	return c.key, nil
}
