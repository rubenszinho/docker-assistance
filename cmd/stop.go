package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop a Docker container",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println(getMessage("Por favor, forneça o nome ou ID do container.", "Please provide the container name or ID."))
			return
		}
		container := args[0]
		command := "docker stop " + container
		err := exec.Command("bash", "-c", command).Run()
		if err != nil {
			fmt.Println(getMessage("Erro ao parar o container:", "Error stopping container:"), err)
			return
		}
		fmt.Println(getMessage("Container parado com sucesso.", "Container stopped successfully."))
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
