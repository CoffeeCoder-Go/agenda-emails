BIN := ./bin/agenda-emails-v1.exe
CMD := ./cmd/server/main.go

OS := linux
ARCH := amd64

build:
	@echo "Construindo..."
	@GOOS=${OS} GOARCH=${ARCH} go build -o ${BIN} ${CMD}
	@echo "Build terminada!"

run:
	@mkdir db
	@${BIN}