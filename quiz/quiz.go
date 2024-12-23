package quiz

import (
	"context"
	"fmt"
	"log"
	"time"

	"scorelingo/terminal"
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

	questions := questions(db)
	fmt.Println(questions)

	time.Sleep(10 * time.Second)
}

func (t *Quiz) Title() string {
	return "Quiz"
}

func (t *Quiz) draw(s string) {
	terminal.ClearScreen()
	terminal.Draw(s)
}
