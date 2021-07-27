package Routes
import (
	"github.com/Day4-solution/Controllers"
	"github.com/gin-gonic/gin"
)
//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/customer-api")
	{
		grp1.GET("customer", Controllers.GetCustomers)
		grp1.POST("customer", Controllers.CreateCustomer)
		grp1.GET("customer/:id", Controllers.GetCustomerByID)
		grp1.PUT("customer/:id", Controllers.UpdateCustomer)
		grp1.DELETE("customer/:id", Controllers.DeleteCustomer)
	}
	grp2 := r.Group("/product-api")
	{
		grp2.GET("product", Controllers.GetProducts)
		grp2.POST("product", Controllers.CreateProduct)
		grp2.GET("product/:id", Controllers.GetProductByID)
		grp2.PUT("product/:id", Controllers.UpdateProduct)
		grp2.DELETE("product/:id", Controllers.DeleteProduct)
	}
	grp3 := r.Group("/order-api")
	{
		grp3.GET("order", Controllers.GetAllOrders)
		grp3.POST("order", Controllers.GoroutineForOrders)
		grp3.GET("order/:id", Controllers.GetOrderForCustomerID)
		grp3.PUT("order/:id", Controllers.UpdateOrder)
	}

	return r
}