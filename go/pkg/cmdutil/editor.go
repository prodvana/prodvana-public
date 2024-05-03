package cmdutil

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/prodvana/prodvana-public/go/pkg/protohelpers"
	"github.com/prodvana/prodvana-public/go/pkg/protoyaml"

	common_config_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/common_config"

	"github.com/google/shlex"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
)

type editOptions[T any] struct {
	rawPostProcess   func([]byte) ([]byte, error)
	typedPostProcess func(T) (T, error)
}

type EditOptions[T any] func(editOptions[T]) editOptions[T]

func WithRawPostProcess[T any](postProcess func([]byte) ([]byte, error)) EditOptions[T] {
	return func(eo editOptions[T]) editOptions[T] {
		eo.rawPostProcess = postProcess
		return eo
	}
}

func WithTypedPostProcess[T any](postProcess func(T) (T, error)) EditOptions[T] {
	return func(eo editOptions[T]) editOptions[T] {
		eo.typedPostProcess = postProcess
		return eo
	}
}

func compileEditOptions[T any](options ...EditOptions[T]) editOptions[T] {
	eo := editOptions[T]{
		rawPostProcess:   func(s []byte) ([]byte, error) { return s, nil },
		typedPostProcess: func(t T) (T, error) { return t, nil },
	}
	for _, o := range options {
		eo = o(eo)
	}
	return eo
}

var Editor string

func LaunchEditor(file string) error {
	editor := Editor
	if editor == "" {
		for _, env := range []string{"EDITOR", "GIT_EDITOR"} {
			editor = os.Getenv(env)
			if editor != "" {
				break
			}
		}
	}
	if editor == "" {
		// TODO(naphat) some default?
		return errors.Errorf("Could not determine default editor")
	}
	args, err := shlex.Split(editor)
	if err != nil {
		return errors.Wrapf(err, "failed to parse editor command: %s", editor)
	}
	args = append(args, file)

	log.Printf("Please edit file in your editor.\n")
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return errors.Wrap(err, "editor error")
	}
	return nil
}

func EditUntilValid[Parsed proto.Message](ctx context.Context, file string, parse func(file string) (Parsed, error), validate func(ctx context.Context, m Parsed) error) (emptyParsed Parsed, err error) {
	prompt := false
	for {
		select {
		case <-ctx.Done():
			return emptyParsed, ctx.Err()
		default:
		}
		if prompt {
			if !PromptYesNo("Edit again?") {
				return emptyParsed, errors.Errorf("Aborted")
			}
		}
		prompt = true
		err := LaunchEditor(file)
		if err != nil {
			return emptyParsed, err
		}
		message, err := parse(file)
		if err != nil {
			fmt.Printf("Parse error: %s\n", err.Error())
			continue
		}
		err = validate(ctx, message)
		if err != nil {
			fmt.Printf("%s\n", err.Error())
			continue
		}
		return message, nil
	}
}

func EditOrReadConfig[ConfigT proto.Message](ctx context.Context, inputFileType, maybeInputFile, startingYaml string, config ConfigT, validate func(context.Context, ConfigT) ([]*common_config_pb.DangerousAction, error), finalPromptMessage string, noPrompt bool, options ...EditOptions[ConfigT]) ([]*common_config_pb.DangerousAction, error) {
	compiledOptions := compileEditOptions(options...)
	var dangerousActions []*common_config_pb.DangerousAction
	if maybeInputFile == "" {
		configFile, err := os.CreateTemp("", "config-*.yaml")
		if err != nil {
			return nil, err
		}
		defer os.Remove(configFile.Name())
		if _, err := configFile.WriteString(startingYaml); err != nil {
			return nil, err
		}
		if err := configFile.Close(); err != nil {
			return nil, err
		}
		config, err = EditUntilValid(ctx, configFile.Name(), func(file string) (ConfigT, error) {
			var defaultConfig ConfigT
			bytes, err := os.ReadFile(file)
			if err != nil {
				return defaultConfig, errors.Wrap(err, "failed to read file")
			}
			bytes, err = compiledOptions.rawPostProcess(bytes)
			if err != nil {
				return defaultConfig, err
			}
			err = protoyaml.Unmarshal(bytes, config)
			if err != nil {
				return defaultConfig, err
			}
			config, err := compiledOptions.typedPostProcess(config)
			if err != nil {
				return defaultConfig, err
			}
			return config, nil
		}, func(ctx context.Context, config ConfigT) error {
			var err error
			dangerousActions, err = validate(ctx, config)
			if err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return nil, err
		}
	} else {
		bytes, err := os.ReadFile(maybeInputFile)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to read %s", maybeInputFile)
		}
		bytes, err = compiledOptions.rawPostProcess(bytes)
		if err != nil {
			return nil, err
		}
		if err := protohelpers.Unmarshal(inputFileType, maybeInputFile, bytes, config, false); err != nil {
			return nil, err
		}
		config, err = compiledOptions.typedPostProcess(config)
		if err != nil {
			return nil, err
		}
		dangerousActions, err = validate(ctx, config)
		if err != nil {
			return nil, err
		}
	}
	// TODO(naphat) config diff
	if !noPrompt {
		if len(dangerousActions) > 0 {
			PromptDangerousActions(dangerousActions)
		} else {
			if !PromptYesNo(finalPromptMessage) {
				return nil, errors.Errorf("Aborted")
			}
		}
	}
	return dangerousActions, nil
}
