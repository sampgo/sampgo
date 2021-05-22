// -
// sampgo
// interop.h
// -

#include "main.h"

AMX_NATIVE_INFO native_list[] = {
	{ "sampgo_CallEvent", n_CallEvent },
	{ NULL, NULL }
};

// GoInt32: sampgo_CallEvent(const event[32], const format[], {Float,_}:...);
cell AMX_NATIVE_CALL n_CallEvent(AMX* amx, cell* params)
{
    int
        len = (int) NULL
    ;

    cell *addr  = NULL;

    amx_GetAddr(amx, params[1], &addr);
    amx_StrLen(addr, &len);

    if (!len) {
        sampgdk_logprintf("(C) sampgo: Empty event name passed to n_CallEvent");
        return false;
    }

    ++ len;
    char* event = malloc( sizeof(char) * (len));
    amx_GetString(event, addr, 0, len);

    len = (int) NULL;

    amx_GetAddr(amx, params[2], &addr);
    amx_StrLen(addr, &len);

    ++ len;
    char* format = malloc( sizeof(char) * (len));
    amx_GetString(format, addr, 0, len);

    bool retval = callEvent(&amx, event, format, params);

    sampgdk_logprintf("(C) sampgo: Received event name (%s) with format (%s) and retval (%i)", event, format, retval);

    free(event);
    free(format);
    return retval;
}

/**
 * \ingroup callbacks
 * \see <a href="http://wiki.sa-mp.com/wiki/OnGameModeInit">OnGameModeInit on SA-MP Wiki</a>
 */
PLUGIN_EXPORT bool PLUGIN_CALL OnGameModeInit()
{
    onGameModeInit();
    return true;
}

/**
 * \ingroup callbacks
 * \see <a href="http://wiki.sa-mp.com/wiki/OnGameModeExit">OnGameModeExit on SA-MP Wiki</a>
 */
PLUGIN_EXPORT bool PLUGIN_CALL OnGameModeExit()
{
    onGameModeExit();
    return true;
}

/**
 * \ingroup callbacks
 * \see <a href="http://wiki.sa-mp.com/wiki/OnPlayerConnect">OnPlayerConnect on SA-MP Wiki</a>
 */
PLUGIN_EXPORT bool PLUGIN_CALL OnPlayerConnect(int playerid)
{
    onPlayerConnect(playerid);
    return true;
}

/**
 * \ingroup callbacks
 * \see <a href="http://wiki.sa-mp.com/wiki/OnPlayerDisconnect">OnPlayerDisconnect on SA-MP Wiki</a>
 */
PLUGIN_EXPORT bool PLUGIN_CALL OnPlayerDisconnect(int playerid, int reason)
{
    onPlayerDisconnect(playerid, reason);
    return true;
}

/**
 * \ingroup callbacks
 * \see <a href="http://wiki.sa-mp.com/wiki/OnPlayerSpawn">OnPlayerSpawn on SA-MP Wiki</a>
 */
PLUGIN_EXPORT bool PLUGIN_CALL OnPlayerSpawn(int playerid)
{
    onPlayerSpawn(playerid);
    return true;
}

/**
 * \ingroup callbacks
 * \see <a href="http://wiki.sa-mp.com/wiki/OnPlayerDeath">OnPlayerDeath on SA-MP Wiki</a>
 */
PLUGIN_EXPORT bool PLUGIN_CALL OnPlayerDeath(int playerid, int killerid, int reason)
{
    onPlayerDeath(playerid, killerid, reason);
    return true;
}

/**
 * \ingroup callbacks
 * \see <a href="http://wiki.sa-mp.com/wiki/OnVehicleSpawn">OnVehicleSpawn on SA-MP Wiki</a>
 */
PLUGIN_EXPORT bool PLUGIN_CALL OnVehicleSpawn(int vehicleid)
{
    onVehicleSpawn(vehicleid);
    return true;
}

/**
 * \ingroup callbacks
 * \see <a href="http://wiki.sa-mp.com/wiki/OnVehicleDeath">OnVehicleDeath on SA-MP Wiki</a>
 */
PLUGIN_EXPORT bool PLUGIN_CALL OnVehicleDeath(int vehicleid, int killerid)
{
    onVehicleDeath(vehicleid, killerid);
    return true;
}

/**
 * \ingroup callbacks
 * \see <a href="http://wiki.sa-mp.com/wiki/OnPlayerText">OnPlayerText on SA-MP Wiki</a>
 */
PLUGIN_EXPORT bool PLUGIN_CALL OnPlayerText(int playerid, const char* text)
{
    return onPlayerText(playerid, text);
}

/**
 * \ingroup callbacks
 * \see <a href="http://wiki.sa-mp.com/wiki/OnPlayerCommandText">OnPlayerCommandText on SA-MP Wiki</a>
 */
PLUGIN_EXPORT bool PLUGIN_CALL OnPlayerCommandText(int playerid, const char* cmdtext)
{
    return onPlayerCommandText(playerid, cmdtext);
}

/**
 * \ingroup callbacks
 * \see <a href="http://wiki.sa-mp.com/wiki/OnPlayerRequestClass">OnPlayerRequestClass on SA-MP Wiki</a>
 */
PLUGIN_EXPORT bool PLUGIN_CALL OnPlayerRequestClass(int playerid, int classid)
{
    onPlayerRequestClass(playerid, classid);
    return true;
}

/**
 * \ingroup callbacks
 * \see <a href="http://wiki.sa-mp.com/wiki/OnPlayerEnterVehicle">OnPlayerEnterVehicle on SA-MP Wiki</a>
 */
PLUGIN_EXPORT bool PLUGIN_CALL OnPlayerEnterVehicle(int playerid, int vehicleid, bool ispassenger)
{
    onPlayerEnterVehicle(playerid, vehicleid, ispassenger);
    return true;
}

/**
 * \ingroup callbacks
 * \see <a href="http://wiki.sa-mp.com/wiki/OnPlayerExitVehicle">OnPlayerExitVehicle on SA-MP Wiki</a>
 */
PLUGIN_EXPORT bool PLUGIN_CALL OnPlayerExitVehicle(int playerid, int vehicleid)
{
    onPlayerExitVehicle(playerid, vehicleid);
    return true;
}

/**
 * \ingroup callbacks
 * \see <a href="http://wiki.sa-mp.com/wiki/OnPlayerStateChange">OnPlayerStateChange on SA-MP Wiki</a>
 */
PLUGIN_EXPORT bool PLUGIN_CALL OnPlayerStateChange(int playerid, int newstate, int oldstate)
{
    onPlayerStateChange(playerid, newstate, oldstate);
    return true;
}

/**
 * \ingroup callbacks
 * \see <a href="http://wiki.sa-mp.com/wiki/OnPlayerEnterCheckpoint">OnPlayerEnterCheckpoint on SA-MP Wiki</a>
 */
PLUGIN_EXPORT bool PLUGIN_CALL OnPlayerEnterCheckpoint(int playerid)
{
    onPlayerEnterCheckpoint(playerid);
    return true;
}

/**
 * \ingroup callbacks
 * \see <a href="http://wiki.sa-mp.com/wiki/OnPlayerLeaveCheckpoint">OnPlayerLeaveCheckpoint on SA-MP Wiki</a>
 */
PLUGIN_EXPORT bool PLUGIN_CALL OnPlayerLeaveCheckpoint(int playerid)
{
    onPlayerLeaveCheckpoint(playerid);
    return true;
}

/**
 * \ingroup callbacks
 * \see <a href="http://wiki.sa-mp.com/wiki/OnPlayerEnterRaceCheckpoint">OnPlayerEnterRaceCheckpoint on SA-MP Wiki</a>
 */
PLUGIN_EXPORT bool PLUGIN_CALL OnPlayerEnterRaceCheckpoint(int playerid)
{
    onPlayerEnterRaceCheckpoint(playerid);
    return true;
}

/**
 * \ingroup callbacks
 * \see <a href="http://wiki.sa-mp.com/wiki/OnPlayerLeaveRaceCheckpoint">OnPlayerLeaveRaceCheckpoint on SA-MP Wiki</a>
 */
PLUGIN_EXPORT bool PLUGIN_CALL OnPlayerLeaveRaceCheckpoint(int playerid)
{
    onPlayerLeaveRaceCheckpoint(playerid);
    return true;
}

/**
 * \ingroup callbacks
 * \see <a href="http://wiki.sa-mp.com/wiki/OnRconCommand">OnRconCommand on SA-MP Wiki</a>
 */
PLUGIN_EXPORT bool PLUGIN_CALL OnRconCommand(const char* cmd)
{
    onRconCommand(cmd);
    return true;
}

/**
 * \ingroup callbacks
 * \see <a href="http://wiki.sa-mp.com/wiki/OnPlayerRequestSpawn">OnPlayerRequestSpawn on SA-MP Wiki</a>
 */
PLUGIN_EXPORT bool PLUGIN_CALL OnPlayerRequestSpawn(int playerid)
{
    onPlayerRequestSpawn(playerid);
    return true;
}

/**
 * \ingroup callbacks
 * \see <a href="http://wiki.sa-mp.com/wiki/OnObjectMoved">OnObjectMoved on SA-MP Wiki</a>
 */
PLUGIN_EXPORT bool PLUGIN_CALL OnObjectMoved(int objectid)
{
    onObjectMoved(objectid);
    return true;
}

/**
 * \ingroup callbacks
 * \see <a href="http://wiki.sa-mp.com/wiki/OnPlayerObjectMoved">OnPlayerObjectMoved on SA-MP Wiki</a>
 */
PLUGIN_EXPORT bool PLUGIN_CALL OnPlayerObjectMoved(int playerid, int objectid)
{
    onPlayerObjectMoved(playerid, objectid);
    return true;
}

/**
 * \ingroup callbacks
 * \see <a href="http://wiki.sa-mp.com/wiki/OnPlayerPickUpPickup">OnPlayerPickUpPickup on SA-MP Wiki</a>
 */
PLUGIN_EXPORT bool PLUGIN_CALL OnPlayerPickUpPickup(int playerid, int pickupid)
{
    onPlayerPickUpPickup(playerid, pickupid);
    return true;
}

/**
 * \ingroup callbacks
 * \see <a href="http://wiki.sa-mp.com/wiki/OnVehicleMod">OnVehicleMod on SA-MP Wiki</a>
 */
PLUGIN_EXPORT bool PLUGIN_CALL OnVehicleMod(int playerid, int vehicleid, int componentid)
{
    onVehicleMod(playerid, vehicleid, componentid);
    return true;
}

/**
 * \ingroup callbacks
 * \see <a href="http://wiki.sa-mp.com/wiki/OnEnterExitModShop">OnEnterExitModShop on SA-MP Wiki</a>
 */
PLUGIN_EXPORT bool PLUGIN_CALL OnEnterExitModShop(int playerid, bool enterexit, int interiorid)
{
    onEnterExitModShop(playerid, enterexit, interiorid);
    return true;
}

/**
 * \ingroup callbacks
 * \see <a href="http://wiki.sa-mp.com/wiki/OnVehiclePaintjob">OnVehiclePaintjob on SA-MP Wiki</a>
 */
PLUGIN_EXPORT bool PLUGIN_CALL OnVehiclePaintjob(int playerid, int vehicleid, int paintjobid)
{
    onVehiclePaintjob(playerid, vehicleid, paintjobid);
    return true;
}

/**
 * \ingroup callbacks
 * \see <a href="http://wiki.sa-mp.com/wiki/OnVehicleRespray">OnVehicleRespray on SA-MP Wiki</a>
 */
PLUGIN_EXPORT bool PLUGIN_CALL OnVehicleRespray(int playerid, int vehicleid, int color1, int color2)
{
    onVehicleRespray(playerid, vehicleid, color1, color2);
    return true;
}

/**
 * \ingroup callbacks
 * \see <a href="http://wiki.sa-mp.com/wiki/OnVehicleDamageStatusUpdate">OnVehicleDamageStatusUpdate on SA-MP Wiki</a>
 */
PLUGIN_EXPORT bool PLUGIN_CALL OnVehicleDamageStatusUpdate(int vehicleid, int playerid)
{
    onVehicleDamageStatusUpdate(vehicleid, playerid);
    return true;
}

/**
 * \ingroup callbacks
 * \see <a href="http://wiki.sa-mp.com/wiki/OnUnoccupiedVehicleUpdate">OnUnoccupiedVehicleUpdate on SA-MP Wiki</a>
 */
PLUGIN_EXPORT bool PLUGIN_CALL OnUnoccupiedVehicleUpdate(int vehicleid, int playerid, int passenger_seat, float new_x, float new_y, float new_z, float vel_x, float vel_y, float vel_z)
{
    onUnoccupiedVehicleUpdate(vehicleid, playerid, passenger_seat, new_x, new_y, new_z, vel_x, vel_y, vel_z);
    return true;
}

/**
 * \ingroup callbacks
 * \see <a href="http://wiki.sa-mp.com/wiki/OnPlayerSelectedMenuRow">OnPlayerSelectedMenuRow on SA-MP Wiki</a>
 */
PLUGIN_EXPORT bool PLUGIN_CALL OnPlayerSelectedMenuRow(int playerid, int row)
{
    onPlayerSelectedMenuRow(playerid, row);
    return true;
}

/**
 * \ingroup callbacks
 * \see <a href="http://wiki.sa-mp.com/wiki/OnPlayerExitedMenu">OnPlayerExitedMenu on SA-MP Wiki</a>
 */
PLUGIN_EXPORT bool PLUGIN_CALL OnPlayerExitedMenu(int playerid)
{
    onPlayerExitedMenu(playerid);
    return true;
}

/**
 * \ingroup callbacks
 * \see <a href="http://wiki.sa-mp.com/wiki/OnPlayerInteriorChange">OnPlayerInteriorChange on SA-MP Wiki</a>
 */
PLUGIN_EXPORT bool PLUGIN_CALL OnPlayerInteriorChange(int playerid, int newinteriorid, int oldinteriorid)
{
    onPlayerInteriorChange(playerid, newinteriorid, oldinteriorid);
    return true;
}

/**
 * \ingroup callbacks
 * \see <a href="http://wiki.sa-mp.com/wiki/OnPlayerKeyStateChange">OnPlayerKeyStateChange on SA-MP Wiki</a>
 */
PLUGIN_EXPORT bool PLUGIN_CALL OnPlayerKeyStateChange(int playerid, int newkeys, int oldkeys)
{
    onPlayerKeyStateChange(playerid, newkeys, oldkeys);
    return true;
}

/**
 * \ingroup callbacks
 * \see <a href="http://wiki.sa-mp.com/wiki/OnRconLoginAttempt">OnRconLoginAttempt on SA-MP Wiki</a>
 */
PLUGIN_EXPORT bool PLUGIN_CALL OnRconLoginAttempt(const char* ip, const char* password, bool success)
{
    onRconLoginAttempt(ip, password, success);
    return true;
}

/**
 * \ingroup callbacks
 * \see <a href="http://wiki.sa-mp.com/wiki/OnPlayerUpdate">OnPlayerUpdate on SA-MP Wiki</a>
 */
PLUGIN_EXPORT bool PLUGIN_CALL OnPlayerUpdate(int playerid)
{
    onPlayerUpdate(playerid);
    return true;
}

/**
 * \ingroup callbacks
 * \see <a href="http://wiki.sa-mp.com/wiki/OnPlayerStreamIn">OnPlayerStreamIn on SA-MP Wiki</a>
 */
PLUGIN_EXPORT bool PLUGIN_CALL OnPlayerStreamIn(int playerid, int forplayerid)
{
    onPlayerStreamIn(playerid, forplayerid);
    return true;
}

/**
 * \ingroup callbacks
 * \see <a href="http://wiki.sa-mp.com/wiki/OnPlayerStreamOut">OnPlayerStreamOut on SA-MP Wiki</a>
 */
PLUGIN_EXPORT bool PLUGIN_CALL OnPlayerStreamOut(int playerid, int forplayerid)
{
    onPlayerStreamOut(playerid, forplayerid);
    return true;
}

/**
 * \ingroup callbacks
 * \see <a href="http://wiki.sa-mp.com/wiki/OnVehicleStreamIn">OnVehicleStreamIn on SA-MP Wiki</a>
 */
PLUGIN_EXPORT bool PLUGIN_CALL OnVehicleStreamIn(int vehicleid, int forplayerid)
{
    onVehicleStreamIn(vehicleid, forplayerid);
    return true;
}

/**
 * \ingroup callbacks
 * \see <a href="http://wiki.sa-mp.com/wiki/OnVehicleStreamOut">OnVehicleStreamOut on SA-MP Wiki</a>
 */
PLUGIN_EXPORT bool PLUGIN_CALL OnVehicleStreamOut(int vehicleid, int forplayerid)
{
    onVehicleStreamOut(vehicleid, forplayerid);
    return true;
}

/**
 * \ingroup callbacks
 * \see <a href="http://wiki.sa-mp.com/wiki/OnActorStreamIn">OnActorStreamIn on SA-MP Wiki</a>
 */
PLUGIN_EXPORT bool PLUGIN_CALL OnActorStreamIn(int actorid, int forplayerid)
{
    onActorStreamIn(actorid, forplayerid);
    return true;
}

/**
 * \ingroup callbacks
 * \see <a href="http://wiki.sa-mp.com/wiki/OnActorStreamOut">OnActorStreamOut on SA-MP Wiki</a>
 */
PLUGIN_EXPORT bool PLUGIN_CALL OnActorStreamOut(int actorid, int forplayerid)
{
    onActorStreamOut(actorid, forplayerid);
    return true;
}

/**
 * \ingroup callbacks
 * \see <a href="http://wiki.sa-mp.com/wiki/OnDialogResponse">OnDialogResponse on SA-MP Wiki</a>
 */
PLUGIN_EXPORT bool PLUGIN_CALL OnDialogResponse(int playerid, int dialogid, int response, int listitem, const char* inputtext)
{
    onDialogResponse(playerid, dialogid, response, listitem, inputtext);
    return true;
}

/**
 * \ingroup callbacks
 * \see <a href="http://wiki.sa-mp.com/wiki/OnPlayerTakeDamage">OnPlayerTakeDamage on SA-MP Wiki</a>
 */
PLUGIN_EXPORT bool PLUGIN_CALL OnPlayerTakeDamage(int playerid, int issuerid, float amount, int weaponid, int bodypart)
{
    onPlayerTakeDamage(playerid, issuerid, amount, weaponid, bodypart);
    return true;
}

/**
 * \ingroup callbacks
 * \see <a href="http://wiki.sa-mp.com/wiki/OnPlayerGiveDamage">OnPlayerGiveDamage on SA-MP Wiki</a>
 */
PLUGIN_EXPORT bool PLUGIN_CALL OnPlayerGiveDamage(int playerid, int damagedid, float amount, int weaponid, int bodypart)
{
    onPlayerGiveDamage(playerid, damagedid, amount, weaponid, bodypart);
    return true;
}

/**
 * \ingroup callbacks
 * \see <a href="http://wiki.sa-mp.com/wiki/OnPlayerGiveDamageActor">OnPlayerGiveDamageActor on SA-MP Wiki</a>
 */
PLUGIN_EXPORT bool PLUGIN_CALL OnPlayerGiveDamageActor(int playerid, int damaged_actorid, float amount, int weaponid, int bodypart)
{
    onPlayerGiveDamageActor(playerid, damaged_actorid, amount, weaponid, bodypart);
    return true;
}

/**
 * \ingroup callbacks
 * \see <a href="http://wiki.sa-mp.com/wiki/OnPlayerClickMap">OnPlayerClickMap on SA-MP Wiki</a>
 */
PLUGIN_EXPORT bool PLUGIN_CALL OnPlayerClickMap(int playerid, float fX, float fY, float fZ)
{
    onPlayerClickMap(playerid, fX, fY, fZ);
    return true;
}

/**
 * \ingroup callbacks
 * \see <a href="http://wiki.sa-mp.com/wiki/OnPlayerClickTextDraw">OnPlayerClickTextDraw on SA-MP Wiki</a>
 */
PLUGIN_EXPORT bool PLUGIN_CALL OnPlayerClickTextDraw(int playerid, int clickedid)
{
    onPlayerClickTextDraw(playerid, clickedid);
    return true;
}

/**
 * \ingroup callbacks
 * \see <a href="http://wiki.sa-mp.com/wiki/OnPlayerClickPlayerTextDraw">OnPlayerClickPlayerTextDraw on SA-MP Wiki</a>
 */
PLUGIN_EXPORT bool PLUGIN_CALL OnPlayerClickPlayerTextDraw(int playerid, int playertextid)
{
    onPlayerClickPlayerTextDraw(playerid, playertextid);
    return true;
}

/**
 * \ingroup callbacks
 * \see <a href="http://wiki.sa-mp.com/wiki/OnIncomingConnection">OnIncomingConnection on SA-MP Wiki</a>
 */
PLUGIN_EXPORT bool PLUGIN_CALL OnIncomingConnection(int playerid, const char* ip_address, int port)
{
    onIncomingConnection(playerid, ip_address, port);
    return true;
}

/**
 * \ingroup callbacks
 * \see <a href="http://wiki.sa-mp.com/wiki/OnTrailerUpdate">OnTrailerUpdate on SA-MP Wiki</a>
 */
PLUGIN_EXPORT bool PLUGIN_CALL OnTrailerUpdate(int playerid, int vehicleid)
{
    onTrailerUpdate(playerid, vehicleid);
    return true;
}

/**
 * \ingroup callbacks
 * \see <a href="http://wiki.sa-mp.com/wiki/OnVehicleSirenStateChange">OnVehicleSirenStateChange on SA-MP Wiki</a>
 */
PLUGIN_EXPORT bool PLUGIN_CALL OnVehicleSirenStateChange(int playerid, int vehicleid, int newstate)
{
    onVehicleSirenStateChange(playerid, vehicleid, newstate);
    return true;
}

/**
 * \ingroup callbacks
 * \see <a href="http://wiki.sa-mp.com/wiki/OnPlayerClickPlayer">OnPlayerClickPlayer on SA-MP Wiki</a>
 */
PLUGIN_EXPORT bool PLUGIN_CALL OnPlayerClickPlayer(int playerid, int clickedplayerid, int source)
{
    onPlayerClickPlayer(playerid, clickedplayerid, source);
    return true;
}

/**
 * \ingroup callbacks
 * \see <a href="http://wiki.sa-mp.com/wiki/OnPlayerEditObject">OnPlayerEditObject on SA-MP Wiki</a>
 */
PLUGIN_EXPORT bool PLUGIN_CALL OnPlayerEditObject(int playerid, bool playerobject, int objectid, int response, float fX, float fY, float fZ, float fRotX, float fRotY, float fRotZ)
{
    onPlayerEditObject(playerid, playerobject, objectid, response, fX, fY, fZ, fRotX, fRotY, fRotZ);
    return true;
}

/**
 * \ingroup callbacks
 * \see <a href="http://wiki.sa-mp.com/wiki/OnPlayerEditAttachedObject">OnPlayerEditAttachedObject on SA-MP Wiki</a>
 */
PLUGIN_EXPORT bool PLUGIN_CALL OnPlayerEditAttachedObject(int playerid, int response, int index, int modelid, int boneid, float fOffsetX, float fOffsetY, float fOffsetZ, float fRotX, float fRotY, float fRotZ, float fScaleX, float fScaleY, float fScaleZ)
{
    onPlayerEditAttachedObject(playerid, response, index, modelid, boneid, fOffsetX, fOffsetY, fOffsetZ, fRotX, fRotY, fRotZ, fScaleX, fScaleY, fScaleZ);
    return true;
}

/**
 * \ingroup callbacks
 * \see <a href="http://wiki.sa-mp.com/wiki/OnPlayerSelectObject">OnPlayerSelectObject on SA-MP Wiki</a>
 */
PLUGIN_EXPORT bool PLUGIN_CALL OnPlayerSelectObject(int playerid, int type, int objectid, int modelid, float fX, float fY, float fZ)
{
    onPlayerSelectObject(playerid, type, objectid, modelid, fX, fY, fZ);
    return true;
}

/**
 * \ingroup callbacks
 * \see <a href="http://wiki.sa-mp.com/wiki/OnPlayerWeaponShot">OnPlayerWeaponShot on SA-MP Wiki</a>
 */
PLUGIN_EXPORT bool PLUGIN_CALL OnPlayerWeaponShot(int playerid, int weaponid, int hittype, int hitid, float fX, float fY, float fZ)
{
    onPlayerWeaponShot(playerid, weaponid, hittype, hitid, fX, fY, fZ);
    return true;
}

/**
 * \ingroup callbacks
 * \see <a href="http://wiki.sa-mp.com/wiki/OnPlayerRequestDownload">OnPlayerRequestDownload on SA-MP Wiki</a>
 */
PLUGIN_EXPORT bool PLUGIN_CALL OnPlayerRequestDownload(int playerid, int type, int crc)
{
    onPlayerRequestDownload(playerid, type, crc);
    return true;
}

PLUGIN_EXPORT unsigned int PLUGIN_CALL Supports()
{
    return sampgdk_Supports() | SUPPORTS_PROCESS_TICK | SUPPORTS_AMX_NATIVES;
}

PLUGIN_EXPORT bool PLUGIN_CALL Load(void** ppData)
{
    sampgdk_Load(ppData, 0);
    return true;
}

PLUGIN_EXPORT void PLUGIN_CALL Unload()
{
    sampgdk_Unload(0);
}

PLUGIN_EXPORT void PLUGIN_CALL ProcessTick()
{
    onTick();
    sampgdk_ProcessTick(0);
}

PLUGIN_EXPORT int PLUGIN_CALL AmxLoad(AMX *amx) {
    return amx_Register(amx, native_list, -1);
}

PLUGIN_EXPORT int PLUGIN_CALL AmxUnload(AMX *amx) {
    return AMX_ERR_NONE;
}

void goLogprintf(char* text)
{
    sampgdk_logprintf((const char*)text);
}

char* constToNonConst(const char* text)
{
    return (char*)text;
}

const char* nonConstToConst(char* text)
{
    return (const char*)text;
}
