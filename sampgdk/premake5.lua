workspace "sampgdk"
    configurations { "Release" }
    location "build"

project "sampgdk"
    kind "StaticLib"
    language "C"
    architecture "x32"

    files { "sampgdk.h", "sampgdk.h" }
    includedirs { "samp-plugin-sdk", "samp-plugin-sdk/amx" }

    buildoptions { "-Wno-attributes" }

    filter { "configurations:Release" }
        defines { }
        optimize "On"