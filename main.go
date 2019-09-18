package main

import (
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/context"
	"github.com/OrderSystem_WeiZhang/middleware/cors"
	"github.com/OrderSystem_WeiZhang/controllers/order"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	r := gin.Default()
	//Set Only allow Patch GET POST OPTION Method
	r.Use(cors.AllowAllOrigins())
	
	r.GET("/orders", order.GetOrders)
	r.POST("/orders", order.PostOrders)
	r.PATCH("/orders/:id",order.PatchOrders)

	http.ListenAndServe(":3000", context.ClearHandler(r))

}

