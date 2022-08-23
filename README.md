# L0-WB
## Docker start
```
docker-compose build; docker-compose up
```
or
```
docker-compose build && docker-compose up
```
## Start service
```
go run main.go
```
## Publish messages
```
cd test/nats/publisher
go run ./publisher.go ## if publish standart model with filename "model.json"
go run ./publisher.go <filename1 filename2...> ## if publish few model and any filename
```
## Testing
```
cd test/testing
go test -v
```
