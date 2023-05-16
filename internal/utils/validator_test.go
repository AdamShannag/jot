package utils

import "testing"

func TestDefaultValidator_ValidateName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid name 1",
			args: args{
				name: "user",
			},
			want: true,
		},
		{
			name: "valid name 2",
			args: args{
				name: "system-user",
			},
			want: true,
		},
		{
			name: "empty name",
			args: args{
				name: "",
			},
			want: false,
		},
		{
			name: "invalid name 1",
			args: args{
				name: "user.system",
			},
			want: false,
		},
		{
			name: "invalid name 2",
			args: args{
				name: "user*sys/tem",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			de := &DefaultValidator{}
			if got := de.ValidateName(tt.args.name); got != tt.want {
				t.Errorf("ValidateName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefaultValidator_ValidatePath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid path 1",
			args: args{
				path: "./",
			},
			want: true,
		},
		{
			name: "valid path 2",
			args: args{
				path: "/",
			},
			want: true,
		},
		{
			name: "valid path 3",
			args: args{
				path: "/user/system/",
			},
			want: true,
		},
		{
			name: "valid path 4",
			args: args{
				path: "./user/system/",
			},
			want: true,
		},
		{
			name: "valid path 5",
			args: args{
				path: "../user/system/",
			},
			want: true,
		},
		{
			name: "empty path",
			args: args{
				path: "",
			},
			want: false,
		},
		{
			name: "invalid path 1",
			args: args{
				path: "user.system",
			},
			want: false,
		},
		{
			name: "invalid path 2",
			args: args{
				path: "user*sys/tem",
			},
			want: false,
		},
		{
			name: "invalid path 4",
			args: args{
				path: "system/",
			},
			want: false,
		},
		{
			name: "invalid path 5",
			args: args{
				path: "/system",
			},
			want: false,
		},
		{
			name: "invalid path 6",
			args: args{
				path: "/system//user",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			de := &DefaultValidator{}
			if got := de.ValidatePath(tt.args.path); got != tt.want {
				t.Errorf("ValidatePath() = %v, want %v", got, tt.want)
			}
		})
	}
}
