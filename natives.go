package sampgo

/*
#cgo CFLAGS: -Wno-attributes

#cgo linux.386 LDFLAGS: -L. -l:sampgdk/build/bin/Debug/libsampgdk.a
#cgo linux.386 LDFLAGS: -ldl

#cgo windows.386 LDFLAGS: -L. -l:sampgdk/build/bin/Debug/sampgdk.lib

#ifndef GOLANG_APP
#define GOLANG_APP

#include "sampgdk/main.h"

#endif
*/
import "C"
import "unsafe"

// CreateActor implements
func CreateActor(modelid int, x float32, y float32, z float32, rotation float32) int {
	return int(C.CreateActor(C.int(modelid), C.float(x), C.float(y), C.float(z), C.float(rotation)))
}

// DestroyActor implements
func DestroyActor(actorid int) bool {
	return bool(C.DestroyActor(C.int(actorid)))
}

// IsActorStreamedIn implements
func IsActorStreamedIn(actorid int, forplayer Player) bool {
	return bool(C.IsActorStreamedIn(C.int(actorid), C.int(forplayer.ID)))
}

// SetActorVirtualWorld implements
func SetActorVirtualWorld(actorid int, vworld int) bool {
	return bool(C.SetActorVirtualWorld(C.int(actorid), C.int(vworld)))
}

// GetActorVirtualWorld implements
func GetActorVirtualWorld(actorid int) int {
	return int(C.GetActorVirtualWorld(C.int(actorid)))
}

// ApplyActorAnimation implements
func ApplyActorAnimation(actorid int, animlib, animname string, fDelta float32, loop, lockx, locky, freeze bool, time int) bool {
	lib := C.CString(animlib)
	name := C.CString(animname)

	defer C.free(unsafe.Pointer(lib))
	defer C.free(unsafe.Pointer(name))

	return bool(C.ApplyActorAnimation(C.int(actorid), C.nonConstToConst(lib), C.nonConstToConst(name), C.float(fDelta), C.bool(loop), C.bool(lockx), C.bool(locky), C.bool(freeze), C.int(time)))
}

// ClearActorAnimations implements
func ClearActorAnimations(actorid int) bool {
	return bool(C.ClearActorAnimations(C.int(actorid)))
}

// SetActorPos implements
func SetActorPos(actorid int, x float32, y float32, z float32) bool {
	return bool(C.SetActorPos(C.int(actorid), C.float(x), C.float(y), C.float(z)))
}

// GetActorPos implements
func GetActorPos(actorid int) (float32, float32, float32) {
	var x, y, z C.float
	C.GetActorPos(C.int(actorid), &x, &y, &z)
	return float32(x), float32(y), float32(z)
}

// SetActorFacingAngle implements
func SetActorFacingAngle(actorid int, angle float32) bool {
	return bool(C.SetActorFacingAngle(C.int(actorid), C.float(angle)))
}

// GetActorFacingAngle implements
func GetActorFacingAngle(actorid int) float32 {
	var a C.float
	C.GetActorFacingAngle(C.int(actorid), &a)
	return float32(a)
}

/*
// SetActorHealth implements
func SetActorHealth(actorid int, health float32) bool {

}

// GetActorHealth implements
func GetActorHealth(actorid int, health float32) bool {

}

// SetActorInvulnerable implements
func SetActorInvulnerable(actorid int, invulnerable bool) bool {

}

// IsActorInvulnerable implements
func IsActorInvulnerable(actorid int) bool {

}

// IsValidActor implements
func IsValidActor(actorid int) bool {

}*/

// IsValidVehicle implements
func IsValidVehicle(vehicleid int) bool {
	return bool(C.IsValidVehicle(C.int(vehicleid)))
}

// GetVehicleDistanceFromPoint returns how far a vehicle is from a certain point in relation to a certain XYZ coordinate.
func GetVehicleDistanceFromPoint(vehicleid int, x float32, y float32, z float32) float32 {
	return float32(C.GetVehicleDistanceFromPoint(C.int(vehicleid), C.float(x), C.float(y), C.float(z)))
}

// CreateVehicle allows you to create a vehicle.
func CreateVehicle(vehicletype int, x float32, y float32, z float32, rotation float32, color1 int, color2 int, respawnDelay int, addsiren bool) int {
	return int(C.CreateVehicle(C.int(vehicletype), C.float(x), C.float(y), C.float(z), C.float(rotation), C.int(color1), C.int(color2), C.int(respawnDelay), C.bool(addsiren)))
}

// DestroyVehicle allows you to destroy a vehicle.
func DestroyVehicle(vehicleid int) bool {
	return bool(C.DestroyVehicle(C.int(vehicleid)))
}

/*
// IsVehicleStreamedIn implements
func IsVehicleStreamedIn(vehicleid int, forplayerid int) bool {

}

// GetVehiclePos implements
func GetVehiclePos(vehicleid int, x float32, y float32, z float32) bool {

}

// SetVehiclePos implements
func SetVehiclePos(vehicleid int, x float32, y float32, z float32) bool {

}

// GetVehicleZAngle implements
func GetVehicleZAngle(vehicleid int, zAngle float32) bool {

}

// GetVehicleRotationQuat implements
func GetVehicleRotationQuat(vehicleid int, w float32, x float32, y float32, z float32) bool {

}

// SetVehicleZAngle implements
func SetVehicleZAngle(vehicleid int, zAngle float32) bool {

}

// SetVehicleParamsForPlayer implements
func SetVehicleParamsForPlayer(vehicleid int, playerid int, objective int, doorslocked int) bool {

}

// ManualVehicleEngineAndLights implements
func ManualVehicleEngineAndLights() bool {

}

// SetVehicleParamsEx implements
func SetVehicleParamsEx(vehicleid int, engine int, lights int, alarm int, doors int, bonnet int, boot int, objective int) bool {

}

// GetVehicleParamsEx implements
func GetVehicleParamsEx(vehicleid int, engine int, lights int, alarm int, doors int, bonnet int, boot int, objective int) bool {

}

// GetVehicleParamsSirenState implements
func GetVehicleParamsSirenState(vehicleid int) int {

}

// SetVehicleParamsCarDoors implements
func SetVehicleParamsCarDoors(vehicleid int, driver int, passenger int, backleft int, backright int) bool {

}

// GetVehicleParamsCarDoors implements
func GetVehicleParamsCarDoors(vehicleid int, driver int, passenger int, backleft int, backright int) bool {

}

// SetVehicleParamsCarWindows implements
func SetVehicleParamsCarWindows(vehicleid int, driver int, passenger int, backleft int, backright int) bool {

}

// GetVehicleParamsCarWindows implements
func GetVehicleParamsCarWindows(vehicleid int, driver int, passenger int, backleft int, backright int) bool {

}

// SetVehicleToRespawn implements
func SetVehicleToRespawn(vehicleid int) bool {

}

// LinkVehicleToInterior implements
func LinkVehicleToInterior(vehicleid int, interiorid int) bool {

}

// AddVehicleComponent implements
func AddVehicleComponent(vehicleid int, componentid int) bool {

}

// RemoveVehicleComponent implements
func RemoveVehicleComponent(vehicleid int, componentid int) bool {

}

// ChangeVehicleColor implements
func ChangeVehicleColor(vehicleid int, color1 int, color2 int) bool {

}

// ChangeVehiclePaintjob implements
func ChangeVehiclePaintjob(vehicleid int, paintjobid int) bool {

}

// SetVehicleHealth implements
func SetVehicleHealth(vehicleid int, health float32) bool {

}

// GetVehicleHealth implements
func GetVehicleHealth(vehicleid int, health float32) bool {

}

// AttachTrailerToVehicle implements
func AttachTrailerToVehicle(trailerid int, vehicleid int) bool {

}

// DetachTrailerFromVehicle implements
func DetachTrailerFromVehicle(vehicleid int) bool {

}

// IsTrailerAttachedToVehicle implements
func IsTrailerAttachedToVehicle(vehicleid int) bool {

}

// GetVehicleTrailer implements
func GetVehicleTrailer(vehicleid int) int {

}

// SetVehicleNumberPlate implements
func SetVehicleNumberPlate(vehicleid int, numberplate string) bool {

}

// GetVehicleModel implements
func GetVehicleModel(vehicleid int) int {

}

// GetVehicleComponentInSlot implements
func GetVehicleComponentInSlot(vehicleid int, slot int) int {

}

// GetVehicleComponentType implements
func GetVehicleComponentType(component int) int {

}

// RepairVehicle implements
func RepairVehicle(vehicleid int) bool {

}

// GetVehicleVelocity implements
func GetVehicleVelocity(vehicleid int, X float32, Y float32, Z float32) bool {

}

// SetVehicleVelocity implements
func SetVehicleVelocity(vehicleid int, X float32, Y float32, Z float32) bool {

}

// SetVehicleAngularVelocity implements
func SetVehicleAngularVelocity(vehicleid int, X float32, Y float32, Z float32) bool {

}

// GetVehicleDamageStatus implements
func GetVehicleDamageStatus(vehicleid int, panels int, doors int, lights int, tires int) bool {

}

// UpdateVehicleDamageStatus implements
func UpdateVehicleDamageStatus(vehicleid int, panels int, doors int, lights int, tires int) bool {

}

// SetVehicleVirtualWorld implements
func SetVehicleVirtualWorld(vehicleid int, worldid int) bool {

}

// GetVehicleVirtualWorld implements
func GetVehicleVirtualWorld(vehicleid int) int {

}

// GetVehicleModelInfo implements
func GetVehicleModelInfo(model int, infotype int, X float32, Y float32, Z float32) bool {

}*/

// SendClientMessage implements
func SendClientMessage(playerid, color int, message string) error {
	msg := C.CString(message)
	defer C.free(unsafe.Pointer(msg))

	C.SendClientMessage(C.int(playerid), C.int(color), C.nonConstToConst(msg))
	return nil
}

// SendClientMessageToAll implements
func SendClientMessageToAll(color int, message string) error {
	msg := C.CString(message)
	defer C.free(unsafe.Pointer(msg))

	C.SendClientMessageToAll(C.int(color), C.nonConstToConst(msg))
	return nil
}

/*
// SendDeathMessage implements
func SendDeathMessage(killer int, killee int, weapon int) bool {

}

// SendDeathMessageToPlayer implements
func SendDeathMessageToPlayer(playerid int, killer int, killee int, weapon int) bool {

}

// GameTextForAll implements
func GameTextForAll(text string, time int, style int) bool {

}

// GameTextForPlayer implements
func GameTextForPlayer(playerid int, text string, time int, style int) bool {

}

// GetTickCount implements
func GetTickCount() int {

}

// GetMaxPlayers implements
func GetMaxPlayers() int {

}

// VectorSize implements
func VectorSize(x float32, y float32, z float32) float32 {

}

// GetPlayerPoolSize implements
func GetPlayerPoolSize() int {

}

// GetVehiclePoolSize implements
func GetVehiclePoolSize() int {

}

// GetActorPoolSize implements
func GetActorPoolSize() int {

}

// SetGameModeText implements
func SetGameModeText(text string) bool {

}

// SetTeamCount implements
func SetTeamCount(count int) bool {

}

// AddPlayerClass implements
func AddPlayerClass(modelid int, spawnX float32, spawnY float32, spawnZ float32, zAngle float32, weapon1 int, weapon1Ammo int, weapon2 int, weapon2Ammo int, weapon3 int, weapon3Ammo int) int {

}

// AddPlayerClassEx implements
func AddPlayerClassEx(teamid int, modelid int, spawn_x float32, spawn_y float32, spawn_z float32, z_angle float32, weapon1 int, weapon1_ammo int, weapon2 int, weapon2_ammo int, weapon3 int, weapon3_ammo int) int {

}

// AddStaticVehicle implements
func AddStaticVehicle(modelid int, spawn_x float32, spawn_y float32, spawn_z float32, z_angle float32, color1 int, color2 int) int {

}

// AddStaticVehicleEx implements
func AddStaticVehicleEx(modelid int, spawn_x float32, spawn_y float32, spawn_z float32, z_angle float32, color1 int, color2 int, respawn_delay int, addsiren bool) int {

}

// AddStaticPickup implements
func AddStaticPickup(model int, pType int, x float32, y float32, z float32, virtualworld int) int {

}

// CreatePickup implements
func CreatePickup(model int, pType int, x float32, y float32, z float32, virtualworld int) int {

}

// DestroyPickup implements
func DestroyPickup(pickup int) bool {

}

// ShowNameTags implements
func ShowNameTags(show bool) bool {

}

// ShowPlayerMarkers implements
func ShowPlayerMarkers(mode int) bool {

}

// GameModeExit implements
func GameModeExit() bool {

}

// SetWorldTime implements
func SetWorldTime(hour int) bool {

}

// GetWeaponName implements
func GetWeaponName(weaponid int, name string, size int) bool {

}

// EnableTirePopping implements
func EnableTirePopping(enable bool) bool {

}

// EnableVehicleFriendlyFire implements
func EnableVehicleFriendlyFire() bool {

}

// AllowInteriorWeapons implements
func AllowInteriorWeapons(allow bool) bool {

}

// SetWeather implements
func SetWeather(weatherid int) bool {

}

// SetGravity implements
func SetGravity(gravity float32) bool {

}

// GetGravity implements
func GetGravity() float32 {

}

// AllowAdminTeleport implements
func AllowAdminTeleport(allow bool) bool {

}

// SetDeathDropAmount implements
func SetDeathDropAmount(amount int) bool {

}

// CreateExplosion implements
func CreateExplosion(x float32, y float32, z float32, eType int, radius float32) bool {

}

// EnableZoneNames implements
func EnableZoneNames(enable bool) bool {

}

// UsePlayerPedAnims implements
func UsePlayerPedAnims() bool {

}

// DisableInteriorEnterExits implements
func DisableInteriorEnterExits() bool {

}

// SetNameTagDrawDistance implements
func SetNameTagDrawDistance(distance float32) bool {

}

// DisableNameTagLOS implements
func DisableNameTagLOS() bool {

}

// LimitGlobalChatRadius implements
func LimitGlobalChatRadius(chat_radius float32) bool {

}

// LimitPlayerMarkerRadius implements
func LimitPlayerMarkerRadius(marker_radius float32) bool {

}

// ConnectNPC implements
func ConnectNPC(name string, script string) bool {

}

// IsPlayerNPC implements
func IsPlayerNPC(playerid int) bool {

}

// IsPlayerAdmin implements
func IsPlayerAdmin(playerid int) bool {

}

// Kick implements
func Kick(playerid int) bool {

}

// Ban implements
func Ban(playerid int) bool {

}

// BanEx implements
func BanEx(playerid int, reason string) bool {

}

// SendRconCommand implements
func SendRconCommand(command string) bool {

}

// GetPlayerNetworkStats implements
func GetPlayerNetworkStats(playerid int, retstr string, size int) bool {

}

// GetNetworkStats implements
func GetNetworkStats(retstr string, size int) bool {

}

// GetPlayerVersion implements
func GetPlayerVersion(playerid int, version string, len int) bool {

}

// BlockIpAddress implements
func BlockIPAddress(ip_address string, timems int) bool {

}

// UnBlockIpAddress implements
func UnBlockIPAddress(ip_address string) bool {

}

// GetServerVarAsString implements
func GetServerVarAsString(varname string, value string, size int) bool {

}

// GetServerVarAsimplements int
func GetServerVarAsInt(varname string) int {

}

// GetServerVarAsimplements bool
func GetServerVarAsBool(varname string) bool {

}

// GetConsoleVarAsString implements
func GetConsoleVarAsString(varname string, buffer string, len int) bool {

}

// GetConsoleVarAsimplements int
func GetConsoleVarAsInt(varname string) int {

}

// GetConsoleVarAsimplements bool
func GetConsoleVarAsBool(varname string) bool {

}

// GetServerTickRate implements
func GetServerTickRate() int {

}

// NetStats_GetConnectedTime implements
func NetStats_GetConnectedTime(playerid int) int {

}

// NetStats_MessagesReceived implements
func NetStats_MessagesReceived(playerid int) int {

}

// NetStats_BytesReceived implements
func NetStats_BytesReceived(playerid int) int {

}

// NetStats_MessagesSent implements
func NetStats_MessagesSent(playerid int) int {

}

// NetStats_BytesSent implements
func NetStats_BytesSent(playerid int) int {

}

// NetStats_MessagesRecvPerSecond implements
func NetStats_MessagesRecvPerSecond(playerid int) int {

}

// NetStats_PacketLossPercent implements
func NetStats_PacketLossPercent(playerid int) float32 {

}

// NetStats_ConnectionStatus implements
func NetStats_ConnectionStatus(playerid int) int {

}

// NetStats_GetIpPort implements
func NetStats_GetIpPort(playerid int, ip_port string, ip_port_len int) bool {

}

// CreateMenu implements
func CreateMenu(title string, columns int, x float32, y float32, col1width float32, col2width float32) int {

}

// DestroyMenu implements
func DestroyMenu(menuid int) bool {

}

// AddMenuItem implements
func AddMenuItem(menuid int, column int, menutext string) int {

}

// SetMenuColumnHeader implements
func SetMenuColumnHeader(menuid int, column int, columnheader string) bool {

}

// ShowMenuForPlayer implements
func ShowMenuForPlayer(menuid int, playerid int) bool {

}

// HideMenuForPlayer implements
func HideMenuForPlayer(menuid int, playerid int) bool {

}

// IsValidMenu implements
func IsValidMenu(menuid int) bool {

}

// DisableMenu implements
func DisableMenu(menuid int) bool {

}

// DisableMenuRow implements
func DisableMenuRow(menuid int, row int) bool {

}

// GetPlayerMenu implements
func GetPlayerMenu(playerid int) int {

}

// TextDrawCreate implements
func TextDrawCreate(x float32, y float32, text string) int {

}

// TextDrawDestroy implements
func TextDrawDestroy(text int) bool {

}

// TextDrawLetterSize implements
func TextDrawLetterSize(text int, x float32, y float32) bool {

}

// TextDrawTextSize implements
func TextDrawTextSize(text int, x float32, y float32) bool {

}

// TextDrawAlignment implements
func TextDrawAlignment(text int, alignment int) bool {

}

// TextDrawColor implements
func TextDrawColor(text int, color int) bool {

}

// TextDrawUseBox implements
func TextDrawUseBox(text int, use bool) bool {

}

// TextDrawBoxColor implements
func TextDrawBoxColor(text int, color int) bool {

}

// TextDrawSetShadow implements
func TextDrawSetShadow(text int, size int) bool {

}

// TextDrawSetOutline implements
func TextDrawSetOutline(text int, size int) bool {

}

// TextDrawBackgroundColor implements
func TextDrawBackgroundColor(text int, color int) bool {

}

// TextDrawFont implements
func TextDrawFont(text int, font int) bool {

}

// TextDrawSetProportional implements
func TextDrawSetProportional(text int, set bool) bool {

}

// TextDrawSetSelectable implements
func TextDrawSetSelectable(text int, set bool) bool {

}

// TextDrawShowForPlayer implements
func TextDrawShowForPlayer(playerid int, text int) bool {

}

// TextDrawHideForPlayer implements
func TextDrawHideForPlayer(playerid int, text int) bool {

}

// TextDrawShowForAll implements
func TextDrawShowForAll(text int) bool {

}

// TextDrawHideForAll implements
func TextDrawHideForAll(text int) bool {

}

// TextDrawSetString implements
func TextDrawSetString(text int, string string) bool {

}

// TextDrawSetPreviewModel implements
func TextDrawSetPreviewModel(text int, modelindex int) bool {

}

// TextDrawSetPreviewRot implements
func TextDrawSetPreviewRot(text int, fRotX float32, fRotY float32, fRotZ float32, fZoom float32) bool {

}

// TextDrawSetPreviewVehCol implements
func TextDrawSetPreviewVehCol(text int, color1 int, color2 int) bool {

}

// SelectTextDraw implements
func SelectTextDraw(playerid int, hovercolor int) bool {

}

// CancelSelectTextDraw implements
func CancelSelectTextDraw(playerid int) bool {

}

// GangZoneCreate implements
func GangZoneCreate(minx float32, miny float32, maxx float32, maxy float32) int {

}

// GangZoneDestroy implements
func GangZoneDestroy(zone int) bool {

}

// GangZoneShowForPlayer implements
func GangZoneShowForPlayer(playerid int, zone int, color int) bool {

}

// GangZoneShowForAll implements
func GangZoneShowForAll(zone int, color int) bool {

}

// GangZoneHideForPlayer implements
func GangZoneHideForPlayer(playerid int, zone int) bool {

}

// GangZoneHideForAll implements
func GangZoneHideForAll(zone int) bool {

}

// GangZoneFlashForPlayer implements
func GangZoneFlashForPlayer(playerid int, zone int, flashcolor int) bool {

}

// GangZoneFlashForAll implements
func GangZoneFlashForAll(zone int, flashcolor int) bool {

}

// GangZoneStopFlashForPlayer implements
func GangZoneStopFlashForPlayer(playerid int, zone int) bool {

}

// GangZoneStopFlashForAll implements
func GangZoneStopFlashForAll(zone int) bool {

}

// Create3DTextLabel implements
func Create3DTextLabel(text string, color int, x float32, y float32, z float32, DrawDistance float32, virtualworld int, testLOS bool) int {

}

// Delete3DTextLabel implements
func Delete3DTextLabel(id int) bool {

}

// Attach3DTextLabelToPlayer implements
func Attach3DTextLabelToPlayer(id int, playerid int, OffsetX float32, OffsetY float32, OffsetZ float32) bool {

}

// Attach3DTextLabelToVehicle implements
func Attach3DTextLabelToVehicle(id int, vehicleid int, OffsetX float32, OffsetY float32, OffsetZ float32) bool {

}

// Update3DTextLabelText implements
func Update3DTextLabelText(id int, color int, text string) bool {

}

// CreatePlayer3DTextLabel implements
func CreatePlayer3DTextLabel(playerid int, text string, color int, x float32, y float32, z float32, DrawDistance float32, attachedplayer int, attachedvehicle int, testLOS bool) int {

}

// DeletePlayer3DTextLabel implements
func DeletePlayer3DTextLabel(playerid int, id int) bool {

}

// UpdatePlayer3DTextLabelText implements
func UpdatePlayer3DTextLabelText(playerid int, id int, color int, text string) bool {

}

// ShowPlayerDialog implements
func ShowPlayerDialog(playerid int, dialogid int, style int, caption string, info string, button1 string, button2 string) bool {

}

// gpci implements
func Gpci(playerid int, buffer string, size int) bool {

}

// AddCharModel implements
func AddCharModel(baseid int, newid int, dffname string, txdname string) int {

}

// AddSimpleModel implements
func AddSimpleModel(virtualworld int, baseid int, newid int, dffname string, txdname string) int {

}

// AddSimpleModelTimed implements
func AddSimpleModelTimed(virtualworld int, baseid int, newid int, dffname string, txdname string, timeon int, timeoff int) int {

}

// FindModelFileNameFromCRC implements
func FindModelFileNameFromCRC(crc int, model_str string, model_str_len int) bool {

}

// FindTextureFileNameFromCRC implements
func FindTextureFileNameFromCRC(crc int, texture_str string, texture_str_len int) bool {

}

// RedirectDownload implements
func RedirectDownload(playerid int, url string) bool {

}

// CreateObject implements
func CreateObject(modelid int, x float32, y float32, z float32, rX float32, rY float32, rZ float32, DrawDistance float32) int {

}

// AttachObjectToVehicle implements
func AttachObjectToVehicle(objectid int, vehicleid int, fOffsetX float32, fOffsetY float32, fOffsetZ float32, fRotX float32, fRotY float32, fRotZ float32) bool {

}

// AttachObjectToObject implements
func AttachObjectToObject(objectid int, attachtoid int, fOffsetX float32, fOffsetY float32, fOffsetZ float32, fRotX float32, fRotY float32, fRotZ float32, SyncRotation bool) bool {

}

// AttachObjectToPlayer implements
func AttachObjectToPlayer(objectid int, playerid int, fOffsetX float32, fOffsetY float32, fOffsetZ float32, fRotX float32, fRotY float32, fRotZ float32) bool {

}

// SetObjectPos implements
func SetObjectPos(objectid int, x float32, y float32, z float32) bool {

}

// GetObjectPos implements
func GetObjectPos(objectid int, x float32, y float32, z float32) bool {

}

// SetObjectRot implements
func SetObjectRot(objectid int, rotX float32, rotY float32, rotZ float32) bool {

}

// GetObjectRot implements
func GetObjectRot(objectid int, rotX float32, rotY float32, rotZ float32) bool {

}

// GetObjectModel implements
func GetObjectModel(objectid int) int {

}

// SetObjectNoCameraCol implements
func SetObjectNoCameraCol(objectid int) bool {

}

// IsValidObject implements
func IsValidObject(objectid int) bool {

}

// DestroyObject implements
func DestroyObject(objectid int) bool {

}

// MoveObject implements
func MoveObject(objectid int, X float32, Y float32, Z float32, Speed float32, RotX float32, RotY float32, RotZ float32) int {

}

// StopObject implements
func StopObject(objectid int) bool {

}

// IsObjectMoving implements
func IsObjectMoving(objectid int) bool {

}

// EditObject implements
func EditObject(playerid int, objectid int) bool {

}

// EditPlayerObject implements
func EditPlayerObject(playerid int, objectid int) bool {

}

// SelectObject implements
func SelectObject(playerid int) bool {

}

// CancelEdit implements
func CancelEdit(playerid int) bool {

}

// CreatePlayerObject implements
func CreatePlayerObject(playerid int, modelid int, x float32, y float32, z float32, rX float32, rY float32, rZ float32, DrawDistance float32) int {

}

// AttachPlayerObjectToPlayer implements
func AttachPlayerObjectToPlayer(objectplayer int, objectid int, attachplayer int, OffsetX float32, OffsetY float32, OffsetZ float32, rX float32, rY float32, rZ float32) bool {

}

// AttachPlayerObjectToVehicle implements
func AttachPlayerObjectToVehicle(playerid int, objectid int, vehicleid int, fOffsetX float32, fOffsetY float32, fOffsetZ float32, fRotX float32, fRotY float32, RotZ float32) bool {

}

// SetPlayerObjectPos implements
func SetPlayerObjectPos(playerid int, objectid int, x float32, y float32, z float32) bool {

}

// GetPlayerObjectPos implements
func GetPlayerObjectPos(playerid int, objectid int, x float32, y float32, z float32) bool {

}

// SetPlayerObjectRot implements
func SetPlayerObjectRot(playerid int, objectid int, rotX float32, rotY float32, rotZ float32) bool {

}

// GetPlayerObjectRot implements
func GetPlayerObjectRot(playerid int, objectid int, rotX float32, rotY float32, rotZ float32) bool {

}

// GetPlayerObjectModel implements
func GetPlayerObjectModel(playerid int, objectid int) int {

}

// SetPlayerObjectNoCameraCol implements
func SetPlayerObjectNoCameraCol(playerid int, objectid int) bool {

}

// IsValidPlayerObject implements
func IsValidPlayerObject(playerid int, objectid int) bool {

}

// DestroyPlayerObject implements
func DestroyPlayerObject(playerid int, objectid int) bool {

}

// MovePlayerObject implements
func MovePlayerObject(playerid int, objectid int, x float32, y float32, z float32, Speed float32, RotX float32, RotY float32, RotZ float32) int {

}

// StopPlayerObject implements
func StopPlayerObject(playerid int, objectid int) bool {

}

// IsPlayerObjectMoving implements
func IsPlayerObjectMoving(playerid int, objectid int) bool {

}

// SetObjectMaterial implements
func SetObjectMaterial(objectid int, materialindex int, modelid int, txdname string, texturename string, materialcolor int) bool {

}

// SetPlayerObjectMaterial implements
func SetPlayerObjectMaterial(playerid int, objectid int, materialindex int, modelid int, txdname string, texturename string, materialcolor int) bool {

}

// SetObjectMaterialText implements
func SetObjectMaterialText(objectid int, text string, materialindex int, materialsize int, fontface string, fontsize int, bold bool, fontcolor int, backcolor int, textalignment int) bool {

}

// SetPlayerObjectMaterialText implements
func SetPlayerObjectMaterialText(playerid int, objectid int, text string, materialindex int, materialsize int, fontface string, fontsize int, bold bool, fontcolor int, backcolor int, textalignment int) bool {

}

// SetObjectsDefaultCameraCol implements
func SetObjectsDefaultCameraCol(disable bool) bool {

}

// SetSpawnInfo implements
func SetSpawnInfo(playerid int, team int, skin int, x float32, y float32, z float32, rotation float32, weapon1 int, weapon1_ammo int, weapon2 int, weapon2_ammo int, weapon3 int, weapon3_ammo int) bool {

}*/

// SpawnPlayer spawns the player.
func SpawnPlayer(playerid int) bool {
	return bool(C.SpawnPlayer(C.int(playerid)))
}

// SetPlayerPos sets the player's current position.
func SetPlayerPos(playerid int, x float32, y float32, z float32) bool {
	return bool(C.SetPlayerPos(C.int(playerid), C.float(x), C.float(y), C.float(z)))
}

/*
// SetPlayerPosFindZ implements
func SetPlayerPosFindZ(playerid int, x float32, y float32, z float32) bool {

}*/

// GetPlayerPos returns the player's current position.
func GetPlayerPos(playerid int) (float32, float32, float32) {
	var x, y, z C.float
	C.GetPlayerPos(C.int(playerid), &x, &y, &z)
	return float32(x), float32(y), float32(z)
}

// SetPlayerFacingAngle sets the player's facing angle.
func SetPlayerFacingAngle(playerid int, angle float32) bool {
	return bool(C.SetPlayerFacingAngle(C.int(playerid), C.float(angle)))
}

// GetPlayerFacingAngle gets the player's facing angle.
func GetPlayerFacingAngle(playerid int) float32 {
	var a C.float
	C.GetPlayerFacingAngle(C.int(playerid), &a)
	return float32(a)
}

// IsPlayerInRangeOfPoint returns whether a player is or isnt in range of a set of coordinates.
func IsPlayerInRangeOfPoint(playerid int, pRange, x, y, z float32) bool {
	return bool(C.IsPlayerInRangeOfPoint(C.int(playerid), C.float(pRange), C.float(x), C.float(y), C.float(z)))
}

/*
// GetPlayerDistanceFromPoimplements int
func GetPlayerDistanceFromPoint(playerid int, x float32, y float32, z float32) float32 {

}

// IsPlayerStreamedIn implements
func IsPlayerStreamedIn(playerid int, forplayerid int) bool {

}

// SetPlayerInterior implements
func SetPlayerInterior(playerid int, interiorid int) bool {

}

// GetPlayerInterior implements
func GetPlayerInterior(playerid int) int {

}

// SetPlayerHealth implements
func SetPlayerHealth(playerid int, health float32) bool {

}

// GetPlayerHealth implements
func GetPlayerHealth(playerid int, health float32) bool {

}

// SetPlayerArmour implements
func SetPlayerArmour(playerid int, armour float32) bool {

}

// GetPlayerArmour implements
func GetPlayerArmour(playerid int, armour float32) bool {

}

// SetPlayerAmmo implements
func SetPlayerAmmo(playerid int, weaponid int, ammo int) bool {

}

// GetPlayerAmmo implements
func GetPlayerAmmo(playerid int) int {

}

// GetPlayerWeaponState implements
func GetPlayerWeaponState(playerid int) int {

}

// GetPlayerTargetPlayer implements
func GetPlayerTargetPlayer(playerid int) int {

}

// GetPlayerTargetActor implements
func GetPlayerTargetActor(playerid int) int {

}

// SetPlayerTeam implements
func SetPlayerTeam(playerid int, teamid int) bool {

}

// GetPlayerTeam implements
func GetPlayerTeam(playerid int) int {

}

// SetPlayerScore implements
func SetPlayerScore(playerid int, score int) bool {

}

// GetPlayerScore implements
func GetPlayerScore(playerid int) int {

}

// GetPlayerDrunkLevel implements
func GetPlayerDrunkLevel(playerid int) int {

}

// SetPlayerDrunkLevel implements
func SetPlayerDrunkLevel(playerid int, level int) bool {

}

// SetPlayerColor implements
func SetPlayerColor(playerid int, color int) bool {

}

// GetPlayerColor implements
func GetPlayerColor(playerid int) int {

}

// SetPlayerSkin implements
func SetPlayerSkin(playerid int, skinid int) bool {

}

// GetPlayerSkin implements
func GetPlayerSkin(playerid int) int {

}

// GivePlayerWeapon implements
func GivePlayerWeapon(playerid int, weaponid int, ammo int) bool {

}

// ResetPlayerWeapons implements
func ResetPlayerWeapons(playerid int) bool {

}

// SetPlayerArmedWeapon implements
func SetPlayerArmedWeapon(playerid int, weaponid int) bool {

}

// GetPlayerWeaponData implements
func GetPlayerWeaponData(playerid int, slot int, weapon int, ammo int) bool {

}

// GivePlayerMoney implements
func GivePlayerMoney(playerid int, money int) bool {

}

// ResetPlayerMoney implements
func ResetPlayerMoney(playerid int) bool {

}

// SetPlayerName implements
func SetPlayerName(playerid int, name string) int {

}

// GetPlayerMoney implements
func GetPlayerMoney(playerid int) int {

}

// GetPlayerState implements
func GetPlayerState(playerid int) int {

}

// GetPlayerIp implements
func GetPlayerIp(playerid int, ip string, size int) bool {

}

// GetPlayerPing implements
func GetPlayerPing(playerid int) int {

}

// GetPlayerWeapon implements
func GetPlayerWeapon(playerid int) int {

}

// GetPlayerKeys implements
func GetPlayerKeys(playerid int, keys int, updown int, leftright int) bool {

}

// GetPlayerName implements
func GetPlayerName(playerid int, name string, size int) int {

}

// SetPlayerTime implements
func SetPlayerTime(playerid int, hour int, minute int) bool {

}

// GetPlayerTime implements
func GetPlayerTime(playerid int, hour int, minute int) bool {

}

// TogglePlayerClock implements
func TogglePlayerClock(playerid int, toggle bool) bool {

}

// SetPlayerWeather implements
func SetPlayerWeather(playerid int, weather int) bool {

}

// ForceClassSelection implements
func ForceClassSelection(playerid int) bool {

}

// SetPlayerWantedLevel implements
func SetPlayerWantedLevel(playerid int, level int) bool {

}

// GetPlayerWantedLevel implements
func GetPlayerWantedLevel(playerid int) int {

}

// SetPlayerFightingStyle implements
func SetPlayerFightingStyle(playerid int, style int) bool {

}

// GetPlayerFightingStyle implements
func GetPlayerFightingStyle(playerid int) int {

}

// SetPlayerVelocity implements
func SetPlayerVelocity(playerid int, x float32, y float32, z float32) bool {

}

// GetPlayerVelocity implements
func GetPlayerVelocity(playerid int, x float32, y float32, z float32) bool {

}

// PlayCrimeReportForPlayer implements
func PlayCrimeReportForPlayer(playerid int, suspectid int, crime int) bool {

}

// PlayAudioStreamForPlayer implements
func PlayAudioStreamForPlayer(playerid int, url string, posX float32, posY float32, posZ float32, distance float32, usepos bool) bool {

}

// StopAudioStreamForPlayer implements
func StopAudioStreamForPlayer(playerid int) bool {

}

// SetPlayerShopName implements
func SetPlayerShopName(playerid int, shopname string) bool {

}

// SetPlayerSkillLevel implements
func SetPlayerSkillLevel(playerid int, skill int, level int) bool {

}

// GetPlayerSurfingVehicleID implements
func GetPlayerSurfingVehicleID(playerid int) int {

}

// GetPlayerSurfingObjectID implements
func GetPlayerSurfingObjectID(playerid int) int {

}

// RemoveBuildingForPlayer implements
func RemoveBuildingForPlayer(playerid int, modelid int, fX float32, fY float32, fZ float32, fRadius float32) bool {

}

// GetPlayerLastShotVectors implements
func GetPlayerLastShotVectors(playerid int, fOriginX float32, fOriginY float32, fOriginZ float32, fHitPosX float32, fHitPosY float32, fHitPosZ float32) bool {

}

// SetPlayerAttachedObject implements
func SetPlayerAttachedObject(playerid int, index int, modelid int, bone int, fOffsetX float32, fOffsetY float32, fOffsetZ float32, fRotX float32, fRotY float32, fRotZ float32, fScaleX float32, fScaleY float32, fScaleZ float32, materialcolor1 int, materialcolor2 int) bool {

}

// RemovePlayerAttachedObject implements
func RemovePlayerAttachedObject(playerid int, index int) bool {

}

// IsPlayerAttachedObjectSlotUsed implements
func IsPlayerAttachedObjectSlotUsed(playerid int, index int) bool {

}

// EditAttachedObject implements
func EditAttachedObject(playerid int, index int) bool {

}

// CreatePlayerTextDraw implements
func CreatePlayerTextDraw(playerid int, x float32, y float32, text string) int {

}

// PlayerTextDrawDestroy implements
func PlayerTextDrawDestroy(playerid int, text int) bool {

}

// PlayerTextDrawLetterSize implements
func PlayerTextDrawLetterSize(playerid int, text int, x float32, y float32) bool {

}

// PlayerTextDrawTextSize implements
func PlayerTextDrawTextSize(playerid int, text int, x float32, y float32) bool {

}

// PlayerTextDrawAlignment implements
func PlayerTextDrawAlignment(playerid int, text int, alignment int) bool {

}

// PlayerTextDrawColor implements
func PlayerTextDrawColor(playerid int, text int, color int) bool {

}

// PlayerTextDrawUseBox implements
func PlayerTextDrawUseBox(playerid int, text int, use bool) bool {

}

// PlayerTextDrawBoxColor implements
func PlayerTextDrawBoxColor(playerid int, text int, color int) bool {

}

// PlayerTextDrawSetShadow implements
func PlayerTextDrawSetShadow(playerid int, text int, size int) bool {

}

// PlayerTextDrawSetOutline implements
func PlayerTextDrawSetOutline(playerid int, text int, size int) bool {

}

// PlayerTextDrawBackgroundColor implements
func PlayerTextDrawBackgroundColor(playerid int, text int, color int) bool {

}

// PlayerTextDrawFont implements
func PlayerTextDrawFont(playerid int, text int, font int) bool {

}

// PlayerTextDrawSetProportional implements
func PlayerTextDrawSetProportional(playerid int, text int, set bool) bool {

}

// PlayerTextDrawSetSelectable implements
func PlayerTextDrawSetSelectable(playerid int, text int, set bool) bool {

}

// PlayerTextDrawShow implements
func PlayerTextDrawShow(playerid int, text int) bool {

}

// PlayerTextDrawHide implements
func PlayerTextDrawHide(playerid int, text int) bool {

}

// PlayerTextDrawSetString implements
func PlayerTextDrawSetString(playerid int, text int, string string) bool {

}

// PlayerTextDrawSetPreviewModel implements
func PlayerTextDrawSetPreviewModel(playerid int, text int, modelindex int) bool {

}

// PlayerTextDrawSetPreviewRot implements
func PlayerTextDrawSetPreviewRot(playerid int, text int, fRotX float32, fRotY float32, fRotZ float32, fZoom float32) bool {

}

// PlayerTextDrawSetPreviewVehCol implements
func PlayerTextDrawSetPreviewVehCol(playerid int, text int, color1 int, color2 int) bool {

}

// SetPlayerChatBubble implements
func SetPlayerChatBubble(playerid int, text string, color int, drawdistance float32, expiretime int) bool {

}

// PutPlayerInVehicle implements
func PutPlayerInVehicle(playerid int, vehicleid int, seatid int) bool {

}

// GetPlayerVehicleID implements
func GetPlayerVehicleID(playerid int) int {

}

// GetPlayerVehicleSeat implements
func GetPlayerVehicleSeat(playerid int) int {

}

// RemovePlayerFromVehicle implements
func RemovePlayerFromVehicle(playerid int) bool {

}

// TogglePlayerControllable implements
func TogglePlayerControllable(playerid int, toggle bool) bool {

}

// PlayerPlaySound implements
func PlayerPlaySound(playerid int, soundid int, x float32, y float32, z float32) bool {

}

// ApplyAnimation implements
func ApplyAnimation(playerid int, animlib string, animname string, fDelta float32, loop bool, lockx bool, locky bool, freeze bool, time int, forcesync bool) bool {

}

// ClearAnimations implements
func ClearAnimations(playerid int, forcesync bool) bool {

}

// GetPlayerAnimationIndex implements
func GetPlayerAnimationIndex(playerid int) int {

}

// GetAnimationName implements
func GetAnimationName(index int, animlib string, animlib_size int, animname string, animname_size int) bool {

}

// GetPlayerSpecialAction implements
func GetPlayerSpecialAction(playerid int) int {

}

// SetPlayerSpecialAction implements
func SetPlayerSpecialAction(playerid int, actionid int) bool {

}

// DisableRemoteVehicleCollisions implements
func DisableRemoteVehicleCollisions(playerid int, disable bool) bool {

}

// SetPlayerCheckpoimplements int
func SetPlayerCheckpoint(playerid int, x float32, y float32, z float32, size float32) bool {

}

// DisablePlayerCheckpoimplements int
func DisablePlayerCheckpoint(playerid int) bool {

}

// SetPlayerRaceCheckpoimplements int
func SetPlayerRaceCheckpoint(playerid int, cType int, x float32, y float32, z float32, nextx float32, nexty float32, nextz float32, size float32) bool {

}

// DisablePlayerRaceCheckpoimplements int
func DisablePlayerRaceCheckpoint(playerid int) bool {

}

// SetPlayerWorldBounds implements
func SetPlayerWorldBounds(playerid int, x_max float32, x_min float32, y_max float32, y_min float32) bool {

}

// SetPlayerMarkerForPlayer implements
func SetPlayerMarkerForPlayer(playerid int, showplayerid int, color int) bool {

}

// ShowPlayerNameTagForPlayer implements
func ShowPlayerNameTagForPlayer(playerid int, showplayerid int, show bool) bool {

}

// SetPlayerMapIcon implements
func SetPlayerMapIcon(playerid int, iconid int, x float32, y float32, z float32, markertype int, color int, style int) bool {

}

// RemovePlayerMapIcon implements
func RemovePlayerMapIcon(playerid int, iconid int) bool {

}

// AllowPlayerTeleport implements
func AllowPlayerTeleport(playerid int, allow bool) bool {

}

// SetPlayerCameraPos implements
func SetPlayerCameraPos(playerid int, x float32, y float32, z float32) bool {

}

// SetPlayerCameraLookAt implements
func SetPlayerCameraLookAt(playerid int, x float32, y float32, z float32, cut int) bool {

}

// SetCameraBehindPlayer implements
func SetCameraBehindPlayer(playerid int) bool {

}

// GetPlayerCameraPos implements
func GetPlayerCameraPos(playerid int, x float32, y float32, z float32) bool {

}

// GetPlayerCameraFrontVector implements
func GetPlayerCameraFrontVector(playerid int, x float32, y float32, z float32) bool {

}

// GetPlayerCameraMode implements
func GetPlayerCameraMode(playerid int) int {

}

// EnablePlayerCameraTarget implements
func EnablePlayerCameraTarget(playerid int, enable bool) bool {

}

// GetPlayerCameraTargetObject implements
func GetPlayerCameraTargetObject(playerid int) int {

}

// GetPlayerCameraTargetVehicle implements
func GetPlayerCameraTargetVehicle(playerid int) int {

}

// GetPlayerCameraTargetPlayer implements
func GetPlayerCameraTargetPlayer(playerid int) int {

}

// GetPlayerCameraTargetActor implements
func GetPlayerCameraTargetActor(playerid int) int {

}

// GetPlayerCameraAspectRatio implements
func GetPlayerCameraAspectRatio(playerid int) float32 {

}

// GetPlayerCameraZoom implements
func GetPlayerCameraZoom(playerid int) float32 {

}

// AttachCameraToObject implements
func AttachCameraToObject(playerid int, objectid int) bool {

}

// AttachCameraToPlayerObject implements
func AttachCameraToPlayerObject(playerid int, playerobjectid int) bool {

}

// InterpolateCameraPos implements
func InterpolateCameraPos(playerid int, FromX float32, FromY float32, FromZ float32, ToX float32, ToY float32, ToZ float32, time int, cut int) bool {

}

// InterpolateCameraLookAt implements
func InterpolateCameraLookAt(playerid int, FromX float32, FromY float32, FromZ float32, ToX float32, ToY float32, ToZ float32, time int, cut int) bool {

}

// IsPlayerConnected implements
func IsPlayerConnected(playerid int) bool {

}

// IsPlayerInVehicle implements
func IsPlayerInVehicle(playerid int, vehicleid int) bool {

}

// IsPlayerInAnyVehicle implements
func IsPlayerInAnyVehicle(playerid int) bool {

}

// IsPlayerInCheckpoimplements int
func IsPlayerInCheckpoint(playerid int) bool {

}

// IsPlayerInRaceCheckpoimplements int
func IsPlayerInRaceCheckpoint(playerid int) bool {

}

// SetPlayerVirtualWorld implements
func SetPlayerVirtualWorld(playerid int, worldid int) bool {

}

// GetPlayerVirtualWorld implements
func GetPlayerVirtualWorld(playerid int) int {

}

// EnableStuntBonusForPlayer implements
func EnableStuntBonusForPlayer(playerid int, enable bool) bool {

}

// EnableStuntBonusForAll implements
func EnableStuntBonusForAll(enable bool) bool {

}

// TogglePlayerSpectating implements
func TogglePlayerSpectating(playerid int, toggle bool) bool {

}

// PlayerSpectatePlayer implements
func PlayerSpectatePlayer(playerid int, targetplayerid int, mode int) bool {

}

// PlayerSpectateVehicle implements
func PlayerSpectateVehicle(playerid int, targetvehicleid int, mode int) bool {

}

// StartRecordingPlayerData implements
func StartRecordingPlayerData(playerid int, recordtype int, recordname string) bool {

}

// StopRecordingPlayerData implements
func StopRecordingPlayerData(playerid int) bool {

}

// CreateExplosionForPlayer implements
func CreateExplosionForPlayer(playerid int, X float32, Y float32, Z float32, eType int, Radius float32) bool {

}
*/
