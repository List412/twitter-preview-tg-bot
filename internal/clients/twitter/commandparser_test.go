package twitter

import (
	"github.com/list412/twitter-preview-tg-bot/internal/commands"
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
			name: "tweeter",
			args: args{text: "https://twitter.com/katyaarenina/status/1578299885459492864?s=52&t=OV9qAdNnakuCvMwwg5VZow"},
			want: commands.ParsedCmdUrl{
				OriginalUrl: "https://twitter.com/katyaarenina/status/1578299885459492864?s=52&t=OV9qAdNnakuCvMwwg5VZow",
				Key:         "1578299885459492864",
				StrippedUrl: "https://twitter.com/katyaarenina/status/1578299885459492864",
			},
			wantErr: false,
		},
		{
			name:    "not tweeter",
			args:    args{text: "https://gogle.com/maps/mom"},
			want:    commands.ParsedCmdUrl{},
			wantErr: true,
		},
		{
			name:    "not url",
			args:    args{text: "dsdsdssd.dsfsd.f.sdf"},
			want:    commands.ParsedCmdUrl{},
			wantErr: true,
		},
		{
			name: "tweeter 2",
			args: args{text: "https://twitter.com/AntonBelyayev1/status/1579008825738526720"},
			want: commands.ParsedCmdUrl{
				OriginalUrl: "https://twitter.com/AntonBelyayev1/status/1579008825738526720",
				Key:         "1579008825738526720",
				StrippedUrl: "https://twitter.com/AntonBelyayev1/status/1579008825738526720",
			},
			wantErr: false,
		},
		{
			name: "new twitter url: x.com",
			args: args{text: "https://x.com/gazpachomachine/status/1692189825816789058?s=46&t=RU5XEuJSgxHmd53V-3wyuQ"},
			want: commands.ParsedCmdUrl{
				OriginalUrl: "https://x.com/gazpachomachine/status/1692189825816789058?s=46&t=RU5XEuJSgxHmd53V-3wyuQ",
				Key:         "1692189825816789058",
				StrippedUrl: "https://x.com/gazpachomachine/status/1692189825816789058",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			receiver := CommandParser{}
			got, err := receiver.Parse(tt.args.text)
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
