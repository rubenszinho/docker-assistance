package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var attachCmd = &cobra.Command{
	Use:   "attach",
	Short: "Attach to a running Docker container",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println(getMessage("Por favor, forneça o nome ou ID do container.", "Please provide the container name or ID."))
			return
		}
		container := args[0]
		command := "docker exec -it " + container + " bash"
		err := exec.Command("bash", "-c", command).Run()
		if err != nil {
			fmt.Println(getMessage("Erro ao acessar o container:", "Error attaching to container:"), err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(attachCmd)
}
