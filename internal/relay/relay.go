package relay

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/whookdev/cli/internal/config"
)

type RelayResponse struct {
	WSUrl   string `json:"ws_url"`
	RelayId string `json:"id"`
}

type RelayRequest struct {
	ProjectName string `json:"project_name"`
}

type Relay struct {
	cfg    *config.Config
	client *http.Client
}

func NewRelay(cfg *config.Config) *Relay {
	return &Relay{
		cfg: cfg,
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (r *Relay) GetWSUrl() (string, error) {
	pName := generateProjectName()
	jsonData, err := json.Marshal(&RelayRequest{
		ProjectName: pName,
	})
	if err != nil {
		return "", fmt.Errorf("marshalling request: %w", err)
	}

	req, err := http.NewRequest("POST", r.cfg.ConductorUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := r.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("sending request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("reading response body: %w", err)
	}

	var jsonBody RelayResponse
	if err := json.Unmarshal(body, &jsonBody); err != nil {
		return "", fmt.Errorf("unmarshalling json: %w", err)
	}

	if jsonBody.WSUrl == "" {
		return "", fmt.Errorf("empty websocket URL in response")
	}

	return jsonBody.WSUrl, nil
}

func generateProjectName() string {
	adjectives := []string{
		"fiscal", "happy", "brave", "quiet", "bright", "calm", "eager", "faithful", "gentle", "jolly",
		"kind", "lively", "nice", "proud", "silly", "thankful", "witty", "zealous", "bold", "cheerful",
		"delightful", "elegant",
	}
	nouns := []string{
		"penguin", "tiger", "lion", "elephant", "rabbit", "bear", "wolf", "fox", "deer", "giraffe",
		"zebra", "kangaroo", "panda", "koala", "leopard", "cheetah", "raccoon", "squirrel", "moose",
		"otter", "seal", "whale",
	}
	places := []string{
		"mountain", "forest", "river", "ocean", "desert", "valley", "canyon", "island", "plateau",
		"prairie", "savannah", "tundra", "glacier", "volcano", "lagoon", "swamp", "jungle", "reef",
		"cliff", "meadow", "grove", "bay",
	}

	adjective := adjectives[rand.Intn(len(adjectives))]
	noun := nouns[rand.Intn(len(nouns))]
	place := places[rand.Intn(len(places))]

	return strings.Join([]string{adjective, noun, place}, "-")
}
