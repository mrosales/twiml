package ssml

const (
	// No pause should be outputted. This can be used to remove a pause that would normally occur (such as after a period).
	PauseStrengthNone PauseStrength = "none"
	// No pause should be outputted (same as none).
	PauseStrengthExtraWeak PauseStrength = "x-weak"
	// Treat adjacent words as if separated by a single comma (equivalent to  medium).
	PauseStrengthWeak PauseStrength = "weak"
	// Treat adjacent words as if separated by a single comma.
	PauseStrengthMedium PauseStrength = "medium"
	// Make a sentence break (equivalent to using the <s> tag).
	PauseStrengthStrong PauseStrength = "strong"
	// Make a paragraph break (equivalent to using the <p> tag).
	PauseStrengthExtraStrong PauseStrength = "x-strong"
)

type Role string

const (
	roleBase           = "amazon:"
	RoleVerb           = Role(roleBase + "VB")
	RolePastParticiple = Role(roleBase + "VBD")
	RoleNoun           = Role(roleBase + "NN")
	RoleSense1         = Role(roleBase + "SENSE_1")
)

type InterpretAs string

const (
	InterpretAsCharacters   = InterpretAs("characters")
	InterpretAsSpellOut     = InterpretAs("spell-out")
	InterpretAsCardinal     = InterpretAs("cardinal")
	InterpretAsNumber       = InterpretAs("number")
	InterpretAsOrdinal      = InterpretAs("ordinal")
	InterpretAsDigits       = InterpretAs("digits")
	InterpretAsUnit         = InterpretAs("unit")
	InterpretAsDate         = InterpretAs("date")
	InterpretAsTime         = InterpretAs("time")
	InterpretAsTelephone    = InterpretAs("telephone")
	InterpretAsAddress      = InterpretAs("address")
	InterpretAsInterjection = InterpretAs("interjection")
	InterpretAsExpletive    = InterpretAs("expletive")
)

type DateFormat string

const (
	DateFormatMonthDayYear = DateFormat("mdy")
	DateFormatDayMonthYear = DateFormat("dmy")
	DateFormatYearMonthDay = DateFormat("ymd")
	DateFormatMonthDay     = DateFormat("md")
	DateFormatDayMonth     = DateFormat("dm")
	DateFormatYearMonth    = DateFormat("ym")
	DateFormatMonthYear    = DateFormat("my")
	DateFormatDay          = DateFormat("d")
	DateFormatMonth        = DateFormat("m")
	DateFormatYear         = DateFormat("y")
	DateFormatDefault      = DateFormat("ymd")
)

type Alphabet string

const (
	AlphabetInternational = Alphabet("ipa")
	AlphabetExtendedSAMPA = Alphabet("x-sampa")
)

type EmphasisLevel string

const (
	// Increase the volume and slow down the speaking rate so the speech is louder and slower.
	EmphasisLevelStrong = EmphasisLevel("strong")
	// Increase the volume and slow down the speaking rate, but not as much as when set to strong. This is used as a default if level is not provided.
	EmphasisLevelModerate = EmphasisLevel("moderate")
	// Decrease the volume and speed up the speaking rate. The speech is softer and faster.
	EmphasisLevelReduced = EmphasisLevel("reduced")
)

type VoiceLanguage string

const (
	VoiceLanguage_en_US = VoiceLanguage("en-US")
	VoiceLanguage_en_GB = VoiceLanguage("en-GB")
	VoiceLanguage_en_IN = VoiceLanguage("en-IN")
	VoiceLanguage_en_AU = VoiceLanguage("en-AU")
	VoiceLanguage_en_CA = VoiceLanguage("en-CA")
	VoiceLanguage_de_DE = VoiceLanguage("de-DE")
	VoiceLanguage_es_ES = VoiceLanguage("es-ES")
	VoiceLanguage_it_IT = VoiceLanguage("it-IT")
	VoiceLanguage_ja_JP = VoiceLanguage("ja-JP")
	VoiceLanguage_fr_FR = VoiceLanguage("fr-FR")
)

type Volume string

const (
	VolumeSilent     = "silent"
	VolumeExtraSoft  = "x-soft"
	VolumeSoft       = "soft"
	VolumeMedium     = "medium"
	VolumeLoud       = "loud"
	VolumeExtraLound = "x-loud"
)

// You can  also use a custom value as a percent
// n% (20%, 275%, 150%). You cannot specify lower than 20%
type Rate string

const (
	RateExtraSlow = "x-slow"
	RateSlow      = "slow"
	RateMedium    = "medium"
	RateFast      = "fast"
	RateExtraFast = "x-fast"
)

type Pitch string

const (
	PitchExtraLow  = "x-low"
	PitchLow       = "low"
	PitchMedium    = "medium"
	PitchHigh      = "high"
	PitchExtraHigh = "x-high"
)
