package defs

type Token struct {
	StatusCode string `json:"StatusCode"`
	AccessKeyId string `json:"AccessKeyId"`
	AccessKeySecret string `json:"AccessKeySecret"`
	SecurityToken string `json:"SecurityToken"`
	Expiration string `json:"Expiration"`
}