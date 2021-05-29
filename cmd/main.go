package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// Commander is a struct for command system.
type Commander struct {
	root     *cobra.Command
	logLevel string
}

func (c *Commander) setRootCommand() {
	c.root = &cobra.Command{
		Use:          "vault-unseal",
		Short:        "Vault Auto Unseal",
		Long:         "CLI command to automatically unseal Vault",
		SilenceUsage: true,
	}
}

func (c *Commander) setPersistentFlags() {
	c.root.PersistentFlags().String("logLevel", "info", "Set the logging level. One of: debug|info|warn|error")
}

func (c *Commander) setLogger() {
	c.logLevel, _ = c.root.Flags().GetString("logLevel")

	level, err := log.ParseLevel(c.logLevel)
	if err != nil {
		log.Fatalf("Lethal damage: %s\n\n", err)
	}

	log.SetLevel(level)
	log.SetFormatter(&log.TextFormatter{
		TimestampFormat: "02-01-2006 15:04:05",
		FullTimestamp:   true,
	})
}

// Execute is entrypoint for the commands.
func (c *Commander) Execute() error {
	cobra.OnInitialize(func() {
		c.setLogger()
	})

	c.setRootCommand()
	c.setPersistentFlags()

	c.root.AddCommand(Version())
	c.root.AddCommand(Unseal())

	return c.root.Execute()
}
