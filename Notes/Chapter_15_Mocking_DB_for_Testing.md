# Mocking DB for tests
1. Using the golang mock library to generate the mocks, which is now maintained by uber `https://github.com/uber-go/mock`.
2. After setting up the library, ran the command `mockgen --destination db/mock/store.go github.com/prakhar30/bankoindiana/db/sqlc Store` to generate mocks for the Store.