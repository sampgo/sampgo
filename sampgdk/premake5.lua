workspace "sampgdk"
    configurations { "Debug", "Release" }
    location "build"

project "sampgdk"
    kind "StaticLib"
    language "C"
    files { "**.h", "**.c" }
    architecture "x32"

    
    includedirs { "amx" }

    filter { "toolset:gcc or toolset:clang" }
        buildoptions { "-Wno-attributes", "-include stddef.h" }

    filter { "configurations:Debug" }
        defines { "DEBUG", "LINUX" }
        symbols "On"

    filter { "configurations:Release" }
        defines { "NDEBUG", "LINUX" }
        optimize "On"