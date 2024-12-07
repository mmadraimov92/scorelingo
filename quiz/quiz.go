package quiz

import (
	"context"
	"fmt"
	"log"
	"scorelingo/terminal"
	"time"
)

type Quiz struct {
	input chan terminal.KeyEvent
}

func New(input chan terminal.KeyEvent) *Quiz {
	return &Quiz{
		input: input,
	}
}

func (t *Quiz) Render(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case keyEvent := <-t.input:
				if keyEvent == terminal.DeleteKey || keyEvent == terminal.EscapeKey {
					cancel()
					return
				}
			}
		}
	}()

	db, err := load()
	if err != nil {
		log.Fatal(err)
	}

	t.draw(fmt.Sprint(db.words[:10]))
	time.Sleep(10 * time.Second)
	// todo: start quiz
	// todo: multiple choices
}

func (t *Quiz) Title() string {
	return "Quiz"
}

func (t *Quiz) draw(s string) {
	terminal.ClearScreen()
	terminal.Draw(s)
}
