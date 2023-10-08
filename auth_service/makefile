run:
	go run main.go -configFile=config.json -logFilePath=authServiceLogs.txt -useFileLog=true -port=8080

build:
	make clear
	GOOS=linux GOARCH=amd64 go build -o XMAuthAPI main.go 
	zip main.zip XMAuthAPI
	rm -f XMAuthAPI

clear:
	rm -f XMAuthAPI main.zip

test:
	go test `go list ./... | grep -v examples` -coverprofile=../auth_coverage.txt -covermode=atomic