package protoyaml

import (
	"github.com/pkg/errors"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"sigs.k8s.io/yaml"
)

type UnmarshalOptions struct {
	protojson.UnmarshalOptions
}

var defaultUnmarshalOptions = UnmarshalOptions{}

func (o UnmarshalOptions) Unmarshal(bytes []byte, message proto.Message) error {
	jsonBytes, err := yaml.YAMLToJSON(bytes)
	if err != nil {
		return errors.Wrap(err, "failed to convert to json")
	}
	if err := o.UnmarshalOptions.Unmarshal(jsonBytes, message); err != nil {
		return errors.Wrap(err, "failed to parse to proto")
	}
	return nil
}

func Unmarshal(bytes []byte, message proto.Message) error {
	return defaultUnmarshalOptions.Unmarshal(bytes, message)
}

func Marshal(message proto.Message) ([]byte, error) {
	jsonBytes, err := protojson.Marshal(message)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal proto to json")
	}
	yamlBytes, err := yaml.JSONToYAML(jsonBytes)
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert to yaml")
	}
	return yamlBytes, nil
}
