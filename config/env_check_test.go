package config_test

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckEnv(t *testing.T) {
	_, err := os.Stat("../.env")
	assert.False(t, errors.Is(err, os.ErrNotExist), "deve existir o arquivo de variáveis de ambiente .env")
}
