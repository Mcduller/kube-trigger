//go:build !ignore_autogenerated

/*
Copyright  The KubeVela Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by ../../../../../hack/generate-go-const-from-file.sh. DO NOT EDIT.

// Instead, edit properties.cue and regenerate this using go generate ./...

package config

const propertiesCUETemplate = `// This is a validator for properties of k8s-resource-watcher

#eventType: "update" | "create" | "delete"

apiVersion: string
kind:       string
namespace:  *"" | string
events: *[] | [...#eventType]
`
