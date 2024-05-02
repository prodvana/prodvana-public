package protohelpers

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"path/filepath"
	"text/template"

	"github.com/prodvana/prodvana-public/go/pkg/protoyaml"

	go_errors "errors"

	"github.com/klauspost/compress/zstd"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protopath"
	"google.golang.org/protobuf/reflect/protorange"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var (
	ErrUser          = fmt.Errorf("user error")
	ErrTemplateParse = errors.Wrap(ErrUser, "template parse error")
	ErrTemplateExec  = errors.Wrap(ErrUser, "template exec error")
)

type VTMessage interface {
	proto.Message
	StableEqualMessageVT(proto.Message) bool
	EqualMessageVT(proto.Message) bool
}

// faster version of proto.Equal that doesn't involve allocations
func Equal(m1, m2 proto.Message) bool {
	if m1 == nil && m2 == nil {
		return true
	}
	if m1 == nil || m2 == nil {
		return false
	}
	return m1.(VTMessage).EqualMessageVT(m2)
}

const FileTypeInfer = "infer"

func Unmarshal(fileType, fileName string, bytes []byte, message proto.Message, discardUnknown bool) error {
	if fileType == FileTypeInfer {
		switch filepath.Ext(fileName) {
		case ".json":
			fileType = "json"
		case ".yaml", ".yml":
			fileType = "yaml"
		case ".pbtxt":
			fileType = "pbtxt"
		default:
			return errors.Errorf("Don't know how to unmarshal %s", fileName)
		}
	}
	switch fileType {
	case "json":
		err := protojson.UnmarshalOptions{DiscardUnknown: discardUnknown}.Unmarshal(bytes, message)
		if err != nil {
			return errors.Wrap(err, "failed to unmarshal")
		}
		return nil
	case "yaml":
		return protoyaml.UnmarshalOptions{UnmarshalOptions: protojson.UnmarshalOptions{DiscardUnknown: discardUnknown}}.Unmarshal(bytes, message)
	case "pbtxt":
		err := prototext.UnmarshalOptions{DiscardUnknown: discardUnknown}.Unmarshal(bytes, message)
		if err != nil {
			return errors.Wrap(err, "failed to unmarshal")
		}
		return nil
	default:
		return errors.Errorf("Don't know how to unmarshal file type %s", fileType)
	}
}

func ProtoListEqual[T proto.Message](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for idx := range a {
		if !proto.Equal(a[idx], b[idx]) {
			return false
		}
	}
	return true
}

func MarshalBytesGzipped(msg proto.Message) ([]byte, error) {
	marshaled, err := proto.Marshal(msg)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal")
	}

	return GzipBytes(marshaled)
}

func MarshalBytesZstd(msg proto.Message) ([]byte, error) {
	marshaled, err := proto.Marshal(msg)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal")
	}

	return ZstdCompressBytes(marshaled)
}

func MarshalBytesAndJson(msg proto.Message) ([]byte, []byte, error) {
	bytes, err := proto.Marshal(msg)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to marshal")
	}
	jsonBytes, err := protojson.Marshal(msg)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to marshal json")
	}
	return bytes, jsonBytes, nil
}

var (
	// zstdEncoder is a shared zstd.Writer that is used to compress bytes.
	// It is safe to use concurrently (https://github.com/klauspost/compress/blob/master/zstd/README.md#blocks)
	// A lot of time is spent in the initialization of the zstd.Writer, so we keep a single instance around
	// as a performance optimization.
	zstdEncoder, _ = zstd.NewWriter(nil)
)

func ZstdCompressBytes(b []byte) ([]byte, error) {
	return zstdEncoder.EncodeAll(b, nil), nil
}

func GzipBytes(b []byte) ([]byte, error) {
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)

	_, err := zw.Write(b)
	if err != nil {
		return nil, errors.Wrap(err, "failed to write gzip")
	}

	err = zw.Close()
	if err != nil {
		return nil, errors.Wrap(err, "failed to close gzip")
	}
	return buf.Bytes(), nil
}

func UnmarshalGzippedProto(b []byte, msg proto.Message) error {
	uncompressed, err := UngzipBytes(b)
	if err != nil {
		return err
	}

	return proto.Unmarshal(uncompressed, msg)
}

func UnmarshalZstdProto(b []byte, msg proto.Message) error {
	uncompressed, err := ZstdUncompressBytes(b)
	if err != nil {
		return err
	}

	return proto.Unmarshal(uncompressed, msg)
}

func ZstdUncompressBytes(b []byte) (uncompressed []byte, finalErr error) {
	zr, err := zstd.NewReader(bytes.NewReader(b))
	if err != nil {
		return nil, errors.Wrap(err, "failed to make zstd reader")
	}
	defer zr.Close()
	bytes, err := io.ReadAll(zr)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read gzip")
	}
	return bytes, nil
}

func UngzipBytes(b []byte) (uncompressed []byte, finalErr error) {
	zr, err := gzip.NewReader(bytes.NewReader(b))
	if err != nil {
		return nil, errors.Wrap(err, "failed to open gzip")
	}
	defer func() {
		err := zr.Close()
		if finalErr == nil {
			finalErr = errors.Wrap(err, "failed to close gzip")
		}
	}()
	bytes, err := io.ReadAll(zr)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read gzip")
	}
	return bytes, nil
}

func TransformStrings(msg proto.Message, transform func(string) (string, error)) error {
	// mostly copied from https://github.com/prodvana/prodvana/pull/5254
	return protorange.Range(msg.ProtoReflect(), func(p protopath.Values) error {
		last := p.Index(-1)
		switch last.Step.Kind() {
		case protopath.MapIndexStep:
			// transform key if it is a string
			key := last.Step.MapIndex()
			if oldKey, ok := key.Interface().(string); ok {
				newKey, err := transform(oldKey)
				if err != nil {
					return err
				}
				if newKey != oldKey {
					beforeLast := p.Index(-2)
					ms := beforeLast.Value.Map()
					oldValue := ms.Get(key)
					ms.Clear(key)
					newKeyVal := protoreflect.ValueOfString(newKey).MapKey()
					ms.Set(newKeyVal, oldValue)
				}
			}

			// transform value if it is a string
			if oldValue, ok := last.Value.Interface().(string); ok {
				newValue, err := transform(oldValue)
				if err != nil {
					return err
				}
				if newValue != oldValue {
					beforeLast := p.Index(-2)
					ms := beforeLast.Value.Map()
					ms.Set(last.Step.MapIndex(), protoreflect.ValueOfString(newValue))
				}
			}
			return nil
		default:
			var oldS string
			var newS string
			var err error
			var ok bool
			oldS, ok = last.Value.Interface().(string)
			if !ok {
				return nil
			}
			newS, err = transform(oldS)
			if err != nil {
				return err
			}

			if newS == oldS {
				return nil
			}

			// Store the transformed string back into the message.
			beforeLast := p.Index(-2)
			switch last.Step.Kind() {
			case protopath.FieldAccessStep:
				m := beforeLast.Value.Message()
				fd := last.Step.FieldDescriptor()
				m.Set(fd, protoreflect.ValueOfString(newS))
			case protopath.ListIndexStep:
				ls := beforeLast.Value.List()
				i := last.Step.ListIndex()
				ls.Set(i, protoreflect.ValueOfString(newS))
			case protopath.MapIndexStep:
				panic("impossible") // MapIndexStep handled above
			}
			return nil
		}
	})
}

func TemplateExecute(msg proto.Message, templateData map[string]interface{}) error {
	return TransformStrings(msg, func(s string) (string, error) {
		tmpl, err := template.New("").Option("missingkey=error").Parse(s)
		if err != nil {
			return "", errors.Wrap(ErrTemplateParse, "failed to parse template")
		}
		var buffer bytes.Buffer
		if err := tmpl.Execute(&buffer, templateData); err != nil {
			var execErr template.ExecError
			if go_errors.As(err, &execErr) {
				return "", errors.Wrap(ErrTemplateExec, err.Error())
			}
			return "", errors.Wrap(err, "failed to execute template")
		}
		return buffer.String(), nil
	})
}
