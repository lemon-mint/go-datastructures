/*
Copyright 2015 Workiva, LLC

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

package mock

import (
	"github.com/stretchr/testify/mock"

	"github.com/lemon-mint/go-datastructures/batcher"
)

var _ batcher.Batcher = new(Batcher)

type Batcher struct {
	mock.Mock
	PutChan chan bool
}

func (m *Batcher) Put(items interface{}) error {
	args := m.Called(items)
	if m.PutChan != nil {
		m.PutChan <- true
	}
	return args.Error(0)
}

func (m *Batcher) Get() ([]interface{}, error) {
	args := m.Called()
	return args.Get(0).([]interface{}), args.Error(1)
}

func (m *Batcher) Flush() error {
	args := m.Called()
	return args.Error(0)
}

func (m *Batcher) Dispose() {
	m.Called()
}

func (m *Batcher) IsDisposed() bool {
	args := m.Called()
	return args.Bool(0)
}
