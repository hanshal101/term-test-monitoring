package model

import "gorm.io/gorm"

type Class struct {
	Room     string `json:"room"`
	Capacity int64  `json:"capacity"`
}

type DualAllocationReq struct {
	Class1 Year    `json:"class1"`
	Class2 Year    `json:"class2"`
	Class  []Class `json:"class"`
}
type SingleAllocReq struct {
	Class1 Year    `json:"class1"`
	Class  []Class `json:"class"`
}

type Year struct {
	Name     string `json:"name"`
	Strength int64  `json:"strength"`
}

type AllocationResult struct {
	gorm.Model
	ClassRoom string `json:"classroom"`
	ClassName string `json:"classname"`
	Start     int64  `json:"start"`
	End       int64  `json:"end"`
}

type Students struct {
	Name       string `json:"name"`
	RollNo     int    `json:"roll_no"`
	Email      string `json:"email"`
	Class      string `json:"class"`
	Department string `json:"department"`
}

type StudentsDB struct {
	gorm.Model
	Name       string `json:"name"`
	RollNo     int    `json:"roll_no"`
	Email      string `json:"email"`
	Class      string `json:"class"`
	Department string `json:"department"`
}

type CreateTimeTable struct {
	gorm.Model
	Year       string `json:"year"`
	Sem        string `json:"sem"`
	Subject    string `json:"subject"`
	Date       string `json:"date"`
	Start_Time string `json:"start_time"`
	End_Time   string `json:"end_time"`
}
type Main_Teachers struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phno"`
}

type Teachers struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phno"`
}

type Co_Teachers struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phno"`
}
type Main_TeachersArr struct {
	ReqAll_MT []Main_Teachers `json:"reqAll_mt"`
}
type Co_TeachersArr struct {
	ReqAll_CO []Co_Teachers `json:"reqAll_co"`
}
type TeacherAllocation struct {
	gorm.Model
	Classroom    string `json:"classroom"`
	Date         string `json:"date"`
	Start_Time   string `json:"start_time"`
	End_Time     string `json:"end_time"`
	Main_Teacher string `json:"main_teacher"`
	Co_Teacher   string `json:"co_teacher"`
}
type TeacherAllocationArr struct {
	ReqAll []TeacherAllocation `json:"reqAll"`
}
type AttPerSt struct {
	gorm.Model
	Name              string `json:"name"`
	RollNo            int    `json:"rollNo"`
	IsPresent         bool   `json:"isPresent"`
	AttendanceModelID uint   `json:"attendanceModelID"`
}

type AttendanceModel struct {
	gorm.Model
	Attendence   []AttPerSt `gorm:"foreignKey:AttendanceModelID" json:"attendance"`
	Class_Name   string     `json:"className"`
	Subject      string     `json:"subject"`
	Main_Teacher string     `json:"mainTeacher"`
}
type AttendenceArr struct {
	ReqAll []AttendanceModel `json:"reqAll"`
}

type Subjects struct {
	gorm.Model
	Name string `json:"name"`
	Year string `json:"year"`
	SEM  int    `json:"sem"`
}

type Attendence_Models struct {
	gorm.Model
	M_Teacher  string `json:"m_teacher"`
	C_Teacher  string `json:"c_teacher"`
	Classroom  string `json:"class_room"`
	Date       string `json:"date"`
	Start_Time string `json:"start_time"`
	End_Time   string `json:"end_time"`
	Subject    string `json:"subject"`
	Year       string `json:"year"`
	Name       string `json:"name"`
	RollNo     int    `json:"roll_no"`
	Class      string `json:"class"`
	IsPresent  bool   `json:"is_present"`
	Supplement string `json:"supplement"`
}

type Subject struct {
	gorm.Model
	Year string `json:"year"`
	Sem  int    `json"sem"`
	Name string `json:"name"`
}

// PAPER REQUESTS MODELS
type PaperModel struct {
	gorm.Model
	TeacherName string `json:"teacher_name"`
	Request     bool   `json:"request"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
}

// DQC Members
type DQCMembers struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

type DQCReview struct {
	gorm.Model
	Name        string `json:"name"`
	Email       string `json:"email"`
	Subject     string `json:"subject"`
	Year        string `json:"year"`
	Semester    string `json:"semester"`
	QPLink      string `json:"qplink"`
	APLink      string `json:"aplink"`
	Request     bool   `json:"request"`
	Status      bool   `json:"status"`
	Approver    string `json:"approver"`
	Description string `json:"description"`
}

type EAuthReq struct {
	Email string `json:"email"`
}

type AllocationCount struct {
	gorm.Model
	Type  string `json:"type"`
	Count string `json:"count" gorm:"default:5"`
}
