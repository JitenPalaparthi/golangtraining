// /*package main

// import (
// 	"fmt"
// 	"errors"
// )

// type Employee struct {
// 	FirstName, LastName string
// }

// type EmployeeError struct {
// 	Details string
// }

// func (er *EmployeeError) Error() string {
// 	//er.Details ="Seems Firstname or Lastname is not provided"
// 	return "Some problem in creating Employee instance"
// }

// func New(fn, ln string) (emp *Employee, err error) {

// 	// return nil,fmt.Errorf("")

// 	// return nil, fmt.Errorf("some error")

// 	// return nil errors.New("some issue")

// 	if fn == "" && ln == "" {
// 		return nil, &EmployeeError{Details: "FirstName and LastName seems to be empty string"}
// 	}

// 	if fn == "" {
// 		return nil, &EmployeeError{Details: "FirstName seems to be empty string"}
// 	}

// 	if ln == "" {
// 		return nil, &EmployeeError{Details: "LastName seems to be empty string"}
// 	}

// 	return &Employee{FirstName: fn, LastName: ln}, nil
// }

// func main() {

// 	emp, err := New("Jiten", "P")
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println(emp)
// 	}

// 	emp1, err := New("", "")

// 	fmt.Println(emp1, err.(*EmployeeError).Details)

// 	emp2, err := New("Jiten", "")

// 	fmt.Println(emp2, err.(*EmployeeError).Details)

// 	emp3, err := New("", "P")

// 	fmt.Println(emp3, err.(*EmployeeError).Details)

// 	_, err = New("J", "")

// 	fmt.Println(err)

// }
// */
// // package main

// // import (
// // 	"fmt"
// // )

// // func fullName(firstName *string, lastName *string) {
// // 	defer fmt.Println("1- deferred call in fullName") //3
// // 	defer fmt.Println("2- deferred call in fullName") //3
// // 	if firstName == nil {
// // 		func() {
// // 			panic("runtime error: first name cannot be nil")
// // 		}()
// // 	}
// // 	if lastName == nil {
// // 		func() {
// // 			panic("runtime error: last name cannot be nil")
// // 		}()
// // 	}
// // 	fmt.Printf("%s %s\n", *firstName, *lastName)
// // 	fmt.Println("returned normally from fullName")

// // }

// // func main() {
// // 	defer fmt.Println("deferred call in main")     //1
// // 	defer fmt.Println("2nd deffered call in main") //2
// // 	firstName := "Elon"
// // 	//lastname :=  "Musk"
// // 	fullName(&firstName, nil)
// // 	fmt.Println("returned normally from main")
// // }

// // net/http
// // gorilla mux
// // beego
// // gin

// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"time"

// 	_ "github.com/gin-gonic/gin"
// )

// type myWriter struct{}

// func (mw myWriter) Write(data []byte) (int, error) {
// 	fmt.Println(string(data))
// 	return 0, nil
// }

// func main() {

// 	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintln(w, "pong")
// 	})

// 	//http.HandleFunc("/ping1", Ping)

// 	http.HandleFunc("/", HelloWorld)

// 	fmt.Println("Web Server started on port 8080")

// 	http.ListenAndServe(":8080", nil)
// }

// func HelloWorld(w http.ResponseWriter, r *http.Request) {
// 	/*f, err := os.OpenFile("hello.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
// 	defer f.Close()
// 	if err != nil {
// 		panic(err)
// 	}*/
// 	//mw := myWriter{}
// 	fmt.Fprintf(w, "Hello World! %s", fmt.Sprintln(time.Now().Unix()))
// }

// func Ping(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "pong")
// }

/*package main

import (
	"net/http"
	_ "time"










	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// This handler will match /user/john but will not match /user/ or /user
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	// However, this one will match /user/john/ and also /user/john/send
	// If no other routers match /user/john, it will redirect to /user/john/
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	// For each matched request Context will hold the route definition
	router.POST("/user/:name/*action", func(c *gin.Context) {
		b := c.FullPath() == "/user/:name/*action" // true
		c.String(http.StatusOK, "%t", b)
	})

	// This handler will add a new router for /user/groups.
	// Exact routes are resolved before param routes, regardless of the order they were defined.
	// Routes starting with /user/groups are never interpreted as /user/:name/... routes
	router.GET("/user/groups", func(c *gin.Context) {
		c.String(http.StatusOK, "The available groups are [...]")
	})

	router.Run(":8080")
}
*/

// package main

// import (
// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )
// type Product struct {
// 	gorm.Model
// 	Code  string
// 	Price uint
// }
// func main() {
// 	dsn := "root:admin@tcp(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 		if err != nil {
// 		panic("failed to connect database")
// 	}
// 	// Migrate the schema
// 	db.AutoMigrate(&Product{})
// 	// Create
// 	db.Create(&Product{Code: "D42", Price: 100})
// 	// Read
// 	var product Product
// 	db.First(&product, 1)                 // find product with integer primary key
// 	db.First(&product, "code = ?", "D42") // find product with code D42
// 	// Update - update product's price to 200
// 	db.Model(&product).Update("Price", 200)
// 	// Update - update multiple fields
// 	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
// 	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
// 	// Delete - delete product
// 	db.Delete(&product, 1)
// }

package main

import (
	"flag"
	"os"

	"example/database"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

var (
	DBConnection = "root:admin@tcp(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"
	DBName       = "demo"
)

func main() {
	portPtr := flag.String("port", "8080", "give port=8080 or preffered port")
	flag.Parse()

	if os.Getenv("DB_CONNECTION") != "" {
		DBConnection = os.Getenv("DB_CONNECTION")
	}
	if os.Getenv("DB_NAME") != "" {
		DBName = os.Getenv("DB_NAME")
	}
	glog.Infoln("Application has been started using port:", *portPtr)

	_, err := database.GetConnection(DBConnection, DBName)
	if err != nil {
		glog.Fatalln("Database Error:", err)
	} else {
		glog.Infoln("Successfully connected to the database")
	}

	glog.Flush()

	gin.ForceConsoleColor()

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run(":" + *portPtr)

}
