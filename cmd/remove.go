package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a Docker container",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println(getMessage("Por favor, forneça o nome ou ID do container.", "Please provide the container name or ID."))
			return
		}
		container := args[0]
		command := "docker rm " + container
		err := exec.Command("bash", "-c", command).Run()
		if err != nil {
			fmt.Println(getMessage("Erro ao remover o container:", "Error removing container:"), err)
			return
		}
		fmt.Println(getMessage("Container removido com sucesso.", "Container removed successfully."))
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
