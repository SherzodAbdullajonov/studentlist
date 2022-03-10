package cache

import "studentList/models"
type StudentCache interface{
	Set(key string, value *models.Student)
	Get(key string) *models.Student
}
