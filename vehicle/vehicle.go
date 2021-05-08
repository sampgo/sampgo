package vehicle

import (
	"fmt"
	"github.com/sampgo/sampgo"
	"github.com/sampgo/sampgo/native"
	"github.com/sampgo/sampgo/player"
	"math"
)

type Seat int

const (
	SeatDriver Seat = iota
	SeatFront
	SeatBackLeft
	SeatBackRight
)

type Vehicle struct {
	ID    int
	Model Model
}

type Params struct {
	Engine    int
	Lights    int
	Alarm     int
	Doors     int
	Bonnet    int
	Boot      int
	Objective int
}

func NewVehicle(model Model, x, y, z, rotation float32, color1, color2 uint8, respawn_delay int, addsiren bool) (Vehicle, error) {
	var v Vehicle
	if !IsValidVehicleModel(model) {
		return v, fmt.Errorf("invalid vehicle model")
	}
	v.ID = native.CreateVehicle(int(model), x, y, z, rotation, int(color1), int(color2), respawn_delay, addsiren)
	v.Model = model
	if v.ID == sampgo.InvalidVehicleId {
		return v, fmt.Errorf("couldn't create vehicle")
	}
	return v, nil
}

func (v *Vehicle) Destroy() error {
	if !native.DestroyVehicle(v.ID) {
		return fmt.Errorf("vehicle doesn't exist")
	}
	return nil
}

func (v *Vehicle) GetSpeedFloat64() float64 {
	var x, y, z float32
	native.GetVehicleVelocity(v.ID, &x, &y, &z)

	return math.Sqrt(float64((x*x)+(y*y)+(z*z))) * 136.666667
}

func (v *Vehicle) GetSpeedFloat32() float32 {
	return float32(v.GetSpeedFloat64())
}

func (v *Vehicle) GetSpeedInt() int {
	return int(math.Round(v.GetSpeedFloat64()))
}

func (v *Vehicle) PutPlayer(p *player.Player, seat int) error {
	if !native.PutPlayerInVehicle(p.ID, v.ID, seat) {
		return fmt.Errorf("player or vehicle doesn't exist")
	}
	return nil
}

func (v *Vehicle) GetParams() Params {
	var params Params
	native.GetVehicleParamsEx(v.ID, &params.Engine, &params.Lights, &params.Alarm, &params.Doors, &params.Bonnet, &params.Boot, &params.Objective)
	return params
}

func (v *Vehicle) SetParams(params Params) bool {
	return native.SetVehicleParamsEx(v.ID, params.Engine, params.Lights, params.Alarm, params.Doors, params.Bonnet, params.Boot, params.Objective)
}

func (v *Vehicle) GetModel() Model {
	return v.Model
}

func (v *Vehicle) GetModelName() string {
	return GetVehicleModelName(v.Model)
}
