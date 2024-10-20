package twitter

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
			name:    "tweeter",
			args:    args{text: "https://twitter.com/katyaarenina/status/1578299885459492864?s=52&t=OV9qAdNnakuCvMwwg5VZow"},
			want:    "1578299885459492864",
			wantErr: false,
		},
		{
			name:    "not tweeter",
			args:    args{text: "https://gogle.com/maps/mom"},
			want:    "",
			wantErr: true,
		},
		{
			name:    "not url",
			args:    args{text: "dsdsdssd.dsfsd.f.sdf"},
			want:    "",
			wantErr: true,
		},
		{
			name:    "tweeter 2",
			args:    args{text: "https://twitter.com/AntonBelyayev1/status/1579008825738526720"},
			want:    "1579008825738526720",
			wantErr: false,
		},
		{
			name:    "new twitter url: x.com",
			args:    args{text: "https://x.com/gazpachomachine/status/1692189825816789058?s=46&t=RU5XEuJSgxHmd53V-3wyuQ"},
			want:    "1692189825816789058",
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
			if got != tt.want {
				t.Errorf("Parse() got = %v, want %v", got, tt.want)
			}
		})
	}
}
