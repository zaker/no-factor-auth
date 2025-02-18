FROM golang as base

RUN update-ca-certificates

WORKDIR /code

COPY ./go.mod ./go.sum ./
RUN go mod download
COPY . .

ENV CGO_ENABLED=0

RUN go test ./...

RUN go build -ldflags "-X main.Version=$TAG" -o no-factor-auth.linux

ARG AUTHSERVER
ARG TENANT_ID

EXPOSE 8089
CMD ["./no-factor-auth"]
