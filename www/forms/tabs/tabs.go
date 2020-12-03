package tabs

import (
	"github.com/golangee/forms-example/www/forms/property"
	. "github.com/golangee/forms-example/www/forms/view"
)

type Tabs struct {
	View
	tabPanes []*tabPane
}

func NewTabs() *Tabs {
	return &Tabs{}
}

func (c *Tabs) AddPane(caption, content Renderable) *Tabs {
	c.tabPanes = append(c.tabPanes, &tabPane{
		caption: caption,
		content: content,
	})

	if len(c.tabPanes) == 1 {
		c.tabPanes[0].active.Set(true)
	}

	c.Invalidate()
	return c
}

func (c *Tabs) Render() Node {
	return Div(


		Ul(Class("list-reset flex"),
			ForEach(len(c.tabPanes), func(i int) Renderable {
				tab := c.tabPanes[i]
				return Li(Class("-mb-px"),

					A(Class("py-2 px-4 transform ease-in-out transition-colors duration-250 hover:bg-primary"),
						If(&tab.active, AddClass("border-b border-primary text-primary"), RemoveClass("border-b border-primary text-primary")),
						AddClickListener(func() {
							c.SetActive(i)
						}),
						tab.caption,
					),
				)
			}),
		),

		ForEach(len(c.tabPanes), func(i int) Renderable {
			tab := c.tabPanes[i]
			return Div(Class("pt-2 ease-in-out transition-opacity duration-500 opacity-0 absolute"),
				If(&tab.active,
					WithModifiers(
						Visibility("visible"),
						AddClass("opacity-100"),
					),
					WithModifiers(
						Visibility("hidden"),
						RemoveClass("opacity-100"),
					)),

				tab.content,

			)
		}),

	)
}

func (c *Tabs) Self(ref **Tabs) *Tabs {
	*ref = c
	return c
}

func (c *Tabs) With(f func(c *Tabs)) *Tabs {
	f(c)
	return c
}

func (c *Tabs) SetActive(idx int) *Tabs {
	for i, pane := range c.tabPanes {
		pane.active.Set(i == idx)
	}

	return c
}

type tabPane struct {
	caption, content Renderable
	active           property.Bool
}
