package config

import (
	"fmt"

	"github.com/zhenghaoz/gorse/client"
)

type gorseDB struct {
	Gorse *client.GorseClient
}

func (g *gorseDB) GetGorseClient() *client.GorseClient {
	return g.Gorse

}

func NewGorseClient(cfg *Config) gorseDB {

	var url string

	switch cfg.Env {
	case "dev":
		url = fmt.Sprintf("http://%s:%d", cfg.DevGorseConfig.Host, cfg.DevGorseConfig.Port)
	case "qa":
		url = fmt.Sprintf("http://%s:%d", cfg.QaGorseConfig.Host, cfg.QaGorseConfig.Port)
	case "prod":
		url = fmt.Sprintf("http://%s:%d", cfg.ProdGorseConfig.Host, cfg.ProdGorseConfig.Port)
	default:
		url = fmt.Sprintf("http://%s:%d", cfg.DevGorseConfig.Host, cfg.DevGorseConfig.Port)
	}

	gorse := client.NewGorseClient(url, "")
	return gorseDB{Gorse: gorse}
}
