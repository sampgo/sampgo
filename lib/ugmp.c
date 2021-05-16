#include "ugmp.h"

bool ugmp_OnVehicleResprayAtGarage(int playerid, int vehicleid, int colour1, int colour2, int colour3, int colour4) {
	// give cgo a hard time.
	return true;
}

bool ugmp_OnLaserDotUpdate(int playerid, int weaponid, float x, float y, float z) {
	// give cgo a hard time.
	return true;
}

bool ugmp_OnPlayerChangeRadio(int playerid, int vehicleid, int oldstation, int newstation) {
	// give cgo a hard time.
	return true;
}

PLUGIN_EXPORT bool PLUGIN_CALL OnPublicCall(AMX* amx, const char* name, cell* params, cell* retval) {
	
	if (isequal(name, "OnVehicleResprayAtGarage")) {
		int 
			playerid = (int)params[1],
			vehicleid = (int)params[2],
			colour1 = (int)params[3],
			colour2 = (int)params[4],
			colour3 = (int)params[5],
			colour4 = (int)params[6];
	
		return ugmp_OnVehicleResprayAtGarage(playerid, vehicleid, colour1, colour2, colour3, colour4);
	} else if (isequal(name, "OnLaserDotUpdate")) {
		int
			playerid = (int)params[1],
			weaponid = (int)params[2];

		float
			x = amx_ctof(params[3]),
			y = amx_ctof(params[4]),
			z = amx_ctof(params[5]);

		return ugmp_OnLaserDotUpdate(playerid, weaponid, x, y, z);
	} else if (isequal(name, "OnPlayerChangeRadio")) {
		int
			playerid = (int)params[1],
			vehicleid = (int)params[2],
			oldstation = (int)params[3],
			newstation = (int)params[4];
		
		return ugmp_OnPlayerChangeRadio(playerid, vehicleid, oldstation, newstation);
	}

	return true;
}

int ugmp_MT19937_Random(int a, int b) {
	AMX_NATIVE Native = sampgdk_FindNative("MT19937_Random");
	if (Native != NULL) return sampgdk_InvokeNative(Native, "ii", a, b);
	else return sampgdk_logprintf( "'" __FILE__ "' / '%s' - Function not discovered !", __func__ ), 0;
}

void ugmp_ApplyMapOffsetToCoords(int type, float x, float y, float z) {
	AMX_NATIVE Native = sampgdk_FindNative("ApplyMapOffsetToCoords");
	if (Native != NULL) sampgdk_InvokeNative(Native, "fff", x, y, z);
	else sampgdk_logprintf( "'" __FILE__ "' / '%s' - Function not discovered !", __func__ );
}

void ugmp_ChangePlayerPedColour(int playerid, int colour1, int colour2, int colour3, int colour4) {
	AMX_NATIVE Native = sampgdk_FindNative("ChangePlayerPedColour");
	if (Native != NULL) sampgdk_InvokeNative(Native, "iiii", playerid, colour1, colour2, colour3, colour4);
	else sampgdk_logprintf( "'" __FILE__ "' / '%s' - Function not discovered !", __func__ );
}

void ugmp_EnableRadioAutoTune(bool toggle) {
	AMX_NATIVE Native = sampgdk_FindNative("EnableRadioAutoTune");
	if (Native != NULL) sampgdk_InvokeNative(Native, "b", toggle);
	else sampgdk_logprintf( "'" __FILE__ "' / '%s' - Function not discovered !", __func__ );
}

void ugmp_SetVehicleRadioStation(int vehicleid, int station) {
	AMX_NATIVE Native = sampgdk_FindNative("SetVehicleRadioStation");
	if (Native != NULL) sampgdk_InvokeNative(Native, "ii", vehicleid, station);
	else sampgdk_logprintf( "'" __FILE__ "' / '%s' - Function not discovered !", __func__ );
}

int ugmp_GetVehicleRadioStation(int vehicleid) {
	AMX_NATIVE Native = sampgdk_FindNative("GetVehicleRadioStation");
	if (Native != NULL) return sampgdk_InvokeNative(Native, "i", vehicleid);
	else return sampgdk_logprintf( "'" __FILE__ "' / '%s' - Function not discovered !", __func__ ), 0;
}

void ugmp_GetRadioStationName(int station, char* str, int len) {
	AMX_NATIVE Native = sampgdk_FindNative("GetRadioStationName");
	if (Native != NULL) sampgdk_InvokeNative(Native, "iRi", station, str, len);
	else sampgdk_logprintf( "'" __FILE__ "' / '%s' - Function not discovered !", __func__ );
}

bool ugmp_IsRadioAutoTuneEnabled() {
	AMX_NATIVE Native = sampgdk_FindNative("IsRadioAutoTuneEnabled");
	if (Native != NULL) return sampgdk_InvokeNative(Native, "");
	else return sampgdk_logprintf( "'" __FILE__ "' / '%s' - Function not discovered !", __func__ ), 0;
}

void ugmp_ChangeVehicleColourRGB(int vehicleid, int colour, int red, int green, int blue) {
	AMX_NATIVE Native = sampgdk_FindNative("ChangeVehicleColourRGB");
	if (Native != NULL) sampgdk_InvokeNative(Native, "iiiii", vehicleid, colour, red, green, blue);
	else sampgdk_logprintf( "'" __FILE__ "' / '%s' - Function not discovered !", __func__ );
}

void ugmp_ChangeVehiclePearlColour(int vehicleid, int colour) {
	AMX_NATIVE Native = sampgdk_FindNative("ChangeVehiclePearlColour");
	if (Native != NULL) sampgdk_InvokeNative(Native, "ii", vehicleid, colour);
	else sampgdk_logprintf( "'" __FILE__ "' / '%s' - Function not discovered !", __func__ );
}

void ugmp_ChangeVehiclePearlColourRGB(int vehicleid, int red, int green, int blue) {
	AMX_NATIVE Native = sampgdk_FindNative("ChangeVehiclePearlColourRGB");
	if (Native != NULL) sampgdk_InvokeNative(Native, "iiii", vehicleid, red, green, blue);
	else sampgdk_logprintf( "'" __FILE__ "' / '%s' - Function not discovered !", __func__ );
}

int  ugmp_SetExtraWeatherParam_RGB(int extraweatherid, int hour, int param, int red, int green, int blue) {
	AMX_NATIVE Native = sampgdk_FindNative("SetExtraWeatherParam_RGB");
	if (Native != NULL) return sampgdk_InvokeNative(Native, "iiiii", extraweatherid, hour, param, red, green, blue);
	else return sampgdk_logprintf( "'" __FILE__ "' / '%s' - Function not discovered !", __func__ ), 0;
}
int  ugmp_SetExtraWeatherParam_RGBA(int extraweatherid, int hour, int param, float red, float green, float blue, float alpha) {
	AMX_NATIVE Native = sampgdk_FindNative("SetExtraWeatherParam_RGBA");
	if (Native != NULL) return sampgdk_InvokeNative(Native, "iiiffff", extraweatherid, hour, param, red, green, blue, alpha);
	else return sampgdk_logprintf( "'" __FILE__ "' / '%s' - Function not discovered !", __func__ ), 0;
}
int  ugmp_SetExtraWeatherParam_Float(int extraweatherid, int hour, int param, float fParam) {
	AMX_NATIVE Native = sampgdk_FindNative("SetExtraWeatherParam_Float");
	if (Native != NULL) return sampgdk_InvokeNative(Native, "iiii", extraweatherid, hour, param, fParam);
	else return sampgdk_logprintf( "'" __FILE__ "' / '%s' - Function not discovered !", __func__ ), 0;
}

int  ugmp_SetExtraWeatherParam_Int(int extraweatherid, int hour, int param, int nParam) {
	AMX_NATIVE Native = sampgdk_FindNative("SetExtraWeatherParam_Int");
	if (Native != NULL) return sampgdk_InvokeNative(Native, "iiii", extraweatherid, hour, param, nParam);
	else return sampgdk_logprintf( "'" __FILE__ "' / '%s' - Function not discovered !", __func__ ), 0;
}

int  ugmp_SetExtraWeatherParam_Flags(int extraweatherid, int flags) {
	AMX_NATIVE Native = sampgdk_FindNative("SetExtraWeatherParam_Flags");
	if (Native != NULL) return sampgdk_InvokeNative(Native, "ii", extraweatherid, flags);
	else return sampgdk_logprintf( "'" __FILE__ "' / '%s' - Function not discovered !", __func__ ), 0;
}

int  ugmp_SetExtraWeatherParam_Windyness(int extraweatherid, float windyness) {
	AMX_NATIVE Native = sampgdk_FindNative("SetExtraWeatherParam_Windyness");
	if (Native != NULL) return sampgdk_InvokeNative(Native, "ii", extraweatherid, windyness);
	else return sampgdk_logprintf( "'" __FILE__ "' / '%s' - Function not discovered !", __func__ ), 0;
}

int  ugmp_CreateExtraWeather(int weatherHandlingType, int weatherType, int colourFilter) {
	AMX_NATIVE Native = sampgdk_FindNative("CreateExtraWeather");
	if (Native != NULL) return sampgdk_InvokeNative(Native, "iii", weatherHandlingType, weatherType, colourFilter);
	else return sampgdk_logprintf( "'" __FILE__ "' / '%s' - Function not discovered !", __func__ ), 0;
}

int  ugmp_DestroyExtraWeather(int extraWeatherID) {
	AMX_NATIVE Native = sampgdk_FindNative("DestroyExtraWeather");
	if (Native != NULL) return sampgdk_InvokeNative(Native, "i", extraWeatherID);
	else return sampgdk_logprintf( "'" __FILE__ "' / '%s' - Function not discovered !", __func__ ), 0;
}

void ugmp_ToggleMoon(bool enable) {
	AMX_NATIVE Native = sampgdk_FindNative("ToggleMoon");
	if (Native != NULL) sampgdk_InvokeNative(Native, "b", enable);
	else sampgdk_logprintf( "'" __FILE__ "' / '%s' - Function not discovered !", __func__ );
}

void ugmp_ToggleStars(bool enable) {
	AMX_NATIVE Native = sampgdk_FindNative("ToggleStars");
	if (Native != NULL) sampgdk_InvokeNative(Native, "b", enable);
	else sampgdk_logprintf( "'" __FILE__ "' / '%s' - Function not discovered !", __func__ );
}

void ugmp_ToggleLowClouds(bool enable) {
	AMX_NATIVE Native = sampgdk_FindNative("ToggleLowClouds");
	if (Native != NULL) sampgdk_InvokeNative(Native, "b", enable);
	else sampgdk_logprintf( "'" __FILE__ "' / '%s' - Function not discovered !", __func__ );
}

void ugmp_ToggleFluffyClouds(bool enable) {
	AMX_NATIVE Native = sampgdk_FindNative("ToggleFluffyClouds");
	if (Native != NULL) sampgdk_InvokeNative(Native, "b", enable);
	else sampgdk_logprintf( "'" __FILE__ "' / '%s' - Function not discovered !", __func__ );
}

void ugmp_ToggleRainbow(bool enable) {
	AMX_NATIVE Native = sampgdk_FindNative("ToggleRainbow");
	if (Native != NULL) sampgdk_InvokeNative(Native, "b", enable);
	else sampgdk_logprintf( "'" __FILE__ "' / '%s' - Function not discovered !", __func__ );
}

void ugmp_SetCurrentSeason(int season) {
	AMX_NATIVE Native = sampgdk_FindNative("SetCurrentSeason");
	if (Native != NULL) sampgdk_InvokeNative(Native, "i", season);
	else sampgdk_logprintf( "'" __FILE__ "' / '%s' - Function not discovered !", __func__ );
}

void ugmp_TogglePlayerMoon(int playerid, bool enable) {
	AMX_NATIVE Native = sampgdk_FindNative("TogglePlayerMoon");
	if (Native != NULL) sampgdk_InvokeNative(Native, "ib", playerid, enable);
	else sampgdk_logprintf( "'" __FILE__ "' / '%s' - Function not discovered !", __func__ );
}

void ugmp_TogglePlayerStars(int playerid, bool enable) {
	AMX_NATIVE Native = sampgdk_FindNative("TogglePlayerStars");
	if (Native != NULL) sampgdk_InvokeNative(Native, "ib", playerid, enable);
	else sampgdk_logprintf( "'" __FILE__ "' / '%s' - Function not discovered !", __func__ );
}

void ugmp_TogglePlayerLowClouds(int playerid, bool enable) {
	AMX_NATIVE Native = sampgdk_FindNative("TogglePlayerLowClouds");
	if (Native != NULL) sampgdk_InvokeNative(Native, "ib", playerid, enable);
	else sampgdk_logprintf( "'" __FILE__ "' / '%s' - Function not discovered !", __func__ );
}

void ugmp_TogglePlayerFluffyClouds(int playerid, bool enable) {
	AMX_NATIVE Native = sampgdk_FindNative("TogglePlayerFluffyClouds");
	if (Native != NULL) sampgdk_InvokeNative(Native, "ib", playerid, enable);
	else sampgdk_logprintf( "'" __FILE__ "' / '%s' - Function not discovered !", __func__ );
}

void ugmp_TogglePlayerRainbow(int playerid, bool enable) {
	AMX_NATIVE Native = sampgdk_FindNative("TogglePlayerRainbow");
	if (Native != NULL) sampgdk_InvokeNative(Native, "ib", playerid, enable);
	else sampgdk_logprintf( "'" __FILE__ "' / '%s' - Function not discovered !", __func__ );
}

void ugmp_SetPlayerSeason(int playerid, int season) {
	AMX_NATIVE Native = sampgdk_FindNative("SetPlayerSeason");
	if (Native != NULL) sampgdk_InvokeNative(Native, "ii", playerid, season);
	else sampgdk_logprintf( "'" __FILE__ "' / '%s' - Function not discovered !", __func__ );
}

void ugmp_ToggleSASunPositionFormula(bool enable) {
	AMX_NATIVE Native = sampgdk_FindNative("ToggleSASunPositionFormula");
	if (Native != NULL) sampgdk_InvokeNative(Native, "b", enable);
	else sampgdk_logprintf( "'" __FILE__ "' / '%s' - Function not discovered !", __func__ );
}

void ugmp_ToggleSASunPositionFormula_Pl(int playerid, bool enable) {
	AMX_NATIVE Native = sampgdk_FindNative("ToggleSASunPositionFormula_Pl");
	if (Native != NULL) sampgdk_InvokeNative(Native, "ib", playerid, enable);
	else sampgdk_logprintf( "'" __FILE__ "' / '%s' - Function not discovered !", __func__ );
}

int  ugmp_IsSASunPositionFormulaEnabled(int playerid) {
	AMX_NATIVE Native = sampgdk_FindNative("IsSASunPositionFormulaEnabled");
	if (Native != NULL) return sampgdk_InvokeNative(Native, "i", playerid);
	else return sampgdk_logprintf( "'" __FILE__ "' / '%s' - Function not discovered !", __func__ ), 0;
}

int  ugmp_GetNumPedModels() {
	AMX_NATIVE Native = sampgdk_FindNative("GetNumPedModels");
	if (Native != NULL) return sampgdk_InvokeNative(Native, "");
	else return sampgdk_logprintf( "'" __FILE__ "' / '%s' - Function not discovered !", __func__ ), 0;
}

int  ugmp_GetNumVehicleModels() {
	AMX_NATIVE Native = sampgdk_FindNative("GetNumVehicleModels");
	if (Native != NULL) return sampgdk_InvokeNative(Native, "");
	else return sampgdk_logprintf( "'" __FILE__ "' / '%s' - Function not discovered !", __func__ ), 0;
}

int  ugmp_GetNumWeaponModels() {
	AMX_NATIVE Native = sampgdk_FindNative("GetNumWeaponModels");
	if (Native != NULL) return sampgdk_InvokeNative(Native, "");
	else return sampgdk_logprintf( "'" __FILE__ "' / '%s' - Function not discovered !", __func__ ), 0;
}

void ugmp_ToggleVehicleColorRGB(int vehicleid, int color, bool enable) {
	AMX_NATIVE Native = sampgdk_FindNative("ToggleVehicleColorRGB");
	if (Native != NULL) sampgdk_InvokeNative(Native, "iib", vehicleid, color, enable);
	else sampgdk_logprintf( "'" __FILE__ "' / '%s' - Function not discovered !", __func__ );
}

int  ugmp_GetVehicleComponentTypeEx(int component) {
	AMX_NATIVE Native = sampgdk_FindNative("GetVehicleComponentTypeEx");
	if (Native != NULL) return sampgdk_InvokeNative(Native, "i", component);
	else return sampgdk_logprintf( "'" __FILE__ "' / '%s' - Function not discovered !", __func__ ), 0;
}

void ugmp_TogglePlayerNightVision(int playerid, bool enable) {
	AMX_NATIVE Native = sampgdk_FindNative("TogglePlayerNightVision");
	if (Native != NULL) sampgdk_InvokeNative(Native, "ib", playerid, enable);
	else sampgdk_logprintf( "'" __FILE__ "' / '%s' - Function not discovered !", __func__ );
}

void ugmp_TogglePlayerInfraRed(int playerid, bool enable) {
	AMX_NATIVE Native = sampgdk_FindNative("TogglePlayerInfraRed");
	if (Native != NULL) sampgdk_InvokeNative(Native, "ib", playerid, enable);
	else sampgdk_logprintf( "'" __FILE__ "' / '%s' - Function not discovered !", __func__ );
}

void ugmp_TogglePlayerCCTV(int playerid, bool enable) {
	AMX_NATIVE Native = sampgdk_FindNative("TogglePlayerCCTV");
	if (Native != NULL) sampgdk_InvokeNative(Native, "ib", playerid, enable);
	else sampgdk_logprintf( "'" __FILE__ "' / '%s' - Function not discovered !", __func__ );
}

void ugmp_TogglePlayerFogOverlay(int playerid, bool enable) {
	AMX_NATIVE Native = sampgdk_FindNative("TogglePlayerFogOverlay");
	if (Native != NULL) sampgdk_InvokeNative(Native, "ib", playerid, enable);
	else sampgdk_logprintf( "'" __FILE__ "' / '%s' - Function not discovered !", __func__ );
}

void ugmp_TogglePlayerDarknessFilter(int playerid, bool enable, int darknessAlpha) {
	AMX_NATIVE Native = sampgdk_FindNative("TogglePlayerDarknessFilter");
	if (Native != NULL) sampgdk_InvokeNative(Native, "ibi", playerid, enable, darknessAlpha);
	else sampgdk_logprintf( "'" __FILE__ "' / '%s' - Function not discovered !", __func__ );
}

void ugmp_TogglePlayerVideoCameraOverlay(int playerid, bool enable) {
	AMX_NATIVE Native = sampgdk_FindNative("TogglePlayerVideoCameraOverlay");
	if (Native != NULL) sampgdk_InvokeNative(Native, "ib", playerid, enable);
	else sampgdk_logprintf( "'" __FILE__ "' / '%s' - Function not discovered !", __func__ );
}

bool ugmp_IsValidVehicleModel(int modelid) {
	AMX_NATIVE Native = sampgdk_FindNative("IsValidVehicleModel");
	if (Native != NULL) return sampgdk_InvokeNative(Native, "i", modelid);
	else return sampgdk_logprintf( "'" __FILE__ "' / '%s' - Function not discovered !", __func__ ), 0;
}

int ugmp_GetValidVehicleModelAt(int index) {
	AMX_NATIVE Native = sampgdk_FindNative("GetValidVehicleModelAt");
	if (Native != NULL) return sampgdk_InvokeNative(Native, "i", index);
	else return sampgdk_logprintf( "'" __FILE__ "' / '%s' - Function not discovered !", __func__ ), 0;
}

/*int  ugmp_GetValidVehicleModelAtEx(int index, int* modelid, int* vehtype, int* modifyas, char** name, size_t namelen = sizeof (name)) {
	AMX_NATIVE Native = sampgdk_FindNative("GetValidVehicleModelAtEx");
	if (Native != NULL) return sampgdk_InvokeNative(Native, "iRRRRi", index, modelid, vehtype, modifyas, name, namelen);
	else return sampgdk_logprintf( "'" __FILE__ "' / '%s' - Function not discovered !", __func__ ), 0;
}*/

int  ugmp_GetRandomVehicleModel() {
	AMX_NATIVE Native = sampgdk_FindNative("GetRandomVehicleModel");
	if (Native != NULL) return sampgdk_InvokeNative(Native, "");
	else return sampgdk_logprintf( "'" __FILE__ "' / '%s' - Function not discovered !", __func__ ), 0;
}

bool ugmp_IsValidPedModel(int modelid) {
	AMX_NATIVE Native = sampgdk_FindNative("IsValidPedModel");
	if (Native != NULL) return sampgdk_InvokeNative(Native, "i", modelid);
	else return sampgdk_logprintf( "'" __FILE__ "' / '%s' - Function not discovered !", __func__ ), 0;
}

int  ugmp_GetValidPedModelAt(int index) {
	AMX_NATIVE Native = sampgdk_FindNative("GetValidPedModelAt");
	if (Native != NULL) return sampgdk_InvokeNative(Native, "i", index);
	else return sampgdk_logprintf( "'" __FILE__ "' / '%s' - Function not discovered !", __func__ ), 0;
}

int  ugmp_GetRandomPedModel() {
	AMX_NATIVE Native = sampgdk_FindNative("GetRandomPedModel");
	if (Native != NULL) return sampgdk_InvokeNative(Native, "");
	else return sampgdk_logprintf( "'" __FILE__ "' / '%s' - Function not discovered !", __func__ ), 0;
}

int  ugmp_GetWeaponHighestParentType(int weapontype) {
	AMX_NATIVE Native = sampgdk_FindNative("GetWeaponHighestParentType");
	if (Native != NULL) return sampgdk_InvokeNative(Native, "i", weapontype);
	else return sampgdk_logprintf( "'" __FILE__ "' / '%s' - Function not discovered !", __func__ ), 0;
}

int  ugmp_GetValidWeaponModelAt(int index) {
	AMX_NATIVE Native = sampgdk_FindNative("GetValidWeaponModelAt");
	if (Native != NULL) return sampgdk_InvokeNative(Native, "i", index);
	else return sampgdk_logprintf( "'" __FILE__ "' / '%s' - Function not discovered !", __func__ ), 0;
}

int  ugmp_GetValidWeaponTypeAt(int index) {
	AMX_NATIVE Native = sampgdk_FindNative("GetValidWeaponTypeAt");
	if (Native != NULL) return sampgdk_InvokeNative(Native, "i", index);
	else return sampgdk_logprintf( "'" __FILE__ "' / '%s' - Function not discovered !", __func__ ), 0;
}

bool ugmp_IsValidWeaponType(int weapontype) {
	AMX_NATIVE Native = sampgdk_FindNative("IsValidWeaponType");
	if (Native != NULL) return sampgdk_InvokeNative(Native, "i", weapontype);
	else return sampgdk_logprintf( "'" __FILE__ "' / '%s' - Function not discovered !", __func__ ), 0;
}

bool ugmp_IsValidWeaponModel(int modelid) {
	AMX_NATIVE Native = sampgdk_FindNative("IsValidWeaponModel");
	if (Native != NULL) return sampgdk_InvokeNative(Native, "i", modelid);
	else return sampgdk_logprintf( "'" __FILE__ "' / '%s' - Function not discovered !", __func__ ), 0;
}

int  ugmp_GetRandomWeaponModel() {
	AMX_NATIVE Native = sampgdk_FindNative("GetRandomWeaponModel");
	if (Native != NULL) return sampgdk_InvokeNative(Native, "");
	else return sampgdk_logprintf( "'" __FILE__ "' / '%s' - Function not discovered !", __func__ ), 0;
}

int  ugmp_GetRandomWeaponType() {
	AMX_NATIVE Native = sampgdk_FindNative("GetRandomWeaponType");
	if (Native != NULL) return sampgdk_InvokeNative(Native, "");
	else return sampgdk_logprintf( "'" __FILE__ "' / '%s' - Function not discovered !", __func__ ), 0;
}

void ugmp_InitialiseDiscordRichPresence(char_t* applicationID) {
	AMX_NATIVE Native = sampgdk_FindNative("InitialiseDiscordRichPresence");
	if (Native != NULL) sampgdk_InvokeNative(Native, "s", applicationID);
	else sampgdk_logprintf( "'" __FILE__ "' / '%s' - Function not discovered !", __func__ );
}

void ugmp_UpdateDiscordRichPresence(char_t* smallImageKey, char_t* smallImageText, char_t* largeImageKey, char_t* largeImageText, char_t* details) {
	AMX_NATIVE Native = sampgdk_FindNative("UpdateDiscordRichPresence");
	if (Native != NULL) sampgdk_InvokeNative(Native, "sssss", smallImageKey, smallImageText, largeImageKey, largeImageText, details);
	else sampgdk_logprintf( "'" __FILE__ "' / '%s' - Function not discovered !", __func__ );
}

void ugmp_ShutdownDiscordRichPresence() {
	AMX_NATIVE Native = sampgdk_FindNative("ShutdownDiscordRichPresence");
	if (Native != NULL) sampgdk_InvokeNative(Native, "");
	else sampgdk_logprintf( "'" __FILE__ "' / '%s' - Function not discovered !", __func__ );
}

void ugmp_SetPlayerKnockedOffBikeState(int playerid, int knockState) {
	AMX_NATIVE Native = sampgdk_FindNative("SetPlayerKnockedOffBikeState");
	if (Native != NULL) sampgdk_InvokeNative(Native, "ii", playerid, knockState);
	else sampgdk_logprintf( "'" __FILE__ "' / '%s' - Function not discovered !", __func__ );
}

void ugmp_SetPlayerFireProof(int playerid, bool enable) {
	AMX_NATIVE Native = sampgdk_FindNative("SetPlayerFireProof");
	if (Native != NULL) sampgdk_InvokeNative(Native, "ib", playerid, enable);
	else sampgdk_logprintf( "'" __FILE__ "' / '%s' - Function not discovered !", __func__ );
}

void ugmp_TogglePlayerInfiniteSprint(int playerid, bool enable) {
	AMX_NATIVE Native = sampgdk_FindNative("TogglePlayerInfiniteSprint");
	if (Native != NULL) sampgdk_InvokeNative(Native, "ib", playerid, enable);
	else sampgdk_logprintf( "'" __FILE__ "' / '%s' - Function not discovered !", __func__ );
}

void ugmp_TogglePlayerSun(int playerid, bool enable) {
	AMX_NATIVE Native = sampgdk_FindNative("TogglePlayerSun");
	if (Native != NULL) sampgdk_InvokeNative(Native, "ib", playerid, enable);
	else sampgdk_logprintf( "'" __FILE__ "' / '%s' - Function not discovered !", __func__ );
}

void ugmp_TogglePlayerRubbish(int playerid, bool enable) {
	AMX_NATIVE Native = sampgdk_FindNative("TogglePlayerRubbish");
	if (Native != NULL) sampgdk_InvokeNative(Native, "ib", playerid, enable);
	else sampgdk_logprintf( "'" __FILE__ "' / '%s' - Function not discovered !", __func__ );
}

bool ugmp_IsRubbishVisibleForPlayer(int playerid) {
	AMX_NATIVE Native = sampgdk_FindNative("IsRubbishVisibleForPlayer");
	if (Native != NULL) return sampgdk_InvokeNative(Native, "i", playerid);
	else return sampgdk_logprintf( "'" __FILE__ "' / '%s' - Function not discovered !", __func__ ), false;
}

void ugmp_TogglePlayerGrass(int playerid, bool enable) {
	AMX_NATIVE Native = sampgdk_FindNative("TogglePlayerGrass");
	if (Native != NULL) sampgdk_InvokeNative(Native, "ib", playerid, enable);
	else sampgdk_logprintf( "'" __FILE__ "' / '%s' - Function not discovered !", __func__ );
}

bool ugmp_IsValidAnimationAndLibrary(char_t* animlib, char_t* animname) {
	AMX_NATIVE Native = sampgdk_FindNative("IsValidAnimationAndLibrary");
	if (Native != NULL) return sampgdk_InvokeNative(Native, "RR", animlib, animname);
	else return sampgdk_logprintf( "'" __FILE__ "' / '%s' - Function not discovered !", __func__ ), false;
}

bool ugmp_SetAircraftHeightLimitForPlayer(int playerid, float limit) {
	AMX_NATIVE Native = sampgdk_FindNative("SetAircraftHeightLimitForPlayer");
	if (Native != NULL) return sampgdk_InvokeNative(Native, "if", playerid, limit);
	else sampgdk_logprintf( "'" __FILE__ "' / '%s' - Function not discovered !", __func__ ), 0;
}