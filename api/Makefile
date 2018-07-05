.PHONY: server
server: build/opus47-server

build:
	mkdir -p build

build/opus47-server: restapi/handlers.go cmd/opus47-server/main.go | build 
	swagger generate server -f openapi.yml
	dep ensure
	go build -o build/opus47-server ./cmd/opus47-server

.PHONY: container
container:
	docker build -f api.dock -t opus47/api:latest .


.PHONY: clean
clean:
	rm -rf models
	rm -rf build
