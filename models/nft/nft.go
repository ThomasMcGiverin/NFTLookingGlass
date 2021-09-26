package nft

import (
	"github.com/thomasmcgiverin/NFTLookingGlass/models"
	"github.com/thomasmcgiverin/NFTLookingGlass/pkg/opensea"
	"time"
)

type Nft struct {
	OwnerAddress        *string    `json:"owner_address"`
	TokenID             *string    `json:"token_id"`
	Name                *string    `json:"name"`
	OwnerName           *string    `json:"owner_name"`
	Description         *string    `json:"description"`
	ImageURL            *string    `json:"image_url"`
	ImagePreviewURL     *string    `json:"image_preview_url"`
	ImageThumbnailURL   *string    `json:"image_thumbnail_url"`
	ContractAddress     *string    `json:"contract_address"`
	ContractName        *string    `json:"contract_name"`
	ContractSymbol      *string    `json:"contract_symbol"`
	ContractDescription *string    `json:"contract_description"`
	CreatedAt           *time.Time `json:"created_at"`
}

func SelectNft(owner string) ([]*Nft, error) {
	db, err := models.DB()
	if err != nil {
		return nil, err
	}

	result := make([]*Nft, 0)
	res := db.Table("nft").Where("owner_address = ?", owner).Find(&result)
	if res.Error != nil {
		return nil, res.Error
	}

	return result, nil
}

func InsertNft(osNfts *opensea.NFTResponse) error {
	db, err := models.DB()
	if err != nil {
		return err
	}

	flatNfts := make([]*Nft, 0)
	for _, osNft := range osNfts.Assets {
		flatNfts = append(flatNfts, &Nft{
			OwnerAddress:        osNft.OwnerData.Address,
			TokenID:             osNft.TokenID,
			Name:                osNft.Name,
			OwnerName:           osNft.OwnerData.UserData.Username,
			Description:         osNft.Description,
			ImageURL:            osNft.ImageURL,
			ImagePreviewURL:     osNft.ImagePreviewURL,
			ImageThumbnailURL:   osNft.ImageThumbnailURL,
			ContractAddress:     osNft.AssetContractData.Address,
			ContractName:        osNft.AssetContractData.Name,
			ContractSymbol:      osNft.AssetContractData.Symbol,
			ContractDescription: osNft.AssetContractData.Description,
		})
	}

	result := db.Table("nft").Select(
		"owner_address",
		"token_id",
		"name",
		"owner_name",
		"description",
		"image_url",
		"image_preview_url",
		"image_thumbnail_url",
		"contract_address",
		"contract_name",
		"contract_symbol",
		"contract_description").Create(&flatNfts)

	return result.Error
}

func DeleteNft(owner string) error {
	db, err := models.DB()
	if err != nil {
		return err
	}

	result := db.Table("nft").Delete(&Nft{}, &owner)
	return result.Error
}
