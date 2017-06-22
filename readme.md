# Example Ginkgo/Gomega Golang test framework

An example test framework to get started with API testing using Golang. Uses the TheyWorkForYou.com API as an example.

## Tech stack

- Go
- Ginkgo: test runner & framework
- Gomega: assertion library
- Resty: http client

## How to get started

- Get an API key from https://www.theyworkforyou.com/api/key - you'll have to sign up for an account
- Export the API key to your environment: `export TWFY_KEY=...`
- Install go
- get this project: `go get github.com/burythehammer/golang-api-test-quickstart`
- Download dependencies (see below)
- Run your tests: `go test`

NOTE: some tests will fail. This is intentional! The theyworkforyou.com API returns `200` when it can't find an MP. I consider that a bug, don't you? :)

## Dependencies

```
go get github.com/onsi/ginkgo/ginkgo
go get github.com/onsi/gomega
go get -u gopkg.in/resty.v0
```