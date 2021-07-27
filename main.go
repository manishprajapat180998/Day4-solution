package main
import (
	"fmt"
	"github.com/Day4-solution/Config"
	"github.com/Day4-solution/Models"
	"github.com/Day4-solution/Routes"
	"github.com/jinzhu/gorm"
)
var err error
func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Customer{},&Models.Order{},Models.Product{})
	r := Routes.SetupRouter()
	//running
	r.Run()
}