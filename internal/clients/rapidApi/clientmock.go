package rapidApi

import (
	"context"
	"net/url"
)

type CtxUrlKey struct {
}

type RapidApiMock struct {
	mocks    map[string][]byte
	QueryKey string
}

func (m *RapidApiMock) SaveResponse(key string, body []byte) {
	if m.mocks == nil {
		m.mocks = make(map[string][]byte)
	}
	m.mocks[key] = body
}

func (m *RapidApiMock) DoRequest(ctx context.Context, host string, method string, query url.Values) ([]byte, error) {
	key := ctx.Value(CtxUrlKey{}).(string)
	return m.mocks[key], nil
}
