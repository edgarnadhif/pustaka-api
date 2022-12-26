postgres:
	distrobox-host-exec podman run -h pustaka-api-server --name pustaka-api-server -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:15-alpine3.17
	
createdb:
	distrobox-host-exec podman exec -ti pustaka-api-server createdb --username=root --owner=root pustaka_api

dropdb:
	distrobox-host-exec podman exec -ti pustaka-api-server dropdb pustaka_api

server:
	go run main.go

.PHONY: postgres createdb dropdb server