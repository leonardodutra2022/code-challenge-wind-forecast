package server_test

import (
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewServer(t *testing.T) {
	host := "localhost"
	port := "9000"
	timeout := time.Second * 3
	_, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
	assert.NotNil(t, err, "porta deve estar dosponível para o servidor da aplicação ser inicializado")
}
