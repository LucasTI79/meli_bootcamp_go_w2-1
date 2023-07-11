
# Instalação

Instale o MySQL: 
```bash
brew install mysql Option2: -arm64 brew install mysql
```

Verifique o status do MySQL: execute o comando 'make build-database' da raiz do projeto (comando declarado no Makefile).
```bash
make build-database
```
> Observação: não coloque senha.

Execute o comando de criação do banco de dados: Verifique com o status 'mysql.server' para verificar se o serviço foi inicializado. 
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
