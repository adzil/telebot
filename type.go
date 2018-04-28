package telebot

// ParseMode represents the parsing mode of a text or caption.
type ParseMode string

const (
	// Markdown represents markdown parsing mode.
	Markdown ParseMode = "Markdown"

	// HTML represents HTML parsing mode.
	HTML ParseMode = "HTML"
)

// UpdateType represents the selected update event for updates.
type UpdateType string

const (
	// MessageUpdate represents the message update type.
	MessageUpdate UpdateType = "message"

	// EditedMessageUpdate represents the edited message update type.
	EditedMessageUpdate UpdateType = "edited_message"

	// ChannelPostUpdate represents the channel post update type.
	ChannelPostUpdate UpdateType = "channel_post"

	// EditedChannelPostUpdate represents the edited channel post update type.
	EditedChannelPostUpdate UpdateType = "edited_channel_post"

	// InlineQueryUpdate represents the inline query update type.
	InlineQueryUpdate UpdateType = "inline_query"

	// ChosenInlineResultUpdate represents the chosen inline result update type.
	ChosenInlineResultUpdate UpdateType = "chosen_inline_result"

	// CallbackQueryUpdate represents the callback query update type.
	CallbackQueryUpdate UpdateType = "callback_query"
)

// Update represents an incoming update. At most one of the optional parameters
// can be present in any given update.
type Update struct {
	ID                 int64               `json:"update_id"`
	Message            *Message            `json:"message,omitempty"`
	EditedMessage      *Message            `json:"edited_message,omitempty"`
	ChannelPost        *Message            `json:"channel_post,omitempty"`
	EditedChannelPost  *Message            `json:"edited_channel_post,omitempty"`
	InlineQuery        *InlineQuery        `json:"inline_query,omitempty"`
	ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result,omitempty"`
	CallbackQuery      *CallbackQuery      `json:"callback_query,omitempty"`
}

// Type gets the update type.
func (u *Update) Type() UpdateType {
	if u.Message != nil {
		return MessageUpdate
	} else if u.EditedMessage != nil {
		return EditedMessageUpdate
	} else if u.ChannelPost != nil {
		return ChannelPostUpdate
	} else if u.EditedChannelPost != nil {
		return EditedChannelPostUpdate
	} else if u.InlineQuery != nil {
		return InlineQueryUpdate
	} else if u.ChosenInlineResult != nil {
		return ChosenInlineResultUpdate
	} else if u.CallbackQuery != nil {
		return CallbackQueryUpdate
	}
	// Cannot determine the update type
	return ""
}

// WebhookInfo contains information about the current status of a webhook.
type WebhookInfo struct {
	URL                  string   `json:"url"`
	HasCustomCertificate bool     `json:"has_custom_certificate"`
	PendingUpdateCount   int      `json:"pending_update_count"`
	LastErrorDate        int64    `json:"last_error_date,omitempty"`
	LastErrorMessage     string   `json:"last_error_message,omitempty"`
	MaxConnections       int      `json:"max_connections,omitempty"`
	AllowedUpdates       []string `json:"allowed_updates,omitempty"`
}

// User represents a Telegram user or bot.
type User struct {
	ID           int64  `json:"id"`
	IsBot        bool   `json:"is_bot"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name,omitempty"`
	Username     string `json:"username,omitempty"`
	LanguageCode string `json:"language_code,omitempty"`
}

// ChatType represents the chat type.
type ChatType string

const (
	// PrivateChat represents private chat type.
	PrivateChat ChatType = "private"

	// GroupChat represents group chat type.
	GroupChat ChatType = "group"

	// SupergroupChat represents supergroup chat type.
	SupergroupChat ChatType = "supergroup"

	// ChannelChat represents channel chat type.
	ChannelChat ChatType = "channel"
)

// Chat represents a chat.
type Chat struct {
	ID                          int64    `json:"id"`
	Type                        ChatType `json:"type"`
	Title                       string   `json:"title,omitempty"`
	Username                    string   `json:"username,omitempty"`
	FirstName                   string   `json:"first_name,omitempty"`
	LastName                    string   `json:"last_name,omitempty"`
	AllMembersAreAdministrators bool     `json:"all_members_are_administrators,omitempty"`
	Description                 string   `json:"description,omitempty"`
	InviteLink                  string   `json:"invite_link,omitempty"`
	PinnedMessage               *Message `json:"pinned_message,omitempty"`
}

// Message represents a message.
type Message struct {
	ID                   int64            `json:"message_id"`
	From                 *User            `json:"from,omitempty"`
	Date                 int64            `json:"date"`
	Chat                 *Chat            `json:"chat"`
	ForwardFrom          *User            `json:"forward_from,omitempty"`
	ForwardFromChat      *Chat            `json:"forward_from_chat,omitempty"`
	ForwardFromMessageID int64            `json:"forward_from_message_id,omitempty"`
	ForwardSignature     string           `json:"forward_signature,omitempty"`
	ForwardDate          int64            `json:"forward_date,omitempty"`
	ReplyToMessage       *Message         `json:"reply_to_message,omitempty"`
	EditDate             int64            `json:"edit_date,omitempty"`
	MediaGroupID         string           `json:"media_group_id,omitempty"`
	AuthorSignature      string           `json:"author_signature,omitempty"`
	Text                 string           `json:"text,omitempty"`
	Entities             []*MessageEntity `json:"entities,omitempty"`
	CaptionEntities      []*MessageEntity `json:"caption_entities,omitempty"`
	Caption              string           `json:"caption,omitempty"`
	Contact              *Contact         `json:"contact,omitempty"`
	Location             *Location        `json:"location,omitempty"`
	Venue                *Venue           `json:"venue,omitempty"`
	NewChatMembers       []*User          `json:"new_chat_members,omitempty"`
	LeftChatMember       *User            `json:"left_chat_member,omitempty"`
	NewChatTitle         string           `json:"new_chat_title,omitempty"`
	PinnedMessage        *Message         `json:"pinned_message,omitempty"`
	ConnectedWebsite     string           `json:"connected_website,omitempty"`
}

// MessageEntityType represents the message entity type.
type MessageEntityType string

const (
	// MentionEntity represents mention entity type.
	MentionEntity MessageEntityType = "mention"

	// HashtagEntity represents hashtag entity type.
	HashtagEntity MessageEntityType = "hashtag"

	// BotCommandEntity represents bot command entity type.
	BotCommandEntity MessageEntityType = "bot_command"

	// URLEntity represents URL entity type.
	URLEntity MessageEntityType = "url"

	// EmailEntity represents email entity type.
	EmailEntity MessageEntityType = "email"

	// BoldEntity represents bold entity type.
	BoldEntity MessageEntityType = "bold"

	// ItalicEntity represents italic entity type.
	ItalicEntity MessageEntityType = "italic"

	// CodeEntity represents code entity type.
	CodeEntity MessageEntityType = "code"

	// PreEntity represents pre entity type.
	PreEntity MessageEntityType = "pre"

	// TextLinkEntity represents text link entity type.
	TextLinkEntity MessageEntityType = "text_link"

	// TextMention represents text mention entity type.
	TextMention MessageEntityType = "text_mention"
)

// MessageEntity represents one special entity in a text message. For example,
// hashtags, usernames, URLs, etc.
type MessageEntity struct {
	Type   MessageEntityType `json:"type"`
	Offset int               `json:"offset"`
	Length int               `json:"length"`
	URL    string            `json:"url,omitempty"`
	User   *User             `json:"user,omitempty"`
}

// Contact represents a phone contact.
type Contact struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name,omitempty"`
	UserID      int64  `json:"user_id,omitempty"`
}

// Location represents a point on the map.
type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// Venue represents a venue.
type Venue struct {
	Location     *Location `json:"location"`
	Title        string    `json:"title"`
	Address      string    `json:"address"`
	FoursquareID string    `json:"foursquare_id,omitempty"`
}

// ReplyKeyboardMarkup represents a custom keyboard with reply options.
type ReplyKeyboardMarkup struct {
	Keyboard        [][]*KeyboardButton `json:"keyboard"`
	ResizeKeyboard  bool                `json:"resize_keyboard,omitempty"`
	OneTimeKeyboard bool                `json:"one_time_keyboard,omitempty"`
	Selective       bool                `json:"selective,omitempty"`
}

// Type implements the ReplyMarkup interface.
func (r *ReplyKeyboardMarkup) Type() ReplyMarkupType {
	return ReplyKeyboardMarkupType
}

// KeyboardButton represents one button of the reply keyboard.
type KeyboardButton struct {
	Text            string `json:"text"`
	RequestContact  bool   `json:"request_contact,omitempty"`
	RequestLocation bool   `json:"request_location,omitempty"`
}

// NewKeyboard is a helper function to instantiate new keyboard.
func NewKeyboard(rows ...[]*KeyboardButton) [][]*KeyboardButton {
	return rows
}

// NewKeyboardRow is a helper function to instantiate new keyboard row.
func NewKeyboardRow(cols ...*KeyboardButton) []*KeyboardButton {
	return cols
}

// ReplyKeyboardRemove removes the current custom keyboard and display the
// default letter-keyboard. RemoveKeyboard field value is always true.
type ReplyKeyboardRemove struct {
	RemoveKeyboard bool `json:"remove_keyboard"`
	Selective      bool `json:"selective,omitempty"`
}

// Type implements the ReplyMarkup interface.
func (i *ReplyKeyboardRemove) Type() ReplyMarkupType {
	return ReplyKeyboardRemoveType
}

// ReplyMarkupType represents the reply markup type.
type ReplyMarkupType string

const (
	// InlineKeyboardMarkupType represents the inline keyboard markup type.
	InlineKeyboardMarkupType ReplyMarkupType = "inline_keyboard_markup"

	// ReplyKeyboardMarkupType represents the reply keyboard markup type.
	ReplyKeyboardMarkupType ReplyMarkupType = "reply_keyboard_markup"

	// ReplyKeyboardRemoveType represents the reply keyboard remove type.
	ReplyKeyboardRemoveType ReplyMarkupType = "reply_keyboard_remove"

	// ForceReplyType represents the force reply type.
	ForceReplyType ReplyMarkupType = "force_reply"
)

// ReplyMarkup represents additional interface options when sending messages.
type ReplyMarkup interface {
	Type() ReplyMarkupType
}

// InlineKeyboardMarkup represents an inline keyboard that appears right next
// to the message it belongs to.
type InlineKeyboardMarkup struct {
	InlineKeyboard [][]*InlineKeyboardButton `json:"inline_keyboard"`
}

// Type implements the ReplyMarkup interface.
func (i *InlineKeyboardMarkup) Type() ReplyMarkupType {
	return InlineKeyboardMarkupType
}

// InlineKeyboardButton represents one button of an inline keyboard. You must
// use exactly one of the optional fields.
type InlineKeyboardButton struct {
	Text                         string `json:"text"`
	URL                          string `json:"url,omitempty"`
	CallbackData                 string `json:"callback_data,omitempty"`
	SwitchInlineQuery            string `json:"switch_inline_query,omitempty"`
	SwitchInlineQueryCurrentChat string `json:"switch_inline_query_current_chat,omitempty"`
}

// NewInlineKeyboard is a helper function to instantiate new inline keyboard.
func NewInlineKeyboard(rows ...[]*InlineKeyboardButton) [][]*InlineKeyboardButton {
	return rows
}

// NewInlineKeyboardRow is a helper function to instantiate new inline keyboard
// row.
func NewInlineKeyboardRow(cols ...*InlineKeyboardButton) []*InlineKeyboardButton {
	return cols
}

// CallbackQuery represents an incoming callback query from a callback button
// in an inline keyboard. If the button that originated the query was attached
// to a message sent by the bot, the field message will be present. If the
// button was attached to a message sent via the bot (in inline mode), the
// field InlineMessageID will be present.
type CallbackQuery struct {
	ID              string   `json:"id"`
	From            *User    `json:"from"`
	Message         *Message `json:"message,omitempty"`
	InlineMessageID string   `json:"inline_message_id,omitempty"`
	ChatInstance    string   `json:"chat_instance"`
	Data            string   `json:"data,omitempty"`
}

// ForceReply will display a reply interface to the user (act as if the user
// has selected the bot‘s message and tapped 'Reply'). ForceReply field value
// is always true.
type ForceReply struct {
	ForceReply bool `json:"force_reply"`
	Selective  bool `json:"selective,omitempty"`
}

// Type implements the ReplyMarkup interface.
func (f *ForceReply) Type() ReplyMarkupType {
	return ForceReplyType
}

// ChatMemberStatus represents the status of a chat member.
type ChatMemberStatus string

const (
	// CreatorMember represents the creator status.
	CreatorMember ChatMemberStatus = "creator"

	// AdministratorMember represents the administrator status.
	AdministratorMember ChatMemberStatus = "administrator"

	// Member represents the member status.
	Member ChatMemberStatus = "member"

	// RestrictedMember represents the restricted status.
	RestrictedMember ChatMemberStatus = "restricted"

	// LeftMember represents the left status.
	LeftMember ChatMemberStatus = "left"

	// KickedMember represents the kicked status.
	KickedMember ChatMemberStatus = "kicked"
)

// ChatMember contains information about one member of a chat.
type ChatMember struct {
	User                  *User            `json:"user"`
	Status                ChatMemberStatus `json:"status"`
	UntilDate             int64            `json:"until_date,omitempty"`
	CanBeEdited           bool             `json:"can_be_edited,omitempty"`
	CanChangeInfo         bool             `json:"can_change_info,omitempty"`
	CanPostMessages       bool             `json:"can_post_messages,omitempty"`
	CanEditMessages       bool             `json:"can_edit_messages,omitempty"`
	CanDeleteMessages     bool             `json:"can_delete_messages,omitempty"`
	CanInviteUsers        bool             `json:"can_invite_users,omitempty"`
	CanRestrictMembers    bool             `json:"can_restrict_members,omitempty"`
	CanPinMessages        bool             `json:"can_pin_messages,omitempty"`
	CanPromoteMembers     bool             `json:"can_promote_members,omitempty"`
	CanSendMessages       bool             `json:"can_send_messages,omitempty"`
	CanSendMediaMessages  bool             `json:"can_send_media_messages,omitempty"`
	CanSendOtherMessages  bool             `json:"can_send_other_messages,omitempty"`
	CanAddWebPagePreviews bool             `json:"can_add_web_page_previews,omitempty"`
}

// ResponseParameters contains information about why a request was unsuccessful.
type ResponseParameters struct {
	MigrateToChatID int64 `json:"migrate_to_chat_id,omitempty"`
	RetryAfter      int   `json:"retry_after,omitempty"`
}
