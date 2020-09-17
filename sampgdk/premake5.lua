workspace "sampgdk"
    configurations { "Release" }
    location "build"

project "sampgdk"
    kind "StaticLib"
    language "C"
    architecture "x32"

    files { "src/*.c", "src/*.h", "lib/sampgdk/*.c", "lib/sampgdk/*.h", "lib/samp-plugin-sdk/**.c", "lib/samp-plugin-sdk/**.h" }
    includedirs { "lib/sampgdk", "lib/samp-plugin-sdk", "lib/samp-plugin-sdk/amx" }

    buildoptions { "-Wno-attributes" }

    filter { "configurations:Release" }
        optimize "On"