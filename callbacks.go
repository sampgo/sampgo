package sampgo

/*
#cgo windows CFLAGS: -I./lib -I./lib/amx -Wno-attributes -Wno-implicit-function-declaration
#cgo windows CFLAGS: -DHAVE_INTTYPES_H -DHAVE_MALLOC_H -DHAVE_STDINT_H -DWIN32
#cgo windows LDFLAGS: -Wl,--subsystem,windows,--kill-at

#cgo linux CFLAGS: -I./lib -I./lib/amx -Wno-attributes -Wno-implicit-function-declaration
#cgo linux CFLAGS: -DHAVE_INTTYPES_H -DHAVE_MALLOC_H -DHAVE_STDINT_H -DLINUX -D_GNU_SOURCE
#cgo linux LDFLAGS: -ldl

#ifndef GOLANG_APP
#define GOLANG_APP

#include "main.h"

#endif
*/
import "C"
import "reflect"

//export callEvent
func callEvent(funcName *C.char_t, format *C.char_t, params []interface{}) bool {
	const CallbackMaxParams = 32

	name := C.GoString(C.constToNonConst(funcName))
	_, ok := events[name]
	if !ok {
		Print("Called an event that is not registered by sampgo.")
		return false
	}

	Print("callEvent (1)")

	// fn, ok := evt.Handler.(func(event) bool)
	// _ = fn
	// if !ok {
	// 	Print("Failed to handle an event that is registered by sampgo.")
	// 	return false
	// }

	Print("callEvent (2)")

	if len(params) > CallbackMaxParams {
		Print("The number of maximum parameters is 32.")
		return false
	}

	Print("callEvent (3)")

	// TODO: Reflect doesn't know of the C data types. This means that we need to manually map over to Go types!
	// TODO: POC code below

	f := reflect.ValueOf(events[name])
	in := make([]reflect.Value, len(params))
	// for k, param := range params {
	// 	// in[k] = reflect.ValueOf(param)
	// 	if param == nil {
	// 		Print("It is a nil")
	// 	} else {
	// 		switch param.(type) {
	// 		case C.int:
	// 			Print("It is a int")
	// 			in[k] = int(param)
	// 		case C.char_t:
	// 			Print("It is a string")
	// 			in[k] = C.GoString(C.constToNonConst(param))
	// 		case C.bool:
	// 			Print("It is a bool")
	// 			in[k] = bool(param)
	// 		case C.float:
	// 			Print("It is a float")
	// 			in[k] = float32(param)
	// 		default:
	// 			Print("Unknown type")
	// 		}
	// 	}
	// }

	// Print("callEvent (4)")

	// if len(params) != f.Type().NumIn() {
	// 	Print("The number of parameters is out of index.")
	// 	return false
	// }

	// Print("callEvent (5)")

	f.Call(in)
	// fn(event{Handler: params})

	Print("callEvent (6)")
	return true
}

//export onGameModeInit
func onGameModeInit() bool {
	evt, ok := events["goModeInit"]
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

//export onGameModeExit
func onGameModeExit() bool {
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
	fn(Player{ID: int(playerid)})
	return true
}

//export onPlayerDisconnect
func onPlayerDisconnect(playerid C.int, reason C.int) bool {
	evt, ok := events["playerDisconnect"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(Player, int) bool)
	if !ok {
		return false
	}
	fn(Player{ID: int(playerid)}, int(reason))
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
	fn(Player{ID: int(playerid)})
	return true
}

//export onPlayerDeath
func onPlayerDeath(playerid C.int, killerid C.int, reason C.int) bool {
	evt, ok := events["playerDeath"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(Player, Player, int) bool)
	if !ok {
		return false
	}
	fn(Player{ID: int(playerid)}, Player{ID: int(killerid)}, int(reason))
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
func onVehicleDeath(vehicleid C.int, killerid C.int) bool {
	evt, ok := events["vehicleDeath"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(int, Player) bool)
	if !ok {
		return false
	}
	fn(int(vehicleid), Player{ID: int(killerid)})
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
	fn(Player{ID: int(playerid)}, C.GoString(C.constToNonConst(text)))
	return true
}

//export onPlayerCommandText
func onPlayerCommandText(playerid C.int, cmdtext *C.char_t) bool {
	evt, ok := events["playerCommandText"]
	if !ok {
		return true
	}

	fn, ok := evt.Handler.(func(Player, string) bool)
	if !ok {
		return true
	}
	return fn(Player{ID: int(playerid)}, C.GoString(C.constToNonConst(cmdtext)))
}

//export onPlayerRequestClass
func onPlayerRequestClass(playerid C.int, classid C.int) bool {
	evt, ok := events["playerRequestClass"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(Player, int) bool)
	if !ok {
		return false
	}
	fn(Player{ID: int(playerid)}, int(classid))
	return true
}

//export onPlayerEnterVehicle
func onPlayerEnterVehicle(playerid C.int, vehicleid C.int, ispassenger C.bool) bool {
	evt, ok := events["playerEnterVehicle"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(Player, int, bool) bool)
	if !ok {
		return false
	}
	fn(Player{ID: int(playerid)}, int(vehicleid), bool(ispassenger))
	return true
}

//export onPlayerExitVehicle
func onPlayerExitVehicle(playerid C.int, vehicleid C.int) bool {
	evt, ok := events["playerExitVehicle"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(Player, int) bool)
	if !ok {
		return false
	}
	fn(Player{ID: int(playerid)}, int(vehicleid))
	return true
}

//export onPlayerStateChange
func onPlayerStateChange(playerid C.int, newstate C.int, oldstate C.int) bool {
	evt, ok := events["playerStateChange"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(Player, int, int) bool)
	if !ok {
		return false
	}
	fn(Player{ID: int(playerid)}, int(newstate), int(oldstate))
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
	fn(Player{ID: int(playerid)})
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
	fn(Player{ID: int(playerid)})
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
	fn(Player{ID: int(playerid)})
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
	fn(Player{ID: int(playerid)})
	return true
}

//export onRconCommand
func onRconCommand(cmd *C.char_t) bool {
	evt, ok := events["rconCommand"]
	if !ok {
		return true
	}

	fn, ok := evt.Handler.(func(string) bool)
	if !ok {
		return true
	}
	fn(C.GoString(C.constToNonConst(cmd)))
	return false
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
	fn(Player{ID: int(playerid)})
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
func onPlayerObjectMoved(playerid C.int, objectid C.int) bool {
	evt, ok := events["playerObjectMoved"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(Player, int) bool)
	if !ok {
		return false
	}
	fn(Player{ID: int(playerid)}, int(objectid))
	return true
}

//export onPlayerPickUpPickup
func onPlayerPickUpPickup(playerid C.int, pickupid C.int) bool {
	evt, ok := events["playerPickUpPickup"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(Player, int) bool)
	if !ok {
		return false
	}
	fn(Player{ID: int(playerid)}, int(pickupid))
	return true
}

//export onVehicleMod
func onVehicleMod(playerid C.int, vehicleid C.int, componentid C.int) bool {
	evt, ok := events["vehicleMod"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(Player, int, int) bool)
	if !ok {
		return false
	}
	fn(Player{ID: int(playerid)}, int(vehicleid), int(componentid))
	return true
}

//export onEnterExitModShop
func onEnterExitModShop(playerid C.int, enterexit C.bool, interiorid C.int) bool {
	evt, ok := events["enterExitModShop"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(Player, bool, int) bool)
	if !ok {
		return false
	}
	fn(Player{ID: int(playerid)}, bool(enterexit), int(interiorid))
	return true
}

//export onVehiclePaintjob
func onVehiclePaintjob(playerid C.int, vehicleid C.int, paintjobid C.int) bool {
	evt, ok := events["vehiclePaintjob"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(Player, int, int) bool)
	if !ok {
		return false
	}
	fn(Player{ID: int(playerid)}, int(vehicleid), int(paintjobid))
	return true
}

//export onVehicleRespray
func onVehicleRespray(playerid C.int, vehicleid C.int, color1 C.int, color2 C.int) bool {
	evt, ok := events["vehicleRespray"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(Player, int, int, int) bool)
	if !ok {
		return false
	}
	fn(Player{ID: int(playerid)}, int(vehicleid), int(color1), int(color2))
	return true
}

//export onVehicleDamageStatusUpdate
func onVehicleDamageStatusUpdate(vehicleid C.int, playerid C.int) bool {
	evt, ok := events["vehicleDamageStatusUpdate"]
	if !ok {
		return true
	}

	fn, ok := evt.Handler.(func(int, Player) bool)
	if !ok {
		return true
	}
	fn(int(vehicleid), Player{ID: int(playerid)})
	return false
}

//export onUnoccupiedVehicleUpdate
func onUnoccupiedVehicleUpdate(vehicleid C.int, playerid C.int, passenger_seat C.int, new_x C.float, new_y C.float, new_z C.float, vel_x C.float, vel_y C.float, vel_z C.float) bool {
	evt, ok := events["unoccupiedVehicleUpdate"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(int, Player, int, float32, float32, float32, float32, float32, float32) bool)
	if !ok {
		return false
	}
	fn(int(vehicleid), Player{ID: int(playerid)}, int(passenger_seat), float32(new_x), float32(new_y), float32(new_z), float32(vel_x), float32(vel_y), float32(vel_z))
	return true
}

//export onPlayerSelectedMenuRow
func onPlayerSelectedMenuRow(playerid C.int, row C.int) bool {
	evt, ok := events["playerSelectedMenuRow"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(Player, int) bool)
	if !ok {
		return false
	}
	fn(Player{ID: int(playerid)}, int(row))
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
	fn(Player{ID: int(playerid)})
	return true
}

//export onPlayerInteriorChange
func onPlayerInteriorChange(playerid C.int, newinteriorid C.int, oldinteriorid C.int) bool {
	evt, ok := events["playerInteriorChange"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(Player, int, int) bool)
	if !ok {
		return false
	}
	fn(Player{ID: int(playerid)}, int(newinteriorid), int(oldinteriorid))
	return true
}

//export onPlayerKeyStateChange
func onPlayerKeyStateChange(playerid C.int, newkeys C.int, oldkeys C.int) bool {
	evt, ok := events["playerKeyStateChange"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(Player, int, int) bool)
	if !ok {
		return false
	}
	fn(Player{ID: int(playerid)}, int(newkeys), int(oldkeys))
	return true
}

//export onRconLoginAttempt
func onRconLoginAttempt(ip *C.char_t, password *C.char_t, success C.bool) bool {
	evt, ok := events["rconLoginAttempt"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(string, string, bool) bool)
	if !ok {
		return false
	}
	fn(C.GoString(C.constToNonConst(ip)), C.GoString(C.constToNonConst(password)), bool(success))
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
	fn(Player{ID: int(playerid)})
	return true
}

//export onPlayerStreamIn
func onPlayerStreamIn(playerid C.int, forplayerid C.int) bool {
	evt, ok := events["playerStreamIn"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(Player, int) bool)
	if !ok {
		return false
	}
	fn(Player{ID: int(playerid)}, int(forplayerid))
	return true
}

//export onPlayerStreamOut
func onPlayerStreamOut(playerid C.int, forplayerid C.int) bool {
	evt, ok := events["playerStreamOut"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(Player, int) bool)
	if !ok {
		return false
	}
	fn(Player{ID: int(playerid)}, int(forplayerid))
	return true
}

//export onVehicleStreamIn
func onVehicleStreamIn(vehicleid C.int, forplayerid C.int) bool {
	evt, ok := events["vehicleStreamIn"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(int, int) bool)
	if !ok {
		return false
	}
	fn(int(vehicleid), int(forplayerid))
	return true
}

//export onVehicleStreamOut
func onVehicleStreamOut(vehicleid C.int, forplayerid C.int) bool {
	evt, ok := events["vehicleStreamOut"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(int, int) bool)
	if !ok {
		return false
	}
	fn(int(vehicleid), int(forplayerid))
	return true
}

//export onActorStreamIn
func onActorStreamIn(actorid C.int, forplayerid C.int) bool {
	evt, ok := events["actorStreamIn"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(int, int) bool)
	if !ok {
		return false
	}
	fn(int(actorid), int(forplayerid))
	return true
}

//export onActorStreamOut
func onActorStreamOut(actorid C.int, forplayerid C.int) bool {
	evt, ok := events["actorStreamOut"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(int, int) bool)
	if !ok {
		return false
	}
	fn(int(actorid), int(forplayerid))
	return true
}

//export onDialogResponse
func onDialogResponse(playerid C.int, dialogid C.int, response C.int, listitem C.int, inputtext *C.char_t) bool {
	evt, ok := events["dialogResponse"]
	if !ok {
		return true
	}

	fn, ok := evt.Handler.(func(Player, int, int, int, string) bool)
	if !ok {
		return true
	}
	fn(Player{ID: int(playerid)}, int(dialogid), int(response), int(listitem), C.GoString(C.constToNonConst(inputtext)))
	return false
}

//export onPlayerTakeDamage
func onPlayerTakeDamage(playerid C.int, issuerid C.int, amount C.float, weaponid C.int, bodypart C.int) bool {
	evt, ok := events["playerTakeDamage"]
	if !ok {
		return true
	}

	fn, ok := evt.Handler.(func(Player, Player, float32, int, int) bool)
	if !ok {
		return true
	}
	fn(Player{ID: int(playerid)}, Player{ID: int(issuerid)}, float32(amount), int(weaponid), int(bodypart))
	return false
}

//export onPlayerGiveDamage
func onPlayerGiveDamage(playerid C.int, damagedid C.int, amount C.float, weaponid C.int, bodypart C.int) bool {
	evt, ok := events["playerGiveDamage"]
	if !ok {
		return true
	}

	fn, ok := evt.Handler.(func(Player, int, float32, int, int) bool)
	if !ok {
		return true
	}
	fn(Player{ID: int(playerid)}, int(damagedid), float32(amount), int(weaponid), int(bodypart))
	return false
}

//export onPlayerGiveDamageActor
func onPlayerGiveDamageActor(playerid C.int, damaged_actorid C.int, amount C.float, weaponid C.int, bodypart C.int) bool {
	evt, ok := events["playerGiveDamageActor"]
	if !ok {
		return true
	}

	fn, ok := evt.Handler.(func(Player, int, float32, int, int) bool)
	if !ok {
		return true
	}
	fn(Player{ID: int(playerid)}, int(damaged_actorid), float32(amount), int(weaponid), int(bodypart))
	return false
}

//export onPlayerClickMap
func onPlayerClickMap(playerid C.int, fX C.float, fY C.float, fZ C.float) bool {
	evt, ok := events["playerClickMap"]
	if !ok {
		return true
	}

	fn, ok := evt.Handler.(func(Player, float32, float32, float32) bool)
	if !ok {
		return true
	}
	fn(Player{ID: int(playerid)}, float32(fX), float32(fY), float32(fZ))
	return false
}

//export onPlayerClickTextDraw
func onPlayerClickTextDraw(playerid C.int, clickedid C.int) bool {
	evt, ok := events["playerClickTextDraw"]
	if !ok {
		return true
	}

	fn, ok := evt.Handler.(func(Player, int) bool)
	if !ok {
		return true
	}
	fn(Player{ID: int(playerid)}, int(clickedid))
	return false
}

//export onPlayerClickPlayerTextDraw
func onPlayerClickPlayerTextDraw(playerid C.int, playertextid C.int) bool {
	evt, ok := events["playerClickPlayerTextDraw"]
	if !ok {
		return true
	}

	fn, ok := evt.Handler.(func(Player, int) bool)
	if !ok {
		return true
	}
	fn(Player{ID: int(playerid)}, int(playertextid))
	return false
}

//export onIncomingConnection
func onIncomingConnection(playerid C.int, ip_address *C.char_t, port C.int) bool {
	evt, ok := events["incomingConnection"]
	if !ok {
		return true
	}

	fn, ok := evt.Handler.(func(Player, string, int) bool)
	if !ok {
		return true
	}
	fn(Player{ID: int(playerid)}, C.GoString(C.constToNonConst(ip_address)), int(port))
	return false
}

//export onTrailerUpdate
func onTrailerUpdate(playerid C.int, vehicleid C.int) bool {
	evt, ok := events["trailerUpdate"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(Player, int) bool)
	if !ok {
		return false
	}
	fn(Player{ID: int(playerid)}, int(vehicleid))
	return true
}

//export onVehicleSirenStateChange
func onVehicleSirenStateChange(playerid C.int, vehicleid C.int, newstate C.int) bool {
	evt, ok := events["vehicleSirenStateChange"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(Player, int, int) bool)
	if !ok {
		return false
	}
	fn(Player{ID: int(playerid)}, int(vehicleid), int(newstate))
	return true
}

//export onPlayerClickPlayer
func onPlayerClickPlayer(playerid C.int, clickedplayerid C.int, source C.int) bool {
	evt, ok := events["playerClickPlayer"]
	if !ok {
		return true
	}

	fn, ok := evt.Handler.(func(Player, int, int) bool)
	if !ok {
		return true
	}
	fn(Player{ID: int(playerid)}, int(clickedplayerid), int(source))
	return false
}

//export onPlayerEditObject
func onPlayerEditObject(playerid C.int, playerobject C.bool, objectid C.int, response C.int, fX C.float, fY C.float, fZ C.float, fRotX C.float, fRotY C.float, fRotZ C.float) bool {
	evt, ok := events["playerEditObject"]
	if !ok {
		return true
	}

	fn, ok := evt.Handler.(func(Player, bool, int, int, float32, float32, float32, float32, float32, float32) bool)
	if !ok {
		return true
	}
	fn(Player{ID: int(playerid)}, bool(playerobject), int(objectid), int(response), float32(fX), float32(fY), float32(fZ), float32(fRotX), float32(fRotY), float32(fRotZ))
	return false
}

//export onPlayerEditAttachedObject
func onPlayerEditAttachedObject(playerid C.int, response C.int, index C.int, modelid C.int, boneid C.int, fOffsetX C.float, fOffsetY C.float, fOffsetZ C.float, fRotX C.float, fRotY C.float, fRotZ C.float, fScaleX C.float, fScaleY C.float, fScaleZ C.float) bool {
	evt, ok := events["playerEditAttachedObject"]
	if !ok {
		return true
	}

	fn, ok := evt.Handler.(func(Player, int, int, int, int, float32, float32, float32, float32, float32, float32, float32, float32, float32) bool)
	if !ok {
		return true
	}
	fn(Player{ID: int(playerid)}, int(response), int(index), int(modelid), int(boneid), float32(fOffsetX), float32(fOffsetY), float32(fOffsetZ), float32(fRotX), float32(fRotY), float32(fRotZ), float32(fScaleX), float32(fScaleY), float32(fScaleZ))
	return false
}

//export onPlayerSelectObject
func onPlayerSelectObject(playerid C.int, type_ C.int, objectid C.int, modelid C.int, fX C.float, fY C.float, fZ C.float) bool {
	evt, ok := events["playerSelectObject"]
	if !ok {
		return true
	}

	fn, ok := evt.Handler.(func(Player, int, int, int, float32, float32, float32) bool)
	if !ok {
		return true
	}
	fn(Player{ID: int(playerid)}, int(type_), int(objectid), int(modelid), float32(fX), float32(fY), float32(fZ))
	return false
}

//export onPlayerWeaponShot
func onPlayerWeaponShot(playerid C.int, weaponid C.int, hittype C.int, hitid C.int, fX C.float, fY C.float, fZ C.float) bool {
	evt, ok := events["playerWeaponShot"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(Player, int, int, int, float32, float32, float32) bool)
	if !ok {
		return false
	}
	fn(Player{ID: int(playerid)}, int(weaponid), int(hittype), int(hitid), float32(fX), float32(fY), float32(fZ))
	return true
}

//export onPlayerRequestDownload
func onPlayerRequestDownload(playerid C.int, type_ C.int, crc C.int) bool {
	evt, ok := events["playerRequestDownload"]
	if !ok {
		return true
	}

	fn, ok := evt.Handler.(func(Player, int, int) bool)
	if !ok {
		return true
	}
	fn(Player{ID: int(playerid)}, int(type_), int(crc))
	return false
}
