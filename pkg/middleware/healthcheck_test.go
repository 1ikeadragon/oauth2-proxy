package middleware

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func upstreamHandler() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(rw, "upstream")
	})
}

func TestHealthCheck(t *testing.T) {
	testCases := []struct {
		name         string
		path         string
		userAgent    string
		expectedCode int
		expectedBody string
	}{
		{
			name:         "configured ping path is served",
			path:         "/ping",
			expectedCode: http.StatusOK,
			expectedBody: "OK",
		},
		{
			name:         "configured user agent is served",
			path:         "/ready",
			userAgent:    "Health-Checker/1.0",
			expectedCode: http.StatusOK,
			expectedBody: "OK",
		},
		{
			name:         "configured user agent is served on any path",
			path:         "/private/resource",
			userAgent:    "Health-Checker/1.0",
			expectedCode: http.StatusOK,
			expectedBody: "OK",
		},
		{
			name:         "unrelated request reaches the next handler",
			path:         "/private/resource",
			userAgent:    "Mozilla/5.0",
			expectedCode: http.StatusForbidden,
			expectedBody: "upstream",
		},
	}

	handler := NewHealthCheck([]string{"/ping"}, []string{"Health-Checker/1.0"})(upstreamHandler())

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, testCase.path, nil)
			if testCase.userAgent != "" {
				request.Header.Set(userAgentHeader, testCase.userAgent)
			}

			recorder := httptest.NewRecorder()
			handler.ServeHTTP(recorder, request)

			if recorder.Code != testCase.expectedCode {
				t.Errorf("expected status %d, got %d", testCase.expectedCode, recorder.Code)
			}
			if recorder.Body.String() != testCase.expectedBody {
				t.Errorf("expected body %q, got %q", testCase.expectedBody, recorder.Body.String())
			}
		})
	}
}
