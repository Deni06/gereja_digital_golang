package dto

type User struct {
	UserID      string
	Username    string
	AppID   string
	AppName         string
	SubscriptionType   string
	SubscriptionEnd  string
	StatusRegistrasi string
	StatusDesign string
	Permission ListPermission
}

type ListPermission struct {
	ListPermission []Permission `json:"ListPermission"`
}

type Permission struct {
	PermissionID      string `json:"PermissionID"`
	PermissionName    string `json:"PermissionName"`
	UserID string `json:"UserID"`
}

type Login struct {
	Username    string
	Password   string
}

type ResponseToken struct {
	AppID string `json:"app_Id"`
	AppName string `json:"app_name"`
	Authorized bool `json:"authorized"`
	Exp int32 `json:"exp"`
	Permission ListPermission `json:"permission"`
	StatusDesign string `json:"status_Design"`
	SubscriptionType string `json:"subscription_type"`
	UserID string `json:"user_id"`
	Username string `json:"username"`
}