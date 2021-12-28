load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix github.com/YarikRevich/hide-seek-client
gazelle(name = "gazelle")

gazelle(
    name = "update-repos-gazelle",
    args = [
        "-from_file=go.mod",
        "-prune",
    ],
    command = "update-repos",
)
