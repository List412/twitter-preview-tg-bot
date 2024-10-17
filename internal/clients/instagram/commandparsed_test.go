package instagram

import "testing"

func TestCommandParser_Parse(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "url",
			args: args{
				text: "https://www.instagram.com/p/DBOehiuxJtX/?igsh=aGlxbTdxOHd2aml3",
			},
			want:    "https://www.instagram.com/p/DBOehiuxJtX/?igsh=aGlxbTdxOHd2aml3",
			wantErr: false,
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
			if got != tt.want {
				t.Errorf("Parse() got = %v, want %v", got, tt.want)
			}
		})
	}
}
