package model

type TokensByCodeRequestParams struct {
	OxdId string `json:"oxd_id"`
	Code  string `json:"code"`
	State string `json:"state"`
}

type TokensByCodeResponseParams struct {
	AccessToken string `json:"access_token"`
	ExpiresIn int64 `json:"expires_in"`
	IdToken	string `json:"id_token"`
	RefreshToken	string `json:"refreshToken"`

}
