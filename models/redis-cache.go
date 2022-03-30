package models

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/garyburd/redigo/redis"
)

type Students []Student

var currentStudentId int

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}
func RedisConnect() redis.Conn {
	c, err := redis.Dial("tcp", ":6379")
	HandleError(err)
	return c
}

// Give us some seed data
func init() {
	PostStudent(Student{

		Name:       "Mirkamol",
		Surname:    "Egamberdiev",
		ID:         1810235,
		Course:     4,
		Adress:     "Namangan",
		Department: "Telecomunication",
		Phone:      901557877,
		Email:      "mirkamol@gmail.com",
	})

	PostStudent(Student{
		Name:       "Muhammadyusuf",
		Surname:    "Saliev",
		ID:         1810237,
		Course:     3,
		Adress:     "Fergana",
		Department: "Computer Science",
		Phone:      901557477,
		Email:      "muhammadyusuf@gmail.com",
	})
}
func FindAll() Students {

	var students []Student

	c := RedisConnect()
	defer c.Close()

	keys, err := c.Do("KEYS", "student:*")
	HandleError(err)

	for _, k := range keys.([]interface{}) {

		var student Student

		reply, err := c.Do("GET", k.([]byte))
		HandleError(err)

		if err := json.Unmarshal(reply.([]byte), &student); err != nil {
			panic(err)
		}
		students = append(students, student)
	}
	return students
}

func FindStudent(id int) Student {

	var student Student

	c := RedisConnect()
	defer c.Close()

	reply, err := c.Do("GET", "student:"+strconv.Itoa(id))
	HandleError(err)

	fmt.Println("GET OK")

	if err = json.Unmarshal(reply.([]byte), &student); err != nil {
		panic(err)
	}
	return student
}

func PostStudent(c Student) {

	currentStudentId += 1

	c.ID = currentStudentId

	p := RedisConnect()
	defer p.Close()

	b, err := json.Marshal(c)
	HandleError(err)

	// Save JSON blob to Redis
	reply, err := p.Do("SET", "student:"+strconv.Itoa(c.ID), b)
	HandleError(err)

	fmt.Println("GET ", reply)
}
