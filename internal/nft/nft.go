package nft

import (
	"fmt"
	"github.com/thomasmcgiverin/NFTLookingGlass/models/nft"
	"github.com/thomasmcgiverin/NFTLookingGlass/pkg/opensea"
	"github.com/thomasmcgiverin/NFTLookingGlass/util"
)

const (
	maxRetries = 5
)

func ListNft(owner string) ([]*nft.Nft, error) {
	return nft.SelectNft(owner)
}

func CreateNft(osNfts *opensea.NFTResponse) error {
	retries := 0
	inserted := false
	for !inserted && retries < maxRetries {
		err := nft.InsertNft(osNfts)
		if err != nil {
			fmt.Println(err)
			retries++
			continue
		}
		inserted = true
	}

	if !inserted {
		return util.ErrInsert
	}

	return nil
}

func DeleteNft(owner string) error {
	return nft.DeleteNft(owner)
}
