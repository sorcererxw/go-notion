package notion

// RichTextType is type of RichText.
type RichTextType string

// RichTextType enums.
const (
	RichTextText     RichTextType = "text"
	RichTextMention  RichTextType = "mention"
	RichTextEquation RichTextType = "equation"
)

// RichText objects contain data for displaying formatted text, mentions, and equations.
// RichText object also contains annotations for style information.
// Arrays of rich text objects are used within property objects
// and property value objects to create what a user sees as a single text value in Notion.
type RichText struct {
	// The plain text without annotations.
	PlainText string `json:"plain_text,omitempty"`
	// The URL of any link or internal Notion mention in this text, if any.
	Href string `json:"href,omitempty"`
	// All annotations that apply to this rich text.
	// Annotations include colors and bold/italics/underline/strikethrough.
	Annotations Annotation `json:"annotations"`
	// Type of this rich text object.
	Type     RichTextType `json:"type,omitempty"`
	Text     *Text        `json:"text,omitempty"`
	Mention  *Mention     `json:"mention,omitempty"`
	Equation *Equation    `json:"equation,omitempty"`
}

// Annotation is style information which applies to the whole rich text object.
type Annotation struct {
	// Whether the text is bolded.
	Bold bool `json:"bold,omitempty"`
	// Whether the text is italicized.
	Italic bool `json:"italic,omitempty"`
	// Whether the text is struck through.
	Strikethrough bool `json:"strikethrough,omitempty"`
	// Whether the text is underlined.
	Underline bool `json:"underline,omitempty"`
	// Whether the text is code style.
	Code bool `json:"code,omitempty"`
	// Color of the text.
	Color Color `json:"color,omitempty"`
}

// Color is color definition of Notion.
type Color string

// Color enums.
const (
	ColorDefault          Color = "default"
	ColorGray             Color = "gray"
	ColorBrown            Color = "brown"
	ColorOrange           Color = "orange"
	ColorYellow           Color = "yellow"
	ColorGreen            Color = "green"
	ColorBlue             Color = "blue"
	ColorPurple           Color = "purple"
	ColorPink             Color = "pink"
	ColorRed              Color = "red"
	ColorGrayBackground   Color = "gray_background"
	ColorBrownBackground  Color = "brown_background"
	ColorOrangeBackground Color = "orange_background"
	ColorYellowBackground Color = "yellow_background"
	ColorGreenBackground  Color = "green_background"
	ColorBlueBackground   Color = "blue_background"
	ColorPurpleBackground Color = "purple_background"
	ColorPinkBackground   Color = "pink_background"
	ColorRedBackground    Color = "red_background"
)

// Text is the content of rich text.
type Text struct {
	// Text content. This field contains the actual content of your text and is probably the field you'll use most often.
	Content string `json:"content,omitempty"`
	// Any inline link in this text.
	Link *Link `json:"link,omitempty"`
}

// Link objects contain a type key whose value is always "url" and a url key whose value is a web address.
type Link struct {
	URL string `json:"url,omitempty"`
	// Type is always be "url".
	Type string `json:"type,omitempty"`
}

// MentionType is type of Mention.
type MentionType string

// MentionType enums.
const (
	MentionUser     MentionType = "user"
	MentionPage     MentionType = "page"
	MentionDatabase MentionType = "database"
	MentionDate     MentionType = "date"
)

// Mention objects represent an inline mention of a user, page, database, or date.
// In the app these are created by typing @ followed by the name of a user, page, database, or a date.
//
// Mention objects contain a type key. In addition, mention objects contain a key corresponding with the value of type.
// The value is an object containing type-specific configuration. The type-specific configurations are described in the sections below.
type Mention struct {
	// Type of the inline mention.
	Type MentionType `json:"type,omitempty"`
	// User mentions contain a user object within the user property.
	User *User `json:"user,omitempty"`
	// Page mentions contain a page reference within the page property.
	// A page reference is an object with an id property, with a string value (UUIDv4) corresponding to a page ID.
	Page *ObjectReference `json:"page,omitempty"`
	// Database mentions contain a database reference within the database property.
	// A database reference is an object with an id property, with a string value (UUIDv4) corresponding to a database ID.
	Database *ObjectReference `json:"database,omitempty"`
	// Date mentions contain a date property value object within the date property.
	Date *Date `json:"date,omitempty"`
}

// Equation .
type Equation struct {
	// Expression The LaTeX string representing this inline equation.
	Expression string `json:"expression,omitempty"`
}
