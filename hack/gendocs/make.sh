#!/usr/bin/env bash

# Copyright AppsCode Inc. and Contributors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

pushd $GOPATH/src/go.searchlight.dev/icinga-operator/hack/gendocs
go run main.go

cd $GOPATH/src/go.searchlight.dev/icinga-operator/docs/reference/hyperalert
sed -i 's/######\ Auto\ generated\ by.*//g' *

cd $GOPATH/src/go.searchlight.dev/icinga-operator/docs/reference/searchlight
sed -i 's/######\ Auto\ generated\ by.*//g' *

cd $GOPATH/src/go.searchlight.dev/icinga-operator/docs/reference/hostfacts
sed -i 's/######\ Auto\ generated\ by.*//g' *
popd
