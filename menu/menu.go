package menu

import (
	"context"
	"fmt"

	"scorelingo/cyclic"
	"scorelingo/quiz"
	"scorelingo/terminal"
)

type item interface {
	Render(context.Context)
	Title() string
}

type App struct {
	input        chan terminal.KeyEvent
	items        []item
	selectedItem *cyclic.Number
}

func New(input chan terminal.KeyEvent, cancel context.CancelFunc) *App {
	quiz := quiz.New(input)

	exit := exit{
		cancel: cancel,
	}

	items := []item{quiz, &exit}
	return &App{
		input:        input,
		items:        items,
		selectedItem: cyclic.NewNumber(int8(len(items) - 1)),
	}
}

func (m *App) Run(ctx context.Context) {
	m.draw(ctx, nil)

	for {
		select {
		case <-ctx.Done():
			return
		case keyEvent := <-m.input:
			m.draw(ctx, &keyEvent)
		}
	}
}

func (m *App) draw(ctx context.Context, pressedKey *terminal.KeyEvent) {
	if ctx.Err() != nil {
		return
	}

	if pressedKey != nil {
		switch *pressedKey {
		case terminal.UpArrowKey:
			m.selectedItem.Decrement()
		case terminal.DownArrowKey:
			m.selectedItem.Increment()
		case terminal.EnterKey:
			m.items[m.selectedItem.Current()].Render(ctx)
			m.draw(ctx, nil)
			return
		default:
			return
		}
	}

	terminal.ClearScreen()
	for i, item := range m.items {
		row := fmt.Sprintf("* %s", item.Title())
		if i == int(m.selectedItem.Current()) {
			row += " <-"
		}
		row += "\n"
		terminal.Draw(row)
	}
}
