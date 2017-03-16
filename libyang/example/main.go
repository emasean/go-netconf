/**
 * @file example.go
 * @author Mislav Novakovic <mislav.novakovic@sartura.hr>
 * @brief libyang go example.
 *
 * @copyright
 * Copyright 2017 Deutsche Telekom AG.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"fmt"
)

/*
#cgo LDFLAGS: -lyang
#cgo LDFLAGS: -lpcre
#include <libyang/libyang.h>
*/
import "C"

func main() {
	// create libyang context
	ctx := C.ly_ctx_new(C.CString("./"))

	if ctx != nil {
		fmt.Printf("Libyang context successfully created.\n")
	} else {
		fmt.Printf("Failed to load libyang.\n")
	}
}
