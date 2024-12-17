package admin

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/mariemalysheva/tokenized-reinvestment/internal/handler/wrapper"
	"testing"
)

func TestAddClients(t *testing.T) {
	var users []wrapper.UserReq

	i := 11
	for i <= 15 {
		users = append(users, wrapper.UserReq{
			FirstName: fmt.Sprintf("stepa %d", i),
			LastName:  fmt.Sprintf("masha %d", i),
			Savings:   int64((i + 1) * 150),
		})
		i++
	}

	data, err := json.Marshal(wrapper.PostAddUsersReq{Clients: users})
	if err != nil {
		t.Fatalf(err.Error())
	}

	fmt.Println(string(data))
}

func TestTransferClients(t *testing.T) {
	users := wrapper.PostTransferUsersReq{ToAddress: common.HexToAddress("0x0392C32f71B97899Be6C55AB5549706a63d6fb9B").Hex()}

	i := int64(5)
	for i < 16 {
		users.ClientIDs = append(users.ClientIDs, i)
		i++
	}

	data, err := json.Marshal(users)
	if err != nil {
		t.Fatalf(err.Error())
	}

	fmt.Println(string(data))
}
