package apigateway

import (
	"github.com/gin-gonic/gin"
	topic_client "github.com/yxm0513/go-micro-service/client/topic"
	"github.com/yxm0513/go-micro-service/proto/topic"
	"golang.org/x/net/context"
	"net/http"
	"strconv"
)

func RegisterTopic(router *gin.RouterGroup) {
	r := router.Group("/topic")
	r.GET("/view", view)
}

func view(c *gin.Context) {
	topicID, err := strconv.ParseInt(c.Query("topic_id"), 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	req := &topic.GetTopicRequest{topicID}
	resp, err := topic_client.GetClient().GetTopic(context.Background(), req)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, resp)
}
