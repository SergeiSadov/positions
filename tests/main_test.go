package tests

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"testing"

	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttputil"

	"github.com/sergeisadov/positions/internal/config"
	"github.com/sergeisadov/positions/internal/di"
	"github.com/sergeisadov/positions/internal/service/http_api/controllers"
)

var (
	controller *controllers.Controller
	cfg        config.Config
)

func TestMain(m *testing.M) {
	var err error
	cfg, err = config.Load("../config.test.json")
	if err != nil {
		log.Fatal(err)
	}

	testDI, err := di.Register(cfg)
	if err != nil {
		log.Fatal(err)
	}

	_, err = testDI.DB.Exec(`create table positions
(
    keyword  text,
    position integer,
    domain   text,
    url      text,
    volume   integer,
    results  integer,
    cpc      float,
    updated  datetime,
    primary key (domain, url, keyword)
);`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = testDI.DB.Exec(`
INSERT INTO positions (keyword, position, domain, url, volume, results, cpc, updated) VALUES ('t', 4, 'test.org', 'https://test.org/test.html', 37460000, 270000000, 25.76, '1495248847');
INSERT INTO positions (keyword, position, domain, url, volume, results, cpc, updated) VALUES ('test', 3, 'test.org', 'https://test.org/test.html', 37460000, 270000000, 25.76, '1495248847');
INSERT INTO positions (keyword, position, domain, url, volume, results, cpc, updated) VALUES ('qwert', 2, 'qwert.net', 'https://qwert.net/bt.html', 37460000, 270000000, 25.76, '1495248847');
INSERT INTO positions (keyword, position, domain, url, volume, results, cpc, updated) VALUES ('tt', 1, 'test.org', 'https://test.org/test.html', 37460000, 270000000, 25.76, '1495248847');`)
	if err != nil {
		log.Fatal(err)
	}

	defer testDI.DB.Close()

	controller = controllers.New(testDI.Repository, &testDI.Logger)
	code := m.Run()

	os.Exit(code)
}

func serve(handler fasthttp.RequestHandler, req *http.Request) (*http.Response, error) {
	ln := fasthttputil.NewInmemoryListener()
	defer ln.Close()

	go func() {
		err := fasthttp.Serve(ln, handler)
		if err != nil {
			panic(fmt.Errorf("failed to serve: %v", err))
		}
	}()

	client := http.Client{
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				return ln.Dial()
			},
		},
	}

	return client.Do(req)
}
