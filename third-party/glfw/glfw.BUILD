load("@io_bazel_rules_go//go:def.bzl", "go_library")

WIN32_DEFINES = [
    "_GLFW_WIN32",
]

WIN32_HDRS = [
    "src/win32_joystick.h",
    "src/win32_platform.h",
    "src/wgl_context.h", 
]

WIN32_SRCS = [
    "src/win32_init.c",
    "src/win32_joystick.c",
    "src/win32_monitor.c",
    "src/win32_thread.c",
    "src/win32_time.c",
    "src/win32_window.c",
    "src/wgl_context.c",
]

WIN32_LINKOPTS = [
    "-DEFAULTLIB:user32.lib",
    "-DEFAULTLIB:gdi32.lib",
    "-DEFAULTLIB:shell32.lib",
]


DARWIN_DEFINES = [
    "_GLFW_COCOA",
    "_GLFW_NSGL",
    "_GLFW_NO_DLOAD_WINMM",
    "_GLFW_USE_OPENGL",
]

DARWIN_HDRS = [
    "src/cocoa_joystick.h",
    "src/cocoa_platform.h",
    "src/glx_context.h",
    "src/nsgl_context.h",
    "src/null_joystick.h",
    "src/null_platform.h",
    "src/posix_thread.h",
    "src/wl_platform.h",
]

DARWIN_SRCS = [
    "src/cocoa_time.c",
    "src/posix_thread.c",
]

DARWIN_LINKOPTS = [
    "-framework OpenGL",
    "-framework Cocoa",
    "-framework IOKit",
    "-framework CoreFoundation"
]
LINUX_DEFINES = [
    "-D_GLFW_HAS_XF86VM",
    "-D_GLFW_X11",
    "-std=c11",
]

# LINUX_SRCS = glob([
#     "v3.3/glfw/glfw/src/glx_context.*",
#     "v3.3/glfw/glfw/src/posix_*.*",
#     "v3.3/glfw/glfw/src/wl_init.c",
#     "v3.3/glfw/glfw/src/wl_platform.h",
#     "v3.3/glfw/glfw/src/x11_*.*",
#     "v3.3/glfw/glfw/src/xkb_*.*",
#     "v3.3/glfw/glfw/src/egl_*.*",
#     "v3.3/glfw/glfw/src/osmesa_*.*",
#     "v3.3/glfw/glfw/src/linux_*.*",
# ]) + [
#     "v3.3/glfw/c_glfw_lin.go",
    
# ]

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
    visibility = ["//visibility:public"],
)

# go_library(
#     name = "glfw-go",
#     srcs = [
#             "v3.3/glfw/glfw/include/GLFW/glfw3.h",
#             "v3.3/glfw/glfw/include/GLFW/glfw3native.h",
#             "v3.3/glfw/glfw/src/init.c",
#             "v3.3/glfw/glfw/src/context.c",
#             "v3.3/glfw/glfw/src/input.c",
#             "v3.3/glfw/glfw/src/monitor.c",
#             "v3.3/glfw/glfw/src/vulkan.c",
#             "v3.3/glfw/glfw/src/window.c",
#             "v3.3/glfw/glfw/src/mappings.h",
#             "v3.3/glfw/glfw/src/osmesa_context.c",
#             "v3.3/glfw/error.c",
#             "v3.3/glfw/input.c",
#             "v3.3/glfw/window.c",
#         ] + [
#             "v3.3/glfw/build_cgo_hack.go",
#             "v3.3/glfw/build.go",
#             "v3.3/glfw/c_glfw.go",
#             "v3.3/glfw/error.go",
#             "v3.3/glfw/glfw_tree_rebuild.go",
#             "v3.3/glfw/glfw.go",
#             "v3.3/glfw/input.go",
#             "v3.3/glfw/monitor.go",
#             "v3.3/glfw/time.go",
#             "v3.3/glfw/util.go",
#             "v3.3/glfw/vulkan.go",
#             "v3.3/glfw/window.go",
#         ],
#     # cdeps = [":glfw"],
#     cgo = True,
#     clinkopts = select({
#         "@io_bazel_rules_go//go/platform:darwin": [
#             "-framework Cocoa -framework IOKit -framework CoreVideo",
#             "-Wno-implicit-function-declaration",
#             "-framework OpenGL",
#             "-x objective-c",
#         ],
#         "@io_bazel_rules_go//go/platform:linux": LINUX_LINKOPTS,
#         "@io_bazel_rules_go//go/platform:windows": WIN32_LINKOPTS,
#     }),
#     copts = select({
#         "@io_bazel_rules_go//go/platform:darwin": DARWIN_DEFINES,
#         "@io_bazel_rules_go//go/platform:linux": LINUX_DEFINES,
#         "@io_bazel_rules_go//go/platform:windows": WIN32_DEFINES}) + [
#             "-Wimplicit-function-declaration"],
#     importpath = "github.com/go-gl/glfw/v3.3/glfw",
#     visibility = ["//visibility:public"],
# )


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
    ],
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
