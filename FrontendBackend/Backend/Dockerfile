FROM golang:1.14-alpine as builder

WORKDIR /Backend

# These lines are crucial so that the Docker has dependencies installed and available to use
#COPY go.mod ./
#RUN go mod download

COPY server.go .
RUN go build -o /app server.go

FROM alpine:3.10
# Define GOTRACEBACK to mark this container as using the Go language runtime
# for `skaffold debug` (https://skaffold.dev/docs/workflows/debug/).
ENV GOTRACEBACK=single
CMD ["./app"]
COPY --from=builder /app .