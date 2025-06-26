package clients

type ClientConfiguration struct {
	AuthUrl      string
	ClientId     string
	ClientSecret string
	GrantType    string
}

func NewClientConfiguration(authUrl string, clientId string, clientSecret string, grantType string) ClientConfiguration {
	return ClientConfiguration{
		AuthUrl:      authUrl,
		ClientId:     clientId,
		ClientSecret: clientSecret,
		GrantType:    grantType,
	}
}
