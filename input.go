package astar

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Input struct {
	LeftMousePressed  bool
	RightMousePressed bool
	MouseX            int
	MouseY            int
	EnterPressed      bool
	ResetPressed      bool
}

func (i *Input) Update() {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
		i.RightMousePressed = true
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		i.LeftMousePressed = true
		i.MouseX, i.MouseY = ebiten.CursorPosition()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		i.EnterPressed = true
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		i.ResetPressed = true
	}
}

func (i *Input) CopyAndReset() *Input {
	result := &Input{
		RightMousePressed: i.RightMousePressed,
		LeftMousePressed:  i.LeftMousePressed,
		MouseX:            i.MouseX,
		MouseY:            i.MouseY,
		EnterPressed:      i.EnterPressed,
		ResetPressed:      i.ResetPressed,
	}

	i.RightMousePressed = false
	i.LeftMousePressed = false
	i.MouseX = 0
	i.MouseY = 0
	i.EnterPressed = false
	i.ResetPressed = false

	return result
}
