# Copyright 2021 Linka Cloud  All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
# http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

$(shell mkdir -p .bin)

export GOBIN=$(PWD)/.bin

export PATH := $(GOBIN):$(PATH)

.PHONY: install
install:
	@go install .
	@go install github.com/alta/protopatch/cmd/protoc-gen-go-patch

PROTO_IMPORTS = -I. \
	-I $(shell go list -m -f {{.Dir}} google.golang.org/protobuf) \
	-I $(shell go list -m -f {{.Dir}} github.com/alta/protopatch)

.PHONY: gen-tests
gen-tests:
	@@protoc $(PROTO_IMPORTS) --go-patch_out=plugin=debug,"tests:." tests/pb/external/ext.proto
	@protoc $(PROTO_IMPORTS) --go-patch_out=plugin=debug,"tests:." tests/pb/test.proto

PROTO_OPTS = paths=source_relative

.PHONY: gen-example
gen-example: install
	@protoc $(PROTO_IMPORTS) --go-patch_out=plugin=go,$(PROTO_OPTS):. tests/pb/external/ext.proto
	@protoc $(PROTO_IMPORTS) --go-patch_out=plugin=go,$(PROTO_OPTS):. --go-patch_out=plugin=go-grpc,$(PROTO_OPTS):. --go-patch_out=plugin=client,$(PROTO_OPTS):. tests/pb/test.proto
