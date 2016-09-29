package server

import (
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

	championId, err := ParseChampionId(query, "id")
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

	region, err := ParseRegion(query)
	if err != nil {
		Failure(w, r, err, http.StatusBadRequest)
		return
	}

	role, err := ParseRole(query)
	if err != nil {
		Failure(w, r, err, http.StatusBadRequest)
		return
	}

	champion, err := h.Apollo.GetChampion(h.Context, &apb.GetChampionRequest{
		ChampionId: championId,
		Patch:      patch,
		Tier:       tier,
		Region:     region,
		Role:       role,
	})
	if err != nil {
		Failure(w, r, err, http.StatusInternalServerError)
		return
	}

	Success(w, r, champion)
}

func (h *Handlers) HandleMatchup(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	focusId, err := ParseChampionId(query, "focus")
	if err != nil {
		Failure(w, r, err, http.StatusBadRequest)
		return
	}

	enemyId, err := ParseChampionId(query, "enemy")
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

	region, err := ParseRegion(query)
	if err != nil {
		Failure(w, r, err, http.StatusBadRequest)
		return
	}

	role, err := ParseRole(query)
	if err != nil {
		Failure(w, r, err, http.StatusBadRequest)
		return
	}

	matchup, err := h.Apollo.GetMatchup(h.Context, &apb.GetMatchupRequest{
		FocusChampionId: focusId,
		EnemyChampionId: enemyId,
		Patch:           patch,
		Tier:            tier,
		Region:          region,
		Role:            role,
	})
	if err != nil {
		Failure(w, r, err, http.StatusInternalServerError)
		return
	}

	Success(w, r, matchup)
}
