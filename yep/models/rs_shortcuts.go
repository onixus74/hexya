// Copyright 2016 NDP Systèmes. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package models

// Create is a shortcut function for rs.Call("Create") on the current RecordSet.
// Data can be either a struct pointer or a FieldMap.
func (rs RecordCollection) Create(data interface{}) *RecordCollection {
	return rs.Call("Create", data).(*RecordCollection)
}

// Write is a shortcut for rs.Call("Write") on the current RecordSet.
// Data can be either a struct pointer or a FieldMap.
func (rs RecordCollection) Write(data interface{}) bool {
	return rs.Call("Write", data).(bool)
}

// Unlink is a shortcut for rs.Call("Unlink") on the current RecordSet.
func (rs RecordCollection) Unlink() int64 {
	return rs.Call("Unlink").(int64)
}