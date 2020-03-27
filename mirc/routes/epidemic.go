package routes

import (
	"github.com/alimy/mir/v2"
	"github.com/alimy/mir/v2/engine"
)

func init() {
	engine.AddEntry(new(Epidemic))
}

// Epidemic contents fetch interface define
type Epidemic struct {
	Group             mir.Group `mir:"ncovh5api"`
	GetContents       mir.Post  `mir:"THPneumoniaService/getContents"`
	GetAreaContents   mir.Post  `mir:"THPneumoniaService/getAreaContents"`
	GetInfoBatch      mir.Post  `mir:"THPneumoniaOuterDataService/getInfoBatch"`
	GetChinaTotal     mir.Post  `mir:"THPneumoniaOuterDataService/getChinaTotal"`
	GetAreaInfo       mir.Post  `mir:"THPneumoniaOuterDataService/getAreaInfo"`
	GetForeignTotal   mir.Post  `mir:"THPneumoniaOuterDataService/getForeignTotal"`
	GetForeignInfo    mir.Post  `mir:"THPneumoniaOuterDataService/getForeignInfo"`
	GetForeignHistory mir.Post  `mir:"THPneumoniaOuterDataService/getForeignHistory"`
	GetHubeiInfo      mir.Post  `mir:"THPneumoniaOuterDataService/getHubeiInfo"`
	GetRate           mir.Post  `mir:"THPneumoniaOuterDataService/getRate"`
	GetCityInfoByCode mir.Post  `mir:"THPneumoniaOuterDataService/getCityInfoByCode"`
}
