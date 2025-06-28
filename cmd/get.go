package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/MegeKaplan/gobox/internal/messages"
	"github.com/MegeKaplan/gobox/internal/storage"
	"github.com/MegeKaplan/gobox/internal/utils"
	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:     "get [package]",
	Short:   "Download and save a Go package",
	Long:    "Fetches a Go package using 'go get' and saves it in the gobox storage for future use.",
	Example: "  gobox get github.com/MegeKaplan/gobox",
	Args:    cobra.ExactArgs(1),
	Run:     runGetPackage,
}

func runGetPackage(cmd *cobra.Command, args []string) {
	packageName := args[0]

	wd, _ := os.Getwd()

	base := filepath.Base(wd)

	exists := utils.FileExists("go.mod")

	noGoModFoundConfirm := false
	noGoModFoundPrompt := &survey.Confirm{
		Message: messages.ErrNoGoModFound,
		Default: true,
	}

	if !exists {
		if err := survey.AskOne(noGoModFoundPrompt, &noGoModFoundConfirm); err != nil {
			cmd.Println(color.RedString(messages.ErrPromptFailed))
			return
		}

		if noGoModFoundConfirm {
			err := exec.Command("go", "mod", "init", base).Run()
			if err != nil {
				cmd.Println(messages.ErrGoModInitFailed)
				return
			}
			cmd.Println(color.GreenString(messages.SuccessGoModCreated))
		} else {
			cmd.Println(color.YellowString(messages.ErrOperationAborted))
			return
		}
	}

	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Suffix = color.BlueString(fmt.Sprintf(messages.StatusInstallingPackage, packageName))
	s.FinalMSG = color.GreenString(fmt.Sprintf(messages.SuccessPackageInstalled, packageName))
	s.Color("blue")
	s.Start()

	_, err := exec.Command("go", "get", packageName).CombinedOutput()
	if err != nil {
		s.FinalMSG = color.RedString(fmt.Sprintf(messages.ErrPackageInstallFailed, packageName))
		s.Stop()
		return
	}

	if err := storage.SavePackage(packageName); err != nil {
		cmd.Println(color.RedString(messages.ErrPackageSaveFailed, packageName))
		return
	}

	s.Stop()
}

func init() {
	rootCmd.AddCommand(getCmd)
}
