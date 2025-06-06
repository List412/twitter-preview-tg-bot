package graphql

import (
	"context"
	"reflect"
	"testing"
)

func TestClient_GetVideo(t *testing.T) {
	type fields struct {
		host string
	}
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *ParsedPost
		wantErr bool
	}{
		{
			name:   "success",
			fields: fields{host: "localhost"},
			args: args{
				ctx: context.Background(),
				id:  "BAY8zQd4B_",
			},
			want:    &ParsedPost{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				host: tt.fields.host,
			}
			got, err := c.GetVideo(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetVideo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetVideo() got = %v, want %v", got, tt.want)
			}
		})
	}
}
