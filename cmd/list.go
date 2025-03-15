package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List running Docker containers",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(getMessage("Listando containers em execução...", "Listing running containers..."))
		command := "docker ps"
		output, err := exec.Command("bash", "-c", command).CombinedOutput()
		if err != nil {
			fmt.Println(getMessage("Erro ao listar containers:", "Error listing containers:"), err)
			return
		}
		fmt.Println(string(output))
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
