package sampgo

/*

#cgo CFLAGS: -I./sampgdk

#ifndef GOLANG_APP
#define GOLANG_APP

#include "main.h"

#endif
*/
import "C"

//export goModeInit
func goModeInit() bool {
	evt, ok := events["goModeInit"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func() bool)
	if !ok {
		return false
	}

	C.OnPlayerConnect(C.int(0))

	fn()
	return true
}

//export goModeExit
func goModeExit() bool {
	evt, ok := events["goModeExit"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func() bool)
	if !ok {
		return false
	}

	fn()
	return true
}

//export onPlayerConnect
func onPlayerConnect(playerid C.int) bool {
	evt, ok := events["playerConnect"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(Player) bool)
	if !ok {
		return false
	}

	p := Player{ID: int(playerid)}
	fn(p)
	return true
}

//export onPlayerDisconnect
func onPlayerDisconnect(playerid, reason C.int) bool {
	evt, ok := events["playerDisconnect"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(Player, int) bool)
	if !ok {
		return false
	}

	p := Player{ID: int(playerid)}
	fn(p, int(reason))
	return true
}

//export onPlayerSpawn
func onPlayerSpawn(playerid C.int) bool {
	evt, ok := events["playerSpawn"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(Player) bool)
	if !ok {
		return false
	}

	p := Player{ID: int(playerid)}
	fn(p)
	return true
}

//export onPlayerDeath
func onPlayerDeath(playerid, killerid, reason C.int) bool {
	evt, ok := events["playerDeath"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(Player, Player, int) bool)
	if !ok {
		return false
	}

	p := Player{ID: int(playerid)}
	k := Player{ID: int(killerid)}
	fn(p, k, int(reason))
	return true
}

//export onVehicleSpawn
func onVehicleSpawn(vehicleid C.int) bool {
	evt, ok := events["vehicleSpawn"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(int) bool)
	if !ok {
		return false
	}

	fn(int(vehicleid))
	return true
}

//export onVehicleDeath
func onVehicleDeath(vehicleid, killerid C.int) bool {
	evt, ok := events["vehicleDeath"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(int, Player) bool)
	if !ok {
		return false
	}

	k := Player{ID: int(killerid)}
	fn(int(vehicleid), k)
	return true
}

//export onPlayerText
func onPlayerText(playerid C.int, text *C.char_t) bool {
	evt, ok := events["playerText"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(Player, string) bool)
	if !ok {
		return false
	}

	p := Player{ID: int(playerid)}
	msg := C.constToNonConst(text)

	fn(p, C.GoString(msg))
	return true
}

//export onPlayerCommandText
func onPlayerCommandText(playerid C.int, cmdtext *C.char_t) bool {
	evt, ok := events["playerCommandText"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(Player, string) bool)
	if !ok {
		return false
	}

	p := Player{ID: int(playerid)}
	msg := C.constToNonConst(cmdtext)

	fn(p, C.GoString(msg))
	return true
}

//export onPlayerRequestClass
func onPlayerRequestClass(playerid, classid C.int) bool {
	evt, ok := events["playerRequestClass"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(Player, int) bool)
	if !ok {
		return false
	}

	p := Player{ID: int(playerid)}
	fn(p, int(classid))
	return true
}

//export onPlayerEnterVehicle
func onPlayerEnterVehicle(playerid, vehicleid C.int, ispassenger bool) bool {
	evt, ok := events["playerEnterVehicle"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(Player, int, bool) bool)
	if !ok {
		return false
	}

	p := Player{ID: int(playerid)}
	fn(p, int(vehicleid), ispassenger)
	return true
}

//export onPlayerExitVehicle
func onPlayerExitVehicle(playerid, vehicleid C.int) bool {
	evt, ok := events["playerExitVehicle"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(Player, int) bool)
	if !ok {
		return false
	}

	p := Player{ID: int(playerid)}
	fn(p, int(vehicleid))
	return true
}

//export onPlayerStateChange
func onPlayerStateChange(playerid, newstate, oldstate C.int) bool {
	evt, ok := events["playerStateChange"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(Player, int, int) bool)
	if !ok {
		return false
	}

	p := Player{ID: int(playerid)}
	fn(p, int(newstate), int(oldstate))
	return true
}

//export onPlayerEnterCheckpoint
func onPlayerEnterCheckpoint(playerid C.int) bool {
	evt, ok := events["playerEnterCheckpoint"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(Player) bool)
	if !ok {
		return false
	}

	p := Player{ID: int(playerid)}
	fn(p)
	return true
}

//export onPlayerLeaveCheckpoint
func onPlayerLeaveCheckpoint(playerid C.int) bool {
	evt, ok := events["playerLeaveCheckpoint"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(Player) bool)
	if !ok {
		return false
	}

	p := Player{ID: int(playerid)}
	fn(p)
	return true
}

//export onPlayerEnterRaceCheckpoint
func onPlayerEnterRaceCheckpoint(playerid C.int) bool {
	evt, ok := events["playerEnterRaceCheckpoint"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(Player) bool)
	if !ok {
		return false
	}

	p := Player{ID: int(playerid)}
	fn(p)
	return true
}

//export onPlayerLeaveRaceCheckpoint
func onPlayerLeaveRaceCheckpoint(playerid C.int) bool {
	evt, ok := events["playerLeaveRaceCheckpoint"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(Player) bool)
	if !ok {
		return false
	}

	p := Player{ID: int(playerid)}
	fn(p)
	return true
}

//export onRconCommand
func onRconCommand(cmd *C.char_t) bool {
	evt, ok := events["rconCommand"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(string) bool)
	if !ok {
		return false
	}

	Cmd := C.constToNonConst(cmd)

	fn(C.GoString(Cmd))
	return true
}

//export onPlayerRequestSpawn
func onPlayerRequestSpawn(playerid C.int) bool {
	evt, ok := events["playerRequestSpawn"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(Player) bool)
	if !ok {
		return false
	}

	p := Player{ID: int(playerid)}
	fn(p)
	return true
}

//export onObjectMoved
func onObjectMoved(objectid C.int) bool {
	evt, ok := events["objectMoved"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(int) bool)
	if !ok {
		return false
	}

	fn(int(objectid))
	return true
}

//export onPlayerObjectMoved
func onPlayerObjectMoved(playerid, objectid C.int) bool {
	evt, ok := events["playerObjectMoved"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(Player, int) bool)
	if !ok {
		return false
	}

	p := Player{ID: int(playerid)}
	fn(p, int(objectid))
	return true
}

//export onPlayerPickUpPickup
func onPlayerPickUpPickup(playerid, pickupid C.int) bool {
	evt, ok := events["playerPickUpPickup"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(Player, int) bool)
	if !ok {
		return false
	}

	p := Player{ID: int(playerid)}
	fn(p, int(pickupid))
	return true
}

//export onVehicleMod
func onVehicleMod(playerid, vehicleid, componentid C.int) bool {
	evt, ok := events["vehicleMod"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(Player, int, int) bool)
	if !ok {
		return false
	}

	p := Player{ID: int(playerid)}
	fn(p, int(vehicleid), int(componentid))
	return true
}

//export onEnterExitModShop
func onEnterExitModShop(playerid C.int, enterexit bool, interiorid C.int) bool {
	evt, ok := events["enterExitModShop"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(Player, bool, int) bool)
	if !ok {
		return false
	}

	p := Player{ID: int(playerid)}
	fn(p, enterexit, int(interiorid))
	return true
}

//export onVehiclePaintjob
func onVehiclePaintjob(playerid, vehicleid, paintjobid C.int) bool {
	evt, ok := events["vehiclePaintjob"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(Player, int, int) bool)
	if !ok {
		return false
	}

	p := Player{ID: int(playerid)}
	fn(p, int(vehicleid), int(paintjobid))
	return true
}

//export onVehicleRespray
func onVehicleRespray(playerid, vehicleid, color1, color2 C.int) bool {
	evt, ok := events["vehicleRespray"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(Player, int, int, int) bool)
	if !ok {
		return false
	}

	p := Player{ID: int(playerid)}
	fn(p, int(vehicleid), int(color1), int(color2))
	return true
}

//export onVehicleDamageStatusUpdate
func onVehicleDamageStatusUpdate(vehicleid, playerid C.int) bool {
	evt, ok := events["vehicleDamageStatusUpdate"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(int, Player) bool)
	if !ok {
		return false
	}

	p := Player{ID: int(playerid)}
	fn(int(vehicleid), p)
	return true
}

//export onUnoccupiedVehicleUpdate
func onUnoccupiedVehicleUpdate(vehicleid, playerid, passengerSeat C.int, newX, newY, newZ, velX, velY, velZ C.float) bool {
	evt, ok := events["unoccupiedVehicleUpdate"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(int, Player, int, float32, float32, float32, float32, float32, float32) bool)
	if !ok {
		return false
	}

	p := Player{ID: int(playerid)}
	fn(int(vehicleid), p, int(passengerSeat), float32(newX), float32(newY), float32(newZ), float32(velX), float32(velY), float32(velZ))
	return true
}

//export onPlayerSelectedMenuRow
func onPlayerSelectedMenuRow(playerid, row C.int) bool {
	evt, ok := events["playerSelectedMenuRow"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(Player, int) bool)
	if !ok {
		return false
	}

	p := Player{ID: int(playerid)}
	fn(p, int(row))
	return true
}

//export onPlayerExitedMenu
func onPlayerExitedMenu(playerid C.int) bool {
	evt, ok := events["playerExitedMenu"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(Player) bool)
	if !ok {
		return false
	}

	p := Player{ID: int(playerid)}
	fn(p)
	return true
}

//export onPlayerInteriorChange
func onPlayerInteriorChange(playerid, newinteriorid, oldinteriorid C.int) bool {
	evt, ok := events["playerInteriorChange"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(Player, int, int) bool)
	if !ok {
		return false
	}

	p := Player{ID: int(playerid)}
	fn(p, int(newinteriorid), int(oldinteriorid))
	return true
}

//export onPlayerKeyStateChange
func onPlayerKeyStateChange(playerid, newkeys, oldkeys C.int) bool {
	evt, ok := events["playerKeyStateChange"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(Player, int, int) bool)
	if !ok {
		return false
	}

	p := Player{ID: int(playerid)}
	fn(p, int(newkeys), int(oldkeys))
	return true
}

//export onRconLoginAttempt
func onRconLoginAttempt(ip, password *C.char_t, success bool) bool {
	evt, ok := events["rconLoginAttempt"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(string, string, bool) bool)
	if !ok {
		return false
	}

	IP := C.constToNonConst(ip)
	Password := C.constToNonConst(password)

	fn(C.GoString(IP), C.GoString(Password), success)
	return true
}

//export onPlayerUpdate
func onPlayerUpdate(playerid C.int) bool {
	evt, ok := events["playerUpdate"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(Player) bool)
	if !ok {
		return false
	}

	p := Player{ID: int(playerid)}
	fn(p)
	return true
}

//export onPlayerStreamIn
func onPlayerStreamIn(playerid, forplayerid C.int) bool {
	evt, ok := events["playerStreamIn"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(Player, Player) bool)
	if !ok {
		return false
	}

	p := Player{ID: int(playerid)}
	f := Player{ID: int(forplayerid)}
	fn(p, f)
	return true
}

//export onPlayerStreamOut
func onPlayerStreamOut(playerid, forplayerid C.int) bool {
	evt, ok := events["playerStreamOut"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(Player, Player) bool)
	if !ok {
		return false
	}

	p := Player{ID: int(playerid)}
	f := Player{ID: int(forplayerid)}
	fn(p, f)
	return true
}

//export onVehicleStreamIn
func onVehicleStreamIn(vehicleid, forplayerid C.int) bool {
	evt, ok := events["vehicleStreamIn"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(int, Player) bool)
	if !ok {
		return false
	}

	f := Player{ID: int(forplayerid)}
	fn(int(vehicleid), f)
	return true
}

//export onVehicleStreamOut
func onVehicleStreamOut(vehicleid, forplayerid C.int) bool {
	evt, ok := events["vehicleStreamOut"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(int, Player) bool)
	if !ok {
		return false
	}

	f := Player{ID: int(forplayerid)}
	fn(int(vehicleid), f)
	return true
}

//export onActorStreamIn
func onActorStreamIn(actorid, forplayerid C.int) bool {
	evt, ok := events["actorStreamIn"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(int, Player) bool)
	if !ok {
		return false
	}

	f := Player{ID: int(forplayerid)}
	fn(int(actorid), f)
	return true
}

//export onActorStreamOut
func onActorStreamOut(actorid, forplayerid C.int) bool {
	evt, ok := events["actorStreamOut"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(int, Player) bool)
	if !ok {
		return false
	}

	f := Player{ID: int(forplayerid)}
	fn(int(actorid), f)
	return true
}

//export onDialogResponse
func onDialogResponse(playerid, dialogid, response, listitem C.int, inputtext *C.char_t) bool {
	evt, ok := events["dialogResponse"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(Player, int, int, int, string) bool)
	if !ok {
		return false
	}

	p := Player{ID: int(playerid)}
	text := C.constToNonConst(inputtext)

	fn(p, int(dialogid), int(response), int(listitem), C.GoString(text))
	return true
}

//export onPlayerTakeDamage
func onPlayerTakeDamage(playerid, issuerid C.int, amount C.float, weaponid, bodypart C.int) bool {
	evt, ok := events["playerTakeDamage"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(Player, Player, float32, int, int) bool)
	if !ok {
		return false
	}

	p := Player{ID: int(playerid)}
	i := Player{ID: int(issuerid)}

	fn(p, i, float32(amount), int(weaponid), int(bodypart))
	return true
}

//export onPlayerGiveDamage
func onPlayerGiveDamage(playerid, damagedid C.int, amount C.float, weaponid, bodypart C.int) bool {
	evt, ok := events["playerGiveDamage"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(Player, Player, float32, int, int) bool)
	if !ok {
		return false
	}

	p := Player{ID: int(playerid)}
	d := Player{ID: int(damagedid)}

	fn(p, d, float32(amount), int(weaponid), int(bodypart))
	return true
}

//export onPlayerGiveDamageActor
func onPlayerGiveDamageActor(playerid, damagedActorid C.int, amount C.float, weaponid, bodypart C.int) bool {
	evt, ok := events["playerGiveDamageActor"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(Player, int, float32, int, int) bool)
	if !ok {
		return false
	}

	p := Player{ID: int(playerid)}

	fn(p, int(damagedActorid), float32(amount), int(weaponid), int(bodypart))
	return true
}

//export onPlayerClickMap
func onPlayerClickMap(playerid C.int, fX, fY, fZ C.float) bool {
	evt, ok := events["playerClickMap"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(Player, float32, float32, float32) bool)
	if !ok {
		return false
	}

	p := Player{ID: int(playerid)}

	fn(p, float32(fX), float32(fY), float32(fZ))
	return true
}

//export onPlayerClickTextDraw
func onPlayerClickTextDraw(playerid, clickedid C.int) bool {
	evt, ok := events["playerClickTextDraw"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(Player, int) bool)
	if !ok {
		return false
	}

	p := Player{ID: int(playerid)}

	fn(p, int(clickedid))
	return true
}

//export onPlayerClickPlayerTextDraw
func onPlayerClickPlayerTextDraw(playerid, playertextid C.int) bool {
	evt, ok := events["playerClickPlayerTextDraw"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(Player, int) bool)
	if !ok {
		return false
	}

	p := Player{ID: int(playerid)}

	fn(p, int(playertextid))
	return true
}

//export onIncomingConnection
func onIncomingConnection(playerid C.int, ipAddress *C.char_t, port C.int) bool {
	evt, ok := events["incomingConnection"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(Player, string, int) bool)
	if !ok {
		return false
	}

	p := Player{ID: int(playerid)}
	IP := C.constToNonConst(ipAddress)

	fn(p, C.GoString(IP), int(port))
	return true
}

//export onTrailerUpdate
func onTrailerUpdate(playerid, vehicleid C.int) bool {
	evt, ok := events["trailerUpdate"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(Player, int) bool)
	if !ok {
		return false
	}

	p := Player{ID: int(playerid)}

	fn(p, int(vehicleid))
	return true
}

//export onVehicleSirenStateChange
func onVehicleSirenStateChange(playerid, vehicleid, newstate C.int) bool {
	evt, ok := events["vehicleSirenStateChange"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(Player, int, int) bool)
	if !ok {
		return false
	}

	p := Player{ID: int(playerid)}

	fn(p, int(vehicleid), int(newstate))
	return true
}

//export onPlayerClickPlayer
func onPlayerClickPlayer(playerid, clickedplayerid, source C.int) bool {
	evt, ok := events["playerClickPlayer"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(Player, Player, int) bool)
	if !ok {
		return false
	}

	p := Player{ID: int(playerid)}
	c := Player{ID: int(clickedplayerid)}

	fn(p, c, int(source))
	return true
}

//export onPlayerEditObject
func onPlayerEditObject(playerid C.int, playerobject bool, objectid, response C.int, fX, fY, fZ, fRotX, fRotY, fRotZ C.float) bool {
	evt, ok := events["playerEditObject"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(Player, bool, int, int, float32, float32, float32, float32, float32, float32) bool)
	if !ok {
		return false
	}

	p := Player{ID: int(playerid)}

	fn(p, playerobject, int(objectid), int(response), float32(fX), float32(fY), float32(fZ), float32(fRotX), float32(fRotY), float32(fRotZ))
	return true
}

//export onPlayerEditAttachedObject
func onPlayerEditAttachedObject(playerid, response, index, modelid, boneid C.int, fOffsetX, fOffsetY, fOffsetZ, fRotX, fRotY, fRotZ, fScaleX, fScaleY, fScaleZ C.float) bool {
	evt, ok := events["playerEditAttachedObject"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(Player, int, int, int, int, float32, float32, float32, float32, float32, float32, float32, float32, float32) bool)
	if !ok {
		return false
	}

	p := Player{ID: int(playerid)}

	fn(p, int(response), int(index), int(modelid), int(boneid), float32(fOffsetX), float32(fOffsetY), float32(fOffsetZ), float32(fRotX), float32(fRotY), float32(fRotZ), float32(fScaleX), float32(fScaleY), float32(fScaleZ))
	return true
}

//export onPlayerSelectObject
func onPlayerSelectObject(playerid, otype, objectid, modelid C.int, fX, fY, fZ C.float) bool {
	evt, ok := events["playerSelectObject"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(Player, int, int, int, float32, float32, float32) bool)
	if !ok {
		return false
	}

	p := Player{ID: int(playerid)}

	fn(p, int(otype), int(objectid), int(modelid), float32(fX), float32(fY), float32(fZ))
	return true
}

//export onPlayerWeaponShot
func onPlayerWeaponShot(playerid, weaponid, hittype, hitid C.int, fX, fY, fZ C.float) bool {
	evt, ok := events["playerWeaponShot"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(Player, int, int, Player, float32, float32, float32) bool)
	if !ok {
		return false
	}

	p := Player{ID: int(playerid)}
	h := Player{ID: int(hitid)}

	fn(p, int(weaponid), int(hittype), h, float32(fX), float32(fY), float32(fZ))
	return true
}

//export onPlayerRequestDownload
func onPlayerRequestDownload(playerid, dtype, crc C.int) bool {
	evt, ok := events["playerRequestDownload"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(Player, int, int) bool)
	if !ok {
		return false
	}

	p := Player{ID: int(playerid)}

	fn(p, int(dtype), int(crc))
	return true
}
