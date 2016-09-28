package server

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	apb "github.com/asunaio/helios/gen-go/asuna"
)

func ParseChampionId(query map[string][]string) (uint32, error) {
	if champion, ok := query["champion"]; ok {
		championId, err := strconv.ParseUint(champion[0], 10, 32)
		if err != nil {
			return 0, errors.New(ErrorChampionIdInvalid)
		}
		return uint32(championId), nil
	}

	return 0, errors.New(ErrorChampionIdNotFound)
}

func ParsePatch(query map[string][]string) (*apb.PatchRange, error) {
	if patch, ok := query["patch"]; ok {
		patches := strings.Split(patch[0], "-")
		if len(patches) == 1 {
			patches = append(patches, patches[0])
		}

		return &apb.PatchRange{
			Min: patches[0],
			Max: patches[1],
		}, nil
	}

	return nil, errors.New(ErrorPatchNotFound)
}

func ParseTier(query map[string][]string) (*apb.TierRange, error) {
	if tier, ok := query["tier"]; ok {
		tiers := strings.Split(tier[0], "-")
		if len(tiers) == 1 {
			tiers = append(tiers, tiers[0])
		}
		fmt.Println(tier)
		fmt.Println(tiers)
		minTier, err := strconv.ParseUint(tiers[0], 0, 32)
		if err != nil {
			return nil, errors.New(ErrorTierInvalid)
		}

		maxTier, err := strconv.ParseUint(tiers[1], 0, 32)
		if err != nil {
			return nil, errors.New(ErrorTierInvalid)
		}

		return &apb.TierRange{
			Min: uint32(minTier),
			Max: uint32(maxTier),
		}, nil
	}

	return nil, errors.New(ErrorTierNotFound)
}
