package configs

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"

	"github.com/prodvana/prodvana-public/go/pkg/cmdutil"

	"github.com/prodvana/prodvana-public/go/cmd/pvnctl/cmd/configs/internal/utils"
	"github.com/prodvana/prodvana-public/go/cmd/pvnctl/cmd/configs/internal/utils/helm"
	"github.com/prodvana/prodvana-public/go/cmd/pvnctl/cmd/configs/internal/utils/kustomize"
	"github.com/prodvana/prodvana-public/go/cmd/pvnctl/cmd/services"

	"github.com/bradenaw/juniper/xslices"
	"github.com/go-git/go-git/v5/plumbing/format/gitignore"
	"github.com/prodvana/prodvana-public/go/pkg/perrors"
	"github.com/prodvana/prodvana-public/go/pkg/protoyaml"
	"github.com/prodvana/prodvana-public/go/pkg/sets"
	"github.com/prodvana/prodvana-public/go/pkg/tarball"
	application_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/application"
	"github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/blobs"
	common_config_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/common_config"
	config_file_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/config_file"
	delivery_extension_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/delivery_extension"
	environment_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/environment"
	fly_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/fly"
	"github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/organization"
	protection_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/protection"
	service_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/service"
	version_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/version"

	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
	grpc_codes "google.golang.org/grpc/codes"
	"google.golang.org/protobuf/proto"

	"github.com/spf13/cobra"
)

var (
	pushFlag     bool
	noPromptFlag bool
	forceFlag    bool
)

func uploadDirectoryAsTgz(ctx context.Context, basePath string) (string, error) {
	stat, err := os.Stat(basePath)
	if err != nil {
		return "", errors.Wrapf(err, "failed to stat %s", basePath)
	}
	if !stat.IsDir() {
		return "", fmt.Errorf("path %s is not a directory", basePath)
	}
	var buffer bytes.Buffer
	err = tarball.ArchiveTgz(basePath, &buffer)
	if err != nil {
		return "", err
	}
	return uploadBytes(ctx, buffer.Bytes())
}

func constructParallelDirectoryWithPatterns(dir string, includePatterns, excludePatterns []string) (string, error) {
	tmpDir, err := os.MkdirTemp("", "pvn-upload")
	if err != nil {
		return "", errors.Wrap(err, "failed to create temp dir")
	}
	var includeMatcher gitignore.Matcher
	if len(includePatterns) > 0 {
		includeMatcher = gitignore.NewMatcher(xslices.Map(includePatterns, func(p string) gitignore.Pattern {
			return gitignore.ParsePattern(p, nil)
		}))
	}
	excludeMatcher := gitignore.NewMatcher(xslices.Map(excludePatterns, func(p string) gitignore.Pattern {
		return gitignore.ParsePattern(p, nil)
	}))

	err = filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			// We create directories if required by the files inside them.
			return nil
		}

		relPath, err := filepath.Rel(dir, path)
		if err != nil {
			return errors.Wrapf(err, "internal error: %s is not a relative path of %s", dir, path)
		}
		relPathComponents := strings.Split(relPath, string(os.PathSeparator))

		// By default, include everything. If includePatterns are specified, only include files that match the patterns.
		if includeMatcher != nil {
			if !includeMatcher.Match(relPathComponents, d.IsDir()) {
				return nil
			}
		}

		if excludeMatcher.Match(relPathComponents, d.IsDir()) {
			return nil
		}

		dstPath := filepath.Join(tmpDir, relPath)

		if err := os.MkdirAll(filepath.Dir(dstPath), 0755); err != nil {
			return errors.Wrapf(err, "failed to create directory %s", filepath.Dir(dstPath))
		}

		// log.Printf("Uploading file %s", relPath)
		f, err := os.Create(dstPath)
		if err != nil {
			return errors.Wrapf(err, "failed to create %s", dstPath)
		}

		src, err := os.Open(path)
		if err != nil {
			return errors.Wrapf(err, "failed to open %s", path)
		}

		_, err = io.Copy(f, src)
		if err != nil {
			return errors.Wrapf(err, "failed to copy %s to %s", path, dstPath)
		}

		return f.Close()
	})
	if err != nil {
		return "", err
	}

	return tmpDir, nil
}

func uploadBytes(ctx context.Context, bytes []byte) (string, error) {
	strm, err := cmdutil.GetBlobsManagerClient().UploadCasBlob(ctx)
	if err != nil {
		return "", err
	}

	const chunkSizeBytes = 4 * 1024 * 1024 // 4mb
	for i := 0; i < len(bytes); i += chunkSizeBytes {
		end := i + chunkSizeBytes
		if end > len(bytes) {
			end = len(bytes)
		}
		err := strm.Send(&blobs.UploadCasBlobReq{
			Bytes: bytes[i:end],
		})
		if err != nil {
			return "", errors.Wrap(err, "failed to upload tarball")
		}
	}
	resp, err := strm.CloseAndRecv()
	if err != nil {
		return "", errors.Wrap(err, "failed to close send")
	}
	return resp.Id, nil
}

type remoteOrInlinedConfig struct {
	inlined string
	remote  *common_config_pb.RemoteConfig
}

func processLocalConfig(ctx context.Context, dir string, local *common_config_pb.LocalConfig) (*remoteOrInlinedConfig, error) {
	if local.SubPath != "" {
		directoryToUpload := local.GetPath()
		if directoryToUpload == "" {
			return nil, fmt.Errorf("path must be set when using subPath")
		}
		directoryToUpload = filepath.Join(dir, directoryToUpload)
		prunedDir, err := constructParallelDirectoryWithPatterns(directoryToUpload, local.IncludePatterns, local.ExcludePatterns)
		defer os.RemoveAll(prunedDir)
		if err != nil {
			return nil, err
		}
		blobId, err := uploadDirectoryAsTgz(ctx, prunedDir)
		if err != nil {
			return nil, err
		}
		return &remoteOrInlinedConfig{
			remote: &common_config_pb.RemoteConfig{
				RemoteOneof: &common_config_pb.RemoteConfig_TarballBlobId{
					TarballBlobId: blobId,
				},
				SubPath: local.SubPath,
			},
		}, nil
	} else {
		var fullPath string
		switch innerLocal := local.PathOneof.(type) {
		case *common_config_pb.LocalConfig_Path:
			fullPath = filepath.Join(dir, innerLocal.Path)
		default:
			return nil, fmt.Errorf("unknown local path %+v", local.PathOneof)
		}
		bytes, err := os.ReadFile(fullPath)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to read %s", fullPath)
		}
		return &remoteOrInlinedConfig{
			inlined: string(bytes),
		}, nil
	}
}

// Given a directory, run kustomize on all directories with a kustomize file, and construct a parallel directory structure
// with the compiled files.
// Caller is responsible for cleaning up the parallel directory when done.
func preprocessKustomizeConfigs(ctx context.Context, dir, base string, excludePatterns []string) (string, error) {
	absBase := filepath.Join(dir, base)
	tmpDir, err := os.MkdirTemp("", "pvn-kustomize")
	if err != nil {
		return "", errors.Wrap(err, "failed to create temp dir")
	}
	patterns := xslices.Map(excludePatterns, func(p string) gitignore.Pattern {
		return gitignore.ParsePattern(p, nil)
	})
	excludeMatcher := gitignore.NewMatcher(patterns)
	matchedDirs := sets.NewSet[string]()
	err = filepath.WalkDir(absBase, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		dir, file := filepath.Split(path)
		switch file {
		case "kustomization.yaml", "kustomization.yml", "Kustomization":
			relPath, err := filepath.Rel(absBase, path)
			if err != nil {
				return errors.Wrapf(err, "internal error: %s is not a relative path of %s", absBase, path)
			}
			relPathComponents := strings.Split(relPath, string(os.PathSeparator))
			if !excludeMatcher.Match(relPathComponents, d.IsDir()) {
				matchedDirs.Add(dir)
			}
		}
		return nil
	})
	if err != nil {
		return "", err
	}
	err = matchedDirs.Iterate(func(kustomizeDir string) error {
		relPath, err := filepath.Rel(absBase, kustomizeDir)
		if err != nil {
			return errors.Wrapf(err, "internal error: %s is not a relative path of %s", absBase, kustomizeDir)
		}
		newDir := filepath.Join(tmpDir, relPath)
		if err := os.MkdirAll(newDir, 0755); err != nil {
			return errors.Wrap(err, "failed to mkdir")
		}
		cfg, err := kustomize.GetRawConfig(ctx, kustomizeDir)
		if err != nil {
			return err
		}
		return errors.Wrap(os.WriteFile(filepath.Join(newDir, "compiled.yaml"), cfg, 0755), "failed to write file")
	})
	if err != nil {
		return "", err
	}
	return tmpDir, nil
}

func expandKubernetesConfig(ctx context.Context, dir string, extConfig *common_config_pb.KubernetesConfig) (*common_config_pb.KubernetesConfig, error) {
	switch source := extConfig.SourceOneof.(type) {
	case *common_config_pb.KubernetesConfig_Local:
		if source.Local.SubPath != "" {
			basePath := source.Local.GetPath()
			if basePath == "" {
				return nil, fmt.Errorf("path must be set when using subPath")
			}
			var directoryToUpload string
			switch extConfig.Type {
			case common_config_pb.KubernetesConfig_KUBERNETES:
				directoryToUpload = filepath.Join(dir, basePath)
			case common_config_pb.KubernetesConfig_KUSTOMIZE:
				tmpDir, err := preprocessKustomizeConfigs(ctx, dir, basePath, source.Local.ExcludePatterns)
				if err != nil {
					return nil, err
				}
				defer os.RemoveAll(tmpDir)
				directoryToUpload = tmpDir
			default:
				return nil, fmt.Errorf("unimplemented config type %v", extConfig.Type)
			}
			blobId, err := uploadDirectoryAsTgz(ctx, directoryToUpload)
			if err != nil {
				return nil, err
			}
			cfg := proto.Clone(extConfig).(*common_config_pb.KubernetesConfig)
			cfg.Type = common_config_pb.KubernetesConfig_KUBERNETES // compiled down to raw k8s config
			cfg.SourceOneof = &common_config_pb.KubernetesConfig_Remote{
				Remote: &common_config_pb.RemoteConfig{
					RemoteOneof: &common_config_pb.RemoteConfig_TarballBlobId{
						TarballBlobId: blobId,
					},
					SubPath: source.Local.SubPath,
				},
			}
			return cfg, nil
		} else {
			var dirOrFiles []string
			if source.Local.PathOneof != nil && len(source.Local.Paths) > 0 {
				return nil, fmt.Errorf("cannot specify both `paths` and `path`")
			}
			if source.Local.PathOneof != nil {
				switch inner := source.Local.PathOneof.(type) {
				case *common_config_pb.LocalConfig_Path:
					dirOrFiles = append(dirOrFiles, filepath.Join(dir, inner.Path))
				default:
					return nil, fmt.Errorf("unknown local path %+v", source.Local.PathOneof)
				}
			} else {
				for _, path := range source.Local.Paths {
					dirOrFiles = append(dirOrFiles, filepath.Join(dir, path))
				}
			}
			switch extConfig.Type {
			case common_config_pb.KubernetesConfig_KUBERNETES:
				configFiles, err := utils.ProcessConfigDirectories(dirOrFiles, []string{".yaml", ".yml"})
				if err != nil {
					return nil, err
				}
				var yamlFiles []string
				for _, file := range configFiles {
					if utils.HasAnySuffix(file.ConfigFile, utils.PvnExts) {
						// special case pvn config file
						continue
					}
					bytes, err := os.ReadFile(file.ConfigFile)
					if err != nil {
						return nil, errors.Wrapf(err, "failed to read %s", file.ConfigFile)
					}
					yamlFiles = append(yamlFiles, string(bytes))
				}
				if len(yamlFiles) == 0 {
					return nil, errors.Errorf("No kubernetes configs found at %s", source.Local.GetPath())
				}
				cfg := proto.Clone(extConfig).(*common_config_pb.KubernetesConfig)
				cfg.Type = common_config_pb.KubernetesConfig_KUBERNETES // compiled down to raw k8s config
				cfg.SourceOneof = &common_config_pb.KubernetesConfig_Inlined{
					Inlined: strings.Join(yamlFiles, "---\n"),
				}
				return cfg, nil
			case common_config_pb.KubernetesConfig_KUSTOMIZE:
				var buffer bytes.Buffer
				for _, dirOrFile := range dirOrFiles {
					bytes, err := kustomize.GetRawConfig(ctx, dirOrFile)
					if err != nil {
						return nil, err
					}
					buffer.Write(bytes)
					buffer.WriteString("---\n")
				}
				cfg := proto.Clone(extConfig).(*common_config_pb.KubernetesConfig)
				cfg.Type = common_config_pb.KubernetesConfig_KUBERNETES
				cfg.SourceOneof = &common_config_pb.KubernetesConfig_Inlined{
					Inlined: buffer.String(),
				}
				return cfg, nil
			default:
				panic("not implemented")
			}
		}
	}
	return extConfig, nil
}

func expandHelmConfig(ctx context.Context, dir string, helmConfig *common_config_pb.HelmConfig) (*common_config_pb.HelmConfig, error) {
	switch inner := helmConfig.ChartOneof.(type) {
	case *common_config_pb.HelmConfig_Local:
		var fullPath string
		switch innerLocal := inner.Local.PathOneof.(type) {
		case *common_config_pb.LocalConfig_Path:
			fullPath = filepath.Join(dir, innerLocal.Path)
		default:
			return nil, fmt.Errorf("unknown local path %+v", inner.Local.PathOneof)
		}

		chartTgz, done, err := helm.Package(ctx, fullPath)
		if err != nil {
			return nil, err
		}
		defer done()
		chartBytes, err := os.ReadFile(chartTgz)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to read %s", chartTgz)
		}
		blobId, err := uploadBytes(ctx, chartBytes)
		if err != nil {
			return nil, err
		}

		helmConfig.ChartOneof = &common_config_pb.HelmConfig_HelmTarballBlobId{
			HelmTarballBlobId: blobId,
		}
	}
	for _, val := range helmConfig.ValuesOverrides {
		switch inner := val.OverrideOneof.(type) {
		case *common_config_pb.HelmValuesOverrides_Local:
			res, err := processLocalConfig(ctx, dir, inner.Local)
			if err != nil {
				return nil, err
			}
			if res.remote == nil {
				val.OverrideOneof = &common_config_pb.HelmValuesOverrides_Inlined{
					Inlined: res.inlined,
				}
			} else {
				val.OverrideOneof = &common_config_pb.HelmValuesOverrides_Remote{
					Remote: res.remote,
				}
			}
		}
	}
	return helmConfig, nil
}

func expandGoogleCloudRunConfig(ctx context.Context, dir string, cloudRunConfig *service_pb.GoogleCloudRunConfig) (*service_pb.GoogleCloudRunConfig, error) {
	switch inner := cloudRunConfig.SpecOneof.(type) {
	case *service_pb.GoogleCloudRunConfig_Local:
		res, err := processLocalConfig(ctx, dir, inner.Local)
		if err != nil {
			return nil, err
		}
		if res.remote == nil {
			cloudRunConfig.SpecOneof = &service_pb.GoogleCloudRunConfig_Inlined{
				Inlined: res.inlined,
			}
		} else {
			cloudRunConfig.SpecOneof = &service_pb.GoogleCloudRunConfig_Remote{
				Remote: res.remote,
			}
		}
	}
	return cloudRunConfig, nil
}

func expandFlyConfig(ctx context.Context, dir string, flyCfg *fly_pb.FlyConfig) (*fly_pb.FlyConfig, error) {
	switch inner := flyCfg.TomlOneof.(type) {
	case *fly_pb.FlyConfig_Local:
		res, err := processLocalConfig(ctx, dir, inner.Local)
		if err != nil {
			return nil, err
		}
		if res.remote == nil {
			flyCfg.TomlOneof = &fly_pb.FlyConfig_Inlined{
				Inlined: res.inlined,
			}
		} else {
			flyCfg.TomlOneof = &fly_pb.FlyConfig_Remote{
				Remote: res.remote,
			}
		}
	}
	return flyCfg, nil
}

func expandEcsSpec(ctx context.Context, dir string, spec *service_pb.AwsEcsConfig_Spec) error {
	if spec == nil {
		return nil
	}
	switch inner := spec.SpecOneof.(type) {
	case *service_pb.AwsEcsConfig_Spec_Local:
		res, err := processLocalConfig(ctx, dir, inner.Local)
		if err != nil {
			return err
		}
		if res.remote == nil {
			spec.SpecOneof = &service_pb.AwsEcsConfig_Spec_Inlined{
				Inlined: res.inlined,
			}
		} else {
			spec.SpecOneof = &service_pb.AwsEcsConfig_Spec_Remote{
				Remote: res.remote,
			}
		}
	}
	return nil
}

func expandEcsConfig(ctx context.Context, dir string, ecsConfig *service_pb.AwsEcsConfig) (*service_pb.AwsEcsConfig, error) {
	err := expandEcsSpec(ctx, dir, ecsConfig.TaskDefinition)
	if err != nil {
		return nil, err
	}
	err = expandEcsSpec(ctx, dir, ecsConfig.ServiceDefinition)
	if err != nil {
		return nil, err
	}
	return ecsConfig, nil
}

func showDiff(ctx context.Context, oldVersionName string, oldConfig proto.Message, newConfig proto.Message) error {
	oldConfigYAML, err := protoyaml.Marshal(oldConfig)
	if err != nil {
		return errors.Errorf("Unable to marshal existing config to diff against: %v", err)
	} else {
		newConfigYAML, err := protoyaml.Marshal(newConfig)
		if err != nil {
			return errors.Errorf("Unable to marshal new config to diff against: %v", err)
		} else {
			oldConfigYAMLStr := string(oldConfigYAML)
			newConfigYAMLStr := string(newConfigYAML)
			if oldConfigYAMLStr == newConfigYAMLStr {
				fmt.Println("No differences found")
			} else {
				tmpDir, err := os.MkdirTemp("", "pvnctl-diff")
				if err != nil {
					return errors.Wrap(err, "failed to make tmpdir")
				}
				defer os.RemoveAll(tmpDir)
				oldAbsFileName := filepath.Join(tmpDir, "old-version.yaml") // do not use oldVersionName here because it can contain /
				err = os.WriteFile(oldAbsFileName, oldConfigYAML, 0644)
				if err != nil {
					return errors.Wrapf(err, "failed to write %s", oldAbsFileName)
				}
				newAbsFileName := filepath.Join(tmpDir, "new-version.yaml")
				err = os.WriteFile(newAbsFileName, newConfigYAML, 0644)
				if err != nil {
					return errors.Wrapf(err, "failed to write %s", newAbsFileName)
				}
				cmd := exec.Command("diff", "-u", "--label", oldVersionName, "--label", "new-version.yaml", oldAbsFileName, newAbsFileName)

				cmd.Stdin = os.Stdin
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				err = cmd.Run()

				if err != nil {
					if exitErr, ok := err.(*exec.ExitError); ok {
						if exitErr.ExitCode() == 127 { // not found
							fmt.Println("Could not find the builtin diff binary. Please install diff to enable generating config diffs. Continuing...")
						} else if exitErr.ExitCode() != 1 {
							return err // diff command returns exit code 1 if differences are found
						}
					}
				}
			}
		}
	}
	return nil
}

func expandConvergenceExtension(ctx context.Context, cfgFile *utils.ConfigFile, cfg *service_pb.DeliveryExtensionConfig) error {
	switch inner := cfg.Definition.(type) {
	case *service_pb.DeliveryExtensionConfig_Inlined:
		err := expandDeliveryExtensionConfig(ctx, cfgFile, inner.Inlined)
		if err != nil {
			return err
		}
	}
	for _, dep := range cfg.Dependencies {
		err := expandConvergenceExtension(ctx, cfgFile, dep)
		if err != nil {
			return err
		}
	}
	return nil
}

func expandConvergenceExtensionInstance(ctx context.Context, cfgFile *utils.ConfigFile, cfg *service_pb.DeliveryExtensionInstance) error {
	switch inner := cfg.Definition.(type) {
	case *service_pb.DeliveryExtensionInstance_Inlined:
		err := expandDeliveryExtensionConfig(ctx, cfgFile, inner.Inlined)
		if err != nil {
			return err
		}
	}
	for _, dep := range cfg.Dependencies {
		err := expandConvergenceExtension(ctx, cfgFile, dep)
		if err != nil {
			return err
		}
	}
	return nil
}

func processServiceConfig(ctx context.Context, cfgFile *utils.ConfigFile, cfg *service_pb.ServiceConfig, meta *service_pb.ServiceUserMetadata, validateOnly bool) error {
	switch inner := cfg.ConfigOneof.(type) {
	case *service_pb.ServiceConfig_ExternalConfig:
		extConfig, err := expandKubernetesConfig(ctx, filepath.Dir(cfgFile.ConfigFile), inner.ExternalConfig)
		if err != nil {
			return err
		}
		cfg.ConfigOneof = &service_pb.ServiceConfig_KubernetesConfig{
			KubernetesConfig: extConfig,
		}
	case *service_pb.ServiceConfig_KubernetesConfig:
		extConfig, err := expandKubernetesConfig(ctx, filepath.Dir(cfgFile.ConfigFile), inner.KubernetesConfig)
		if err != nil {
			return err
		}
		cfg.ConfigOneof = &service_pb.ServiceConfig_KubernetesConfig{
			KubernetesConfig: extConfig,
		}
	case *service_pb.ServiceConfig_Helm:
		helmConfig, err := expandHelmConfig(ctx, filepath.Dir(cfgFile.ConfigFile), inner.Helm)
		if err != nil {
			return err
		}
		cfg.ConfigOneof = &service_pb.ServiceConfig_Helm{
			Helm: helmConfig,
		}
	case *service_pb.ServiceConfig_AwsEcs:
		ecsConfig, err := expandEcsConfig(ctx, filepath.Dir(cfgFile.ConfigFile), inner.AwsEcs)
		if err != nil {
			return err
		}
		cfg.ConfigOneof = &service_pb.ServiceConfig_AwsEcs{
			AwsEcs: ecsConfig,
		}
	case *service_pb.ServiceConfig_GoogleCloudRun:
		cloudRunConfig, err := expandGoogleCloudRunConfig(ctx, filepath.Dir(cfgFile.ConfigFile), inner.GoogleCloudRun)
		if err != nil {
			return err
		}
		cfg.ConfigOneof = &service_pb.ServiceConfig_GoogleCloudRun{
			GoogleCloudRun: cloudRunConfig,
		}
	case *service_pb.ServiceConfig_Fly:
		flyCfg, err := expandFlyConfig(ctx, filepath.Dir(cfgFile.ConfigFile), inner.Fly)
		if err != nil {
			return err
		}
		cfg.ConfigOneof = &service_pb.ServiceConfig_Fly{
			Fly: flyCfg,
		}
	}
	for _, ext := range cfg.DeliveryExtensions {
		err := expandConvergenceExtension(ctx, cfgFile, ext)
		if err != nil {
			return err
		}
	}
	for _, ext := range cfg.ConvergenceExtensions {
		err := expandConvergenceExtension(ctx, cfgFile, ext)
		if err != nil {
			return err
		}
	}
	for _, ext := range cfg.DeliveryExtensionInstances {
		err := expandConvergenceExtensionInstance(ctx, cfgFile, ext)
		if err != nil {
			return err
		}
	}
	for _, ext := range cfg.ConvergenceExtensionInstances {
		err := expandConvergenceExtensionInstance(ctx, cfgFile, ext)
		if err != nil {
			return err
		}
	}
	for _, perRc := range cfg.PerReleaseChannel {
		switch inner := perRc.ConfigOneof.(type) {
		case *service_pb.PerReleaseChannelConfig_ExternalConfig:
			extConfig, err := expandKubernetesConfig(ctx, filepath.Dir(cfgFile.ConfigFile), inner.ExternalConfig)
			if err != nil {
				return err
			}
			perRc.ConfigOneof = &service_pb.PerReleaseChannelConfig_KubernetesConfig{
				KubernetesConfig: extConfig,
			}
		case *service_pb.PerReleaseChannelConfig_KubernetesConfig:
			extConfig, err := expandKubernetesConfig(ctx, filepath.Dir(cfgFile.ConfigFile), inner.KubernetesConfig)
			if err != nil {
				return err
			}
			perRc.ConfigOneof = &service_pb.PerReleaseChannelConfig_KubernetesConfig{
				KubernetesConfig: extConfig,
			}
		case *service_pb.PerReleaseChannelConfig_Helm:
			helmConfig, err := expandHelmConfig(ctx, filepath.Dir(cfgFile.ConfigFile), inner.Helm)
			if err != nil {
				return err
			}
			perRc.ConfigOneof = &service_pb.PerReleaseChannelConfig_Helm{
				Helm: helmConfig,
			}
		case *service_pb.PerReleaseChannelConfig_AwsEcs:
			ecsConfig, err := expandEcsConfig(ctx, filepath.Dir(cfgFile.ConfigFile), inner.AwsEcs)
			if err != nil {
				return err
			}
			perRc.ConfigOneof = &service_pb.PerReleaseChannelConfig_AwsEcs{
				AwsEcs: ecsConfig,
			}
		case *service_pb.PerReleaseChannelConfig_GoogleCloudRun:
			cloudRunConfig, err := expandGoogleCloudRunConfig(ctx, filepath.Dir(cfgFile.ConfigFile), inner.GoogleCloudRun)
			if err != nil {
				return err
			}
			perRc.ConfigOneof = &service_pb.PerReleaseChannelConfig_GoogleCloudRun{
				GoogleCloudRun: cloudRunConfig,
			}
		case *service_pb.PerReleaseChannelConfig_Fly:
			flyCfg, err := expandFlyConfig(ctx, filepath.Dir(cfgFile.ConfigFile), inner.Fly)
			if err != nil {
				return err
			}
			perRc.ConfigOneof = &service_pb.PerReleaseChannelConfig_Fly{
				Fly: flyCfg,
			}
		}
		for _, ext := range perRc.DeliveryExtensions {
			err := expandConvergenceExtension(ctx, cfgFile, ext)
			if err != nil {
				return err
			}
		}
		for _, ext := range perRc.ConvergenceExtensions {
			err := expandConvergenceExtension(ctx, cfgFile, ext)
			if err != nil {
				return err
			}
		}
	}
	if cfg.Application == "" {
		return errors.Errorf("Must specify application in service config file.")
	}

	validationResp, err := cmdutil.GetServiceManagerClient().ValidateConfigureService(ctx, &service_pb.ConfigureServiceReq{
		Application:           cfg.Application,
		ServiceConfig:         cfg,
		SkipRuntimeValidation: !validateOnly, // optimization, since ConfigureService itself will do runtime validation
	})
	if err != nil {
		return errors.Wrapf(err, "validation failed")
	}
	if validateOnly {
		return nil
	}

	err = handleConfigDiffPrompt(ctx, "service", validationResp.CompiledConfig, func(ctx context.Context) (string, proto.Message, error) {
		getLatestConfigResp, err := cmdutil.GetServiceManagerClient().GetServiceConfig(ctx, &service_pb.GetServiceConfigReq{
			Application: cfg.Application,
			Service:     cfg.Name,
		})
		if err != nil {
			return "", nil, err
		}
		return getLatestConfigResp.ConfigVersion, getLatestConfigResp.CompiledConfig, nil
	})
	if err != nil {
		return err
	}

	resp, err := cmdutil.GetServiceManagerClient().ConfigureService(ctx, &service_pb.ConfigureServiceReq{
		Application:           cfg.Application,
		ServiceConfig:         cfg,
		Source:                version_pb.Source_CONFIG_FILE,
		SourceMetadata:        utils.DetectSourceMetadata(cfgFile),
		ForceCreateNewVersion: forceFlag,
	})
	if err != nil {
		return errors.Wrap(err, "failed to configure service")
	} else {
		if resp.CreatedNewVersion {
			fmt.Printf("Successfully configured %s\n", resp.ConfigVersion)
			orgResp, err := cmdutil.GetOrganizationManagerClient().GetOrganization(ctx, &organization.GetOrganizationReq{})
			if err != nil {
				return errors.Wrap(err, "failed to get organization")
			}
			fmt.Printf("To update service to latest configuration, create a new release with the new configuration at %s/applications/%s/services/%s/desired-state-editor/latest\n", orgResp.Organization.UiAddress, cfg.Application, cfg.Name)
		} else {
			fmt.Printf("Reused current latest configuration %s\n", resp.ConfigVersion)
		}
	}

	if pushFlag {
		resp, err := cmdutil.GetServiceManagerClient().ApplyParameters(ctx, &service_pb.ApplyParametersReq{
			Oneof: &service_pb.ApplyParametersReq_ServiceConfigVersion{
				ServiceConfigVersion: &service_pb.ServiceConfigVersionReference{
					Application:          cfg.Application,
					Service:              resp.ServiceId,
					ServiceConfigVersion: resp.ConfigVersion,
				},
			},
			Application:    cfg.Application,
			Source:         version_pb.Source_CONFIG_FILE,
			SourceMetadata: utils.DetectSourceMetadata(cfgFile),
		})
		if err != nil {
			return errors.Wrap(err, "failed to apply parameters")
		}
		err = services.PushService(ctx, cfg.Application, cfg.Name, resp.Version)
		if err != nil {
			return err
		}
	}
	if meta != nil {
		_, err = cmdutil.GetServiceManagerClient().SetServiceMetadata(ctx,
			&service_pb.SetServiceMetadataReq{
				Application: cfg.Application,
				Service:     resp.ServiceId,
				Metadata:    meta,
			},
		)
		if err != nil {
			return errors.Wrap(err, "failed to set service metadata")
		}
	}
	return nil
}

func handleConfigDiffPrompt(ctx context.Context, noun string, newConfig proto.Message, getOldConfig func(context.Context) (string, proto.Message, error)) error {
	oldVersion, oldConfig, err := getOldConfig(ctx)
	if err != nil {
		isNew := perrors.GrpcErrCode(err) == grpc_codes.NotFound
		if isNew {
			fmt.Printf("This is a new %s. No config diff to show.\n", noun)
			if !noPromptFlag {
				proceed := cmdutil.PromptYesNo("Proceed?")
				if !proceed {
					return errors.Errorf("Aborted")
				}
			}
		} else {
			if noPromptFlag {
				fmt.Printf("Unable to fetch existing config to diff against: %v\n", err)
			} else {
				proceed := cmdutil.PromptYesNo(fmt.Sprintf("Unable to fetch existing config to diff against: %v\n\n, proceed anyway?", err))
				if !proceed {
					return errors.Errorf("Aborted")
				}
			}
		}
	} else {
		err = showDiff(ctx, fmt.Sprintf("%s.yaml", oldVersion), oldConfig, newConfig)
		if err != nil {
			if noPromptFlag {
				fmt.Printf("Unable to generate config diff: %v\n", err)
			} else {
				proceed := cmdutil.PromptYesNo(fmt.Sprintf("Unable to generate config diff: %v\n\nProceed anyway?", err))
				if !proceed {
					return errors.Errorf("Aborted")
				}
			}
		} else {
			if !noPromptFlag {
				proceed := cmdutil.PromptYesNo("Proceed?")
				if !proceed {
					return errors.Errorf("Aborted")
				}
			}
		}
	}
	return nil
}

func processApplicationConfig(ctx context.Context, cfgFile *utils.ConfigFile, cfg *application_pb.ApplicationConfig, meta *application_pb.ApplicationUserMetadata, validateOnly bool) error {
	validationResp, err := cmdutil.GetApplicationManagerClient().ValidateConfigureApplication(ctx, &application_pb.ConfigureApplicationReq{
		ApplicationConfig: cfg,
	})
	if err != nil {
		return errors.Wrapf(err, "validation failed")
	}
	if validateOnly {
		return nil
	}
	var approvedIds []string
	for _, action := range validationResp.DangerousActions {
		approvedIds = append(approvedIds, action.Id)
	}

	err = handleConfigDiffPrompt(ctx, "application", validationResp.CompiledConfig, func(ctx context.Context) (string, proto.Message, error) {
		getLatestConfigResp, err := cmdutil.GetApplicationManagerClient().GetApplicationConfig(ctx, &application_pb.GetApplicationConfigReq{
			Application: cfg.Name,
		})
		if err != nil {
			return "", nil, err
		}
		return getLatestConfigResp.Version, getLatestConfigResp.CompiledConfig, nil
	})
	if err != nil {
		return err
	}

	resp, err := cmdutil.GetApplicationManagerClient().ConfigureApplication(ctx,
		&application_pb.ConfigureApplicationReq{
			ApplicationConfig: cfg,
			// TODO(naphat) what's the appropriate flow here to avoid dangerous actions?
			ApprovedDangerousActionIds: approvedIds,
			Source:                     version_pb.Source_CONFIG_FILE,
			SourceMetadata:             utils.DetectSourceMetadata(cfgFile),
			ForceCreateNewVersion:      forceFlag,
		},
	)
	if err != nil {
		return errors.Wrap(err, "failed to configure application")
	} else {
		if resp.CreatedNewVersion {
			fmt.Printf("Successfully configured %s\n", resp.Meta.Version)
		} else {
			fmt.Printf("Reused current latest configuration %s\n", resp.Meta.Version)
		}
	}
	if meta != nil {
		_, err = cmdutil.GetApplicationManagerClient().SetApplicationMetadata(ctx,
			&application_pb.SetApplicationMetadataReq{
				Application: resp.Meta.Id,
				Metadata:    meta,
			},
		)
		if err != nil {
			return errors.Wrap(err, "failed to set application metadata")
		}
	}
	return nil
}

func processProtectionConfig(ctx context.Context, cfgFile *utils.ConfigFile, cfg *protection_pb.ProtectionConfig, validateOnly bool) error {
	switch inner := cfg.ExecConfig.(type) {
	case *protection_pb.ProtectionConfig_KubernetesConfig:
		extConfig, err := expandKubernetesConfig(ctx, filepath.Dir(cfgFile.ConfigFile), inner.KubernetesConfig)
		if err != nil {
			return err
		}
		cfg.ExecConfig = &protection_pb.ProtectionConfig_KubernetesConfig{
			KubernetesConfig: extConfig,
		}
	}

	validationResp, err := cmdutil.GetProtectionManagerClient().ValidateConfigureProtection(ctx, &protection_pb.ConfigureProtectionReq{
		ProtectionConfig: cfg,
	})
	if err != nil {
		return err
	}
	if validateOnly {
		return nil
	}
	err = handleConfigDiffPrompt(ctx, "protection", validationResp.CompiledConfig, func(ctx context.Context) (string, proto.Message, error) {
		resp, err := cmdutil.GetProtectionManagerClient().GetProtectionConfig(ctx, &protection_pb.GetProtectionConfigReq{
			Protection: cfg.Name,
		})
		if err != nil {
			return "", nil, err
		}
		return resp.Version, resp.CompiledConfig, nil
	})
	if err != nil {
		return err
	}

	resp, err := cmdutil.GetProtectionManagerClient().ConfigureProtection(ctx,
		&protection_pb.ConfigureProtectionReq{
			ProtectionConfig: cfg,
			Source:           version_pb.Source_CONFIG_FILE,
			SourceMetadata:   utils.DetectSourceMetadata(cfgFile),
		},
	)

	if err != nil {
		return errors.Wrap(err, "failed to configure protection")
	} else {
		fmt.Printf("Successfully configured %s\n", resp.Version)
	}
	return nil
}

func expandDeliveryExtensionConfig(ctx context.Context, cfgFile *utils.ConfigFile, cfg *delivery_extension_pb.DeliveryExtensionConfig) error {
	switch inner := cfg.ExecConfig.(type) {
	case *delivery_extension_pb.DeliveryExtensionConfig_KubernetesConfig:
		extConfig, err := expandKubernetesConfig(ctx, filepath.Dir(cfgFile.ConfigFile), inner.KubernetesConfig)
		if err != nil {
			return err
		}
		cfg.ExecConfig = &delivery_extension_pb.DeliveryExtensionConfig_KubernetesConfig{
			KubernetesConfig: extConfig,
		}
	}
	return nil
}

func processDeliveryExtensionConfig(ctx context.Context, cfgFile *utils.ConfigFile, cfg *delivery_extension_pb.DeliveryExtensionConfig, validateOnly bool) error {
	if err := expandDeliveryExtensionConfig(ctx, cfgFile, cfg); err != nil {
		return err
	}

	validationResp, err := cmdutil.GetDeliveryExtensionManagerClient().ValidateConfigureDeliveryExtension(ctx, &delivery_extension_pb.ConfigureDeliveryExtensionReq{
		DeliveryExtensionConfig: cfg,
	})
	if err != nil {
		return err
	}
	if validateOnly {
		return nil
	}
	err = handleConfigDiffPrompt(ctx, "delivery extension", validationResp.CompiledConfig, func(ctx context.Context) (string, proto.Message, error) {
		resp, err := cmdutil.GetDeliveryExtensionManagerClient().GetDeliveryExtensionConfig(ctx, &delivery_extension_pb.GetDeliveryExtensionConfigReq{
			DeliveryExtension: cfg.Name,
		})
		if err != nil {
			return "", nil, err
		}
		return resp.Version, resp.CompiledConfig, nil
	})
	if err != nil {
		return err
	}

	resp, err := cmdutil.GetDeliveryExtensionManagerClient().ConfigureDeliveryExtension(ctx,
		&delivery_extension_pb.ConfigureDeliveryExtensionReq{
			DeliveryExtensionConfig: cfg,
			Source:                  version_pb.Source_CONFIG_FILE,
			SourceMetadata:          utils.DetectSourceMetadata(cfgFile),
		},
	)
	if err != nil {
		return errors.Wrap(err, "failed to configure delivery extension")
	} else {
		fmt.Printf("Successfully configured %s\n", resp.Version)
	}
	return nil
}

func expandRuntimeExtensionApplyCommand(ctx context.Context, cfgFile *utils.ConfigFile, cfg *environment_pb.ExtensionApplyCommand) error {
	if cfg == nil {
		return nil
	}
	switch inner := cfg.ExecConfig.(type) {
	case *environment_pb.ExtensionApplyCommand_KubernetesConfig:
		expanded, err := expandKubernetesConfig(ctx, filepath.Dir(cfgFile.ConfigFile), inner.KubernetesConfig)
		if err != nil {
			return err
		}
		inner.KubernetesConfig = expanded
	}
	return nil
}

func expandRuntimeExtensionFetchCommand(ctx context.Context, cfgFile *utils.ConfigFile, cfg *environment_pb.ExtensionFetchCommand) error {
	if cfg == nil {
		return nil
	}
	switch inner := cfg.ExecConfig.(type) {
	case *environment_pb.ExtensionFetchCommand_KubernetesConfig:
		expanded, err := expandKubernetesConfig(ctx, filepath.Dir(cfgFile.ConfigFile), inner.KubernetesConfig)
		if err != nil {
			return err
		}
		inner.KubernetesConfig = expanded
	}
	return nil
}

func expandRuntimeExtensionGetInfoCommand(ctx context.Context, cfgFile *utils.ConfigFile, cfg *environment_pb.ExtensionGetInfoCommand) error {
	if cfg == nil {
		return nil
	}
	switch inner := cfg.ExecConfig.(type) {
	case *environment_pb.ExtensionGetInfoCommand_KubernetesConfig:
		expanded, err := expandKubernetesConfig(ctx, filepath.Dir(cfgFile.ConfigFile), inner.KubernetesConfig)
		if err != nil {
			return err
		}
		inner.KubernetesConfig = expanded
	}
	return nil
}

func expandRuntimeConfig(ctx context.Context, cfgFile *utils.ConfigFile, cfg *environment_pb.ClusterConfig) error {
	for _, ext := range []*environment_pb.ExtensionClusterConfig{
		cfg.GetExtension(),
		cfg.GetCustom(),
	} {
		if ext != nil {
			if err := expandRuntimeExtensionApplyCommand(ctx, cfgFile, ext.Apply); err != nil {
				return err
			}
			if err := expandRuntimeExtensionFetchCommand(ctx, cfgFile, ext.Fetch); err != nil {
				return err
			}
			if err := expandRuntimeExtensionGetInfoCommand(ctx, cfgFile, ext.Debug); err != nil {
				return err
			}
			if err := expandRuntimeExtensionGetInfoCommand(ctx, cfgFile, ext.GetInfo); err != nil {
				return err
			}
		}
	}
	return nil
}

func processRuntimeConfig(ctx context.Context, cfgFile *utils.ConfigFile, cfg *environment_pb.ClusterConfig, validateOnly bool) error {
	err := expandRuntimeConfig(ctx, cfgFile, cfg)
	if err != nil {
		return err
	}

	validationResp, err := cmdutil.GetEnvironmentManagerClient().ValidateConfigureCluster(ctx, &environment_pb.ConfigureClusterReq{
		RuntimeName: cfg.Name,
		Config:      cfg,
	})
	if err != nil {
		return err
	}
	if validateOnly {
		return nil
	}

	err = handleConfigDiffPrompt(ctx, "runtime", validationResp.CompiledConfig, func(ctx context.Context) (string, proto.Message, error) {
		getLatestConfigResp, err := cmdutil.GetEnvironmentManagerClient().GetClusterConfig(ctx, &environment_pb.GetClusterConfigReq{
			RuntimeName: cfg.Name,
		})
		if err != nil {
			return "", nil, err
		}
		return "latest" /* runtimes are unversioned */, getLatestConfigResp.Config, nil
	})
	if err != nil {
		return err
	}

	_, err = cmdutil.GetEnvironmentManagerClient().ConfigureCluster(ctx,
		&environment_pb.ConfigureClusterReq{
			RuntimeName:    cfg.Name,
			Config:         cfg,
			Source:         version_pb.Source_CONFIG_FILE,
			SourceMetadata: utils.DetectSourceMetadata(cfgFile),
		},
	)
	if err != nil {
		return errors.Wrap(err, "failed to configure runtime")
	} else {
		fmt.Printf("Successfully configured %s\n", cfg.Name)
	}
	return nil
}

func processConfig(ctx context.Context, cfgFile *utils.ConfigFile, cfg *config_file_pb.ProdvanaConfig, validateOnly bool) error {
	switch inner := cfg.ConfigOneof.(type) {
	case *config_file_pb.ProdvanaConfig_Application:
		return processApplicationConfig(ctx, cfgFile, inner.Application, cfg.GetApplicationMetadata(), validateOnly)
	case *config_file_pb.ProdvanaConfig_Protection:
		return processProtectionConfig(ctx, cfgFile, inner.Protection, validateOnly)
	case *config_file_pb.ProdvanaConfig_DeliveryExtension:
		return processDeliveryExtensionConfig(ctx, cfgFile, inner.DeliveryExtension, validateOnly)
	case *config_file_pb.ProdvanaConfig_Service:
		return processServiceConfig(ctx, cfgFile, inner.Service, cfg.GetServiceMetadata(), validateOnly)
	case *config_file_pb.ProdvanaConfig_Runtime:
		return processRuntimeConfig(ctx, cfgFile, inner.Runtime, validateOnly)
	default:
		return errors.Errorf("unsupported config %T", cfg.ConfigOneof)
	}
}

func getApplyPriority(cfg *config_file_pb.ProdvanaConfig) int {
	switch cfg.ConfigOneof.(type) {
	case *config_file_pb.ProdvanaConfig_Protection:
		// need to be first as anything can reference a protection
		return 1
	case *config_file_pb.ProdvanaConfig_Runtime:
		return 2
	case *config_file_pb.ProdvanaConfig_DeliveryExtension:
		// just need to be before applications and services
		return 3
	case *config_file_pb.ProdvanaConfig_Application:
		return 4
	case *config_file_pb.ProdvanaConfig_Service:
		return 5
	default:
		return 0
	}
}

type cfgWithFile struct {
	cfgFile *utils.ConfigFile
	cfg     *config_file_pb.ProdvanaConfig
}

var applyCmd = &cobra.Command{
	Use:   "apply <name>",
	Short: "Apply a Prodvana config.",
	Long: `Apply a Prodvana config.

pvnctl configs apply my-service.pvn.yaml
pvnctl configs apply services/my-service/  # will look for all *.pvn.yaml inside this directory
`,
	Args:   cobra.MinimumNArgs(1),
	PreRun: cmdutil.ValidateAndRefreshAuthPreRun,
	Run: func(cmd *cobra.Command, args []string) {
		cfgFiles, err := utils.ProcessPvnConfigDirectories(args)
		if err != nil {
			log.Fatal(err)
		}
		ctx := context.Background()
		var overallErr error
		withFiles := make([]cfgWithFile, 0, len(cfgFiles))
		for _, f := range cfgFiles {
			cfg, err := utils.ParseConfig(f)
			if err != nil {
				overallErr = multierror.Append(overallErr, errors.Wrapf(err, "failed to process %s", f.ConfigFile))
				continue
			}
			withFiles = append(withFiles, cfgWithFile{
				cfgFile: f,
				cfg:     cfg,
			})
		}
		if overallErr != nil {
			log.Fatal(overallErr)
		}
		// stable sort to somewhat respect order the files were passed in
		sort.SliceStable(withFiles, func(i, j int) bool { return getApplyPriority(withFiles[i].cfg) < getApplyPriority(withFiles[j].cfg) })

		for _, cfgWithFile := range withFiles {
			log.Printf("Applying %s\n", cfgWithFile.cfgFile.ConfigFile)
			err = processConfig(ctx, cfgWithFile.cfgFile, cfgWithFile.cfg, false)
			if err != nil {
				overallErr = multierror.Append(overallErr, errors.Wrapf(err, "failed to process %s", cfgWithFile.cfgFile.ConfigFile))
				continue
			}
		}
		if overallErr != nil {
			log.Fatal(overallErr)
		}
	},
}

func init() {
	RootCmd.AddCommand(applyCmd)
	// TODO(naphat) delete once pvnctl has a push command that understands parameters
	applyCmd.Flags().BoolVar(&pushFlag, "push", false, "Push after applying with default parameter values. Exposed only for testing purposes, not meant to be used externally.")
	applyCmd.Flags().BoolVar(&noPromptFlag, "no-prompt", false, "Whether to prompt for confirmation before applying.")
	applyCmd.Flags().BoolVar(&forceFlag, "force", false, "Create a new config version even if there are no changes detected from the previous version.")
	cmdutil.Must(applyCmd.Flags().MarkHidden("push"))
}
