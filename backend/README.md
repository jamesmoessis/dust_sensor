# Backend 

## API

The backend has two routes 

* `/api/settings` -> GET threshold and on/off
* `/api/settings` -> POST change the threshold and toggle.
* `/api/measurements` -> POST dust reading


## Development Prerequisites

* Golang 1.20 or above
* AWS CLI installed, with AWS credentials for the relevant account
* `zip` command line utility

## Test

```
go test ./...
```

## Running Local

There's a slightly janky local setup which makes it slightly easier than running the 
localstack/lambda/docker setup. Note, this bypasses the lambda and just runs as a regular
golang `net/http` server, so it does not test the lambda features of the application, 
particularly those under `./cmd/lambda`.

First, you need the AWS CLI setup locally so the AWS SDK can find your credentials to connect
to Dynamo DB.

Then run `go run ./cmd/local` and it will start the server locally on port 8080.

### Example Requests

#### Get settings 

```shell
$ curl http://localhost:8080/api/settings --fail
```

#### Update Settings

```shell
curl -X PUT http://localhost:8080/api/settings -d '{"isOn":true,"threshold":200}' --fail``
```

## Building and Deploying 

```shell
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o bootstrap ./cmd/lambda && \
zip lambda.zip bootstrap && \
aws lambda update-function-code --function-name dust_sensor_settings_api --zip-file fileb://lambda.zip
```

## Architecture Overview

The built golang applications is run from AWS Lambda, and stores the state in a simple Dynamo DB table. 
The public facing endpoint is API gateway, and this takes the request and marshals it into the 
correct event type that the function will understand.
