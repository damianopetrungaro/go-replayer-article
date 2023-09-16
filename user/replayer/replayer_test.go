package replayer

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/google/go-replayers/httpreplay"

	"github.com/go-replayer-article/user"
)

func TestClient_Save(t *testing.T) {
	t.Run("saved", func(t *testing.T) {
		cl := getClient(t,
			func(rec *httpreplay.Recorder) { rec.ScrubBody(`"CreatedAt":\s*".*?"`) },
			func(rec *httpreplay.Recorder) { rec.RemoveRequestHeaders("Content-Length") },
		)

		u := &user.User{
			ID:        "an-existing-id",
			Name:      "Luke",
			CreatedAt: time.Now(),
		}

		ctx := context.Background()

		if err := cl.Save(ctx, u); err != nil {
			t.Errorf("could not save existing user: %s", err)
		}
	})
}

func TestClient_Get(t *testing.T) {
	t.Run("found", func(t *testing.T) {
		cl := getClient(t)

		id := "an-existing-id"
		name := "Luke"

		ctx := context.Background()
		u, err := cl.Get(ctx, id)
		if err != nil {
			t.Errorf("could not get existing user: %s", err)
		}

		if u.ID != id {
			t.Errorf("could not match existing user id: %s", u.ID)
		}

		if u.Name != name {
			t.Errorf("could not match existing user name: %s", u.Name)
		}
	})

	t.Run("not found", func(t *testing.T) {
		cl := getClient(t)

		id := "a-non-existing-id"

		ctx := context.Background()
		u, err := cl.Get(ctx, id)
		if !errors.Is(err, user.ErrNotFound) {
			t.Errorf("could not match non existing user error: %s", err)
		}

		if u != nil {
			t.Errorf("could get non existing user: %v", u)
		}
	})
}
