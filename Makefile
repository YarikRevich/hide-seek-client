#DO NOT EDIT! USED FOR DEVELOPMENT

.PHONY: gen_proto

gen_proto:
	@bazel build //api:server_external_api
	@bazel build //api:services_external_api
	@cp bazel-bin/api/server_external_api_/github.com/YarikRevich/hide-seek-client/internal/core/networking/api/server_external/server_external.pb.go internal/core/networking/api/server_external
	@cp bazel-bin/api/services_external_api_/github.com/YarikRevich/hide-seek-client/internal/core/networking/api/services_external/services_external.pb.go internal/core/networking/api/services_external
