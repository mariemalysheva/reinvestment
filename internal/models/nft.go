package models

import "github.com/ethereum/go-ethereum/common"

// NFTInputParams is an auto generated low-level Go binding around an user-defined struct.
type NFTInputParams struct {
	MaxTokenId           int64
	MaxNFTsPerOneTokenId int64
	Uri                  string
	Name                 string
	Symbol               string
}

type NFT struct {
	ID           uint64         `json:"id,omitempty"`
	Address      common.Address `json:"address"`
	Name         string         `json:"name"`
	Description  string         `json:"description"`
	Image        string         `json:"image"`
	AnimationUrl string         `json:"animation_url"`
	Title        string         `json:"title,omitempty"`
	Duration     int64          `json:"duration,omitempty"`
}
