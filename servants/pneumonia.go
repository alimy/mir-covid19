package servants

import (
	api "github.com/alimy/mir-covid19/mirc/gen/api/ncovh5api/THPneumoniaService"
	"github.com/gin-gonic/gin"
)

type pneumonia struct {
	*baseServant
}

func (e *pneumonia) GetContents(c *gin.Context) {
	// TODO
}

func (e *pneumonia) GetAreaContents(c *gin.Context) {
	// TODO
}

// NewEpidemic return an Pneumonia service instance
func NewPneumonia() api.Pneumonia {
	return &pneumonia{
		baseServant: baseSrv,
	}
}
