# Go Expert Challenge - Stress Test CLI

Implementação de um CLI em Go para realizar testes de stress em um endereço web.

## Arquitetura

As requisições são realizadas de forma concorrente de acordo com a quantidade informada e são distribuídas em um pool de workers. Cada worker é responsável por realizar uma requisição HTTP e armazenar o resultado em um canal de comunicação. O resultado é então processado e exibido ao final da execução.

## Como executar

### Via Docker

O projeto está disponível no Docker Hub, e para executar o container basta utilizar o comando abaixo, substituindo os valores de `--url`, `--requests` e `--concurrency` pelos valores desejados.

```sh
docker run mathcale/goexpert-stresstest-cli-challenge:latest \
    --url https://google.com.br \
    --requests 100 \
    --concurrency 10
```

### Localmente

Para executar o projeto localmente, é necessário ter o Go instalado na máquina. Após a instalação, basta executar o comando abaixo, substituindo os valores de `--url`, `--requests` e `--concurrency` pelos valores desejados.

```sh
go run cmd/cli/main.go \
    --url https://google.com.br \
    --requests 100 \
    --concurrency 10
```

## Testes

Para executar os testes de unidade, basta executar o comando abaixo.

```sh
make test
```
