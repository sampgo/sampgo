#if defined _sampgo_included
	#endinput
#endif

#define _sampgo_included

// Should return 1 if this works out all fine.
native GoInt32: sampgo_CallEvent(const function[32], const format[] = "", {Float,_}:...);