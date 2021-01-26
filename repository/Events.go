package repository

import (
	"fmt"
	"gintest/dto"
	"gintest/helper"
	"log"
)

func GetAllEvents(request dto.Events)(*dto.ListEvents,error){
	dbInstance := helper.GetDBInstance()
	dbInstance.LogMode(true)
	query := fmt.Sprintf("SELECT event_id, event_name,event_start, event_detail, event_location, created_date, created_by," +
		"event_date, event_end, event_img, app_id, status FROM events where app_id = %v ", request.AppID)
	rows, err := dbInstance.Raw(query).Rows()
	if err != nil {
		log.Print(err)
		return nil,err
	}
	defer rows.Close()

	events := make([]dto.Events, 0)
	for rows.Next() {
		event := new(dto.Events)
		err := rows.Scan(&event.EventID,&event.EventName, &event.EventStart, &event.EventDetail, &event.EventLocation, &event.CreatedDate,
			&event.CreatedBy, &event.EventDate, &event.EventEnd, &event.EventImg, &event.AppID, &event.Status)
		if err != nil {
			log.Print(err)
		}
		events = append(events, *event)
	}
	if err = rows.Err(); err != nil {
		log.Print(err)
		return nil,err
	}
	result := dto.ListEvents{
		Events:events,
	}
	return &result,nil
}