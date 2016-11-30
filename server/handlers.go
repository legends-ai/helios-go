package server

import (
	"net/http"

	"golang.org/x/net/context"

	apb "github.com/asunaio/helios/gen-go/asuna"
	gin "gopkg.in/gin-gonic/gin.v1"
)

type Handlers struct {
	Lucinda apb.LucindaClient
	Vulgate apb.VulgateClient
	Context context.Context
}

func (h *Handlers) HandleChampion(ctx *gin.Context) {
	championId, err := parseChampionId(ctx, "id")
	if err != nil {
		Failure(ctx, err, http.StatusBadRequest)
		return
	}

	tier, err := parseTier(ctx)
	if err != nil {
		Failure(ctx, err, http.StatusBadRequest)
		return
	}

	region, err := parseRegion(ctx)
	if err != nil {
		Failure(ctx, err, http.StatusBadRequest)
		return
	}

	champion, err := h.Lucinda.GetChampion(h.Context, &apb.LucindaRpc_GetChampionRequest{
		ChampionId:  championId,
		Patch:       parsePatch(ctx),
		Tier:        tier,
		Region:      region,
		Role:        parseRole(ctx),
		MinPlayRate: parseMinPlayRate(ctx),
	})

	if err != nil {
		Failure(ctx, err, http.StatusInternalServerError)
		return
	}

	Success(ctx, champion)
}

func (h *Handlers) HandleMatchup(ctx *gin.Context) {
	focusId, err := parseChampionId(ctx, "focus")
	if err != nil {
		Failure(ctx, err, http.StatusBadRequest)
		return
	}

	enemyId, err := parseChampionId(ctx, "enemy")
	if err != nil {
		Failure(ctx, err, http.StatusBadRequest)
		return
	}

	tier, err := parseTier(ctx)
	if err != nil {
		Failure(ctx, err, http.StatusBadRequest)
		return
	}

	region, err := parseRegion(ctx)
	if err != nil {
		Failure(ctx, err, http.StatusBadRequest)
		return
	}

	matchup, err := h.Lucinda.GetMatchup(h.Context, &apb.LucindaRpc_GetMatchupRequest{
		FocusChampionId: focusId,
		EnemyChampionId: enemyId,
		Patch:           parsePatch(ctx),
		Tier:            tier,
		Region:          region,
		Role:            parseRole(ctx),
		MinPlayRate:     parseMinPlayRate(ctx),
	})

	if err != nil {
		Failure(ctx, err, http.StatusInternalServerError)
		return
	}

	Success(ctx, matchup)
}

func (h *Handlers) HandleStaticEntry(ctx *gin.Context) {
	region, err := parseRegion(ctx)
	if err != nil {
		Failure(ctx, err, http.StatusBadRequest)
		return
	}

	context := &apb.VulgateRpc_Context{
		Locale: parseLocale(ctx),
		Region: region,
	}

	if ctx.Query("version") != "" {
		context.Release = &apb.VulgateRpc_Context_Version{
			Version: ctx.Query("version"),
		}
	} else if ctx.Query("patch") != "" {
		context.Release = &apb.VulgateRpc_Context_Patch{
			Patch: ctx.Query("patch"),
		}
	}

	entry, err := h.Vulgate.GetEntry(ctx, &apb.VulgateRpc_GetEntryRequest{
		Context: context,
		Format:  apb.VulgateRpc_GetEntryRequest_Format(apb.VulgateRpc_GetEntryRequest_Format_value[ctx.Query("format")]),
	})

	if err != nil {
		Failure(ctx, err, http.StatusInternalServerError)
		return
	}

	Success(ctx, entry)
}
