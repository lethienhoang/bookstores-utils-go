package jwt_auth

type TokenDetailsDto struct {
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
	UserId       int64
}
