package controller

import (
	"context"
	"net/http"

	"github.com/kengkeng852/adventure-shop-api/pkg/custom"
	_oauth2Exception "github.com/kengkeng852/adventure-shop-api/pkg/oauth2/exception"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
)

func (c *googleOAuth2Controller) PlayerAuthorizing(pctx echo.Context, next echo.HandlerFunc) error {
	ctx := context.Background()

	tokenSource, err := c.getTokenSource(pctx)
	if err != nil {
		return custom.CustomError(pctx, http.StatusUnauthorized, err.Error())
	}

	if !tokenSource.Valid() {
		tokenSource, err = c.playerTokenRefreshing(pctx, tokenSource)
		if err != nil {
			return custom.CustomError(pctx, http.StatusUnauthorized, err.Error())
		}
	}

	client := playerGoogleOAuth2.Client(ctx, tokenSource)

	userInfo, err := c.getUserInfo(client)
	if err !=  nil {
		return custom.CustomError(pctx, http.StatusUnauthorized, err.Error())
	}

	if !c.oauth2Service.IsPlayerValid(userInfo.ID) {
		return custom.CustomError(pctx, http.StatusUnauthorized, err.Error())
	}

	pctx.Set("playerID", userInfo.ID)

	return next(pctx)

}

func (c *googleOAuth2Controller) AdminAuthorizing(pctx echo.Context, next echo.HandlerFunc) error {
	ctx := context.Background()

	tokenSource, err := c.getTokenSource(pctx)
	if err != nil {
		return custom.CustomError(pctx, http.StatusUnauthorized, err.Error())
	}

	if !tokenSource.Valid() {
		tokenSource, err = c.adminTokenRefreshing(pctx, tokenSource)
		if err != nil {
			return custom.CustomError(pctx, http.StatusUnauthorized, err.Error())
		}
	}

	client := adminGoogleOAuth2.Client(ctx, tokenSource)

	userInfo, err := c.getUserInfo(client)
	if err !=  nil {
		return custom.CustomError(pctx, http.StatusUnauthorized, err.Error())
	}

	if !c.oauth2Service.IsAdminValid(userInfo.ID) {
		return custom.CustomError(pctx, http.StatusUnauthorized, err.Error())
	}

	pctx.Set("adminID", userInfo.ID)

	return next(pctx)
}

func (c *googleOAuth2Controller) playerTokenRefreshing(pctx echo.Context, token *oauth2.Token) (*oauth2.Token, error) {
	ctx := context.Background()

	updatedToken, err := playerGoogleOAuth2.TokenSource(ctx, token).Token()
	if err != nil {
		return nil, &_oauth2Exception.Unauthorized{}
	}

	c.setSameSiteCookie(pctx, accessTokenCookieName, updatedToken.AccessToken)
	c.setSameSiteCookie(pctx, refreshTokenCookieName, updatedToken.RefreshToken)

	return updatedToken, nil
}

func (c *googleOAuth2Controller) adminTokenRefreshing(pctx echo.Context, token *oauth2.Token) (*oauth2.Token, error) {
	ctx := context.Background()

	updatedToken, err := adminGoogleOAuth2.TokenSource(ctx, token).Token()
	if err != nil {
		return nil, &_oauth2Exception.Unauthorized{}
	}

	c.setSameSiteCookie(pctx, accessTokenCookieName, updatedToken.AccessToken)
	c.setSameSiteCookie(pctx, refreshTokenCookieName, updatedToken.RefreshToken)

	return updatedToken, nil
}

func (c *googleOAuth2Controller) getTokenSource(pctx echo.Context) (*oauth2.Token, error) {
	accessToken, err := pctx.Cookie(accessTokenCookieName)
	if err != nil {
		return nil, &_oauth2Exception.Unauthorized{}
	}

	refreshToken, err := pctx.Cookie(refreshTokenCookieName)
	if err != nil {
		return nil, &_oauth2Exception.Unauthorized{}
	}

	return &oauth2.Token{
		AccessToken:  accessToken.Value,
		RefreshToken: refreshToken.Value,
	}, nil
}
