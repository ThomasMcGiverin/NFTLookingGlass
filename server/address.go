package server

import (
	"fmt"
	"github.com/thomasmcgiverin/NFTLookingGlass/internal/nft"
	"github.com/thomasmcgiverin/NFTLookingGlass/pkg"
	"github.com/thomasmcgiverin/NFTLookingGlass/util"
	"net/http"
	"strings"
)

func getNftHandler(w http.ResponseWriter, r *http.Request) {
	owner := util.GetURLParam(r, "owner")
	owner = strings.ToLower(owner)
	refresh := util.GetURLBoolParam(r, "refresh")

	if refresh {
		// Get NFT data from opensea
		res, err := pkg.OpenSeaClient.GetNFTInfo(owner)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Delete old NFT entries owned by owner address
		err = nft.DeleteNft(owner)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Insert NFTs from OpenSea
		err = nft.CreateNft(owner, res)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	// Get NFTs owned by owner from the database
	nfts, err := nft.ListNft(owner)
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(nfts) == 0 && !refresh {
		// There are no nfts of the owner in the database so pull from OpenSea, otherwise owner has no nfts
		// Get NFT data from opensea
		res, err := pkg.OpenSeaClient.GetNFTInfo(owner)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Insert NFTs from OpenSea
		err = nft.CreateNft(owner, res)
		if err != nil {
			fmt.Println(err)
			return
		}

		nfts, err = nft.ListNft(owner)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	util.ServeJson(w, nfts)
}
