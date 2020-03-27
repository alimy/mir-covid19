package servants

import (
	"github.com/alimy/mir-covid19/mirc/gen/api/ncovh5api"
	"github.com/gin-gonic/gin"
)

type epidemic struct {
	// TODO
}

func (e *epidemic) GetContents(c *gin.Context) {
	// TODO
}

func (e *epidemic) GetAreaContents(c *gin.Context) {
	// TODO
}

func (e *epidemic) GetInfoBatch(c *gin.Context) {
	// TODO
}

func (e *epidemic) GetChinaTotal(c *gin.Context) {
	// TODO
}

func (e *epidemic) GetAreaInfo(c *gin.Context) {
	// TODO
}

func (e *epidemic) GetForeignTotal(c *gin.Context) {
	// TODO
}

func (e *epidemic) GetForeignInfo(c *gin.Context) {
	// TODO
}

func (e *epidemic) GetForeignHistory(c *gin.Context) {
	// TODO
}

func (e *epidemic) GetHubeiInfo(c *gin.Context) {
	// TODO
}

func (e *epidemic) GetRate(c *gin.Context) {
	// TODO
}

func (e *epidemic) GetCityInfoByCode(c *gin.Context) {
	// TODO
}

// NewEpidemic return an ncovh5api.Epidemic instance
func NewEpidemic() ncovh5api.Epidemic {
	return &epidemic{}
}
