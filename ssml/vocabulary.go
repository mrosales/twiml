package ssml

import "encoding/xml"

// Phrases
type Text string

func (t Text) Phrase() {}
func (t Text) Encode(e *xml.Encoder, s xml.StartElement) error {
	return e.EncodeToken(xml.CharData([]byte(t)))
}

type Paragraph struct {
	XMLName xml.Name `xml:"p"`
	Child   Builder  `xml:",innerxml"`
}

func (p *Paragraph) Phrase()         {}
func (p *Paragraph) Type() string    { return "paragraph" }
func (p *Paragraph) Validate() error { return nil }

type Sentence struct {
	XMLName xml.Name `xml:"s"`
	Child   Builder  `xml:",innerxml"`
}

func (s *Sentence) Phrase()         {}
func (s *Sentence) Type() string    { return "sentence" }
func (s *Sentence) Validate() error { return nil }

// Represents a pause in the speech. Set the length of the pause with the strength or time attributes.
type Break struct {
	XMLName xml.Name `xml:"break"`
	// The strength of the pause. See PauseStrength constants for details
	Strength PauseStrength `xml:"strength,attr,omitempty"`

	// Duration of the pause; up to 10 seconds (10s) or 10000 milliseconds (10000ms).
	Time Duration `xml:"time,attr,omitempty"`
}

func (b *Break) Phrase()         {}
func (b *Break) Type() string    { return "break" }
func (b *Break) Validate() error { return nil }

// The audio tag lets you provide the URL for an MP3 file that the Alexa service
// can play while rendering a response. You can use this to embed short, pre-recorded
// audio within your service's response. For example, you could include sound effects
// alongside your text-to-speech responses, or provide responses using a voice
// associated with your brand. For more information, see
// "Including Short Pre-Recorded Audio in your Response".
type Audio struct {
	XMLName xml.Name `xml:"audio"`

	// Specifies the URL for the MP3 file. Note the following requirements and limitations:
	//
	// The MP3 must be hosted at an Internet-accessible HTTPS endpoint. HTTPS is required, and the domain hosting the MP3 file must present a valid, trusted SSL certificate. Self-signed certificates cannot be used.
	// The MP3 must not contain any customer-specific or other sensitive information.
	// The MP3 must be a valid MP3 file (MPEG version 2).
	// The audio file cannot be longer than 240 seconds.
	// The bit rate must be 48 kbps. Note that this bit rate gives a good result when used with spoken content, but is generally not a high enough quality for music.
	// The sample rate must be 22050Hz, 24000Hz, or 16000Hz.
	// You may need to use converter software to convert your MP3 files to the required codec version (MPEG version 2) and bit rate (48 kbps).
	Source *URL `xml:"src,attr,omitempty"`
}

func (a *Audio) Phrase()         {}
func (a *Audio) Type() string    { return "audio" }
func (a *Audio) Validate() error { return nil }

type Word struct {
	XMLName xml.Name `xml:"w"`
	Role    Role     `xml:"role,attr,omitempty"`
	Text    string   `xml:",chardata"`
}

func (w *Word) Phrase()         {}
func (w *Word) Type() string    { return "word" }
func (w *Word) Validate() error { return nil }

// Describes how the text should be interpreted. This lets you provide additional context
// to the text and eliminate any ambiguity on how Alexa should render the text. Indicate
// how Alexa should interpret the text with the interpret-as attribute.
type SayAs struct {
	XMLName xml.Name `xml:"say-as"`

	// Describes how the text should be interpreted
	InterpretAs InterpretAs `xml:"interpret-as,attr,omitempty"`

	// Only used when interpret-as is set to date.
	Format DateFormat `xml:"format,attr,omitempty"`

	// Text to control the interpretation of
	Text string `xml:",chardata"`
}

func (s *SayAs) Phrase()         {}
func (s *SayAs) Type() string    { return "say-as" }
func (s *SayAs) Validate() error { return nil }

type Phoneme struct {
	XMLName       xml.Name `xml:"phoneme"`
	Alphabet      Alphabet `xml:"alphabet,attr,omitempty"`
	Pronunciation string   `xml:"ph,attr,omitempty"`
	Text          string   `xml:",chardata"`
}

func (p *Phoneme) Phrase()         {}
func (p *Phoneme) Type() string    { return "phoneme" }
func (p *Phoneme) Validate() error { return nil }

type Prosody struct {
	XMLName xml.Name `xml:"prosody"`
	Volume  Volume   `xml:"volume,attr,omitempty"`
	Pitch   Pitch    `xml:"pitch,attr,omitempty"`
	Rate    Rate     `xml:"rate,attr,omitempty"`
	Child   Builder  `xml:",innerxml"`
}

func (p *Prosody) Phrase()         {}
func (p *Prosody) Type() string    { return "prosody" }
func (p *Prosody) Validate() error { return nil }

type Emphasis struct {
	XMLName xml.Name      `xml:"emphasis"`
	Level   EmphasisLevel `xml:"level,attr"`
	Child   Builder       `xml:",innerxml"`
}

func (e *Emphasis) Phrase()         {}
func (e *Emphasis) Type() string    { return "emphasis" }
func (e *Emphasis) Validate() error { return nil }

type Sub struct {
	XMLName xml.Name `xml:"sub"`
	Alias   string   `xml:"alias,attr"`
	Text    string   `xml:",chardata"`
}

func (s *Sub) Phrase()         {}
func (s *Sub) Type() string    { return "sub" }
func (s *Sub) Validate() error { return nil }

type Language struct {
	XMLName  xml.Name      `xml:"lang"`
	Language VoiceLanguage `xml:"xml:lang,attr"`
	Child    Builder       `xml:",innerxml"`
}

func (l *Language) Phrase()         {}
func (l *Language) Type() string    { return "language" }
func (l *Language) Validate() error { return nil }
