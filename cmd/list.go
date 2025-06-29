package cmd

import (
	"github.com/MegeKaplan/gobox/internal/messages"
	"github.com/MegeKaplan/gobox/internal/storage"
	"github.com/MegeKaplan/gobox/internal/utils"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List all installed packages",
	Long:    "Lists all Go packages that have been installed and saved using gobox, along with usage statistics.",
	Example: "  gobox list",
	Run:     runListPackages,
}

func runListPackages(cmd *cobra.Command, args []string) {
	packages, err := storage.LoadPackages()
	if err != nil {
		cmd.Println(color.RedString(messages.ErrLoadingPackagesFailed))
		return
	}

	if len(packages) == 0 {
		cmd.Println(color.YellowString(messages.StatusNoPackagesFound))
		return
	}

	utils.SortPackages(&packages, "last_used", false)

	cmd.Println(color.HiBlueString("Installed packages:"))
	for i, pkg := range packages {
		cmd.Printf("%s%d%s %s %s %s\n",
			color.HiCyanString("["),
			i+1,
			color.HiCyanString("]"),
			color.GreenString(pkg.Name),
			color.MagentaString("last used:"),
			color.HiMagentaString(pkg.LastUsed.Format("2006-01-02 15:04:05")),
		)
		cmd.Printf("     ↳ %s %s\n",
			color.BlueString("Used"),
			color.YellowString("%d times", pkg.UsageCount),
		)
	}

}

func init() {
	rootCmd.AddCommand(listCmd)
}
