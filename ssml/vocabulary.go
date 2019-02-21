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

func (p *Paragraph) Phrase() {}

type Sentence struct {
	XMLName xml.Name `xml:"s"`
	Child   Builder  `xml:",innerxml"`
}

func (s *Sentence) Phrase() {}

// Represents a pause in the speech. Set the length of the pause with the strength or time attributes.
type Break struct {
	XMLName xml.Name `xml:"b"`
	// The strength of the pause. See PauseStrength constants for details
	Strength PauseStrength `xml:"strength,attr,omitempty"`

	// Duration of the pause; up to 10 seconds (10s) or 10000 milliseconds (10000ms).
	Duration Duration `xml:"duration,attr,omitempty"`
}

func (b *Break) Phrase() {}

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

func (a *Audio) Phrase() {}

type Word struct {
	XMLName xml.Name `xml:"w"`
	Role    Role     `xml:"role,attr,omitempty"`
	Text    string   `xml:",chardata"`
}

func (w *Word) Phrase() {}

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

func (s *SayAs) Phrase() {}

type Phoneme struct {
	XMLName       xml.Name `xml:"phoneme"`
	Alphabet      Alphabet `xml:"alphabet,attr,omitempty"`
	Pronunciation string   `xml:"ph,attr,omitempty"`
	Text          string   `xml:",chardata"`
}

func (p *Phoneme) Phrase() {}

type Prosody struct {
	XMLName xml.Name `xml:"prosody"`
	Volume  Volume   `xml:"volume,attr,omitempty"`
	Pitch   Pitch    `xml:"pitch,attr,omitempty"`
	Rate    Rate     `xml:"rate,attr,omitempty"`
	Child   Builder  `xml:",innerxml"`
}

func (p *Prosody) Phrase() {}

type Emphasis struct {
	XMLName xml.Name      `xml:"emphasis"`
	Level   EmphasisLevel `xml:"level,attr"`
	Child   Builder       `xml:",innerxml"`
}

func (e *Emphasis) Phrase() {}

type Sub struct {
	XMLName xml.Name `xml:"sub"`
	Alias   string   `xml:"alias,attr"`
	Text    string   `xml:",chardata"`
}

func (s *Sub) Phrase() {}

type Language struct {
	XMLName  xml.Name      `xml:"lang"`
	Language VoiceLanguage `xml:"xml:lang,attr"`
	Child    Builder       `xml:",innerxml"`
}

func (l *Language) Phrase() {}
