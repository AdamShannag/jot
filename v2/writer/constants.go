package writer

const (
	API_DIR             string = "api"
	API_ENDPOINTS_DIR   string = "endpoints"
	API_MIDDLEWARES_DIR string = "middlewares"
	BIN_DIR             string = "bin"
	CMD_DIR             string = "cmd"
	DEPLOY_DIR          string = "deploy"
)

const (
	API_FILE  = "api.go"
	MAIN_FILE = "main.go"
)

func serviceDFsMap() dirFilesMap {
	return dirFilesMap{
		API_DIR:    []File{{API_FILE, "", "", nil}},
		BIN_DIR:    []File{},
		CMD_DIR:    []File{},
		DEPLOY_DIR: []File{},
	}
}

func serviceDFs() []*Dir {
	dirs := []*Dir{}

	for k, v := range serviceDFsMap() {
		dirs = append(dirs, &Dir{Name: k, Files: v, Dirs: nil})
	}

	return dirs
}
