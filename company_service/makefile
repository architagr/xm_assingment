run:
	go run main.go -configFile=config.json -logFilePath=companyServiceLogs.txt -useFileLog=true -port=8080

build:
	make clear
	GOOS=linux GOARCH=amd64 go build -o XMCompanyAPI main.go 
	zip main.zip XMCompanyAPI
	rm -f XMCompanyAPI

clear:
	rm -f XMCompanyAPI main.zip

test:
	go test `go list ./... | grep -v examples` -coverprofile=../company_coverage.txt -covermode=atomic