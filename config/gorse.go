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

	url := fmt.Sprintf("http://%s:%d", cfg.GorseConfig.Host, cfg.GorseConfig.Port)
	gorse := client.NewGorseClient(url, "")
	return gorseDB{Gorse: gorse}
}
