package Controllers
import (
	"encoding/json"
	"fmt"
	"github.com/Day4-solution/Config"
	"github.com/Day4-solution/Models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"sync"
	"time"
)
// Get all customers
func GetCustomers(c *gin.Context) {
	var customer []Models.Customer
	err := Models.GetAllCustomers(&customer)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, customer)
	}
}
//Insert a new customer
func CreateCustomer(c *gin.Context) {
	var customer Models.Customer
	c.BindJSON(&customer)
	customer.Wait=true
	err := Models.CreateCustomer(&customer)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, customer)
	}
}
// Get customer using id
func GetCustomerByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var customer Models.Customer
	err := Models.GetCustomerByID(&customer, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, customer)
	}
}
// update values of customer
func UpdateCustomer(c *gin.Context) {
	var customer Models.Customer
	id := c.Params.ByName("id")
	err := Models.GetCustomerByID(&customer, id)
	if err != nil {
		c.JSON(http.StatusNotFound, customer)
	}
	c.BindJSON(&customer)
	err = Models.UpdateCustomer(&customer, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, customer)
	}
}
// Delete the customer row
func DeleteCustomer(c *gin.Context) {
	var customer Models.Customer
	id := c.Params.ByName("id")
	err := Models.DeleteCustomer(&customer, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}

// Get all orders of customer
func GetAllOrders(c *gin.Context) {
	var order []Models.Order
	err := Models.GetAllOrders(&order)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
	}
}


//  Create order for a customer
func CreateOrder(c *gin.Context) {
	var order Models.Order
	c.BindJSON(&order)
	err := Models.CreateOrder(&order)
	if err!=nil{
		fmt.Println("error.")
	}
	var No_of_Buying_quantity= order.Quantity
	var product Models.Product
	var customer Models.Customer
	err_:=Models.GetProductByID(&product,strconv.FormatUint(uint64(order.ProdId), 10))
	if err_ != nil {
		fmt.Println("Error error",http.StatusNotFound)
	}
	error_:=Models.GetCustomerByID(&customer,strconv.FormatUint(uint64(order.CustId), 10))
	if error_ != nil {
		fmt.Println("Error error",http.StatusNotFound)
	}
	var mutex = &sync.Mutex{}
	mutex.Lock()
	if product.Quantity < No_of_Buying_quantity && customer.Wait==false{
		order.Status = "Failed"
		prod,_ := json.Marshal(product)
		fmt.Println(string(prod),order.ProdId)
	} else{
		order.Status = "Order Placed"
		Remaining_quantity := product.Quantity-No_of_Buying_quantity
		//Config.DB.Save(product)
		Config.DB.Model(&product).Where("prod_id = ?", product.ProdId).Update("quantity", Remaining_quantity)
	}
	mutex.Unlock()
	Config.DB.Model(&customer).Where("cust_id = ?", customer.CustId).Update("wait", false)
	time.Sleep(100*time.Second)
	Config.DB.Model(&customer).Where("cust_id = ?", customer.CustId).Update("wait", true)
}

// Get Customer by id
func GetOrderForCustomerID(c *gin.Context) {
	id := c.Params.ByName("cust_id")
	var order Models.Order
	err := Models.GetOrderForCustomerID(&order, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
	}
}

//  Update the order information
func UpdateOrder(c *gin.Context) {
	var order Models.Order
	id := c.Params.ByName("id")
	err := Models.GetOrderByID(&order, id)
	if err != nil {
		c.JSON(http.StatusNotFound, order)
	}
	error_ := c.BindJSON(&order)
	if error_ != nil {
		fmt.Println(error_.Error())
	}
	err = Models.UpdateOrder(&order, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
	}
}
// Groutine for the each customer
func GoroutineForOrders(c *gin.Context){
	go CreateOrder(c)
}

func GetProducts(c *gin.Context) {
	var product []Models.Product
	err := Models.GetProducts(&product)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}

//  Create product
func CreateProduct(c *gin.Context) {
	var product Models.Product
	c.BindJSON(&product)
	err := Models.CreateProduct(&product)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}

// Get product by id
func GetProductByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var product Models.Product
	err := Models.GetProductByID(&product, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}

// Update the product information
func UpdateProduct(c *gin.Context) {
	var product Models.Product
	id := c.Params.ByName("id")
	err := Models.GetProductByID(&product, id)
	if err != nil {
		c.JSON(http.StatusNotFound, product)
	}
	error_ := c.BindJSON(&product)
	if error_ != nil {
		fmt.Println(error_.Error())
	}
	err = Models.UpdateProduct(&product, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}

// Delete the product record
func DeleteProduct(c *gin.Context) {
	var product Models.Product
	id := c.Params.ByName("id")
	err := Models.DeleteProduct(&product, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}