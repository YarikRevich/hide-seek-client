load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "imgui-go",
    srcs = glob(["*.cpp", "*.go", "*.h"], ["*_test.go"]),
    # srcs = [
    #     "wrapper/WrapperConverter.cpp",
    #     "wrapper.cpp",
    #     "wrapper/Color.cpp",
    #     "wrapper/Context.cpp",
    #     "wrapper/Focus.cpp",
    #     "wrapper/DragDrop.cpp",
    #     "wrapper/DrawCommand.cpp",
    #     "wrapper/DrawData.cpp",
    #     "wrapper/DrawList.cpp",
    #     "wrapper/Font.cpp",
    #     "wrapper/FontAtlas.cpp",
    #     "wrapper/FontConfig.cpp",
    #     "wrapper/InputTextCallbackData.cpp",
    #     "wrapper/IO.cpp",
    #     "wrapper/Layout.cpp",
    #     "wrapper/ListClipper.cpp",
    #     "wrapper/Main.cpp",
    #     "wrapper/Popup.cpp",
    #     "wrapper/Scroll.cpp",
    #     "wrapper/State.cpp",
    #     "wrapper/Style.cpp",
    #     "wrapper/Tables.cpp",
    #     "wrapper/Widgets.cpp",
    #     "wrapper/Window.cpp",
    #     "wrapper/Settings.cpp",
    #     "imgui/imgui.cpp",
    #     "imgui/imgui_demo.cpp",
    #     "imgui/imgui_draw.cpp",
    #     "imgui/imgui_tables.cpp",
    #     "imgui/imgui_widgets.cpp",
    #     "imgui/misc/freetype/imgui_freetype.cpp"] + [
    #     "wrapper.go",
    #     "AllocatedGlyphRanges.go",
    #     "Assert.go",
    #     "Color.go",
    #     "Condition.go",
    #     "Context.go",
    #     "DragDrop.go",
    #     "DrawCommand.go",
    #     "DrawData.go",
    #     "DrawList.go",
    #     "Focus.go",
    #     "Font.go",
    #     "FontAtlas.go",
    #     "FontConfig.go",
    #     "FreeType.go",
    #     "GlyphRanges.go",
    #     "IO.go",
    #     "InputTextCallbackData.go",
    #     "Layout.go",
    #     "ListClipper.go",
    #     "Main.go",
    #     "PackedColor.go",
    #     "Popup.go",
    #     "Scroll.go",
    #     "Settings.go",
    #     "State.go",
    #     "Style.go",
    #     "Tables.go",
    #     "TextureID.go",
    #     "Vectors.go",
    #     "Widgets.go",
    #     "Window.go",
    #     "WrapperConverter.go"] + glob(["**/*.h"]),
    
    importpath = "github.com/inkyblackness/imgui-go/v2",
    cgo = True,
    cdeps = ["@org_freetype_freetype2//:freetype2"],
    visibility = ["//visibility:public"],
)

alias(
    name = "go_default_library",
    actual = ":imgui-go",
    visibility = ["//visibility:public"],
)