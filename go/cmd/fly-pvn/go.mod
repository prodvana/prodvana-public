module github.com/prodvana/prodvana-public/go/cmd/fly-pvn

go 1.22

require (
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.19.1 // indirect
	github.com/spf13/cobra v1.8.0
	google.golang.org/grpc v1.63.2
	google.golang.org/protobuf v1.33.0 // indirect
)

replace github.com/prodvana/prodvana-public/go/prodvana-sdk => ../../../go-sdk

require (
	github.com/bradenaw/juniper v0.15.3
	github.com/google/go-containerregistry v0.19.1
	github.com/hashicorp/go-multierror v1.1.1
	github.com/pelletier/go-toml/v2 v2.2.0
	github.com/prodvana/prodvana-public/go/prodvana-sdk v0.0.0-00010101000000-000000000000
	golang.org/x/exp v0.0.0-20231006140011-7918f672742d
)

require (
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/envoyproxy/protoc-gen-validate v1.0.4 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/planetscale/vtprotobuf v0.6.0 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/net v0.23.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20240314234333-6e1732d8331c // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240401170217-c3f982113cda // indirect
)

require (
	github.com/pkg/errors v0.9.1
	golang.org/x/sys v0.18.0 // indirect
	golang.org/x/text v0.14.0 // indirect
)
