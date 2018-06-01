package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/remoteview/service-blocks/common"
	"github.com/remoteview/service-blocks/models"
)

// @Summary List blocks
// @Description List blocks
// @Accept  json
// @Produce  json
// @Router /blocks [get]
func ListBlocksHandler(c *gin.Context) {
	db := common.GetDB()
	var data []models.Block
	err := db.All(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "json decoding : " + err.Error(),
			"status": http.StatusBadRequest,
		})
		return
	}
	if len(data) > 0 {
		c.JSON(200, data)
	} else {
		c.JSON(200, make([]string, 0))
	}
}

// @Summary Get block
// @Description Get block
// @Accept  json
// @Produce  json
// @Router /blocks/{id} [get]
func GetBlockHandler(c *gin.Context) {
	id := c.Params.ByName("id")
	var block models.Block
	if err := common.GetDB().Find(&block, id); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, block)
	}
}

// @Summary Create block
// @Description List blocks
// @Accept  json
// @Produce  json
// @Router /blocks [post]
func CreateBlockHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json; charset=utf-8")

	var err error
	var block models.Block

	if err = c.BindJSON(&block); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "json decoding : " + err.Error(),
			"status": http.StatusBadRequest,
		})
		return
	}
	if err = common.GetDB().Create(&block); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "json decoding : " + err.Error(),
			"status": http.StatusBadRequest,
		})
		return
	}
	c.JSON(201, gin.H{"success": block})
}

// @Summary Delete block
// @Description Delete block
// @Accept  json
// @Produce  json
// @Router /blocks/{id} [delete]
func DeleteBlockHandler(c *gin.Context) {
	id := c.Params.ByName("id")
	var block models.Block
	if err := common.GetDB().Find(&block, id); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		if err = common.GetDB().Destroy(&block); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":  "json decoding : " + err.Error(),
				"status": http.StatusBadRequest,
			})
			return
		}
	}
}
