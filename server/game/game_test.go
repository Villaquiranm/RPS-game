package game

import (
	"RPS_gaming/game/logic"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPlayerVSComputer(t *testing.T) {
	tests := map[string]struct {
		option         string
		expectedStatus int
	}{
		"invalid option": {
			option:         "INVALID",
			expectedStatus: http.StatusBadRequest,
		},
		"ok": {
			option:         logic.MovesNames[logic.LIZARD],
			expectedStatus: http.StatusOK,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			game := NewGame()
			req, err := http.NewRequest("GET", "/play", nil)
			if err != nil {
				t.Fatal(err)
			}
			q := req.URL.Query()
			q.Add("option", tc.option)
			req.URL.RawQuery = q.Encode()
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(game.PlayerVsComputer)
			handler.ServeHTTP(rr, req)

			// Check the status code is what we expect.
			if status := rr.Code; status != tc.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tc.expectedStatus)
			}
		})
	}
}

func TestComputerVSComputer(t *testing.T) {
	tests := map[string]struct {
		expectedStatus int
	}{
		"ok": {
			expectedStatus: http.StatusOK,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			game := NewGame()
			req, err := http.NewRequest("GET", "/play", nil)
			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(game.ComputerVsComputer)
			handler.ServeHTTP(rr, req)

			// Check the status code is what we expect.
			if status := rr.Code; status != tc.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tc.expectedStatus)
			}
		})
	}
}
