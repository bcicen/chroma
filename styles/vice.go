package styles

import (
	"github.com/alecthomas/chroma"
)

// Vice style.
var Vice = Register(chroma.MustNewStyle("vice", chroma.StyleEntries{
	chroma.Text:                  "#AEAFB0",
	chroma.Error:                 "bg:#131416 #ff00ff",
	chroma.Comment:               "#cccc99",
	chroma.Keyword:               "#12e4dc",
	chroma.KeywordNamespace:      "#7b95ff",
	chroma.KeywordType:           "#00ffaf",
	chroma.KeywordConstant:       "#fff35b",
	chroma.KeywordDeclaration:    "#ff3aa8",
	chroma.Operator:              "#6D6E70",
	chroma.Punctuation:           "#8B8C8E",
	chroma.Name:                  "#AEAFB0",
	chroma.NameAttribute:         "#F4F5F7",
	chroma.NameClass:             "#ffb766",
	chroma.NameBuiltin:           "#ff3aa8",
	chroma.NameConstant:          "#ef6155",
	chroma.NameDecorator:         "#5bc4bf",
	chroma.NameException:         "#ef6155",
	chroma.NameFunction:          "#D6D7D9",
	chroma.NameNamespace:         "#AEAFB0",
	chroma.NameBuiltinPseudo:     "#A9AAAC",
	chroma.NameOther:             "#AEAFB0",
	chroma.NameTag:               "#5bc4bf",
	chroma.NameVariable:          "#f2f2f2",
	chroma.LiteralNumber:         "#9cb0ff",
	chroma.Literal:               "#00ffaf",
	chroma.LiteralDate:           "#00ffaf",
	chroma.LiteralString:         "#7fefcc",
	chroma.LiteralStringChar:     "#e7e9db",
	chroma.LiteralStringDoc:      "#7fefcc",
	chroma.LiteralStringEscape:   "#ff00ff",
	chroma.LiteralStringInterpol: "#9cb0ff",
	chroma.GenericDeleted:        "#ef6155",
	chroma.GenericEmph:           "italic",
	chroma.GenericHeading:        "bold #e7e9db",
	chroma.GenericInserted:       "#48b685",
	chroma.GenericPrompt:         "bold #776e71",
	chroma.GenericStrong:         "bold",
	chroma.GenericSubheading:     "bold #5bc4bf",
	chroma.Background:            "bg:#121217",
}))
