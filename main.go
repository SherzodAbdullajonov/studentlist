//Package classificaton of Studen API
//
//Documentation for Student API
//
//
//Schemes:http
// BasePath:/
//Version: 1.0.0
//
//Consume:
//-aplication/json
//
//
// Students:
// - aplication/json
//
//swagger: meta
package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"studentList/models"

	"github.com/gin-gonic/gin"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var MPosDB *sql.DB
var MPosGORM *gorm.DB

func InitGormPostgres() {
	MPosGORM, err = gorm.Open("postgres", "user=postgres dbname=studentlist password=sherzod sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
}

// var (
// 	student  = &models.Student{Name: "Sherzod", Surname: "Abdullajonov", Id: 1, Course: 4, Department: "Socie", Adress: "Fergana", Phone: 901666989}
// 	student2 = &models.Student{Name: "Khurshid", Surname: "Kabilov", Id: 2, Course: 4, Department: "Socie", Adress: "Fergana", Phone: 911146761}
// 	student3 = &models.Student{Name: "Shokhruhk", Surname: "Gafurov", Id: 3, Course: 4, Department: "Socie", Adress: "Fergana", Phone: 916771185}
// 	student4 = &models.Student{Name: "Jahongir", Surname: "Makhkamov", Id: 4, Course: 4, Department: "Socie", Adress: "Fergana", Phone: 906330463}
// 	student5 = &models.Student{Name: "Shahzod", Surname: "Burkhanov", Id: 5, Course: 4, Department: "Socie", Adress: "Jizzakh", Phone: 995558076}
// 	student6 = &models.Student{Name: "Jaloliddin", Surname: "Abdullayev", Id: 6, Course: 4, Department: "Socie", Adress: "Jizzakh", Phone: 994522399}
// )
var db *gorm.DB
var err error

func main() {
	dialect := os.Getenv("DIALECT")
	host := os.Getenv("HOST")
	dbPort := os.Getenv("DBPORT")
	user := os.Getenv("USER")
	dbname := os.Getenv("NAME")
	password := os.Getenv("PASSWORD")
	//Database connection string
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, dbname, password, dbPort)
	//Openning connection to database
	db, err = gorm.Open(dialect, dbURI)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Successfully connected to database")
	}
	defer db.Close()

	InitGormPostgres()
	defer MPosGORM.Close()

	// Set the router as the default one shipped with Gin
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	//api := router.Group("/api")
	// db.AutoMigrate(&models.Student{})
	// db.Create(&student)
	// db.Create(&student2)
	// db.Create(&student3)
	// db.Create(&student4)
	// db.Create(&student5)
	// db.Create(&student6)
	router.GET("/", GetStudents)
	router.GET("/student", GetStudent)
	router.GET("/student/:id", GetStudentById)
	router.POST("/student", CreateStudent)
	router.PUT("/student/:id", UpdateStudent)
	router.DELETE("/student/:id", DeleteStudent)

	log.Fatal(router.Run(":4000"))

}
func GetStudents(c *gin.Context) {
	var student []models.Student
	if err := db.Find(&student).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, student)
	}
}
func GetStudentById(c *gin.Context) {
	id := c.Params.ByName("id")
	var student models.Student
	if err := db.Where("id = ?", id).First(&student).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, student)
	}
}
func GetStudent(c *gin.Context) {
	var student []models.Student
	if err := db.Find(&student).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, student)
	}
}
func CreateStudent(c *gin.Context) {
	var student models.Student
	c.BindJSON(&student)
	db.Create(&student)
	c.JSON(200, student)
}
func UpdateStudent(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	if err := db.Where("id = ?", id).First(&student).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&student)
	db.Save(student)
	c.JSON(200, student)
}

func DeleteStudent(c *gin.Context) {
	id := c.Params.ByName("id")
	var student models.Student
	d := db.Where("id = ?", id).Delete(&student)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
