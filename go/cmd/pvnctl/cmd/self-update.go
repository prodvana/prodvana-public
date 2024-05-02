package cmd

import (
	_ "embed"
	"log"
	"os"

	"github.com/blang/semver"
	"github.com/pkg/errors"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
	"github.com/spf13/cobra"
)

//go:embed version.txt
var version string

const pvnctlRepo = "prodvana/pvnctl"

func makePvnctlUpdater() (*selfupdate.Updater, error) {
	return selfupdate.NewUpdater(selfupdate.Config{
		Filters: []string{"^pvnctl"},
	})
}

func parseCurrentVersion() (semver.Version, error) {
	v, err := semver.ParseTolerant(version)
	if err != nil {
		return semver.Version{}, errors.Wrap(err, "Invalid pvnctl version")
	}
	return v, nil
}

// very hacky implementation to figure out if self-update is the command being run.
func isSelfUpdateCmd() bool {
	for _, arg := range os.Args {
		if arg == "self-update" {
			return true
		}
	}
	return false
}

func printUpdateWarning() error {
	if os.Getenv("PVNCTL_SKIP_VERSION_CHECK") == "1" {
		return nil
	}
	updater, err := makePvnctlUpdater()
	if err != nil {
		return err
	}

	current, err := parseCurrentVersion()
	if err != nil {
		return errors.Errorf("Could not detect pvnctl version. Please run: pvnctl self-update")
	}

	latest, found, err := updater.DetectLatest(pvnctlRepo)
	if err != nil {
		return errors.Wrapf(err, "WARNING: Could not detect latest pvnctl version. Check your network connectivity?")
	}
	if !found {
		return errors.Errorf("WARNING: Could not find pvnctl repo. Please contact Prodvana support.")
	}
	if latest.Version.GT(current) {
		return errors.Errorf("Stale pvnctl version (current %s vs. latest %s). Please run: pvnctl self-update", current, latest.Version)
	}
	return nil
}

var selfUpdateCmd = &cobra.Command{
	Use:   "self-update",
	Short: "Update pvnctl in place",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		v, err := parseCurrentVersion()
		if err != nil {
			// Allow self-update to proceed even if currently installed binary has bad version.
			log.Printf("Version in current pvnctl version is invalid. Pulling latest pvnctl")
		}
		updater, err := makePvnctlUpdater()
		if err != nil {
			log.Fatal(err)
		}
		latest, err := updater.UpdateSelf(v, pvnctlRepo)
		if err != nil {
			log.Println("Binary update failed:", err)
			return
		}
		if latest.Version.Equals(v) {
			log.Println("Current binary is the latest version", version)
		} else {
			log.Println("Successfully updated to version", latest.Version)
		}
	},
}
