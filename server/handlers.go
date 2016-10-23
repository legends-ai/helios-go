package server

import (
	"net/http"

	"golang.org/x/net/context"

	apb "github.com/asunaio/helios/gen-go/asuna"
	gin "gopkg.in/gin-gonic/gin.v1"
)

type Handlers struct {
	Apollo  apb.ApolloClient
	Context context.Context
}

func (h *Handlers) HandleChampion(ctx *gin.Context) {
	championId, err := ParseChampionId(ctx, "id")
	if err != nil {
		Failure(ctx, err, http.StatusBadRequest)
		return
	}

	patch, err := ParsePatch(ctx)
	if err != nil {
		Failure(ctx, err, http.StatusBadRequest)
		return
	}

	tier, err := ParseTier(ctx)
	if err != nil {
		Failure(ctx, err, http.StatusBadRequest)
		return
	}

	region, err := ParseRegion(ctx)
	if err != nil {
		Failure(ctx, err, http.StatusBadRequest)
		return
	}

	role, err := ParseRole(ctx)
	if err != nil {
		Failure(ctx, err, http.StatusBadRequest)
		return
	}

	mpr, err := ParseMinPlayRate(ctx)
	if err != nil {
		Failure(ctx, err, http.StatusBadRequest)
		return
	}

	champion, err := h.Apollo.GetChampion(h.Context, &apb.GetChampionRequest{
		ChampionId:  championId,
		Patch:       patch,
		Tier:        tier,
		Region:      region,
		Role:        role,
		MinPlayRate: mpr,
	})
	if err != nil {
		Failure(ctx, err, http.StatusInternalServerError)
		return
	}

	Success(ctx, champion)
}

func (h *Handlers) HandleMatchup(ctx *http.Context) {
	focusId, err := ParseChampionId(ctx, "focus")
	if err != nil {
		Failure(ctx, err, http.StatusBadRequest)
		return
	}

	enemyId, err := ParseChampionId(ctx, "enemy")
	if err != nil {
		Failure(ctx, err, http.StatusBadRequest)
		return
	}

	patch, err := ParsePatch(ctx)
	if err != nil {
		Failure(ctx, err, http.StatusBadRequest)
		return
	}

	tier, err := ParseTier(ctx)
	if err != nil {
		Failure(ctx, err, http.StatusBadRequest)
		return
	}

	region, err := ParseRegion(ctx)
	if err != nil {
		Failure(ctx, err, http.StatusBadRequest)
		return
	}

	role, err := ParseRole(ctx)
	if err != nil {
		Failure(ctx, err, http.StatusBadRequest)
		return
	}

	mpr, err := ParseMinPlayRate(ctx)
	if err != nil {
		Failure(ctx, err, http.StatusBadRequest)
		return
	}

	matchup, err := h.Apollo.GetMatchup(h.Context, &apb.GetMatchupRequest{
		FocusChampionId: focusId,
		EnemyChampionId: enemyId,
		Patch:           patch,
		Tier:            tier,
		Region:          region,
		Role:            role,
		MinPlayRate:     mpr,
	})
	if err != nil {
		Failure(ctx, err, http.StatusInternalServerError)
		return
	}

	Success(ctx, matchup)
}
