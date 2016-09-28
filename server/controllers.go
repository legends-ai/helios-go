package server

import (
	"net/http"

	apb "github.com/asunaio/helios/gen-go/asuna"
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
)

type Controllers struct {
	Apollo  apb.ApolloClient
	Context context.Context
}

func (c *Controllers) HandleChampion(w http.ResponseWriter, r *http.Request) {
}

func (c *Controllers) HandleMatchup(w http.ResponseWriter, r *http.Request) {

}
