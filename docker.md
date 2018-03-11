# Docker and Microservices

Docker is a container technology that shares the Linux kernel between containers that are partitioned from each other and maintain  separate namespaces.  This seperation allows containers to very lightweight and hold only the dependencies they need to run a particular application.  They do not try to emulate the hardware of a server like a Virtual Machine (VM), they share the hosts Linux kernel.  They can be used to develop very small, dedicated microservers that contain only the executables needed to perform a particular role.

Containers find application in Continuous Integration and Deployment. They can be used for the build process, taking in source code and returning an executable.  A smaller container containing just the executable becomes a microservice container.

How do we use Docker to build a Go program?  We will need a container image that contains a Golang compiler and its dependencies.  We then need to add the source code and dependent packages and then compile into an executable.

How to we use Docker to create a microservice container?  We need an empty 'scratch container' and then transfer the previously compiled executable.

How and where the microservice might be deployed depends on the hosting service used. Cloud hosting services support containers.  Typically they provide a VM that run an operating system that supports containers (most Linux variants) and runs the Docker runtime server.  A VM can be configured to run containers ad-hoc or at startup by configuring SystemD.











