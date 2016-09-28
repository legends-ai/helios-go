package server

import (
	"net/http"

	apb "github.com/asunaio/helios/gen-go/asuna"
	"golang.org/x/net/context"
)

type Controllers struct {
	Apollo  apb.ApolloClient
	Context context.Context
}

func (c *Controllers) HandleChampion(w http.ResponseWriter, r *http.Request) {
	champion, err := c.Apollo.GetChampion(c.Context, &apb.GetChampionRequest{
		ChampionId: 103,
		Patch: &apb.PatchRange{
			Min: "6.16",
			Max: "6.18",
		},
		Tier: &apb.TierRange{
			Min: 0x0000,
			Max: 0x1000,
		},
		Region: apb.Region_NA,
		Role:   apb.Role_JUNGLE,
	})

	if err != nil {
		Failure(w, r, err, http.StatusInternalServerError)
		return
	}

	Success(w, r, champion)
}

func (c *Controllers) HandleMatchup(w http.ResponseWriter, r *http.Request) {

}
