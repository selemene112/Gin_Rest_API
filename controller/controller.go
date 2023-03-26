package controller

import (
	// "API_PSQL/models"
	"database/sql"
	"log"
	"strconv"

	// "fmt"

	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Product struct{
	id int `json:"id"`
	name	string `json:"name"`
	email string `json:"email"`
	age string `json:"age"`
	division string `json:"division"`


}

var (
db *sql.DB
err error
)

// func Create(c *gin.Context){
// 	var Product = product{}

// 	sqlStatement := ` 
// 	INSERT INTO Product (Name, Email, Age, Division)
// 	VALUES ($1, $2, $3, $4)
// 	Return *
// 	`

// 	err = db.QueryRow(sqlStatement, "Ari","Ariyanda@gmail.com", 21 , "IT").
// 		Scan(&Product.id, &Product.Name, &Product.Email, &Product.Age, &Product.Division)

// 	if err != nil {

// 		panic(err)
// 	}
// 	fmt.Printf("new data update : %+v\n", Product)
// }

// func Show(c *gin.Context){
// 	var results = []product{}

// 	sqlStatement := `SELECT * FROM users`
// 	rows, err := db.Query(sqlStatement)
	
// 	if err !=nil {
// 		// fmt.Println("err =>", err)
// 		panic(err)
// 	}

// 	defer rows.Close()


// 	for rows.Next(){
// 		var Product = product{}

// 		err = rows.Scan(&Product.id,&Product.name, &Product.email, &Product.age, &Product.division)

// 		if err !=nil {
// 			panic(err)
// 		}
// 		results = append(results, Product)
		
// 	}

// 	fmt.Println("Data Product : ", results)

// }

// func Index(c *gin.Context){
// 	var Product []product
// 	db.Fine(&Product);
// 	c.JSON(http.StatusOK, gin.H{"Product": Product})

// }

// func Index(c *gin.Context) {
// 	anu := []Product{}
//     c.JSON(http.StatusOK, anu)
// 	return
// }

func Create(c *gin.Context) {
    var req Product
    if err := c.ShouldBindJSON(&req); err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    log.Printf("Received new user: %+v\n", req)
    c.JSON(http.StatusCreated, gin.H{"message": "User created"})
}


func Show(c *gin.Context){
	rows, err := db.Query("SELECT id, name, email, age, division FROM users")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query database"})
        return
    }
    defer rows.Close()

   
    var users []Product
    for rows.Next() {
        var user Product
        err := rows.Scan(&user.id, &user.name, &user.email, &user.age, &user.division)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan row"})
            return
        }
        users = append(users, user)
    }

    
    c.JSON(http.StatusOK, users)
}

func Getbyid(c *gin.Context) {
   
    userID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    

   
    row := db.QueryRow("SELECT id, name, email, age, division FROM users WHERE id = $1", userID)

    
    var user Product
    err = row.Scan(&user.id, &user.name, &user.email, &user.age, &user.division)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

   
    c.JSON(http.StatusOK, user)
}

func Update(c *gin.Context) {
    
    userID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

  
    var user Product
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    row := db.QueryRow("SELECT id FROM users WHERE id = $1", userID)
    var id int
    err = row.Scan(&id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

   
    _, err = db.Exec("UPDATE users SET name = $1, email = $2, age = $3, division = $4 WHERE id = $5", user.name, user.email, user.age, user.division, userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
        return
    }

    
    c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func Delete(c *gin.Context) {
    
    userID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    
    _, err = db.Exec("DELETE FROM users WHERE id = $1", userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
