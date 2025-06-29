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

var initCmd = &cobra.Command{
	Use:     "init [project_name]",
	Short:   "Initialize a new Go project",
	Long:    "Sets up a new Go project with a go.mod file and optionally installs selected packages from the gobox storage.",
	Example: "  gobox init myproject",
	Args:    cobra.MaximumNArgs(1),
	Run:     runInitProject,
}

func runInitProject(cmd *cobra.Command, args []string) {
	cmd.Println(color.BlueString(messages.StatusProjectInitializing))

	pw, _ := os.Getwd()

	projectPath := pw

	if len(args) > 0 {
		projectPath = filepath.Join(pw, args[0])
	}

	_, err := os.Stat(projectPath)
	if err == nil {
		entries, _ := os.ReadDir(projectPath)
		if len(entries) > 0 {
			cmd.Println(color.RedString(messages.ErrProjectAlreadyExists))
			return
		}
	}

	moduleName := filepath.Base(projectPath)

	moduleNamePrompt := &survey.Input{
		Message: messages.PromptModuleName,
		Default: moduleName,
	}
	survey.AskOne(moduleNamePrompt, &moduleName)

	packages, _ := storage.LoadPackages()

	utils.SortPackages(&packages, "last_used", false)

	os.MkdirAll(projectPath, 0755)
	os.Chdir(projectPath)

	if err := exec.Command("go", "mod", "init", moduleName).Run(); err != nil {
		cmd.Println(color.RedString(messages.ErrGoModInitFailed))
		return
	}

	cmd.Println(color.GreenString(messages.SuccessGoModCreated))

	if len(packages) == 0 {
		cmd.Println(color.YellowString(messages.StatusNoPackagesFound))
	} else {
		var options []string
		for _, pkg := range packages {
			options = append(options, pkg.Name)
		}

		var selected []string
		packageSelectPrompt := &survey.MultiSelect{
			Message: messages.PromptSelectPackagesToInstall,
			Options: options,
		}
		survey.AskOne(packageSelectPrompt, &selected)

		for _, pkg := range selected {
			s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
			s.Suffix = color.BlueString(fmt.Sprintf(messages.StatusInstallingPackage, pkg))
			s.FinalMSG = color.GreenString(fmt.Sprintf(messages.SuccessPackageInstalled, pkg))
			s.Color("blue")
			s.Start()

			_, err := exec.Command("go", "get", pkg).CombinedOutput()
			if err != nil {
				s.FinalMSG = color.RedString(messages.ErrPackageInstallFailed)
				s.Stop()
				return
			}

			if err := storage.SavePackage(pkg); err != nil {
				s.FinalMSG = color.RedString(messages.ErrPackageSaveFailed)
				s.Stop()
				return
			}

			s.Stop()
		}
	}

	cmd.Println(color.GreenString(fmt.Sprintf(messages.SuccessProjectInitialized, moduleName)))
}

func init() {
	rootCmd.AddCommand(initCmd)
}
