package server

import (
	"errors"
	"strconv"
	"strings"

	"gopkg.in/gin-gonic/gin.v1"

	apb "github.com/asunaio/helios/gen-go/asuna"
)

func ParseChampionId(ctx *gin.Context, field string) (uint32, error) {
	id := ctx.Param(field)
	championId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		switch field {
		case "focus":
			return 0, errors.New(ErrorInvalidFocusId)
		case "enemy":
			return 0, errors.New(ErrorInvalidEnemyId)
		default:
			return 0, errors.New(ErrorInvalidChampionId)
		}
	}
	return uint32(championId), nil
}

func ParsePatch(ctx *gin.Context) (*apb.PatchRange, error) {
	patch := ctx.Query("patch")
	patches := strings.Split(patch, "-")
	if len(patches) == 1 {
		patches = append(patches, patches[0])
	}

	if patches[0] == "" || patches[1] == "" {
		return nil, errors.New(ErrorInvalidPatch)
	}

	return &apb.PatchRange{
		Min: patches[0],
		Max: patches[1],
	}, nil
}

func ParseTier(ctx *gin.Context) (*apb.TierRange, error) {
	tier := ctx.Query("tier")
	tiers := strings.Split(tier, "-")
	if len(tiers) == 1 {
		tiers = append(tiers, tiers[0])
	}
	minTier, err := strconv.ParseUint(tiers[0], 0, 32)
	if err != nil {
		return nil, errors.New(ErrorInvalidTier)
	}

	maxTier, err := strconv.ParseUint(tiers[1], 0, 32)
	if err != nil {
		return nil, errors.New(ErrorInvalidTier)
	}

	return &apb.TierRange{
		Min: uint32(minTier),
		Max: uint32(maxTier),
	}, nil
}

func ParseRegion(ctx *gin.Context) (apb.Region, error) {
	switch ctx.Query("region") {
	case "BR":
		return apb.Region_BR, nil
	case "EUNE":
		return apb.Region_EUNE, nil
	case "EUW":
		return apb.Region_EUW, nil
	case "JP":
		return apb.Region_JP, nil
	case "KR":
		return apb.Region_KR, nil
	case "LAN":
		return apb.Region_LAN, nil
	case "LAS":
		return apb.Region_LAS, nil
	case "NA":
		return apb.Region_NA, nil
	case "OCE":
		return apb.Region_OCE, nil
	case "TR":
		return apb.Region_TR, nil
	case "RU":
		return apb.Region_RU, nil
	case "PBE":
		return apb.Region_PBE, nil
	case "GLOBAL":
		return apb.Region_GLOBAL, nil
	default:
		return apb.Region_UNDEFINED_REGION, errors.New(ErrorInvalidRegion)
	}
}

func ParseRole(ctx *gin.Context) (apb.Role, error) {
	switch ctx.Query("role") {
	case "TOP":
		return apb.Role_TOP, nil
	case "JUNGLE":
		return apb.Role_JUNGLE, nil
	case "MID":
		return apb.Role_MID, nil
	case "BOT":
		return apb.Role_BOT, nil
	case "SUPPORT":
		return apb.Role_SUPPORT, nil
	default:
		return apb.Role_UNDEFINED_ROLE, errors.New(ErrorInvalidRole)
	}
}

func ParseMinPlayRate(ctx *gin.Context) (float64, error) {
	raw := ctx.Query("min_play_rate")
	mpr, err := strconv.ParseFloat(raw, 64)
	if err != nil {
		return 0, errors.New(ErrorInvalidMinPlayRate)
	}
	return mpr, nil

}
