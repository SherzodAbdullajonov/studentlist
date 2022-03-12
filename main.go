package main

import (
	"fmt"
	"log"
	"os"
	"studentList/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	_ "studentList/doc/studentlist"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var db *gorm.DB
var err error

// @title           Gin Swagger Example API
// @version         1.0
// @description     This is a sample server server.
// @termsOfService  http://swagger.io/terms/
// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @host      localhost:4000
// @BasePath  /
// @schemes   []string{"http"}
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

	// Set the router as the default one shipped with Gin
	gin.SetMode(gin.ReleaseMode)
	//router := gin.Default()
	//api := router.Group("/api")
	// db.AutoMigrate(&models.Student{})
	// db.Create(&student)
	// db.Create(&student2)
	// db.Create(&student3)
	// db.Create(&student4)
	// db.Create(&student5)
	// db.Create(&student6)

	r := gin.Default()
	url := ginSwagger.URL("http://localhost:4000/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	r.GET("/student", GetStudents)
	r.GET("/student/:id", GetStudentById)
	r.POST("/student", CreateStudent)
	r.PUT("/student/:id", UpdateStudent)
	r.DELETE("/student/:id", DeleteStudent)

	log.Fatal(r.Run(":4000"))
	//log.Fatal(router.Run(":4000"))

}

// GetStudents godoc
// @Summary      Show student list.
// @ID get-all-Students
// @Description  get all students from the database.
// @Tags         Students
// @Accept       json
// @Produce      json
// @Success      200  {struct}  models.Student
// @Failure      404  {object}  models.Student
// @Router       /student [get]
func GetStudents(c *gin.Context) {
	var student []models.Student
	if err := db.Find(&student).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, student)
	}
}

// GetStudentById godoc
// @Summary      Show a  student from the list.
// @ID get-Student-by-id
// @Description  get one student by id from the database.
// @Tags         Students
// @Accept       json
// @Produce      json
// @Param  	     id path string true "Student ID"
// @Success      200  {struct}  models.Student
// @Failure      404  {object}  models.Student
// @Router       /student/:id [get]
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

// CreateStudent godoc
// @Summary      Create a new student.
// @Description  create a student and add to the database.
// @Tags         Students
// @Accept       json
// @Produce      json
// @Param 		 models.Student
// @Success      200  {object}  models.Student
// @Failure      404  {object}  models.Student
// @Router       /student [post]
func CreateStudent(c *gin.Context) {
	var student models.Student
	c.BindJSON(&student)
	db.Create(&student)
	c.JSON(200, student)
}

// UpdtadeStudent godoc
// @Summary      Update a student.
// @Description  update an existing student by ID.
// @Tags         Students
// @Accept       json
// @Produce      json
// @Success      200  {struct}  models.Student
// @Failure      404  {object}  models.Student
// @Router       /student/:id [put]
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

// DeleteStudent godoc
// @Summary      Delete a student.
// @Description  delete an existing student by ID.
// @Tags         Students
// @Accept       json
// @Produce      json
// @Success      200  {struct}  models.Student
// @Failure      404  {object}  models.Student
// @Router       /student/:id [delete]
func DeleteStudent(c *gin.Context) {
	id := c.Params.ByName("id")
	var student models.Student
	d := db.Where("id = ?", id).Delete(&student)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}

// var (
// 	student  = &models.Student{Name: "Sherzod", Surname: "Abdullajonov", Id: 1, Course: 4, Department: "Socie", Adress: "Fergana", Phone: 901666989}
// 	studen4000/swagger/doc.jsont2 = &models.Student{Name: "Khurshid", Surname: "Kabilov", Id: 2, Course: 4, Department: "Socie", Adress: "Fergana", Phone: 911146761}
// 	student3 = &models.Student{Name: "Shokhruhk", Surname: "Gafurov", Id: 3, Course: 4, Department: "Socie", Adress: "Fergana", Phone: 916771185}
// 	student4 = &models.Student{Name: "Jahongir", Surname: "Makhkamov", Id: 4, Course: 4, Department: "Socie", Adress: "Fergana", Phone: 906330463}
// 	student5 = &models.Student{Name: "Shahzod", Surname: "Burkhanov", Id: 5, Course: 4, Department: "Socie", Adress: "Jizzakh", Phone: 995558076}
// 	student6 = &models.Student{Name: "Jaloliddin", Surname: "Abdullayev", Id: 6, Course: 4, Department: "Socie", Adress: "Jizzakh", Phone: 994522399}
// )
