# Deployment Notes

1. Create a Dockerfile with all the necessary info. Important point to remember about EXPOSE port it, it doesn't really exposes the port, but just serves as a documentation purpose between the builder and the deployer.
2. Once the docker file is setup, we will use the command `docker build -t bankoindiana:latest .` to build the app, supplying a tag using the -t argument. 
3. A multistage docker file is used to just copy the binary to the image, thereby massively reducing the size.