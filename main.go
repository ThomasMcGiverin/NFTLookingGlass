package main

import (
	"NFTLookingGlass/models/nft"
	"NFTLookingGlass/pkg"
	//"NFTLookingGlass/server"
	"fmt"
)

func main() {
	//server.InitServer()
	res, err := pkg.OpenSeaClient.GetNFTInfo("0xA8eED2608fa8BD71C4c8aC540C420612146C1ffd")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.Assets)

	err = nft.InsertNft(res)
	if err != nil {
		fmt.Println(err)
	}
}
