swagger: swagger-init
	swag init -g ./internal/server/http/router.go -o api

mockgen-install:
	go install github.com/golang/mock/mockgen@v1.6.0

mockgen: mockgen-install
	mockgen -source=questioner.go -destination=questioner_mock.go -package=main

