package sampgo

/*
#cgo CFLAGS: -I./vendor/sdk -I./vendor/gdk -I./interop -Wno-attributes -Wno-implicit-function-declaration -DHAVE_INTTYPES_H -DHAVE_MALLOC_H -DHAVE_STDINT_H

#cgo linux CFLAGS: -DLINUX
#cgo linux LDFLAGS: -ldl

#cgo windows CFLAGS: -DWIN32

#ifndef GOLANG_APP
#define GOLANG_APP

#include "interop/unitybuild.c"
#include "interop.h"

#endif
*/
import "C"

import "unsafe"

// For documentation, please visit https://open.mp/docs/scripting/functions/CreateActor
func CreateActor(modelid int, x float32, y float32, z float32, rotation float32) int {
	return int(C.CreateActor(C.int(modelid), C.float(x), C.float(y), C.float(z), C.float(rotation)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/DestroyActor
func DestroyActor(actorid int) bool {
	return bool(C.DestroyActor(C.int(actorid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/IsActorStreamedIn
func IsActorStreamedIn(actorid int, forplayerid int) bool {
	return bool(C.IsActorStreamedIn(C.int(actorid), C.int(forplayerid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SetActorVirtualWorld
func SetActorVirtualWorld(actorid int, vworld int) bool {
	return bool(C.SetActorVirtualWorld(C.int(actorid), C.int(vworld)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetActorVirtualWorld
func GetActorVirtualWorld(actorid int) int {
	return int(C.GetActorVirtualWorld(C.int(actorid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/ApplyActorAnimation
func ApplyActorAnimation(actorid int, animlib string, animname string, fDelta float32, loop bool, lockx bool, locky bool, freeze bool, time int) bool {
	csanimlib := C.CString(animlib)
	defer C.free(unsafe.Pointer(csanimlib))
	csanimname := C.CString(animname)
	defer C.free(unsafe.Pointer(csanimname))
	return bool(C.ApplyActorAnimation(C.int(actorid), C.nonConstToConst(csanimlib), C.nonConstToConst(csanimname), C.float(fDelta), C.bool(loop), C.bool(lockx), C.bool(locky), C.bool(freeze), C.int(time)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/ClearActorAnimations
func ClearActorAnimations(actorid int) bool {
	return bool(C.ClearActorAnimations(C.int(actorid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SetActorPos
func SetActorPos(actorid int, x float32, y float32, z float32) bool {
	return bool(C.SetActorPos(C.int(actorid), C.float(x), C.float(y), C.float(z)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetActorPos
func GetActorPos(actorid int, x *float32, y *float32, z *float32) bool {
	var ret bool
	var cx C.float
	var cy C.float
	var cz C.float
	ret = bool(C.GetActorPos(C.int(actorid), &cx, &cy, &cz))
	*x = float32(cx)
	*y = float32(cy)
	*z = float32(cz)
	return ret
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SetActorFacingAngle
func SetActorFacingAngle(actorid int, angle float32) bool {
	return bool(C.SetActorFacingAngle(C.int(actorid), C.float(angle)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetActorFacingAngle
func GetActorFacingAngle(actorid int, angle *float32) bool {
	var ret bool
	var cangle C.float
	ret = bool(C.GetActorFacingAngle(C.int(actorid), &cangle))
	*angle = float32(cangle)
	return ret
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SetActorHealth
func SetActorHealth(actorid int, health float32) bool {
	return bool(C.SetActorHealth(C.int(actorid), C.float(health)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetActorHealth
func GetActorHealth(actorid int, health *float32) bool {
	var ret bool
	var chealth C.float
	ret = bool(C.GetActorHealth(C.int(actorid), &chealth))
	*health = float32(chealth)
	return ret
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SetActorInvulnerable
func SetActorInvulnerable(actorid int, invulnerable bool) bool {
	return bool(C.SetActorInvulnerable(C.int(actorid), C.bool(invulnerable)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/IsActorInvulnerable
func IsActorInvulnerable(actorid int) bool {
	return bool(C.IsActorInvulnerable(C.int(actorid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/IsValidActor
func IsValidActor(actorid int) bool {
	return bool(C.IsValidActor(C.int(actorid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SetSpawnInfo
func SetSpawnInfo(playerid int, team int, skin int, x float32, y float32, z float32, rotation float32, weapon1 int, weapon1_ammo int, weapon2 int, weapon2_ammo int, weapon3 int, weapon3_ammo int) bool {
	return bool(C.SetSpawnInfo(C.int(playerid), C.int(team), C.int(skin), C.float(x), C.float(y), C.float(z), C.float(rotation), C.int(weapon1), C.int(weapon1_ammo), C.int(weapon2), C.int(weapon2_ammo), C.int(weapon3), C.int(weapon3_ammo)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SpawnPlayer
func SpawnPlayer(playerid int) bool {
	return bool(C.SpawnPlayer(C.int(playerid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SetPlayerPos
func SetPlayerPos(playerid int, x float32, y float32, z float32) bool {
	return bool(C.SetPlayerPos(C.int(playerid), C.float(x), C.float(y), C.float(z)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SetPlayerPosFindZ
func SetPlayerPosFindZ(playerid int, x float32, y float32, z float32) bool {
	return bool(C.SetPlayerPosFindZ(C.int(playerid), C.float(x), C.float(y), C.float(z)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetPlayerPos
func GetPlayerPos(playerid int, x *float32, y *float32, z *float32) bool {
	var ret bool
	var cx C.float
	var cy C.float
	var cz C.float
	ret = bool(C.GetPlayerPos(C.int(playerid), &cx, &cy, &cz))
	*x = float32(cx)
	*y = float32(cy)
	*z = float32(cz)
	return ret
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SetPlayerFacingAngle
func SetPlayerFacingAngle(playerid int, angle float32) bool {
	return bool(C.SetPlayerFacingAngle(C.int(playerid), C.float(angle)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetPlayerFacingAngle
func GetPlayerFacingAngle(playerid int, angle *float32) bool {
	var ret bool
	var cangle C.float
	ret = bool(C.GetPlayerFacingAngle(C.int(playerid), &cangle))
	*angle = float32(cangle)
	return ret
}

// For documentation, please visit https://open.mp/docs/scripting/functions/IsPlayerInRangeOfPoint
func IsPlayerInRangeOfPoint(playerid int, range_ float32, x float32, y float32, z float32) bool {
	return bool(C.IsPlayerInRangeOfPoint(C.int(playerid), C.float(range_), C.float(x), C.float(y), C.float(z)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetPlayerDistanceFromPoint
func GetPlayerDistanceFromPoint(playerid int, x float32, y float32, z float32) float32 {
	return float32(C.GetPlayerDistanceFromPoint(C.int(playerid), C.float(x), C.float(y), C.float(z)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/IsPlayerStreamedIn
func IsPlayerStreamedIn(playerid int, forplayerid int) bool {
	return bool(C.IsPlayerStreamedIn(C.int(playerid), C.int(forplayerid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SetPlayerInterior
func SetPlayerInterior(playerid int, interiorid int) bool {
	return bool(C.SetPlayerInterior(C.int(playerid), C.int(interiorid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetPlayerInterior
func GetPlayerInterior(playerid int) int {
	return int(C.GetPlayerInterior(C.int(playerid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SetPlayerHealth
func SetPlayerHealth(playerid int, health float32) bool {
	return bool(C.SetPlayerHealth(C.int(playerid), C.float(health)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetPlayerHealth
func GetPlayerHealth(playerid int, health *float32) bool {
	var ret bool
	var chealth C.float
	ret = bool(C.GetPlayerHealth(C.int(playerid), &chealth))
	*health = float32(chealth)
	return ret
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SetPlayerArmour
func SetPlayerArmour(playerid int, armour float32) bool {
	return bool(C.SetPlayerArmour(C.int(playerid), C.float(armour)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetPlayerArmour
func GetPlayerArmour(playerid int, armour *float32) bool {
	var ret bool
	var carmour C.float
	ret = bool(C.GetPlayerArmour(C.int(playerid), &carmour))
	*armour = float32(carmour)
	return ret
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SetPlayerAmmo
func SetPlayerAmmo(playerid int, weaponid int, ammo int) bool {
	return bool(C.SetPlayerAmmo(C.int(playerid), C.int(weaponid), C.int(ammo)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetPlayerAmmo
func GetPlayerAmmo(playerid int) int {
	return int(C.GetPlayerAmmo(C.int(playerid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetPlayerWeaponState
func GetPlayerWeaponState(playerid int) int {
	return int(C.GetPlayerWeaponState(C.int(playerid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetPlayerTargetPlayer
func GetPlayerTargetPlayer(playerid int) int {
	return int(C.GetPlayerTargetPlayer(C.int(playerid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetPlayerTargetActor
func GetPlayerTargetActor(playerid int) int {
	return int(C.GetPlayerTargetActor(C.int(playerid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SetPlayerTeam
func SetPlayerTeam(playerid int, teamid int) bool {
	return bool(C.SetPlayerTeam(C.int(playerid), C.int(teamid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetPlayerTeam
func GetPlayerTeam(playerid int) int {
	return int(C.GetPlayerTeam(C.int(playerid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SetPlayerScore
func SetPlayerScore(playerid int, score int) bool {
	return bool(C.SetPlayerScore(C.int(playerid), C.int(score)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetPlayerScore
func GetPlayerScore(playerid int) int {
	return int(C.GetPlayerScore(C.int(playerid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetPlayerDrunkLevel
func GetPlayerDrunkLevel(playerid int) int {
	return int(C.GetPlayerDrunkLevel(C.int(playerid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SetPlayerDrunkLevel
func SetPlayerDrunkLevel(playerid int, level int) bool {
	return bool(C.SetPlayerDrunkLevel(C.int(playerid), C.int(level)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SetPlayerColor
func SetPlayerColor(playerid int, color int) bool {
	return bool(C.SetPlayerColor(C.int(playerid), C.int(color)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetPlayerColor
func GetPlayerColor(playerid int) int {
	return int(C.GetPlayerColor(C.int(playerid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SetPlayerSkin
func SetPlayerSkin(playerid int, skinid int) bool {
	return bool(C.SetPlayerSkin(C.int(playerid), C.int(skinid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetPlayerSkin
func GetPlayerSkin(playerid int) int {
	return int(C.GetPlayerSkin(C.int(playerid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GivePlayerWeapon
func GivePlayerWeapon(playerid int, weaponid int, ammo int) bool {
	return bool(C.GivePlayerWeapon(C.int(playerid), C.int(weaponid), C.int(ammo)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/ResetPlayerWeapons
func ResetPlayerWeapons(playerid int) bool {
	return bool(C.ResetPlayerWeapons(C.int(playerid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SetPlayerArmedWeapon
func SetPlayerArmedWeapon(playerid int, weaponid int) bool {
	return bool(C.SetPlayerArmedWeapon(C.int(playerid), C.int(weaponid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetPlayerWeaponData
func GetPlayerWeaponData(playerid int, slot int, weapon *int, ammo *int) bool {
	var ret bool
	var cweapon C.int
	var cammo C.int
	ret = bool(C.GetPlayerWeaponData(C.int(playerid), C.int(slot), &cweapon, &cammo))
	*weapon = int(cweapon)
	*ammo = int(cammo)
	return ret
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GivePlayerMoney
func GivePlayerMoney(playerid int, money int) bool {
	return bool(C.GivePlayerMoney(C.int(playerid), C.int(money)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/ResetPlayerMoney
func ResetPlayerMoney(playerid int) bool {
	return bool(C.ResetPlayerMoney(C.int(playerid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SetPlayerName
func SetPlayerName(playerid int, name string) int {
	csname := C.CString(name)
	defer C.free(unsafe.Pointer(csname))
	return int(C.SetPlayerName(C.int(playerid), C.nonConstToConst(csname)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetPlayerMoney
func GetPlayerMoney(playerid int) int {
	return int(C.GetPlayerMoney(C.int(playerid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetPlayerState
func GetPlayerState(playerid int) int {
	return int(C.GetPlayerState(C.int(playerid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetPlayerIp
func GetPlayerIp(playerid int, ip *string, size int) bool {
	var ret bool
	var cip *C.char
	cip = (*C.char)(C.malloc(C.uint(size)))
	defer C.free(unsafe.Pointer(cip))
	ret = bool(C.GetPlayerIp(C.int(playerid), cip, C.int(size)))
	*ip = C.GoString(C.constToNonConst(cip))
	return ret
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetPlayerPing
func GetPlayerPing(playerid int) int {
	return int(C.GetPlayerPing(C.int(playerid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetPlayerWeapon
func GetPlayerWeapon(playerid int) int {
	return int(C.GetPlayerWeapon(C.int(playerid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetPlayerKeys
func GetPlayerKeys(playerid int, keys *int, updown *int, leftright *int) bool {
	var ret bool
	var ckeys C.int
	var cupdown C.int
	var cleftright C.int
	ret = bool(C.GetPlayerKeys(C.int(playerid), &ckeys, &cupdown, &cleftright))
	*keys = int(ckeys)
	*updown = int(cupdown)
	*leftright = int(cleftright)
	return ret
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetPlayerName
func GetPlayerName(playerid int, name *string, size int) int {
	var ret int
	var cname *C.char
	cname = (*C.char)(C.malloc(C.uint(size)))
	defer C.free(unsafe.Pointer(cname))
	ret = int(C.GetPlayerName(C.int(playerid), cname, C.int(size)))
	*name = C.GoString(C.constToNonConst(cname))
	return ret
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SetPlayerTime
func SetPlayerTime(playerid int, hour int, minute int) bool {
	return bool(C.SetPlayerTime(C.int(playerid), C.int(hour), C.int(minute)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetPlayerTime
func GetPlayerTime(playerid int, hour *int, minute *int) bool {
	var ret bool
	var chour C.int
	var cminute C.int
	ret = bool(C.GetPlayerTime(C.int(playerid), &chour, &cminute))
	*hour = int(chour)
	*minute = int(cminute)
	return ret
}

// For documentation, please visit https://open.mp/docs/scripting/functions/TogglePlayerClock
func TogglePlayerClock(playerid int, toggle bool) bool {
	return bool(C.TogglePlayerClock(C.int(playerid), C.bool(toggle)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SetPlayerWeather
func SetPlayerWeather(playerid int, weather int) bool {
	return bool(C.SetPlayerWeather(C.int(playerid), C.int(weather)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/ForceClassSelection
func ForceClassSelection(playerid int) bool {
	return bool(C.ForceClassSelection(C.int(playerid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SetPlayerWantedLevel
func SetPlayerWantedLevel(playerid int, level int) bool {
	return bool(C.SetPlayerWantedLevel(C.int(playerid), C.int(level)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetPlayerWantedLevel
func GetPlayerWantedLevel(playerid int) int {
	return int(C.GetPlayerWantedLevel(C.int(playerid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SetPlayerFightingStyle
func SetPlayerFightingStyle(playerid int, style int) bool {
	return bool(C.SetPlayerFightingStyle(C.int(playerid), C.int(style)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetPlayerFightingStyle
func GetPlayerFightingStyle(playerid int) int {
	return int(C.GetPlayerFightingStyle(C.int(playerid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SetPlayerVelocity
func SetPlayerVelocity(playerid int, x float32, y float32, z float32) bool {
	return bool(C.SetPlayerVelocity(C.int(playerid), C.float(x), C.float(y), C.float(z)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetPlayerVelocity
func GetPlayerVelocity(playerid int, x *float32, y *float32, z *float32) bool {
	var ret bool
	var cx C.float
	var cy C.float
	var cz C.float
	ret = bool(C.GetPlayerVelocity(C.int(playerid), &cx, &cy, &cz))
	*x = float32(cx)
	*y = float32(cy)
	*z = float32(cz)
	return ret
}

// For documentation, please visit https://open.mp/docs/scripting/functions/PlayCrimeReportForPlayer
func PlayCrimeReportForPlayer(playerid int, suspectid int, crime int) bool {
	return bool(C.PlayCrimeReportForPlayer(C.int(playerid), C.int(suspectid), C.int(crime)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/PlayAudioStreamForPlayer
func PlayAudioStreamForPlayer(playerid int, url string, posX float32, posY float32, posZ float32, distance float32, usepos bool) bool {
	csurl := C.CString(url)
	defer C.free(unsafe.Pointer(csurl))
	return bool(C.PlayAudioStreamForPlayer(C.int(playerid), C.nonConstToConst(csurl), C.float(posX), C.float(posY), C.float(posZ), C.float(distance), C.bool(usepos)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/StopAudioStreamForPlayer
func StopAudioStreamForPlayer(playerid int) bool {
	return bool(C.StopAudioStreamForPlayer(C.int(playerid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SetPlayerShopName
func SetPlayerShopName(playerid int, shopname string) bool {
	csshopname := C.CString(shopname)
	defer C.free(unsafe.Pointer(csshopname))
	return bool(C.SetPlayerShopName(C.int(playerid), C.nonConstToConst(csshopname)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SetPlayerSkillLevel
func SetPlayerSkillLevel(playerid int, skill int, level int) bool {
	return bool(C.SetPlayerSkillLevel(C.int(playerid), C.int(skill), C.int(level)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetPlayerSurfingVehicleID
func GetPlayerSurfingVehicleID(playerid int) int {
	return int(C.GetPlayerSurfingVehicleID(C.int(playerid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetPlayerSurfingObjectID
func GetPlayerSurfingObjectID(playerid int) int {
	return int(C.GetPlayerSurfingObjectID(C.int(playerid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/RemoveBuildingForPlayer
func RemoveBuildingForPlayer(playerid int, modelid int, fX float32, fY float32, fZ float32, fRadius float32) bool {
	return bool(C.RemoveBuildingForPlayer(C.int(playerid), C.int(modelid), C.float(fX), C.float(fY), C.float(fZ), C.float(fRadius)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetPlayerLastShotVectors
func GetPlayerLastShotVectors(playerid int, fOriginX *float32, fOriginY *float32, fOriginZ *float32, fHitPosX *float32, fHitPosY *float32, fHitPosZ *float32) bool {
	var ret bool
	var cfOriginX C.float
	var cfOriginY C.float
	var cfOriginZ C.float
	var cfHitPosX C.float
	var cfHitPosY C.float
	var cfHitPosZ C.float
	ret = bool(C.GetPlayerLastShotVectors(C.int(playerid), &cfOriginX, &cfOriginY, &cfOriginZ, &cfHitPosX, &cfHitPosY, &cfHitPosZ))
	*fOriginX = float32(cfOriginX)
	*fOriginY = float32(cfOriginY)
	*fOriginZ = float32(cfOriginZ)
	*fHitPosX = float32(cfHitPosX)
	*fHitPosY = float32(cfHitPosY)
	*fHitPosZ = float32(cfHitPosZ)
	return ret
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SetPlayerAttachedObject
func SetPlayerAttachedObject(playerid int, index int, modelid int, bone int, fOffsetX float32, fOffsetY float32, fOffsetZ float32, fRotX float32, fRotY float32, fRotZ float32, fScaleX float32, fScaleY float32, fScaleZ float32, materialcolor1 int, materialcolor2 int) bool {
	return bool(C.SetPlayerAttachedObject(C.int(playerid), C.int(index), C.int(modelid), C.int(bone), C.float(fOffsetX), C.float(fOffsetY), C.float(fOffsetZ), C.float(fRotX), C.float(fRotY), C.float(fRotZ), C.float(fScaleX), C.float(fScaleY), C.float(fScaleZ), C.int(materialcolor1), C.int(materialcolor2)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/RemovePlayerAttachedObject
func RemovePlayerAttachedObject(playerid int, index int) bool {
	return bool(C.RemovePlayerAttachedObject(C.int(playerid), C.int(index)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/IsPlayerAttachedObjectSlotUsed
func IsPlayerAttachedObjectSlotUsed(playerid int, index int) bool {
	return bool(C.IsPlayerAttachedObjectSlotUsed(C.int(playerid), C.int(index)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/EditAttachedObject
func EditAttachedObject(playerid int, index int) bool {
	return bool(C.EditAttachedObject(C.int(playerid), C.int(index)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/CreatePlayerTextDraw
func CreatePlayerTextDraw(playerid int, x float32, y float32, text string) int {
	cstext := C.CString(text)
	defer C.free(unsafe.Pointer(cstext))
	return int(C.CreatePlayerTextDraw(C.int(playerid), C.float(x), C.float(y), C.nonConstToConst(cstext)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/PlayerTextDrawDestroy
func PlayerTextDrawDestroy(playerid int, text int) bool {
	return bool(C.PlayerTextDrawDestroy(C.int(playerid), C.int(text)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/PlayerTextDrawLetterSize
func PlayerTextDrawLetterSize(playerid int, text int, x float32, y float32) bool {
	return bool(C.PlayerTextDrawLetterSize(C.int(playerid), C.int(text), C.float(x), C.float(y)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/PlayerTextDrawTextSize
func PlayerTextDrawTextSize(playerid int, text int, x float32, y float32) bool {
	return bool(C.PlayerTextDrawTextSize(C.int(playerid), C.int(text), C.float(x), C.float(y)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/PlayerTextDrawAlignment
func PlayerTextDrawAlignment(playerid int, text int, alignment int) bool {
	return bool(C.PlayerTextDrawAlignment(C.int(playerid), C.int(text), C.int(alignment)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/PlayerTextDrawColor
func PlayerTextDrawColor(playerid int, text int, color int) bool {
	return bool(C.PlayerTextDrawColor(C.int(playerid), C.int(text), C.int(color)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/PlayerTextDrawUseBox
func PlayerTextDrawUseBox(playerid int, text int, use bool) bool {
	return bool(C.PlayerTextDrawUseBox(C.int(playerid), C.int(text), C.bool(use)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/PlayerTextDrawBoxColor
func PlayerTextDrawBoxColor(playerid int, text int, color int) bool {
	return bool(C.PlayerTextDrawBoxColor(C.int(playerid), C.int(text), C.int(color)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/PlayerTextDrawSetShadow
func PlayerTextDrawSetShadow(playerid int, text int, size int) bool {
	return bool(C.PlayerTextDrawSetShadow(C.int(playerid), C.int(text), C.int(size)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/PlayerTextDrawSetOutline
func PlayerTextDrawSetOutline(playerid int, text int, size int) bool {
	return bool(C.PlayerTextDrawSetOutline(C.int(playerid), C.int(text), C.int(size)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/PlayerTextDrawBackgroundColor
func PlayerTextDrawBackgroundColor(playerid int, text int, color int) bool {
	return bool(C.PlayerTextDrawBackgroundColor(C.int(playerid), C.int(text), C.int(color)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/PlayerTextDrawFont
func PlayerTextDrawFont(playerid int, text int, font int) bool {
	return bool(C.PlayerTextDrawFont(C.int(playerid), C.int(text), C.int(font)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/PlayerTextDrawSetProportional
func PlayerTextDrawSetProportional(playerid int, text int, set bool) bool {
	return bool(C.PlayerTextDrawSetProportional(C.int(playerid), C.int(text), C.bool(set)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/PlayerTextDrawSetSelectable
func PlayerTextDrawSetSelectable(playerid int, text int, set bool) bool {
	return bool(C.PlayerTextDrawSetSelectable(C.int(playerid), C.int(text), C.bool(set)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/PlayerTextDrawShow
func PlayerTextDrawShow(playerid int, text int) bool {
	return bool(C.PlayerTextDrawShow(C.int(playerid), C.int(text)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/PlayerTextDrawHide
func PlayerTextDrawHide(playerid int, text int) bool {
	return bool(C.PlayerTextDrawHide(C.int(playerid), C.int(text)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/PlayerTextDrawSetString
func PlayerTextDrawSetString(playerid int, text int, string string) bool {
	csstring := C.CString(string)
	defer C.free(unsafe.Pointer(csstring))
	return bool(C.PlayerTextDrawSetString(C.int(playerid), C.int(text), C.nonConstToConst(csstring)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/PlayerTextDrawSetPreviewModel
func PlayerTextDrawSetPreviewModel(playerid int, text int, modelindex int) bool {
	return bool(C.PlayerTextDrawSetPreviewModel(C.int(playerid), C.int(text), C.int(modelindex)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/PlayerTextDrawSetPreviewRot
func PlayerTextDrawSetPreviewRot(playerid int, text int, fRotX float32, fRotY float32, fRotZ float32, fZoom float32) bool {
	return bool(C.PlayerTextDrawSetPreviewRot(C.int(playerid), C.int(text), C.float(fRotX), C.float(fRotY), C.float(fRotZ), C.float(fZoom)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/PlayerTextDrawSetPreviewVehCol
func PlayerTextDrawSetPreviewVehCol(playerid int, text int, color1 int, color2 int) bool {
	return bool(C.PlayerTextDrawSetPreviewVehCol(C.int(playerid), C.int(text), C.int(color1), C.int(color2)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SetPVarInt
func SetPVarInt(playerid int, varname string, value int) bool {
	csvarname := C.CString(varname)
	defer C.free(unsafe.Pointer(csvarname))
	return bool(C.SetPVarInt(C.int(playerid), C.nonConstToConst(csvarname), C.int(value)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetPVarInt
func GetPVarInt(playerid int, varname string) int {
	csvarname := C.CString(varname)
	defer C.free(unsafe.Pointer(csvarname))
	return int(C.GetPVarInt(C.int(playerid), C.nonConstToConst(csvarname)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SetPVarString
func SetPVarString(playerid int, varname string, value string) bool {
	csvarname := C.CString(varname)
	defer C.free(unsafe.Pointer(csvarname))
	csvalue := C.CString(value)
	defer C.free(unsafe.Pointer(csvalue))
	return bool(C.SetPVarString(C.int(playerid), C.nonConstToConst(csvarname), C.nonConstToConst(csvalue)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetPVarString
func GetPVarString(playerid int, varname string, value *string, size int) bool {
	csvarname := C.CString(varname)
	defer C.free(unsafe.Pointer(csvarname))
	var ret bool
	var cvalue *C.char
	cvalue = (*C.char)(C.malloc(C.uint(size)))
	defer C.free(unsafe.Pointer(cvalue))
	ret = bool(C.GetPVarString(C.int(playerid), C.nonConstToConst(csvarname), cvalue, C.int(size)))
	*value = C.GoString(C.constToNonConst(cvalue))
	return ret
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SetPVarFloat
func SetPVarFloat(playerid int, varname string, value float32) bool {
	csvarname := C.CString(varname)
	defer C.free(unsafe.Pointer(csvarname))
	return bool(C.SetPVarFloat(C.int(playerid), C.nonConstToConst(csvarname), C.float(value)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetPVarFloat
func GetPVarFloat(playerid int, varname string) float32 {
	csvarname := C.CString(varname)
	defer C.free(unsafe.Pointer(csvarname))
	return float32(C.GetPVarFloat(C.int(playerid), C.nonConstToConst(csvarname)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/DeletePVar
func DeletePVar(playerid int, varname string) bool {
	csvarname := C.CString(varname)
	defer C.free(unsafe.Pointer(csvarname))
	return bool(C.DeletePVar(C.int(playerid), C.nonConstToConst(csvarname)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetPVarsUpperIndex
func GetPVarsUpperIndex(playerid int) int {
	return int(C.GetPVarsUpperIndex(C.int(playerid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetPVarNameAtIndex
func GetPVarNameAtIndex(playerid int, index int, varname *string, size int) bool {
	var ret bool
	var cvarname *C.char
	cvarname = (*C.char)(C.malloc(C.uint(size)))
	defer C.free(unsafe.Pointer(cvarname))
	ret = bool(C.GetPVarNameAtIndex(C.int(playerid), C.int(index), cvarname, C.int(size)))
	*varname = C.GoString(C.constToNonConst(cvarname))
	return ret
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetPVarType
func GetPVarType(playerid int, varname string) int {
	csvarname := C.CString(varname)
	defer C.free(unsafe.Pointer(csvarname))
	return int(C.GetPVarType(C.int(playerid), C.nonConstToConst(csvarname)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SetPlayerChatBubble
func SetPlayerChatBubble(playerid int, text string, color int, drawdistance float32, expiretime int) bool {
	cstext := C.CString(text)
	defer C.free(unsafe.Pointer(cstext))
	return bool(C.SetPlayerChatBubble(C.int(playerid), C.nonConstToConst(cstext), C.int(color), C.float(drawdistance), C.int(expiretime)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/PutPlayerInVehicle
func PutPlayerInVehicle(playerid int, vehicleid int, seatid int) bool {
	return bool(C.PutPlayerInVehicle(C.int(playerid), C.int(vehicleid), C.int(seatid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetPlayerVehicleID
func GetPlayerVehicleID(playerid int) int {
	return int(C.GetPlayerVehicleID(C.int(playerid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetPlayerVehicleSeat
func GetPlayerVehicleSeat(playerid int) int {
	return int(C.GetPlayerVehicleSeat(C.int(playerid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/RemovePlayerFromVehicle
func RemovePlayerFromVehicle(playerid int) bool {
	return bool(C.RemovePlayerFromVehicle(C.int(playerid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/TogglePlayerControllable
func TogglePlayerControllable(playerid int, toggle bool) bool {
	return bool(C.TogglePlayerControllable(C.int(playerid), C.bool(toggle)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/PlayerPlaySound
func PlayerPlaySound(playerid int, soundid int, x float32, y float32, z float32) bool {
	return bool(C.PlayerPlaySound(C.int(playerid), C.int(soundid), C.float(x), C.float(y), C.float(z)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/ApplyAnimation
func ApplyAnimation(playerid int, animlib string, animname string, fDelta float32, loop bool, lockx bool, locky bool, freeze bool, time int, forcesync bool) bool {
	csanimlib := C.CString(animlib)
	defer C.free(unsafe.Pointer(csanimlib))
	csanimname := C.CString(animname)
	defer C.free(unsafe.Pointer(csanimname))
	return bool(C.ApplyAnimation(C.int(playerid), C.nonConstToConst(csanimlib), C.nonConstToConst(csanimname), C.float(fDelta), C.bool(loop), C.bool(lockx), C.bool(locky), C.bool(freeze), C.int(time), C.bool(forcesync)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/ClearAnimations
func ClearAnimations(playerid int, forcesync bool) bool {
	return bool(C.ClearAnimations(C.int(playerid), C.bool(forcesync)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetPlayerAnimationIndex
func GetPlayerAnimationIndex(playerid int) int {
	return int(C.GetPlayerAnimationIndex(C.int(playerid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetAnimationName
func GetAnimationName(index int, animlib *string, animlib_size int, animname *string, animname_size int) bool {
	var ret bool
	var canimlib *C.char
	canimlib = (*C.char)(C.malloc(C.uint(animlib_size)))
	defer C.free(unsafe.Pointer(canimlib))
	var canimname *C.char
	canimname = (*C.char)(C.malloc(C.uint(animname_size)))
	defer C.free(unsafe.Pointer(canimname))
	ret = bool(C.GetAnimationName(C.int(index), canimlib, C.int(animlib_size), canimname, C.int(animname_size)))
	*animlib = C.GoString(C.constToNonConst(canimlib))
	*animname = C.GoString(C.constToNonConst(canimname))
	return ret
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetPlayerSpecialAction
func GetPlayerSpecialAction(playerid int) int {
	return int(C.GetPlayerSpecialAction(C.int(playerid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SetPlayerSpecialAction
func SetPlayerSpecialAction(playerid int, actionid int) bool {
	return bool(C.SetPlayerSpecialAction(C.int(playerid), C.int(actionid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/DisableRemoteVehicleCollisions
func DisableRemoteVehicleCollisions(playerid int, disable bool) bool {
	return bool(C.DisableRemoteVehicleCollisions(C.int(playerid), C.bool(disable)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SetPlayerCheckpoint
func SetPlayerCheckpoint(playerid int, x float32, y float32, z float32, size float32) bool {
	return bool(C.SetPlayerCheckpoint(C.int(playerid), C.float(x), C.float(y), C.float(z), C.float(size)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/DisablePlayerCheckpoint
func DisablePlayerCheckpoint(playerid int) bool {
	return bool(C.DisablePlayerCheckpoint(C.int(playerid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SetPlayerRaceCheckpoint
func SetPlayerRaceCheckpoint(playerid int, type_ int, x float32, y float32, z float32, nextx float32, nexty float32, nextz float32, size float32) bool {
	return bool(C.SetPlayerRaceCheckpoint(C.int(playerid), C.int(type_), C.float(x), C.float(y), C.float(z), C.float(nextx), C.float(nexty), C.float(nextz), C.float(size)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/DisablePlayerRaceCheckpoint
func DisablePlayerRaceCheckpoint(playerid int) bool {
	return bool(C.DisablePlayerRaceCheckpoint(C.int(playerid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SetPlayerWorldBounds
func SetPlayerWorldBounds(playerid int, x_max float32, x_min float32, y_max float32, y_min float32) bool {
	return bool(C.SetPlayerWorldBounds(C.int(playerid), C.float(x_max), C.float(x_min), C.float(y_max), C.float(y_min)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SetPlayerMarkerForPlayer
func SetPlayerMarkerForPlayer(playerid int, showplayerid int, color int) bool {
	return bool(C.SetPlayerMarkerForPlayer(C.int(playerid), C.int(showplayerid), C.int(color)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/ShowPlayerNameTagForPlayer
func ShowPlayerNameTagForPlayer(playerid int, showplayerid int, show bool) bool {
	return bool(C.ShowPlayerNameTagForPlayer(C.int(playerid), C.int(showplayerid), C.bool(show)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SetPlayerMapIcon
func SetPlayerMapIcon(playerid int, iconid int, x float32, y float32, z float32, markertype int, color int, style int) bool {
	return bool(C.SetPlayerMapIcon(C.int(playerid), C.int(iconid), C.float(x), C.float(y), C.float(z), C.int(markertype), C.int(color), C.int(style)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/RemovePlayerMapIcon
func RemovePlayerMapIcon(playerid int, iconid int) bool {
	return bool(C.RemovePlayerMapIcon(C.int(playerid), C.int(iconid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/AllowPlayerTeleport
func AllowPlayerTeleport(playerid int, allow bool) bool {
	return bool(C.AllowPlayerTeleport(C.int(playerid), C.bool(allow)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SetPlayerCameraPos
func SetPlayerCameraPos(playerid int, x float32, y float32, z float32) bool {
	return bool(C.SetPlayerCameraPos(C.int(playerid), C.float(x), C.float(y), C.float(z)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SetPlayerCameraLookAt
func SetPlayerCameraLookAt(playerid int, x float32, y float32, z float32, cut int) bool {
	return bool(C.SetPlayerCameraLookAt(C.int(playerid), C.float(x), C.float(y), C.float(z), C.int(cut)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SetCameraBehindPlayer
func SetCameraBehindPlayer(playerid int) bool {
	return bool(C.SetCameraBehindPlayer(C.int(playerid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetPlayerCameraPos
func GetPlayerCameraPos(playerid int, x *float32, y *float32, z *float32) bool {
	var ret bool
	var cx C.float
	var cy C.float
	var cz C.float
	ret = bool(C.GetPlayerCameraPos(C.int(playerid), &cx, &cy, &cz))
	*x = float32(cx)
	*y = float32(cy)
	*z = float32(cz)
	return ret
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetPlayerCameraFrontVector
func GetPlayerCameraFrontVector(playerid int, x *float32, y *float32, z *float32) bool {
	var ret bool
	var cx C.float
	var cy C.float
	var cz C.float
	ret = bool(C.GetPlayerCameraFrontVector(C.int(playerid), &cx, &cy, &cz))
	*x = float32(cx)
	*y = float32(cy)
	*z = float32(cz)
	return ret
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetPlayerCameraMode
func GetPlayerCameraMode(playerid int) int {
	return int(C.GetPlayerCameraMode(C.int(playerid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/EnablePlayerCameraTarget
func EnablePlayerCameraTarget(playerid int, enable bool) bool {
	return bool(C.EnablePlayerCameraTarget(C.int(playerid), C.bool(enable)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetPlayerCameraTargetObject
func GetPlayerCameraTargetObject(playerid int) int {
	return int(C.GetPlayerCameraTargetObject(C.int(playerid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetPlayerCameraTargetVehicle
func GetPlayerCameraTargetVehicle(playerid int) int {
	return int(C.GetPlayerCameraTargetVehicle(C.int(playerid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetPlayerCameraTargetPlayer
func GetPlayerCameraTargetPlayer(playerid int) int {
	return int(C.GetPlayerCameraTargetPlayer(C.int(playerid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetPlayerCameraTargetActor
func GetPlayerCameraTargetActor(playerid int) int {
	return int(C.GetPlayerCameraTargetActor(C.int(playerid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetPlayerCameraAspectRatio
func GetPlayerCameraAspectRatio(playerid int) float32 {
	return float32(C.GetPlayerCameraAspectRatio(C.int(playerid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetPlayerCameraZoom
func GetPlayerCameraZoom(playerid int) float32 {
	return float32(C.GetPlayerCameraZoom(C.int(playerid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/AttachCameraToObject
func AttachCameraToObject(playerid int, objectid int) bool {
	return bool(C.AttachCameraToObject(C.int(playerid), C.int(objectid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/AttachCameraToPlayerObject
func AttachCameraToPlayerObject(playerid int, playerobjectid int) bool {
	return bool(C.AttachCameraToPlayerObject(C.int(playerid), C.int(playerobjectid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/InterpolateCameraPos
func InterpolateCameraPos(playerid int, FromX float32, FromY float32, FromZ float32, ToX float32, ToY float32, ToZ float32, time int, cut int) bool {
	return bool(C.InterpolateCameraPos(C.int(playerid), C.float(FromX), C.float(FromY), C.float(FromZ), C.float(ToX), C.float(ToY), C.float(ToZ), C.int(time), C.int(cut)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/InterpolateCameraLookAt
func InterpolateCameraLookAt(playerid int, FromX float32, FromY float32, FromZ float32, ToX float32, ToY float32, ToZ float32, time int, cut int) bool {
	return bool(C.InterpolateCameraLookAt(C.int(playerid), C.float(FromX), C.float(FromY), C.float(FromZ), C.float(ToX), C.float(ToY), C.float(ToZ), C.int(time), C.int(cut)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/IsPlayerConnected
func IsPlayerConnected(playerid int) bool {
	return bool(C.IsPlayerConnected(C.int(playerid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/IsPlayerInVehicle
func IsPlayerInVehicle(playerid int, vehicleid int) bool {
	return bool(C.IsPlayerInVehicle(C.int(playerid), C.int(vehicleid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/IsPlayerInAnyVehicle
func IsPlayerInAnyVehicle(playerid int) bool {
	return bool(C.IsPlayerInAnyVehicle(C.int(playerid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/IsPlayerInCheckpoint
func IsPlayerInCheckpoint(playerid int) bool {
	return bool(C.IsPlayerInCheckpoint(C.int(playerid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/IsPlayerInRaceCheckpoint
func IsPlayerInRaceCheckpoint(playerid int) bool {
	return bool(C.IsPlayerInRaceCheckpoint(C.int(playerid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SetPlayerVirtualWorld
func SetPlayerVirtualWorld(playerid int, worldid int) bool {
	return bool(C.SetPlayerVirtualWorld(C.int(playerid), C.int(worldid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetPlayerVirtualWorld
func GetPlayerVirtualWorld(playerid int) int {
	return int(C.GetPlayerVirtualWorld(C.int(playerid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/EnableStuntBonusForPlayer
func EnableStuntBonusForPlayer(playerid int, enable bool) bool {
	return bool(C.EnableStuntBonusForPlayer(C.int(playerid), C.bool(enable)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/EnableStuntBonusForAll
func EnableStuntBonusForAll(enable bool) bool {
	return bool(C.EnableStuntBonusForAll(C.bool(enable)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/TogglePlayerSpectating
func TogglePlayerSpectating(playerid int, toggle bool) bool {
	return bool(C.TogglePlayerSpectating(C.int(playerid), C.bool(toggle)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/PlayerSpectatePlayer
func PlayerSpectatePlayer(playerid int, targetplayerid int, mode int) bool {
	return bool(C.PlayerSpectatePlayer(C.int(playerid), C.int(targetplayerid), C.int(mode)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/PlayerSpectateVehicle
func PlayerSpectateVehicle(playerid int, targetvehicleid int, mode int) bool {
	return bool(C.PlayerSpectateVehicle(C.int(playerid), C.int(targetvehicleid), C.int(mode)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/StartRecordingPlayerData
func StartRecordingPlayerData(playerid int, recordtype int, recordname string) bool {
	csrecordname := C.CString(recordname)
	defer C.free(unsafe.Pointer(csrecordname))
	return bool(C.StartRecordingPlayerData(C.int(playerid), C.int(recordtype), C.nonConstToConst(csrecordname)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/StopRecordingPlayerData
func StopRecordingPlayerData(playerid int) bool {
	return bool(C.StopRecordingPlayerData(C.int(playerid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/CreateExplosionForPlayer
func CreateExplosionForPlayer(playerid int, X float32, Y float32, Z float32, type_ int, Radius float32) bool {
	return bool(C.CreateExplosionForPlayer(C.int(playerid), C.float(X), C.float(Y), C.float(Z), C.int(type_), C.float(Radius)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/IsValidVehicle
func IsValidVehicle(vehicleid int) bool {
	return bool(C.IsValidVehicle(C.int(vehicleid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetVehicleDistanceFromPoint
func GetVehicleDistanceFromPoint(vehicleid int, x float32, y float32, z float32) float32 {
	return float32(C.GetVehicleDistanceFromPoint(C.int(vehicleid), C.float(x), C.float(y), C.float(z)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/CreateVehicle
func CreateVehicle(vehicletype int, x float32, y float32, z float32, rotation float32, color1 int, color2 int, respawn_delay int, addsiren bool) int {
	return int(C.CreateVehicle(C.int(vehicletype), C.float(x), C.float(y), C.float(z), C.float(rotation), C.int(color1), C.int(color2), C.int(respawn_delay), C.bool(addsiren)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/DestroyVehicle
func DestroyVehicle(vehicleid int) bool {
	return bool(C.DestroyVehicle(C.int(vehicleid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/IsVehicleStreamedIn
func IsVehicleStreamedIn(vehicleid int, forplayerid int) bool {
	return bool(C.IsVehicleStreamedIn(C.int(vehicleid), C.int(forplayerid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetVehiclePos
func GetVehiclePos(vehicleid int, x *float32, y *float32, z *float32) bool {
	var ret bool
	var cx C.float
	var cy C.float
	var cz C.float
	ret = bool(C.GetVehiclePos(C.int(vehicleid), &cx, &cy, &cz))
	*x = float32(cx)
	*y = float32(cy)
	*z = float32(cz)
	return ret
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SetVehiclePos
func SetVehiclePos(vehicleid int, x float32, y float32, z float32) bool {
	return bool(C.SetVehiclePos(C.int(vehicleid), C.float(x), C.float(y), C.float(z)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetVehicleZAngle
func GetVehicleZAngle(vehicleid int, z_angle *float32) bool {
	var ret bool
	var cz_angle C.float
	ret = bool(C.GetVehicleZAngle(C.int(vehicleid), &cz_angle))
	*z_angle = float32(cz_angle)
	return ret
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetVehicleRotationQuat
func GetVehicleRotationQuat(vehicleid int, w *float32, x *float32, y *float32, z *float32) bool {
	var ret bool
	var cw C.float
	var cx C.float
	var cy C.float
	var cz C.float
	ret = bool(C.GetVehicleRotationQuat(C.int(vehicleid), &cw, &cx, &cy, &cz))
	*w = float32(cw)
	*x = float32(cx)
	*y = float32(cy)
	*z = float32(cz)
	return ret
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SetVehicleZAngle
func SetVehicleZAngle(vehicleid int, z_angle float32) bool {
	return bool(C.SetVehicleZAngle(C.int(vehicleid), C.float(z_angle)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SetVehicleParamsForPlayer
func SetVehicleParamsForPlayer(vehicleid int, playerid int, objective int, doorslocked int) bool {
	return bool(C.SetVehicleParamsForPlayer(C.int(vehicleid), C.int(playerid), C.int(objective), C.int(doorslocked)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/ManualVehicleEngineAndLights
func ManualVehicleEngineAndLights() bool {
	return bool(C.ManualVehicleEngineAndLights())
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SetVehicleParamsEx
func SetVehicleParamsEx(vehicleid int, engine int, lights int, alarm int, doors int, bonnet int, boot int, objective int) bool {
	return bool(C.SetVehicleParamsEx(C.int(vehicleid), C.int(engine), C.int(lights), C.int(alarm), C.int(doors), C.int(bonnet), C.int(boot), C.int(objective)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetVehicleParamsEx
func GetVehicleParamsEx(vehicleid int, engine *int, lights *int, alarm *int, doors *int, bonnet *int, boot *int, objective *int) bool {
	var ret bool
	var cengine C.int
	var clights C.int
	var calarm C.int
	var cdoors C.int
	var cbonnet C.int
	var cboot C.int
	var cobjective C.int
	ret = bool(C.GetVehicleParamsEx(C.int(vehicleid), &cengine, &clights, &calarm, &cdoors, &cbonnet, &cboot, &cobjective))
	*engine = int(cengine)
	*lights = int(clights)
	*alarm = int(calarm)
	*doors = int(cdoors)
	*bonnet = int(cbonnet)
	*boot = int(cboot)
	*objective = int(cobjective)
	return ret
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetVehicleParamsSirenState
func GetVehicleParamsSirenState(vehicleid int) int {
	return int(C.GetVehicleParamsSirenState(C.int(vehicleid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SetVehicleParamsCarDoors
func SetVehicleParamsCarDoors(vehicleid int, driver int, passenger int, backleft int, backright int) bool {
	return bool(C.SetVehicleParamsCarDoors(C.int(vehicleid), C.int(driver), C.int(passenger), C.int(backleft), C.int(backright)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetVehicleParamsCarDoors
func GetVehicleParamsCarDoors(vehicleid int, driver *int, passenger *int, backleft *int, backright *int) bool {
	var ret bool
	var cdriver C.int
	var cpassenger C.int
	var cbackleft C.int
	var cbackright C.int
	ret = bool(C.GetVehicleParamsCarDoors(C.int(vehicleid), &cdriver, &cpassenger, &cbackleft, &cbackright))
	*driver = int(cdriver)
	*passenger = int(cpassenger)
	*backleft = int(cbackleft)
	*backright = int(cbackright)
	return ret
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SetVehicleParamsCarWindows
func SetVehicleParamsCarWindows(vehicleid int, driver int, passenger int, backleft int, backright int) bool {
	return bool(C.SetVehicleParamsCarWindows(C.int(vehicleid), C.int(driver), C.int(passenger), C.int(backleft), C.int(backright)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetVehicleParamsCarWindows
func GetVehicleParamsCarWindows(vehicleid int, driver *int, passenger *int, backleft *int, backright *int) bool {
	var ret bool
	var cdriver C.int
	var cpassenger C.int
	var cbackleft C.int
	var cbackright C.int
	ret = bool(C.GetVehicleParamsCarWindows(C.int(vehicleid), &cdriver, &cpassenger, &cbackleft, &cbackright))
	*driver = int(cdriver)
	*passenger = int(cpassenger)
	*backleft = int(cbackleft)
	*backright = int(cbackright)
	return ret
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SetVehicleToRespawn
func SetVehicleToRespawn(vehicleid int) bool {
	return bool(C.SetVehicleToRespawn(C.int(vehicleid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/LinkVehicleToInterior
func LinkVehicleToInterior(vehicleid int, interiorid int) bool {
	return bool(C.LinkVehicleToInterior(C.int(vehicleid), C.int(interiorid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/AddVehicleComponent
func AddVehicleComponent(vehicleid int, componentid int) bool {
	return bool(C.AddVehicleComponent(C.int(vehicleid), C.int(componentid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/RemoveVehicleComponent
func RemoveVehicleComponent(vehicleid int, componentid int) bool {
	return bool(C.RemoveVehicleComponent(C.int(vehicleid), C.int(componentid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/ChangeVehicleColor
func ChangeVehicleColor(vehicleid int, color1 int, color2 int) bool {
	return bool(C.ChangeVehicleColor(C.int(vehicleid), C.int(color1), C.int(color2)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/ChangeVehiclePaintjob
func ChangeVehiclePaintjob(vehicleid int, paintjobid int) bool {
	return bool(C.ChangeVehiclePaintjob(C.int(vehicleid), C.int(paintjobid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SetVehicleHealth
func SetVehicleHealth(vehicleid int, health float32) bool {
	return bool(C.SetVehicleHealth(C.int(vehicleid), C.float(health)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetVehicleHealth
func GetVehicleHealth(vehicleid int, health *float32) bool {
	var ret bool
	var chealth C.float
	ret = bool(C.GetVehicleHealth(C.int(vehicleid), &chealth))
	*health = float32(chealth)
	return ret
}

// For documentation, please visit https://open.mp/docs/scripting/functions/AttachTrailerToVehicle
func AttachTrailerToVehicle(trailerid int, vehicleid int) bool {
	return bool(C.AttachTrailerToVehicle(C.int(trailerid), C.int(vehicleid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/DetachTrailerFromVehicle
func DetachTrailerFromVehicle(vehicleid int) bool {
	return bool(C.DetachTrailerFromVehicle(C.int(vehicleid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/IsTrailerAttachedToVehicle
func IsTrailerAttachedToVehicle(vehicleid int) bool {
	return bool(C.IsTrailerAttachedToVehicle(C.int(vehicleid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetVehicleTrailer
func GetVehicleTrailer(vehicleid int) int {
	return int(C.GetVehicleTrailer(C.int(vehicleid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SetVehicleNumberPlate
func SetVehicleNumberPlate(vehicleid int, numberplate string) bool {
	csnumberplate := C.CString(numberplate)
	defer C.free(unsafe.Pointer(csnumberplate))
	return bool(C.SetVehicleNumberPlate(C.int(vehicleid), C.nonConstToConst(csnumberplate)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetVehicleModel
func GetVehicleModel(vehicleid int) int {
	return int(C.GetVehicleModel(C.int(vehicleid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetVehicleComponentInSlot
func GetVehicleComponentInSlot(vehicleid int, slot int) int {
	return int(C.GetVehicleComponentInSlot(C.int(vehicleid), C.int(slot)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetVehicleComponentType
func GetVehicleComponentType(component int) int {
	return int(C.GetVehicleComponentType(C.int(component)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/RepairVehicle
func RepairVehicle(vehicleid int) bool {
	return bool(C.RepairVehicle(C.int(vehicleid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetVehicleVelocity
func GetVehicleVelocity(vehicleid int, X *float32, Y *float32, Z *float32) bool {
	var ret bool
	var cX C.float
	var cY C.float
	var cZ C.float
	ret = bool(C.GetVehicleVelocity(C.int(vehicleid), &cX, &cY, &cZ))
	*X = float32(cX)
	*Y = float32(cY)
	*Z = float32(cZ)
	return ret
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SetVehicleVelocity
func SetVehicleVelocity(vehicleid int, X float32, Y float32, Z float32) bool {
	return bool(C.SetVehicleVelocity(C.int(vehicleid), C.float(X), C.float(Y), C.float(Z)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SetVehicleAngularVelocity
func SetVehicleAngularVelocity(vehicleid int, X float32, Y float32, Z float32) bool {
	return bool(C.SetVehicleAngularVelocity(C.int(vehicleid), C.float(X), C.float(Y), C.float(Z)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetVehicleDamageStatus
func GetVehicleDamageStatus(vehicleid int, panels *int, doors *int, lights *int, tires *int) bool {
	var ret bool
	var cpanels C.int
	var cdoors C.int
	var clights C.int
	var ctires C.int
	ret = bool(C.GetVehicleDamageStatus(C.int(vehicleid), &cpanels, &cdoors, &clights, &ctires))
	*panels = int(cpanels)
	*doors = int(cdoors)
	*lights = int(clights)
	*tires = int(ctires)
	return ret
}

// For documentation, please visit https://open.mp/docs/scripting/functions/UpdateVehicleDamageStatus
func UpdateVehicleDamageStatus(vehicleid int, panels int, doors int, lights int, tires int) bool {
	return bool(C.UpdateVehicleDamageStatus(C.int(vehicleid), C.int(panels), C.int(doors), C.int(lights), C.int(tires)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SetVehicleVirtualWorld
func SetVehicleVirtualWorld(vehicleid int, worldid int) bool {
	return bool(C.SetVehicleVirtualWorld(C.int(vehicleid), C.int(worldid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetVehicleVirtualWorld
func GetVehicleVirtualWorld(vehicleid int) int {
	return int(C.GetVehicleVirtualWorld(C.int(vehicleid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetVehicleModelInfo
func GetVehicleModelInfo(model int, infotype int, X *float32, Y *float32, Z *float32) bool {
	var ret bool
	var cX C.float
	var cY C.float
	var cZ C.float
	ret = bool(C.GetVehicleModelInfo(C.int(model), C.int(infotype), &cX, &cY, &cZ))
	*X = float32(cX)
	*Y = float32(cY)
	*Z = float32(cZ)
	return ret
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SendClientMessage
func SendClientMessage(playerid int, color int, message string) bool {
	csmessage := C.CString(message)
	defer C.free(unsafe.Pointer(csmessage))
	return bool(C.SendClientMessage(C.int(playerid), C.int(color), C.nonConstToConst(csmessage)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SendClientMessageToAll
func SendClientMessageToAll(color int, message string) bool {
	csmessage := C.CString(message)
	defer C.free(unsafe.Pointer(csmessage))
	return bool(C.SendClientMessageToAll(C.int(color), C.nonConstToConst(csmessage)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SendPlayerMessageToPlayer
func SendPlayerMessageToPlayer(playerid int, senderid int, message string) bool {
	csmessage := C.CString(message)
	defer C.free(unsafe.Pointer(csmessage))
	return bool(C.SendPlayerMessageToPlayer(C.int(playerid), C.int(senderid), C.nonConstToConst(csmessage)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SendPlayerMessageToAll
func SendPlayerMessageToAll(senderid int, message string) bool {
	csmessage := C.CString(message)
	defer C.free(unsafe.Pointer(csmessage))
	return bool(C.SendPlayerMessageToAll(C.int(senderid), C.nonConstToConst(csmessage)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SendDeathMessage
func SendDeathMessage(killer int, killee int, weapon int) bool {
	return bool(C.SendDeathMessage(C.int(killer), C.int(killee), C.int(weapon)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SendDeathMessageToPlayer
func SendDeathMessageToPlayer(playerid int, killer int, killee int, weapon int) bool {
	return bool(C.SendDeathMessageToPlayer(C.int(playerid), C.int(killer), C.int(killee), C.int(weapon)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GameTextForAll
func GameTextForAll(text string, time int, style int) bool {
	cstext := C.CString(text)
	defer C.free(unsafe.Pointer(cstext))
	return bool(C.GameTextForAll(C.nonConstToConst(cstext), C.int(time), C.int(style)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GameTextForPlayer
func GameTextForPlayer(playerid int, text string, time int, style int) bool {
	cstext := C.CString(text)
	defer C.free(unsafe.Pointer(cstext))
	return bool(C.GameTextForPlayer(C.int(playerid), C.nonConstToConst(cstext), C.int(time), C.int(style)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetTickCount
func GetTickCount() int {
	return int(C.sampgdk_GetTickCount())
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetMaxPlayers
func GetMaxPlayers() int {
	return int(C.GetMaxPlayers())
}

// For documentation, please visit https://open.mp/docs/scripting/functions/VectorSize
func VectorSize(x float32, y float32, z float32) float32 {
	return float32(C.VectorSize(C.float(x), C.float(y), C.float(z)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetPlayerPoolSize
func GetPlayerPoolSize() int {
	return int(C.GetPlayerPoolSize())
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetVehiclePoolSize
func GetVehiclePoolSize() int {
	return int(C.GetVehiclePoolSize())
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetActorPoolSize
func GetActorPoolSize() int {
	return int(C.GetActorPoolSize())
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SHA256_PassHash
func SHA256_PassHash(password string, salt string, ret_hash *string, ret_hash_len int) bool {
	cspassword := C.CString(password)
	defer C.free(unsafe.Pointer(cspassword))
	cssalt := C.CString(salt)
	defer C.free(unsafe.Pointer(cssalt))
	var ret bool
	var cret_hash *C.char
	cret_hash = (*C.char)(C.malloc(C.uint(ret_hash_len)))
	defer C.free(unsafe.Pointer(cret_hash))
	ret = bool(C.SHA256_PassHash(C.nonConstToConst(cspassword), C.nonConstToConst(cssalt), cret_hash, C.int(ret_hash_len)))
	*ret_hash = C.GoString(C.constToNonConst(cret_hash))
	return ret
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SetSVarInt
func SetSVarInt(varname string, int_value int) bool {
	csvarname := C.CString(varname)
	defer C.free(unsafe.Pointer(csvarname))
	return bool(C.SetSVarInt(C.nonConstToConst(csvarname), C.int(int_value)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetSVarInt
func GetSVarInt(varname string) int {
	csvarname := C.CString(varname)
	defer C.free(unsafe.Pointer(csvarname))
	return int(C.GetSVarInt(C.nonConstToConst(csvarname)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SetSVarString
func SetSVarString(varname string, string_value string) bool {
	csvarname := C.CString(varname)
	defer C.free(unsafe.Pointer(csvarname))
	csstring_value := C.CString(string_value)
	defer C.free(unsafe.Pointer(csstring_value))
	return bool(C.SetSVarString(C.nonConstToConst(csvarname), C.nonConstToConst(csstring_value)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetSVarString
func GetSVarString(varname string, string_return *string, len_ int) bool {
	csvarname := C.CString(varname)
	defer C.free(unsafe.Pointer(csvarname))
	var ret bool
	var cstring_return *C.char
	cstring_return = (*C.char)(C.malloc(C.uint(len_)))
	defer C.free(unsafe.Pointer(cstring_return))
	ret = bool(C.GetSVarString(C.nonConstToConst(csvarname), cstring_return, C.int(len_)))
	*string_return = C.GoString(C.constToNonConst(cstring_return))
	return ret
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SetSVarFloat
func SetSVarFloat(varname string, float_value float32) bool {
	csvarname := C.CString(varname)
	defer C.free(unsafe.Pointer(csvarname))
	return bool(C.SetSVarFloat(C.nonConstToConst(csvarname), C.float(float_value)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetSVarFloat
func GetSVarFloat(varname string) float32 {
	csvarname := C.CString(varname)
	defer C.free(unsafe.Pointer(csvarname))
	return float32(C.GetSVarFloat(C.nonConstToConst(csvarname)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/DeleteSVar
func DeleteSVar(varname string) bool {
	csvarname := C.CString(varname)
	defer C.free(unsafe.Pointer(csvarname))
	return bool(C.DeleteSVar(C.nonConstToConst(csvarname)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetSVarsUpperIndex
func GetSVarsUpperIndex() int {
	return int(C.GetSVarsUpperIndex())
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetSVarNameAtIndex
func GetSVarNameAtIndex(index int, ret_varname *string, ret_len int) bool {
	var ret bool
	var cret_varname *C.char
	cret_varname = (*C.char)(C.malloc(C.uint(ret_len)))
	defer C.free(unsafe.Pointer(cret_varname))
	ret = bool(C.GetSVarNameAtIndex(C.int(index), cret_varname, C.int(ret_len)))
	*ret_varname = C.GoString(C.constToNonConst(cret_varname))
	return ret
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetSVarType
func GetSVarType(varname string) int {
	csvarname := C.CString(varname)
	defer C.free(unsafe.Pointer(csvarname))
	return int(C.GetSVarType(C.nonConstToConst(csvarname)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SetGameModeText
func SetGameModeText(text string) bool {
	cstext := C.CString(text)
	defer C.free(unsafe.Pointer(cstext))
	return bool(C.SetGameModeText(C.nonConstToConst(cstext)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SetTeamCount
func SetTeamCount(count int) bool {
	return bool(C.SetTeamCount(C.int(count)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/AddPlayerClass
func AddPlayerClass(modelid int, spawn_x float32, spawn_y float32, spawn_z float32, z_angle float32, weapon1 int, weapon1_ammo int, weapon2 int, weapon2_ammo int, weapon3 int, weapon3_ammo int) int {
	return int(C.AddPlayerClass(C.int(modelid), C.float(spawn_x), C.float(spawn_y), C.float(spawn_z), C.float(z_angle), C.int(weapon1), C.int(weapon1_ammo), C.int(weapon2), C.int(weapon2_ammo), C.int(weapon3), C.int(weapon3_ammo)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/AddPlayerClassEx
func AddPlayerClassEx(teamid int, modelid int, spawn_x float32, spawn_y float32, spawn_z float32, z_angle float32, weapon1 int, weapon1_ammo int, weapon2 int, weapon2_ammo int, weapon3 int, weapon3_ammo int) int {
	return int(C.AddPlayerClassEx(C.int(teamid), C.int(modelid), C.float(spawn_x), C.float(spawn_y), C.float(spawn_z), C.float(z_angle), C.int(weapon1), C.int(weapon1_ammo), C.int(weapon2), C.int(weapon2_ammo), C.int(weapon3), C.int(weapon3_ammo)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/AddStaticVehicle
func AddStaticVehicle(modelid int, spawn_x float32, spawn_y float32, spawn_z float32, z_angle float32, color1 int, color2 int) int {
	return int(C.AddStaticVehicle(C.int(modelid), C.float(spawn_x), C.float(spawn_y), C.float(spawn_z), C.float(z_angle), C.int(color1), C.int(color2)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/AddStaticVehicleEx
func AddStaticVehicleEx(modelid int, spawn_x float32, spawn_y float32, spawn_z float32, z_angle float32, color1 int, color2 int, respawn_delay int, addsiren bool) int {
	return int(C.AddStaticVehicleEx(C.int(modelid), C.float(spawn_x), C.float(spawn_y), C.float(spawn_z), C.float(z_angle), C.int(color1), C.int(color2), C.int(respawn_delay), C.bool(addsiren)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/AddStaticPickup
func AddStaticPickup(model int, type_ int, x float32, y float32, z float32, virtualworld int) int {
	return int(C.AddStaticPickup(C.int(model), C.int(type_), C.float(x), C.float(y), C.float(z), C.int(virtualworld)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/CreatePickup
func CreatePickup(model int, type_ int, x float32, y float32, z float32, virtualworld int) int {
	return int(C.CreatePickup(C.int(model), C.int(type_), C.float(x), C.float(y), C.float(z), C.int(virtualworld)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/DestroyPickup
func DestroyPickup(pickup int) bool {
	return bool(C.DestroyPickup(C.int(pickup)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/ShowNameTags
func ShowNameTags(show bool) bool {
	return bool(C.ShowNameTags(C.bool(show)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/ShowPlayerMarkers
func ShowPlayerMarkers(mode int) bool {
	return bool(C.ShowPlayerMarkers(C.int(mode)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GameModeExit
func GameModeExit() bool {
	return bool(C.GameModeExit())
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SetWorldTime
func SetWorldTime(hour int) bool {
	return bool(C.SetWorldTime(C.int(hour)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetWeaponName
func GetWeaponName(weaponid int, name *string, size int) bool {
	var ret bool
	var cname *C.char
	cname = (*C.char)(C.malloc(C.uint(size)))
	defer C.free(unsafe.Pointer(cname))
	ret = bool(C.GetWeaponName(C.int(weaponid), cname, C.int(size)))
	*name = C.GoString(C.constToNonConst(cname))
	return ret
}

// For documentation, please visit https://open.mp/docs/scripting/functions/EnableTirePopping
func EnableTirePopping(enable bool) bool {
	return bool(C.EnableTirePopping(C.bool(enable)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/EnableVehicleFriendlyFire
func EnableVehicleFriendlyFire() bool {
	return bool(C.EnableVehicleFriendlyFire())
}

// For documentation, please visit https://open.mp/docs/scripting/functions/AllowInteriorWeapons
func AllowInteriorWeapons(allow bool) bool {
	return bool(C.AllowInteriorWeapons(C.bool(allow)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SetWeather
func SetWeather(weatherid int) bool {
	return bool(C.SetWeather(C.int(weatherid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SetGravity
func SetGravity(gravity float32) bool {
	return bool(C.SetGravity(C.float(gravity)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetGravity
func GetGravity() float32 {
	return float32(C.GetGravity())
}

// For documentation, please visit https://open.mp/docs/scripting/functions/AllowAdminTeleport
func AllowAdminTeleport(allow bool) bool {
	return bool(C.AllowAdminTeleport(C.bool(allow)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SetDeathDropAmount
func SetDeathDropAmount(amount int) bool {
	return bool(C.SetDeathDropAmount(C.int(amount)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/CreateExplosion
func CreateExplosion(x float32, y float32, z float32, type_ int, radius float32) bool {
	return bool(C.CreateExplosion(C.float(x), C.float(y), C.float(z), C.int(type_), C.float(radius)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/EnableZoneNames
func EnableZoneNames(enable bool) bool {
	return bool(C.EnableZoneNames(C.bool(enable)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/UsePlayerPedAnims
func UsePlayerPedAnims() bool {
	return bool(C.UsePlayerPedAnims())
}

// For documentation, please visit https://open.mp/docs/scripting/functions/DisableInteriorEnterExits
func DisableInteriorEnterExits() bool {
	return bool(C.DisableInteriorEnterExits())
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SetNameTagDrawDistance
func SetNameTagDrawDistance(distance float32) bool {
	return bool(C.SetNameTagDrawDistance(C.float(distance)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/DisableNameTagLOS
func DisableNameTagLOS() bool {
	return bool(C.DisableNameTagLOS())
}

// For documentation, please visit https://open.mp/docs/scripting/functions/LimitGlobalChatRadius
func LimitGlobalChatRadius(chat_radius float32) bool {
	return bool(C.LimitGlobalChatRadius(C.float(chat_radius)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/LimitPlayerMarkerRadius
func LimitPlayerMarkerRadius(marker_radius float32) bool {
	return bool(C.LimitPlayerMarkerRadius(C.float(marker_radius)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/ConnectNPC
func ConnectNPC(name string, script string) bool {
	csname := C.CString(name)
	defer C.free(unsafe.Pointer(csname))
	csscript := C.CString(script)
	defer C.free(unsafe.Pointer(csscript))
	return bool(C.ConnectNPC(C.nonConstToConst(csname), C.nonConstToConst(csscript)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/IsPlayerNPC
func IsPlayerNPC(playerid int) bool {
	return bool(C.IsPlayerNPC(C.int(playerid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/IsPlayerAdmin
func IsPlayerAdmin(playerid int) bool {
	return bool(C.IsPlayerAdmin(C.int(playerid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/Kick
func Kick(playerid int) bool {
	return bool(C.Kick(C.int(playerid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/Ban
func Ban(playerid int) bool {
	return bool(C.Ban(C.int(playerid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/BanEx
func BanEx(playerid int, reason string) bool {
	csreason := C.CString(reason)
	defer C.free(unsafe.Pointer(csreason))
	return bool(C.BanEx(C.int(playerid), C.nonConstToConst(csreason)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SendRconCommand
func SendRconCommand(command string) bool {
	cscommand := C.CString(command)
	defer C.free(unsafe.Pointer(cscommand))
	return bool(C.SendRconCommand(C.nonConstToConst(cscommand)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetPlayerNetworkStats
func GetPlayerNetworkStats(playerid int, retstr *string, size int) bool {
	var ret bool
	var cretstr *C.char
	cretstr = (*C.char)(C.malloc(C.uint(size)))
	defer C.free(unsafe.Pointer(cretstr))
	ret = bool(C.GetPlayerNetworkStats(C.int(playerid), cretstr, C.int(size)))
	*retstr = C.GoString(C.constToNonConst(cretstr))
	return ret
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetNetworkStats
func GetNetworkStats(retstr *string, size int) bool {
	var ret bool
	var cretstr *C.char
	cretstr = (*C.char)(C.malloc(C.uint(size)))
	defer C.free(unsafe.Pointer(cretstr))
	ret = bool(C.GetNetworkStats(cretstr, C.int(size)))
	*retstr = C.GoString(C.constToNonConst(cretstr))
	return ret
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetPlayerVersion
func GetPlayerVersion(playerid int, version *string, len_ int) bool {
	var ret bool
	var cversion *C.char
	cversion = (*C.char)(C.malloc(C.uint(len_)))
	defer C.free(unsafe.Pointer(cversion))
	ret = bool(C.GetPlayerVersion(C.int(playerid), cversion, C.int(len_)))
	*version = C.GoString(C.constToNonConst(cversion))
	return ret
}

// For documentation, please visit https://open.mp/docs/scripting/functions/BlockIpAddress
func BlockIpAddress(ip_address string, timems int) bool {
	csip_address := C.CString(ip_address)
	defer C.free(unsafe.Pointer(csip_address))
	return bool(C.BlockIpAddress(C.nonConstToConst(csip_address), C.int(timems)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/UnBlockIpAddress
func UnBlockIpAddress(ip_address string) bool {
	csip_address := C.CString(ip_address)
	defer C.free(unsafe.Pointer(csip_address))
	return bool(C.UnBlockIpAddress(C.nonConstToConst(csip_address)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetServerVarAsString
func GetServerVarAsString(varname string, value *string, size int) bool {
	csvarname := C.CString(varname)
	defer C.free(unsafe.Pointer(csvarname))
	var ret bool
	var cvalue *C.char
	cvalue = (*C.char)(C.malloc(C.uint(size)))
	defer C.free(unsafe.Pointer(cvalue))
	ret = bool(C.GetServerVarAsString(C.nonConstToConst(csvarname), cvalue, C.int(size)))
	*value = C.GoString(C.constToNonConst(cvalue))
	return ret
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetServerVarAsInt
func GetServerVarAsInt(varname string) int {
	csvarname := C.CString(varname)
	defer C.free(unsafe.Pointer(csvarname))
	return int(C.GetServerVarAsInt(C.nonConstToConst(csvarname)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetServerVarAsBool
func GetServerVarAsBool(varname string) bool {
	csvarname := C.CString(varname)
	defer C.free(unsafe.Pointer(csvarname))
	return bool(C.GetServerVarAsBool(C.nonConstToConst(csvarname)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetConsoleVarAsString
func GetConsoleVarAsString(varname string, buffer *string, len_ int) bool {
	csvarname := C.CString(varname)
	defer C.free(unsafe.Pointer(csvarname))
	var ret bool
	var cbuffer *C.char
	cbuffer = (*C.char)(C.malloc(C.uint(len_)))
	defer C.free(unsafe.Pointer(cbuffer))
	ret = bool(C.GetConsoleVarAsString(C.nonConstToConst(csvarname), cbuffer, C.int(len_)))
	*buffer = C.GoString(C.constToNonConst(cbuffer))
	return ret
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetConsoleVarAsInt
func GetConsoleVarAsInt(varname string) int {
	csvarname := C.CString(varname)
	defer C.free(unsafe.Pointer(csvarname))
	return int(C.GetConsoleVarAsInt(C.nonConstToConst(csvarname)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetConsoleVarAsBool
func GetConsoleVarAsBool(varname string) bool {
	csvarname := C.CString(varname)
	defer C.free(unsafe.Pointer(csvarname))
	return bool(C.GetConsoleVarAsBool(C.nonConstToConst(csvarname)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetServerTickRate
func GetServerTickRate() int {
	return int(C.GetServerTickRate())
}

// For documentation, please visit https://open.mp/docs/scripting/functions/NetStats_GetConnectedTime
func NetStats_GetConnectedTime(playerid int) int {
	return int(C.NetStats_GetConnectedTime(C.int(playerid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/NetStats_MessagesReceived
func NetStats_MessagesReceived(playerid int) int {
	return int(C.NetStats_MessagesReceived(C.int(playerid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/NetStats_BytesReceived
func NetStats_BytesReceived(playerid int) int {
	return int(C.NetStats_BytesReceived(C.int(playerid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/NetStats_MessagesSent
func NetStats_MessagesSent(playerid int) int {
	return int(C.NetStats_MessagesSent(C.int(playerid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/NetStats_BytesSent
func NetStats_BytesSent(playerid int) int {
	return int(C.NetStats_BytesSent(C.int(playerid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/NetStats_MessagesRecvPerSecond
func NetStats_MessagesRecvPerSecond(playerid int) int {
	return int(C.NetStats_MessagesRecvPerSecond(C.int(playerid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/NetStats_PacketLossPercent
func NetStats_PacketLossPercent(playerid int) float32 {
	return float32(C.NetStats_PacketLossPercent(C.int(playerid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/NetStats_ConnectionStatus
func NetStats_ConnectionStatus(playerid int) int {
	return int(C.NetStats_ConnectionStatus(C.int(playerid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/NetStats_GetIpPort
func NetStats_GetIpPort(playerid int, ip_port *string, ip_port_len int) bool {
	var ret bool
	var cip_port *C.char
	cip_port = (*C.char)(C.malloc(C.uint(ip_port_len)))
	defer C.free(unsafe.Pointer(cip_port))
	ret = bool(C.NetStats_GetIpPort(C.int(playerid), cip_port, C.int(ip_port_len)))
	*ip_port = C.GoString(C.constToNonConst(cip_port))
	return ret
}

// For documentation, please visit https://open.mp/docs/scripting/functions/CreateMenu
func CreateMenu(title string, columns int, x float32, y float32, col1width float32, col2width float32) int {
	cstitle := C.CString(title)
	defer C.free(unsafe.Pointer(cstitle))
	return int(C.sampgdk_CreateMenu(C.nonConstToConst(cstitle), C.int(columns), C.float(x), C.float(y), C.float(col1width), C.float(col2width)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/DestroyMenu
func DestroyMenu(menuid int) bool {
	return bool(C.sampgdk_DestroyMenu(C.int(menuid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/AddMenuItem
func AddMenuItem(menuid int, column int, menutext string) int {
	csmenutext := C.CString(menutext)
	defer C.free(unsafe.Pointer(csmenutext))
	return int(C.AddMenuItem(C.int(menuid), C.int(column), C.nonConstToConst(csmenutext)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SetMenuColumnHeader
func SetMenuColumnHeader(menuid int, column int, columnheader string) bool {
	cscolumnheader := C.CString(columnheader)
	defer C.free(unsafe.Pointer(cscolumnheader))
	return bool(C.SetMenuColumnHeader(C.int(menuid), C.int(column), C.nonConstToConst(cscolumnheader)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/ShowMenuForPlayer
func ShowMenuForPlayer(menuid int, playerid int) bool {
	return bool(C.ShowMenuForPlayer(C.int(menuid), C.int(playerid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/HideMenuForPlayer
func HideMenuForPlayer(menuid int, playerid int) bool {
	return bool(C.HideMenuForPlayer(C.int(menuid), C.int(playerid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/IsValidMenu
func IsValidMenu(menuid int) bool {
	return bool(C.IsValidMenu(C.int(menuid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/DisableMenu
func DisableMenu(menuid int) bool {
	return bool(C.DisableMenu(C.int(menuid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/DisableMenuRow
func DisableMenuRow(menuid int, row int) bool {
	return bool(C.DisableMenuRow(C.int(menuid), C.int(row)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetPlayerMenu
func GetPlayerMenu(playerid int) int {
	return int(C.GetPlayerMenu(C.int(playerid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/TextDrawCreate
func TextDrawCreate(x float32, y float32, text string) int {
	cstext := C.CString(text)
	defer C.free(unsafe.Pointer(cstext))
	return int(C.TextDrawCreate(C.float(x), C.float(y), C.nonConstToConst(cstext)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/TextDrawDestroy
func TextDrawDestroy(text int) bool {
	return bool(C.TextDrawDestroy(C.int(text)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/TextDrawLetterSize
func TextDrawLetterSize(text int, x float32, y float32) bool {
	return bool(C.TextDrawLetterSize(C.int(text), C.float(x), C.float(y)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/TextDrawTextSize
func TextDrawTextSize(text int, x float32, y float32) bool {
	return bool(C.TextDrawTextSize(C.int(text), C.float(x), C.float(y)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/TextDrawAlignment
func TextDrawAlignment(text int, alignment int) bool {
	return bool(C.TextDrawAlignment(C.int(text), C.int(alignment)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/TextDrawColor
func TextDrawColor(text int, color int) bool {
	return bool(C.TextDrawColor(C.int(text), C.int(color)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/TextDrawUseBox
func TextDrawUseBox(text int, use bool) bool {
	return bool(C.TextDrawUseBox(C.int(text), C.bool(use)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/TextDrawBoxColor
func TextDrawBoxColor(text int, color int) bool {
	return bool(C.TextDrawBoxColor(C.int(text), C.int(color)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/TextDrawSetShadow
func TextDrawSetShadow(text int, size int) bool {
	return bool(C.TextDrawSetShadow(C.int(text), C.int(size)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/TextDrawSetOutline
func TextDrawSetOutline(text int, size int) bool {
	return bool(C.TextDrawSetOutline(C.int(text), C.int(size)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/TextDrawBackgroundColor
func TextDrawBackgroundColor(text int, color int) bool {
	return bool(C.TextDrawBackgroundColor(C.int(text), C.int(color)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/TextDrawFont
func TextDrawFont(text int, font int) bool {
	return bool(C.TextDrawFont(C.int(text), C.int(font)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/TextDrawSetProportional
func TextDrawSetProportional(text int, set bool) bool {
	return bool(C.TextDrawSetProportional(C.int(text), C.bool(set)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/TextDrawSetSelectable
func TextDrawSetSelectable(text int, set bool) bool {
	return bool(C.TextDrawSetSelectable(C.int(text), C.bool(set)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/TextDrawShowForPlayer
func TextDrawShowForPlayer(playerid int, text int) bool {
	return bool(C.TextDrawShowForPlayer(C.int(playerid), C.int(text)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/TextDrawHideForPlayer
func TextDrawHideForPlayer(playerid int, text int) bool {
	return bool(C.TextDrawHideForPlayer(C.int(playerid), C.int(text)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/TextDrawShowForAll
func TextDrawShowForAll(text int) bool {
	return bool(C.TextDrawShowForAll(C.int(text)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/TextDrawHideForAll
func TextDrawHideForAll(text int) bool {
	return bool(C.TextDrawHideForAll(C.int(text)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/TextDrawSetString
func TextDrawSetString(text int, string string) bool {
	csstring := C.CString(string)
	defer C.free(unsafe.Pointer(csstring))
	return bool(C.TextDrawSetString(C.int(text), C.nonConstToConst(csstring)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/TextDrawSetPreviewModel
func TextDrawSetPreviewModel(text int, modelindex int) bool {
	return bool(C.TextDrawSetPreviewModel(C.int(text), C.int(modelindex)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/TextDrawSetPreviewRot
func TextDrawSetPreviewRot(text int, fRotX float32, fRotY float32, fRotZ float32, fZoom float32) bool {
	return bool(C.TextDrawSetPreviewRot(C.int(text), C.float(fRotX), C.float(fRotY), C.float(fRotZ), C.float(fZoom)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/TextDrawSetPreviewVehCol
func TextDrawSetPreviewVehCol(text int, color1 int, color2 int) bool {
	return bool(C.TextDrawSetPreviewVehCol(C.int(text), C.int(color1), C.int(color2)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SelectTextDraw
func SelectTextDraw(playerid int, hovercolor int) bool {
	return bool(C.SelectTextDraw(C.int(playerid), C.int(hovercolor)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/CancelSelectTextDraw
func CancelSelectTextDraw(playerid int) bool {
	return bool(C.CancelSelectTextDraw(C.int(playerid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GangZoneCreate
func GangZoneCreate(minx float32, miny float32, maxx float32, maxy float32) int {
	return int(C.GangZoneCreate(C.float(minx), C.float(miny), C.float(maxx), C.float(maxy)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GangZoneDestroy
func GangZoneDestroy(zone int) bool {
	return bool(C.GangZoneDestroy(C.int(zone)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GangZoneShowForPlayer
func GangZoneShowForPlayer(playerid int, zone int, color int) bool {
	return bool(C.GangZoneShowForPlayer(C.int(playerid), C.int(zone), C.int(color)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GangZoneShowForAll
func GangZoneShowForAll(zone int, color int) bool {
	return bool(C.GangZoneShowForAll(C.int(zone), C.int(color)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GangZoneHideForPlayer
func GangZoneHideForPlayer(playerid int, zone int) bool {
	return bool(C.GangZoneHideForPlayer(C.int(playerid), C.int(zone)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GangZoneHideForAll
func GangZoneHideForAll(zone int) bool {
	return bool(C.GangZoneHideForAll(C.int(zone)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GangZoneFlashForPlayer
func GangZoneFlashForPlayer(playerid int, zone int, flashcolor int) bool {
	return bool(C.GangZoneFlashForPlayer(C.int(playerid), C.int(zone), C.int(flashcolor)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GangZoneFlashForAll
func GangZoneFlashForAll(zone int, flashcolor int) bool {
	return bool(C.GangZoneFlashForAll(C.int(zone), C.int(flashcolor)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GangZoneStopFlashForPlayer
func GangZoneStopFlashForPlayer(playerid int, zone int) bool {
	return bool(C.GangZoneStopFlashForPlayer(C.int(playerid), C.int(zone)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GangZoneStopFlashForAll
func GangZoneStopFlashForAll(zone int) bool {
	return bool(C.GangZoneStopFlashForAll(C.int(zone)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/Create3DTextLabel
func Create3DTextLabel(text string, color int, x float32, y float32, z float32, DrawDistance float32, virtualworld int, testLOS bool) int {
	cstext := C.CString(text)
	defer C.free(unsafe.Pointer(cstext))
	return int(C.Create3DTextLabel(C.nonConstToConst(cstext), C.int(color), C.float(x), C.float(y), C.float(z), C.float(DrawDistance), C.int(virtualworld), C.bool(testLOS)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/Delete3DTextLabel
func Delete3DTextLabel(id int) bool {
	return bool(C.Delete3DTextLabel(C.int(id)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/Attach3DTextLabelToPlayer
func Attach3DTextLabelToPlayer(id int, playerid int, OffsetX float32, OffsetY float32, OffsetZ float32) bool {
	return bool(C.Attach3DTextLabelToPlayer(C.int(id), C.int(playerid), C.float(OffsetX), C.float(OffsetY), C.float(OffsetZ)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/Attach3DTextLabelToVehicle
func Attach3DTextLabelToVehicle(id int, vehicleid int, OffsetX float32, OffsetY float32, OffsetZ float32) bool {
	return bool(C.Attach3DTextLabelToVehicle(C.int(id), C.int(vehicleid), C.float(OffsetX), C.float(OffsetY), C.float(OffsetZ)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/Update3DTextLabelText
func Update3DTextLabelText(id int, color int, text string) bool {
	cstext := C.CString(text)
	defer C.free(unsafe.Pointer(cstext))
	return bool(C.Update3DTextLabelText(C.int(id), C.int(color), C.nonConstToConst(cstext)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/CreatePlayer3DTextLabel
func CreatePlayer3DTextLabel(playerid int, text string, color int, x float32, y float32, z float32, DrawDistance float32, attachedplayer int, attachedvehicle int, testLOS bool) int {
	cstext := C.CString(text)
	defer C.free(unsafe.Pointer(cstext))
	return int(C.CreatePlayer3DTextLabel(C.int(playerid), C.nonConstToConst(cstext), C.int(color), C.float(x), C.float(y), C.float(z), C.float(DrawDistance), C.int(attachedplayer), C.int(attachedvehicle), C.bool(testLOS)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/DeletePlayer3DTextLabel
func DeletePlayer3DTextLabel(playerid int, id int) bool {
	return bool(C.DeletePlayer3DTextLabel(C.int(playerid), C.int(id)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/UpdatePlayer3DTextLabelText
func UpdatePlayer3DTextLabelText(playerid int, id int, color int, text string) bool {
	cstext := C.CString(text)
	defer C.free(unsafe.Pointer(cstext))
	return bool(C.UpdatePlayer3DTextLabelText(C.int(playerid), C.int(id), C.int(color), C.nonConstToConst(cstext)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/ShowPlayerDialog
func ShowPlayerDialog(playerid int, dialogid int, style int, caption string, info string, button1 string, button2 string) bool {
	cscaption := C.CString(caption)
	defer C.free(unsafe.Pointer(cscaption))
	csinfo := C.CString(info)
	defer C.free(unsafe.Pointer(csinfo))
	csbutton1 := C.CString(button1)
	defer C.free(unsafe.Pointer(csbutton1))
	csbutton2 := C.CString(button2)
	defer C.free(unsafe.Pointer(csbutton2))
	return bool(C.ShowPlayerDialog(C.int(playerid), C.int(dialogid), C.int(style), C.nonConstToConst(cscaption), C.nonConstToConst(csinfo), C.nonConstToConst(csbutton1), C.nonConstToConst(csbutton2)))
}

// Native SetTimer was not generated because it is marked as 'noimpl'
// Native KillTimer was not generated because it is marked as 'noimpl'
// For documentation, please visit https://open.mp/docs/scripting/functions/gpci
func gpci(playerid int, buffer *string, size int) bool {
	var ret bool
	var cbuffer *C.char
	cbuffer = (*C.char)(C.malloc(C.uint(size)))
	defer C.free(unsafe.Pointer(cbuffer))
	ret = bool(C.gpci(C.int(playerid), cbuffer, C.int(size)))
	*buffer = C.GoString(C.constToNonConst(cbuffer))
	return ret
}

// For documentation, please visit https://open.mp/docs/scripting/functions/AddCharModel
func AddCharModel(baseid int, newid int, dffname string, txdname string) int {
	csdffname := C.CString(dffname)
	defer C.free(unsafe.Pointer(csdffname))
	cstxdname := C.CString(txdname)
	defer C.free(unsafe.Pointer(cstxdname))
	return int(C.AddCharModel(C.int(baseid), C.int(newid), C.nonConstToConst(csdffname), C.nonConstToConst(cstxdname)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/AddSimpleModel
func AddSimpleModel(virtualworld int, baseid int, newid int, dffname string, txdname string) int {
	csdffname := C.CString(dffname)
	defer C.free(unsafe.Pointer(csdffname))
	cstxdname := C.CString(txdname)
	defer C.free(unsafe.Pointer(cstxdname))
	return int(C.AddSimpleModel(C.int(virtualworld), C.int(baseid), C.int(newid), C.nonConstToConst(csdffname), C.nonConstToConst(cstxdname)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/AddSimpleModelTimed
func AddSimpleModelTimed(virtualworld int, baseid int, newid int, dffname string, txdname string, timeon int, timeoff int) int {
	csdffname := C.CString(dffname)
	defer C.free(unsafe.Pointer(csdffname))
	cstxdname := C.CString(txdname)
	defer C.free(unsafe.Pointer(cstxdname))
	return int(C.AddSimpleModelTimed(C.int(virtualworld), C.int(baseid), C.int(newid), C.nonConstToConst(csdffname), C.nonConstToConst(cstxdname), C.int(timeon), C.int(timeoff)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/FindModelFileNameFromCRC
func FindModelFileNameFromCRC(crc int, model_str *string, model_str_len int) bool {
	var ret bool
	var cmodel_str *C.char
	cmodel_str = (*C.char)(C.malloc(C.uint(model_str_len)))
	defer C.free(unsafe.Pointer(cmodel_str))
	ret = bool(C.FindModelFileNameFromCRC(C.int(crc), cmodel_str, C.int(model_str_len)))
	*model_str = C.GoString(C.constToNonConst(cmodel_str))
	return ret
}

// For documentation, please visit https://open.mp/docs/scripting/functions/FindTextureFileNameFromCRC
func FindTextureFileNameFromCRC(crc int, texture_str *string, texture_str_len int) bool {
	var ret bool
	var ctexture_str *C.char
	ctexture_str = (*C.char)(C.malloc(C.uint(texture_str_len)))
	defer C.free(unsafe.Pointer(ctexture_str))
	ret = bool(C.FindTextureFileNameFromCRC(C.int(crc), ctexture_str, C.int(texture_str_len)))
	*texture_str = C.GoString(C.constToNonConst(ctexture_str))
	return ret
}

// For documentation, please visit https://open.mp/docs/scripting/functions/RedirectDownload
func RedirectDownload(playerid int, url string) bool {
	csurl := C.CString(url)
	defer C.free(unsafe.Pointer(csurl))
	return bool(C.RedirectDownload(C.int(playerid), C.nonConstToConst(csurl)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/CreateObject
func CreateObject(modelid int, x float32, y float32, z float32, rX float32, rY float32, rZ float32, DrawDistance float32) int {
	return int(C.CreateObject(C.int(modelid), C.float(x), C.float(y), C.float(z), C.float(rX), C.float(rY), C.float(rZ), C.float(DrawDistance)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/AttachObjectToVehicle
func AttachObjectToVehicle(objectid int, vehicleid int, fOffsetX float32, fOffsetY float32, fOffsetZ float32, fRotX float32, fRotY float32, fRotZ float32) bool {
	return bool(C.AttachObjectToVehicle(C.int(objectid), C.int(vehicleid), C.float(fOffsetX), C.float(fOffsetY), C.float(fOffsetZ), C.float(fRotX), C.float(fRotY), C.float(fRotZ)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/AttachObjectToObject
func AttachObjectToObject(objectid int, attachtoid int, fOffsetX float32, fOffsetY float32, fOffsetZ float32, fRotX float32, fRotY float32, fRotZ float32, SyncRotation bool) bool {
	return bool(C.AttachObjectToObject(C.int(objectid), C.int(attachtoid), C.float(fOffsetX), C.float(fOffsetY), C.float(fOffsetZ), C.float(fRotX), C.float(fRotY), C.float(fRotZ), C.bool(SyncRotation)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/AttachObjectToPlayer
func AttachObjectToPlayer(objectid int, playerid int, fOffsetX float32, fOffsetY float32, fOffsetZ float32, fRotX float32, fRotY float32, fRotZ float32) bool {
	return bool(C.AttachObjectToPlayer(C.int(objectid), C.int(playerid), C.float(fOffsetX), C.float(fOffsetY), C.float(fOffsetZ), C.float(fRotX), C.float(fRotY), C.float(fRotZ)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SetObjectPos
func SetObjectPos(objectid int, x float32, y float32, z float32) bool {
	return bool(C.SetObjectPos(C.int(objectid), C.float(x), C.float(y), C.float(z)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetObjectPos
func GetObjectPos(objectid int, x *float32, y *float32, z *float32) bool {
	var ret bool
	var cx C.float
	var cy C.float
	var cz C.float
	ret = bool(C.GetObjectPos(C.int(objectid), &cx, &cy, &cz))
	*x = float32(cx)
	*y = float32(cy)
	*z = float32(cz)
	return ret
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SetObjectRot
func SetObjectRot(objectid int, rotX float32, rotY float32, rotZ float32) bool {
	return bool(C.SetObjectRot(C.int(objectid), C.float(rotX), C.float(rotY), C.float(rotZ)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetObjectRot
func GetObjectRot(objectid int, rotX *float32, rotY *float32, rotZ *float32) bool {
	var ret bool
	var crotX C.float
	var crotY C.float
	var crotZ C.float
	ret = bool(C.GetObjectRot(C.int(objectid), &crotX, &crotY, &crotZ))
	*rotX = float32(crotX)
	*rotY = float32(crotY)
	*rotZ = float32(crotZ)
	return ret
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetObjectModel
func GetObjectModel(objectid int) int {
	return int(C.GetObjectModel(C.int(objectid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SetObjectNoCameraCol
func SetObjectNoCameraCol(objectid int) bool {
	return bool(C.SetObjectNoCameraCol(C.int(objectid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/IsValidObject
func IsValidObject(objectid int) bool {
	return bool(C.IsValidObject(C.int(objectid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/DestroyObject
func DestroyObject(objectid int) bool {
	return bool(C.DestroyObject(C.int(objectid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/MoveObject
func MoveObject(objectid int, X float32, Y float32, Z float32, Speed float32, RotX float32, RotY float32, RotZ float32) int {
	return int(C.MoveObject(C.int(objectid), C.float(X), C.float(Y), C.float(Z), C.float(Speed), C.float(RotX), C.float(RotY), C.float(RotZ)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/StopObject
func StopObject(objectid int) bool {
	return bool(C.StopObject(C.int(objectid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/IsObjectMoving
func IsObjectMoving(objectid int) bool {
	return bool(C.IsObjectMoving(C.int(objectid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/EditObject
func EditObject(playerid int, objectid int) bool {
	return bool(C.EditObject(C.int(playerid), C.int(objectid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/EditPlayerObject
func EditPlayerObject(playerid int, objectid int) bool {
	return bool(C.EditPlayerObject(C.int(playerid), C.int(objectid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SelectObject
func SelectObject(playerid int) bool {
	return bool(C.sampgdk_SelectObject(C.int(playerid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/CancelEdit
func CancelEdit(playerid int) bool {
	return bool(C.CancelEdit(C.int(playerid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/CreatePlayerObject
func CreatePlayerObject(playerid int, modelid int, x float32, y float32, z float32, rX float32, rY float32, rZ float32, DrawDistance float32) int {
	return int(C.CreatePlayerObject(C.int(playerid), C.int(modelid), C.float(x), C.float(y), C.float(z), C.float(rX), C.float(rY), C.float(rZ), C.float(DrawDistance)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/AttachPlayerObjectToPlayer
func AttachPlayerObjectToPlayer(objectplayer int, objectid int, attachplayer int, OffsetX float32, OffsetY float32, OffsetZ float32, rX float32, rY float32, rZ float32) bool {
	return bool(C.AttachPlayerObjectToPlayer(C.int(objectplayer), C.int(objectid), C.int(attachplayer), C.float(OffsetX), C.float(OffsetY), C.float(OffsetZ), C.float(rX), C.float(rY), C.float(rZ)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/AttachPlayerObjectToVehicle
func AttachPlayerObjectToVehicle(playerid int, objectid int, vehicleid int, fOffsetX float32, fOffsetY float32, fOffsetZ float32, fRotX float32, fRotY float32, RotZ float32) bool {
	return bool(C.AttachPlayerObjectToVehicle(C.int(playerid), C.int(objectid), C.int(vehicleid), C.float(fOffsetX), C.float(fOffsetY), C.float(fOffsetZ), C.float(fRotX), C.float(fRotY), C.float(RotZ)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SetPlayerObjectPos
func SetPlayerObjectPos(playerid int, objectid int, x float32, y float32, z float32) bool {
	return bool(C.SetPlayerObjectPos(C.int(playerid), C.int(objectid), C.float(x), C.float(y), C.float(z)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetPlayerObjectPos
func GetPlayerObjectPos(playerid int, objectid int, x *float32, y *float32, z *float32) bool {
	var ret bool
	var cx C.float
	var cy C.float
	var cz C.float
	ret = bool(C.GetPlayerObjectPos(C.int(playerid), C.int(objectid), &cx, &cy, &cz))
	*x = float32(cx)
	*y = float32(cy)
	*z = float32(cz)
	return ret
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SetPlayerObjectRot
func SetPlayerObjectRot(playerid int, objectid int, rotX float32, rotY float32, rotZ float32) bool {
	return bool(C.SetPlayerObjectRot(C.int(playerid), C.int(objectid), C.float(rotX), C.float(rotY), C.float(rotZ)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetPlayerObjectRot
func GetPlayerObjectRot(playerid int, objectid int, rotX *float32, rotY *float32, rotZ *float32) bool {
	var ret bool
	var crotX C.float
	var crotY C.float
	var crotZ C.float
	ret = bool(C.GetPlayerObjectRot(C.int(playerid), C.int(objectid), &crotX, &crotY, &crotZ))
	*rotX = float32(crotX)
	*rotY = float32(crotY)
	*rotZ = float32(crotZ)
	return ret
}

// For documentation, please visit https://open.mp/docs/scripting/functions/GetPlayerObjectModel
func GetPlayerObjectModel(playerid int, objectid int) int {
	return int(C.GetPlayerObjectModel(C.int(playerid), C.int(objectid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SetPlayerObjectNoCameraCol
func SetPlayerObjectNoCameraCol(playerid int, objectid int) bool {
	return bool(C.SetPlayerObjectNoCameraCol(C.int(playerid), C.int(objectid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/IsValidPlayerObject
func IsValidPlayerObject(playerid int, objectid int) bool {
	return bool(C.IsValidPlayerObject(C.int(playerid), C.int(objectid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/DestroyPlayerObject
func DestroyPlayerObject(playerid int, objectid int) bool {
	return bool(C.DestroyPlayerObject(C.int(playerid), C.int(objectid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/MovePlayerObject
func MovePlayerObject(playerid int, objectid int, x float32, y float32, z float32, Speed float32, RotX float32, RotY float32, RotZ float32) int {
	return int(C.MovePlayerObject(C.int(playerid), C.int(objectid), C.float(x), C.float(y), C.float(z), C.float(Speed), C.float(RotX), C.float(RotY), C.float(RotZ)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/StopPlayerObject
func StopPlayerObject(playerid int, objectid int) bool {
	return bool(C.StopPlayerObject(C.int(playerid), C.int(objectid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/IsPlayerObjectMoving
func IsPlayerObjectMoving(playerid int, objectid int) bool {
	return bool(C.IsPlayerObjectMoving(C.int(playerid), C.int(objectid)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SetObjectMaterial
func SetObjectMaterial(objectid int, materialindex int, modelid int, txdname string, texturename string, materialcolor int) bool {
	cstxdname := C.CString(txdname)
	defer C.free(unsafe.Pointer(cstxdname))
	cstexturename := C.CString(texturename)
	defer C.free(unsafe.Pointer(cstexturename))
	return bool(C.SetObjectMaterial(C.int(objectid), C.int(materialindex), C.int(modelid), C.nonConstToConst(cstxdname), C.nonConstToConst(cstexturename), C.int(materialcolor)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SetPlayerObjectMaterial
func SetPlayerObjectMaterial(playerid int, objectid int, materialindex int, modelid int, txdname string, texturename string, materialcolor int) bool {
	cstxdname := C.CString(txdname)
	defer C.free(unsafe.Pointer(cstxdname))
	cstexturename := C.CString(texturename)
	defer C.free(unsafe.Pointer(cstexturename))
	return bool(C.SetPlayerObjectMaterial(C.int(playerid), C.int(objectid), C.int(materialindex), C.int(modelid), C.nonConstToConst(cstxdname), C.nonConstToConst(cstexturename), C.int(materialcolor)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SetObjectMaterialText
func SetObjectMaterialText(objectid int, text string, materialindex int, materialsize int, fontface string, fontsize int, bold bool, fontcolor int, backcolor int, textalignment int) bool {
	cstext := C.CString(text)
	defer C.free(unsafe.Pointer(cstext))
	csfontface := C.CString(fontface)
	defer C.free(unsafe.Pointer(csfontface))
	return bool(C.SetObjectMaterialText(C.int(objectid), C.nonConstToConst(cstext), C.int(materialindex), C.int(materialsize), C.nonConstToConst(csfontface), C.int(fontsize), C.bool(bold), C.int(fontcolor), C.int(backcolor), C.int(textalignment)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SetPlayerObjectMaterialText
func SetPlayerObjectMaterialText(playerid int, objectid int, text string, materialindex int, materialsize int, fontface string, fontsize int, bold bool, fontcolor int, backcolor int, textalignment int) bool {
	cstext := C.CString(text)
	defer C.free(unsafe.Pointer(cstext))
	csfontface := C.CString(fontface)
	defer C.free(unsafe.Pointer(csfontface))
	return bool(C.SetPlayerObjectMaterialText(C.int(playerid), C.int(objectid), C.nonConstToConst(cstext), C.int(materialindex), C.int(materialsize), C.nonConstToConst(csfontface), C.int(fontsize), C.bool(bold), C.int(fontcolor), C.int(backcolor), C.int(textalignment)))
}

// For documentation, please visit https://open.mp/docs/scripting/functions/SetObjectsDefaultCameraCol
func SetObjectsDefaultCameraCol(disable bool) bool {
	return bool(C.SetObjectsDefaultCameraCol(C.bool(disable)))
}
