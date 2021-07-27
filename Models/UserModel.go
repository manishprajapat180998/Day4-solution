package Models

type Customer struct {
	CustId  	uint   `json:"cust_id"`
	CustName    string `json:"cust_name"`
	Wait 		bool    `json:"wait"`
}
func (b *Customer) TableName() string {
	return "Customer"
}

type Order struct {
	OrderId   uint   `json:"order_id" `
	CustId    uint   `json:"cust_id"`
	ProdId    uint   `json:"prod_id"`
	Status    string `json:"status"`
	Quantity  int    `json:"quantity"`
}
func (o *Order) TableName() string {
	return "order"
}
type Product struct {
	ProdId     	uint 	`json:"prod_id" `
	ProdName   	string 	`json:"prod_name"`
	Price      	int  	`json:"price"`
	Quantity  	int  	`json:"quantity"`
	Message		string	`json:"message"`
}
func (p *Product) TableName() string {
	return "product"
}