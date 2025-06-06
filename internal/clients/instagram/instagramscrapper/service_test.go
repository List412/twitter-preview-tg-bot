package instagramscrapper

import (
	"context"
	"encoding/json"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/list412/twitter-preview-tg-bot/internal/clients/rapidApi"
	"github.com/list412/twitter-preview-tg-bot/internal/commands"
	"github.com/list412/twitter-preview-tg-bot/internal/events/telegram/tgTypes"
	"github.com/list412/twitter-preview-tg-bot/internal/projectpath"

	"github.com/go-test/deep"
)

func TestService_GetPost(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		ctx context.Context
		id  commands.ParsedCmdUrl
	}
	type testCase struct {
		name    string
		fields  fields
		args    args
		want    tgTypes.TweetThread
		wantErr bool
	}

	baseMocksPaths := path.Join(projectpath.ProjectPath(), "tests/mocks/instagramscrapper.Service")

	dirEntries, err := os.ReadDir(baseMocksPaths)
	if err != nil {
		t.Fatal(err)
	}

	rapidApiMock := rapidApi.RapidApiMock{}

	client := NewClient(&rapidApiMock, "host")

	var tests []testCase
	for _, dirEntry := range dirEntries {
		responseFile, err := os.ReadFile(path.Join(baseMocksPaths, dirEntry.Name(), "result.json"))
		if err != nil {
			t.Fatal(err)
		}
		result := tgTypes.TweetThread{}
		err = json.Unmarshal(responseFile, &result)
		if err != nil {
			t.Fatal(err)
		}
		apiResponseFile, err := os.ReadFile(path.Join(baseMocksPaths, dirEntry.Name(), "api_response.json"))
		if err != nil {
			t.Fatal(err)
		}
		url := strings.ReplaceAll(dirEntry.Name(), "_", "/")
		rapidApiMock.SaveResponse(dirEntry.Name(), apiResponseFile)
		if dirEntry.IsDir() {
			tc := testCase{
				name: url,
				fields: fields{
					client: client,
				},
				args: args{
					ctx: context.WithValue(context.Background(), rapidApi.CtxUrlKey{}, dirEntry.Name()),
					id: commands.ParsedCmdUrl{
						OriginalUrl: url,
						Key:         "",
						StrippedUrl: url,
					},
				},
				want:    result,
				wantErr: false,
			}
			tests = append(tests, tc)
		}
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Service{
				client: tt.fields.client,
			}
			got, err := s.GetPost(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := deep.Equal(got, tt.want); diff != nil {
				t.Errorf("GetPost() diff = %v", diff)
			}
		})
	}
}
