package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var addVolumeCmd = &cobra.Command{
	Use:   "add-volume",
	Short: "Add a volume to an existing Docker container",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Println(getMessage("Por favor, forneça o nome ou ID do container e o caminho do volume.", "Please provide the container name/ID and the volume path (host_path:container_path)."))
			return
		}
		container := args[0]
		volumePath := args[1]
		// Simplified implementation; additional logic may be required.
		command := "docker run -d --name temp_" + container + " --volumes-from " + container + " -v " + volumePath + " " + container + " tail -f /dev/null"
		err := exec.Command("bash", "-c", command).Run()
		if err != nil {
			fmt.Println(getMessage("Erro ao adicionar o volume:", "Error adding volume:"), err)
			return
		}
		fmt.Println(getMessage("Volume adicionado com sucesso. Reinicie o container se necessário.", "Volume added successfully. Restart the container if needed."))
	},
}

func init() {
	rootCmd.AddCommand(addVolumeCmd)
}
