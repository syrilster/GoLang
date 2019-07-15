# A Simple Multi-Stage Dockerfile
* Using multi-stage Dockerfiles, we can pick apart the tasks of building and running our Go applications into different stages. Typically, we start off with a large image which includes all of the necessary dependencies, packages, etc. needed to compile the binary executable of our Go application. This would be classed as our builder stage.
* We then take a far more lightweight image for our run stage which includes only what is absolutely needed in order to run a binary executable. This would typically be classed as a production stage or something similar.
