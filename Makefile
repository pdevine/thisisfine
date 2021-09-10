build:
	CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static"' -o bin/thisisfine tif.go

run:
	go run tif.go

compile:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static"' -o bin/thisisfine-linux-amd64 tif.go
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static"' -o bin/thisisfine-windows-amd64.exe tif.go
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static"' -o bin/thisisfine-darwin-amd64 tif.go



