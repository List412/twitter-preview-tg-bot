package telegram

import "testing"

func Test_parseTweeterUrl(t *testing.T) {
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseTweeterUrl(tt.args.text)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseTweeterUrl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("parseTweeterUrl() got = %v, want %v", got, tt.want)
			}
		})
	}
}
