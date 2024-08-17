package webhook

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"log/slog"
	"os"
	"strings"
)

var (
	tlsCert string
	tlsKey  string
	port    uint
	cfgPath string
)

var longDesc = `Example showing how to implement a basic webhook webhook in Kubernetes.

Example:
$ webhook-webhook --tls-cert <tls_cert> --tls-key <tls_key> --port <port> --cfgPath <cfgPath>`

var rootCmd = &cobra.Command{
	Use:   "webhook-webhook",
	Short: "Kubernetes webhook webhook example",
	Long:  longDesc,
	Run: func(cmd *cobra.Command, args []string) {
		logger := initLogger()
		var (
			cfg SparkAppConfig
			wh  WebHook
		)

		if tlsCert == "" || tlsKey == "" {
			fmt.Println("--tls-cert and --tls-key required")
			slog.Warn("")
		}

		cfg, err := GetConf(cfgPath)
		if err != nil {
			panic(err)
		}
		slog.Debug(fmt.Sprintf("Config: %+v", cfg))
		wh = WebHook{MutateConfig: &cfg}
		slog.Info("Success reading config")
		wh.RunWebhookServer(tlsCert, tlsKey, port, logger)
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

func initLogger() *log.Logger {
	var logLevel slog.Level
	logLevelValue := strings.ToUpper(os.Getenv("LOG_LEVEL"))
	switch logLevelValue {
	case "":
		logLevel = slog.LevelInfo
	case "INFO":
		logLevel = slog.LevelInfo
	case "DEBUG":
		logLevel = slog.LevelDebug
	case "ERROR":
		logLevel = slog.LevelError
	}

	opts := &slog.HandlerOptions{
		Level: logLevel,
	}
	logger := slog.NewLogLogger(slog.NewJSONHandler(os.Stdout, nil), logLevel)
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, opts)))
	return logger
}
