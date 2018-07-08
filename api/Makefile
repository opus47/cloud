.PHONY: server
server: build/opus47-server

build:
	mkdir -p build

cmd/opus47-server/main.go: openapi.yml
	swagger generate server -f openapi.yml

build/opus47-server: restapi/handlers.go restapi/db.go cmd/opus47-server/main.go | build 
	dep ensure
	go build -o $@ ./cmd/opus47-server

.PHONY: container
container:
	docker build -f api.dock -t opus47/api:latest .


.PHONY: clean
clean:
	rm -rf models
	rm -rf build
	rm -rf cmd
