package models

//TokenDetails Token details
type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUUID   string
	RefreshUUID  string
	AtExpires    int64
	RtExpires    int64
}

//AccessDetails Access details
type AccessDetails struct {
	AccessUUID string
	UserID     string
}
