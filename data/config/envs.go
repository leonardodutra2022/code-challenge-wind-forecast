package config

/*
Estrutura de dados para configurações gerais utilizadas na aplicação, que serão oriundos de variáveis de ambientes
*/
type Config struct {
	HostDB             string  `env:"DB_HOST" envDefault:"localhost"`          // Host do banco de dados
	PortDB             int     `env:"DB_PORT" envDefault:"5432"`               // Porta banco de dados
	PassDB             string  `env:"DB_PASS" envDefault:"aksQmB93fA00"`       // Senha do banco de dados
	DatabaseName       string  `env:"DB_NAME" envDefault:"code_challenge"`     // Nome do Banco de Dados
	UserDB             string  `env:"DB_USER" envDefault:"postgres"`           // Usuário banco de dados
	IsProduction       bool    `env:"PRODUCTION" envDefault:"false"`           // Indicação se está em produção ou debug
	IsTestMode         bool    `env:"TEST_MODE" envDefault:"false"`            // Indicação se está em modo de teste
	SSLMode            string  `env:"SSL_MDOE" envDefault:"disable"`           // Conexão SSL do banco de dados
	HostApi            string  `env:"HOST_API" envDefault:"localhost"`         // Host API
	PortApi            string  `env:"PORT_API" envDefault:"9000"`              // Porta API
	LatitudeMonitor    float64 `env:"LATITUDE" envDefault:"-27.6289"`          // Latitude geogrática para monitoramento
	LongitudeMonitor   float64 `env:"LONGITUDE" envDefault:"-48.4478"`         // Longitude geogrática para monitoramento
	TimeZone           string  `env:"TIMEZONE" envDefault:"America/Sao_Paulo"` // Timezone
	CheckTimeInSeconds float64 `env:"CHECKTIME_SECONDS" envDefault:"600.0"`    // Tempo em segundos para a aplicação verificar periodicamente a API externa
}
