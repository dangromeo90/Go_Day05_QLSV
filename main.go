package main

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)
type Student struct{
	ID    int    
    Name  string 
    Age   int    
    Email string 
}
var students = []Student{
    {ID: 1, Name: "Nguyen Vu Nhat Dang", Age: 34, Email: "nhadang0915111916@gmail.com"},
    {ID: 2, Name: "Le Minh Tan", Age: 22, Email: "leminhtan@gmail.com"},
}


func main() {
 r := gin.Default()

  // Lấy Danh sách sinh viên
  r.GET("/students",getListStudent)

 // lấy 1 sinh viên
 r.GET("/student/:id",getOneStudent)

  // Thêm sinh viên
r.POST("/student",addStudent)

// Cập nhật sinh viên
r.PUT("update-student/:id",updateStudent)

// Xoá sinh viên
r.DELETE("/delele-student/:id",deleteStudent)
r.Run() 
}

// Danh sách sinh viên
func getListStudent(c * gin.Context)  {
	c.JSON(http.StatusOK,students)
}
// Lấy 1 sinh viên
func getOneStudent(c * gin.Context){
	idST := c.Param("id")
	id ,err := strconv.Atoi(idST)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error" :"Invalid ID"})
	}
	for _, student := range students{
		if student.ID == id {
			c.JSON(http.StatusOK,student)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error" : "Student not found"})
}

// Thêm sinh viên
func addStudent(c *gin.Context) {
	var newStudent Student
	if err := c.BindJSON(&newStudent); err != nil{
		c.JSON(http.StatusBadRequest , gin.H{"error": "Invalid input"})
		return
	}
	newStudent.ID = len(students) + 1
	students =append(students, newStudent)
	c.JSON(http.StatusOK, newStudent)
}

// Cập nhật sinh viên
func updateStudent( c * gin.Context)  {
	idSt := c.Param("id")
	id, err := strconv.Atoi(idSt)
	if err != nil{
		c.JSON(http.StatusBadRequest , gin.H{"error" : "Invalid ID"})
		return
	}
	var updateStudent Student
	if err := c.BindJSON(&updateStudent) ; err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"error": "Invalid input"})
		return
	}
	for i, student := range students{
		if student.ID == id {
			students[i].Name = updateStudent.Name
			students[i].Age = updateStudent.Age
			students[i].Email = updateStudent.Email
			c.JSON(http.StatusOK, students[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
}

// Xoá Sinh Viên
func deleteStudent( c *gin.Context)  {
	idST := c.Param("id")
	id , err := strconv.Atoi(idST)
	if err != nil {
		c.JSON(http.StatusBadRequest , gin.H{"error" : "Invalid ID"})
		return
	}
	
	for i, student := range students{
		if student.ID == id{
			students = append(students[:i],students[i+1:]... )
			c.JSON(http.StatusOK, gin.H{"message": "Student deleted"})
            return
		}
	}
    c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
}