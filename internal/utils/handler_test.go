package utils

import (
	"testing"
)

func TestDefaultHandler_HandleName(t *testing.T) {

	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "given valid name",
			args: args{
				name: "user",
			},
			want:    "user",
			wantErr: false,
		},
		{
			name: "given empty name",
			args: args{
				name: "",
			},
			want:    "project-name",
			wantErr: false,
		},
		{
			name: "given invalid name",
			args: args{
				name: "system.user",
			},
			want:    "system.user",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &DefaultHandler{
				formatter: &DefaultFormatter{},
				validator: &DefaultValidator{},
			}
			if err := h.HandleName(&tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("HandleName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.args.name != tt.want {
				t.Errorf("HandleName() got = %v, want %v", tt.args.name, tt.want)
			}
		})
	}
}

func TestDefaultHandler_HandlePath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "given empty path",
			args: args{
				path: "",
			},
			want:    "./",
			wantErr: false,
		},
		{
			name: "given valid relative path",
			args: args{
				path: "./user",
			},
			want:    "./user/",
			wantErr: false,
		},
		{
			name: "given invalid path",
			args: args{
				path: "//user",
			},
			want:    "//user/",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &DefaultHandler{
				formatter: &DefaultFormatter{},
				validator: &DefaultValidator{},
			}
			if err := h.HandlePath(&tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("HandlePath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.args.path != tt.want {
				t.Errorf("HandlePath() got = %v, want %v", tt.args.path, tt.want)
			}
		})
	}
}
