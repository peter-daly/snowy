// +build documentation

package documents

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/SimonRichardson/betwixt"
	"github.com/SimonRichardson/betwixt/pkg/output"
	"github.com/go-kit/kit/log"
	"github.com/golang/mock/gomock"
	"github.com/trussle/snowy/pkg/document"
	metricMocks "github.com/trussle/snowy/pkg/metrics/mocks"
	"github.com/trussle/snowy/pkg/repository"
	repoMocks "github.com/trussle/snowy/pkg/repository/mocks"
	"github.com/trussle/snowy/pkg/uuid"
)

func TestDocumentation_Flow(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	file, err := os.Create("README.md")
	if err != nil {
		t.Fatal(err)
	}

	var (
		clients  = metricMocks.NewMockGauge(ctrl)
		duration = metricMocks.NewMockHistogramVec(ctrl)
		observer = metricMocks.NewMockObserver(ctrl)
		repo     = repoMocks.NewMockRepository(ctrl)

		api = NewAPI(repo, log.NewNopLogger(), clients, duration)

		outputs = []betwixt.Output{
			output.NewMarkdown(file, output.Options{
				Header:    "# Snowy",
				Optionals: true,
			}),
		}
		capture = betwixt.New(api, outputs)
		server  = httptest.NewServer(capture)

		uid         = uuid.New()
		tags        = []string{"abc", "def", "g"}
		inputDoc, _ = document.Build(
			document.WithAuthorID(uid.String()),
			document.WithName("document-name"),
			document.WithTags(tags),
		)
		outputDoc, _ = document.Build(
			document.WithResourceID(uid),
			document.WithAuthorID(uuid.New().String()),
			document.WithName("document-name"),
			document.WithTags(tags),
			document.WithCreatedOn(time.Now()),
			document.WithDeletedOn(time.Time{}),
		)
	)

	defer func() {
		if err := capture.Output(); err != nil {
			t.Fatal(err)
		}

		file.Sync()
		file.Close()
	}()

	t.Run("get", func(t *testing.T) {
		clients.EXPECT().Inc().Times(1)
		clients.EXPECT().Dec().Times(1)

		duration.EXPECT().WithLabelValues("GET", "/", "200").Return(observer).Times(1)
		observer.EXPECT().Observe(Float64()).Times(1)

		repo.EXPECT().GetDocument(uid, repository.Query{
			Tags: tags,
		}).Times(1).Return(outputDoc, nil)

		resp, err := http.Get(fmt.Sprintf("%s?resource_id=%s&query.tags=%s", server.URL, uid, strings.Join(tags, ",")))
		if err != nil {
			t.Fatal(err)
		}
		defer resp.Body.Close()
	})

	t.Run("get multiple", func(t *testing.T) {
		clients.EXPECT().Inc().Times(1)
		clients.EXPECT().Dec().Times(1)

		duration.EXPECT().WithLabelValues("GET", "/multiple", "200").Return(observer).Times(1)
		observer.EXPECT().Observe(Float64()).Times(1)

		repo.EXPECT().GetDocuments(uid, repository.Query{
			Tags: tags,
		}).Times(1).Return([]document.Document{
			outputDoc,
		}, nil)

		resp, err := http.Get(fmt.Sprintf("%s/multiple?resource_id=%s&query.tags=%s", server.URL, uid, strings.Join(tags, ",")))
		if err != nil {
			t.Fatal(err)
		}
		defer resp.Body.Close()
	})

	t.Run("insert", func(t *testing.T) {
		clients.EXPECT().Inc().Times(1)
		clients.EXPECT().Dec().Times(1)

		duration.EXPECT().WithLabelValues("POST", "/", "200").Return(observer).Times(1)
		observer.EXPECT().Observe(Float64()).Times(1)

		repo.EXPECT().InsertDocument(Document(inputDoc)).Times(1).Return(outputDoc, nil)

		b, err := json.Marshal(struct {
			Name     string   `json:"name"`
			AuthorID string   `json:"author_id"`
			Tags     []string `json:"tags"`
		}{
			Name:     "document-name",
			AuthorID: uid.String(),
			Tags:     []string{"abc", "def", "g"},
		})
		if err != nil {
			t.Fatal(err)
		}

		resp, err := http.Post(server.URL, "application/json", bytes.NewReader(b))
		if err != nil {
			t.Fatal(err)
		}
		defer resp.Body.Close()
	})

	t.Run("append", func(t *testing.T) {
		clients.EXPECT().Inc().Times(1)
		clients.EXPECT().Dec().Times(1)

		duration.EXPECT().WithLabelValues("PUT", "/", "200").Return(observer).Times(1)
		observer.EXPECT().Observe(Float64()).Times(1)

		repo.EXPECT().AppendDocument(uid, Document(inputDoc)).Return(outputDoc, nil).Times(1)

		b, err := json.Marshal(struct {
			Name     string   `json:"name"`
			AuthorID string   `json:"author_id"`
			Tags     []string `json:"tags"`
		}{
			Name:     "document-name",
			AuthorID: uid.String(),
			Tags:     []string{"abc", "def", "g"},
		})
		if err != nil {
			t.Fatal(err)
		}

		resp, err := Put(fmt.Sprintf("%s?resource_id=%s", server.URL, uid), "application/json", bytes.NewReader(b))
		if err != nil {
			t.Fatal(err)
		}
		defer resp.Body.Close()

	})
}
