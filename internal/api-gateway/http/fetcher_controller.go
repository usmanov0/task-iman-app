package http

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"test-project-iman/internal/api-gateway/delivery/grpc_server_client"
	"test-project-iman/internal/api-gateway/delivery/grpc_server_client/fetch_datas/pb"
	"time"
)

type FetchController interface {
	FetchData(ctx *gin.Context)
}

type fetchController struct {
	s grpc_server_client.ServiceManager
}

func New(s grpc_server_client.ServiceManager) FetchController {
	return &fetchController{s: s}
}

func (f *fetchController) FetchData(ctx *gin.Context) {
	_, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	res, err := f.s.FetchData().CollectorPosts(ctx, &pb.Empty{})
	log.Println(">>>>", res)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": "Posts fetched successfully and saved to database successfully"})
}
