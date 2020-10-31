package main

import (
	// "fmt"
	"time"	
	"strconv"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"math/rand"
	"github.com/go-vgo/robotgo"
)

func move(nr_moves int){
	
	for m :=0; m < nr_moves; m++ {

		x := rand.Intn(1920)
		y := rand.Intn(1080)
		robotgo.MoveMouse(x, y)
    	time.Sleep(time.Second/10) 

	}

}

func setCountDown(value int, clock *widget.Label){
	formatted := strconv.Itoa(value)
	clock.SetText(formatted)
}


func main() {

	nr_moves := 100
	initial_time_count := 60

	a := app.New()
	w := a.NewWindow("Move THAT Mouse")

	clock := widget.NewLabel("")
	clock_title := widget.NewLabel("")

	nr_moves_text := widget.NewLabel("\n\n")
	time_away_text := widget.NewLabel("\n\n")



	w.SetContent(widget.NewVBox(
		widget.NewLabel("- Don't wanna work?"),
		widget.NewLabel("- Worried somebody will notice if your away from Microsoft Teams?"),
		widget.NewLabel("- Just click 'Start' and we'll randomly move your mouse for you every other minute!"),
		widget.NewLabel("\n\n"),
		clock_title,
		clock,
		nr_moves_text,
		time_away_text,
		widget.NewLabel("\n\n"),
		widget.NewButton("Start", func() {
			time_away_text.SetText("\n\nSeconds you have been away: \n" + strconv.Itoa(0))

			go func() {
				time_away := 0
				t := time.NewTicker(time.Second)
				for range t.C {
					time_away ++
					time_away_text.SetText("\n\nSeconds you have been away: \n" + strconv.Itoa(time_away))
				}
			}()

			go func() {

				for i :=0; true; i++ {
					clock_title.SetText("Seconds until next move:")
					nr_moves_text.SetText("\n\nTotal number of moves: \n" + strconv.Itoa(i))

					setCountDown(initial_time_count, clock)
					var time_count int = initial_time_count - 1

					t := time.NewTicker(time.Second)

					for range t.C {


						if time_count == 0 {
							time_count = initial_time_count
							break
						}
						setCountDown(time_count, clock)
						time_count -= 1
					}

					clock.SetText("")
					clock_title.SetText("MOVING!")
					move(nr_moves)
				}

			}()
		}),
		widget.NewButton("Quit", func() {
			a.Quit()
		}),
	))

	
	w.ShowAndRun()
}
