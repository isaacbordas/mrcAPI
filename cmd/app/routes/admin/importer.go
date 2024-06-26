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
	case "genres":
		importGenres(c, importerProvider)
	case "developers":
		importDevelopers(c, importerProvider)
	case "publishers":
		importPublishers(c, importerProvider)
	case "gameByName":
		name := c.Param("name")
		importGameByName(c, importerProvider, name)
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

func importGenres(c *gin.Context, i importer.Importer) {
	errImport := i.ImportGenres()
	if errImport != nil {
		c.JSON(http.StatusBadRequest, errImport.Error())
	}

	c.JSON(http.StatusOK, nil)
}

func importDevelopers(c *gin.Context, i importer.Importer) {
	errImport := i.ImportDevelopers()
	if errImport != nil {
		c.JSON(http.StatusBadRequest, errImport.Error())
	}

	c.JSON(http.StatusOK, nil)
}

func importPublishers(c *gin.Context, i importer.Importer) {
	errImport := i.ImportPublishers()
	if errImport != nil {
		c.JSON(http.StatusBadRequest, errImport.Error())
	}

	c.JSON(http.StatusOK, nil)
}

func importGameByName(c *gin.Context, i importer.Importer, name string) {
	games, errImport := i.ImportGameByName(name)
	if errImport != nil {
		c.JSON(http.StatusBadRequest, errImport.Error())
	}

	c.JSON(http.StatusOK, games)
}
