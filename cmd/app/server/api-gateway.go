package server

import (
	"test-project-iman/internal/api-gateway/delivery/grpc_server_client"
	router2 "test-project-iman/internal/api-gateway/router"
)

func ApiGatewayServer() {
	s := grpc_server_client.NewApiClient()

	router := router2.SetupRouter(s)

	router.Run(":8082")
}
