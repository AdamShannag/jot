package makefile

import (
	"embed"
	"fmt"
	"os/exec"
	"text/template"
	"time"

	"github.com/AdamShannag/jot/v1/internal/cleanup"
	"github.com/AdamShannag/jot/v1/internal/io"
	"github.com/AdamShannag/jot/v1/internal/spinner"
)

type Makefile struct {
	servicePath string
	data        map[string]any
	commands    []string
	out         chan output
	timeout     time.Duration
	spinner     *spinner.CustomSpinner
}

type output struct {
	err error
}

const (
	fname = "Makefile.jot"
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

func New(servicePath string, timeout time.Duration, spinner *spinner.CustomSpinner) *Makefile {
	return &Makefile{
		servicePath: servicePath,
		data:        make(map[string]any),
		out:         make(chan output),
		timeout:     timeout,
		spinner:     spinner,
	}
}

func (m *Makefile) InitMod(name string) {
	m.data["ModuleName"] = name
	m.commands = append(m.commands, initmod)
}

func (m *Makefile) GetGoModules(names ...string) {
	m.data["GoModules"] = names
	m.commands = append(m.commands, getmods)
}

func (m *Makefile) GoTidy() {
	m.commands = append(m.commands, gotidy)
}

func (m *Makefile) GoFmt() {
	m.commands = append(m.commands, gofmt)
}

func (m *Makefile) Close() {
	cleanup.Add(func() error {
		m.spinner.Suffix(fmt.Sprintf("%s %s...", "Removing", fname)).Start()
		defer m.spinner.Stop()
		io.RemoveFile(m.servicePath + fname)
		return nil
	})
	close(m.out)
}

func (m *Makefile) execute(command ...string) {
	cms := []string{"-f", fname, "-C", m.servicePath}
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

func (m *Makefile) Build() {
	m.spinner.Start()

	defer m.spinner.Stop()
	defer m.Close()

	io.TplToFile(tmpl, path, m.servicePath, fname, m.data)
	for _, cmd := range m.commands {
		m.spinner.Suffix(fmt.Sprintf("%s: %s...", "Running", cmd))
		m.execute(cmd)
	}
	m.data = make(map[string]any)
}
