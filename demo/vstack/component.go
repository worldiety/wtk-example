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

package vstack

import (
	. "github.com/golangee/forms"
)

const Path = "/demo/vstack"

type ContentView struct {
	*VStack
}

func NewContentView() *ContentView {
	view := &ContentView{}
	view.VStack = NewVStack().AddViews(
		NewText("VStack").Style(Font(Headline1)),
		NewText("VStack or vertical stack is a single column grid.").Style(Font(Body)),

		NewVStack(
			NewText("a").Style(BackgroundColor(Red50)),
			NewText("b").Style(BackgroundColor(Blue50)),
			NewText("c").Style(BackgroundColor(Yellow50)),
			NewText("d").Style(BackgroundColor(Green50)),
		).Style(BackgroundColor(BlueGray50), Padding()).SetGap(Pixel(DefaultPadding)),
		NewCode(GoSyntax, code),
	)
	return view
}

func FromQuery(Query) View {
	return NewContentView()
}

const code = `package vstack

import (
	. "github.com/golangee/forms"
)

const Path = "/demo/vstack"

type ContentView struct {
	*VStack
}

func NewContentView() *ContentView {
	view := &ContentView{}
	view.VStack = NewVStack().AddViews(
		NewText("VStack").Style(Font(Headline1)),
		NewText("VStack is a single row grid.").Style(Font(Body)),

		NewVStack(
			NewText("a").Style(BackgroundColor(Red50)),
			NewText("b").Style(BackgroundColor(Blue50)),
			NewText("c").Style(BackgroundColor(Yellow50)),
			NewText("d").Style(BackgroundColor(Green50)),
		).Style(BackgroundColor(BlueGray50), Padding()).SetGap(Pixel(DefaultPadding)),
		NewCode(GoSyntax, code),
	)
	return view
}`
