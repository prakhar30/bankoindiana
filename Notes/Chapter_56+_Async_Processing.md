# Async Processing

- For this tutorial, we will be using a in-dev third party library called Asyncq, which is backed by Redis. `https://github.com/hibiken/asynq`
- To run a redis server on our local machine, we will use a docker image, which we can search on `hub.docker.com` and use the official redis image from there. 
- We used a DB transaction to send verify email after creation of a user, mainly cause of any errors received from redis will give the client a failure message, which might make the client retry. However, because the DB call would be successful, there will already be a record created in the DB. This is not a good way to do it, hence DB transactions are an elegant way to do it.