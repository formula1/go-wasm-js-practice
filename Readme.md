Just messing around with Go <=> Javascript

# Build

- [Install Go](https://go.dev/doc/install)
- [Install Docker Compose](https://docs.docker.com/compose/install/)
- terminal - `cd path/to/this/repo`
- terminal - `GOOS=js GOARCH=wasm go build -o ./build/static/main.wasm src/main.go`

# Run
- terminal - `cd path/to/this/repo`
- terminal - `cd build`
- terminal - `docker-compose up`
- web browser - go to `http://localhost:8080`