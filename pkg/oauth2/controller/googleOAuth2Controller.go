package service

import (
	"math/rand"
	"sync"

	"github.com/kengkeng852/adventure-shop-api/config"
	_oauth2Service "github.com/kengkeng852/adventure-shop-api/pkg/oauth2/service"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
)

type googleOAuth2Controller struct {
	oauth2Service _oauth2Service.OAuth2Service
	oauth2Conf    *config.OAuth2
	logger        echo.Logger
}

var (
	playerGoogleOAuth2 *oauth2.Config
	adminGoogleOAUth2 *oauth2.Config
	once 	sync.Once

	accessTokenCookieName = "act"
	refreshTokenCookieName = "rtf"
	stateCookieName = "state"

	letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func NewGoogleOAuth2Controller(oauth2Service _oauth2Service.OAuth2Service, oauth2Conf *config.OAuth2, logger echo.Logger) OAuth2Controller {
	once.Do(func() {
		setGoogleOAuth2Config(oauth2Conf)
	})

	return &googleOAuth2Controller{oauth2Service, oauth2Conf, logger}
}

func setGoogleOAuth2Config(oauth2Conf *config.OAuth2) {
	playerGoogleOAuth2 = &oauth2.Config {
		ClientID: oauth2Conf.ClientID,
		ClientSecret: oauth2Conf.ClientSecret,
		RedirectURL: oauth2Conf.PlayerRedirectUrl,
		Scopes: oauth2Conf.Scopes,
		Endpoint: oauth2.Endpoint{
			AuthURL: oauth2Conf.Endpoints.AuthUrl,
			TokenURL: oauth2Conf.Endpoints.TokenUrl,
			DeviceAuthURL: oauth2Conf.Endpoints.DeviceAuthUrl,
			AuthStyle: oauth2.AuthStyleInParams,
		},

	}
}

func RandomStringBtyes() string {
	b := make([]byte, 16)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
