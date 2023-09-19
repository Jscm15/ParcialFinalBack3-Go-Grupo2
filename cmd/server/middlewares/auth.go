package middlewares

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Auth struct {
	privateKey, publicKey string
}

func NewAuth(publicKey, privateKey string) *Auth {
	return &Auth{
		publicKey:  publicKey,
		privateKey: privateKey,
	}
}

func (a *Auth) AuthHeader(ctx *gin.Context){
	headerPublicKey := ctx.GetHeader("PUBLIC_KEY")
	headerPrivateKey := ctx.GetHeader("PRIVATE_KEY")

	if a.publicKey != headerPublicKey || a.privateKey != headerPrivateKey{
		ctx.AbortWithError(http.StatusUnauthorized, errors.New("unauthorize access"))
	}
}