package utils

import (
	"log"
	"path/filepath"
	"prodvana/cmd/cmdutil"

	version_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/version"
)

func DetectSourceMetadata(cfgFile *ConfigFile) *version_pb.SourceMetadata {
	ciDefaults := cmdutil.DetectCIDefaults()
	meta := &version_pb.SourceMetadata{}
	var baseDir string
	if ciDefaults != nil {
		meta.Commit = ciDefaults.Commit
		meta.RepoUrl = ciDefaults.Repository
		baseDir = ciDefaults.BaseDirectory
	}
	if baseDir != "" {
		relPath, err := filepath.Rel(baseDir, cfgFile.ConfigFile)
		if err == nil {
			meta.FilePath = relPath
		} else {
			log.Printf("Warning: config directory %s not in detected repository directory %s", cfgFile.ConfigFile, baseDir)
		}
	}
	return meta
}
