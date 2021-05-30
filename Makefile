go-build:
	@go build -v -o go-accounts-app *.go

go-run:
	@./go-accounts-app

run:
	make go-build
	make go-run

compile-proto:
	@protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative rpc/accounts.proto