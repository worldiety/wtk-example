// Copyright 2020 Torben Schinke
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package section3

import . "github.com/golangee/forms"

type ContentView struct {
	*VStack
}

func NewContentView() *ContentView {
	return &ContentView{
		VStack: NewVStack().AddViews(
			NewText("Turtle Rock").
				Style(Font(Headline1)),
			NewHStack().AddViews(
				NewText("Joshua Tree National Park").
					Style(Font(Headline2)),
				NewSpacer(),
				NewText("California").
					Style(Font(Headline2)),
			),

		),
	}
}
