package Models
import (
	"fmt"
	"github.com/Day4-solution/Config"
	_ "github.com/go-sql-driver/mysql"
)

// Get product details from user data
func GetProducts(product *[]Product) (err error) {
	if err = Config.DB.Find(product).Error; err != nil {
		return err
	}
	return nil
}

// Create Product
func CreateProduct(product *Product) (err error) {
	if err = Config.DB.Create(product).Error; err != nil {
		return err
	}
	return nil
}

// Get product using a product id
func GetProductByID(product *Product, id string) (err error) {
	if err = Config.DB.Where("prod_id = ?", id).First(product).Error; err != nil {
		return err
	}
	return nil
}

// Update product
func UpdateProduct(product *Product, id string) (err error) {
	fmt.Println(product)
	Config.DB.Save(product)
	return nil
}

// Delete a row from product table
func DeleteProduct(product *Product, id string) (err error) {
	Config.DB.Where("prod_id = ?", id).Delete(product)
	return nil
}

// Show all the orders
func GetAllOrders(order *[]Order) (err error) {
	if err = Config.DB.Find(order).Error; err != nil {
		return err
	}
	return nil
}

// New data insertion in order table
func CreateOrder(order *Order) (err error) {
	if err = Config.DB.Create(order).Error; err != nil {
		return err
	}
	return nil
}

// Get order for customer id
func GetOrderForCustomerID(order *Order, id string) (err error) {
	if err = Config.DB.Where("cust_id = ?", id).Error; err != nil {
		return err
	}
	return nil
}

// Update order data
func UpdateOrder(order *Order, id string) (err error) {
	fmt.Println(order)
	Config.DB.Save(order)
	return nil
}

// Get order details for a particular id
func GetOrderByID(order *Order, id string) (err error) {
	if err = Config.DB.Where("order_id = ?", id).First(order).Error; err != nil {
		return err
	}
	return nil
}

// Get all user data
func GetAllCustomers(customer *[]Customer) (err error) {
	if err = Config.DB.Find(customer).Error; err != nil {
		return err
	}
	return nil
}
// Insert a new data
func CreateCustomer(customer *Customer) (err error) {
	if err = Config.DB.Create(customer).Error; err != nil {
		return err
	}
	return nil
}
// Get customer data using id
func GetCustomerByID(customer *Customer, id string) (err error) {
	if err = Config.DB.Where("cust_id = ?", id).First(customer).Error; err != nil {
		return err
	}
	return nil
}
// Update customer data
func UpdateCustomer(customer *Customer, id string) (err error) {
	fmt.Println(customer)
	Config.DB.Save(customer)
	return nil
}
// Delete customer data
func DeleteCustomer(customer *Customer, id string) (err error) {
	Config.DB.Where("cust_id = ?", id).Delete(customer)
	return nil
}