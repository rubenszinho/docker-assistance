package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var imageName, containerName string
var useGPU, useVolume, usePort bool
var volumePath, portMapping string

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create and start a new Docker container",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(getMessage("Criando container:", "Creating container:"), containerName, getMessage("com a imagem:", "with image:"), imageName)

		options := ""
		if useGPU {
			options += "--gpus all "
		}
		if useVolume {
			options += "-v " + volumePath + " "
		}
		if usePort {
			options += "-p " + portMapping + " "
		}

		command := fmt.Sprintf("docker run -d %s --name %s --user root %s tail -f /dev/null", options, containerName, imageName)
		fmt.Println(getMessage("Executando:", "Executing:"), command)

		cmdOut := exec.Command("bash", "-c", command)
		err := cmdOut.Run()
		if err != nil {
			fmt.Println(getMessage("Erro ao criar container:", "Error creating container:"), err)
			return
		}
		fmt.Println(getMessage("Container iniciado com sucesso.", "Container started successfully."))
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().StringVarP(&imageName, "image", "i", "", getMessage("Imagem Docker a ser utilizada (obrigatório)", "Docker image to use (required)"))
	createCmd.Flags().StringVarP(&containerName, "name", "n", "", getMessage("Nome do container (obrigatório)", "Container name (required)"))
	createCmd.Flags().BoolVar(&useGPU, "gpu", false, getMessage("Ativar suporte GPU", "Enable GPU support"))
	createCmd.Flags().BoolVar(&useVolume, "volume", false, getMessage("Adicionar volume", "Add volume"))
	createCmd.Flags().StringVar(&volumePath, "volume-path", "", getMessage("Caminho do volume (host_path:container_path)", "Volume path (host_path:container_path)"))
	createCmd.Flags().BoolVar(&usePort, "port", false, getMessage("Mapear porta", "Map a port"))
	createCmd.Flags().StringVar(&portMapping, "port-mapping", "", getMessage("Mapeamento de porta (host_port:container_port)", "Port mapping (host_port:container_port)"))
}
