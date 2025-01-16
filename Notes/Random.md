# Random notes

- Use command `history | grep "some_keyword"` to search through your command history.
- Postgres isolation level documentation `https://www.postgresql.org/docs/current/transaction-iso.html`
- To use postgres with github actions, we need to use a service for it `https://docs.github.com/en/actions/use-cases-and-examples/using-containerized-services/creating-postgresql-service-containers`
- This course is available in textual form at `https://dev.to/techschoolguru/series/7172`
- To generate a random key we can use the command `openssl rand -hex 64 | head -c 32`
- jq is an excellent command line JSON processor `https://jqlang.github.io/jq/` which can be used to process json responses from stuff like AWS secret manager or something.
- To make any file executable, like the wait-for.sh we need to use the `chmod +x wait-for.sh` command.
- We can use the Go Statik library to bundle front end files in our BE binary. `https://github.com/rakyll/statik`
- To write structured JSON logs, we are going to use a library called Zerolog, `https://github.com/rs/zerolog`
- We have modified the VS code settings for `go test flag` to 
```
"go.testFlags": [
        "-v",
        "-count=1",
    ]
```
so that we run the tests with verbose logs and never caching the tests, which is something that VS code does, annoyingly.
- To update go version in the project, it needs to reflect in 2 places - Dockerfile, should specify the correct image, and the go.mod file. Then go mod tidy to resolve any dependencies. Also in the ci-test.yml for the Github test workflow.
- Can use `go env GOPATH` to check where the go bin and other folders are located.
- To handle CORS, we can use a library like `https://github.com/rs/cors`