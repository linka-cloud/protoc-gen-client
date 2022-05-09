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

.PHONY: install
install:
	@go install .

PROTO_IMPORTS = -I. \
	-I $(shell go list -m -f {{.Dir}} google.golang.org/protobuf)

.PHONY: gen-tests
gen-tests:
	@protoc $(PROTO_IMPORTS) --debug_out="tests:." tests/pb/test.proto

PROTO_OPTS = paths=source_relative

.PHONY: gen-example
gen-example: install
	@protoc  $(PROTO_IMPORTS) --go_out=$(PROTO_OPTS):. --go-grpc_out=$(PROTO_OPTS):. --client_out=$(PROTO_OPTS):. tests/pb/test.proto
