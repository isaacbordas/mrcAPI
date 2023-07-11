package admin

import (
	"github.com/gin-gonic/gin"
	"mrcAPI/pkg/errors"
	"mrcAPI/pkg/importer"
	"mrcAPI/pkg/importer/wire"
	"net/http"
)

func ImportByEntity(c *gin.Context) {
	importerProvider, err := wire.ProvideImport()
	if err != nil {
		panic(err)
	}

	entity := c.Param("entity")

	switch entity {
	case "platforms":
		importPlatforms(c, importerProvider)
	default:
		panic(errors.ErrEntityNotAllowed{Entity: entity})
	}
}

func importPlatforms(c *gin.Context, i importer.Importer) {
	errImport := i.ImportPlatforms()
	if errImport != nil {
		c.JSON(http.StatusBadRequest, errImport.Error())
	}

	c.JSON(http.StatusOK, nil)
}
