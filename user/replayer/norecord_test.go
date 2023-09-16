//go:build !golden
// +build !golden

package replayer

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-replayers/httpreplay"

	"github.com/go-replayer-article/user"
)

func getClient(t *testing.T, _ ...func(*httpreplay.Recorder)) user.Client {
	var cl *http.Client

	filename := fmt.Sprintf("testdata/%s", t.Name())

	rep, err := httpreplay.NewReplayer(filename)
	if err != nil {
		t.Fatalf("could not create recorder: %s", err)
	}

	t.Cleanup(func() {
		if err := rep.Close(); err != nil {
			t.Errorf("could not close recorder: %s", err)
		}
	})

	cl = rep.Client()

	const url = "http://localhost:8080"
	return user.Client{BaseClient: cl, BaseURL: url}
}
