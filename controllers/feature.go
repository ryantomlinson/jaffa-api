package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ryantomlinson/jaffa-api/model"
)

type FeatureController struct{}

var featureModel = new (model.Feature)

func(ctrl FeatureController) GetAllFeatures(c *gin.Context) {
	features, err := featureModel.All()

	if err != nil {
		c.JSON(406, gin.H{"Message": "Could not find any features", "error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"data": features})
}