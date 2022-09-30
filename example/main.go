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
	fmt.Println(svc.GetStringEx("test.hello", nil, nil))
	testData(svc, "test.book", 1)
	testData(svc, "test.book", 2)
	testData(svc, "test.book", 3)
	testData(svc, "test.pen", 1)
	testData(svc, "test.pen", 2)
	testData(svc, "test.pen", 3)
	err := svc.Localize("en-US")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("=========================== Localize en-US ===========================")
	fmt.Println(svc.GetString("test.hello"))
	fmt.Println(svc.GetStringEx("test.hello", nil, nil))
	testData(svc, "test.book", 1)
	testData(svc, "test.book", 2)
	testData(svc, "test.book", 3)
	testData(svc, "test.pen", 1)
	testData(svc, "test.pen", 2)
	testData(svc, "test.pen", 3)
}

func testData(svc i18n.I18n, id string, total int) {
	fmt.Println("GetString:\t\t\t\t", svc.GetString(id, i18n.PluralCount, total, "total", total))
	fmt.Println("GetString with kv:\t\t\t", svc.GetString(id, i18n.KeyValue().Plural(total).Add("total", total)...))
	fmt.Println("GetStringEx:\t\t\t\t", svc.GetStringEx(id, map[string]int{"total": total}, total))
}
