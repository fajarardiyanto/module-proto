module github.com/fajarardiyanto/module-proto

go 1.14

replace github.com/fajarardiyanto/module-proto => ./

require (
	github.com/favadi/protoc-go-inject-tag v1.3.0 // indirect
	google.golang.org/grpc v1.49.0
	google.golang.org/protobuf v1.27.1
)
