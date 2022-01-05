load("@io_bazel_rules_go//go:def.bzl", "go_library")

WIN32_DEFINES = [
    "-D_GLFW_WIN32",
]

WIN32_HDRS = [
    "v3.3/glfw/glfw/src/win32_joystick.h",
    "v3.3/glfw/glfw/src/win32_platform.h",
    "v3.3/glfw/glfw/src/wgl_context.h", 
]

WIN32_SRCS = [
    "v3.3/glfw/glfw/src/win32_init.c",
    "v3.3/glfw/glfw/src/win32_joystick.c",
    "v3.3/glfw/glfw/src/win32_monitor.c",
    "v3.3/glfw/glfw/src/win32_thread.c",
    "v3.3/glfw/glfw/src/win32_time.c",
    "v3.3/glfw/glfw/src/win32_window.c",
    "v3.3/glfw/glfw/src/wgl_context.c",
]

WIN32_LINKOPTS = [
    "-DEFAULTLIB:user32.lib",
    "-DEFAULTLIB:gdi32.lib",
    "-DEFAULTLIB:shell32.lib",
]


DARWIN_DEFINES = [
    "-D_GLFW_COCOA",
    "-D_GLFW_NSGL",
    "-D_GLFW_NO_DLOAD_WINMM",
    # "-D_GLFW_USE_OPENGL",
]

DARWIN_HDRS = [
    "v3.3/glfw/glfw/src/cocoa_joystick.h",
    "v3.3/glfw/glfw/src/cocoa_platform.h",
    "v3.3/glfw/glfw/src/glx_context.h",
    "v3.3/glfw/glfw/src/nsgl_context.h",
    "v3.3/glfw/glfw/src/null_joystick.h",
    "v3.3/glfw/glfw/src/null_platform.h",
    "v3.3/glfw/glfw/src/posix_thread.h",
    "v3.3/glfw/glfw/src/wl_platform.h",
    "v3.3/glfw/glfw/src/cocoa_window.m",
    "v3.3/glfw/glfw/src/cocoa_joystick.m",
    "v3.3/glfw/glfw/src/nsgl_context.m",
    "v3.3/glfw/glfw/src/cocoa_monitor.m",
    "v3.3/glfw/glfw/src/cocoa_init.m",
]

DARWIN_SRCS = [
    "v3.3/glfw/glfw/src/cocoa_time.c",
    "v3.3/glfw/glfw/src/posix_thread.c",
]

DARWIN_LINKOPTS = [
    "-framework OpenGL",
    "-framework Cocoa",
    "-framework IOKit",
    "-framework CoreFoundation",
    "-framework CoreVideo",
    "-x objective-c",
]
LINUX_DEFINES = [
    "-D_GLFW_HAS_XF86VM",
    "-D_GLFW_X11",
    "-std=c11",
]

LINUX_SRCS = [
    "v3.3/glfw/glfw/src/glx_context.c",
    "v3.3/glfw/glfw/src/linux_joystick.c",
    "v3.3/glfw/glfw/src/posix_thread.c",
    "v3.3/glfw/glfw/src/posix_time.c",
    "v3.3/glfw/glfw/src/x11_init.c",
    "v3.3/glfw/glfw/src/x11_monitor.c",
    "v3.3/glfw/glfw/src/x11_window.c",
]

LINUX_HDRS = [
    "v3.3/glfw/glfw/src/glx_context.h",
    "v3.3/glfw/glfw/src/linux_joystick.h",
    "v3.3/glfw/glfw/src/posix_thread.h",
    "v3.3/glfw/glfw/src/posix_time.h",
    "v3.3/glfw/glfw/src/x11_platform.h",
]

LINUX_LINKOPTS = [
    "-lX11",
    "-lXext",
    "-lXdamage",
    "-lXfixes",
    "-ldl",
    "-lGL",
    "-std=c11",
]

cc_library(
    name = "glfw",
    hdrs = [
        "v3.3/glfw/glfw/include/GLFW/glfw3.h",
        "v3.3/glfw/glfw/include/GLFW/glfw3native.h",
        "v3.3/glfw/glfw/src/egl_context.h",
        "v3.3/glfw/glfw/src/internal.h",
        "v3.3/glfw/glfw/src/osmesa_context.h",
        "v3.3/glfw/glfw/src/mappings.h",
        "v3.3/glfw/glfw/src/xkb_unicode.h",
    ] + select({
        "@io_bazel_rules_go//go/platform:darwin": DARWIN_HDRS,
        "@io_bazel_rules_go//go/platform:linux": LINUX_HDRS,
        "@io_bazel_rules_go//go/platform:windows": WIN32_HDRS}),
    srcs = [
        "v3.3/glfw/glfw/src/context.c",
        "v3.3/glfw/glfw/src/egl_context.c",
        "v3.3/glfw/glfw/src/init.c",
        "v3.3/glfw/glfw/src/input.c",
        "v3.3/glfw/glfw/src/osmesa_context.c",
        "v3.3/glfw/glfw/src/monitor.c",
        "v3.3/glfw/glfw/src/vulkan.c",
        "v3.3/glfw/glfw/src/window.c",
        "v3.3/glfw/glfw/src/xkb_unicode.c",
    ] + select({
        "@io_bazel_rules_go//go/platform:darwin": DARWIN_SRCS,
        "@io_bazel_rules_go//go/platform:linux": LINUX_SRCS,
        "@io_bazel_rules_go//go/platform:windows": WIN32_SRCS}),
    linkopts = select({
        "@io_bazel_rules_go//go/platform:darwin": DARWIN_LINKOPTS,
        "@io_bazel_rules_go//go/platform:linux": LINUX_LINKOPTS,
        "@io_bazel_rules_go//go/platform:windows": WIN32_LINKOPTS,
    }),
    copts = select({
        "@io_bazel_rules_go//go/platform:darwin": DARWIN_DEFINES,
        "@io_bazel_rules_go//go/platform:linux": LINUX_DEFINES,
        "@io_bazel_rules_go//go/platform:windows": WIN32_DEFINES}),
    visibility = ["//visibility:public"],
)

go_library(
    name = "glfw-go",
    srcs = [
        "v3.3/glfw/glfw/include/GLFW/glfw3.h",
        "v3.3/glfw/glfw/include/GLFW/glfw3native.h",
        "v3.3/glfw/build_cgo_hack.go",
        "v3.3/glfw/build.go",
        "v3.3/glfw/error.go",
        "v3.3/glfw/glfw_tree_rebuild.go",
        "v3.3/glfw/glfw.go",
        "v3.3/glfw/input.go",
        "v3.3/glfw/monitor.go",
        "v3.3/glfw/time.go",
        "v3.3/glfw/util.go",
        "v3.3/glfw/vulkan.go",
        "v3.3/glfw/window.go",
        "v3.3/glfw/context.go",
        "v3.3/glfw/input.c",
        "v3.3/glfw/monitor.c",
        "v3.3/glfw/error.c",
        "v3.3/glfw/window.c",
    ] + select({
        "@io_bazel_rules_go//go/platform:darwin": [
            "v3.3/glfw/glfw/src/cocoa_window.m",
            "v3.3/glfw/glfw/src/cocoa_monitor.m",
            "v3.3/glfw/glfw/src/cocoa_joystick.m",
            "v3.3/glfw/glfw/src/nsgl_context.m",
            "v3.3/glfw/glfw/src/cocoa_init.m",
            "v3.3/glfw/native_darwin.go",
        ],
    }),
    cdeps = [":glfw"], 
    cgo = True,
    clinkopts = select({
        "@io_bazel_rules_go//go/platform:darwin": [
            "-framework Cocoa -framework IOKit -framework CoreVideo",
            "-Wno-implicit-function-declaration",
            "-framework OpenGL",
            "-x objective-c",
        ],
        "@io_bazel_rules_go//go/platform:linux": LINUX_LINKOPTS,
        "@io_bazel_rules_go//go/platform:windows": WIN32_LINKOPTS,
    }),
    copts = select({
        "@io_bazel_rules_go//go/platform:darwin": DARWIN_DEFINES,
        "@io_bazel_rules_go//go/platform:linux": LINUX_DEFINES,
        "@io_bazel_rules_go//go/platform:windows": WIN32_DEFINES}) + [
            "-Wimplicit-function-declaration"],
    importpath = "github.com/go-gl/glfw/v3.3/glfw",
    visibility = ["//visibility:public"],
)



alias(
    name = "go_default_library",
    actual = ":glfw-go",
    visibility = ["//visibility:public"],
)
