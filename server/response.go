package server

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

type failure struct {
	Code    int
	Message string
	Method  string
	Path    string
}

func Success(w http.ResponseWriter, r *http.Request, pb proto.Message) {
	var jsonb bytes.Buffer
	if err := (&jsonpb.Marshaler{
		EnumsAsInts:  false,
		EmitDefaults: true,
		OrigName:     false,
	}).Marshal(&jsonb, pb); err != nil {
		Failure(w, r, err, http.StatusInternalServerError)
	}
	w.Write(jsonb.Bytes())
}

func Failure(w http.ResponseWriter, r *http.Request, e error, code int) {
	log.Printf("[ERROR] %s %s: %v", r.Method, r.URL.Path, e)
	w.WriteHeader(code)
	jsonb, err := json.Marshal(failure{
		Code:    code,
		Message: e.Error(),
		Method:  r.Method,
		Path:    r.URL.Path,
	})

	if err != nil {
		w.Write([]byte("{ 'Message': 'JSON Marshal Failure.'}"))
	} else {
		w.Write(jsonb)
	}
}
