package order

import (
	"log"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/OrderSystem_WeiZhang/models"
)



func PostOrders(c *gin.Context) {
	var resultErr model.ResultErr
	var orderRequest model.OrderRequest
	errBinding := c.BindWith(&orderRequest, binding.JSON)
	if errBinding != nil {
		log.Fatal(errBinding.Error())
		resultErr.Error = "ERR_INPUT"
		c.JSON(400, resultErr)
		return	
	}

	err := validateCoordinateInput(orderRequest)
	if err != nil {
		log.Fatal(err.Error())
		resultErr.Error = "ERR_INPUT"
		c.JSON(400, resultErr)
		return			
	}

	c.JSON(200, map[string]bool{"success": true})
	return
}

func PatchOrders(c *gin.Context) {
	var resultErr model.ResultErr

	//Get Order Id
	orderId, errParse := strconv.Atoi(c.Params.ByName("id"))
	if errParse != nil {
		resultErr.Error = "ERR_PARAMETER"
		c.JSON(400, resultErr)
		return	
	}

	//Check Order Id with Row
	var order model.OrderRespond
	order.Id = orderId
	effectRow , errTx := order.Update()
	if errTx != nil || effectRow != 1 {
		resultErr.Error = "ERR_ORDER_BE_TAKEN"
		c.JSON(409, resultErr)
		return			
	}

	c.JSON(200, map[string]string{"status":"SUCCESS"})
}

func GetOrders(c *gin.Context) {
	var resultErr model.ResultErr
	var order model.OrderRespond
	var returnOrders []model.OrderRespond
	var err error

	//page and limit default value set to 1 and 20
	page, limit := 1, 20
	if c.Request.URL.Query()["page"] != nil {
		page, err = strconv.Atoi(c.Request.URL.Query()["page"][0])
		if err != nil || page <= 0 {
			resultErr.Error = "ERR_PARAMETER"
			c.JSON(400, resultErr)
			return		
		}
	}
	if c.Request.URL.Query()["limit"] != nil {
		limit, err = strconv.Atoi(c.Request.URL.Query()["limit"][0])
		if err != nil || limit <= 0 {
			resultErr.Error = "ERR_PARAMETER"
			c.JSON(400, resultErr)
			return		
		}
	}

	//Get all record from database
	//NOTE: return slice already sort by primary key Id
	var orders []model.OrderRespond
	orders, err = order.GetAll()
	if err != nil {
		log.Println(err)
		resultErr.Error = "ERR_DATA"
		c.JSON(400, resultErr)
		return		
	}


	//Get Slice lenght and return page count part of slice
	if len(orders) > (page -1) * limit {
		if len(orders) > page * limit {
			returnOrders = orders[(page-1) * limit : page * limit]
			}else {
				returnOrders = orders[(page-1) * limit:]
			}
	}else{
		returnOrders = orders[0:0]
	}

	c.JSON(200, returnOrders)
}


