FROM golang:1.23

WORKDIR /app
ADD . /app/

RUN go build -o ./bin/inspiration-tech-case ./cmd/app/main.go

EXPOSE 9999

ENTRYPOINT ["./bin/inspiration-tech-case"]