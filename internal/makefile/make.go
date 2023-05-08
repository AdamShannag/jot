package makefile

import (
	"embed"
	"fmt"
	"os/exec"
	"text/template"
	"time"

	"github.com/AdamShannag/jot/internal/io"
	"github.com/briandowns/spinner"
)

type makefile struct {
	servicePath string
	data        map[string]any
	commands    []string
	out         chan output
	timeout     time.Duration
}

type output struct {
	err error
}

const (
	fname = "Makefile"
	path  = "util.makefile.gotpl"

	initmod = "initmod"
	getmods = "getmods"
	gotidy  = "gotidy"
	gofmt   = "gofmt"
)

var (
	//go:embed files/*
	resources embed.FS
	tmpl      = template.Must(template.ParseFS(resources, "files/*"))
)

func New(servicePath string, timeout time.Duration) *makefile {
	return &makefile{
		servicePath: servicePath,
		data:        make(map[string]any),
		out:         make(chan output),
		timeout:     timeout,
	}
}

func (m *makefile) InitMod(name string) {
	m.data["ModuleName"] = name
	m.commands = append(m.commands, initmod)
}

func (m *makefile) GetGoModules(names ...string) {
	m.data["GoModules"] = names
	m.commands = append(m.commands, getmods)
}

func (m *makefile) GoTidy() {
	m.commands = append(m.commands, gotidy)
}

func (m *makefile) GoFmt() {
	m.commands = append(m.commands, gofmt)
}

func (m *makefile) Close() {
	close(m.out)
}

func (m *makefile) execute(command ...string) {
	cms := []string{"-C", m.servicePath}
	cms = append(cms, command...)

	go func() {
		cmd := exec.Command("make", cms...)
		err := cmd.Run()
		m.out <- output{err}
	}()

	select {
	case <-time.After(m.timeout * time.Second):
		fmt.Println("timed out")
	case x := <-m.out:
		if x.err != nil {
			fmt.Printf("program errored: %s\n", x.err)
		}
	}
}

func (m *makefile) Build() {
	defer m.Close()

	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Suffix = " Jotting..."
	s.Color("cyan")

	s.Start()
	defer s.Stop()

	io.TplToFile(tmpl, path, m.servicePath, fname, m.data)
	for _, cmd := range m.commands {
		m.execute(cmd)
	}
}
