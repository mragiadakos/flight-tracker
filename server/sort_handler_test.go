package server

import (
	"encoding/json"
	"flight-tracker/domain"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestSortFailures(t *testing.T) {
	sorter := domain.Domain{}
	h := NewServer(sorter)
	e := echo.New()
	failures := map[int]string{
		1: `{}`,
		2: `{"paths":[[]]}`,
		3: `{"paths":[["ATL", "EWR"], ["SFO", "SFO"]]}`,
		4: `{"paths":[["EWR", "IND"], ["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]]}`,
		5: `{"paths":[["ATL", "EWR"], ["SFO", "IND", "ATL"]]}`,
		6: `{"paths":[["ATL", "EWR"], ["SFO"]]}`,
		7: `{"paths":[["ATL", "EWR"], [}`,
	}
	for i, v := range failures {
		req := httptest.NewRequest(http.MethodGet, "/sort", strings.NewReader(v))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, h.sortHandler(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code, fmt.Sprint("failed on ", i))
			resp := ResponseSort{}
			err := json.Unmarshal(rec.Body.Bytes(), &resp)
			assert.NoError(t, err, fmt.Sprint("failed on ", i))
			assert.Empty(t, resp.Path, fmt.Sprint("failed on ", i))
			assert.NotNil(t, resp.Message, fmt.Sprint("failed on ", i))
			assert.EqualValues(t, resp.Status, StatusFailure, fmt.Sprint("failed on ", i))
			assert.Greater(t, len(*resp.Message), 0, fmt.Sprint("failed on ", i))
		}
	}
}

func TestSortSuccesses(t *testing.T) {
	sorter := domain.Domain{}
	h := NewServer(sorter)
	e := echo.New()
	successes := map[int]string{
		1: `{"paths":[["SFO", "EWR"]]}`,
		2: `{"paths":[["ATL", "EWR"], ["SFO", "ATL"]]}`,
		3: `{"paths":[["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]]}`,
	}
	expected := []string{"SFO", "EWR"}
	for i, v := range successes {
		req := httptest.NewRequest(http.MethodGet, "/sort", strings.NewReader(v))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, h.sortHandler(c)) {
			assert.Equal(t, http.StatusOK, rec.Code, fmt.Sprint("failed on ", i))
			resp := ResponseSort{}
			err := json.Unmarshal(rec.Body.Bytes(), &resp)
			assert.NoError(t, err, fmt.Sprint("failed on ", i))
			assert.EqualValues(t, resp.Status, StatusSuccess, fmt.Sprint("failed on ", i))
			assert.Nil(t, resp.Message, fmt.Sprint("failed on ", i))
			assert.EqualValues(t, resp.Path, expected)
		}
	}
}
