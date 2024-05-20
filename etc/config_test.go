package config_test

import (
	config "eulabs/products/etc"
	"testing"
)

func TestConfig(t *testing.T) {

	t.Run("teste obter config", func(t *testing.T) {
		conf, error := config.Load("config", "../")
		if error != nil {
			t.Error(error)
			return
		}
		t.Logf("Testando config %v\n", conf)
	})
}
