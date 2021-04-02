package jtypes

type Grievance struct {
	StudentID   int64  `json:"student_id"`
	HostelID    int64  `json:"hostel_id"`
	Description string `json:"descr"`
}
