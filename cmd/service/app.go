package main

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	configurationFilePathFlag string

	rootCmd = &cobra.Command{
		Use:   "service",
		Short: "Example service",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			// configuration.InitConfig(configurationFilePathFlag)

			// appConfiguration := configuration.GetConfiguration()

			// service.Run(appConfiguration)
		},
	}
)

func init() {
	logger, _ := zap.NewProduction()
	zap.ReplaceGlobals(logger)
	setupFlags()
}

func setupFlags() {
	// Setup flags
	rootCmd.PersistentFlags().StringVar(&configurationFilePathFlag, "config", "", "config file path")

	// Bind flags to viper
	_ = viper.BindPFlag("tracing.enable", rootCmd.PersistentFlags().Lookup("tracing"))
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		zap.L().Fatal("Unable to run", zap.Error(err))
	}
}
