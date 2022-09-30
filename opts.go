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

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v2"
	"io"
	"io/fs"
	"io/ioutil"
	"path/filepath"
)

type options struct{}

var Options options

// 配置默认语言
func (o options) DefaultLanguage(lang language.Tag) opt {
	return func(n *defaultI18n) {
		n.bundle = i18n.NewBundle(lang)
		n.lang = lang.String()
		n.localizer = i18n.NewLocalizer(n.bundle, n.lang)
	}
}

// 注册反序列化函数
func (o options) RegisterUnmarshalFunc(format string, unmarshalFunc func(data []byte, v interface{}) error) opt {
	return func(n *defaultI18n) {
		n.dataFuncs = append(n.dataFuncs, func(bundle *i18n.Bundle) {
			bundle.RegisterUnmarshalFunc(format, unmarshalFunc)
		})
	}
}

// 注册支持yaml格式
func (o options) SupportYaml() opt {
	return func(n *defaultI18n) {
		n.dataFuncs = append(n.dataFuncs, func(bundle *i18n.Bundle) {
			bundle.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)
		})
	}
}

// 注册支持toml格式
func (o options) SupportToml() opt {
	return func(n *defaultI18n) {
		n.dataFuncs = append(n.dataFuncs, func(bundle *i18n.Bundle) {
			bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
		})
	}
}

// 从文件加载message数据
func (o options) LoadMessageFile(filepath string) opt {
	return func(n *defaultI18n) {
		n.dataFuncs = append(n.dataFuncs, func(bundle *i18n.Bundle) {
			bundle.MustLoadMessageFile(filepath)
		})
	}
}

// 从目录加载message数据
func (o options) LoadMessageDir(dir string, filter func(path string) bool) opt {
	return func(n *defaultI18n) {
		n.dataFuncs = append(n.dataFuncs, func(bundle *i18n.Bundle) {
			_ = filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if !info.IsDir() {
					if filter != nil && !filter(path) {
						return nil
					}
					_, err = bundle.LoadMessageFile(path)
					if err != nil {
						return err
					}
				}
				return nil
			})
		})
	}
}

// 从data加载message数据
func (o options) LoadMessageData(data []byte, lang, format string) opt {
	return func(n *defaultI18n) {
		n.dataFuncs = append(n.dataFuncs, func(bundle *i18n.Bundle) {
			bundle.MustParseMessageFileBytes(data, fmt.Sprintf("%s.%s", lang, format))
		})
	}
}

// 从reader加载message数据
func (o options) LoadMessageReader(r io.Reader, lang, format string) opt {
	return func(n *defaultI18n) {
		n.dataFuncs = append(n.dataFuncs, func(bundle *i18n.Bundle) {
			d, err := ioutil.ReadAll(r)
			if err != nil {
				panic(err)
			}
			bundle.MustParseMessageFileBytes(d, fmt.Sprintf("%s.%s", lang, format))
		})
	}
}
