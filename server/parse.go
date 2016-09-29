package server

import (
	"errors"
	"strconv"
	"strings"

	apb "github.com/asunaio/helios/gen-go/asuna"
)

func ParseChampionId(query map[string][]string, field string) (uint32, error) {
	if id, ok := query[field]; ok {
		championId, err := strconv.ParseUint(id[0], 10, 32)
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

func ParseRegion(query map[string][]string) (apb.Region, error) {
	if region, ok := query["region"]; ok {
		switch region[0] {
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
			return apb.Region_UNKNOWN_REGION, errors.New(ErrorRegionInvalid)
		}
	}

	return apb.Region_UNKNOWN_REGION, errors.New(ErrorRegionNotFound)
}

func ParseRole(query map[string][]string) (apb.Role, error) {
	if role, ok := query["role"]; ok {
		switch role[0] {
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
			return apb.Role_UNKNOWN_ROLE, errors.New(ErrorRoleInvalid)
		}
	}
	return apb.Role_UNKNOWN_ROLE, errors.New(ErrorRoleNotFound)
}
