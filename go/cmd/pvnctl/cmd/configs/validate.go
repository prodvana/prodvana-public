package configs

import (
	"context"
	"log"
	"sort"

	"github.com/prodvana/prodvana-public/go/cmd/pvnctl/cmd/configs/internal/utils"

	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var validateFlags = struct {
	remoteValidation bool
}{}

var validateCmd = &cobra.Command{
	Use:   "validate <name>",
	Short: "Validate a Prodvana config.",
	Long: `Validate a Prodvana config.

pvnctl configs validate my-service.pvn.yaml
pvnctl configs vwalidate services/my-service/  # will look for all *.pvn.yaml inside this directory
`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cfgFiles, err := utils.ProcessPvnConfigDirectories(args)
		if err != nil {
			log.Fatal(err)
		}
		ctx := context.Background()
		var overallErr error
		withFiles := make([]cfgWithFile, 0, len(cfgFiles))
		for _, f := range cfgFiles {
			log.Printf("Parsing %s\n", f.ConfigFile)
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

		// stable sort to somewhat respect order the files were passed in
		sort.SliceStable(withFiles, func(i, j int) bool { return getApplyPriority(withFiles[i].cfg) < getApplyPriority(withFiles[j].cfg) })

		if validateFlags.remoteValidation {
			for _, cfgWithFile := range withFiles {
				log.Printf("Validating %s\n", cfgWithFile.cfgFile.ConfigFile)
				err = processConfig(ctx, cfgWithFile.cfgFile, cfgWithFile.cfg, true)
				if err != nil {
					overallErr = multierror.Append(overallErr, errors.Wrapf(err, "failed to process %s", cfgWithFile.cfgFile.ConfigFile))
					continue
				}
			}
		}

		if overallErr != nil {
			log.Fatal(overallErr)
		}
	},
}

func init() {
	RootCmd.AddCommand(validateCmd)
	validateCmd.Flags().BoolVar(&validateFlags.remoteValidation, "remote", false, "Do remote validations against the underlying runtime.")
}
