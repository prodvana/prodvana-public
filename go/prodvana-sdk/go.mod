module github.com/prodvana/prodvana-public/go/prodvana-sdk

go 1.22

require (
	github.com/envoyproxy/protoc-gen-validate v1.0.4
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.20.0
	github.com/pkg/errors v0.9.1
	github.com/planetscale/vtprotobuf v0.6.0
	google.golang.org/genproto/googleapis/api v0.0.0-20240513163218-0867130af1f8
	google.golang.org/grpc v1.64.0
	google.golang.org/protobuf v1.34.1
)

require (
	golang.org/x/net v0.23.0 // indirect
	golang.org/x/sys v0.18.0 // indirect
	golang.org/x/text v0.15.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240513163218-0867130af1f8 // indirect
)

replace github.com/planetscale/vtprotobuf v0.6.0 => github.com/prodvana/vtprotobuf v0.6.1-pvn
