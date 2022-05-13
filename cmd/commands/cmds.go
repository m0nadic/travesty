package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"travesty/internal/app/model"
	"travesty/internal/app/server"
	"travesty/internal/app/service"
	"travesty/internal/app/util"
)

const defaultFile = "travesty.yaml"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Displays version information",
	Long:  "Display version information and exit.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("1.0.0")
	},
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Creates sample config files",
	Long:  "Creates sample config files for services in the current folder.",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			fmt.Println("Generating", defaultFile)
			err := util.GenerateFile(defaultFile)
			if err != nil {
				_, _ = fmt.Fprintln(os.Stderr, "ERROR:", err)
				return
			}
			return
		}

		for _, arg := range args {
			fmt.Println("Generating", arg)
			err := util.GenerateFile(arg)
			if err != nil {
				_, _ = fmt.Fprintln(os.Stderr, "ERROR:", err)
				continue
			}
		}

	},
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the services",
	Long:  "Starts the mock services defined in the input files supplied as arguments",
	Run: func(cmd *cobra.Command, args []string) {
		services := make([]model.Service, 0)
		for _, arg := range args {
			fmt.Println("Starting mock service from file", arg)
			svc, err := util.LoadFile(arg)
			if err != nil {
				_, _ = fmt.Fprintln(os.Stderr, "ERROR:", err)
				continue
			}
			spawner := service.NewSpawner()
			spawner.Start(svc)
			services = append(services, svc)
		}

		server.StartAdminServer("0.0.0.0", 3040, services)
	},
}
