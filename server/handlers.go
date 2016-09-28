package server

import (
	"fmt"
	"net/http"

	apb "github.com/asunaio/helios/gen-go/asuna"
	"golang.org/x/net/context"
)

type Handlers struct {
	Apollo  apb.ApolloClient
	Context context.Context
}

func (h *Handlers) HandleChampion(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	championId, err := ParseChampionId(query)
	if err != nil {
		Failure(w, r, err, http.StatusBadRequest)
		return
	}

	patch, err := ParsePatch(query)
	if err != nil {
		Failure(w, r, err, http.StatusBadRequest)
		return
	}

	tier, err := ParseTier(query)
	if err != nil {
		Failure(w, r, err, http.StatusBadRequest)
		return
	}

	fmt.Println(championId)
	fmt.Println(patch)
	fmt.Println(tier)

	champion, err := h.Apollo.GetChampion(h.Context, &apb.GetChampionRequest{
		ChampionId: championId,
		Patch:      patch,
		Tier:       tier,
		Region:     apb.Region_NA,
		Role:       apb.Role_JUNGLE,
	})

	if err != nil {
		Failure(w, r, err, http.StatusInternalServerError)
		return
	}

	Success(w, r, champion)
}

func (h *Handlers) HandleMatchup(w http.ResponseWriter, r *http.Request) {

}
