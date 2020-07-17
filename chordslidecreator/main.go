package main

import (
	"fmt"
	"github.com/Preston-PLB/choRenderer"
	"github.com/leaanthony/mewn"
	"github.com/mitchellh/mapstructure"
	"github.com/wailsapp/wails"
	"log"
)

func main() {

	js := mewn.String("./frontend/dist/app.js")
	css := mewn.String("./frontend/dist/app.css")

	app := wails.CreateApp(&wails.AppConfig{
		Width:  1024,
		Height: 383,
		Title:  "ChordSlideCreator",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})

	app.Bind(initSong)
	app.Run()
}

func initSong(data map[string]interface{}) {
	var result choRenderer.SongSettings

	fmt.Printf("data: %#v\n", data)

	err := mapstructure.Decode(data, &result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("result: %#v\n", result)

	song := choRenderer.Song{}

	song.LoadSettings(&result)

	fmt.Println(song)
}
