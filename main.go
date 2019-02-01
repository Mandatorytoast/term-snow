package main

import (
	"github.com/nsf/termbox-go"
	"math/rand"
	"time"
)

type snowdrop struct {
	XPos   int
	YPos   int
	Symbol rune
	FColor termbox.Attribute
	BColor termbox.Attribute
}

const (
	ForegroundColor = termbox.ColorWhite
	BackgroundColor = termbox.ColorDefault
	Symbol          = '*'
)

func main() {
	rand.Seed(time.Now().UnixNano())
	snowflakes := []snowdrop{}
	err := termbox.Init()
	change_dir := true
	width, _ := termbox.Size() //_ == height
	if err != nil {
		panic(err)
	}

	defer termbox.Close()

	event_queue := make(chan termbox.Event)
	go func() {
		for {
			event_queue <- termbox.PollEvent()
		}
	}()

	draw_tick := time.NewTicker(80 * time.Millisecond)

loop:
	for {
		select {
		case ev := <-event_queue:
			if ev.Type == termbox.EventKey && ev.Key == termbox.KeyEsc {
				break loop
			}
		case <-draw_tick.C:
			termbox.Clear(termbox.ColorWhite, termbox.ColorDefault)
			for i := 0; i < 3; i++ {
				snowflakes = append(snowflakes, SpawnSnowflake(width))
			}
			for i, s := range snowflakes {
				termbox.SetCell(s.XPos, s.YPos, s.Symbol, s.FColor, s.BColor)
				snowflakes[i].YPos += 1
				if change_dir {
					snowflakes[i].XPos += 1
					change_dir = false
				} else {
					snowflakes[i].XPos -= 1
					change_dir = true
				}
			}
			termbox.Flush()
		}
	}
}

func SpawnSnowflake(width int) snowdrop {
	spawn_x := rand.Intn(width)
	snowflake := snowdrop{spawn_x, 0, Symbol, ForegroundColor, BackgroundColor}
	return snowflake
}
