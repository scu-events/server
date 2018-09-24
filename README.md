# Server
Back-end code for scu-events

# Dependencies

```
go get google.golang.org/api/calendar/v3
go get golang.org/x/oauth2
go get github.com/mmcdole/gofeed
go get cloud.google.com/go
```

# Google Calendar API keys

1. Create OAuth Client ID (Application type: Other)
2. Download the credentials, save it as client_secret.json
3. Run the server `go run server.go`
4. Follow the instruction in the log and the browser to retrieve token.json

# Run

`go run server.go`

# Deploy
`GOOS="linux" GOARCH="amd64" go build server.go`
Then deploy the binary executable to a linux based machine
