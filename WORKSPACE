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
    urls = ["https://github.com/bazelbuild/rules_cc/releases/download/0.0.1/rules_cc-0.0.1.tar.gz"],
    sha256 = "4dccbfd22c0def164c8f47458bd50e0c7148f3d92002cdb459c2a96a68498241",
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

go_repository(
    name = "co_honnef_go_tools",
    importpath = "honnef.co/go/tools",
    sum = "h1:/hemPrYIhOhy8zYrNj+069zDB68us2sMGsfkFJO0iZs=",
    version = "v0.0.0-20190523083050-ea95bdfd59fc",
)

go_repository(
    name = "com_github_alimasyhur_is_connect",
    importpath = "github.com/alimasyhur/is-connect",
    sum = "h1:W51dH0udo7NVhSJJI6hMCI2ELQlqWhTre8yRbFBo5/I=",
    version = "v0.0.0-20180112042527-a7e9ece095d0",
)

go_repository(
    name = "com_github_antihax_optional",
    importpath = "github.com/antihax/optional",
    sum = "h1:xK2lYat7ZLaVVcIuj82J8kIro4V6kDe0AUDFboUCwcg=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_atotto_clipboard",
    importpath = "github.com/atotto/clipboard",
    sum = "h1:EH0zSVneZPSuFR11BlR9YppQTVDbh5+16AmcJi4g1z4=",
    version = "v0.1.4",
)

go_repository(
    name = "com_github_burntsushi_toml",
    importpath = "github.com/BurntSushi/toml",
    sum = "h1:GaI7EiDXDRfa8VshkTj7Fym7ha+y8/XxIgD2okUIjLw=",
    version = "v0.4.1",
)

go_repository(
    name = "com_github_burntsushi_xgb",
    importpath = "github.com/BurntSushi/xgb",
    sum = "h1:1BDTz0u9nC3//pOCMdNH+CiXJVYJh5UQNCOBG7jbELc=",
    version = "v0.0.0-20160522181843-27f122750802",
)

go_repository(
    name = "com_github_census_instrumentation_opencensus_proto",
    importpath = "github.com/census-instrumentation/opencensus-proto",
    sum = "h1:glEXhBS5PSLLv4IXzLA5yPRVX4bilULVyxxbrfOtDAk=",
    version = "v0.2.1",
)

go_repository(
    name = "com_github_cespare_xxhash_v2",
    importpath = "github.com/cespare/xxhash/v2",
    sum = "h1:6MnRN8NT7+YBpUIWxHtefFZOKTAPgGjpQSxqLNn0+qY=",
    version = "v2.1.1",
)

go_repository(
    name = "com_github_client9_misspell",
    importpath = "github.com/client9/misspell",
    sum = "h1:ta993UF76GwbvJcIo3Y68y/M3WxlpEHPWIGDkJYwzJI=",
    version = "v0.3.4",
)

go_repository(
    name = "com_github_cncf_udpa_go",
    importpath = "github.com/cncf/udpa/go",
    sum = "h1:hzAQntlaYRkVSFEfj9OTWlVV1H155FMD8BTKktLv0QI=",
    version = "v0.0.0-20210930031921-04548b0d99d4",
)

go_repository(
    name = "com_github_cncf_xds_go",
    importpath = "github.com/cncf/xds/go",
    sum = "h1:zH8ljVhhq7yC0MIeUL/IviMtY8hx2mK8cN9wEYb8ggw=",
    version = "v0.0.0-20211011173535-cb28da3451f1",
)

go_repository(
    name = "com_github_davecgh_go_spew",
    importpath = "github.com/davecgh/go-spew",
    sum = "h1:vj9j/u1bqnvCEfJOwUhtlOARqs3+rkHYY13jYWTU97c=",
    version = "v1.1.1",
)

go_repository(
    name = "com_github_engoengine_glm",
    importpath = "github.com/engoengine/glm",
    sum = "h1:rVKp3XvwVYVidU1oOcahpRdoG32EJ1+CLxiwZOADpNc=",
    version = "v0.0.0-20170725114841-9c08f4d1f668",
)

go_repository(
    name = "com_github_engoengine_math",
    importpath = "github.com/EngoEngine/math",
    sum = "h1:ejDfSg48ynB9T6btiu9EHjZmpQgW/zHf3IeC7SqXXv8=",
    version = "v1.0.4",
)

go_repository(
    name = "com_github_envoyproxy_go_control_plane",
    importpath = "github.com/envoyproxy/go-control-plane",
    sum = "h1:fP+fF0up6oPY49OrjPrhIJ8yQfdIM85NXMLkMg1EXVs=",
    version = "v0.9.10-0.20210907150352-cf90f659a021",
)

go_repository(
    name = "com_github_envoyproxy_protoc_gen_validate",
    importpath = "github.com/envoyproxy/protoc-gen-validate",
    sum = "h1:EQciDnbrYxy13PgWoY8AqoxGiPrpgBZ1R8UNe3ddc+A=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_faiface_beep",
    importpath = "github.com/faiface/beep",
    sum = "h1:UB5DiRNmA4erfUYnHbgU4UB6DlBOrsdEFRtcc8sCkdQ=",
    version = "v1.0.2",
)

go_repository(
    name = "com_github_franela_goblin",
    importpath = "github.com/franela/goblin",
    sum = "h1:NrF81UtW8gG2LBGkXFQFqlfNnvMt9WdB46sfdJY4oqc=",
    version = "v0.0.0-20211003143422-0a4f594942bf",
)

go_repository(
    name = "com_github_gabstv_ebiten_imgui",
    importpath = "github.com/gabstv/ebiten-imgui",
    sum = "h1:isljlX+7ObvOGiRRe0IUltWcoioNUj2DX+StcaaO/9M=",
    version = "v0.2.0",
)

go_repository(
    name = "com_github_gdamore_encoding",
    importpath = "github.com/gdamore/encoding",
    sum = "h1:+7OoQ1Bc6eTm5niUzBa0Ctsh6JbMW6Ra+YNuAtDBdko=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_gdamore_tcell",
    importpath = "github.com/gdamore/tcell",
    sum = "h1:U73YL+jMem2XfhvaIUfPO6MpJawaG92B2funXVb9qLs=",
    version = "v1.1.1",
)

go_repository(
    name = "com_github_ghodss_yaml",
    importpath = "github.com/ghodss/yaml",
    sum = "h1:wQHKEahhL6wmXdzwWG11gIVCkOv05bNOh+Rxn0yngAk=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_go_gl_gl",
    importpath = "github.com/go-gl/gl",
    sum = "h1:zDw5v7qm4yH7N8C8uWd+8Ii9rROdgWxQuGoJ9WDXxfk=",
    version = "v0.0.0-20211210172815-726fda9656d6",
)

go_repository(
    name = "com_github_go_gl_glfw",
    importpath = "github.com/go-gl/glfw",
    sum = "h1:um18JldLG6QwC9tj6mSfQnb+kor5aezfPPtq1GmHek0=",
    version = "v0.0.0-20211213063430-748e38ca8aec",
)

go_repository(
    name = "com_github_go_gl_glfw_v3_3_glfw",
    importpath = "github.com/go-gl/glfw/v3.3/glfw",
    sum = "h1:3FLiRYO6PlQFDpUU7OEFlWgjGD1jnBIVSJ5SYRWk+9c=",
    version = "v0.0.0-20211213063430-748e38ca8aec",
)

go_repository(
    name = "com_github_go_gl_glow",
    importpath = "github.com/go-gl/glow",
    sum = "h1:7GinxKZ0LC09PnNQLCU8cHsvDUr4uzkx1P22Ji9UAYY=",
    version = "v0.0.0-20211208232303-9d81eb29c711",
)

go_repository(
    name = "com_github_go_ping_ping",
    importpath = "github.com/go-ping/ping",
    sum = "h1:wtjTfjwAR/BYYMJ+QOLI/3J/qGEI0fgrkZvgsEWK2/Q=",
    version = "v0.0.0-20210911151512-381826476871",
)

go_repository(
    name = "com_github_gofrs_flock",
    importpath = "github.com/gofrs/flock",
    sum = "h1:MSdYClljsF3PbENUUEx85nkWfJSGfzYI9yEBZOJz6CY=",
    version = "v0.8.0",
)

go_repository(
    name = "com_github_golang_freetype",
    importpath = "github.com/golang/freetype",
    sum = "h1:DACJavvAHhabrF08vX0COfcOBJRhZ8lUbR+ZWIs0Y5g=",
    version = "v0.0.0-20170609003504-e2365dfdc4a0",
)

go_repository(
    name = "com_github_golang_glog",
    importpath = "github.com/golang/glog",
    sum = "h1:VKtxabqXZkF25pY9ekfRL6a582T4P37/31XEstQ5p58=",
    version = "v0.0.0-20160126235308-23def4e6c14b",
)

go_repository(
    name = "com_github_golang_mock",
    importpath = "github.com/golang/mock",
    sum = "h1:ErTB+efbowRARo13NNdxyJji2egdxLGQhRaY+DUumQc=",
    version = "v1.6.0",
)

go_repository(
    name = "com_github_golang_protobuf",
    importpath = "github.com/golang/protobuf",
    sum = "h1:JjCZWpVbqXDqFVmTfYWEVTMIYrL/NPdPSCHPJ0T/raM=",
    version = "v1.4.3",
)

go_repository(
    name = "com_github_google_go_cmp",
    importpath = "github.com/google/go-cmp",
    sum = "h1:/QaMHBdZ26BB3SSst0Iwl10Epc+xhTquomWX0oZEB6w=",
    version = "v0.5.0",
)

go_repository(
    name = "com_github_google_uuid",
    importpath = "github.com/google/uuid",
    sum = "h1:t6JiXgmwXMjEs8VusXIJk2BXHsn+wx8BZdTaoZ5fu7I=",
    version = "v1.3.0",
)

go_repository(
    name = "com_github_gopherjs_gopherjs",
    importpath = "github.com/gopherjs/gopherjs",
    sum = "h1:16eHWuMGvCjSfgRJKqIzapE78onvvTbdi1rMkU00lZw=",
    version = "v0.0.0-20180825215210-0210a2f0f73c",
)

go_repository(
    name = "com_github_gopherjs_gopherwasm",
    importpath = "github.com/gopherjs/gopherwasm",
    sum = "h1:32nge/RlujS1Im4HNCJPp0NbBOAeBXFuT1KonUuLl+Y=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_grpc_ecosystem_grpc_gateway",
    importpath = "github.com/grpc-ecosystem/grpc-gateway",
    sum = "h1:gmcG1KaJ57LophUzW0Hy8NmPhnMZb4M0+kPpLofRdBo=",
    version = "v1.16.0",
)

go_repository(
    name = "com_github_hajimehoshi_bitmapfont_v2",
    importpath = "github.com/hajimehoshi/bitmapfont/v2",
    sum = "h1:JefUkL0M4nrdVwVq7MMZxSTh6mSxOylm+C4Anoucbb0=",
    version = "v2.1.3",
)

go_repository(
    name = "com_github_hajimehoshi_ebiten_v2",
    importpath = "github.com/hajimehoshi/ebiten/v2",
    sum = "h1:yx8g5YQy7xnVbT4lCZCAQHx454j50emlRs6Aa78vdPc=",
    version = "v2.1.5",
)

go_repository(
    name = "com_github_hajimehoshi_file2byteslice",
    importpath = "github.com/hajimehoshi/file2byteslice",
    sum = "h1:4IP7CPObI35+mQShFOYg2JMHDJKciLTW5599inhFfkA=",
    version = "v0.0.0-20200812174855-0e5e8a80490e",
)

go_repository(
    name = "com_github_hajimehoshi_go_mp3",
    importpath = "github.com/hajimehoshi/go-mp3",
    sum = "h1:xSYNE2F3lxtOu9BRjCWHHceg7S91IHfXfXp5+LYQI7s=",
    version = "v0.3.2",
)

go_repository(
    name = "com_github_hajimehoshi_oto",
    importpath = "github.com/hajimehoshi/oto",
    sum = "h1:I7maFPz5MBCwiutOrz++DLdbr4rTzBsbBuV2VpgU9kk=",
    version = "v0.7.1",
)

go_repository(
    name = "com_github_inkyblackness_imgui_go",
    importpath = "github.com/inkyblackness/imgui-go",
    sum = "h1:uaxSM5SbbqCTGEx5ig7B2J78hM3g3az4f5NC6b4J7lY=",
    version = "v1.12.0",
)

go_repository(
    name = "com_github_inkyblackness_imgui_go_v2",
    importpath = "github.com/inkyblackness/imgui-go/v2",
    sum = "h1:yD6m1xqTNl/HF0M4ceN1HNhUKRrk+SU7nbGs2KNOXUA=",
    version = "v2.4.1",
)

go_repository(
    name = "com_github_inkyblackness_imgui_go_v4",
    importpath = "github.com/inkyblackness/imgui-go/v4",
    sum = "h1:jY32Xl18aRwTBXaDfyefCmPDxJOtGM8kGfu/kMNJpbE=",
    version = "v4.4.0",
)

go_repository(
    name = "com_github_jakecoffman_cp",
    importpath = "github.com/jakecoffman/cp",
    sum = "h1:bhKvCNbAddYegYHSV5abG3G23vZdsISgqXa4X/lK8Oo=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_jfreymuth_oggvorbis",
    importpath = "github.com/jfreymuth/oggvorbis",
    sum = "h1:MLNGGyhOMiVcvea9Dp5+gbs2SAwqwQbtrWnonYa0M0Y=",
    version = "v1.0.3",
)

go_repository(
    name = "com_github_jfreymuth_vorbis",
    importpath = "github.com/jfreymuth/vorbis",
    sum = "h1:m1xH6+ZI4thH927pgKD8JOH4eaGRm18rEE9/0WKjvNE=",
    version = "v1.0.2",
)

go_repository(
    name = "com_github_kr_pretty",
    importpath = "github.com/kr/pretty",
    sum = "h1:Fmg33tUaq4/8ym9TJN1x7sLJnHVwhP33CNkpYV/7rwI=",
    version = "v0.2.1",
)

go_repository(
    name = "com_github_kr_pty",
    importpath = "github.com/kr/pty",
    sum = "h1:VkoXIwSboBpnk99O/KFauAEILuNHv5DVFKZMBN/gUgw=",
    version = "v1.1.1",
)

go_repository(
    name = "com_github_kr_text",
    importpath = "github.com/kr/text",
    sum = "h1:45sCR5RtlFHMR4UwH9sdQ5TC8v0qDQCHnXt+kaKSTVE=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_lucasb_eyer_go_colorful",
    importpath = "github.com/lucasb-eyer/go-colorful",
    sum = "h1:5MnxBC15uMxFv5FY/J/8vzyaBiArCOkMdFT9Jsw78iY=",
    version = "v0.0.0-20181028223441-12d3b2882a08",
)

go_repository(
    name = "com_github_mattn_go_runewidth",
    importpath = "github.com/mattn/go-runewidth",
    sum = "h1:2BvfKmzob6Bmd4YsL0zygOqfdFnK7GR4QL06Do4/p7Y=",
    version = "v0.0.4",
)

go_repository(
    name = "com_github_mattn_go_sqlite3",
    importpath = "github.com/mattn/go-sqlite3",
    sum = "h1:gDp86IdQsN/xWjIEmr9MF6o9mpksUgh0fu+9ByFxzIU=",
    version = "v1.14.8",
)

go_repository(
    name = "com_github_mbndr_figlet4go",
    importpath = "github.com/mbndr/figlet4go",
    sum = "h1:mQncVDBpKkAecPcH2IMGpKUQYhwowlafQbfkz2QFqkc=",
    version = "v0.0.0-20190224160619-d6cef5b186ea",
)

go_repository(
    name = "com_github_mewkiz_flac",
    importpath = "github.com/mewkiz/flac",
    sum = "h1:dHGW/2kf+/KZ2GGqSVayNEhL9pluKn/rr/h/QqD9Ogc=",
    version = "v1.0.5",
)

go_repository(
    name = "com_github_niemeyer_pretty",
    importpath = "github.com/niemeyer/pretty",
    sum = "h1:fD57ERR4JtEqsWbfPhv4DMiApHyliiK5xCTNVSPiaAs=",
    version = "v0.0.0-20200227124842-a10e7caefd8e",
)

go_repository(
    name = "com_github_pkg_browser",
    importpath = "github.com/pkg/browser",
    sum = "h1:49lOXmGaUpV9Fz3gd7TFZY106KVlPVa5jcYD1gaQf98=",
    version = "v0.0.0-20180916011732-0a3d74bf9ce4",
)

go_repository(
    name = "com_github_pkg_errors",
    importpath = "github.com/pkg/errors",
    sum = "h1:FEBLx1zS214owpjy7qsBeixbURkuhQAwrK5UwLGTwt4=",
    version = "v0.9.1",
)

go_repository(
    name = "com_github_pmezard_go_difflib",
    importpath = "github.com/pmezard/go-difflib",
    sum = "h1:4DBwDE0NGyQoBHbLQYPwSUPoCMWR5BEzIk/f1lZbAQM=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_prometheus_client_model",
    importpath = "github.com/prometheus/client_model",
    sum = "h1:gQz4mCbXsO+nc9n1hCxHcGA3Zx3Eo+UHZoInFGUIXNM=",
    version = "v0.0.0-20190812154241-14fe0d1b01d4",
)

go_repository(
    name = "com_github_rogpeppe_fastuuid",
    importpath = "github.com/rogpeppe/fastuuid",
    sum = "h1:Ppwyp6VYCF1nvBTXL3trRso7mXMlRrw9ooo375wvi2s=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_rs_xid",
    importpath = "github.com/rs/xid",
    sum = "h1:6NjYksEUlhurdVehpc7S7dk6DAmcKv8V9gG0FsVN2U4=",
    version = "v1.3.0",
)

go_repository(
    name = "com_github_sirupsen_logrus",
    importpath = "github.com/sirupsen/logrus",
    sum = "h1:dJKuHgqk1NNQlqoA6BTlM1Wf9DOH3NBjQyu0h9+AZZE=",
    version = "v1.8.1",
)

go_repository(
    name = "com_github_stretchr_objx",
    importpath = "github.com/stretchr/objx",
    sum = "h1:4G4v2dO3VZwixGIRoQ5Lfboy6nUhCyYzaqnIAPPhYs4=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_stretchr_testify",
    importpath = "github.com/stretchr/testify",
    sum = "h1:nwc3DEeHmmLAfoZucVR881uASk0Mfjw8xYJ99tb5CcY=",
    version = "v1.7.0",
)

go_repository(
    name = "com_github_veandco_go_sdl2",
    importpath = "github.com/veandco/go-sdl2",
    sum = "h1:8QoD2bhWl7SbQDflIAUYWfl9Vq+mT8/boJFAUzAScgY=",
    version = "v0.4.10",
)

go_repository(
    name = "com_github_yarikrevich_caching",
    importpath = "github.com/YarikRevich/caching",
    sum = "h1:L4YrgjEy0K4t1KRUP7Eg6o47Q8Xa0bRTi+8IQpjLDj0=",
    version = "v0.0.0-20211008105337-1500fe296f71",
)

go_repository(
    name = "com_github_yarikrevich_game_networking",
    importpath = "github.com/YarikRevich/game-networking",
    sum = "h1:rYx14LpueMIMfdiBiSmNhAkxD82YRYMaOCIWiL7xTV0=",
    version = "v1.1.4-0.20211015150625-86d96e1221b6",
)

go_repository(
    name = "com_github_yarikrevich_gsl",
    importpath = "github.com/YarikRevich/GSL",
    sum = "h1:DVg0bFn4R47AdOPDDqz0JwVGYNPANZxEhyxh9eH9tfU=",
    version = "v0.0.0-20210814192237-382c0edab67a",
)

go_repository(
    name = "com_github_yarikrevich_wrapper",
    importpath = "github.com/YarikRevich/wrapper",
    sum = "h1:Au84OIFwtF42CEEIrg4mELpTCjNCe/4/z79ueRuIOlo=",
    version = "v0.0.0-20210919192625-fd9de6c41066",
)

go_repository(
    name = "com_github_yuin_goldmark",
    importpath = "github.com/yuin/goldmark",
    sum = "h1:dPmz1Snjq0kmkz159iL7S6WzdahUTHnHB5M56WFVifs=",
    version = "v1.3.5",
)

go_repository(
    name = "com_google_cloud_go",
    importpath = "cloud.google.com/go",
    sum = "h1:eOI3/cP2VTU6uZLDYAoic+eyzzB9YyGmJ7eIjl8rOPg=",
    version = "v0.34.0",
)

go_repository(
    name = "com_shuralyov_dmitri_gpu_mtl",
    importpath = "dmitri.shuralyov.com/gpu/mtl",
    sum = "h1:+PdD6GLKejR9DizMAKT5DpSAkKswvZrurk1/eEt9+pw=",
    version = "v0.0.0-20201218220906-28db891af037",
)

go_repository(
    name = "in_gopkg_check_v1",
    importpath = "gopkg.in/check.v1",
    sum = "h1:Hei/4ADfdWqJk1ZMxUNpqntNwaWcugrBjAiHlqqRiVk=",
    version = "v1.0.0-20201130134442-10cb98267c6c",
)

go_repository(
    name = "in_gopkg_data_dog_go_sqlmock_v1",
    importpath = "gopkg.in/DATA-DOG/go-sqlmock.v1",
    sum = "h1:FVCohIoYO7IJoDDVpV2pdq7SgrMH6wHnuTyrdrxJNoY=",
    version = "v1.3.0",
)

go_repository(
    name = "in_gopkg_yaml_v2",
    importpath = "gopkg.in/yaml.v2",
    sum = "h1:fvjTMHxHEw/mxHbtzPi3JCcKXQRAnQTBRo6YCJSVHKI=",
    version = "v2.2.3",
)

go_repository(
    name = "in_gopkg_yaml_v3",
    importpath = "gopkg.in/yaml.v3",
    sum = "h1:dUUwHk2QECo/6vqA44rthZ8ie2QXMNeKRTHCNY2nXvo=",
    version = "v3.0.0-20200313102051-9f266ea9e77c",
)

go_repository(
    name = "io_opentelemetry_go_proto_otlp",
    importpath = "go.opentelemetry.io/proto/otlp",
    sum = "h1:rwOQPCuKAKmwGKq2aVNnYIibI6wnV7EvzgfTCzcdGg8=",
    version = "v0.7.0",
)

go_repository(
    name = "org_golang_google_appengine",
    importpath = "google.golang.org/appengine",
    sum = "h1:/wp5JvzpHIxhs/dumFmF7BXTf3Z+dd4uXta4kVyO508=",
    version = "v1.4.0",
)

go_repository(
    name = "org_golang_google_genproto",
    importpath = "google.golang.org/genproto",
    sum = "h1:+kGHl1aib/qcwaRi1CbqBZ1rk19r85MNUf8HaBghugY=",
    version = "v0.0.0-20200526211855-cb27e3aa2013",
)

go_repository(
    name = "org_golang_google_grpc",
    importpath = "google.golang.org/grpc",
    sum = "h1:XT2/MFpuPFsEX2fWh3YQtHkZ+WYZFQRfaUgLZYj/p6A=",
    version = "v1.42.0",
)

go_repository(
    name = "org_golang_google_protobuf",
    importpath = "google.golang.org/protobuf",
    sum = "h1:Ejskq+SyPohKW+1uil0JJMtmHCgJPJ/qWTxr8qp+R4c=",
    version = "v1.25.0",
)

go_repository(
    name = "org_golang_x_crypto",
    importpath = "golang.org/x/crypto",
    sum = "h1:psW17arqaxU48Z5kZ0CQnkZWQJsqcURM6tKiBApRjXI=",
    version = "v0.0.0-20200622213623-75b288015ac9",
)

go_repository(
    name = "org_golang_x_exp",
    importpath = "golang.org/x/exp",
    sum = "h1:GnGfrp0fiNhiBS/v/aCFTmfEWgkvxW4Qiu8oM2/IfZ4=",
    version = "v0.0.0-20201221025956-e89b829e73ea",
)

go_repository(
    name = "org_golang_x_image",
    importpath = "golang.org/x/image",
    sum = "h1:RNPAfi2nHY7C2srAV8A49jpsYr0ADedCk1wq6fTMTvs=",
    version = "v0.0.0-20210628002857-a66eb6448b8d",
)

go_repository(
    name = "org_golang_x_lint",
    importpath = "golang.org/x/lint",
    sum = "h1:XQyxROzUlZH+WIQwySDgnISgOivlhjIEwaQaJEJrrN0=",
    version = "v0.0.0-20190313153728-d0100b6bd8b3",
)

go_repository(
    name = "org_golang_x_mobile",
    importpath = "golang.org/x/mobile",
    sum = "h1:h+GZ3ubjuWaQjGe8owMGcmMVCqs0xYJtRG5y2bpHaqU=",
    version = "v0.0.0-20210220033013-bdb1ca9a1e08",
)

go_repository(
    name = "org_golang_x_mod",
    importpath = "golang.org/x/mod",
    sum = "h1:Gz96sIWK3OalVv/I/qNygP42zyoKp3xptRVCWRFEBvo=",
    version = "v0.4.2",
)

go_repository(
    name = "org_golang_x_net",
    importpath = "golang.org/x/net",
    sum = "h1:4nGaVu0QrbjT/AK2PRLuQfQuh6DJve+pELhqTdAj3x0=",
    version = "v0.0.0-20210405180319-a5a99cb37ef4",
)

go_repository(
    name = "org_golang_x_oauth2",
    importpath = "golang.org/x/oauth2",
    sum = "h1:TzXSXBo42m9gQenoE3b9BGiEpg5IG2JkU5FkPIawgtw=",
    version = "v0.0.0-20200107190931-bf48bf16ab8d",
)

go_repository(
    name = "org_golang_x_sync",
    importpath = "golang.org/x/sync",
    sum = "h1:5KslGYwFpkhGh+Q16bwMP3cOontH8FOep7tGV86Y7SQ=",
    version = "v0.0.0-20210220032951-036812b2e83c",
)

go_repository(
    name = "org_golang_x_sys",
    importpath = "golang.org/x/sys",
    sum = "h1:gG67DSER+11cZvqIMb8S8bt0vZtiN6xWYARwirrOSfE=",
    version = "v0.0.0-20210510120138-977fb7262007",
)

go_repository(
    name = "org_golang_x_term",
    importpath = "golang.org/x/term",
    sum = "h1:v+OssWQX+hTHEmOBgwxdZxK4zHq3yOs8F9J7mk0PY8E=",
    version = "v0.0.0-20201126162022-7de9c90e9dd1",
)

go_repository(
    name = "org_golang_x_text",
    importpath = "golang.org/x/text",
    sum = "h1:aRYxNxv6iGQlyVaZmk6ZgYEDa+Jg18DxebPSrd6bg1M=",
    version = "v0.3.6",
)

go_repository(
    name = "org_golang_x_tools",
    importpath = "golang.org/x/tools",
    sum = "h1:wGiQel/hW0NnEkJUk8lbzkX2gFJU6PFxf1v5OlCfuOs=",
    version = "v0.1.1",
)

go_repository(
    name = "org_golang_x_xerrors",
    importpath = "golang.org/x/xerrors",
    sum = "h1:go1bK/D/BFZV2I8cIQd1NKEZ+0owSTG1fDTci4IqFcE=",
    version = "v0.0.0-20200804184101-5ec99f83aff1",
)

new_git_repository(
    name = "glfw",
    remote = "https://github.com/glfw/glfw.git",
    commit = "8d7e5cdb49a1a5247df612157ecffdd8e68923d2",
    build_file = "@//:third-party/glfw/glfw.BUILD",
)

go_rules_dependencies()

go_register_toolchains(version = "1.17.2")

gazelle_dependencies()

rules_proto_dependencies()

rules_proto_toolchains()
