# Docker
1. List all running container - `docker ps`. `docker ps -a` lists all containers, even stopped ones.
2. List all images available - `docker images`
3. `hub.docker.com` is like a central place to get any application image that we might need. Can search for `postgres` here and use the official image to get the image. 
4. Can get a specific version from here by using something like `docker pull postgres:12-alping`. But for this course going ahead with the default command and see what how it goes. Basically syntax is `docker pull <image>:<tag>`
5. Basic information on how to run the image in the container will be available on the docker image page. For our usecase we can use the command `docker run --name some-postgres -e POSTGRES_PASSWORD=mysecretpassword -d postgres` here `--name` specifies the name of the container, `-e` flag is to set the environment vars, `-d` augument is to run it in background mode(or detach mode), and the last argument is the name of the image. 
6. We will run a little modified command for this `docker run --name postgresLatest -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=something_secret -d postgres`. An important argument is the `-p` arugment which is for port mapping, basically docker container is run on a separate virtual network which is different than the host network, hence we need to create a kinda bridge to connect to the postgres server.
7. Hitting `docker ps` after the above command will show us that we have successfully started a container with our image. 
8. A `docker exec -it <container_name_or_id> <command> [args]` command lets us run a single command inside the docker container. The `-it` flag tells docker to run the command as an interactive tty(teletype) session.
9. Hence to access the postgres console inside the docker container we run `docker exec -it postgresLatest psql -U root` . We can exit this console by `\q`
10. To see the logs of the container we can use the `docker logs <container_name_or_id>` command. 
11. We will use `TablePlus` DB GUI tool to interact with the DB to speed up things a little bit. Enter the basic details used during the setup and then connect to the database in the tool.
12. Get the generated SQL from the snippet that we created in DB Diagram website in the tool and run the entire script to create the data base schema. On running it successfully and hitting refresh all tables should have been created. 
13. Use the command `docker stop <container_name>` to stop the container.
14. After stopping the command we will need to use `docker start <container_name>` to run the already created container again. 