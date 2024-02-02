package gui

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

var (
	exePath string
	exeName string
	sep string
	cmd *exec.Cmd
)

func StartSerice(g *Gui) {
	ex, err := os.Executable()
	if err != nil {
		g.Message.Set(err.Error())
		return
	}
	exePath = filepath.Dir(ex)
	if runtime.GOOS == "windows"{
		sep = "\\"
		exeName = "server.exe"
	}else {
		sep = "/"
		exeName = "server"
	}
	addr,err := g.ApiStr.Get()
	if err != nil {
		g.Message.Set(err.Error())
		return
	}
	permissions,err := g.Premissions.Get()
	if err != nil {
		g.Message.Set(err.Error())
		return
	}
	rootpath,err := g.RootPath.Get()
	if err != nil {
		g.Message.Set(err.Error())
		return
	}
	cmd = exec.Command(exePath+sep+exeName, "-runpath", exePath,"-addr",addr,"-r",permissions,"-rootpath",rootpath)
	err = cmd.Run()
	if err != nil {
		g.Message.Set(err.Error())
		return
	}
}

func StopService() error {
	err := cmd.Process.Kill()
	if err != nil {
		return err
	}
	return nil
}
