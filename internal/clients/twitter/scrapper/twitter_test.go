package scrapper

import (
	"net/http"
	"reflect"
	"testing"
)

func TestClient_updateToken(t *testing.T) {
	type fields struct {
		host   string
		token  string
		client http.Client
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "1",
			fields: fields{
				host:   "twitter.com",
				token:  "",
				client: http.Client{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				host:   tt.fields.host,
				token:  tt.fields.token,
				client: tt.fields.client,
			}
			if err := c.updateToken(); (err != nil) != tt.wantErr {
				t.Errorf("updateToken() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_GetTweetPage(t *testing.T) {
	type fields struct {
		host   string
		token  string
		client http.Client
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "1",
			fields: fields{
				host:   "twitter.com",
				token:  "1580014184506572802",
				client: http.Client{},
			},
			args:    args{id: "495719809695621121"},
			want:    []byte{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				host:   tt.fields.host,
				token:  tt.fields.token,
				client: tt.fields.client,
			}
			got, err := c.GetTweetPage(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTweetPage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTweetPage() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetTweet(t *testing.T) {
	type fields struct {
		host   string
		token  string
		client http.Client
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "1",
			fields: fields{
				host:   "twitter.com",
				token:  "1580273926558695424",
				client: http.Client{},
			},
			args:    args{id: "1579898619272257537"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				host:   tt.fields.host,
				token:  tt.fields.token,
				client: tt.fields.client,
			}
			replays, err := c.GetTweetSelfReplays(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTweet() error = %v, wantErr %v", err, tt.wantErr)
			}
			_ = replays
		})
	}
}
