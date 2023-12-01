package wework

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	netURL "net/url"
)

type Client struct {
	config *ClientConfig
}

type FileUpload struct {
	FileName  string
	Reader    io.Reader
	FieldName string
}

type requestOptions struct {
	json  any
	file  *FileUpload
	query netURL.Values
	token *TokenDescriptor
}

type requestOption func(*requestOptions)

func WithJSONData(data any) func(*requestOptions) {
	return func(opts *requestOptions) {
		opts.json = data
	}
}

func WithFile(data *FileUpload) func(*requestOptions) {
	return func(opts *requestOptions) {
		opts.file = data
	}
}

func WithQuery(query netURL.Values) func(*requestOptions) {
	return func(opts *requestOptions) {
		opts.query = query
	}
}

func WithToken(token *TokenDescriptor) func(*requestOptions) {
	return func(opts *requestOptions) {
		opts.token = token
	}
}

func (c *Client) NewRequest(ctx context.Context, method string, url string, setters ...requestOption) (req *http.Request, err error) {
	opts := requestOptions{
		query: netURL.Values{},
	}
	header := make(http.Header)
	header.Set("User-Agent", "go-wework/1.0")

	for _, setter := range setters {
		setter(&opts)
	}

	var bodyReader io.Reader
	if opts.file != nil {
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		part, err := writer.CreateFormFile(opts.file.FieldName, opts.file.FileName)
		if err != nil {
			return nil, err
		}

		_, err = io.Copy(part, opts.file.Reader)
		if err != nil {
			return nil, err
		}

		if err = writer.Close(); err != nil {
			return nil, err
		}

		bodyReader = body
		header.Set("Content-Type", writer.FormDataContentType())
	} else if opts.json != nil {
		var reqBytes []byte
		reqBytes, err = json.Marshal(opts.json)
		if err != nil {
			return nil, err
		}

		bodyReader = bytes.NewReader(reqBytes)
		header.Set("Content-Type", "application/json; charset=utf-8")
	}

	if opts.token != nil {
		tk, err := c.GetToken(ctx, opts.token)
		if err != nil {
			return nil, err
		}

		opts.query.Add(opts.token.ParamValue, tk)
	}

	if c.config.DebugMode {
		fmt.Printf("[go-wework] [%s]%s, data: %+v\n", method, url, opts.json)
		opts.query.Add("debug", "1")
	}

	if len(opts.query) > 0 {
		url = url + "?" + opts.query.Encode()
	}
	url = c.config.Options.BaseURL + url

	req, err = http.NewRequestWithContext(ctx, method, url, bodyReader)
	if err != nil {
		return nil, err
	}

	req.Header = header

	return req, nil
}

func (c *Client) SendRequest(req *http.Request, v interface{}) error {

	res, err := c.config.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if isFailureStatusCode(res) {
		return fmt.Errorf("http error: status code %d", res.StatusCode)
	}

	err = decodeResponse(res.Body, v)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) Request(ctx context.Context, method string, url string, v interface{}, setters ...requestOption) error {
	req, err := c.NewRequest(ctx, method, url, setters...)
	if err != nil {
		return err
	}

	err = c.SendRequest(req, &v)
	if err != nil {
		return err
	}

	return nil
}

func isFailureStatusCode(resp *http.Response) bool {
	return resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusBadRequest
}

func decodeResponse(body io.Reader, v interface{}) error {
	if v == nil {
		v = &APIBaseResponse{}
	}

	bodyBytes, err := io.ReadAll(body)

	if err != nil {
		return err
	}

	var apiResp APIBaseResponse
	err = json.Unmarshal(bodyBytes, &apiResp)
	if err != nil {
		return err
	}

	if apiResp.ErrCode != 0 {
		return fmt.Errorf("api error: %d - %s", apiResp.ErrCode, apiResp.ErrMsg)
	}

	if err := json.Unmarshal(bodyBytes, v); err != nil {
		return err
	}

	return nil
}

func NewClient(config *ClientConfig) *Client {
	return &Client{config: config}
}
