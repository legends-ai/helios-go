package server

import (
	"bytes"
	"log"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/kataras/iris"
)

func Success(ctx *iris.Context, pb proto.Message) {
	var jsonb bytes.Buffer
	if err := (&jsonpb.Marshaler{
		EnumsAsInts:  false,
		EmitDefaults: true,
		OrigName:     false,
	}).Marshal(&jsonb, pb); err != nil {
		Failure(ctx, err, iris.StatusInternalServerError)
	}
	ctx.SetHeader("Content-Type", "application/json")
	ctx.Write(jsonb.String())
}

func Failure(ctx *iris.Context, e error, status int) {
	log.Printf("[ERROR] %s %s: %v", ctx.MethodString(), ctx.PathString(), e)
	ctx.JSON(status, iris.Map{
		"status":  status,
		"message": e.Error(),
		"method":  ctx.MethodString(),
		"path":    ctx.PathString(),
		"query":   ctx.URLParams(),
		"uri":     ctx.RequestPath(false),
	})
}
