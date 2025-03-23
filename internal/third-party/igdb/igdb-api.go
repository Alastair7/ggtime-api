package igdb

type IgdbService struct {
	baseUrl string
	auth    Authenticator
}

func NewIgdbService(auth Authenticator) *IgdbService {

	return &IgdbService{
		baseUrl: "https://api.igdb.com/v4",
		auth:    auth,
	}
}

func (ig *IgdbService) Authenticate() (string, error) {

	return ig.auth.GetAccessToken()
}
