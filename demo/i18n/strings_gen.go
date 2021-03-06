// Code generated by go generate; DO NOT EDIT.
// This file was generated by github.com/golangee/i18n

package i18n

import (
	"fmt"
	i18n "github.com/golangee/i18n"
)

func init() {
	var tag string

	// from strings-de-DE.xml
	tag = "de-DE"

	i18n.ImportValue(i18n.NewText(tag, "demo_i18n_active_locale", "Aktive Sprace: %s"))
	i18n.ImportValue(i18n.NewText(tag, "demo_i18n_browser_locale", "browser Sprache: %s"))
	i18n.ImportValue(i18n.NewText(tag, "demo_i18n_caption", "Lokalisierung"))
	i18n.ImportValue(i18n.NewText(tag, "demo_i18n_select_language", "Sprache wählen"))
	i18n.ImportValue(i18n.NewText(tag, "demo_i18n_text", "Dies ist kein Kernmerkmal der Formulare selbst, sondern aus dem Modul ee/i18n. Die Verwendung des Einbettungscodegenerators macht alles typsicher, erfordert aber Platz in Wasm-Binär und Speicher. Es wird empfohlen, nur eine Basissprache zu verwenden und zur Laufzeit eine einheitliche Übersetzungsdatei zu laden."))
	i18n.ImportValue(i18n.NewText(tag, "hello_world", "Hallo Welt"))
	i18n.ImportValue(i18n.NewText(tag, "hello_x", "Hallo %s"))
	i18n.ImportValue(i18n.NewText(tag, "locale", "de-DE"))
	i18n.ImportValue(i18n.NewTextArray(tag, "selector_details_array", "Erste Zeile", "Zweite Zeile", "Dritte Zeile", "Vierte"))
	i18n.ImportValue(i18n.NewQuantityText(tag, "x_cats").One("Eine Katze").Other("%[1]d Katzen"))
	i18n.ImportValue(i18n.NewText(tag, "x_runs_around_Y_and_sings_z", " %[3]s wird gesungen von %[1]s, der um %[2]s herumläuft"))
	_ = tag

	// from strings-en.xml
	tag = "en"

	i18n.ImportValue(i18n.NewText(tag, "demo_i18n_active_locale", "active locale: %s"))
	i18n.ImportValue(i18n.NewText(tag, "demo_i18n_browser_locale", "browser locale: %s"))
	i18n.ImportValue(i18n.NewText(tag, "demo_i18n_caption", "Localisation"))
	i18n.ImportValue(i18n.NewText(tag, "demo_i18n_select_language", "select language"))
	i18n.ImportValue(i18n.NewText(tag, "demo_i18n_text", "This is not a core feature by forms itself, but from the ee/i18n module. Using the embedding code generator, makes everything typesafe but requires space in wasm binary and memory. The recommendation is to just use a base language and load a unified translation file at runtime."))
	i18n.ImportValue(i18n.NewText(tag, "hello_world", "hello world"))
	i18n.ImportValue(i18n.NewText(tag, "hello_x", "Hello %s"))
	i18n.ImportValue(i18n.NewText(tag, "locale", "en"))
	i18n.ImportValue(i18n.NewTextArray(tag, "selector_details_array", "first line", "second line", "third line", "fourth line"))
	i18n.ImportValue(i18n.NewQuantityText(tag, "x_cats").One("One cat").Other("%[1]d cats"))
	i18n.ImportValue(i18n.NewText(tag, "x_runs_around_Y_and_sings_z", "%[1]s runs around the %[2]s and sings %[3]s"))
	_ = tag

}

// Resources wraps the package strings to get invoked safely.
type Resources struct {
	res *i18n.Resources
}

// NewResources creates a new localized resource instance.
func NewResources(locale string) Resources {
	return Resources{i18n.From(locale)}
}

// DemoI18NActiveLocale returns a translated text for "active locale: %s"
func (r Resources) DemoI18NActiveLocale(str0 string) string {
	str, err := r.res.Text("demo_i18n_active_locale", str0)
	if err != nil {
		return fmt.Errorf("MISS!demo_i18n_active_locale: %w", err).Error()
	}
	return str
}

// DemoI18NBrowserLocale returns a translated text for "browser locale: %s"
func (r Resources) DemoI18NBrowserLocale(str0 string) string {
	str, err := r.res.Text("demo_i18n_browser_locale", str0)
	if err != nil {
		return fmt.Errorf("MISS!demo_i18n_browser_locale: %w", err).Error()
	}
	return str
}

// DemoI18NCaption returns a translated text for "Localisation"
func (r Resources) DemoI18NCaption() string {
	str, err := r.res.Text("demo_i18n_caption")
	if err != nil {
		return fmt.Errorf("MISS!demo_i18n_caption: %w", err).Error()
	}
	return str
}

// DemoI18NSelectLanguage returns a translated text for "select language"
func (r Resources) DemoI18NSelectLanguage() string {
	str, err := r.res.Text("demo_i18n_select_language")
	if err != nil {
		return fmt.Errorf("MISS!demo_i18n_select_language: %w", err).Error()
	}
	return str
}

// DemoI18NText returns a translated text for "This is not a core feature by forms itself, but from the ee/i18n module. Using the embedding code generator, makes everything typesafe but requires space in wasm binary and memory. The recommendation is to just use a base language and load a unified translation file at runtime."
func (r Resources) DemoI18NText() string {
	str, err := r.res.Text("demo_i18n_text")
	if err != nil {
		return fmt.Errorf("MISS!demo_i18n_text: %w", err).Error()
	}
	return str
}

// HelloWorld returns a translated text for "hello world"
func (r Resources) HelloWorld() string {
	str, err := r.res.Text("hello_world")
	if err != nil {
		return fmt.Errorf("MISS!hello_world: %w", err).Error()
	}
	return str
}

// HelloX returns a translated text for "Hello %s"
func (r Resources) HelloX(str0 string) string {
	str, err := r.res.Text("hello_x", str0)
	if err != nil {
		return fmt.Errorf("MISS!hello_x: %w", err).Error()
	}
	return str
}

// Locale returns a translated text for "en"
func (r Resources) Locale() string {
	str, err := r.res.Text("locale")
	if err != nil {
		return fmt.Errorf("MISS!locale: %w", err).Error()
	}
	return str
}

// SelectorDetailsArray returns a translated text for "first line"
func (r Resources) SelectorDetailsArray() []string {
	str, err := r.res.TextArray("selector_details_array")
	if err != nil {
		return []string{fmt.Errorf("MISS!selector_details_array: %w", err).Error()}
	}
	return str
}

// XCats returns a translated text for "%[1]d cats"
func (r Resources) XCats(quantity int, num0 int) string {
	str, err := r.res.QuantityText("x_cats", quantity, num0)
	if err != nil {
		return fmt.Errorf("MISS!x_cats: %w", err).Error()
	}
	return str
}

// XRunsAroundYAndSingsZ returns a translated text for "%[1]s runs around the %[2]s and sings %[3]s"
func (r Resources) XRunsAroundYAndSingsZ(str0 string, str1 string, str2 string) string {
	str, err := r.res.Text("x_runs_around_Y_and_sings_z", str0, str1, str2)
	if err != nil {
		return fmt.Errorf("MISS!x_runs_around_Y_and_sings_z: %w", err).Error()
	}
	return str
}

// FuncMap returns the named functions to be used with a template
func (r Resources) FuncMap() map[string]interface{} {
	m := make(map[string]interface{})
	m["DemoI18NActiveLocale"] = r.DemoI18NActiveLocale
	m["DemoI18NBrowserLocale"] = r.DemoI18NBrowserLocale
	m["DemoI18NCaption"] = r.DemoI18NCaption
	m["DemoI18NSelectLanguage"] = r.DemoI18NSelectLanguage
	m["DemoI18NText"] = r.DemoI18NText
	m["HelloWorld"] = r.HelloWorld
	m["HelloX"] = r.HelloX
	m["Locale"] = r.Locale
	m["SelectorDetailsArray"] = r.SelectorDetailsArray
	m["XCats"] = r.XCats
	m["XRunsAroundYAndSingsZ"] = r.XRunsAroundYAndSingsZ
	return m
}
