init:
	go install -v ./vendor/github.com/gogo/protobuf/protoc-gen-gogofast

gen:
	find ./api -name '*.pb.go' -delete
	protoc --proto_path=./ --gogofast_out=Mapi/bar/models.proto=github.com/gebv/go-protoc-gen-example/api/bar,plugins=grpc:. ./api/foo/*.proto
	protoc --proto_path=./ --gogofast_out=grpc:. ./api/bar/*.proto

run:
	go run cmd/main.go
