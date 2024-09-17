package cmd

import (
	"fmt"
	"log/slog"

	"github.com/spf13/cobra"
	wh "sparkapp-k8s-webhook/pkg/webhook"
)

var (
	tlsCert string
	tlsKey  string
	port    uint
	cfgPath string
)

var (
	longDesc = `Example showing how to implement a basic webhook webhook in Kubernetes.`
	example  = `$ webhook-webhook --tls-cert <tls_cert> --tls-key <tls_key> --port <port> --cfgPath <cfgPath>`
)

var rootCmd = &cobra.Command{
	Use:     "spark-webhook",
	Short:   "Kubernetes webhook example",
	Long:    longDesc,
	Example: example,
	Run: func(cmd *cobra.Command, args []string) {
		logger := wh.InitLogger()
		var (
			cfg     wh.SparkAppConfig
			webHook wh.WebHook
		)

		if tlsCert == "" || tlsKey == "" {
			fmt.Println("--tls-cert and --tls-key required")
			slog.Warn("")
		}

		cfg, err := wh.GetConf(cfgPath)
		if err != nil {
			panic(err)
		}
		slog.Debug(fmt.Sprintf("Config: %+v", cfg))
		webHook = wh.WebHook{MutateConfig: &cfg}
		slog.Info("Success reading config")
		webHook.RunWebhookServer(tlsCert, tlsKey, port, logger)
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.Flags().StringVar(&tlsCert, "tls-cert", "", "Certificate for TLS")
	rootCmd.Flags().StringVar(&tlsKey, "tls-key", "", "Private key file for TLS")
	rootCmd.Flags().UintVar(&port, "port", 443, "Port to listen on for HTTPS traffic")
	rootCmd.Flags().StringVar(&cfgPath, "cfgPath", "", "Path to spark webhook config file")
}
