package format

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
			got := ProjectName(tt.args.name)
			if got != tt.want {
				t.Errorf("ProjectName() got = %v, want %v", got, tt.want)
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
			got := Path(tt.args.path)
			if got != tt.want {
				t.Errorf("Path() got = %v, want %v", got, tt.want)
			}
		})
	}
}
