PROTOC := protoc --proto_path=./ --gogofast_out=plugins=grpc:.
# PROTOC := protoc -I. --gogofast_out=plugins=grpc:.

init:
	go install -v ./vendor/github.com/gogo/protobuf/protoc-gen-gogofast

gen:
	find ./api -name '*.pb.go' -delete
	${PROTOC} ./api/foo/*.proto
	${PROTOC} ./api/bar/*.proto

run:
	go run cmd/main.go
