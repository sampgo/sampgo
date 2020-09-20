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
        files { "main.def" }

    filter { "system:linux" }
        defines { "LINUX" }
        
    filter { "toolset:gcc" }
        buildoptions { "-Wno-attributes", "-include stddef.h" }

    filter { "toolset:gcc", "system:windows" }
        buildoptions { "-include stdint.h" }

    filter { "toolset:clang" }
        buildoptions { "-Wno-ignored-attributes", "-include stddef.h" }

    filter { "configurations:Debug" }
        defines { "DEBUG" }
        symbols "On"

    filter { "configurations:Release" }
        defines { "NDEBUG" }
        optimize "On"