package eth

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"time"
)

type (
	callParams struct {
		ctx     context.Context
		method  string
		params  []interface{}
		result  interface{}
		timeout time.Duration
	}

	callResponse struct {
		Result interface{} `json:"result"`
		Error  *rpcError   `json:"error"`
	}

	rpcError struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data,omitempty"`
	}
)

func (m rpcError) Error() string {
	return fmt.Sprintf("%d: %s (%#v)", m.Code, m.Message, m.Data)
}

func (c *Client) call(params callParams) error {
	var (
		data   []byte
		err    error
		buf    = c.buf.Get().(*bytes.Buffer)
		resp   *http.Response
		result = &callResponse{Result: params.result}
	)

	defer c.buf.Put(buf)

	// Marshal request body:
	if data, err = json.Marshal(map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  params.method,
		"id":      time.Now().UnixNano(),
		"params":  params.params,
	}); err != nil {
		return err
	}

	// Put data to buffer
	if _, err = buf.Write(data); err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(params.ctx, params.timeout)
	defer cancel()

	req := c.req.WithContext(ctx)
	req.Body = ioutil.NopCloser(buf)
	req.ContentLength = int64(buf.Len())

	// Make request
	resp, err = c.cli.Do(req)
	if err != nil {
		return err
	}

	// Close body after request
	defer resp.Body.Close()

	// status code must be between 200 and 300 codes:
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("bad response code: %d", resp.StatusCode)
	}

	// only for debug reasons:
	if c.cfg.Debug {
		data, err = httputil.DumpResponse(resp, true)
		c.log.Debugw("eth request",
			"method", params.method,
			"params", params.params,
			"status", resp.StatusCode,
			"response", string(data),
			"error", err)
	}

	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return err
	}

	if result.Error != nil {
		return result.Error
	}

	return nil
}
