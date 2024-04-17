build:
	@go build -o bin/api

run: build 
	@./bin/api

test:
	@go test -v ./...

docker:
	echo "building docker file"
	@docker build -t api .
	echo "running inside Docker container"
	@docker run -p 3000:3000 --net="host" api