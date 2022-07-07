/*
Copyright © 2022 SUSE LLC

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

package config

import (
	"strings"

	"github.com/rancher/wrangler/pkg/data"
	"github.com/rancher/wrangler/pkg/data/convert"
	schemas2 "github.com/rancher/wrangler/pkg/schemas"
	"github.com/rancher/wrangler/pkg/schemas/mappers"
)

type FuzzyNames struct {
	mappers.DefaultMapper
	names map[string]string
}

func (f *FuzzyNames) ToInternal(data data.Object) error {
	for k, v := range data {
		if newK, ok := f.names[strings.ToLower(k)]; ok && newK != k {
			data[newK] = v
		}
	}
	return nil
}

func (f *FuzzyNames) addName(name, toName string) {
	f.names[strings.ToLower(name)] = toName
	f.names[convert.ToYAMLKey(name)] = toName
	f.names[strings.ToLower(convert.ToYAMLKey(name))] = toName
}

func (f *FuzzyNames) ModifySchema(schema *schemas2.Schema, schemas *schemas2.Schemas) error {
	if f.names == nil {
		f.names = map[string]string{}
	}

	for name := range schema.ResourceFields {
		if strings.HasSuffix(name, "s") && len(name) > 1 {
			f.addName(name[:len(name)-1], name)
		}
		if strings.HasSuffix(name, "es") && len(name) > 2 {
			f.addName(name[:len(name)-2], name)
		}
		f.addName(name, name)
	}

	f.names["pass"] = "passphrase"
	f.names["password"] = "passphrase"

	return nil
}
