package server

import (
	"bytes"
	"log"
	"net/http"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"gopkg.in/gin-gonic/gin.v1"
)

func Success(ctx *gin.Context, pb proto.Message) {
	var jsonb bytes.Buffer
	if err := (&jsonpb.Marshaler{
		EnumsAsInts:  false,
		EmitDefaults: true,
		OrigName:     false,
	}).Marshal(&jsonb, pb); err != nil {
		Failure(ctx, err, http.StatusInternalServerError)
		return
	}
	ctx.Header("Content-Type", "application/json")
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.String(http.StatusOK, jsonb.String())
}

func Failure(ctx *gin.Context, e error, status int) {
	log.Printf("[ERROR] %s %s: %v", ctx.Request.Method, ctx.Request.URL.Path, e)
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.JSON(status, gin.H{
		"status":  status,
		"message": e.Error(),
		"method":  ctx.Request.Method,
		"path":    ctx.Request.URL.Path,
		"query":   ctx.Request.URL.Query(),
	})
}
