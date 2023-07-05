package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/url"
	"os"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
	"github.com/yitsushi/go-misskey/services/notes"
)

type Options struct {
	token    string
	url      string
	text     string
	postOpts *PostOptions
}

type PostOptions struct {
	visibility string
}

type Output struct {
	NoteID    string    `json:"note_id"`
	Note      string    `json:"note"`
	CreatedAt time.Time `json:"created_at"`
}

func main() {
	// flag
	var (
		optUrl   = flag.String("url", "", "post misskey instance url")
		envUrl   = os.Getenv("MISSKEY_URL")
		optToken = flag.String("token", "", "post misskey token")
		envToken = os.Getenv("MISSKEY_TOKEN")

		optText       = flag.String("text", "", "post text")
		optVisibility = flag.String("vis", models.VisibilityPublic, "post visibility")

		errText []string
	)
	flag.Parse()

	opts := &Options{
		text: strings.Replace(*optText, "\\n", "\n", -1),
		postOpts: &PostOptions{
			visibility: *optVisibility,
		},
	}

	// token
	switch {
	case *optUrl != "":
		opts.token = *optToken
	case envToken != "":
		opts.token = envToken
	default:
		errText = append(errText, "error: MISSKEY_TOKEN env or --token option is required")
	}

	// instance url
	switch {
	case *optUrl != "":
		opts.url = *optUrl
	case envUrl != "":
		opts.url = envUrl
	default:
		errText = append(errText, "error: MISSKEY_URL env or --url option is required")
	}

	if 0 < len(errText) {
		log.Fatal(strings.Join(errText, "\n"))
	}

	u, err := url.Parse(opts.url)
	if err != nil {
		log.Fatal(fmt.Sprintf("url parse error: %v", err))
	}

	client, err := misskey.NewClientWithOptions(
		misskey.WithAPIToken(opts.token),
		misskey.WithBaseURL(u.Scheme, u.Host, u.Path),
	)
	if err != nil {
		log.Fatal(fmt.Sprintf("create client error: %v", err))
	}

	output, err := postNotes(client, opts.postOpts, opts.text)
	if err != nil {
		log.Fatal(fmt.Sprintf("notes post error: %v", err))
	}

	o, _ := json.Marshal(output)
	log.Info(string(o))

	os.Exit(0)
}

func postNotes(client *misskey.Client, postOpts *PostOptions, text string) (*Output, error) {
	result, err := client.Notes().Create(
		notes.CreateRequest{
			Text:       core.NewString(text),
			Visibility: models.Visibility(postOpts.visibility),
		},
	)
	if err != nil {
		return nil, err
	}
	return &Output{
		NoteID:    result.CreatedNote.ID,
		Note:      result.CreatedNote.Text,
		CreatedAt: result.CreatedNote.CreatedAt,
	}, nil
}
