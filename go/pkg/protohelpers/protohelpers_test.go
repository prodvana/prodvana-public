package protohelpers

import (
	"crypto/rand"
	"fmt"
	"testing"

	"github.com/prodvana/prodvana-public/go/pkg/protohelpers/prototestutil"

	common_config_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/common_config"
	service_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/service"
	"github.com/stretchr/testify/require"
)

func TestTemplate(t *testing.T) {
	msg := service_pb.ServiceConfig{
		Name: "my-service",
		Programs: []*common_config_pb.ProgramConfig{
			{
				Name:  "{{.Foo}}bar",
				Image: "{{.MyImage}}",
				Cmd:   []string{"my-cmd", "--foo={{.Foo}}", "--bar", "--quoted={{.Quoted}}"},
				Env: map[string]*common_config_pb.EnvValue{
					"Key{{.Foo}}": {
						ValueOneof: &common_config_pb.EnvValue_Value{
							Value: "my-value",
						},
					},
					"Key2": {
						ValueOneof: &common_config_pb.EnvValue_Value{
							Value: "my-value{{.Bar}}",
						},
					},
				},
			},
		},
		Replicas: &service_pb.ReplicasConfig{
			ConfigOneof: &service_pb.ReplicasConfig_Fixed{
				Fixed: 1,
			},
		},
	}
	require.NoError(t, TemplateExecute(&msg, map[string]interface{}{
		"Foo":     "1",
		"MyImage": "my-image",
		"Bar":     "baz",
		"Quoted":  `{"quoted"}`,
	}))

	prototestutil.RequireEquals(t, &service_pb.ServiceConfig{
		Name: "my-service",
		Programs: []*common_config_pb.ProgramConfig{
			{
				Name:  "1bar",
				Image: "my-image",
				Cmd:   []string{"my-cmd", "--foo=1", "--bar", `--quoted={"quoted"}`},
				Env: map[string]*common_config_pb.EnvValue{
					"Key1": {
						ValueOneof: &common_config_pb.EnvValue_Value{
							Value: "my-value",
						},
					},
					"Key2": {
						ValueOneof: &common_config_pb.EnvValue_Value{
							Value: "my-valuebaz",
						},
					},
				},
			},
		},
		Replicas: &service_pb.ReplicasConfig{
			ConfigOneof: &service_pb.ReplicasConfig_Fixed{
				Fixed: 1,
			},
		},
	}, &msg)
}

func TestTemplateStringMap(t *testing.T) {
	msg := service_pb.ServiceConfig{
		Name: "my-service",
		ConfigOneof: &service_pb.ServiceConfig_Terraform{
			Terraform: &service_pb.TerraformConfig{
				BackendConfig: map[string]string{
					"key1":         "{{.Foo}}1234",
					"key2{{.Bar}}": "1234",
					"key3":         "{{.Quoted}}",
				},
			},
		},
	}
	require.NoError(t, TemplateExecute(&msg, map[string]interface{}{
		"Foo":     "1",
		"MyImage": "my-image",
		"Bar":     "baz",
		"Quoted":  `{"quoted"}`,
	}))

	prototestutil.RequireEquals(t, &service_pb.ServiceConfig{
		Name: "my-service",
		ConfigOneof: &service_pb.ServiceConfig_Terraform{
			Terraform: &service_pb.TerraformConfig{
				BackendConfig: map[string]string{
					"key1":    "11234",
					"key2baz": "1234",
					"key3":    "{\"quoted\"}",
				},
			},
		},
	}, &msg)
}

func TestCompression(t *testing.T) {
	for _, tc := range []struct {
		bytes []byte
	}{
		{
			bytes: []byte{},
		},
		{
			bytes: []byte("hi"),
		},
		{
			bytes: []byte("hi\n\n1234"),
		},
	} {
		t.Run(fmt.Sprintf("gzip_%s", tc.bytes), func(t *testing.T) {
			gzipped, err := GzipBytes(tc.bytes)
			require.NoError(t, err)
			ungzipped, err := UngzipBytes(gzipped)
			require.NoError(t, err)
			require.Equal(t, tc.bytes, ungzipped)
		})
		t.Run(fmt.Sprintf("zstd_%s", tc.bytes), func(t *testing.T) {
			compressed, err := ZstdCompressBytes(tc.bytes)
			require.NoError(t, err)
			uncompressed, err := ZstdUncompressBytes(compressed)
			require.NoError(t, err)
			require.Equal(t, tc.bytes, uncompressed)
		})
	}
}

func TestParallelZstdCompression(t *testing.T) {
	const numParallel = 20
	const numBytes = 2 * 1024 * 1024 // 2MB

	for i := 0; i < numParallel; i++ {
		go func() {
			bytes := make([]byte, numBytes)
			_, err := rand.Read(bytes)
			require.NoError(t, err)
			compressedBytes, err := ZstdCompressBytes(bytes)
			require.NoError(t, err)
			uncompressed, err := ZstdUncompressBytes(compressedBytes)
			require.NoError(t, err)
			require.Equal(t, bytes, uncompressed)
		}()
	}
}

// M1 Mac Air
// $ go test -bench=. -run=XXX -v github.com/prodvana/prodvana-public/go/pkg/protohelpers
// goos: linux
// goarch: arm64
// pkg: github.com/prodvana/prodvana-public/go/pkg/protohelpers
// BenchmarkZstdCompress
// BenchmarkZstdCompress-8            27088             43961 ns/op        2981.55 MB/s      303104 B/op          2 allocs/op
// PASS
// ok      github.com/prodvana/prodvana-public/go/pkg/protohelpers   2.708s
func BenchmarkZstdCompress(b *testing.B) {
	const lenToCompress = 128 * 1024 // 128KB
	randBytes := make([]byte, lenToCompress)
	_, err := rand.Read(randBytes)
	require.NoError(b, err)

	// Warmup - make sure the zstdEncoder is initialized
	_, _ = ZstdCompressBytes([]byte{0})

	b.ReportAllocs()
	b.SetBytes(lenToCompress)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err = ZstdCompressBytes(randBytes)
		require.NoError(b, err)
	}
}
