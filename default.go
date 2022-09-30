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
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"sync"
)

var (
	defaultLanguage = language.Chinese
)

type defaultI18n struct {
	bundle    *i18n.Bundle
	localizer *i18n.Localizer

	lang string
	lock sync.RWMutex

	dataFuncs []func(bundle *i18n.Bundle)
}

type opt func(*defaultI18n)

func New(opts ...opt) *defaultI18n {
	ret := &defaultI18n{
	}

	for _, opt := range opts {
		opt(ret)
	}

	if ret.bundle == nil {
		ret.bundle = i18n.NewBundle(defaultLanguage)
		ret.lang = defaultLanguage.String()
		ret.localizer = i18n.NewLocalizer(ret.bundle, ret.lang)
	}

	for _, f := range ret.dataFuncs {
		if f != nil {
			f(ret.bundle)
		}
	}

	return ret
}

// Switch language
func (s *defaultI18n) Localize(lang string) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.lang != lang {
		s.lang = lang
		s.localizer = i18n.NewLocalizer(s.bundle, lang)
	}
	return nil
}

// Get i18n string
func (s *defaultI18n) GetString(id string, kvs ...interface{}) (message string) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	size := len(kvs)
	conf := &i18n.LocalizeConfig{
		MessageID: id,
	}
	if size > 0 {
		if size&1 != 0 {
			panic("KeyValue pair is odd ")
		}
		tmplData := make(map[string]interface{}, size>>1)
		var key string
		for i := range kvs {
			if i&1 == 0 {
				key = kvs[i].(string)
			} else {
				if key == PluralCount {
					conf.PluralCount = kvs[i]
				}
				tmplData[key] = kvs[i]
			}
		}
		conf.TemplateData = tmplData
	}
	str, err := s.localizer.Localize(conf)
	if err != nil {
		return id
	}

	return str
}
