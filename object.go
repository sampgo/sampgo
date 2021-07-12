package sampgo

import (
	"fmt"
	"time"
)

type Object struct {
	ID int
}

func (o *Object) GetID() int {
	return o.ID
}

func NewObject(modelid int, x, y, z, rX, rY, rZ, drawDistance float32) (o *Object) {
	o = new(Object)
	o.ID = CreateObject(modelid, x, y, z, rX, rY, rZ, drawDistance)
	return
}

func (o *Object) Destroy() {
	DestroyObject(o.ID)
}

func (o *Object) IsValid() bool {
	return IsValidObject(o.ID)
}

// Returns the time it will take for the object to move
func (o *Object) Move(x, y, z, speed, rotX, rotY, rotZ float32) time.Duration {
	return time.Duration(MoveObject(o.ID, x, y, z, speed, rotX, rotY, rotZ)) * time.Millisecond
}

func (o *Object) IsMoving() bool {
	return IsObjectMoving(o.ID)
}

func (o *Object) Stop() {
	StopObject(o.ID)
}

func (o *Object) SetPos(x, y, z float32) {
	SetObjectPos(o.ID, x, y, z)
}
func (o *Object) SetRot(rx, ry, rz float32) {
	SetObjectRot(o.ID, rx, ry, rz)
}

func (o *Object) GetPos() (x, y, z float32, err error) {
	if !GetObjectPos(o.ID, &x, &y, &z) {
		err = fmt.Errorf("invalid object")
	}
	return
}

func (o *Object) GetRot() (rx, ry, rz float32, err error) {
	if !GetObjectRot(o.ID, &rx, &ry, &rz) {
		err = fmt.Errorf("invalid object")
	}
	return
}

type PlayerObject struct {
	ID     int
	player Player
}

func (o *PlayerObject) GetID() int {
	return o.ID
}

func NewPlayerObject(modelid int, x, y, z, rX, rY, rZ, drawDistance float32) (o *PlayerObject) {
	o = new(PlayerObject)
	o.ID = CreatePlayerObject(o.player.ID, modelid, x, y, z, rX, rY, rZ, drawDistance)
	return
}

func (o *PlayerObject) Destroy() {
	DestroyPlayerObject(o.player.ID, o.ID)
}

func (o *PlayerObject) IsValid() bool {
	return IsValidPlayerObject(o.player.ID, o.ID)
}

// Returns the time it will take for the object to move
func (o *PlayerObject) Move(x, y, z, speed, rotX, rotY, rotZ float32) time.Duration {
	return time.Duration(MovePlayerObject(o.player.ID, o.ID, x, y, z, speed, rotX, rotY, rotZ)) * time.Millisecond
}

func (o *PlayerObject) IsMoving() bool {
	return IsPlayerObjectMoving(o.player.ID, o.ID)
}

func (o *PlayerObject) Stop() {
	StopPlayerObject(o.player.ID, o.ID)
}

func (o *PlayerObject) SetPos(x, y, z float32) {
	SetPlayerObjectPos(o.player.ID, o.ID, x, y, z)
}
func (o *PlayerObject) SetRot(rx, ry, rz float32) {
	SetPlayerObjectRot(o.player.ID, o.ID, rx, ry, rz)
}

func (o *PlayerObject) GetPos() (x, y, z float32, err error) {
	if !GetPlayerObjectPos(o.player.ID, o.ID, &x, &y, &z) {
		err = fmt.Errorf("invalid object")
	}
	return
}

func (o *PlayerObject) GetRot() (rx, ry, rz float32, err error) {
	if !GetPlayerObjectRot(o.player.ID, o.ID, &rx, &ry, &rz) {
		err = fmt.Errorf("invalid object")
	}
	return
}
