## How To Run

1. Run docker compose with command `docker-compose up`
2. Run application with command `go run cmd/app/main.go`

## How To Test

You can use below link for test endpoints

[localhost:9999/swagger/index.html](localhost:9999/swagger/index.html)

Also you can run tests with command

 `go test ./tests/account && go test ./tests/transaction`

## Environment Usage

SERVER.PORT for API port
DATASOURCE.SEED for save existing accounts 
DATASOURCE.CLEANUP_INTERVAL for cache config
DATASOURCE.DEFAULT_EXPIRATION for cache config

## Save Existing Accounts

You must add .csv file for save existing accounts. 
```csv
accountId,balance
1,10.00
2,20.00
3,30.00

```