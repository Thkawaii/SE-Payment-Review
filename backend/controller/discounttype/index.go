package discounttype


import (
	"net/http"
	"SE/config"
	"SE/entity"

	"github.com/gin-gonic/gin"
)


func GetAllD(c *gin.Context) {


   db := config.DB()


   var discounttype []entity.DiscountType

   db.Find(&discounttype)


   c.JSON(http.StatusOK, &discounttype)


}