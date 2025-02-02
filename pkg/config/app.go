// thie purpose of this file is to return a variable db so that other packages can interact with the database
package config
import(
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
 
var db * gorm.DB

// Connect establishes a connection to the MySQL database using GORM.
// It takes in the connection string with format: "username:password@tcp(host:port)/databaseName?param=value".
// If the connection is successful, it assigns the database connection to the global variable `db`.

func Connect() {
	d, err := gorm.Open("mysql", "ansh1:12345@unix(/tmp/mysql.sock)/myTable?charset=utf8&parseTime=True&loc=Local")
    if err != nil {
        panic(err)
    }
    db = d
}


func GetDB() *gorm.DB{
	return db
}