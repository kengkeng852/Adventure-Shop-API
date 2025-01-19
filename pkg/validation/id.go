package validation

import (
	_adminException "github.com/kengkeng852/adventure-shop-api/pkg/admin/exception"
	_playerException "github.com/kengkeng852/adventure-shop-api/pkg/player/exception"
	"github.com/labstack/echo/v4"
)

 func AdminIDGetting(pctx echo.Context) (string, error) {
	if adminID, ok := pctx.Get("adminID").(string); !ok || adminID == ""{
		return "", &_adminException.AdminNotFound{AdminID: "Unknown"}
	}else {
		return adminID, nil
	}
 }

 func PlayerIDGetting(pctx echo.Context) (string, error) {
	if playerID, ok := pctx.Get("playerID").(string); !ok || playerID == ""{
		return "", &_playerException.PlayerNotFound{PlayerID: "Unknown"}
	}else {
		return playerID, nil
	}
 }