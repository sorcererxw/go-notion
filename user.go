package notion

// UserType is type of User.
type UserType string

const (
	// UserPerson means User objects represent people.
	UserPerson UserType = "person"
	// UserBot means User objects represent bot.
	UserBot UserType = "bot"
)

// The User object represents a user in a Notion workspace.
type User struct {
	// Always "user"
	Object ObjectType `json:"object,omitempty"`
	ID     string     `json:"id,omitempty"`
	Type   UserType   `json:"type,omitempty"`
	// Properties only present for non-bot users.
	Person *struct {
		Email string `json:"email,omitempty"`
	} `json:"person,omitempty"`
	// Properties only present for bot users.
	Bot       *struct{} `json:"bot,omitempty"`
	Name      string    `json:"name,omitempty"`
	AvatarURL string    `json:"avatar_url,omitempty"`
}
