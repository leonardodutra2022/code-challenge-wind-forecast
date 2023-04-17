package service_test

import (
	"testing"

	"github.com/leonardodutra2022/code-challenge-wind-forecast/service"
	"github.com/stretchr/testify/assert"
)

func TestFindLastQueryApi(t *testing.T) {
	qrApi, err := service.FindLastQueryApi()
	assert.Nil(t, err, "deve ocorrer uma listagem sem erro")
	assert.Equal(t, uint(1), qrApi.ID, "deve ser equivalente o valor")
}
