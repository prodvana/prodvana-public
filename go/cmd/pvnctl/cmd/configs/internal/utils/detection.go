package utils

import (
	_ "embed"
	go_errors "errors"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/prodvana/prodvana-public/go/pkg/protoyaml"

	"github.com/pkg/errors"
	config_file_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/config_file"
	"github.com/santhosh-tekuri/jsonschema/v5"
	"gopkg.in/yaml.v3"
)

//go:embed schemas/ProdvanaConfig.json
var configSchemaStr string
var compiledConfigSchema = jsonschema.MustCompileString("ProdvanaConfig.json", configSchemaStr)

type ConfigFile struct {
	ConfigFile string
}

var (
	PvnExts = []string{".pvn.yaml", ".pvn.yml"}
)

func ProcessPvnConfigDirectories(args []string) ([]*ConfigFile, error) {
	parsedDirs, err := ProcessConfigDirectories(args, PvnExts)
	if err != nil {
		return nil, err
	}
	if len(parsedDirs) == 0 {
		return nil, errors.Errorf("No Prodvana configs found.")
	}
	return parsedDirs, nil
}

func HasAnySuffix(str string, suffixs []string) bool {
	for _, suffix := range suffixs {
		if strings.HasSuffix(str, suffix) {
			return true
		}
	}
	return false
}

func ProcessConfigDirectories(args []string, exts []string) ([]*ConfigFile, error) {
	type queueItem struct {
		path   string
		noGlob bool
	}
	queue := make([]queueItem, 0, len(args))
	for _, arg := range args {
		queue = append(queue, queueItem{
			path: arg,
		})
	}
	parsedDirs := []*ConfigFile{}
	for len(queue) > 0 {
		last := len(queue) - 1
		item := queue[last]
		queue = queue[:last]

		if HasAnySuffix(item.path, exts) {
			absPath, err := filepath.Abs(item.path)
			if err != nil {
				return nil, errors.Wrapf(err, "failed to get abs path for %s", item.path)
			}
			parsedDirs = append(parsedDirs, &ConfigFile{
				ConfigFile: absPath,
			})
			continue
		}
		if item.noGlob {
			continue
		}
		dir, base := filepath.Split(item.path)
		if dir == "" {
			dir = "."
		}
		if base == "..." {
			err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
				if err != nil {
					return err
				}
				if !d.IsDir() {
					queue = append(queue, queueItem{
						path:   path,
						noGlob: true,
					})
				}
				return nil
			})
			if err != nil {
				return nil, errors.Wrapf(err, "failed to walk dir %s", dir)
			}
			continue
		}
		entries, err := os.ReadDir(item.path)
		if err != nil {
			if go_errors.Is(err, syscall.ENOTDIR) {
				continue
			}
			return nil, errors.Wrapf(err, "failed to read dir %s", item.path)
		}
		for _, entry := range entries {
			queue = append(queue, queueItem{
				path:   filepath.Join(item.path, entry.Name()),
				noGlob: true, // not recursive
			})
		}
	}
	return parsedDirs, nil
}

func ParseConfig(cfgFile *ConfigFile) (*config_file_pb.ProdvanaConfig, error) {
	bytes, err := os.ReadFile(cfgFile.ConfigFile)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read file")
	}
	{
		// validate with json schema first because that gives a much nicer error message than our protoyaml library
		var untyped map[string]interface{}
		if err := yaml.Unmarshal(bytes, &untyped); err != nil {
			return nil, errors.Wrap(err, "failed to unmarshal yaml")
		}
		if err := compiledConfigSchema.Validate(untyped); err != nil {
			return nil, errors.Wrap(err, "schema validation failed")
		}
	}
	var cfg config_file_pb.ProdvanaConfig
	if err := protoyaml.Unmarshal(bytes, &cfg); err != nil {
		return nil, errors.Wrap(err, "failed to parse config file")
	}
	if err := cfg.ValidateAll(); err != nil {
		return nil, errors.Wrap(err, "validation failed")
	}
	if cfg.MetadataOneof != nil {
		switch cfg.ConfigOneof.(type) {
		case *config_file_pb.ProdvanaConfig_Application:
			if cfg.GetApplicationMetadata() == nil {
				return nil, errors.Errorf("Incorrect metadata type for application.")
			}
		case *config_file_pb.ProdvanaConfig_Service:
			if cfg.GetServiceMetadata() == nil {
				return nil, errors.Errorf("Incorrect metadata type for service.")
			}
		default:
			return nil, errors.Errorf("Metadata not supported for %T", cfg.ConfigOneof)
		}
	}
	return &cfg, nil
}
