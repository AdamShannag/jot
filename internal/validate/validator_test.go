package validate

import "testing"

func TestName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "valid name 1",
			args: args{
				name: "user",
			},
			wantErr: false,
		},
		{
			name: "valid name 2",
			args: args{
				name: "system-user",
			},
			wantErr: false,
		},
		{
			name: "empty name",
			args: args{
				name: "",
			},
			wantErr: true,
		},
		{
			name: "invalid name 1",
			args: args{
				name: "user.system",
			},
			wantErr: true,
		},
		{
			name: "invalid name 2",
			args: args{
				name: "user*sys/tem",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Name(tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("Name() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "valid path 1",
			args: args{
				path: "./",
			},
			wantErr: false,
		},
		{
			name: "valid path 2",
			args: args{
				path: "/",
			},
			wantErr: false,
		},
		{
			name: "valid path 3",
			args: args{
				path: "/user/system/",
			},
			wantErr: false,
		},
		{
			name: "valid path 4",
			args: args{
				path: "./user/system/",
			},
			wantErr: false,
		},
		{
			name: "valid path 5",
			args: args{
				path: "../user/system/",
			},
			wantErr: false,
		},
		{
			name: "empty path",
			args: args{
				path: "",
			},
			wantErr: true,
		},
		{
			name: "invalid path 1",
			args: args{
				path: "user.system",
			},
			wantErr: true,
		},
		{
			name: "invalid path 2",
			args: args{
				path: "user*sys/tem",
			},
			wantErr: true,
		},
		{
			name: "invalid path 4",
			args: args{
				path: "system/",
			},
			wantErr: true,
		},
		{
			name: "invalid path 5",
			args: args{
				path: "/system",
			},
			wantErr: true,
		},
		{
			name: "invalid path 6",
			args: args{
				path: "/system//user",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Path(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("Path() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
