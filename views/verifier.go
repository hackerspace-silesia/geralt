package views

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/slack-go/slack"
)

func VerifySecret(c *gin.Context, secret string) error {
	verifier, err := slack.NewSecretsVerifier(c.Request.Header, secret)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return err
	}
	buf, _ := ioutil.ReadAll(c.Request.Body)
	buff_verify := bytes.NewBuffer(buf)
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(buf))

	if _, err := verifier.Write(buff_verify.Bytes()); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return err
	}
	if err := verifier.Ensure(); err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return err
	}
	return nil
}
