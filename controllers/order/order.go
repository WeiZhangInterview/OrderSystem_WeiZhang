package order

import (
	"strconv"
	"github.com/Order_system/db/db"
	"github.com/gin-gonic/gin"
)
/*
	if c.Request.URL.Query()["q"] == nil {
		resp.Err = "ERR_QUERY"
		c.JSON(200, resp)
		return
	}
	drugName := c.Request.URL.Query()["q"][0]
*/

func getPageCount(c *gin.Context) (pg int, cnt int) {
	page := 1
	count := 20

	if c.Request.URL.Query()["page"] != nil {
		page, _ = strconv.Atoi(c.Request.URL.Query()["page"][0])
	}
	if c.Request.URL.Query()["count"] != nil {
		count, _ = strconv.Atoi(c.Request.URL.Query()["count"][0])
	}

	return page, count
}
func PostOrders(c *gin.Context) {
	c.JSON(200, map[string]bool{"success": true})
}

func PatchOrders(c *gin.Context) {
	c.JSON(200, map[string]bool{"success": true})
}

func GetOrders(c *gin.Context) {
	order := db.test()
	c.JSON(200, map[string]string{"success": order})
}


