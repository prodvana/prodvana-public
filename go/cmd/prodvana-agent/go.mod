module github.com/prodvana/prodvana-public/go/cmd/prodvana-agent

go 1.22.0

toolchain go1.22.1

require (
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.19.1 // indirect
	google.golang.org/grpc v1.63.2 // indirect
	google.golang.org/protobuf v1.34.0 // indirect
	k8s.io/apimachinery v0.30.0 // indirect
	k8s.io/client-go v0.30.0
)

replace github.com/prodvana/prodvana-public/go/prodvana-sdk => ../../prodvana-sdk

replace github.com/prodvana/prodvana-public/go/pkg => ../../pkg

require (
	cloud.google.com/go/compute/metadata v0.3.0
	github.com/prodvana/prodvana-public/go/pkg v0.0.0-00010101000000-000000000000
	github.com/prodvana/prodvana-public/go/prodvana-sdk v0.0.0-00010101000000-000000000000
	k8s.io/kubectl v0.29.0
)

require (
	cloud.google.com/go v0.112.2 // indirect
	cloud.google.com/go/auth v0.3.0 // indirect
	cloud.google.com/go/auth/oauth2adapt v0.2.2 // indirect
	cloud.google.com/go/spanner v1.61.0 // indirect
	github.com/GoogleCloudPlatform/grpc-gcp-go/grpcgcp v1.5.0 // indirect
	github.com/aws/aws-sdk-go v1.52.0 // indirect
	github.com/census-instrumentation/opencensus-proto v0.4.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/cncf/xds/go v0.0.0-20231128003011-0fa0005c9caa // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/envoyproxy/go-control-plane v0.12.0 // indirect
	github.com/envoyproxy/protoc-gen-validate v1.0.4 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/getsentry/sentry-go v0.27.0 // indirect
	github.com/go-logr/logr v1.4.1 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/google/s2a-go v0.1.7 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.3.2 // indirect
	github.com/googleapis/gax-go/v2 v2.12.3 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/mxk/go-flowrate v0.0.0-20140419014527-cca7078d478f // indirect
	github.com/planetscale/vtprotobuf v0.6.0 // indirect
	go.opencensus.io v0.24.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.49.0 // indirect
	go.opentelemetry.io/otel v1.26.0 // indirect
	go.opentelemetry.io/otel/metric v1.26.0 // indirect
	go.opentelemetry.io/otel/trace v1.26.0 // indirect
	golang.org/x/crypto v0.22.0 // indirect
	golang.org/x/net v0.24.0 // indirect
	golang.org/x/oauth2 v0.19.0 // indirect
	golang.org/x/sync v0.7.0 // indirect
	golang.org/x/term v0.19.0 // indirect
	golang.org/x/time v0.5.0 // indirect
	golang.org/x/xerrors v0.0.0-20231012003039-104605ab7028 // indirect
	google.golang.org/api v0.176.1 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20240429193739-8cf5692501f6 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240429193739-8cf5692501f6 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	k8s.io/klog/v2 v2.120.1 // indirect
	k8s.io/utils v0.0.0-20230726121419-3b25d923346b // indirect
	sigs.k8s.io/json v0.0.0-20221116044647-bc3834ca7abd // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.4.1 // indirect
)

require (
	github.com/pkg/errors v0.9.1
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.51.0 // indirect
	golang.org/x/sys v0.19.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	k8s.io/api v0.30.0 // indirect
	sigs.k8s.io/yaml v1.4.0 // indirect
)

replace github.com/planetscale/vtprotobuf v0.6.0 => github.com/prodvana/vtprotobuf v0.6.1-pvn