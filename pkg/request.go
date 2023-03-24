package pkg

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
)

func doPostRequest(ctx context.Context, url string, data []byte) ([]byte, error) {
	var req *http.Request
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	var resp *http.Response
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	var b []byte
	b, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("post request failed with code = %d, content = %s", resp.StatusCode, b)
	}
	return b, nil
}

func Post(ctx context.Context, url string, data []byte) ([]byte, error) {
	return doPostRequest(ctx, url, data)
}
