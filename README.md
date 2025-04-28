# flight-itinerary
Flight itinerary using go - echo framework

Itinerary Service

A simple Go web service that reconstructs a flight itinerary from unordered ticket pairs.

Features:
- Endpoint: POST /itinerary
- Input: JSON array of [source, destination] pairs
- Output: JSON array of airport codes in travel order

Tech Stack:
- Go 1.24
- Echo v4
- Docker
- Make
- OpenAPI 3.0 spec (openapi.yaml)

Project Layout:
cmd/
  server/
    main.go            application entrypoint
internal/
  handler/
    handler.go        HTTP layer
    handler_test.go
  service/
    itinerary.go      core logic
    itinerary_test.go
Dockerfile             container image definition
Makefile               common build/test commands
openapi.yaml           OpenAPI 3.0 API specification
go.mod
.gitignore
README.txt

Getting Started:
1. Clone the repo:
   git clone https://github.com/psdhajare/flight-itinerary.git
   cd flight-itinerary

2. Build & run locally:
   make build
   ./bin/itinerary

3. Or run directly:
   make run

4. Docker:
   docker build -t itinerary-service .
   docker run -p 8080:8080 itinerary-service

5. Tests:
   make test

6. Cleanup & tidy:
   make tidy

The server runs on :8080.  
Send a request:
curl -X POST -H "Content-Type: application/json" \
  -d '[["LAX","DXB"],["JFK","LAX"],["SFO","SJC"],["DXB","SFO"]]' \
  http://localhost:8080/itinerary

Expected response:
["JFK","LAX","DXB","SFO","SJC"]

OpenAPI Spec:
- File: openapi.yaml
- Serves as the formal API definition. You can load it into any Swagger-UI or ReDoc instance.

Makefile Targets:
- build: compile binary to ./bin/itinerary
- run: rebuild & run locally
- test: run all unit tests
- tidy: run go mod tidy
- clean: remove binaries
- docker: build Docker image

Design Decisions:
- Echo framework for lightweight, performant routing and JSON binding
- Service/handler separation to keep core logic testable and HTTP concerns isolated
- Two-stage Docker build for minimal runtime image
- Standard library only for simplicity and minimal external dependencies