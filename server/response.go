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
	}
	ctx.Header("Content-Type", "application/json")
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.String(http.StatusOK, jsonb.String())
}

func Failure(ctx *gin.Context, e error, status int) {
	log.Printf("[ERROR] %s %s: %v", ctx.MethodString(), ctx.PathString(), e)
	ctx.JSON(status, gin.H{
		"status":  status,
		"message": e.Error(),
		"method":  ctx.MethodString(),
		"path":    ctx.PathString(),
		"query":   ctx.URLParams(),
		"uri":     ctx.RequestPath(false),
	})
}
