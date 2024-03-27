package config

import "github.com/zhenghaoz/gorse/client"

type gorseDB struct {
	Gorse *client.GorseClient
}

func (g *gorseDB) GetGorseClient() *client.GorseClient {
	return g.Gorse

}

func NewGorseClient(cfg *Config) gorseDB {
	gorse := client.NewGorseClient("http://cp23ssa1.sit.kmutt.ac.th:8088", "")
	return gorseDB{Gorse: gorse}
}
