package app

import (
	"github.com/golangee/forms-example/www/component/text"
	. "github.com/golangee/forms-example/www/component/view"
	"github.com/golangee/forms-example/www/dom"
	"github.com/golangee/forms-example/www/internal/build"
	"github.com/golangee/forms-example/www/style"
	"github.com/golangee/log"
	"github.com/golangee/log/ecs"
)

type Application struct {
	logger log.Logger
}

func NewApplication() *Application {
	return &Application{
		logger: log.NewLogger(ecs.Log("application")),
	}
}

func (a *Application) Run() {

	a.doStuff()
	//a.doStuffWithComponents()

	// keep alive
	select {}
}

type MyCustomView struct {
	View
	fu string
}

func (m *MyCustomView) Render() Node {
	return Div(
		Class(style.Text3xl, style.BgBlack, style.TextBlue600),
		text.NewText(m.fu).Render(),
		text.NewText("b").Render(),
		AddEventListener("click", func() {
			m.Invalidate()
		}),
	)
}

func (a *Application) doStuffWithComponents() {
	defer dom.GlobalPanicHandler()

	fu := "xai"
	myText := NewInlineView(func() Node {
		return Div(
			Class(style.Text3xl, style.BgBlack, style.TextBlue600),
			Class(style.Text3xl, style.BgBlack, style.TextBlue600),
			DebugLog("compose: doStuff"),
			AddEventListenerOnce(dom.EventRelease, func() {
				a.logger.Print(ecs.Msg("released"))
			}),
			Div(Text(fu)),
			Span(text.NewText("bbbb"), Class(style.BgRed500)),
		)
	})

	RenderBody(myText)
}

func (a *Application) doStuff() {
	defer dom.GlobalPanicHandler()

	a.logger.Print(ecs.Msg("application is running30"), log.V("build.commit", build.Commit))

	counter := 0
	var myImg dom.Element
	var nameText dom.Element
	content :=
		Div(Class("rounded overflow-hidden shadow-lg dark:bg-gray-800"),
			Figure(Class("md:flex bg-gray-100 rounded-xl p-8 md:p-0"),
				Img(
					Class("w-32 h-32 md:w-48 md:h-auto md:rounded-none rounded-full mx-auto"),
					Src("https://tailwindcss.com/_next/static/media/sarah-dayan.a8ff3f1095a58085a82e3bb6aab12eb2.jpg"),
					Width("384"),
					Height("512"),
					//	Self(&myImg),
					AddClickListener(func() {

						a.logger.Print(ecs.Msg("clicked it"))
						myImg.SetClassName("rounded-xl")
					}),
				),
				Div(
					Class("pt-6 md:p-8 text-center md:text-left space-y-4"),
					Blockquote(
						P(
							Class("text-lg font-semibold"),
							Text("“Tailwind CSS is the only framework that I've seen scale on large teams. It’s easy to customize, adapts to any design,and the build size is tiny.”"),
						),
					),
					Figcaption(Class("font-medium"),
						Div(Class("text-yellow-400"),
							Text("Sarah Dayan"),
							//Self(&nameText),
							AddEventListenerOnce("click", func() {
								counter++
								nameText.SetTextContent("absolute nice")
								a.logger.Print(ecs.Msg("only once clicked it"))
								myImg.SetClassName("rounded-xl")
								var failTest *int
								*failTest = 5
								_ = failTest
							}),
						),
						Div(Class("text-gray-500"),
							Text(" Staff Engineer, Algolia"),
						),
						Ul(
							Li(Text("hey")),
							Li(Text("ho")),
						),
					),
				),

			),
		)

	RenderBody(content)

	//myImg.Release()
}
