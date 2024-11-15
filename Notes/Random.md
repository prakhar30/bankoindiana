# Random notes

- Use command `history | grep "some_keyword"` to search through your command history.
- Postgres isolation level documentation `https://www.postgresql.org/docs/current/transaction-iso.html`
- To use postgres with github actions, we need to use a service for it `https://docs.github.com/en/actions/use-cases-and-examples/using-containerized-services/creating-postgresql-service-containers`
- This course is available in textual form at `https://dev.to/techschoolguru/series/7172`
- To generate a random key we can use the command `openssl rand -hex 64 | head -c 32`
- jq is an excellent command line JSON processor `https://jqlang.github.io/jq/` which can be used to process json responses from stuff like AWS secret manager or something.
- To make any file executable, like the wait-for.sh we need to use the `chmod +x wait-for.sh` command.