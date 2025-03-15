package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var language string

var rootCmd = &cobra.Command{
	Use:   "docker-assistance",
	Short: "Docker Assistance CLI for data science environments",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(getMessage("Bem-vindo ao Docker Assistance!", "Welcome to Docker Assistance!"))
		fmt.Println(getMessage("Use 'docker-assistance --help' para ver os comandos disponíveis.", "Use 'docker-assistance --help' to see available commands."))
	},
}

func getMessage(pt, en string) string {
	if language == "pt" {
		return pt
	}
	return en
}

func Execute() {
	rootCmd.PersistentFlags().StringVarP(&language, "lang", "l", "en", "Set language (en/pt)")
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
