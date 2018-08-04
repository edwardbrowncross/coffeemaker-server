set GOOS=linux
set GOARCH=amd64
go build -ldflags='-s -w' -o bin\coffeestatechange coffeestatechange\main.go;