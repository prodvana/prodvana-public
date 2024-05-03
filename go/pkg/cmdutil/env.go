package cmdutil

import (
	"fmt"
	"strings"

	common_config_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/common_config"
)

// Converts a list of VAR=VALUE strings to an env variable map
func ProcessEnvVars(envVars []string, secretEnvVars []string) (map[string]*common_config_pb.EnvValue, error) {
	varMap := make(map[string]*common_config_pb.EnvValue, len(envVars))
	for _, env := range envVars {
		components := strings.SplitN(env, "=", 2)
		if len(components) != 2 {
			return nil, fmt.Errorf("--env must be in VAR=VALUE format. Invalid format: %s", env)
		}
		name := components[0]
		value := &common_config_pb.EnvValue{
			ValueOneof: &common_config_pb.EnvValue_Value{
				Value: components[1],
			},
		}
		varMap[name] = value
	}

	for _, env := range secretEnvVars {
		components := strings.SplitN(env, "=", 2)
		if len(components) != 2 {
			return nil, fmt.Errorf("--env must be in VAR=SECRET_NAME format. Invalid format: %s", env)
		}
		name := components[0]
		value := &common_config_pb.EnvValue{
			ValueOneof: &common_config_pb.EnvValue_RawSecret{
				RawSecret: components[1],
			},
		}
		varMap[name] = value
	}
	return varMap, nil
}
