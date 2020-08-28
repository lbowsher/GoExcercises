## ContainerizedDeployment

Simple web server example written in Golang. 

Displays on the web page Hello World as well as the version of this app and the hostname.

Uses an environment variable named `PORT` for the web server's port, or it will default to 8080 if `PORT` is undefined.

Dockerfile setup for containerized deployment and it sets `PORT` to 8080