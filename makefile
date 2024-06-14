# Makefile para automatizar comandos Go

# Limpa o cache do Go
clean:
	go clean -cache

# Executa os testes
test:
	go test ./tests/...

# Comando padrão
all: clean test

# Medir o tempo de execução dos testes
timed-test:
	time go test ./tests/...
