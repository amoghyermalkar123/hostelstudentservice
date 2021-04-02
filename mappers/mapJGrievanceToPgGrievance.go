package mappers

import (
	"hserver/jtypes"
	"hserver/types"
	"time"
)

func MapJGrievanceToPgGrievance(d *jtypes.Grievance) *types.StudentGriev {
	return &types.StudentGriev{
		StudentID: d.StudentID,
		HostelID:  d.HostelID,
		Descr:     d.Description,
		PostedAt:  time.Now(),
	}
}
