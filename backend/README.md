# Backend 

The backend has two routes 

`/api/settings` -> GET threshold and on/off
`/api/settings` -> POST change the threshold and toggle.
`/api/measurements` -> POST dust reading


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

Then run `go run cmd/local/main.go` and it will start the server locally on port 8080.

### Example Requests

#### Get settings 

```
$ curl http://localhost:8080/api/settings --fail
```

#### Update Settings

```
curl -X PUT http://localhost:8080/api/settings -d '{"isOn":true,"threshold":200}' --fail``
```
