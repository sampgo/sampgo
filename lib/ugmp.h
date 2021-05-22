#ifndef UGMP_INCLUDE_H
#define UGMP_INCLUDE_H

// -
// Includes
// -

#include <string.h>
#include <sampgdk.h>

// -
// Macro Definitions
// -


#ifndef UGMP_FUNCTION
/*
	Creates an alias to a function.

	@param X The main function.
	@param Y The function that should point to it.
*/
#define UGMP_FUNCTION(X,Y)
	#ifndef X
	#define X Y
	#endif
#endif

#ifndef isequal
/*
	Checks if two strings are equal, and returns a boolean.

	@param S The source string.
	@param C The string to check against.

	@return The result of the check.
*/
#define isequal(S,C) (strcmp(S,C) == 0)
#endif

// -
// Macros
// -

UGMP_FUNCTION(ugmp_ChangePlayerPedColor, ugmp_ChangePlayerPedColour)
UGMP_FUNCTION(ugmp_ChangeVehiclePearlColor, ugmp_ChangeVehiclePearlColour)
UGMP_FUNCTION(ugmp_ChangeVehiclePearlColorRGB, ugmp_ChangeVehiclePearlColourRGB)

// -
// Types
// -

#ifndef char_t
typedef const char char_t;
#endif

// -
// Enumerators
// -

enum UGMP_MAP_OFFSETS {
	MAPOFFSET_MIAMI,
	MAPOFFSET_NEWYORK,
	MAPOFFSET_NEWENGLAND,
};

enum UGMP_VEHICLE_COLOUR {
	VEHICLE_COLOUR_PRIMARY,
	VEHICLE_COLOUR_SECONDARY,
	VEHICLE_COLOUR_TERTIARY,
	VEHICLE_COLOUR_QUATERNARY
};

// -
// Functions
// -

int ugmp_MT19937_Random(int a, int b);

void ugmp_ApplyMapOffsetToCoords(int type, float x, float y, float z);
void ugmp_ChangePlayerPedColour(int playerid, int colour1, int colour2, int colour3, int colour4);

void ugmp_EnableRadioAutoTune(bool toggle);
void ugmp_SetVehicleRadioStation(int vehicleid, int station);
int ugmp_GetVehicleRadioStation(int vehicleid);
void ugmp_GetRadioStationName(int station, char* str, int len);
bool ugmp_IsRadioAutoTuneEnabled();

void ugmp_ChangeVehicleColourRGB(int vehicleid, int colour, int red, int green, int blue);
void ugmp_ChangeVehiclePearlColour(int vehicleid, int colour);
void ugmp_ChangeVehiclePearlColourRGB(int vehicleid, int red, int green, int blue);

int  ugmp_CreateExtraWeather(int weatherHandlingType, int weatherType, int colourFilter);
int  ugmp_DestroyExtraWeather(int extraWeatherID);

int  ugmp_SetExtraWeatherParam_RGB(int extraweatherid, int hour, int param, int red, int green, int blue);
int  ugmp_SetExtraWeatherParam_RGBA(int extraweatherid, int hour, int param, float red, float green, float blue, float alpha);
int  ugmp_SetExtraWeatherParam_Float(int extraweatherid, int hour, int param, float fParam);
int  ugmp_SetExtraWeatherParam_Int(int extraweatherid, int hour, int param, int nParam);
int  ugmp_SetExtraWeatherParam_Flags(int extraweatherid, int flags);
int  ugmp_SetExtraWeatherParam_Windyness(int extraweatherid, float windyness);

void ugmp_ToggleMoon(bool enable);
void ugmp_ToggleStars(bool enable);
void ugmp_ToggleLowClouds(bool enable);
void ugmp_ToggleFluffyClouds(bool enable);
void ugmp_ToggleRainbow(bool enable);
void ugmp_SetCurrentSeason(int season);
void ugmp_TogglePlayerMoon(int playerid, bool enable);
void ugmp_TogglePlayerStars(int playerid, bool enable);
void ugmp_TogglePlayerLowClouds(int playerid, bool enable);
void ugmp_TogglePlayerFluffyClouds(int playerid, bool enable);
void ugmp_TogglePlayerRainbow(int playerid, bool enable);
void ugmp_SetPlayerSeason(int playerid, int season);
void ugmp_ToggleSASunPositionFormula(bool enable);
void ugmp_ToggleSASunPositionFormula_Pl(int playerid, bool enable);
int  ugmp_IsSASunPositionFormulaEnabled(int playerid);

int  ugmp_GetNumPedModels();

int  ugmp_GetNumVehicleModels();
int  ugmp_GetNumWeaponModels();

void ugmp_ToggleVehicleColorRGB(int vehicleid, int color, bool enable);
int  ugmp_GetVehicleComponentTypeEx(int component);

void ugmp_TogglePlayerNightVision(int playerid, bool enable);
void ugmp_TogglePlayerInfraRed(int playerid, bool enable);
void ugmp_TogglePlayerCCTV(int playerid, bool enable);
void ugmp_TogglePlayerFogOverlay(int playerid, bool enable);
void ugmp_TogglePlayerDarknessFilter(int playerid, bool enable, int darknessAlpha);
void ugmp_TogglePlayerVideoCameraOverlay(int playerid, bool enable);

bool ugmp_IsValidVehicleModel(int modelid);
int  ugmp_GetValidVehicleModelAt(int index);
//int  ugmp_GetValidVehicleModelAtEx(int index, int* modelid, int* vehtype, int* modifyas, char** name, size_t namelen = sizeof (name));
int  ugmp_GetRandomVehicleModel();

bool ugmp_IsValidPedModel(int modelid);
int  ugmp_GetValidPedModelAt(int index);
int  ugmp_GetRandomPedModel();

int  ugmp_GetWeaponHighestParentType(int weapontype);
int  ugmp_GetValidWeaponModelAt(int index);
int  ugmp_GetValidWeaponTypeAt(int index);
bool ugmp_IsValidWeaponType(int weapontype);
bool ugmp_IsValidWeaponModel(int modelid);
int  ugmp_GetRandomWeaponModel();
int  ugmp_GetRandomWeaponType();

void ugmp_InitialiseDiscordRichPresence(char_t* applicationID);
void ugmp_UpdateDiscordRichPresence(char_t* smallImageKey, char_t* smallImageText, char_t* largeImageKey, char_t* largeImageText, char_t* details);
void ugmp_ShutdownDiscordRichPresence();

void ugmp_SetPlayerKnockedOffBikeState(int playerid, int knockState);
void ugmp_SetPlayerFireProof(int playerid, bool enable);
void ugmp_TogglePlayerInfiniteSprint(int playerid, bool enable);
void ugmp_TogglePlayerSun(int playerid, bool enable);
void ugmp_TogglePlayerRubbish(int playerid, bool enable);
bool ugmp_IsRubbishVisibleForPlayer(int playerid);
void ugmp_TogglePlayerGrass(int playerid, bool enable);


bool ugmp_IsValidAnimationAndLibrary(char_t* animlib, char_t* animname);
bool ugmp_SetAircraftHeightLimitForPlayer(int playerid, float limit);

#endif
