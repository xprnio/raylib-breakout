bin/server:
	go build -o bin/server cmd/server/main.go

bin/raygo:
	go build -o bin/raygo cmd/game/main.go

server: bin/server
	bin/server

game: bin/raygo
	bin/raygo
