FROM golang:latest
WORKDIR /usr/local/app
RUN apt update && apt full-upgrade -y
COPY . .
RUN go get -v -u ./...
RUN go build -v -o bin/main -ldflags "-s -w" main.go
CMD [ "/usr/local/app/bin/main" ]
LABEL Name=go-nginx Version=0.0.1
EXPOSE 8080
