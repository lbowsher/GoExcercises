## FrontendBackend
A simple "frontend" web app that sends post requests to a backend web server. The frontend and backend are each meant to be deployed in their own seperate containers, but within the same kubernetes pod. Also setup for development with Skaffold.

### Backend

The Backend server.go sets up a web server that prints "Hello, I love x" when you visit the page `localhost:8080/x` where x can be replaced by whatever you want. 

Additionally, whenever a POST request is sent to `localhost:8080/send`, the server will log all of the data sent to it. After recieving a POST request, it will respond with the message: "Received a POST Request"

### Frontend

The frontend sends a POST request to the backend via localhost port 8080 with the sample data:

    "name":"test",
    "salary":"123",
    "age":"23"

It will then log the response it gets from the server.