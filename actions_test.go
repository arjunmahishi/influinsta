package main

import (
	"testing"
)

func Test_getAction(t *testing.T) {
	type args struct {
		actionName string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "reshare video",
			args:    args{actionName: "reshare-video"},
			wantErr: false,
		},
		{
			name:    "random follow",
			args:    args{actionName: "random-follow"},
			wantErr: false,
		},
		{
			name:    "random comment",
			args:    args{actionName: "random-comments"},
			wantErr: false,
		},
		{
			name:    "wrong action name",
			args:    args{actionName: "wrong-name"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := getAction(tt.args.actionName)
			if (err != nil) != tt.wantErr {
				t.Errorf("getAction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
