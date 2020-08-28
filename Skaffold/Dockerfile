FROM golang:1.14-alpine as builder

RUN apk add git

WORKDIR /Skaffold

# These lines are crucial so that the Docker has dependencies like client-go installed and available to use
COPY go.mod go.sum ./
RUN go mod download

COPY main.go .

RUN go build -o /app main.go

FROM alpine:3.10
# Define GOTRACEBACK to mark this container as using the Go language runtime
# for `skaffold debug` (https://skaffold.dev/docs/workflows/debug/).
ENV GOTRACEBACK=single
CMD ["./app"]
COPY --from=builder /app .
