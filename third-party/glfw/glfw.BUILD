load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "glfw",
    srcs = glob(["v3.3/**/*.h", "v3.3/**/*.m", "v3.3/**/*.c", "v3.3/**/*.go"], ["v3.3/glfw/glfw/deps/vs2008/*"]),
    cgo = True,
    clinkopts = select({
    #     "@io_bazel_rules_go//go/platform:android": [
    #         "-lGL",
    #         "-lX11 -lXrandr -lXxf86vm -lXi -lXcursor -lm -lXinerama -ldl -lrt",
    #     ],
        "@io_bazel_rules_go//go/platform:darwin": [
            "-framework Cocoa -framework IOKit -framework CoreVideo",
            "-Wno-implicit-function-declaration",
            "-framework OpenGL",
            "-x objective-c",
        ],
    #     "@io_bazel_rules_go//go/platform:ios": [
    #         "-framework Cocoa -framework IOKit -framework CoreVideo",
    #         "-framework OpenGL",
    #     ],
    #     "@io_bazel_rules_go//go/platform:linux": [
    #         "-lGL",
    #         "-lX11 -lXrandr -lXxf86vm -lXi -lXcursor -lm -lXinerama -ldl -lrt",
    #     ],
    #     "@io_bazel_rules_go//go/platform:windows": [
    #         "-lgdi32",
    #         "-lopengl32",
    #     ],
    #     "//conditions:default": [
              
    #     ],
    }),
    copts = select({
    #     "@io_bazel_rules_go//go/platform:android": [
    #         "-D_GLFW_X11 -D_GNU_SOURCE",
    #     ],
        "@io_bazel_rules_go//go/platform:darwin": [
            "-D_GLFW_COCOA -Wno-deprecated-declarations",
            "-Wno-implicit-function-declaration",
            # "-D_GLFW_USE_OPENGL",
            # "-D_GLFW_NSGL",
            "-x objective-c",
            # "-std=c89",
        ],
    #     "@io_bazel_rules_go//go/platform:ios": [
    #         "-D_GLFW_COCOA -Wno-deprecated-declarations",
    #         "-x objective-c",
    #     ],
    #     "@io_bazel_rules_go//go/platform:linux": [
    #         "-D_GLFW_X11 -D_GNU_SOURCE",
    #     ],
    #     "@io_bazel_rules_go//go/platform:windows": [
    #         "-D_GLFW_WIN32 -Iglfw/deps/mingw",
    #     ],
    #     "//conditions:default": [],
    }),
    importpath = "github.com/go-gl/glfw/v3.3/glfw",
    visibility = ["//visibility:public"],
)

alias(
    name = "go_default_library",
    actual = ":glfw",
    visibility = ["//visibility:public"],
)
