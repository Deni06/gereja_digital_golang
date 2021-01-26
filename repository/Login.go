package repository

import (
	"fmt"
	"gintest/dto"
	"gintest/helper"
	"gintest/util"
	"log"
)

func Login(request dto.Login)(*dto.User,error){
	dbInstance := helper.GetDBInstance()
	dbInstance.LogMode(true)
	query := fmt.Sprintf("SELECT u.user_id, u.user_name, u.app_id, a.app_name, a.subscription_type, a.subscription_end, a.status_design " +
		"from user u LEFT JOIN app a ON u.app_id = a.app_id WHERE u.user_name = '%v' AND u.password =  '%v' AND u.status_registrasi = 1 " +
		"ORDER BY a.subscription_end DESC LIMIT 1 ", request.Username, request.Password)
	row := dbInstance.Raw(query).Row()
	user := new(dto.User)
	err := row.Scan(&user.UserID, &user.Username, &user.AppID, &user.AppName, &user.SubscriptionType, &user.SubscriptionEnd, &user.StatusDesign)
	if err != nil {
		log.Print(err)
		if user.UserID==""{
			return nil, util.UnhandledError{ErrorMessage:"Username Or Password Incorrect"}
		}
		return nil, err
	}
	return user,nil
}

func GetPermission(request dto.Permission)(*dto.ListPermission,error){
	dbInstance := helper.GetDBInstance()
	dbInstance.LogMode(true)
	query := fmt.Sprintf("SELECT up.user_id, up.permission_id, p.permission_name " +
		"from user_permission up LEFT JOIN permission p ON p.permission_id = up.permission_id WHERE up.user_id = '%v'",
		request.UserID)
	rows, err := dbInstance.Raw(query).Rows()
	if err != nil {
		log.Print(err)
		return nil,err
	}
	defer rows.Close()

	permissions := make([]dto.Permission, 0)
	for rows.Next() {
		permission := new(dto.Permission)
		err := rows.Scan(&permission.UserID,&permission.PermissionID, &permission.PermissionName)
		if err != nil {
			log.Print(err)
		}
		permissions = append(permissions, *permission)
	}
	if err = rows.Err(); err != nil {
		log.Print(err)
		return nil,err
	}
	result := dto.ListPermission{
		ListPermission:permissions,
	}
	return &result,nil
}