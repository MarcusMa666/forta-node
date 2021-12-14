package cmd

import (
	"github.com/forta-protocol/forta-node/config"

	"github.com/spf13/cobra"
)

func handleFortaImages(cmd *cobra.Command, args []string) error {
	cmd.Println("Use images:", config.UseDockerImages)
	cmd.Println("Node:", config.DockerScannerNodeImage)
	return nil
}
