//go:build golden
// +build golden

package replayer

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-replayers/httpreplay"

	"github.com/go-replayer-article/user"
)

func getClient(t *testing.T, recorderOpts ...func(*httpreplay.Recorder)) user.Client {
	var cl *http.Client

	filename := fmt.Sprintf("testdata/%s", t.Name())

	rec, err := httpreplay.NewRecorderWithOpts(filename)
	if err != nil {
		t.Fatalf("could not create recorder: %s", err)
	}

	for i := range recorderOpts {
		recorderOpts[i](rec)
	}

	t.Cleanup(func() {
		if err := rec.Close(); err != nil {
			t.Errorf("could not close recorder: %s", err)
		}
	})

	cl = rec.Client()

	const url = "http://localhost:8080"
	return user.Client{BaseClient: cl, BaseURL: url}
}
