package server

import (
	"github.com/kataras/iris"
	"golang.org/x/net/context"

	apb "github.com/asunaio/helios/gen-go/asuna"
)

type Handlers struct {
	Apollo  apb.ApolloClient
	Context context.Context
}

func (h *Handlers) HandleChampion(ctx *iris.Context) {
	championId, err := ParseChampionId(ctx, "id")
	if err != nil {
		Failure(ctx, err, iris.StatusBadRequest)
		return
	}

	patch, err := ParsePatch(ctx)
	if err != nil {
		Failure(ctx, err, iris.StatusBadRequest)
		return
	}

	tier, err := ParseTier(ctx)
	if err != nil {
		Failure(ctx, err, iris.StatusBadRequest)
		return
	}

	region, err := ParseRegion(ctx)
	if err != nil {
		Failure(ctx, err, iris.StatusBadRequest)
		return
	}

	role, err := ParseRole(ctx)
	if err != nil {
		Failure(ctx, err, iris.StatusBadRequest)
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
		Failure(ctx, err, iris.StatusInternalServerError)
		return
	}

	Success(ctx, champion)
}

func (h *Handlers) HandleMatchup(ctx *iris.Context) {
	focusId, err := ParseChampionId(ctx, "focus")
	if err != nil {
		Failure(ctx, err, iris.StatusBadRequest)
		return
	}

	enemyId, err := ParseChampionId(ctx, "enemy")
	if err != nil {
		Failure(ctx, err, iris.StatusBadRequest)
		return
	}

	patch, err := ParsePatch(ctx)
	if err != nil {
		Failure(ctx, err, iris.StatusBadRequest)
		return
	}

	tier, err := ParseTier(ctx)
	if err != nil {
		Failure(ctx, err, iris.StatusBadRequest)
		return
	}

	region, err := ParseRegion(ctx)
	if err != nil {
		Failure(ctx, err, iris.StatusBadRequest)
		return
	}

	role, err := ParseRole(ctx)
	if err != nil {
		Failure(ctx, err, iris.StatusBadRequest)
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
		Failure(ctx, err, iris.StatusInternalServerError)
		return
	}

	Success(ctx, matchup)
}