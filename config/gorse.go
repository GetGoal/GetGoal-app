package config

import "github.com/zhenghaoz/gorse/client"

type gorseDB struct {
	Gorse *client.GorseClient
}

func (g *gorseDB) GetGorseClient() *client.GorseClient {
	return g.Gorse

}

func NewGorseClient(cfg *Config) gorseDB {
	gorse := client.NewGorseClient("http://127.0.0.1:8088", "")
	return gorseDB{Gorse: gorse}
}
