# Code Challenge Wind Forecast

Code challenge.


### Configurando Projeto localmente

Ao realizar git clone do projeto, acesse o diretório do mesmo e execute: 

1. Baixar dependências e outras verificações ``go mod tidy``
2. Para executar como dev ``go run main.go``
3. Para compilar a aplicação ``go build``


### Mais informações

Ao executar a API, estará executando na porta ``9000``, com endpoins em:

Previsão
````
    http://localhost:9000/api/previsao
````

Alertas
````
    http://localhost:9000/api/alerta
````



# Teste 

Para executar testes no projeto, ao acessar esse diretório, execute:

````
    TEST_MODE=true go test ./...
````


### Cobertura dos testes

Para verificar quanto está a cobertura dos testes no projeto, execute: 

````
    TEST_MODE=true  go test -coverprofile=coverage ./...
````

Em seguida:

````
    go tool cover -html=coverage
````



# Comandos 

Algumas variáveis de ambiente podem ser inicializadas juntamente com o comando de execução da aplicação, como executar em ``modo de produção``:

````
    PRODUCTION=true go run main.go
````

Com uma configuração de banco de dados diferente em seu computador, utilize as variáveis de ambiente ``DB_HOST``, ``DB_USER``, ``DB_PASS``, ``DB_PORT`` e ``DB_NAME``.

````
    DB_HOST=localhost DB_USER=postgres DB_PASS=pass DB_PORT=5432 DB_NAME=dbpostgres go run main.go
````



# Infra

Os arquivos **`Dockerfile`** e **`docker-compose.yaml`** são responsáveis por gerar os recursos de infra e configuração necessários para a(as) aplicação(ções) executar(em), como: ***portas de acesso***, ***banco de dados***, ***senhas de acesso***, ***rede interna*** para acesso entre as mesmas. Eis abaixo algumas recomendações para o respectivo projeto:


### Banco de Dados PostgreSQL

O container do Banco de Dados está declarado commo `forecast_db`, e a configuração para acesso ao mesmo encontra-se na declaração deste container, no arquivo `docker-compose.yaml`, onde estão sendo repassadas algumas variáveis de ambiente: `POSTGRES_PASSWORD`; `POSTGRES_USER` e `POSTGRES_DB` - todas as variáveis necessárias com valores padrão estão no arquivo ``.env.example``, onde é necessário copiar esse arquivo para ``.env`` e alterar os valores das variáveis. O serviço irá executar na porta `5432`. Os arquivos referentes ao banco de dados serão persistidos como volume em `./db-data` no contexto da pasta do projeto.


### Backend Golang

O container do backend está declarado no mesmo arquivo `docker-compose.yaml` como `forecast`. Todas as variáveis necessárias estão declarados também no arquivo ``.env.example``, como `DB_HOST`, `DB_PORT`, `DB_PASS`, `DB_NAME`, `DB_USER`, `PRODUCTION`, `SSL_MDOE`, `PORT_API`, `LATITUDE`, `LONGITUDE` e `TIMEZONE`. 


### Variáveis de Ambiente

Copie o arquivo ``.env.example`` para ``.env`` e atualize seus valores, essas variáveis serão inicializadas quando utilizado o arquivo ``docker-compose.yaml`` em implantação.


### Comandos para executar a Infra

Para construir os containers utilizando Docker Compose, acesse esta pasta do projeto e use o comando ``docker compose up``, para recompilar ``docker compose up --build``. Para desligar e ligar todos os containers, use respectivamente ``docker compose stop`` e ``docker compose start``.


#### Importante

***Ao construir a imagem do container da aplicação em Docker, está sendo executado os testes unitários, antes dessa fase.***