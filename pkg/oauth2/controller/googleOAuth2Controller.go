package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/avast/retry-go"
	"github.com/kengkeng852/adventure-shop-api/config"
	_adminModel "github.com/kengkeng852/adventure-shop-api/pkg/admin/model"
	"github.com/kengkeng852/adventure-shop-api/pkg/custom"
	_oauth2Exception "github.com/kengkeng852/adventure-shop-api/pkg/oauth2/exception"
	_oauth2Model "github.com/kengkeng852/adventure-shop-api/pkg/oauth2/model"
	_oauth2Service "github.com/kengkeng852/adventure-shop-api/pkg/oauth2/service"
	_playerModel "github.com/kengkeng852/adventure-shop-api/pkg/player/model"
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
	adminGoogleOAuth2  *oauth2.Config
	once               sync.Once

	accessTokenCookieName  = "act"
	refreshTokenCookieName = "rtf"
	stateCookieName        = "state"

	letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func NewGoogleOAuth2Controller(oauth2Service _oauth2Service.OAuth2Service, oauth2Conf *config.OAuth2, logger echo.Logger) OAuth2Controller {
	once.Do(func() {
		setGoogleOAuth2Config(oauth2Conf)
	})

	return &googleOAuth2Controller{oauth2Service, oauth2Conf, logger}
}

func setGoogleOAuth2Config(oauth2Conf *config.OAuth2) {
	playerGoogleOAuth2 = &oauth2.Config{
		ClientID:     oauth2Conf.ClientID,
		ClientSecret: oauth2Conf.ClientSecret,
		RedirectURL:  oauth2Conf.PlayerRedirectUrl,
		Scopes:       oauth2Conf.Scopes,
		Endpoint: oauth2.Endpoint{
			AuthURL:       oauth2Conf.Endpoints.AuthUrl,
			TokenURL:      oauth2Conf.Endpoints.TokenUrl,
			DeviceAuthURL: oauth2Conf.Endpoints.DeviceAuthUrl,
			AuthStyle:     oauth2.AuthStyleInParams,
		},
	}

	adminGoogleOAuth2 = &oauth2.Config{
		ClientID:     oauth2Conf.ClientID,
		ClientSecret: oauth2Conf.ClientSecret,
		RedirectURL:  oauth2Conf.AdminRedirectUrl,
		Scopes:       oauth2Conf.Scopes,
		Endpoint: oauth2.Endpoint{
			AuthURL:       oauth2Conf.Endpoints.AuthUrl,
			TokenURL:      oauth2Conf.Endpoints.TokenUrl,
			DeviceAuthURL: oauth2Conf.Endpoints.DeviceAuthUrl,
			AuthStyle:     oauth2.AuthStyleInParams,
		},
	}
}

func (c *googleOAuth2Controller) PlayerLogin(pctx echo.Context) error {
	state := randomState()

	c.setCookie(pctx, stateCookieName, state)

	return pctx.Redirect(http.StatusFound, playerGoogleOAuth2.AuthCodeURL(state))
}

func (c *googleOAuth2Controller) AdminLogin(pctx echo.Context) error {
	state := randomState()

	c.setCookie(pctx, stateCookieName, state)

	return pctx.Redirect(http.StatusFound, adminGoogleOAuth2.AuthCodeURL(state))
}

func (c *googleOAuth2Controller) PlayerLoginCallback(pctx echo.Context) error {
	ctx := context.Background()

	if err := retry.Do(func() error {
		return c.callbackValidating(pctx)
	}, retry.Attempts(3), retry.Delay(3*time.Second)); err != nil {
		c.logger.Errorf("Failed to validate callback: %s", err.Error())
		return custom.CustomError(pctx, http.StatusUnauthorized, err.Error())
	}

	token, err := playerGoogleOAuth2.Exchange(ctx, pctx.QueryParam("code"))
	if err != nil {
		c.logger.Errorf("Failed to exchange token: %s", err.Error())
		return custom.CustomError(pctx, http.StatusUnauthorized, err.Error())
	}

	client := playerGoogleOAuth2.Client(ctx, token)

	userInfo, err := c.getUserInfo(client)
	if err != nil {
		c.logger.Errorf("Failed to get user info: %s", err.Error())
		return custom.CustomError(pctx, http.StatusUnauthorized, err.Error())
	}

	playerCreatingReq := &_playerModel.PlayerCreatingReq{
		ID:     userInfo.ID,
		Email:  userInfo.Email,
		Name:   userInfo.Name,
		Avatar: userInfo.Picture,
	}

	if err := c.oauth2Service.PlayerAccountCreating(playerCreatingReq); err != nil {
		c.logger.Errorf("Failed to create player: %s", err.Error())
		return custom.CustomError(pctx, http.StatusUnauthorized, err.Error())
	}

	c.setSameSiteCookie(pctx, accessTokenCookieName, token.AccessToken)
	c.setSameSiteCookie(pctx, refreshTokenCookieName, token.RefreshToken)

	return pctx.JSON(http.StatusOK, &_oauth2Model.LoginResponse{Message: "Login Success"})

}

func (c *googleOAuth2Controller) AdminLoginCallback(pctx echo.Context) error {
	ctx := context.Background()

	if err := retry.Do(func() error {
		return c.callbackValidating(pctx)
	}, retry.Attempts(3), retry.Delay(3*time.Second)); err != nil {
		c.logger.Errorf("Failed to validate callback: %s", err.Error())
		return custom.CustomError(pctx, http.StatusUnauthorized, err.Error())
	}

	token, err := adminGoogleOAuth2.Exchange(ctx, pctx.QueryParam("code"))
	if err != nil {
		c.logger.Errorf("Failed to exchange token: %s", err.Error())
		return custom.CustomError(pctx, http.StatusUnauthorized, err.Error())
	}

	client := adminGoogleOAuth2.Client(ctx, token)

	userInfo, err := c.getUserInfo(client)
	if err != nil {
		c.logger.Errorf("Failed to get admin info: %s", err.Error())
		return custom.CustomError(pctx, http.StatusUnauthorized, err.Error())
	}

	adminCreatingReq := &_adminModel.AdminCreatingReq{
		ID:     userInfo.ID,
		Email:  userInfo.Email,
		Name:   userInfo.Name,
		Avatar: userInfo.Picture,
	}

	if err := c.oauth2Service.AdminAccountCreating(adminCreatingReq); err != nil {
		c.logger.Errorf("Failed to create admin: %s", err.Error())
		return custom.CustomError(pctx, http.StatusUnauthorized, err.Error())
	}

	c.setSameSiteCookie(pctx, accessTokenCookieName, token.AccessToken)
	c.setSameSiteCookie(pctx, refreshTokenCookieName, token.RefreshToken)

	return pctx.JSON(http.StatusOK, &_oauth2Model.LoginResponse{Message: "Login Success"})
}

func (c *googleOAuth2Controller) Logout(pctx echo.Context) error {
	accessToken, err := pctx.Cookie(accessTokenCookieName)
	if err != nil {
		c.logger.Errorf("Error reading access token: %s", err.Error())
		return custom.CustomError(pctx, http.StatusBadRequest, err.Error())
	}

	if err := c.revokeToken(accessToken.Value); err != nil {
		c.logger.Errorf("Error revoking token: %s", err.Error())
		return custom.CustomError(pctx, http.StatusInternalServerError, err.Error())
	}

	c.destroyCookie(pctx, accessTokenCookieName)
	c.destroyCookie(pctx, refreshTokenCookieName)

	return pctx.JSON(http.StatusOK, &_oauth2Model.LogoutResponse{Message: "Logout sucessful"})
}

func (c *googleOAuth2Controller) revokeToken(accessToken string) error {
	revokeURL := fmt.Sprintf("%s?token=%s", c.oauth2Conf.RevokeUrl, accessToken)

	response, err := http.Post(revokeURL, "application/x-www-form-urlencoded", nil)
	if err != nil {
		fmt.Println("Error revoking token:", err)
		return err
	}

	defer response.Body.Close()

	return nil
}

func (c *googleOAuth2Controller) setCookie(pctx echo.Context, name, value string) {
	cookie := &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		HttpOnly: true,
	}

	pctx.SetCookie(cookie)
}

func (c *googleOAuth2Controller) destroyCookie(pctx echo.Context, name string) {
	cookie := &http.Cookie{
		Name:     name,
		Path:     "/",
		HttpOnly: true,
		MaxAge:   -1, //delete cookie
	}

	pctx.SetCookie(cookie)
}

func (c *googleOAuth2Controller) setSameSiteCookie(pctx echo.Context, name, value string) {
	cookie := &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}

	pctx.SetCookie(cookie)
}

func (c *googleOAuth2Controller) destroySameSiteCookie(pctx echo.Context, name string) {
	cookie := &http.Cookie{
		Name:     name,
		Path:     "/",
		HttpOnly: true,
		MaxAge:   -1, //delete cookie
		SameSite: http.SameSiteStrictMode,
		// Secure: true, for production
	}

	pctx.SetCookie(cookie)
}

func (c *googleOAuth2Controller) getUserInfo(client *http.Client) (*_oauth2Model.UserInfo, error) {
	response, err := client.Get(c.oauth2Conf.UserInfoUrl)
	if err != nil {
		c.logger.Errorf("Error getting user info: %s", err.Error())
		return nil, err
	}

	defer response.Body.Close()

	userInfoInBytes, err := io.ReadAll(response.Body)
	if err != nil {
		c.logger.Errorf("Error getting reading user info: %s", err.Error())
		return nil, err
	}

	userInfo := new(_oauth2Model.UserInfo)
	if err := json.Unmarshal(userInfoInBytes, &userInfo); err != nil {
		c.logger.Errorf("Error unmarshalling user info: %s", err.Error())
		return nil, err
	}

	return userInfo, nil
}

func (c *googleOAuth2Controller) callbackValidating(pctx echo.Context) error {
	state := pctx.QueryParam("state")

	stateFromCookie, err := pctx.Cookie(stateCookieName)
	if err != nil {
		c.logger.Errorf("Failed to get state from cookie: %s", err.Error())
		return &_oauth2Exception.Unauthorized{}
	}
	if state != stateFromCookie.Value {
		c.logger.Errorf("Invalid state %s", state)
		return &_oauth2Exception.InvalidState{}
	}

	c.destroyCookie(pctx, stateCookieName)

	return nil
}

func randomState() string {
	b := make([]byte, 16)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
