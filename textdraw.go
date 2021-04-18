package sampgo

import "fmt"

const (
	FontSanAndreas = iota
	FontClear
	FontCapitalClear
	FontGTA
	FontSprite
)

type PlayerTextDraw struct {
	player   *Player
	textDraw int
	align    int
}

func (p *Player) NewTextDraw(x, y float32, text string) (PlayerTextDraw, error) {
	td := PlayerTextDraw{player: p, textDraw: CreatePlayerTextDraw(p.ID, x, y, text)}
	if td.textDraw == InvalidTextDraw {
		return td, fmt.Errorf("invalid playertextdraw")
	}
	return td, nil
}

func (p *PlayerTextDraw) Destroy() {
	PlayerTextDrawDestroy(p.player.ID, p.textDraw)
}

func (p *PlayerTextDraw) SetString(text string) {
	PlayerTextDrawSetString(p.player.ID, p.textDraw, text)
}

func (p *PlayerTextDraw) Show() {
	PlayerTextDrawShow(p.player.ID, p.textDraw)
}

func (p *PlayerTextDraw) Hide() {
	PlayerTextDrawHide(p.player.ID, p.textDraw)
}

func (p *PlayerTextDraw) Font(font int) {
	PlayerTextDrawFont(p.player.ID, p.textDraw, font)
}

func (p *PlayerTextDraw) UseBox(use bool) {
	PlayerTextDrawUseBox(p.player.ID, p.textDraw, use)
}

func (p *PlayerTextDraw) SetAlignment(align int) {
	p.align = align
	PlayerTextDrawAlignment(p.player.ID, p.textDraw, p.align)
}

func (p *PlayerTextDraw) SetTextSize(x, y float32) {
	if p.align == 2 {
		x, y = y, x
	}
	PlayerTextDrawTextSize(p.player.ID, p.textDraw, x, y)
}

func (p *PlayerTextDraw) SetColor(color int) {
	PlayerTextDrawColor(p.player.ID, p.textDraw, color)
}

var SetColour = (*PlayerTextDraw).SetColor

func (p *PlayerTextDraw) SetBoxColor(color int) {
	PlayerTextDrawBoxColor(p.player.ID, p.textDraw, color)
}

var SetBoxColour = (*PlayerTextDraw).SetBoxColor

func (p *PlayerTextDraw) SetBackgroundColor(color int) {
	PlayerTextDrawBackgroundColor(p.player.ID, p.textDraw, color)
}

var SetBackgroundColour = (*PlayerTextDraw).SetBackgroundColor
