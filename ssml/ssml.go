package ssml

import (
	"encoding/xml"
	"net/url"
	"time"

	"github.com/pkg/errors"
)

type Builder interface {
	Text(string) Builder
	Space() Builder
	Newline() Builder
	Paragraph(string) Builder
	Sentence(string) Builder
	Break(PauseStrength) Builder
	BreakFor(time.Duration) Builder
	Audio(*url.URL) Builder
	Word(string, Role) Builder
	SayAs(string, InterpretAs) Builder
	Date(time.Time, DateFormat) Builder
	Phoneme(string, Alphabet, string) Builder
	Prosody(string, Volume, Pitch, Rate) Builder
	Volume(string, Volume) Builder
	Pitch(string, Pitch) Builder
	Rate(string, Rate) Builder
	Emphasis(string, EmphasisLevel) Builder
	Language(string, VoiceLanguage) Builder
	Sub(text string, alias string) Builder

	AddChild(Builder) Builder
	Phrase
	xml.Marshaler
}

type Phrase interface {
	Phrase()
}

func Build(phrases ...Phrase) Builder {
	return &builder{phrases}
}

///////////////////////////////// Internal Builder Implementation /////////////////////////////////

// interface that allows an object to encode themselves
type xmlEncoder interface {
	Encode(e *xml.Encoder, start xml.StartElement) error
}

type builder struct {
	expression []Phrase
}

func (b *builder) Phrase() {}

func (b *builder) add(p Phrase) Builder {
	b.expression = append(b.expression, p)
	return b
}

func (b *builder) Text(s string) Builder {
	return b.add(Text(s))
}
func (b *builder) Space() Builder {
	return b.add(Text(" "))
}
func (b *builder) Newline() Builder {
	return b.add(Text("\n"))
}

func (b *builder) Paragraph(s string) Builder {
	return b.add(&Paragraph{Child: Build().Text(s)})
}

func (b *builder) Sentence(s string) Builder {
	return b.add(&Sentence{Child: Build().Text(s)})
}

func (b *builder) Break(p PauseStrength) Builder {
	return b.add(&Break{Strength: p})
}

func (b *builder) BreakFor(d time.Duration) Builder {
	return b.add(&Break{Duration: Duration(d)})
}

func (b *builder) Audio(source *url.URL) Builder {
	u := URL(*source)
	return b.add(&Audio{Source: &u})
}

func (b *builder) Word(w string, r Role) Builder {
	return b.add(&Word{Role: r, Text: w})
}

func (b *builder) SayAs(text string, interpretAs InterpretAs) Builder {
	return b.add(&SayAs{InterpretAs: interpretAs, Text: text})
}

func (b *builder) Date(d time.Time, format DateFormat) Builder {
	return b.add(&SayAs{InterpretAs: InterpretAsDate, Format: format, Text: format.Format(d)})
}

func (b *builder) Phoneme(text string, alphabet Alphabet, pronunciation string) Builder {
	return b.add(&Phoneme{Alphabet: alphabet, Pronunciation: pronunciation, Text: text})
}

func (b *builder) Prosody(text string, volume Volume, pitch Pitch, rate Rate) Builder {
	return b.add(&Prosody{Volume: volume, Pitch: pitch, Rate: rate, Child: Build().Text(text)})
}

func (b *builder) Volume(text string, volume Volume) Builder {
	return b.add(&Prosody{Volume: volume, Child: Build().Text(text)})
}

func (b *builder) Pitch(text string, pitch Pitch) Builder {
	return b.add(&Prosody{Pitch: pitch, Child: Build().Text(text)})
}

func (b *builder) Rate(text string, rate Rate) Builder {
	return b.add(&Prosody{Rate: rate, Child: Build().Text(text)})
}

func (b *builder) Emphasis(text string, level EmphasisLevel) Builder {
	return b.add(&Emphasis{Level: level, Child: Build().Text(text)})
}

func (b *builder) Sub(text, alias string) Builder {
	return b.add(&Sub{Alias: alias, Text: text})
}

func (b *builder) Language(text string, language VoiceLanguage) Builder {
	return b.add(&Language{Language: language, Child: Build().Text(text)})
}

func (b *builder) AddChild(child Builder) Builder {
	return b.add(child)
}

// Implement the encodeXML interface in order to support nested builders
func (b *builder) Encode(e *xml.Encoder, start xml.StartElement) error {
	if len(b.expression) == 0 {
		return nil
	}

	for _, phrase := range b.expression {
		var err error
		if enc, ok := phrase.(xmlEncoder); ok {
			err = enc.Encode(e, start)
		} else {
			err = e.Encode(phrase)
		}
		if err != nil {
			return errors.Wrapf(err, "failed encoding element %v to xml", phrase)
		}
	}
	return nil
}

func (b *builder) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return b.Encode(e, start)
}
