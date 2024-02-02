package main

import "launcher/gui"

func main()  {
	g := gui.NewGui()
	g.Init()
	g.Start()
}
