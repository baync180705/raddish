run:
	go run cmd/server/main.go
dev:
	$(shell go env GOPATH)/bin/air
clean:
	go clean
	rm -f tmp/main
docker-dev:
	docker-compose up
docker-build-prod:
	docker build -t raddish-prod .