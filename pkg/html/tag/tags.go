// Code generated - DO NOT EDIT.
package tag

import (
	"github.com/alexisbcz/vauban/pkg/html"
)


type a struct {
	*Tag
}

func A(children ...html.Node) *a {
	return &a{New("a", false, children)}
}

func (e *a) Style(value string) *a {
	e.Attribute("style", value)
	return e
}

func (e *a) StyleIf(condition bool, value string) *a {
	if condition {
		e.Attribute("style", value)
	}
	return e
}

func (e *a) ID(value string) *a {
	e.Attribute("id", value)
	return e
}

func (e *a) IDIf(condition bool, value string) *a {
	if condition {
		e.Attribute("id", value)
	}
	return e
}

func (e *a) Class(value string) *a {
	e.Attribute("class", value)
	return e
}

func (e *a) ClassIf(condition bool, value string) *a {
	if condition {
		e.Attribute("class", value)
	}
	return e
}

func (e *a) Href(value string) *a {
	e.Attribute("href", value)
	return e
}

func (e *a) HrefIf(condition bool, value string) *a {
	if condition {
		e.Attribute("href", value)
	}
	return e
}

func (e *a) Target(value string) *a {
	e.Attribute("target", value)
	return e
}

func (e *a) TargetIf(condition bool, value string) *a {
	if condition {
		e.Attribute("target", value)
	}
	return e
}

func (e *a) Download(value string) *a {
	e.Attribute("download", value)
	return e
}

func (e *a) DownloadIf(condition bool, value string) *a {
	if condition {
		e.Attribute("download", value)
	}
	return e
}

func (e *a) Rel(value string) *a {
	e.Attribute("rel", value)
	return e
}

func (e *a) RelIf(condition bool, value string) *a {
	if condition {
		e.Attribute("rel", value)
	}
	return e
}

func (e *a) Type(value string) *a {
	e.Attribute("type", value)
	return e
}

func (e *a) TypeIf(condition bool, value string) *a {
	if condition {
		e.Attribute("type", value)
	}
	return e
}

func (e *a) Hreflang(value string) *a {
	e.Attribute("hreflang", value)
	return e
}

func (e *a) HreflangIf(condition bool, value string) *a {
	if condition {
		e.Attribute("hreflang", value)
	}
	return e
}

func (e *a) Media(value string) *a {
	e.Attribute("media", value)
	return e
}

func (e *a) MediaIf(condition bool, value string) *a {
	if condition {
		e.Attribute("media", value)
	}
	return e
}

func (e *a) Ping(value string) *a {
	e.Attribute("ping", value)
	return e
}

func (e *a) PingIf(condition bool, value string) *a {
	if condition {
		e.Attribute("ping", value)
	}
	return e
}

func (e *a) Referrerpolicy(value string) *a {
	e.Attribute("referrerpolicy", value)
	return e
}

func (e *a) ReferrerpolicyIf(condition bool, value string) *a {
	if condition {
		e.Attribute("referrerpolicy", value)
	}
	return e
}

func (e *a) Accesskey(value string) *a {
	e.Attribute("accesskey", value)
	return e
}

func (e *a) AccesskeyIf(condition bool, value string) *a {
	if condition {
		e.Attribute("accesskey", value)
	}
	return e
}

func (e *a) Contenteditable(value string) *a {
	e.Attribute("contenteditable", value)
	return e
}

func (e *a) ContenteditableIf(condition bool, value string) *a {
	if condition {
		e.Attribute("contenteditable", value)
	}
	return e
}

func (e *a) Dir(value string) *a {
	e.Attribute("dir", value)
	return e
}

func (e *a) DirIf(condition bool, value string) *a {
	if condition {
		e.Attribute("dir", value)
	}
	return e
}

func (e *a) Draggable(value string) *a {
	e.Attribute("draggable", value)
	return e
}

func (e *a) DraggableIf(condition bool, value string) *a {
	if condition {
		e.Attribute("draggable", value)
	}
	return e
}

func (e *a) Hidden(value string) *a {
	e.Attribute("hidden", value)
	return e
}

func (e *a) HiddenIf(condition bool, value string) *a {
	if condition {
		e.Attribute("hidden", value)
	}
	return e
}

func (e *a) Lang(value string) *a {
	e.Attribute("lang", value)
	return e
}

func (e *a) LangIf(condition bool, value string) *a {
	if condition {
		e.Attribute("lang", value)
	}
	return e
}

func (e *a) Spellcheck(value string) *a {
	e.Attribute("spellcheck", value)
	return e
}

func (e *a) SpellcheckIf(condition bool, value string) *a {
	if condition {
		e.Attribute("spellcheck", value)
	}
	return e
}

func (e *a) Tabindex(value string) *a {
	e.Attribute("tabindex", value)
	return e
}

func (e *a) TabindexIf(condition bool, value string) *a {
	if condition {
		e.Attribute("tabindex", value)
	}
	return e
}

func (e *a) Title(value string) *a {
	e.Attribute("title", value)
	return e
}

func (e *a) TitleIf(condition bool, value string) *a {
	if condition {
		e.Attribute("title", value)
	}
	return e
}

func (e *a) Translate(value string) *a {
	e.Attribute("translate", value)
	return e
}

func (e *a) TranslateIf(condition bool, value string) *a {
	if condition {
		e.Attribute("translate", value)
	}
	return e
}

func (e *a) Role(value string) *a {
	e.Attribute("role", value)
	return e
}

func (e *a) RoleIf(condition bool, value string) *a {
	if condition {
		e.Attribute("role", value)
	}
	return e
}

type abbr struct {
	*Tag
}

func Abbr(children ...html.Node) *abbr {
	return &abbr{New("abbr", false, children)}
}

type address struct {
	*Tag
}

func Address(children ...html.Node) *address {
	return &address{New("address", false, children)}
}

type area struct {
	*Tag
}

func Area(children ...html.Node) *area {
	return &area{New("area", true, children)}
}

func (e *area) Alt(value string) *area {
	e.Attribute("alt", value)
	return e
}

func (e *area) AltIf(condition bool, value string) *area {
	if condition {
		e.Attribute("alt", value)
	}
	return e
}

func (e *area) Coords(value string) *area {
	e.Attribute("coords", value)
	return e
}

func (e *area) CoordsIf(condition bool, value string) *area {
	if condition {
		e.Attribute("coords", value)
	}
	return e
}

func (e *area) Shape(value string) *area {
	e.Attribute("shape", value)
	return e
}

func (e *area) ShapeIf(condition bool, value string) *area {
	if condition {
		e.Attribute("shape", value)
	}
	return e
}

type article struct {
	*Tag
}

func Article(children ...html.Node) *article {
	return &article{New("article", false, children)}
}

type aside struct {
	*Tag
}

func Aside(children ...html.Node) *aside {
	return &aside{New("aside", false, children)}
}

type audio struct {
	*Tag
}

func Audio(children ...html.Node) *audio {
	return &audio{New("audio", false, children)}
}

func (e *audio) Src(value string) *audio {
	e.Attribute("src", value)
	return e
}

func (e *audio) SrcIf(condition bool, value string) *audio {
	if condition {
		e.Attribute("src", value)
	}
	return e
}

func (e *audio) Controls(value string) *audio {
	e.Attribute("controls", value)
	return e
}

func (e *audio) ControlsIf(condition bool, value string) *audio {
	if condition {
		e.Attribute("controls", value)
	}
	return e
}

func (e *audio) Autoplay(value string) *audio {
	e.Attribute("autoplay", value)
	return e
}

func (e *audio) AutoplayIf(condition bool, value string) *audio {
	if condition {
		e.Attribute("autoplay", value)
	}
	return e
}

func (e *audio) Loop(value string) *audio {
	e.Attribute("loop", value)
	return e
}

func (e *audio) LoopIf(condition bool, value string) *audio {
	if condition {
		e.Attribute("loop", value)
	}
	return e
}

func (e *audio) Muted(value string) *audio {
	e.Attribute("muted", value)
	return e
}

func (e *audio) MutedIf(condition bool, value string) *audio {
	if condition {
		e.Attribute("muted", value)
	}
	return e
}

func (e *audio) Preload(value string) *audio {
	e.Attribute("preload", value)
	return e
}

func (e *audio) PreloadIf(condition bool, value string) *audio {
	if condition {
		e.Attribute("preload", value)
	}
	return e
}

type b struct {
	*Tag
}

func B(children ...html.Node) *b {
	return &b{New("b", false, children)}
}

type base struct {
	*Tag
}

func Base(children ...html.Node) *base {
	return &base{New("base", true, children)}
}

type bdi struct {
	*Tag
}

func Bdi(children ...html.Node) *bdi {
	return &bdi{New("bdi", false, children)}
}

type bdo struct {
	*Tag
}

func Bdo(children ...html.Node) *bdo {
	return &bdo{New("bdo", false, children)}
}

type blockquote struct {
	*Tag
}

func Blockquote(children ...html.Node) *blockquote {
	return &blockquote{New("blockquote", false, children)}
}

func (e *blockquote) Cite(value string) *blockquote {
	e.Attribute("cite", value)
	return e
}

func (e *blockquote) CiteIf(condition bool, value string) *blockquote {
	if condition {
		e.Attribute("cite", value)
	}
	return e
}

type body struct {
	*Tag
}

func Body(children ...html.Node) *body {
	return &body{New("body", false, children)}
}

type br struct {
	*Tag
}

func Br(children ...html.Node) *br {
	return &br{New("br", true, children)}
}

type button struct {
	*Tag
}

func Button(children ...html.Node) *button {
	return &button{New("button", false, children)}
}

func (e *button) Value(value string) *button {
	e.Attribute("value", value)
	return e
}

func (e *button) ValueIf(condition bool, value string) *button {
	if condition {
		e.Attribute("value", value)
	}
	return e
}

func (e *button) Name(value string) *button {
	e.Attribute("name", value)
	return e
}

func (e *button) NameIf(condition bool, value string) *button {
	if condition {
		e.Attribute("name", value)
	}
	return e
}

func (e *button) Disabled(value string) *button {
	e.Attribute("disabled", value)
	return e
}

func (e *button) DisabledIf(condition bool, value string) *button {
	if condition {
		e.Attribute("disabled", value)
	}
	return e
}

func (e *button) Form(value string) *button {
	e.Attribute("form", value)
	return e
}

func (e *button) FormIf(condition bool, value string) *button {
	if condition {
		e.Attribute("form", value)
	}
	return e
}

func (e *button) Formaction(value string) *button {
	e.Attribute("formaction", value)
	return e
}

func (e *button) FormactionIf(condition bool, value string) *button {
	if condition {
		e.Attribute("formaction", value)
	}
	return e
}

func (e *button) Formmethod(value string) *button {
	e.Attribute("formmethod", value)
	return e
}

func (e *button) FormmethodIf(condition bool, value string) *button {
	if condition {
		e.Attribute("formmethod", value)
	}
	return e
}

func (e *button) Formenctype(value string) *button {
	e.Attribute("formenctype", value)
	return e
}

func (e *button) FormenctypeIf(condition bool, value string) *button {
	if condition {
		e.Attribute("formenctype", value)
	}
	return e
}

func (e *button) Formtarget(value string) *button {
	e.Attribute("formtarget", value)
	return e
}

func (e *button) FormtargetIf(condition bool, value string) *button {
	if condition {
		e.Attribute("formtarget", value)
	}
	return e
}

func (e *button) Formnovalidate(value string) *button {
	e.Attribute("formnovalidate", value)
	return e
}

func (e *button) FormnovalidateIf(condition bool, value string) *button {
	if condition {
		e.Attribute("formnovalidate", value)
	}
	return e
}

func (e *button) Autofocus(value string) *button {
	e.Attribute("autofocus", value)
	return e
}

func (e *button) AutofocusIf(condition bool, value string) *button {
	if condition {
		e.Attribute("autofocus", value)
	}
	return e
}

type canvas struct {
	*Tag
}

func Canvas(children ...html.Node) *canvas {
	return &canvas{New("canvas", false, children)}
}

func (e *canvas) Width(value string) *canvas {
	e.Attribute("width", value)
	return e
}

func (e *canvas) WidthIf(condition bool, value string) *canvas {
	if condition {
		e.Attribute("width", value)
	}
	return e
}

func (e *canvas) Height(value string) *canvas {
	e.Attribute("height", value)
	return e
}

func (e *canvas) HeightIf(condition bool, value string) *canvas {
	if condition {
		e.Attribute("height", value)
	}
	return e
}

type caption struct {
	*Tag
}

func Caption(children ...html.Node) *caption {
	return &caption{New("caption", false, children)}
}

type cite struct {
	*Tag
}

func Cite(children ...html.Node) *cite {
	return &cite{New("cite", false, children)}
}

type code struct {
	*Tag
}

func Code(children ...html.Node) *code {
	return &code{New("code", false, children)}
}

type col struct {
	*Tag
}

func Col(children ...html.Node) *col {
	return &col{New("col", true, children)}
}

func (e *col) Span(value string) *col {
	e.Attribute("span", value)
	return e
}

func (e *col) SpanIf(condition bool, value string) *col {
	if condition {
		e.Attribute("span", value)
	}
	return e
}

type colgroup struct {
	*Tag
}

func Colgroup(children ...html.Node) *colgroup {
	return &colgroup{New("colgroup", false, children)}
}

type data struct {
	*Tag
}

func Data(children ...html.Node) *data {
	return &data{New("data", false, children)}
}

type datalist struct {
	*Tag
}

func Datalist(children ...html.Node) *datalist {
	return &datalist{New("datalist", false, children)}
}

type dd struct {
	*Tag
}

func Dd(children ...html.Node) *dd {
	return &dd{New("dd", false, children)}
}

type del struct {
	*Tag
}

func Del(children ...html.Node) *del {
	return &del{New("del", false, children)}
}

type details struct {
	*Tag
}

func Details(children ...html.Node) *details {
	return &details{New("details", false, children)}
}

func (e *details) Open(value string) *details {
	e.Attribute("open", value)
	return e
}

func (e *details) OpenIf(condition bool, value string) *details {
	if condition {
		e.Attribute("open", value)
	}
	return e
}

type dfn struct {
	*Tag
}

func Dfn(children ...html.Node) *dfn {
	return &dfn{New("dfn", false, children)}
}

type dialog struct {
	*Tag
}

func Dialog(children ...html.Node) *dialog {
	return &dialog{New("dialog", false, children)}
}

type div struct {
	*Tag
}

func Div(children ...html.Node) *div {
	return &div{New("div", false, children)}
}

type dl struct {
	*Tag
}

func Dl(children ...html.Node) *dl {
	return &dl{New("dl", false, children)}
}

type dt struct {
	*Tag
}

func Dt(children ...html.Node) *dt {
	return &dt{New("dt", false, children)}
}

type em struct {
	*Tag
}

func Em(children ...html.Node) *em {
	return &em{New("em", false, children)}
}

type embed struct {
	*Tag
}

func Embed(children ...html.Node) *embed {
	return &embed{New("embed", true, children)}
}

type fieldset struct {
	*Tag
}

func Fieldset(children ...html.Node) *fieldset {
	return &fieldset{New("fieldset", false, children)}
}

type figcaption struct {
	*Tag
}

func Figcaption(children ...html.Node) *figcaption {
	return &figcaption{New("figcaption", false, children)}
}

type figure struct {
	*Tag
}

func Figure(children ...html.Node) *figure {
	return &figure{New("figure", false, children)}
}

type footer struct {
	*Tag
}

func Footer(children ...html.Node) *footer {
	return &footer{New("footer", false, children)}
}

type form struct {
	*Tag
}

func Form(children ...html.Node) *form {
	return &form{New("form", false, children)}
}

func (e *form) Action(value string) *form {
	e.Attribute("action", value)
	return e
}

func (e *form) ActionIf(condition bool, value string) *form {
	if condition {
		e.Attribute("action", value)
	}
	return e
}

func (e *form) Method(value string) *form {
	e.Attribute("method", value)
	return e
}

func (e *form) MethodIf(condition bool, value string) *form {
	if condition {
		e.Attribute("method", value)
	}
	return e
}

func (e *form) Enctype(value string) *form {
	e.Attribute("enctype", value)
	return e
}

func (e *form) EnctypeIf(condition bool, value string) *form {
	if condition {
		e.Attribute("enctype", value)
	}
	return e
}

func (e *form) Novalidate(value string) *form {
	e.Attribute("novalidate", value)
	return e
}

func (e *form) NovalidateIf(condition bool, value string) *form {
	if condition {
		e.Attribute("novalidate", value)
	}
	return e
}

func (e *form) Autocomplete(value string) *form {
	e.Attribute("autocomplete", value)
	return e
}

func (e *form) AutocompleteIf(condition bool, value string) *form {
	if condition {
		e.Attribute("autocomplete", value)
	}
	return e
}

type h1 struct {
	*Tag
}

func H1(children ...html.Node) *h1 {
	return &h1{New("h1", false, children)}
}

type h2 struct {
	*Tag
}

func H2(children ...html.Node) *h2 {
	return &h2{New("h2", false, children)}
}

type h3 struct {
	*Tag
}

func H3(children ...html.Node) *h3 {
	return &h3{New("h3", false, children)}
}

type h4 struct {
	*Tag
}

func H4(children ...html.Node) *h4 {
	return &h4{New("h4", false, children)}
}

type h5 struct {
	*Tag
}

func H5(children ...html.Node) *h5 {
	return &h5{New("h5", false, children)}
}

type h6 struct {
	*Tag
}

func H6(children ...html.Node) *h6 {
	return &h6{New("h6", false, children)}
}

type head struct {
	*Tag
}

func Head(children ...html.Node) *head {
	return &head{New("head", false, children)}
}

type header struct {
	*Tag
}

func Header(children ...html.Node) *header {
	return &header{New("header", false, children)}
}

type hr struct {
	*Tag
}

func Hr(children ...html.Node) *hr {
	return &hr{New("hr", true, children)}
}

type html_ struct {
	*Tag
}

func (children ...html.Node) *html_ {
	return &html_{New("html", false, children)}
}

type i struct {
	*Tag
}

func I(children ...html.Node) *i {
	return &i{New("i", false, children)}
}

type iframe struct {
	*Tag
}

func Iframe(children ...html.Node) *iframe {
	return &iframe{New("iframe", false, children)}
}

func (e *iframe) Allow(value string) *iframe {
	e.Attribute("allow", value)
	return e
}

func (e *iframe) AllowIf(condition bool, value string) *iframe {
	if condition {
		e.Attribute("allow", value)
	}
	return e
}

func (e *iframe) Allowfullscreen(value string) *iframe {
	e.Attribute("allowfullscreen", value)
	return e
}

func (e *iframe) AllowfullscreenIf(condition bool, value string) *iframe {
	if condition {
		e.Attribute("allowfullscreen", value)
	}
	return e
}

func (e *iframe) Sandbox(value string) *iframe {
	e.Attribute("sandbox", value)
	return e
}

func (e *iframe) SandboxIf(condition bool, value string) *iframe {
	if condition {
		e.Attribute("sandbox", value)
	}
	return e
}

func (e *iframe) Srcdoc(value string) *iframe {
	e.Attribute("srcdoc", value)
	return e
}

func (e *iframe) SrcdocIf(condition bool, value string) *iframe {
	if condition {
		e.Attribute("srcdoc", value)
	}
	return e
}

func (e *iframe) Loading(value string) *iframe {
	e.Attribute("loading", value)
	return e
}

func (e *iframe) LoadingIf(condition bool, value string) *iframe {
	if condition {
		e.Attribute("loading", value)
	}
	return e
}

type img struct {
	*Tag
}

func Img(children ...html.Node) *img {
	return &img{New("img", true, children)}
}

func (e *img) Srcset(value string) *img {
	e.Attribute("srcset", value)
	return e
}

func (e *img) SrcsetIf(condition bool, value string) *img {
	if condition {
		e.Attribute("srcset", value)
	}
	return e
}

func (e *img) Sizes(value string) *img {
	e.Attribute("sizes", value)
	return e
}

func (e *img) SizesIf(condition bool, value string) *img {
	if condition {
		e.Attribute("sizes", value)
	}
	return e
}

func (e *img) Crossorigin(value string) *img {
	e.Attribute("crossorigin", value)
	return e
}

func (e *img) CrossoriginIf(condition bool, value string) *img {
	if condition {
		e.Attribute("crossorigin", value)
	}
	return e
}

func (e *img) Decoding(value string) *img {
	e.Attribute("decoding", value)
	return e
}

func (e *img) DecodingIf(condition bool, value string) *img {
	if condition {
		e.Attribute("decoding", value)
	}
	return e
}

func (e *img) Ismap(value string) *img {
	e.Attribute("ismap", value)
	return e
}

func (e *img) IsmapIf(condition bool, value string) *img {
	if condition {
		e.Attribute("ismap", value)
	}
	return e
}

func (e *img) Usemap(value string) *img {
	e.Attribute("usemap", value)
	return e
}

func (e *img) UsemapIf(condition bool, value string) *img {
	if condition {
		e.Attribute("usemap", value)
	}
	return e
}

type input struct {
	*Tag
}

func Input(children ...html.Node) *input {
	return &input{New("input", true, children)}
}

func (e *input) Placeholder(value string) *input {
	e.Attribute("placeholder", value)
	return e
}

func (e *input) PlaceholderIf(condition bool, value string) *input {
	if condition {
		e.Attribute("placeholder", value)
	}
	return e
}

func (e *input) Required(value string) *input {
	e.Attribute("required", value)
	return e
}

func (e *input) RequiredIf(condition bool, value string) *input {
	if condition {
		e.Attribute("required", value)
	}
	return e
}

func (e *input) Readonly(value string) *input {
	e.Attribute("readonly", value)
	return e
}

func (e *input) ReadonlyIf(condition bool, value string) *input {
	if condition {
		e.Attribute("readonly", value)
	}
	return e
}

func (e *input) Checked(value string) *input {
	e.Attribute("checked", value)
	return e
}

func (e *input) CheckedIf(condition bool, value string) *input {
	if condition {
		e.Attribute("checked", value)
	}
	return e
}

func (e *input) Min(value string) *input {
	e.Attribute("min", value)
	return e
}

func (e *input) MinIf(condition bool, value string) *input {
	if condition {
		e.Attribute("min", value)
	}
	return e
}

func (e *input) Max(value string) *input {
	e.Attribute("max", value)
	return e
}

func (e *input) MaxIf(condition bool, value string) *input {
	if condition {
		e.Attribute("max", value)
	}
	return e
}

func (e *input) Pattern(value string) *input {
	e.Attribute("pattern", value)
	return e
}

func (e *input) PatternIf(condition bool, value string) *input {
	if condition {
		e.Attribute("pattern", value)
	}
	return e
}

type ins struct {
	*Tag
}

func Ins(children ...html.Node) *ins {
	return &ins{New("ins", false, children)}
}

type kbd struct {
	*Tag
}

func Kbd(children ...html.Node) *kbd {
	return &kbd{New("kbd", false, children)}
}

type label struct {
	*Tag
}

func Label(children ...html.Node) *label {
	return &label{New("label", false, children)}
}

func (e *label) For(value string) *label {
	e.Attribute("for", value)
	return e
}

func (e *label) ForIf(condition bool, value string) *label {
	if condition {
		e.Attribute("for", value)
	}
	return e
}

type legend struct {
	*Tag
}

func Legend(children ...html.Node) *legend {
	return &legend{New("legend", false, children)}
}

type li struct {
	*Tag
}

func Li(children ...html.Node) *li {
	return &li{New("li", false, children)}
}

type link struct {
	*Tag
}

func Link(children ...html.Node) *link {
	return &link{New("link", true, children)}
}

func (e *link) Integrity(value string) *link {
	e.Attribute("integrity", value)
	return e
}

func (e *link) IntegrityIf(condition bool, value string) *link {
	if condition {
		e.Attribute("integrity", value)
	}
	return e
}

type main struct {
	*Tag
}

func Main(children ...html.Node) *main {
	return &main{New("main", false, children)}
}

type map_ struct {
	*Tag
}

func Map(children ...html.Node) *map_ {
	return &map_{New("map", false, children)}
}

type mark struct {
	*Tag
}

func Mark(children ...html.Node) *mark {
	return &mark{New("mark", false, children)}
}

type meta struct {
	*Tag
}

func Meta(children ...html.Node) *meta {
	return &meta{New("meta", true, children)}
}

func (e *meta) Content(value string) *meta {
	e.Attribute("content", value)
	return e
}

func (e *meta) ContentIf(condition bool, value string) *meta {
	if condition {
		e.Attribute("content", value)
	}
	return e
}

func (e *meta) Charset(value string) *meta {
	e.Attribute("charset", value)
	return e
}

func (e *meta) CharsetIf(condition bool, value string) *meta {
	if condition {
		e.Attribute("charset", value)
	}
	return e
}

func (e *meta) HttpEquiv(value string) *meta {
	e.Attribute("http-equiv", value)
	return e
}

func (e *meta) HttpEquivIf(condition bool, value string) *meta {
	if condition {
		e.Attribute("http-equiv", value)
	}
	return e
}

type meter struct {
	*Tag
}

func Meter(children ...html.Node) *meter {
	return &meter{New("meter", false, children)}
}

func (e *meter) Low(value string) *meter {
	e.Attribute("low", value)
	return e
}

func (e *meter) LowIf(condition bool, value string) *meter {
	if condition {
		e.Attribute("low", value)
	}
	return e
}

func (e *meter) High(value string) *meter {
	e.Attribute("high", value)
	return e
}

func (e *meter) HighIf(condition bool, value string) *meter {
	if condition {
		e.Attribute("high", value)
	}
	return e
}

func (e *meter) Optimum(value string) *meter {
	e.Attribute("optimum", value)
	return e
}

func (e *meter) OptimumIf(condition bool, value string) *meter {
	if condition {
		e.Attribute("optimum", value)
	}
	return e
}

type nav struct {
	*Tag
}

func Nav(children ...html.Node) *nav {
	return &nav{New("nav", false, children)}
}

type noscript struct {
	*Tag
}

func Noscript(children ...html.Node) *noscript {
	return &noscript{New("noscript", false, children)}
}

type object struct {
	*Tag
}

func Object(children ...html.Node) *object {
	return &object{New("object", false, children)}
}

type ol struct {
	*Tag
}

func Ol(children ...html.Node) *ol {
	return &ol{New("ol", false, children)}
}

func (e *ol) Start(value string) *ol {
	e.Attribute("start", value)
	return e
}

func (e *ol) StartIf(condition bool, value string) *ol {
	if condition {
		e.Attribute("start", value)
	}
	return e
}

func (e *ol) Reversed(value string) *ol {
	e.Attribute("reversed", value)
	return e
}

func (e *ol) ReversedIf(condition bool, value string) *ol {
	if condition {
		e.Attribute("reversed", value)
	}
	return e
}

type optgroup struct {
	*Tag
}

func Optgroup(children ...html.Node) *optgroup {
	return &optgroup{New("optgroup", false, children)}
}

func (e *optgroup) Label(value string) *optgroup {
	e.Attribute("label", value)
	return e
}

func (e *optgroup) LabelIf(condition bool, value string) *optgroup {
	if condition {
		e.Attribute("label", value)
	}
	return e
}

type option struct {
	*Tag
}

func Option(children ...html.Node) *option {
	return &option{New("option", false, children)}
}

func (e *option) Selected(value string) *option {
	e.Attribute("selected", value)
	return e
}

func (e *option) SelectedIf(condition bool, value string) *option {
	if condition {
		e.Attribute("selected", value)
	}
	return e
}

type output struct {
	*Tag
}

func Output(children ...html.Node) *output {
	return &output{New("output", false, children)}
}

type p struct {
	*Tag
}

func P(children ...html.Node) *p {
	return &p{New("p", false, children)}
}

type param struct {
	*Tag
}

func Param(children ...html.Node) *param {
	return &param{New("param", true, children)}
}

type picture struct {
	*Tag
}

func Picture(children ...html.Node) *picture {
	return &picture{New("picture", false, children)}
}

type pre struct {
	*Tag
}

func Pre(children ...html.Node) *pre {
	return &pre{New("pre", false, children)}
}

type progress struct {
	*Tag
}

func Progress(children ...html.Node) *progress {
	return &progress{New("progress", false, children)}
}

type q struct {
	*Tag
}

func Q(children ...html.Node) *q {
	return &q{New("q", false, children)}
}

type rp struct {
	*Tag
}

func Rp(children ...html.Node) *rp {
	return &rp{New("rp", false, children)}
}

type rt struct {
	*Tag
}

func Rt(children ...html.Node) *rt {
	return &rt{New("rt", false, children)}
}

type ruby struct {
	*Tag
}

func Ruby(children ...html.Node) *ruby {
	return &ruby{New("ruby", false, children)}
}

type s struct {
	*Tag
}

func S(children ...html.Node) *s {
	return &s{New("s", false, children)}
}

type samp struct {
	*Tag
}

func Samp(children ...html.Node) *samp {
	return &samp{New("samp", false, children)}
}

type script struct {
	*Tag
}

func Script(children ...html.Node) *script {
	return &script{New("script", false, children)}
}

func (e *script) Async(value string) *script {
	e.Attribute("async", value)
	return e
}

func (e *script) AsyncIf(condition bool, value string) *script {
	if condition {
		e.Attribute("async", value)
	}
	return e
}

func (e *script) Defer(value string) *script {
	e.Attribute("defer", value)
	return e
}

func (e *script) DeferIf(condition bool, value string) *script {
	if condition {
		e.Attribute("defer", value)
	}
	return e
}

func (e *script) Nomodule(value string) *script {
	e.Attribute("nomodule", value)
	return e
}

func (e *script) NomoduleIf(condition bool, value string) *script {
	if condition {
		e.Attribute("nomodule", value)
	}
	return e
}

type section struct {
	*Tag
}

func Section(children ...html.Node) *section {
	return &section{New("section", false, children)}
}

type select_ struct {
	*Tag
}

func Select(children ...html.Node) *select_ {
	return &select_{New("select", false, children)}
}

func (e *select_) Multiple(value string) *select_ {
	e.Attribute("multiple", value)
	return e
}

func (e *select_) MultipleIf(condition bool, value string) *select_ {
	if condition {
		e.Attribute("multiple", value)
	}
	return e
}

func (e *select_) Size(value string) *select_ {
	e.Attribute("size", value)
	return e
}

func (e *select_) SizeIf(condition bool, value string) *select_ {
	if condition {
		e.Attribute("size", value)
	}
	return e
}

type small struct {
	*Tag
}

func Small(children ...html.Node) *small {
	return &small{New("small", false, children)}
}

type source struct {
	*Tag
}

func Source(children ...html.Node) *source {
	return &source{New("source", true, children)}
}

type span struct {
	*Tag
}

func Span(children ...html.Node) *span {
	return &span{New("span", false, children)}
}

type strong struct {
	*Tag
}

func Strong(children ...html.Node) *strong {
	return &strong{New("strong", false, children)}
}

type style struct {
	*Tag
}

func Style(children ...html.Node) *style {
	return &style{New("style", false, children)}
}

type sub struct {
	*Tag
}

func Sub(children ...html.Node) *sub {
	return &sub{New("sub", false, children)}
}

type summary struct {
	*Tag
}

func Summary(children ...html.Node) *summary {
	return &summary{New("summary", false, children)}
}

type sup struct {
	*Tag
}

func Sup(children ...html.Node) *sup {
	return &sup{New("sup", false, children)}
}

type svg struct {
	*Tag
}

func (children ...html.Node) *svg {
	return &svg{New("svg", false, children)}
}

func (e *svg) ViewBox(value string) *svg {
	e.Attribute("viewBox", value)
	return e
}

func (e *svg) ViewBoxIf(condition bool, value string) *svg {
	if condition {
		e.Attribute("viewBox", value)
	}
	return e
}

func (e *svg) PreserveAspectRatio(value string) *svg {
	e.Attribute("preserveAspectRatio", value)
	return e
}

func (e *svg) PreserveAspectRatioIf(condition bool, value string) *svg {
	if condition {
		e.Attribute("preserveAspectRatio", value)
	}
	return e
}

func (e *svg) Xmlns(value string) *svg {
	e.Attribute("xmlns", value)
	return e
}

func (e *svg) XmlnsIf(condition bool, value string) *svg {
	if condition {
		e.Attribute("xmlns", value)
	}
	return e
}

func (e *svg) Version(value string) *svg {
	e.Attribute("version", value)
	return e
}

func (e *svg) VersionIf(condition bool, value string) *svg {
	if condition {
		e.Attribute("version", value)
	}
	return e
}

type table struct {
	*Tag
}

func Table(children ...html.Node) *table {
	return &table{New("table", false, children)}
}

func (e *table) Border(value string) *table {
	e.Attribute("border", value)
	return e
}

func (e *table) BorderIf(condition bool, value string) *table {
	if condition {
		e.Attribute("border", value)
	}
	return e
}

func (e *table) Cellpadding(value string) *table {
	e.Attribute("cellpadding", value)
	return e
}

func (e *table) CellpaddingIf(condition bool, value string) *table {
	if condition {
		e.Attribute("cellpadding", value)
	}
	return e
}

func (e *table) Cellspacing(value string) *table {
	e.Attribute("cellspacing", value)
	return e
}

func (e *table) CellspacingIf(condition bool, value string) *table {
	if condition {
		e.Attribute("cellspacing", value)
	}
	return e
}

type tbody struct {
	*Tag
}

func Tbody(children ...html.Node) *tbody {
	return &tbody{New("tbody", false, children)}
}

type td struct {
	*Tag
}

func Td(children ...html.Node) *td {
	return &td{New("td", false, children)}
}

func (e *td) Colspan(value string) *td {
	e.Attribute("colspan", value)
	return e
}

func (e *td) ColspanIf(condition bool, value string) *td {
	if condition {
		e.Attribute("colspan", value)
	}
	return e
}

func (e *td) Rowspan(value string) *td {
	e.Attribute("rowspan", value)
	return e
}

func (e *td) RowspanIf(condition bool, value string) *td {
	if condition {
		e.Attribute("rowspan", value)
	}
	return e
}

func (e *td) Headers(value string) *td {
	e.Attribute("headers", value)
	return e
}

func (e *td) HeadersIf(condition bool, value string) *td {
	if condition {
		e.Attribute("headers", value)
	}
	return e
}

type template struct {
	*Tag
}

func Template(children ...html.Node) *template {
	return &template{New("template", false, children)}
}

type textarea struct {
	*Tag
}

func Textarea(children ...html.Node) *textarea {
	return &textarea{New("textarea", false, children)}
}

func (e *textarea) Rows(value string) *textarea {
	e.Attribute("rows", value)
	return e
}

func (e *textarea) RowsIf(condition bool, value string) *textarea {
	if condition {
		e.Attribute("rows", value)
	}
	return e
}

func (e *textarea) Cols(value string) *textarea {
	e.Attribute("cols", value)
	return e
}

func (e *textarea) ColsIf(condition bool, value string) *textarea {
	if condition {
		e.Attribute("cols", value)
	}
	return e
}

func (e *textarea) Maxlength(value string) *textarea {
	e.Attribute("maxlength", value)
	return e
}

func (e *textarea) MaxlengthIf(condition bool, value string) *textarea {
	if condition {
		e.Attribute("maxlength", value)
	}
	return e
}

func (e *textarea) Minlength(value string) *textarea {
	e.Attribute("minlength", value)
	return e
}

func (e *textarea) MinlengthIf(condition bool, value string) *textarea {
	if condition {
		e.Attribute("minlength", value)
	}
	return e
}

func (e *textarea) Wrap(value string) *textarea {
	e.Attribute("wrap", value)
	return e
}

func (e *textarea) WrapIf(condition bool, value string) *textarea {
	if condition {
		e.Attribute("wrap", value)
	}
	return e
}

type tfoot struct {
	*Tag
}

func Tfoot(children ...html.Node) *tfoot {
	return &tfoot{New("tfoot", false, children)}
}

type th struct {
	*Tag
}

func Th(children ...html.Node) *th {
	return &th{New("th", false, children)}
}

func (e *th) Scope(value string) *th {
	e.Attribute("scope", value)
	return e
}

func (e *th) ScopeIf(condition bool, value string) *th {
	if condition {
		e.Attribute("scope", value)
	}
	return e
}

type thead struct {
	*Tag
}

func Thead(children ...html.Node) *thead {
	return &thead{New("thead", false, children)}
}

type time struct {
	*Tag
}

func Time(children ...html.Node) *time {
	return &time{New("time", false, children)}
}

func (e *time) Datetime(value string) *time {
	e.Attribute("datetime", value)
	return e
}

func (e *time) DatetimeIf(condition bool, value string) *time {
	if condition {
		e.Attribute("datetime", value)
	}
	return e
}

type title struct {
	*Tag
}

func Title(children ...html.Node) *title {
	return &title{New("title", false, children)}
}

type tr struct {
	*Tag
}

func Tr(children ...html.Node) *tr {
	return &tr{New("tr", false, children)}
}

type track struct {
	*Tag
}

func Track(children ...html.Node) *track {
	return &track{New("track", true, children)}
}

type u struct {
	*Tag
}

func U(children ...html.Node) *u {
	return &u{New("u", false, children)}
}

type ul struct {
	*Tag
}

func Ul(children ...html.Node) *ul {
	return &ul{New("ul", false, children)}
}

type var_ struct {
	*Tag
}

func Var(children ...html.Node) *var_ {
	return &var_{New("var", false, children)}
}

type video struct {
	*Tag
}

func Video(children ...html.Node) *video {
	return &video{New("video", false, children)}
}

func (e *video) Poster(value string) *video {
	e.Attribute("poster", value)
	return e
}

func (e *video) PosterIf(condition bool, value string) *video {
	if condition {
		e.Attribute("poster", value)
	}
	return e
}

type wbr struct {
	*Tag
}

func Wbr(children ...html.Node) *wbr {
	return &wbr{New("wbr", true, children)}
}

type circle struct {
	*Tag
}

func Circle(children ...html.Node) *circle {
	return &circle{New("circle", false, children)}
}

func (e *circle) Cx(value string) *circle {
	e.Attribute("cx", value)
	return e
}

func (e *circle) CxIf(condition bool, value string) *circle {
	if condition {
		e.Attribute("cx", value)
	}
	return e
}

func (e *circle) Cy(value string) *circle {
	e.Attribute("cy", value)
	return e
}

func (e *circle) CyIf(condition bool, value string) *circle {
	if condition {
		e.Attribute("cy", value)
	}
	return e
}

func (e *circle) R(value string) *circle {
	e.Attribute("r", value)
	return e
}

func (e *circle) RIf(condition bool, value string) *circle {
	if condition {
		e.Attribute("r", value)
	}
	return e
}

func (e *circle) Fill(value string) *circle {
	e.Attribute("fill", value)
	return e
}

func (e *circle) FillIf(condition bool, value string) *circle {
	if condition {
		e.Attribute("fill", value)
	}
	return e
}

func (e *circle) Stroke(value string) *circle {
	e.Attribute("stroke", value)
	return e
}

func (e *circle) StrokeIf(condition bool, value string) *circle {
	if condition {
		e.Attribute("stroke", value)
	}
	return e
}

type ellipse struct {
	*Tag
}

func Ellipse(children ...html.Node) *ellipse {
	return &ellipse{New("ellipse", false, children)}
}

type g struct {
	*Tag
}

func G(children ...html.Node) *g {
	return &g{New("g", false, children)}
}

func (e *g) Transform(value string) *g {
	e.Attribute("transform", value)
	return e
}

func (e *g) TransformIf(condition bool, value string) *g {
	if condition {
		e.Attribute("transform", value)
	}
	return e
}

type line struct {
	*Tag
}

func Line(children ...html.Node) *line {
	return &line{New("line", false, children)}
}

func (e *line) X1(value string) *line {
	e.Attribute("x1", value)
	return e
}

func (e *line) X1If(condition bool, value string) *line {
	if condition {
		e.Attribute("x1", value)
	}
	return e
}

func (e *line) Y1(value string) *line {
	e.Attribute("y1", value)
	return e
}

func (e *line) Y1If(condition bool, value string) *line {
	if condition {
		e.Attribute("y1", value)
	}
	return e
}

func (e *line) X2(value string) *line {
	e.Attribute("x2", value)
	return e
}

func (e *line) X2If(condition bool, value string) *line {
	if condition {
		e.Attribute("x2", value)
	}
	return e
}

func (e *line) Y2(value string) *line {
	e.Attribute("y2", value)
	return e
}

func (e *line) Y2If(condition bool, value string) *line {
	if condition {
		e.Attribute("y2", value)
	}
	return e
}

func (e *line) StrokeWidth(value string) *line {
	e.Attribute("stroke-width", value)
	return e
}

func (e *line) StrokeWidthIf(condition bool, value string) *line {
	if condition {
		e.Attribute("stroke-width", value)
	}
	return e
}

type path struct {
	*Tag
}

func Path(children ...html.Node) *path {
	return &path{New("path", false, children)}
}

func (e *path) D(value string) *path {
	e.Attribute("d", value)
	return e
}

func (e *path) DIf(condition bool, value string) *path {
	if condition {
		e.Attribute("d", value)
	}
	return e
}

type polygon struct {
	*Tag
}

func Polygon(children ...html.Node) *polygon {
	return &polygon{New("polygon", false, children)}
}

func (e *polygon) Points(value string) *polygon {
	e.Attribute("points", value)
	return e
}

func (e *polygon) PointsIf(condition bool, value string) *polygon {
	if condition {
		e.Attribute("points", value)
	}
	return e
}

type polyline struct {
	*Tag
}

func Polyline(children ...html.Node) *polyline {
	return &polyline{New("polyline", false, children)}
}

type rect struct {
	*Tag
}

func Rect(children ...html.Node) *rect {
	return &rect{New("rect", false, children)}
}

func (e *rect) X(value string) *rect {
	e.Attribute("x", value)
	return e
}

func (e *rect) XIf(condition bool, value string) *rect {
	if condition {
		e.Attribute("x", value)
	}
	return e
}

func (e *rect) Y(value string) *rect {
	e.Attribute("y", value)
	return e
}

func (e *rect) YIf(condition bool, value string) *rect {
	if condition {
		e.Attribute("y", value)
	}
	return e
}

func (e *rect) Rx(value string) *rect {
	e.Attribute("rx", value)
	return e
}

func (e *rect) RxIf(condition bool, value string) *rect {
	if condition {
		e.Attribute("rx", value)
	}
	return e
}

func (e *rect) Ry(value string) *rect {
	e.Attribute("ry", value)
	return e
}

func (e *rect) RyIf(condition bool, value string) *rect {
	if condition {
		e.Attribute("ry", value)
	}
	return e
}

type text struct {
	*Tag
}

func Text(children ...html.Node) *text {
	return &text{New("text", false, children)}
}

func (e *text) FontSize(value string) *text {
	e.Attribute("font-size", value)
	return e
}

func (e *text) FontSizeIf(condition bool, value string) *text {
	if condition {
		e.Attribute("font-size", value)
	}
	return e
}

func (e *text) TextAnchor(value string) *text {
	e.Attribute("text-anchor", value)
	return e
}

func (e *text) TextAnchorIf(condition bool, value string) *text {
	if condition {
		e.Attribute("text-anchor", value)
	}
	return e
}

type use struct {
	*Tag
}

func Use(children ...html.Node) *use {
	return &use{New("use", false, children)}
}
