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

http_archive(
    name = "net_zlib_zlib",
    sha256 = "6d4d6640ca3121620995ee255945161821218752b551a1a180f4215f7d124d45",
    build_file = "@//:third-party/zlib/zlib.BUILD",
    strip_prefix = "zlib-cacf7f1d4e3d44d871b605da3b647f07d718623f",
    urls = [
        "https://mirror.bazel.build/github.com/madler/zlib/archive/cacf7f1d4e3d44d871b605da3b647f07d718623f.tar.gz",
        "https://github.com/madler/zlib/archive/cacf7f1d4e3d44d871b605da3b647f07d718623f.tar.gz",
    ],
)

http_archive(
    name = "org_libpng_libpng",
    build_file = "@//:third-party/libpng/libpng.BUILD",
    sha256 = "7f415186d38ca71c23058386d7cf5135c8beda821ee1beecdc2a7a26c0356615",
    strip_prefix = "libpng-1.2.57",
    urls = [
        "https://mirror.bazel.build/github.com/glennrp/libpng/archive/v1.2.57.tar.gz",
        "https://github.com/glennrp/libpng/archive/v1.2.57.tar.gz",
    ],
)

http_archive(
    name = "org_freetype_freetype2",
    build_file = "@//:third-party/freetype2/freetype2.BUILD",
    sha256 = "33a28fabac471891d0523033e99c0005b95e5618dc8ffa7fa47f9dadcacb1c9b",
    strip_prefix = "freetype-2.8",
    urls = [
        "https://mirror.bazel.build/download.savannah.gnu.org/releases/freetype/freetype-2.8.tar.gz",
        "http://download.savannah.gnu.org/releases/freetype/freetype-2.8.tar.gz",
    ],
)

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")
load("@rules_proto//proto:repositories.bzl", "rules_proto_dependencies", "rules_proto_toolchains")
load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

# new_git_repository(
#     name = "glfw",
#     build_file = "@//:third-party/glfw/glfw.BUILD",
#     commit = "8d7e5cdb49a1a5247df612157ecffdd8e68923d2",
#     remote = "https://github.com/glfw/glfw.git",
# )

new_git_repository(
    name = "glm",
    build_file = "@//:third-party/glm/glm.BUILD",
    commit = "658d8960d081e0c9c312d49758c7ef919371b428",
    remote = "https://github.com/g-truc/glm.git",
)

# new_git_repository(
#     name = "imgui",
#     build_file = "@//:third-party/imgui/imgui.BUILD",
#     commit = "dea92bb7231cf3441e30b373be87e2655b38a113",
#     remote = "https://github.com/inkyblackness/imgui-go.git",
# )

new_git_repository(
    name = "com_github_inkyblackness_imgui_go_v2",
    remote = "https://github.com/inkyblackness/imgui-go.git",
    commit = "622bfc4ce89aef7136d3e04cd9d220105a474a4a",
    build_file = "@//:third-party/imgui/imgui.BUILD", 
)

new_git_repository(
    name = "com_github_go_gl_glfw_v3_3_glfw",
    build_file = "@//:third-party/glfw/glfw.BUILD",
    remote = "https://github.com/go-gl/glfw.git",
    commit = "748e38ca8aecd6b0646cba97eed16259a4af568f",
)

    # go_repository(
        
    #     # importpath = "github.com/go-gl/glfw/v3.3/glfw",
        
    #     # patches = ["//third-party/glfw:glfw.patch"],  #keep
    #     # patch_args = ["-p0"],
    #     sum = "h1:3FLiRYO6PlQFDpUU7OEFlWgjGD1jnBIVSJ5SYRWk+9c=",
    #     version = "v0.0.0-20211213063430-748e38ca8aec",
    # )


local_repository(
    name = "ebiten_imgui",
    path = "third-party/ebiten_imgui"
)

load("//:deps.bzl", "go_repositories")

# gazelle:repository_macro deps.bzl%go_repositories
go_repositories()


# Proto workspace

rules_proto_dependencies()
rules_proto_toolchains()


# Golang workspace

go_rules_dependencies()
go_register_toolchains(version = "1.17.2")

gazelle_dependencies()

