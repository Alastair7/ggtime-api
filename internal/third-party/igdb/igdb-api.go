package igdb

type Authenticator interface {
	GetAccessToken() (string, error)
}

type Fetcher interface {
	Fetch(url string) (interface{}, error)
}

type IgdbService struct {
	baseUrl string
	auth    Authenticator
}

func NewIgdbService(auth Authenticator) *IgdbService {

	return &IgdbService{
		baseUrl: "http://igdb.es",
		auth:    auth,
	}
}

type IgdbAuthenticator struct{}

func (i *IgdbAuthenticator) GetAccessToken() (string, error) {
	return "access token", nil
}

func (ig *IgdbService) Authenticate() (string, error) {

	return ig.auth.GetAccessToken()
}
