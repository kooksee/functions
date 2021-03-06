package server

import (
	"context"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/iron-io/functions/api/models"
	"github.com/iron-io/runner/common"
)

func handleRouteList(c *gin.Context) {
	ctx := c.MustGet("ctx").(context.Context)
	log := common.Logger(ctx)

	filter := &models.RouteFilter{}

	if img := c.Query("image"); img != "" {
		filter.Image = img
	}

	var routes []*models.Route
	var err error
	if app := c.Param("app"); app != "" {
		routes, err = Api.Datastore.GetRoutesByApp(ctx, app, filter)
	} else {
		routes, err = Api.Datastore.GetRoutes(ctx, filter)
	}

	if err != nil {
		log.WithError(err).Error(models.ErrRoutesGet)
		c.JSON(http.StatusInternalServerError, simpleError(models.ErrRoutesGet))
		return
	}

	log.WithFields(logrus.Fields{"routes": routes}).Debug("Got routes")

	c.JSON(http.StatusOK, routesResponse{"Sucessfully listed routes", routes})
}
