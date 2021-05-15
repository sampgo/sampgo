// -
// sampgo
// interop.h
// -

#ifndef SAMPGO_INTEROP_H
#define SAMPGO_INTEROP_H

#include "ugmp.h"
#include <sampgdk.h>

typedef const char char_t;

#ifndef GOLANG_APP
// All of the callbacks we want to export.

extern bool goModeInit();
extern bool goModeExit();

extern bool onPlayerConnect(int playerid);
extern bool onPlayerDisconnect(int playerid, int reason);

extern bool onPlayerSpawn(int playerid);
extern bool onPlayerDeath(int playerid, int killerid, int reason);

extern bool onVehicleSpawn(int vehicleid);
extern bool onVehicleDeath(int vehicleid, int killerid);

extern bool onPlayerText(int playerid, const char* text);
extern bool onPlayerCommandText(int playerid, const char* cmdtext);

extern bool onPlayerRequestClass(int playerid, int classid);

extern bool onPlayerEnterVehicle(int playerid, int vehicleid, bool ispassenger);
extern bool onPlayerExitVehicle(int playerid, int vehicleid);

extern bool onPlayerStateChange(int playerid, int newstate, int oldstate);

extern bool onPlayerEnterCheckpoint(int playerid);
extern bool onPlayerLeaveCheckpoint(int playerid);

extern bool onPlayerEnterRaceCheckpoint(int playerid);
extern bool onPlayerLeaveRaceCheckpoint(int playerid);

extern bool onRconCommand(const char* cmd);
extern bool onPlayerRequestSpawn(int playerid);

extern bool onObjectMoved(int objectid);
extern bool onPlayerObjectMoved(int playerid, int objectid);

extern bool onPlayerPickUpPickup(int playerid, int pickupid);
extern bool onVehicleMod(int playerid, int vehicleid, int componentid);

extern bool onEnterExitModShop(int playerid, bool enterexit, int interiorid);
extern bool onVehiclePaintjob(int playerid, int vehicleid, int paintjobid);

extern bool onVehicleRespray(int playerid, int vehicleid, int color1, int color2);
extern bool onVehicleDamageStatusUpdate(int vehicleid, int playerid);

extern bool onUnoccupiedVehicleUpdate(int vehicleid, int playerid, int passenger_seat, float new_x, float new_y, float new_z, float vel_x, float vel_y, float vel_z);
extern bool onPlayerSelectedMenuRow(int playerid, int row);

extern bool onPlayerExitedMenu(int playerid);
extern bool onPlayerInteriorChange(int playerid, int newinteriorid, int oldinteriorid);

extern bool onPlayerKeyStateChange(int playerid, int newkeys, int oldkeys);
extern bool onRconLoginAttempt(const char* ip, const char* password, bool success);

extern bool onPlayerUpdate(int playerid);

extern bool onPlayerStreamIn(int playerid, int forplayerid);
extern bool onPlayerStreamOut(int playerid, int forplayerid);

extern bool onVehicleStreamIn(int vehicleid, int forplayerid);
extern bool onVehicleStreamOut(int vehicleid, int forplayerid);

extern bool onActorStreamIn(int actorid, int forplayerid);
extern bool onActorStreamOut(int actorid, int forplayerid);

extern bool onDialogResponse(int playerid, int dialogid, int response, int listitem, const char* inputtext);

extern bool onPlayerTakeDamage(int playerid, int issuerid, float amount, int weaponid, int bodypart);
extern bool onPlayerGiveDamage(int playerid, int damagedid, float amount, int weaponid, int bodypart);

extern bool onPlayerGiveDamageActor(int playerid, int damaged_actorid, float amount, int weaponid, int bodypart);
extern bool onPlayerClickMap(int playerid, float fX, float fY, float fZ);

extern bool onPlayerClickTextDraw(int playerid, int clickedid);
extern bool onPlayerClickPlayerTextDraw(int playerid, int playertextid);

extern bool onIncomingConnection(int playerid, const char* ip_address, int port);
extern bool onTrailerUpdate(int playerid, int vehicleid);

extern bool onVehicleSirenStateChange(int playerid, int vehicleid, int newstate);
extern bool onPlayerClickPlayer(int playerid, int clickedplayerid, int source);

extern bool onPlayerEditObject(int playerid, bool playerobject, int objectid, int response, float fX, float fY, float fZ, float fRotX, float fRotY, float fRotZ);
extern bool onPlayerEditAttachedObject(int playerid, int response, int index, int modelid, int boneid, float fOffsetX, float fOffsetY, float fOffsetZ, float fRotX, float fRotY, float fRotZ, float fScaleX, float fScaleY, float fScaleZ);

extern bool onPlayerSelectObject(int playerid, int type, int objectid, int modelid, float fX, float fY, float fZ);
extern bool onPlayerWeaponShot(int playerid, int weaponid, int hittype, int hitid, float fX, float fY, float fZ);

extern bool onPlayerRequestDownload(int playerid, int type, int crc);

extern void onTick();
#endif

cell n_CallEvent(AMX* amx, cell* params);

// All of the natives we want to export.
extern void goLogprintf(char* text);
extern char* constToNonConst(const char* text);
extern const char* nonConstToConst(char* text);

#endif
