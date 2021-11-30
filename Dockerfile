FROM golang:1.16-alpine

WORKDIR /app

ENV GO111MODULE="on"
ENV GOOS="linux"
ENV CGO_ENABLED=0

# Debug
RUN go install github.com/cosmtrek/air@latest && \
    go install github.com/go-delve/delve/cmd/dlv@latest

# Dependencies
COPY ./go.mod .
COPY ./go.sum .
RUN go mod download && go mod verify

COPY . .

RUN go build -o /usr/local/bin/go-auth .

EXPOSE 8080
EXPOSE 2345

CMD [ "go-auth" ]