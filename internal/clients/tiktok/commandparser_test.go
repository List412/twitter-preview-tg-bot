package tiktok

import (
	"github.com/list412/tweets-tg-bot/internal/commands"
	"testing"
)

func TestCommandParser_Parse(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name    string
		args    args
		want    commands.ParsedCmdUrl
		wantErr bool
	}{
		{
			name: "default",
			args: args{text: "https://vt.tiktok.com/ZS2VK6K9v/"},
			want: commands.ParsedCmdUrl{
				OriginalUrl: "https://vt.tiktok.com/ZS2VK6K9v/",
				Key:         "ZS2VK6K9v",
				StrippedUrl: "https://vt.tiktok.com/ZS2VK6K9v/",
			},
			wantErr: false,
		},
		{
			name:    "wrong site",
			args:    args{text: "https://vt.titcock.com/ZS2VK6K9v/"},
			want:    commands.ParsedCmdUrl{},
			wantErr: true,
		},
		{
			name:    "empty id",
			args:    args{text: "https://vt.tiktok.com/"},
			want:    commands.ParsedCmdUrl{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := CommandParser{}
			got, err := p.Parse(tt.args.text)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if got.OriginalUrl != tt.want.OriginalUrl {
					t.Errorf("OriginalUrl = %v, want %v", got.OriginalUrl, tt.want.OriginalUrl)
				}
				if got.Key != tt.want.Key {
					t.Errorf("Key = %v, want %v", got.Key, tt.want.Key)
				}
				if got.StrippedUrl != tt.want.StrippedUrl {
					t.Errorf("StrippedUrl = %v, want %v", got.StrippedUrl, tt.want.StrippedUrl)
				}
			}
		})
	}
}
