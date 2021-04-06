/*
 * Copyright (c) 2020 Baidu, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package json

import (
	"io"

	"github.com/baidu/openless/pkg/util/json"
)

// Serializer xxx
type Serializer struct {
	pretty bool
}

// NewSerializer xxx
func NewSerializer() *Serializer {
	return &Serializer{
		pretty: true,
	}
}

// Encode xxx
func (s *Serializer) Encode(obj interface{}, w io.Writer) error {
	return nil
}

// Decode xxx
func (s *Serializer) Decode(data []byte, into interface{}) error {
	return json.Unmarshal(data, into)
}
