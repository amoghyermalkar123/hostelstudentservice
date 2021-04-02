package types

import "time"

type StudentGriev struct {
	tableName struct{}  `pg:"student_grievs,alias:student_grievs"`
	ID        int64     `pg:"id"`
	StudentID int64     `pg:"student_id"`
	HostelID  int64     `pg:"hostel_id"`
	Descr     string    `pg:"descr"`
	PostedAt  time.Time `pg:"posted_at"`
}
