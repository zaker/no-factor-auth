FROM golang as base

RUN update-ca-certificates

WORKDIR /code

COPY ./go.mod ./go.sum ./
RUN go mod download
COPY . .

ENV CGO_ENABLED=0
ENV GOOS=linux

RUN go test
RUN go build

FROM scratch

ARG AUTHSERVER
ARG TENANT_ID

COPY --from=base /code/no-factor-auth .

CMD ["./no-factor-auth"]
