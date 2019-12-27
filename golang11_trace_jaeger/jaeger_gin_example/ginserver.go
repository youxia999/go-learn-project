package main

import (
	"bytes"
	"github.com/gin-gonic/gin"
	_ "go-learn-project/golang10_mvc_gin/router/docs"
	"go-learn-project/golang11_trace_jaeger/jaeger_gin_example/util"
	"net/http"
)

func main() {
	engine:=gin.Default()
	engine.Use(util.SetUp("gin sever"))
	engine.GET("/someGet",getting)
	engine.Group("/api/v1")
	engine.POST("/postSome",posting)
	engine.Run(":8082")
}

func posting(context *gin.Context) {
	buf:=bytes.Buffer{}
	buf.ReadFrom(context.Request.Body)
	context.Request.Body.Close();
	s:=buf.String();
	context.String(http.StatusOK,"request is:"+s);
}

func getting(context *gin.Context) {
	msg:=context.Query("one");
	context.JSON(http.StatusOK,gin.H{"status":"OK","message":msg})
}
