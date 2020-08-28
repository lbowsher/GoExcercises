## GracefulShutdown

Starts up a web server and terminates gracefully after 10 seconds.

Once the server is running, visit http://localhost:8080/bananas to see `Hello, I love bananas`, or you can replace bananas with whatever you like.

The main point of this example is to gracefully terminate a web server so that all requests and processes can be completed before the server shuts off. 

The web server shuts down whenever a `SIGTERM` OR `SIGINT` signal is recieved. Currently, after 10 seconds of running, main.go will send its own SIGTERM signal to the server.