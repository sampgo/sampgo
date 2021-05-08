package player

import (
	"fmt"
	"github.com/sampgo/sampgo"
	"github.com/sampgo/sampgo/native"
)

var (
	SetColour           = (*PlayerTextDraw).SetColor
	SetBoxColour        = (*PlayerTextDraw).SetBoxColor
	SetBackgroundColour = (*PlayerTextDraw).SetBackgroundColor
)

type PlayerTextDraw struct {
	player   *Player
	textDraw int
	align    int
}

func (p *Player) NewTextDraw(x, y float32, text string) (PlayerTextDraw, error) {
	td := PlayerTextDraw{player: p, textDraw: native.CreatePlayerTextDraw(p.ID, x, y, text)}
	if td.textDraw == sampgo.InvalidTextDraw {
		return td, fmt.Errorf("invalid playertextdraw")
	}
	return td, nil
}

func (p *PlayerTextDraw) Destroy() {
	native.PlayerTextDrawDestroy(p.player.ID, p.textDraw)
}

func (p *PlayerTextDraw) SetString(text string) {
	native.PlayerTextDrawSetString(p.player.ID, p.textDraw, text)
}

func (p *PlayerTextDraw) Show() {
	native.PlayerTextDrawShow(p.player.ID, p.textDraw)
}

func (p *PlayerTextDraw) Hide() {
	native.PlayerTextDrawHide(p.player.ID, p.textDraw)
}

func (p *PlayerTextDraw) Font(font int) {
	native.PlayerTextDrawFont(p.player.ID, p.textDraw, font)
}

func (p *PlayerTextDraw) UseBox(use bool) {
	native.PlayerTextDrawUseBox(p.player.ID, p.textDraw, use)
}

func (p *PlayerTextDraw) SetAlignment(align int) {
	p.align = align
	native.PlayerTextDrawAlignment(p.player.ID, p.textDraw, p.align)
}

func (p *PlayerTextDraw) SetTextSize(x, y float32) {
	if p.align == 2 {
		x, y = y, x
	}
	native.PlayerTextDrawTextSize(p.player.ID, p.textDraw, x, y)
}

func (p *PlayerTextDraw) SetColor(color int) {
	native.PlayerTextDrawColor(p.player.ID, p.textDraw, color)
}

func (p *PlayerTextDraw) SetBoxColor(color int) {
	native.PlayerTextDrawBoxColor(p.player.ID, p.textDraw, color)
}
func (p *PlayerTextDraw) SetBackgroundColor(color int) {
	native.PlayerTextDrawBackgroundColor(p.player.ID, p.textDraw, color)
}
