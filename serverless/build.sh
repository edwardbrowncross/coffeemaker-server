env GOOS=linux go build -ldflags="-s -w" -o bin/coffeestatechange coffeestatechange/main.go
env GOOS=linux go build -ldflags="-s -w" -o bin/slackwebhook slackwebhook/main.go