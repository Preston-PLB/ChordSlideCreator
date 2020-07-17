package main

import (
	"fmt"
	"github.com/zserge/lorca"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
)

func main() {
	path, _ := os.UserHomeDir()
	path += string(os.PathSeparator) + "ChordSlideCreator"

	ui, err := lorca.New("", path, 720, 320)

	if execErr, ok := err.(*exec.Error); ok {
		lorca.PromptDownload()
		log.Fatalf("\nChrome could not be started. Do you have Chrome installed? %s\n", execErr)
	} else if err != nil {
		log.Fatalf("\nOps. Something went wrong while starting the application: %s", err)
	}

	defer ui.Close()

	renderEvent := make(chan byte)

	ui.Bind("add", func(a, b int) int { return a + b })
	ui.Bind("render", func() { renderEvent <- 1 })

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	go http.Serve(ln, http.FileServer(http.Dir("./ui")))

	ui.Load(fmt.Sprintf("http://%s", ln.Addr()))

	sigc := make(chan os.Signal)
	signal.Notify(sigc, os.Interrupt)

	go func() {
		for {
			select {
			case <-renderEvent:
				filePath := ui.Eval(`document.getElementsByName('files')[0].value`)
				fmt.Print(filePath)
			}
		}
	}()

	select {
	case <-sigc:
	case <-ui.Done():
	}
}
