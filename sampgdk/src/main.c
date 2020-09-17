#include "main.h"

void SAMPGDK_CALL PrintTickCountTimer(int timerid, void *params) {
    sampgdk_logprintf("Tick count: %d", GetTickCount());
}

PLUGIN_EXPORT bool PLUGIN_CALL OnGameModeInit() {
    SetGameModeText("GoMode");
    SetTimer(1000, true, PrintTickCountTimer, 0);
    goModeInit();
    return true;
}

PLUGIN_EXPORT bool PLUGIN_CALL OnPlayerConnect(int playerid) {
    sampgdk_logprintf("OnPlayerConnect called");
    return true;
}

PLUGIN_EXPORT unsigned int PLUGIN_CALL Supports()
{
    sampgdk_logprintf("Supports Called from C");
    return SUPPORTS_VERSION | SUPPORTS_AMX_NATIVES;
}

PLUGIN_EXPORT bool PLUGIN_CALL Load(void** ppData)
{
    sampgdk_logprintf("Load called from C");
    return true;
}