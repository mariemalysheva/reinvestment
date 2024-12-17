package cronjob

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	contract_reinvestment "github.com/mariemalysheva/tokenized-reinvestment/contracts/reinvestment"
	"github.com/mariemalysheva/tokenized-reinvestment/internal/models"
	"math/big"
	"testing"
)

const price = 1000

func TestAddUsers(t *testing.T) {
	node, err := ethclient.Dial("https://bsc-testnet.nodereal.io/v1/2c65af21a7c14fc4825f269763f31068")
	if err != nil {
		t.Fatalf(err.Error())
	}

	priv, err := crypto.HexToECDSA("bbba2f55818d7b7bf92fd0eef48af603118e421a3b56f802bc07899307166c69")
	if err != nil {
		t.Fatalf(err.Error())
	}
	txOpts, err := bind.NewKeyedTransactorWithChainID(priv, big.NewInt(97))
	if err != nil {
		t.Fatalf(err.Error())
	}
	//txOpts.GasPrice = big.NewInt(5500000000)

	reinvestmentManagerAddr := common.HexToAddress("0x0537f3190FEbbd3AEFb1A84B35DEFdeAC03B4981")
	//assetNftAddr := common.HexToAddress("0x71298276C62919871BA5453b71b902FaF2b813F3")

	reinv, err := contract_reinvestment.NewContractReinvestment(reinvestmentManagerAddr, node)
	if err != nil {
		t.Fatalf(err.Error())
	}

	//asset, err := contract_assetnft.NewContractAssetnft(assetNftAddr, node)
	//if err != nil {
	//	t.Fatalf(err.Error())
	//}

	var userIDs []*big.Int
	var salt []string
	var bals []*big.Int
	//var meta []contract_assetnft.AssetNFTAssetMetadata

	i := 120
	for i < 130 {
		user := models.Client{
			ID:        int64(i + 1),
			FirstName: fmt.Sprintf("first%d", i+1),
			LastName:  fmt.Sprintf("last%d", i+1),
		}
		user.Salt, err = user.Sign(priv, reinvestmentManagerAddr)
		if err != nil {
			t.Fatalf(err.Error())
		}

		bal := big.NewInt(int64(10000 + 100*i))

		//am := bal / price

		//m := contract_assetnft.AssetNFTAssetMetadata{
		//	UserID:         big.NewInt(int64(user.ID)),
		//	Salt:           user.Salt,
		//	Amount:         big.NewInt(am),
		//	SavingsBalance: big.NewInt(bal),
		//}

		userIDs = append(userIDs, big.NewInt(int64(user.ID)))
		salt = append(salt, user.Salt)
		//meta = append(meta, m)
		bals = append(bals, bal)
		i++
	}

	tx, err := reinv.AddUserBatch(txOpts, userIDs, salt, bals)
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Log(tx.Hash())
	//tx, err := asset.MintNFTBatch(txOpts, common.HexToAddress("0x61cA265FfD02Fb1af53cE8C4Eb205F0C9Cf978e8"), userIDs, meta)
	//if err != nil {
	//	t.Fatalf(err.Error())
	//}
	//t.Log(tx.Hash())

	//tx, err := reinv.ReinvestSavings(txOpts)
	//if err != nil {
	//	t.Fatalf(err.Error())
	//}

	//tx, err := reinv.Set(txOpts)
	//if err != nil {
	//	t.Fatalf(err.Error())
	//}

	//t.Log(tx.Hash())

}

//rAbi, err := ContractReinvestmentMetaData.GetAbi()
//if err != nil {
//return nil, err
//}
//
//input, err := rAbi.Pack("addUserBatch", userIDs_, salt_, balance_)
//if err != nil {
//return nil, err
//}
//
//fmt.Println(hexutil.Encode(input))
