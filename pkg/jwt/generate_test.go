package jwt_test

import (
	"eulabs/products/pkg/jwt"
	"testing"
)

func Test_JWT(t *testing.T) {

	t.Run("gerar token", func(t *testing.T) {
		token, err := jwt.Generate("@eulabs#", "avelino.bego@gmail.com")
		if err != nil {
			t.Error(err)
		}
		t.Log(token)
	})
}
