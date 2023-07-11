
# Instalação

Instale o MySQL: 
```bash
brew install mysql
```
Opção alternativa: 
```bash
-arm64 brew install mysql
```

Construa o banco de dados executando o seguinte comando a partir da raiz do projeto:
```bash
make build-database
```
> Observação: não coloque senha.

Verifique com o status 'mysql.server' se o MySQL foi inicializado. 
```bash
mysql.server
```

Caso contrário, execute o comando 'mysql.server start':
```bash
mysql.server start
```

# Execução

Para rodar o projeto, execute:
```bash
make start
```

Para executar os testes, execute:
```bash
make test
```

Para executar os testes e ver a análise de cobertura, execute:
```bash
make test-cover
```
