module github.com/prodvana/prodvana-public/go/pkg

go 1.22.0

toolchain go1.22.1

require (
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.19.1 // indirect
	google.golang.org/grpc v1.63.2
	google.golang.org/protobuf v1.33.0
)

replace github.com/prodvana/prodvana-public/go/prodvana-sdk => ../prodvana-sdk

require github.com/prodvana/prodvana-public/go/prodvana-sdk v0.0.0-00010101000000-000000000000

require (
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/planetscale/vtprotobuf v0.6.0 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/rogpeppe/go-internal v1.12.0 // indirect
	golang.org/x/net v0.24.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20240401170217-c3f982113cda // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240415180920-8c6c420018be // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

require (
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.9.0
	golang.org/x/sys v0.19.0 // indirect
	golang.org/x/text v0.14.0 // indirect
)

replace github.com/planetscale/vtprotobuf v0.6.0 => github.com/prodvana/vtprotobuf v0.6.1-pvn
