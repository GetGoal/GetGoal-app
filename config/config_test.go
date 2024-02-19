package config_test

import (
	"testing"

	"github.com/xbklyn/getgoal-app/config"
)

func TestReadConfig(t *testing.T) {
	t.Run("ReadConfig_LocalConfigExists", func(t *testing.T) {
		cfg := config.ReadConfig("../")

		if &cfg == nil {
			t.Errorf("ReadConfig() returned nil config")
		}

		if cfg.App.Port <= 0 {
			t.Errorf("Invalid App Port: %d", cfg.App.Port)
		}

		if cfg.Db.Host == "" {
			t.Errorf("Db Host is empty")
		}

		if cfg.Search.LabelLimit <= 0 {
			t.Errorf("Invalid Search LabelLimit: %d", cfg.Search.LabelLimit)
		}

	})

	t.Run("ReadConfig_LocalConfigNotExists", func(t *testing.T) {

		cfg := config.ReadConfig("../../")

		if &cfg == nil {
			t.Errorf("ReadConfig() returned nil config")
		}

		if cfg.App.Port <= 0 {
			t.Errorf("Invalid App Port: %d", cfg.App.Port)
		}

		if cfg.Db.Host == "" {
			t.Errorf("Db Host is empty")
		}

		if cfg.Search.LabelLimit <= 0 {
			t.Errorf("Invalid Search LabelLimit: %d", cfg.Search.LabelLimit)
		}

	})

	t.Run("ReadConfig_ConfigFilesNotExist", func(t *testing.T) {

		cfg := config.ReadConfig("../../")

		if &cfg == nil {
			t.Errorf("ReadConfig() returned nil config")
		}
	})
}
