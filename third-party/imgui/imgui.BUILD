cc_library(
    name = "imgui_src",
    hdrs = glob(["**/*.h"]),
    srcs = glob(["**/*.cpp", "wrapper.cpp"]),
    includes = [
        "imgui",
        "wrapper"
    ],
        linkopts = ["-lpthread"],
    visibility = ["//visibility:public"],
)

cc_library(
    name = "imgui",
    deps = [":imgui_src"],
    strip_include_prefix="external/imgui",
    visibility = ["//visibility:public"],
)