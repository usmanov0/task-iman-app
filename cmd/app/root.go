package app

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	server2 "test-project-iman/cmd/app/server"
)

var root = &cobra.Command{Use: "run-grpc"}

var postCollectorCmd = &cobra.Command{
	Use: "grpc-server1",
	Run: func(cmd *cobra.Command, args []string) {
		server2.RunGrpcCollectorServer()
	},
}

var postGrpcServerCmd = &cobra.Command{
	Use: "grpc-server2",
	Run: func(cmd *cobra.Command, args []string) {
		server2.RunGrpcServer()
	},
}

var apiGateway = &cobra.Command{
	Use: "api-gateway",
	Run: func(cmd *cobra.Command, args []string) {
		server2.SetupApiGatewayRouter()
	},
}

func Execute() {
	root.AddCommand(postCollectorCmd)
	root.AddCommand(postGrpcServerCmd)
	root.AddCommand(apiGateway)
	if err := root.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error while executing CLI %s", err)
		os.Exit(1)
	}
}
