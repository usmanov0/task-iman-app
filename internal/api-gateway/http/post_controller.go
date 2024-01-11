package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"test-project-iman/internal/api-gateway/delivery/grpc_server_client"
	"test-project-iman/internal/api-gateway/delivery/grpc_server_client/post_service/pb"
)

type PostController interface {
	GetList(ctx *gin.Context)
	GetOne(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type postController struct {
	s grpc_server_client.ServiceManager
}

func NewPostController(s grpc_server_client.ServiceManager) PostController {
	return &postController{s: s}
}

func (p *postController) GetList(ctx *gin.Context) {
	pageStr := ctx.Query("page")
	limitStr := ctx.Query("limit")

	page, err := strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'page' parameter"})
		return
	}

	limit, err := strconv.ParseInt(limitStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'limit' parameter"})
	}

	response, err := p.s.PostService().GetList(ctx, &pb.PostRequestPage{
		Page:  page,
		Limit: limit,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": response.Post})
}

func (p *postController) GetOne(ctx *gin.Context) {
	id := ctx.Param("id")

	postId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'id' parameter"})
	}

	response, err := p.s.PostService().GetPost(ctx, &pb.PostRequestId{
		Id: postId,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": response})
}

func (p *postController) Update(ctx *gin.Context) {
	id := ctx.Param("id")

	postId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'id' parameter"})
	}

	res, err := p.s.PostService().Update(ctx, &pb.PostUpdate{
		Id:    postId,
		Title: ctx.PostForm("title"),
		Body:  ctx.PostForm("body"),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
	}
	ctx.JSON(http.StatusOK, gin.H{"success": res.Success, "message": res.Message})
}

func (p *postController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	postId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'id' parameter"})
	}

	_, err = p.s.PostService().Delete(ctx, &pb.PostRequestId{
		Id: postId,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	ctx.JSON(http.StatusOK, "Deleted successfully")
}
