package dblayer

import (
	"hserver/types"
	"log"
)

func GetStudentDetails() (*types.StudentDetails, error) {
	db, err := getDbSession()

	if err != nil {
		log.Println(err)
		return nil, err
	}
	dets := &types.StudentDetails{}
	err = db.Model(dets).Where("id = ?", int64(1)).Select(dets)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	return dets, nil
}
