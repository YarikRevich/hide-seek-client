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
    if [[ $$OSTYPE == "darwin"* ]]; then\
        log_dir_path="/Users/$$USER/games/HideSeek/log";
        log_file_path="/Users/$$USER/games/HideSeek/log/log.log";
        db_dir_path="/Users/$$USER/games/HideSeek/db";
        db_file_path="/Users/$$USER/games/HideSeek/db/storage.db";
    elif [[ $$OSTYPE == "linux"* ]]; then\
        log_dir_path="/home/$$USER/games/HideSeek/log";
        log_file_path="/home/$$USER/games/HideSeek/log/log.log";
        db_dir_path="/home/$$USER/games/HideSeek/db";
        db_file_path="/home/$$USER/games/HideSeek/db/storage.db";
    fi;\
    
    if ! [[ -d $$log_dir_path ]]; then\
        mkdir -p $$log_dir_path;\
    fi;\
    if ! [[ -d $$db_dir_path ]]; then\
        mkdir -p $$db_dir_path;\
    fi;\

    if ! [[ -f $$log_file_path ]]; then\
        touch $$log_file_path;\
    fi;\
    if ! [[ -f $$db_file_path ]]; then\
        touch $$db_file_path;\
    fi;\

    if [[ $$OSTYPE == "darwin"* ]]; then\
        chmod -R 777 /Users/$$USER/games;\
    elif [[ $$OSTYPE == "linux"* ]]; then\
        chmod -R 777 /home/$$USER/games;\
    fi;\
    echo EOF > $@;
"""

genrule(
    name = "service_env",
    cmd = SERVICE_ENV,
    outs = ["service_env.stub"],
)
