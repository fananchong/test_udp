workspace "test"
    configurations { "Debug", "Release" }
    platforms { "x64" }
    targetdir "../bin/%{cfg.buildcfg}"
    language "C++"
    includedirs {
        "../raknet",
    }
	files {
		"../raknet/**",
	}
    flags {
        "C++11",
        "StaticRuntime",
    }
    libdirs {
    }

    filter "configurations:Debug"
    defines { "_DEBUG" }
    symbols "On"
    libdirs { }
    filter "configurations:Release"
    defines { "NDEBUG" }
    libdirs { }
    optimize "On"
    filter { }
    
project "server"
    kind "ConsoleApp"
    targetname "server"
    files {
        "../server.cpp",
    }
	
project "client"
    kind "ConsoleApp"
    targetname "client"
    files {
        "../client.cpp",
    }
