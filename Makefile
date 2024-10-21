run-http:
	export $(shell cat .env) && \
	go run main.go http

run-grpc:
	export $(shell cat .env) && \
	go run main.go grpc

ent-gen:
	go generate ./ent

ent-new:
	go run -mod=mod entgo.io/ent/cmd/ent new $(Entity)
