package bsd

import (
	"github.com/surechen/go-licenses-toolkie/licenses/common"
	"github.com/surechen/go-licenses-toolkie/licenses/handler"
	"strings"
)

type Bsd2Handler struct {
	Name string
	WordsMap map[string]int
}

var Handler Bsd2Handler

var WordsMap map[string]int

const bsd2Pattern = `Copyright (c) 2013 Richard Musiol. All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are
met:

   * Redistributions of source code must retain the above copyright
notice, this list of conditions and the following disclaimer.
   * Redistributions in binary form must reproduce the above
copyright notice, this list of conditions and the following disclaimer
in the documentation and/or other materials provided with the
distribution.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
"AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
`

func (h *Bsd2Handler) Parse(licenseInfo string) error {
	h.WordsMap = ParseLicense(licenseInfo)
	return nil
}

func (h *Bsd2Handler) Analyse(licenseInfo string) int {
	count := 0
	WordsMap := ParseLicense(licenseInfo)
	for word, value := range WordsMap {
		if info, ok := h.WordsMap[word]; ok {
			if value > info {
				count += value - info
			} else {
				count += info - value
			}
		} else {
			count += value
		}
	}
	return count
}

func ParseLicense(licenseInfo string) (map[string]int) {
	licenseInfo = strings.Replace(licenseInfo, "\n", " ", -1)
	licenseInfo = strings.Replace(licenseInfo, "\r\n", " ", -1)
	wordList := strings.Split(licenseInfo, " ")
	wordsMap := make(map[string]int, len(wordList))
	for _, word := range wordList {
		if count, ok := wordsMap[word]; ok {
			wordsMap[word] = count + 1
		} else {
			wordsMap[word] = 1
		}
	}
	return wordsMap
}

func init() {
	Handler.Name = common.Bsd2
	Handler.Parse(bsd2Pattern)
	handler.InstallHandler(Handler.Name, &Handler)
}