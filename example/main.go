/*
 * Copyright 2022 Xiongfa Li.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
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
	"github.com/xfali/i18n"
	"os"
	"path"
)

func main() {
	svc := i18n.New(i18n.Options.LoadMessageDir(".", func(file string) bool {
		return path.Ext(file) == ".json"
	}))
	fmt.Println(svc.GetString("test.hello"))
	err := svc.Localize("en-US")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(svc.GetString("test.hello"))
}
