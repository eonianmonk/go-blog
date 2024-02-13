proto:
	protoc --experimental_allow_proto3_optional --go_out=./assets/pb_generated --go-grpc_out=./assets/pb_generated/ --proto_path=./assets/pb ./assets/pb/post.proto
	protoc --go_out=./assets/pb_generated --go-grpc_out=./assets/pb_generated/ --proto_path=./assets/pb ./assets/pb/comment.proto

sqlc:
	sqlc generate

serve_rpc:
	echo "Not yet"

serve_gateway:
	echo "Not yet"

build:
	go build -o blogo ./cmd