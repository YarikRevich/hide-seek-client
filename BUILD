load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix github.com/YarikRevich/hide-seek-client
gazelle(name = "gazelle")

gazelle(
    name = "update-repos-gazelle",
    args = [
        "-from_file=go.mod",
        "-to_macro=deps.bzl%go_repositories",
        "-prune",
    ],
    command = "update-repos",
)

DEPS = """
    if [[ $$OSTYPE == "linux"* ]]; then\
        apt-get install libxcursor-dev libxrandr-dev libxinerama-dev libxi-dev -y;\
        apt-get install libgl1-mesa-dev -y;\
        apt-get install libwayland-dev -y;\
        apt-get install libasound2-dev -y;\
    fi;\
    echo EOF > $@;
"""

genrule(
    name = "deps",
    outs = ["deps.stub"],
    cmd = DEPS,
)

SERVICE_ENV = """
    if ! [[ -d "/home/$$USER/games/HideSeek/log" ]]; then\
        mkdir -p /home/$$USER/games/HideSeek/log;\
    fi;\
    if ! [[ -d "/home/$$USER/games/HideSeek/db" ]]; then\
        mkdir -p /home/$$USER/games/HideSeek/db;\
    fi;\
    if ! [[ -f "/home/$$USER/games/HideSeek/log/log.log" ]]; then\
        touch /home/$$USER/games/HideSeek/log/log.log;\
    fi;\
    chmod 666 /home/$$USER/games/HideSeek/log/log.log;\
    echo EOF > $@;
"""

genrule(
    name = "service_env",
    cmd = SERVICE_ENV,
    outs = ["service_env.stub"],
)
