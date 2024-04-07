package vitals

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hanshal101/term-test-monitor/database/model"
	"github.com/hanshal101/term-test-monitor/database/postgres"
)

func GetTeachers(c *gin.Context) {
	teacherType := c.Param("type")
	var data []model.Main_Teachers

	tx := postgres.DB.Begin()
	switch teacherType {
	case "teachingStaff":
		if err := tx.Find(&data).Error; err != nil {
			tx.Rollback()
			c.JSON(500, gin.H{"error": "Failed to fetch teaching staff"})
			return
		}
	case "nonteachingStaff":
		var coTeachers []model.Co_Teachers
		if err := tx.Find(&coTeachers).Error; err != nil {
			tx.Rollback()
			c.JSON(500, gin.H{"error": "Failed to fetch non-teaching staff"})
			return
		}
	default:
		c.JSON(400, gin.H{"error": "Invalid teacher type"})
		return
	}

	tx.Commit()
	c.JSON(200, data)
}

type reqTeacher struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phno"`
	Type  string `json:"type"`
}

func CreateTeacher(c *gin.Context) {
	var req reqTeacher
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error in Binding"})
		return
	}
	teacher_type := req.Type
	tx := postgres.DB.Begin()
	switch teacher_type {
	case "TeachingStaff":
		var data model.Main_Teachers
		data.Name = req.Name
		data.Email = req.Email
		data.Phone = req.Phone
		if err := tx.Create(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "error in creating teacher"})
			return
		}
	case "NonTeachingStaff":
		var data model.Co_Teachers
		data.Name = req.Name
		data.Email = req.Email
		data.Phone = req.Phone
		if err := tx.Create(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "error in creating teacher"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"error": "teacher created successfully"})
	}
}
func EditTeacher(c *gin.Context) {
	var req []reqTeacher
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error in Binding"})
		return
	}
	teacherType := c.Param("type")
	tx := postgres.DB.Begin()
	switch teacherType {
	case "teachingStaff":
		for _, r := range req {
			var data model.Main_Teachers
			if err := tx.First(&data, "email = ?", r.Email).Error; err != nil {
				tx.Rollback()
				c.JSON(http.StatusBadRequest, gin.H{"error": "teacher not found"})
				return
			}
			data.Name = r.Name
			data.Phone = r.Phone
			if err := tx.Save(&data).Error; err != nil {
				tx.Rollback()
				c.JSON(http.StatusBadRequest, gin.H{"error": "error in updating teacher"})
				return
			}
		}
	case "nonteachingStaff":
		for _, r := range req {
			var data model.Co_Teachers
			if err := tx.First(&data, "email = ?", r.Email).Error; err != nil {
				tx.Rollback()
				c.JSON(http.StatusBadRequest, gin.H{"error": "teacher not found"})
				return
			}
			data.Name = r.Name
			data.Phone = r.Phone
			if err := tx.Save(&data).Error; err != nil {
				tx.Rollback()
				c.JSON(http.StatusBadRequest, gin.H{"error": "error in updating teacher"})
				return
			}
		}
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid teacher type"})
		return
	}

	tx.Commit()
	c.JSON(http.StatusOK, gin.H{"message": "teacher updated successfully"})
}

func DeleteTeacher(c *gin.Context) {
	email := c.Param("email")
	teacherType := c.Param("type")

	tx := postgres.DB.Begin()

	switch teacherType {
	case "teachingStaff":
		var data model.Main_Teachers
		if err := tx.Where("email = ?", email).Delete(&data).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusBadRequest, gin.H{"error": "teacher not found or error in deletion"})
			return
		}
	case "nonteachingStaff":
		var data model.Co_Teachers
		if err := tx.Where("email = ?", email).Delete(&data).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusBadRequest, gin.H{"error": "teacher not found or error in deletion"})
			return
		}
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid teacher type"})
		return
	}

	tx.Commit()
	c.JSON(http.StatusOK, gin.H{"message": "teacher deleted successfully"})
}
