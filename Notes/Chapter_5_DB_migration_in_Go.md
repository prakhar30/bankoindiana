# DB Migration in Go
1. This chapter will help us get versed with DB migrations in Go as they are an essential part of any application. Here we will make use of the `golang-migrate` library to do this migration which can be found at `https://github.com/golang-migrate/migrate`
2. We will follow the CLI documentation here. On a mac we can use homebrew to install the library. `brew install golang-migrate`. To check if successfully installed we can run `migrate --version` or `migrate --help` to see the manual. 
3. We will use the command `migrate create -ext sql -dir db/migration -seq init_schema` to create our first migration files. here `-ext` is the extension of the file, `-seq` flag to generate a sequential version of the migration file. 
4. With our docker container running we can create a new db by going inside our container shell by using the command `docker exec -it postgresLatest /bin/sh` and then doing a `createdb --username=root --owner=root banko_indiana`.
5. Create a makefile step to contain the list of all helpful commands that could be executed easily.
6. After the creation of a new container and database, we are ready to execute our first migration - `migrate -path db/migration -database "postgresql://root:something_secret@localhost:5432/banko_indiana?sslmode=disable" -verbose up` here the `sslmode` argument is important cause without it the command will fail as our postgres container by default has SSL disabled. If everything goes alright, all the tables should have been created in the DB. Both these commands are added to the make file as well to execute easily.