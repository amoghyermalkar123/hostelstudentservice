package dblayer

import (
	"hserver/types"
	"log"
)

func CreateStudentGrievance(grievance *types.StudentGriev) error {
	db, err := getDbSession()

	if err != nil {
		log.Println(err)
		return err
	}

	err = db.Insert(grievance)

	if err != nil {
		log.Println("jbubukbb")
		log.Println(err)
		return err
	}
	return nil
}

func GetStudentGrievances(studentID int64) ([]*types.StudentGriev, error) {
	db, err := getDbSession()

	if err != nil {
		log.Println(err)
		return nil, err
	}
	dets := make([]*types.StudentGriev, 0)
	err = db.Model(dets).Where("student_id = ?", studentID).Select(dets)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	return dets, nil
}
