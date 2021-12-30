send:	
		go run ./message_sender/nats_sender.go
main:
		go run ./cmd/main.go ./cmd/web.go ./cmd/stan.go
initDB:
		go run db/db_init.go
test:
		go test ./...
dockerStart:
		docker start wizardly_kirch postgres
dockerStop:
		docker stop wizardly_kirch postgres
dockerRestart: dockerStop dockerStart
