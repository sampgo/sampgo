package player

import (
	"fmt"
	"github.com/sampgo/sampgo"
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
	td := PlayerTextDraw{player: p, textDraw: sampgo.CreatePlayerTextDraw(p.ID, x, y, text)}
	if td.textDraw == sampgo.InvalidTextDraw {
		return td, fmt.Errorf("invalid playertextdraw")
	}
	return td, nil
}

func (p *PlayerTextDraw) Destroy() {
	sampgo.PlayerTextDrawDestroy(p.player.ID, p.textDraw)
}

func (p *PlayerTextDraw) SetString(text string) {
	sampgo.PlayerTextDrawSetString(p.player.ID, p.textDraw, text)
}

func (p *PlayerTextDraw) Show() {
	sampgo.PlayerTextDrawShow(p.player.ID, p.textDraw)
}

func (p *PlayerTextDraw) Hide() {
	sampgo.PlayerTextDrawHide(p.player.ID, p.textDraw)
}

func (p *PlayerTextDraw) Font(font int) {
	sampgo.PlayerTextDrawFont(p.player.ID, p.textDraw, font)
}

func (p *PlayerTextDraw) UseBox(use bool) {
	sampgo.PlayerTextDrawUseBox(p.player.ID, p.textDraw, use)
}

func (p *PlayerTextDraw) SetAlignment(align int) {
	p.align = align
	sampgo.PlayerTextDrawAlignment(p.player.ID, p.textDraw, p.align)
}

func (p *PlayerTextDraw) SetTextSize(x, y float32) {
	if p.align == 2 {
		x, y = y, x
	}
	sampgo.PlayerTextDrawTextSize(p.player.ID, p.textDraw, x, y)
}

func (p *PlayerTextDraw) SetColor(color int) {
	sampgo.PlayerTextDrawColor(p.player.ID, p.textDraw, color)
}

func (p *PlayerTextDraw) SetBoxColor(color int) {
	sampgo.PlayerTextDrawBoxColor(p.player.ID, p.textDraw, color)
}
func (p *PlayerTextDraw) SetBackgroundColor(color int) {
	sampgo.PlayerTextDrawBackgroundColor(p.player.ID, p.textDraw, color)
}
