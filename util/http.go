package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"

	"knit-integration/config"
	"knit-integration/logger"
	"knit-integration/model"

	"github.com/gin-gonic/gin"
)

func Post(ctx *gin.Context, conf *config.AppConfig, logwrapper *logger.Logger, contentType string, body map[string]any) (map[string]any, error) {
	postBody, _ := json.Marshal(body)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, conf.AuthURL, bytes.NewBuffer(postBody))
	baseReq := &model.BaseRequest{CompanyID: 0, AccountID: 0, UserID: 0}
	req.Header.Set("Content-Type", contentType)
	if err != nil {
		logwrapper.Error(baseReq, err)
		return nil, err
	}
	client := http.Client{Timeout: time.Duration(conf.TokenSrvTimeout) * time.Millisecond}
	resp, err := client.Do(req)
	if err != nil {
		logwrapper.Error(baseReq, err)
		return nil, err
	}
	defer resp.Body.Close()
	var responseBody []byte
	responseBody, err = io.ReadAll(resp.Body)
	if err != nil {
		logwrapper.Error(baseReq, err)
		return nil, err
	}
	var responseMap map[string]any
	if err = json.Unmarshal([]byte(string(responseBody)), &responseMap); err != nil {
		return nil, err
	}
	if resp.StatusCode != config.HTTPSuccessCode {
		err = errors.New("invalid response")
	}
	return responseMap, err
}
