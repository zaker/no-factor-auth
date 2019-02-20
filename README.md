# This is a no factor "Authentication" service aka YOLO Auth

It doesn't provide security, it provides a signed token to services that needs to perform e2e testing

## Setup

It needs a pem encode RSA certificate, for now called cert.pem.

And since it's mostly zero config, it just starts on <http://0.0.0.0:8089>, all incomming on local server.

## Run

```bash
go build
./no-factor-auth
```
