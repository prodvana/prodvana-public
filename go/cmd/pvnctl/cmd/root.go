package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/prodvana/prodvana-public/go/pkg/cmdutil"

	"github.com/prodvana/prodvana-public/go/cmd/pvnctl/cmd/apitokens"
	"github.com/prodvana/prodvana-public/go/cmd/pvnctl/cmd/applications"
	"github.com/prodvana/prodvana-public/go/cmd/pvnctl/cmd/auth"
	"github.com/prodvana/prodvana-public/go/cmd/pvnctl/cmd/blobs"
	"github.com/prodvana/prodvana-public/go/cmd/pvnctl/cmd/configs"
	"github.com/prodvana/prodvana-public/go/cmd/pvnctl/cmd/deployments"
	"github.com/prodvana/prodvana-public/go/cmd/pvnctl/cmd/desiredstates"
	"github.com/prodvana/prodvana-public/go/cmd/pvnctl/cmd/healthcheck"
	"github.com/prodvana/prodvana-public/go/cmd/pvnctl/cmd/integration"
	"github.com/prodvana/prodvana-public/go/cmd/pvnctl/cmd/labels"
	"github.com/prodvana/prodvana-public/go/cmd/pvnctl/cmd/organizations"
	"github.com/prodvana/prodvana-public/go/cmd/pvnctl/cmd/preflight"
	"github.com/prodvana/prodvana-public/go/cmd/pvnctl/cmd/protections"
	"github.com/prodvana/prodvana-public/go/cmd/pvnctl/cmd/recipes"
	"github.com/prodvana/prodvana-public/go/cmd/pvnctl/cmd/releasechannels"
	"github.com/prodvana/prodvana-public/go/cmd/pvnctl/cmd/runtimes"
	"github.com/prodvana/prodvana-public/go/cmd/pvnctl/cmd/secrets"
	"github.com/prodvana/prodvana-public/go/cmd/pvnctl/cmd/services"

	"github.com/prodvana/prodvana-public/go/pkg/client"

	"github.com/prodvana/prodvana-public/go/prodvana-sdk/session"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	format        string
	colorDisabled bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:              "pvnctl",
	Short:            "Prodvana CLI",
	Long:             `pvnctl is used to operate Prodvana from the command line.`,
	TraverseChildren: true,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func setOutputFormat() {
	err := cmdutil.SetOutputFormat(format)
	if err != nil {
		log.Fatal(err)
	}
	if colorDisabled || !cmdutil.ShouldPrintColor() {
		color.NoColor = true
	}
}

func init() {
	log.SetFlags(0) // no log prefix
	var traceAPIRequests bool
	s := session.GetSession()

	cobra.OnInitialize(setOutputFormat, func() {
		if traceAPIRequests {
			client.EnableTraceAPIRequests()
		}

		globalApiAddr := "api.prodvana.io"
		if os.Getenv("PVN_GLOBAL_API_OVERRIDE") != "" {
			globalApiAddr = os.Getenv("PVN_GLOBAL_API_OVERRIDE")
		}
		err := client.SetDefaultGlobalApiAddr(globalApiAddr)
		if err != nil {
			log.Fatal(err)
		}

		var addr string
		var skipTls bool
		if cmdutil.ApiServerAddrOverride != "" {
			addr = cmdutil.ApiServerAddrOverride
		} else {
			authCtx, ok := s.Contexts[session.InProcessContext]
			if ok {
				addr = authCtx.Addr
				skipTls = authCtx.SkipTls
			}
		}
		if addr != "" {
			err = client.SetDefaultApiServerAddr(addr, skipTls)
			if err != nil {
				log.Fatal(err)
			}
		}
		err = client.SetWaitForReady(false) // since pvnctl is (mostly) interactive, it is better to fail fast when failing to connect to Prodvana
		if err != nil {
			log.Fatal(err)
		}
		err = client.SetSource("pvnctl")
		if err != nil {
			log.Fatal(err)
		}

		if !isSelfUpdateCmd() {
			if err := printUpdateWarning(); err != nil {
				// soft errors only, stale versions are not necessarily catastrophic
				log.Println(err)
			}
		}
	})

	rootCmd.Flags().StringVar(&cmdutil.ApiServerAddrOverride, "apiserver-addr", "", "Address for apiserver, exposed mainly for internal uses. This does NOT work with --context flag. Use of both together leads to undefined behavior.")
	cmdutil.Must(rootCmd.Flags().MarkHidden("apiserver-addr")) // only used by pipelines
	rootCmd.Flags().BoolVar(&traceAPIRequests, "trace-api-requests", false, "(internal) Trace API requests")
	cmdutil.Must(rootCmd.Flags().MarkHidden("trace-api-requests"))
	rootCmd.Flags().StringVar(&session.InProcessContext, "context", s.CurrentContext, "Override default context for pvnctl")
	rootCmd.PersistentFlags().StringVar(&cmdutil.Editor, "editor", "", "Editor to use for interactive commands. If not set, default to value of EDITOR, then GIT_EDITOR.")
	rootCmd.PersistentFlags().StringVarP(&format, "format", "f", cmdutil.OutputFormat, fmt.Sprintf("Output format. Accepts one of [%s].", strings.Join(cmdutil.ValidOutputFormats(), ", ")))
	rootCmd.PersistentFlags().BoolVar(&colorDisabled, "no-color", false, "Disables color output; this is implicitly set if output == json")
	rootCmd.PersistentFlags().BoolVar(&cmdutil.ForceAsyncSetDesiredState, "async-set-desired-state", os.Getenv("FORCE_ASYNC_SET_DESIRED_STATE") == "1", "Use async mode for SetDesiredState")
	cmdutil.Must(rootCmd.PersistentFlags().MarkHidden("async-set-desired-state"))
	rootCmd.AddCommand(releasechannels.RootCmd)
	rootCmd.AddCommand(organizations.RootCmd)
	rootCmd.AddCommand(services.RootCmd)
	rootCmd.AddCommand(runtimes.RootCmd)
	rootCmd.AddCommand(auth.RootCmd)
	rootCmd.AddCommand(integration.RootCmd)
	rootCmd.AddCommand(healthcheck.RootCmd)
	rootCmd.AddCommand(apitokens.RootCmd)
	rootCmd.AddCommand(applications.RootCmd)
	rootCmd.AddCommand(desiredstates.RootCmd)
	rootCmd.AddCommand(protections.RootCmd)
	rootCmd.AddCommand(configs.RootCmd)
	rootCmd.AddCommand(secrets.RootCmd)
	rootCmd.AddCommand(blobs.RootCmd)
	rootCmd.AddCommand(labels.RootCmd)
	rootCmd.AddCommand(deployments.RootCmd)
	rootCmd.AddCommand(recipes.RootCmd)
	rootCmd.AddCommand(selfUpdateCmd)
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(preflight.RootCmd)
}
