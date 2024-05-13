package lifespan

import (
	"fmt"
	"github.com/spf13/cobra"
	"gitlab.amazmetest.ru/ml/spark-amazme-webhook/pkg/mutating"
	"log/slog"
	"os"
)

var (
	tlsCert string
	tlsKey  string
	port    uint
	cfgPath string
)

var longDesc = `Example showing how to implement a basic mutating webhook in Kubernetes.

Example:
$ mutating-webhook --tls-cert <tls_cert> --tls-key <tls_key> --port <port>`

var rootCmd = &cobra.Command{
	Use:   "mutating-webhook",
	Short: "Kubernetes mutating webhook example",
	Long:  longDesc,
	Run: func(cmd *cobra.Command, args []string) {
		logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
		slog.SetDefault(logger)

		var (
			cfg mutating.SparkAppConfig
			wh  mutating.WebHook
		)

		if tlsCert == "" || tlsKey == "" {
			fmt.Println("--tls-cert and --tls-key required")
			slog.Warn("")
		}

		cfg.GetConf(cfgPath)
		wh = mutating.WebHook(cfg)
		wh.RunWebhookServer(tlsCert, tlsKey, port)
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.Flags().StringVar(&tlsCert, "tls-cert", "", "Certificate for TLS")
	rootCmd.Flags().StringVar(&tlsKey, "tls-key", "", "Private key file for TLS")
	rootCmd.Flags().IntVar(&port, "port", 443, "Port to listen on for HTTPS traffic")
}
