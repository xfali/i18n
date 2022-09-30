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

package i18n

const(
	PluralCount = "i18n.PluralCount"
)

type I18n interface {
	// 指定使用语言
	// lang: 语言 zh | en | jp ...
	// err: 切换语言错误，如不支持该语言等
	Localize(lang string) (err error)

	// 获得i18n字符串
	// id: i18n message ID
	// kvs: 增加参数，格式为key value对，key为string，value为interface
	// 注意：
	// 	1、必须为偶数；
	// 	2、同时参数如果key为i18n.PluralCount，则value作为判定plural选择i18n message的依据
	// msg: 返回i18n字符串
	GetString(id string, kvs ...interface{}) (message string)
}
