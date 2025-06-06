package instagram

import (
	"testing"

	"github.com/list412/tweets-tg-bot/internal/commands"
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
			name: "post",
			args: args{
				text: "https://www.instagram.com/p/DBOehiuxJtX/?igsh=aGlxbTdxOHd2aml3",
			},
			want: commands.ParsedCmdUrl{
				OriginalUrl: "https://www.instagram.com/p/DBOehiuxJtX/?igsh=aGlxbTdxOHd2aml3",
				Key:         "DBOehiuxJtX",
				StrippedUrl: "https://www.instagram.com/p/DBOehiuxJtX/",
			},
			wantErr: false,
		},
		{
			name: "story",
			args: args{
				text: "https://www.instagram.com/stories/rodion_balkov/3480992824345244950?utm_source=ig_story_item_share&igsh=NjRhajZ4eWFnd3Jz",
			},
			want: commands.ParsedCmdUrl{
				OriginalUrl: "https://www.instagram.com/stories/rodion_balkov/3480992824345244950?utm_source=ig_story_item_share&igsh=NjRhajZ4eWFnd3Jz",
				Key:         "3480992824345244950",
				StrippedUrl: "https://www.instagram.com/stories/rodion_balkov/3480992824345244950",
			},
			wantErr: false,
		},
		{
			name: "story_no_id",
			args: args{
				text: "https://www.instagram.com/stories/rodion_balkov/",
			},
			want:    commands.ParsedCmdUrl{},
			wantErr: true,
		},
		{
			name: "reel",
			args: args{
				text: "https://www.instagram.com/reel/DAlNNKvNCiP/?igsh=MTgyNXBxa29xdjU5ZA==",
			},
			want: commands.ParsedCmdUrl{
				OriginalUrl: "https://www.instagram.com/reel/DAlNNKvNCiP/?igsh=MTgyNXBxa29xdjU5ZA==",
				Key:         "DAlNNKvNCiP",
				StrippedUrl: "https://www.instagram.com/reel/DAlNNKvNCiP/",
			},
			wantErr: false,
		},
		{
			name: "reel_2",
			args: args{
				text: "https://www.instagram.com/share/reel/BAY8zQd4B_",
			},
			want: commands.ParsedCmdUrl{
				OriginalUrl: "https://www.instagram.com/share/reel/BAY8zQd4B_",
				Key:         "BAY8zQd4B_",
				StrippedUrl: "https://www.instagram.com/share/reel/BAY8zQd4B_",
			},
			wantErr: false,
		},
		{
			name:    "not a full url",
			args:    args{text: "https://www.instagram.com/reel/"},
			want:    commands.ParsedCmdUrl{},
			wantErr: true,
		},
		{
			name: "base_url",
			args: args{
				text: "https://www.instagram.com",
			},
			want:    commands.ParsedCmdUrl{},
			wantErr: true,
		},
		{
			name: "base_url2",
			args: args{
				text: "https://www.instagram.com/",
			},
			want:    commands.ParsedCmdUrl{},
			wantErr: true,
		},
		{
			name: "wrong website",
			args: args{
				text: "https://facebook.com",
			},
			want:    commands.ParsedCmdUrl{},
			wantErr: true,
		},
		{
			name:    "empty",
			args:    args{},
			want:    commands.ParsedCmdUrl{},
			wantErr: true,
		},
		{
			name: "user page",
			args: args{
				text: "https://www.instagram.com/kate616706",
			},
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
