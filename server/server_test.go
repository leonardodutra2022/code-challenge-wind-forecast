package server_test

import (
	"encoding/json"
	"io"
	"net"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/caarlos0/env"
	"github.com/leonardodutra2022/code-challenge-wind-forecast/data/config"
	"github.com/leonardodutra2022/code-challenge-wind-forecast/server"
	"github.com/stretchr/testify/assert"
)

type Server struct {
	port string
}

type CheckHealth struct {
	Check string `json:"check"`
}

func TestNewServerPort(t *testing.T) {
	host := "localhost"
	port := "9000"
	timeout := time.Second * 3
	_, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
	assert.NotNil(t, err, "porta deve estar dosponível para o servidor da aplicação ser inicializado")
}

func TestNewServer(t *testing.T) {
	srvStruct := server.NewServer()
	srvTest := Server{
		port: "9000",
	}
	assert.Equal(t, srvTest.port, srvStruct.Port, "deve ser idêntico o valor da porta do servidor")
}

func TestRun(t *testing.T) {
	cfg := config.Config{}
	env.Parse(&cfg)
	run := server.NewServer()
	go run.Run()
	var respCheck CheckHealth
	resp, err := http.Get(strings.Join([]string{"http://", cfg.HostApi, ":", cfg.PortApi, "/api/checkHealth"}, ""))
	assert.Nil(t, err, "deve realizar requisição sem erro")
	body, err := io.ReadAll(resp.Body)
	assert.Nil(t, err, "deve efetuar leitura do body da requisição sem erro")
	err = json.Unmarshal(body, &respCheck)
	assert.Nil(t, err, "deve realizar conversão do tipo do body recebido sem erro")
	assert.Equal(t, "ok", respCheck.Check, "deve retornar mensagem OK do endpoint check indicando que o serviço está funcionando")
}
