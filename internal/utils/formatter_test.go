package utils

import "testing"

func TestDefaultFormatter_FormatName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "given name",
			args: args{
				name: "user",
			},
			want: "user",
		}, {
			name: "given empty name",
			args: args{
				name: "",
			},
			want: "project-name",
		}, {
			name: "given name with special character",
			args: args{
				name: "system-user",
			},
			want: "system-user",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			de := &DefaultFormatter{}
			de.FormatName(&tt.args.name)
			if tt.args.name != tt.want {
				t.Errorf("FormatName() got = %v, want %v", tt.args.name, tt.want)
			}
		})
	}
}

func TestDefaultFormatter_FormatPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "given empty path",
			args: args{
				path: "",
			},
			want: "./",
		}, {
			name: "given dot",
			args: args{
				path: ".",
			},
			want: "./",
		},
		{
			name: "given current directory path",
			args: args{
				path: "./",
			},
			want: "./",
		},
		{
			name: "given relative path",
			args: args{
				path: "./user",
			},
			want: "./user/",
		},
		{
			name: "given relative path with ack slash",
			args: args{
				path: "./user/",
			},
			want: "./user/",
		},
		{
			name: "given root path",
			args: args{
				path: "/",
			},
			want: "/",
		},
		{
			name: "given absolute path without slashes",
			args: args{
				path: "user",
			},
			want: "/user/",
		},
		{
			name: "given absolute path",
			args: args{
				path: "/user",
			},
			want: "/user/",
		},
		{
			name: "given absolute path with back slash",
			args: args{
				path: "/user/",
			},
			want: "/user/",
		},
		{
			name: "given relative path with reverse slash",
			args: args{
				path: "system\\user\\",
			},
			want: "/system/user/",
		},
		{
			name: "given path",
			args: args{
				path: "../user",
			},
			want: "../user/",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			de := &DefaultFormatter{}
			de.FormatPath(&tt.args.path)
			if tt.args.path != tt.want {
				t.Errorf("FormatPath() got = %v, want %v", tt.args.path, tt.want)
			}
		})
	}
}
