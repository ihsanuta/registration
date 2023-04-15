.PHONY: run swag-generate cert

run:
	@go run .

swag-generate:
	@swag init -g app/handler/handler.go --outputTypes go

cert:
	openssl genrsa -out cert/id_rsa 4096
	openssl rsa -in cert/id_rsa -pubout -out cert/id_rsa.pub