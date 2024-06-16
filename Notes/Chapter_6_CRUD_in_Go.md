# CRUD in Go

- We will use SQLC to generate Go code that interacts with database for us. We will create normal db relations in go code and SQLC will generate code to interact with DB. `sqlc.dev`. Can find installation instructions there, but for mac currently easiest way is via brew `brew install sqlc`. Run command `sqlc` in terminal to verify installation.
-  To generate the sqlc yaml file, hit `sqlc init` in the root of your project.
- The `sqlc.yaml` file created in the root of your directory will contain all the config for sqlc code generatation. The format and values of this can be taken from the documentation available on Github and customised accordingly. 
- The `sqlc generate` command is also added to the makefile. This command will initially fail as there will be no queries in the `/query` folder. Need to add these for the command to successfully generate code. 
- Once a query has been added to the query folder using the docs, the generate command will generate necessary files and method. 
- An important command to run after this will be to create the `go mod init https://github.com/prakhar30/bankoindiana` command to create the `go.mod` file which is like a dependency management file for go. Once this is done run `go mod tidy` to download and add any dependencies required by the project. This should resolve any errors the generated files have been showing so far.