package sampgo

import (
	"fmt"
)

// Player implements OO players.
type Player struct {
	ID int
}

// GetName returns the players name.
func (p *Player) GetName() string {
	var name string
	GetPlayerName(p.ID, &name, MaxPlayerName)
	return name
}

// SetName sets the players name.
func (p *Player) SetName(name string) error {
	if len(name) > 24 {
		return fmt.Errorf("name length above 24 chars")
	}

	ret := SetPlayerName(p.ID, name)

	switch ret {
	case 1:
		return nil
	case 0:
		return fmt.Errorf("player already has that name")
	case -1:
		return fmt.Errorf("name can not be changed (it's already in use, too long or has invalid characters)")
	}

	return nil
}

// SendMessage allows you to send a player a message.
func (p *Player) SendMessage(colour int, msg string) error {
	if len(msg) < 1 {
		return fmt.Errorf("msg too short")
	}
	if len(msg) > 144 {
		return fmt.Errorf("message too long")
	}

	if !SendClientMessage(p.ID, colour, msg) {
		return fmt.Errorf("the player is not connected")
	}
	return nil
}

// GetPos gets the player's current position.
func (p *Player) GetPos() (float32, float32, float32, error) {
	var x, y, z float32
	if !GetPlayerPos(p.ID, &x, &y, &z) {
		return x, y, z, fmt.Errorf("GetPlayerPos failure (i.e. player not connected)")
	}
	return x, y, z, nil
}

// SetPos sets the player's current position.
func (p *Player) SetPos(x, y, z float32) error {
	if !SetPlayerPos(p.ID, x, y, z) {
		return fmt.Errorf("player not found")
	}
	return nil
}

// Spawn spawns the player.
func (p *Player) Spawn() error {
	if !SpawnPlayer(p.ID) {
		return fmt.Errorf("player was unable to be spawned")
	}
	return nil
}

// Set player spawn info
func (p *Player) SetSpawnInfo(team, skin int, x, y, z, rotation float32, weapon1, weapon1_ammo, weapon2, weapon2_ammo, weapon3, weapon3_ammo int) {
	SetSpawnInfo(p.ID, team, skin, x, y, z, rotation, weapon1, weapon1_ammo, weapon2, weapon2_ammo, weapon3, weapon3_ammo)
}

func (p *Player) ShowDialog(dialogid, style int, caption, info, button1, button2 string) error {
	if !ShowPlayerDialog(p.ID, dialogid, style, caption, info, button1, button2) {
		return fmt.Errorf("couldn't show dialog")
	}
	return nil
}

func (p *Player) GetFacingAngle() (float32, error) {
	var a float32
	if !GetPlayerFacingAngle(p.ID, &a) {
		return a, fmt.Errorf("invalid player")
	}
	return a, nil
}

func (p *Player) GiveMoney(money int) error {
	if !GivePlayerMoney(p.ID, money) {
		return fmt.Errorf("invalid player")
	}
	return nil
}

func (p *Player) GetMoney() int {
	return GetPlayerMoney(p.ID)
}

func (p *Player) ResetMoney() error {
	if !ResetPlayerMoney(p.ID) {
		return fmt.Errorf("invalid player")
	}
	return nil
}

func (p *Player) IsAdmin() bool {
	return IsPlayerAdmin(p.ID)
}

func (p *Player) GetPlayerState() int {
	return GetPlayerState(p.ID)
}

func (p *Player) GetVehicle() (v Vehicle, err error) {
	v.ID = GetPlayerVehicleID(p.ID)
	if v.ID == 0 {
		err = fmt.Errorf("player is not in a vehicle")
	}
	return
}

func (p *Player) IsInVehicle(v *Vehicle) bool {
	return IsPlayerInVehicle(p.ID, v.ID)
}

func (p *Player) IsInAnyVehicle(v *Vehicle) bool {
	return IsPlayerInAnyVehicle(p.ID)
}

func (p *Player) ApplyAnimation(animlib, animname string, fDelta float32, loop, lockx, locky, freeze bool, time int, forcesync bool) {
	ApplyAnimation(p.ID, animlib, animname, fDelta, loop, lockx, locky, freeze, time, forcesync)
}

func (p *Player) ClearAnimations(forcesync bool) {
	ClearAnimations(p.ID, forcesync)
}

func (p *Player) SetSpecialAction(actionid int) error {
	if !SetPlayerSpecialAction(p.ID, actionid) {
		return fmt.Errorf("invalid player")
	}
	return nil
}

func (p *Player) GetSpecialAction() int {
	return GetPlayerSpecialAction(p.ID)
}

func (p *Player) SelectTextDraw(hovercolor int) {
	SelectTextDraw(p.ID, hovercolor)
}

func (p *Player) CancelSelectTextDraw() {
	CancelSelectTextDraw(p.ID)
}
