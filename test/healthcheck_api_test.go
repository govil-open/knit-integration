package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestHelathCheckAPI(t *testing.T) {
	testCases := []struct {
		name          string
		setupHeader   func(t *testing.T, request *http.Request)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			setupHeader: func(t *testing.T, request *http.Request) {
				request.Header.Set("X-Health-Check", "1")
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name: "HealthCheckNotFound",
			setupHeader: func(t *testing.T, request *http.Request) {
				request.Header.Set("X-Health-Check", "100")
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
	}
	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			recorder := httptest.NewRecorder()
			url := "/health"
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)
			tc.setupHeader(t, request)
			testServer.Router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}
