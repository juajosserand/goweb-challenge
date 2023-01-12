package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"runtime"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/juajosserand/goweb-challenge/internal/domain"
	"github.com/juajosserand/goweb-challenge/internal/ticket"
	"github.com/juajosserand/goweb-challenge/mock"
	"github.com/stretchr/testify/assert"
)

func init() {
	_, filename, _, _ := runtime.Caller(0)
	rootDir := path.Join(path.Dir(filename), "../..")
	_ = os.Chdir(rootDir)

	_ = godotenv.Load(rootDir + "/.env")
}

func arrange(method string, endpoint string, headers map[string]string, body []byte) (func() *http.Response, error) {
	var tickets = []domain.Ticket{
		{
			Id:      "1",
			Name:    "Tait Mc Caughan",
			Email:   "tmc0@scribd.com",
			Country: "Finland",
			Time:    "17:11",
			Price:   785.00,
		},
		{
			Id:      "2",
			Name:    "Padget McKee",
			Email:   "pmckee1@hexun.com",
			Country: "China",
			Time:    "20:19",
			Price:   537.00,
		},
		{
			Id:      "3",
			Name:    "Yalonda Jermyn",
			Email:   "yjermyn2@omniture.com",
			Country: "China",
			Time:    "18:11",
			Price:   579.00,
		},
	}

	repo := mock.NewRepositoryTest(&mock.DbMock{Db: tickets})
	svc := ticket.NewService(repo)
	mux := gin.Default()
	NewService(mux, svc)

	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	for h, v := range headers {
		req.Header.Set(h, v)
	}

	if len(body) > 0 {
		req.Header.Set("Content-Type", "application/json")
	}

	res := httptest.NewRecorder()

	return func() *http.Response {
		mux.ServeHTTP(res, req)
		return res.Result()
	}, nil
}

func TestGetTicketsByCountry(t *testing.T) {
	act, err := arrange(http.MethodGet, "/tickets/getByCountry/China", nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	res := act()
	var data map[string]int
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, 2, data["total_tickets"])
}

func TestAverageDestination(t *testing.T) {
	act, err := arrange(http.MethodGet, "/tickets/getAverage/China", nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	res := act()
	var data map[string]float64
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.InEpsilon(t, float64(2)/float64(3)*100, data["avg"], 0.001)
}
