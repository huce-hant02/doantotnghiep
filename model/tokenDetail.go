package model

// TokenDetail details for token authentication
type TokenDetail struct {
	Username     string
	AccessToken  string
	RefreshToken string
	AccessUUID   string
	RefreshUUID  string
	AtExpires    int64
	RtExpires    int64
}
