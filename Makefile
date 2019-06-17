dependency:
	@go get -v ./...

integration-test: docker-start dependency
	@go test -v ./...

docker-start:
	@docker run -p 6379:6379 -d --name some-redis redis
	- docker build -t node-aggregation .
	@docker run -d -p 8000:8000 node-aggregation

unit-test: dependency
	@go test -v -short ./...

clean: docker-clean

docker-clean:
	@echo "Stopping all running containers"
	- docker container stop `docker container ls -aq`
	@echo "Remove all non running containers"
	- docker rm `docker ps -q -f status=exited`
	@echo "Delete all untagged/dangling (<none>) images"
	- docker rmi `docker images -q -f dangling=true`

