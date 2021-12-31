workspace(name = "local")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
load("@bazel_tools//tools/build_defs/repo:git.bzl", "new_git_repository")

http_archive(
    name = "bazel_gazelle",
    sha256 = "62ca106be173579c0a167deb23358fdfe71ffa1e4cfdddf5582af26520f1c66f",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.23.0/bazel-gazelle-v0.23.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.23.0/bazel-gazelle-v0.23.0.tar.gz",
    ],
)

http_archive(
    name = "rules_proto",
    sha256 = "66bfdf8782796239d3875d37e7de19b1d94301e8972b3cbd2446b332429b4df1",
    strip_prefix = "rules_proto-4.0.0",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_proto/archive/refs/tags/4.0.0.tar.gz",
        "https://github.com/bazelbuild/rules_proto/archive/refs/tags/4.0.0.tar.gz",
    ],
)

http_archive(
    name = "rules_cc",
    sha256 = "4dccbfd22c0def164c8f47458bd50e0c7148f3d92002cdb459c2a96a68498241",
    urls = ["https://github.com/bazelbuild/rules_cc/releases/download/0.0.1/rules_cc-0.0.1.tar.gz"],
)

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "8e968b5fcea1d2d64071872b12737bbb5514524ee5f0a4f54f5920266c261acb",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.28.0/rules_go-v0.28.0.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.28.0/rules_go-v0.28.0.zip",
    ],
)

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")
load("@rules_proto//proto:repositories.bzl", "rules_proto_dependencies", "rules_proto_toolchains")
load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

new_git_repository(
    name = "glfw",
    build_file = "@//:third-party/glfw/glfw.BUILD",
    commit = "8d7e5cdb49a1a5247df612157ecffdd8e68923d2",
    remote = "https://github.com/glfw/glfw.git",
)

new_git_repository(
    name = "glm",
    build_file = "@//:third-party/glm/glm.BUILD",
    commit = "658d8960d081e0c9c312d49758c7ef919371b428",
    remote = "https://github.com/g-truc/glm.git",
)

new_git_repository(
    name = "imgui",
    build_file = "@//:third-party/imgui/imgui.BUILD",
    commit = "dea92bb7231cf3441e30b373be87e2655b38a113",
    remote = "https://github.com/inkyblackness/imgui-go.git",
)

load("//:deps.bzl", "go_repositories")

# gazelle:repository_macro deps.bzl%go_repositories
go_repositories()

go_rules_dependencies()

go_register_toolchains(version = "1.17.2")

gazelle_dependencies()

rules_proto_dependencies()

rules_proto_toolchains()
