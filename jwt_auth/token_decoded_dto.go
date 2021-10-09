package jwt_auth

type TokenDecodedDto struct {
	UserId        int64
	AccessTokenId string
	RefeshTokenId string
}