package rdf

import (
	"fmt"
)

// Constants for kinds of RDF Term
const (
	UnknownTerm = iota
	IRITerm
	BlankTerm
	LiteralTerm
)

// A Term represents an element of a triple or quad, one of IRI, Blank Node or Literal
type Term struct {
	Value    string // the lexical value of the IRI, literal or blank node label
	Language string // the language tag of a literal
	Datatype string // the data type IRI of a literal
	Kind     int    // the kind of term
}

func (t *Term) String() string {
	switch t.Kind {
	case IRITerm:
		return fmt.Sprintf("<%s>", t.Value)
	case BlankTerm:
		return fmt.Sprintf("_:%s", t.Value)
	case LiteralTerm:
		if t.Language != "" {
			return fmt.Sprintf("\"%s\"@%s", t.Value, t.Language)
		} else if t.Datatype != "" {
			return fmt.Sprintf("\"%s\"^^<%s>", t.Value, t.Datatype)
		} else {
			return fmt.Sprintf("\"%s\"", t.Value)
		}
	default:
		return fmt.Sprintf("[unknown kind: %s]", t.Value)
	}
}

// IRI constructs an IRI term
func IRI(iri string) Term {
	return Term{Value: iri, Kind: IRITerm}
}

// Blank constructs a blank term
func Blank(label string) Term {
	return Term{Value: label, Kind: BlankTerm}
}

// Literal constructs a literal term without a language or datatype IRI.
func Literal(value string) Term {
	return Term{Value: value, Kind: LiteralTerm}
}

// LiteralWithLanguage constructs a literal term with a language.
func LiteralWithLanguage(value, lang string) Term {
	return Term{Value: value, Language: lang, Kind: LiteralTerm}
}

// LiteralWithDatatype constructs a literal term with a datatype IRI.
func LiteralWithDatatype(value, dt string) Term {
	return Term{Value: value, Datatype: dt, Kind: LiteralTerm}
}
