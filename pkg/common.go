package pkg

import (
	"NFTLookingGlass/pkg/opensea"
	"NFTLookingGlass/server/config"
)

var OpenSeaClient *opensea.Client

func init() {
	OpenSeaClient = opensea.New(config.Cfg.OpenSeaURL, config.Cfg.OpenSeaGetTimeout)
}
