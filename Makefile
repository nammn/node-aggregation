dependency:
	@go get -v ./...

integration-test: docker-start dependency
	@go test -v ./...

docker-start:
	@docker run -p 6379:6379 --name some-redis redis

unit-test: dependency
	@go test -v -short ./...

docker-clean:
	@docker container stop $(docker container ls -aq)
