workspace "sampgdk"
    configurations { "Debug", "Release" }
    location "build"

project "sampgdk"
    kind "StaticLib"
    language "C"
    files { "**.h", "**.c" }
    architecture "x32"

    includedirs { "amx" }

    filter { "system:windows" }
        defines { "WIN32" }

    filter { "system:linux" }
        defines { "LINUX" }
        
    filter { "toolset:gcc or toolset:clang" }
        buildoptions { "-Wno-attributes" }

    filter { "toolset:gcc" }
        buildoptions{ "-include stddef.h" }

    filter { "configurations:Debug" }
        defines { "DEBUG" }
        symbols "On"

    filter { "configurations:Release" }
        defines { "NDEBUG" }
        optimize "On"