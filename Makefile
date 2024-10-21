run-http:
	export $(shell cat .env) && \
	go run main.go http

run-grpc:
	export $(shell cat .env) && \
	go run main.go grpc

visualize:
	atlas schema inspect -u "ent://ent/schema" --dev-url "postgresql://moviedb_owner:TEO5xKFjn0yA@ep-crimson-voice-a1n3bsrl.ap-southeast-1.aws.neon.tech/moviedb?sslmode=require&search_path=public" -w

ent-gen:
	go generate ./ent

ent-new:
	go run -mod=mod entgo.io/ent/cmd/ent new $(Entity)
