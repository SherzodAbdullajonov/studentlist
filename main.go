package main

import (
	"fmt"
	"log"
	"os"
	"studentList/models"

	_ "studentList/doc/studentlist"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var db *gorm.DB
var err error

// @title API document title
// @version version(1.0)
// @description Description of specifications
// @Precautions when using termsOfService specifications

// @contact.name API supporter
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name license(Mandatory)
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:4000
// @BasePath /
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
	// db.AutoMigrate(&models.Student{})

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
// @Param  	     id path int true "id"
// @Success      200  {object}  models.Student
// @Failure      404  {object}  models.Student
// @Router       /student/{id} [get]
func GetStudentById(c *gin.Context) {
	id := c.Params.ByName("id")
	log.Println("ID ", id)
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
// @Param        student body models.Student true "CreateStudent"
// @Success      200  {struct}  models.Student
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
// @Param        student body models.Student true "UpdateStudent"
// @Success      200  {struct}  models.Student
// @Failure      404  {object}  models.Student
// @Router       /student/{id} [put]
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
// @Param  	     id path int true "id"
// @Success      200  {struct}  models.Student
// @Failure      404  {object}  models.Student
// @Router       /student/{id} [delete]
func DeleteStudent(c *gin.Context) {
	id := c.Params.ByName("id")
	var student models.Student
	d := db.Where("id = ?", id).Delete(&student)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
