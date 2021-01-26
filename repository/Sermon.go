package repository

import (
	"fmt"
	"gintest/dto"
	"gintest/helper"
	"log"
)

func GetAllSermon(request dto.Sermon)(*dto.ListSermon,error){
	dbInstance := helper.GetDBInstance()
	dbInstance.LogMode(true)
	query := fmt.Sprintf("SELECT category_id, category_name,category_desc, created_on, created_by, status, app_id from category where app_id = %v ", request.AppID)
	rows, err := dbInstance.Raw(query).Rows()
	if err != nil {
		log.Print(err)
		return nil,err
	}
	defer rows.Close()

	sermons := make([]dto.Sermon, 0)
	for rows.Next() {
		sermon := new(dto.Sermon)
		err := rows.Scan(&sermon.SermonID,&sermon.SermonName, &sermon.SermonDesc, &sermon.CreatedOn, &sermon.CreatedBy, &sermon.Status,
			&sermon.AppID)
		if err != nil {
			log.Print(err)
		}
		sermons = append(sermons, *sermon)
	}
	if err = rows.Err(); err != nil {
		log.Print(err)
		return nil,err
	}
	result := dto.ListSermon{
		Sermon:sermons,
	}
	return &result,nil
}