package pkg

import (
	"github.com/thomasmcgiverin/NFTLookingGlass/pkg/opensea"
	"github.com/thomasmcgiverin/NFTLookingGlass/server/config"
)

var OpenSeaClient *opensea.Client

func init() {
	OpenSeaClient = opensea.New(config.Cfg.OpenSeaURL, config.Cfg.OpenSeaGetTimeout)
}
