/*
Package telebot implements RPC method to Telegram Bot API with convenient high
level polling method for getting updates.
*/
package telebot

import (
	"context"
	"time"
)

const (
	// EndpointURL represents the Telegram Bot API endpoint URL.
	EndpointURL = "https://api.telegram.org/bot"

	// DefaultBackoffPeriod used to hold polling when an error occured.
	DefaultBackoffPeriod = 5 * time.Second

	// DefaultPollTimeout used to set the polling timeout in seconds.
	DefaultPollTimeout = 60
)

// Bot implements the Telegram Bot API interface.
type Bot struct {
	Self          *User
	BackoffPeriod time.Duration
	caller        *Caller
}

// GetUpdates sets parameter for GetUpdates and PollUpdates method.
type GetUpdates struct {
	Offset         int64        `json:"offset,omitempty"`
	Limit          int          `json:"limit,omitempty"`
	Timeout        int          `json:"timeout,omitempty"`
	AllowedUpdates []UpdateType `json:"allowed_updates,omitempty"`
}

// GetUpdates receive incoming updates using polling request. To continuously
// poll updates, use the PollUpdates method instead.
func (b *Bot) GetUpdates(req *GetUpdates) ([]*Update, error) {
	var updates []*Update
	err := b.caller.Poll("getUpdates", req, &updates)
	return updates, err
}

// PollUpdates continously polls for updates with channel. Do not reuse request
// after this call because it internally updates the offset number to fetch the
// next updates. If timeout not set, it will automatically set to 60 seconds to
// prevent short polling. To stop the polling, pass a cancellation context into
// ctx. Otherwise, use nil or background context.
func (b *Bot) PollUpdates(ctx context.Context, req *GetUpdates) (chan *Update, chan error) {
	// Fill context with background if nil
	if ctx == nil {
		ctx = context.Background()
	}
	// Increase timeout value if unset to prevent short polling
	if req.Timeout <= 0 {
		req.Timeout = DefaultPollTimeout
	}
	// Create channel for update and error
	retupd, reterr := make(chan *Update), make(chan error)
	// Create goroutine to enable asynchronous update mechanism
	go func() {
		// Get done channel
		done := ctx.Done()
		// Loop until cancellation
		for {
			select {
			case <-done:
				// Close update and error channel
				close(retupd)
				close(reterr)
				// Exit from goroutine
				return
			default:
				// Get updates
				upds, err := b.GetUpdates(req)
				// Check for error
				if err != nil {
					// Return error to channel and sleep for backoff period
					reterr <- err
					time.Sleep(b.BackoffPeriod)
					continue
				}
				// Iterate over updates and return it to update channel
				for _, upd := range upds {
					retupd <- upd
					// Change offset value and increment it by one
					if req.Offset <= upd.ID {
						req.Offset = upd.ID + 1
					}
				}
			}
		}
	}()
	// Return channels
	return retupd, reterr
}

// SetWebhook sets parameter for SetWebhook method.
type SetWebhook struct {
	URL            string       `json:"url"`
	MaxConnections int          `json:"max_connections,omitempty"`
	AllowedUpdates []UpdateType `json:"allowed_updates,omitempty"`
}

// SetWebhook tells API to receive incoming updates via an outgoing webhook.
// Whenever there is an update for the bot, we will send an HTTPS POST request
// to the specified url, containing a JSON-serialized Update object. For
// automated webhook setup with callback handler, use StartWebhook method
// instead.
func (b *Bot) SetWebhook(req *SetWebhook) error {
	return b.caller.Call("setWebhook", req, nil)
}

// DeleteWebhook removes webhook integration if you decide to switch back to
// GetUpdates or PollUpdates method. Returns True on success.
func (b *Bot) DeleteWebhook() (bool, error) {
	var ok bool
	err := b.caller.Call("deleteWebhook", nil, &ok)
	return ok, err
}

// GetWebhookInfo gets current webhook status.
func (b *Bot) GetWebhookInfo() (*WebhookInfo, error) {
	var info WebhookInfo
	err := b.caller.Call("getWebhookInfo", nil, &info)
	return &info, err
}

// GetMe returns basic information about the bot in form of a User object.
func (b *Bot) GetMe() (*User, error) {
	var me User
	err := b.caller.Call("getMe", nil, &me)
	return &me, err
}

// SendRequestType represents the send request type enumerator.
type SendRequestType string

const (
	// SendMessageRequest represents the send message request type.
	SendMessageRequest SendRequestType = "sendMessage"

	// ForwardMessageRequest represents the forward message request type.
	ForwardMessageRequest SendRequestType = "forwardMessage"

	// SendLocationRequest represents the send location request type.
	SendLocationRequest SendRequestType = "sendLocation"

	// EditMessageLiveLocationRequest represents the edit live location request
	// type.
	EditMessageLiveLocationRequest SendRequestType = "editMessageLiveLocation"

	// StopMessageLiveLocationRequest represents the stop live location request
	// type.
	StopMessageLiveLocationRequest SendRequestType = "stopMessageLiveLocation"

	// SendVenueRequest represents the send venue request type.
	SendVenueRequest SendRequestType = "sendVenue"

	// SendContactRequest represents the send contact request type.
	SendContactRequest SendRequestType = "sendContact"

	// EditMessageTextRequest represents the edit message text request type.
	EditMessageTextRequest SendRequestType = "editMessageText"

	// EditMessageCaptionRequest represents the edit message caption request
	// type.
	EditMessageCaptionRequest SendRequestType = "editMessageCaption"

	// EditMessageReplyMarkupRequest represents the edit message reply markup
	// request type.
	EditMessageReplyMarkupRequest SendRequestType = "editMessageReplyMarkup"
)

// SendRequest represents generic send request that can be distinguished by its
// type.
type SendRequest interface {
	Type() SendRequestType
}

// Send processes send request into a proper API call with type detection.
func (b *Bot) Send(req SendRequest) (*Message, error) {
	var msg Message
	err := b.caller.Call(string(req.Type()), req, &msg)
	return &msg, err
}

// SendMessage send text messages.
type SendMessage struct {
	ChatID                int64       `json:"chat_id"`
	Text                  string      `json:"text"`
	ParseMode             ParseMode   `json:"parse_mode,omitempty"`
	DisableWebPagePreview bool        `json:"disable_web_page_preview,omitempty"`
	DisableNotification   bool        `json:"disable_notification,omitempty"`
	ReplyToMessageID      int64       `json:"reply_to_message_id,omitempty"`
	ReplyMarkup           ReplyMarkup `json:"reply_markup,omitempty"`
}

// Type implements the SendRequest interface.
func (m *SendMessage) Type() SendRequestType {
	return SendMessageRequest
}

// ForwardMessage forward messages of any kind.
type ForwardMessage struct {
	ChatID              int64 `json:"chat_id"`
	FromChatID          int64 `json:"from_chat_id"`
	DisableNotification bool  `json:"disable_notification,omitempty"`
	MessageID           int64 `json:"message_id"`
}

// Type implements the SendRequest interface.
func (m *ForwardMessage) Type() SendRequestType {
	return ForwardMessageRequest
}

// SendLocation send point on the map.
type SendLocation struct {
	ChatID              int64       `json:"chat_id"`
	Latitude            float64     `json:"latitude"`
	Longitude           float64     `json:"longitude"`
	LivePeriod          int         `json:"live_period,omitempty"`
	DisableNotification bool        `json:"disable_notification,omitempty"`
	ReplyToMessageID    int64       `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         ReplyMarkup `json:"reply_markup,omitempty"`
}

// Type implements the SendRequest interface.
func (m *SendLocation) Type() SendRequestType {
	return SendLocationRequest
}

// EditMessageLiveLocation edit live location messages sent by the bot. A
// location can be edited until its live_period expires or editing is explicitly
// disabled by a call to StopMessageLiveLocation.
type EditMessageLiveLocation struct {
	ChatID          int64                 `json:"chat_id,omitempty"`
	MessageID       int64                 `json:"message_id,omitempty"`
	InlineMessageID string                `json:"inline_message_id,omitempty"`
	Latitude        float64               `json:"latitude"`
	Longitude       float64               `json:"longitude"`
	ReplyMarkup     *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// Type implements the SendRequest interface.
func (m *EditMessageLiveLocation) Type() SendRequestType {
	return EditMessageLiveLocationRequest
}

// StopMessageLiveLocation stop updating a live location message sent by the bot
// before LivePeriod expires.
type StopMessageLiveLocation struct {
	ChatID          int64                 `json:"chat_id,omitempty"`
	MessageID       int64                 `json:"message_id,omitempty"`
	InlineMessageID string                `json:"inline_message_id,omitempty"`
	ReplyMarkup     *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// Type implements the SendRequest interface.
func (m *StopMessageLiveLocation) Type() SendRequestType {
	return StopMessageLiveLocationRequest
}

// SendVenue send information about a venue.
type SendVenue struct {
	ChatID              int64       `json:"chat_id"`
	Latitude            float64     `json:"latitude"`
	Longitude           float64     `json:"longitude"`
	Title               string      `json:"title"`
	Address             string      `json:"address"`
	FoursquareID        string      `json:"foursquare_id,omitempty"`
	DisableNotification bool        `json:"disable_notification,omitempty"`
	ReplyToMessageID    int64       `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         ReplyMarkup `json:"reply_markup,omitempty"`
}

// Type implements the SendRequest interface.
func (m *SendVenue) Type() SendRequestType {
	return SendVenueRequest
}

// SendContact send phone contacts.
type SendContact struct {
	ChatID              int64       `json:"chat_id"`
	PhoneNumber         string      `json:"phone_number"`
	FirstName           string      `json:"first_name"`
	LastName            string      `json:"last_name,omitempty"`
	DisableNotification bool        `json:"disable_notification,omitempty"`
	ReplyToMessageID    int64       `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         ReplyMarkup `json:"reply_markup,omitempty"`
}

// Type implements the SendRequest interface.
func (m *SendContact) Type() SendRequestType {
	return SendContactRequest
}

// EditMessageText edit text messages sent by the bot.
type EditMessageText struct {
	ChatID                int64       `json:"chat_id,omitempty"`
	MessageID             int64       `json:"message_id,omitempty"`
	InlineMessageID       string      `json:"inline_message_id,omitempty"`
	Text                  string      `json:"text"`
	ParseMode             ParseMode   `json:"parse_mode,omitempty"`
	DisableWebPagePreview bool        `json:"disable_web_page_preview,omitempty"`
	ReplyMarkup           ReplyMarkup `json:"reply_markup,omitempty"`
}

// Type implements the SendRequest interface.
func (m *EditMessageText) Type() SendRequestType {
	return EditMessageTextRequest
}

// EditMessageCaption edit captions of messages sent by the bot.
type EditMessageCaption struct {
	ChatID          int64       `json:"chat_id,omitempty"`
	MessageID       int64       `json:"message_id,omitempty"`
	InlineMessageID string      `json:"inline_message_id,omitempty"`
	Caption         string      `json:"caption"`
	ParseMode       ParseMode   `json:"parse_mode,omitempty"`
	ReplyMarkup     ReplyMarkup `json:"reply_markup,omitempty"`
}

// Type implements the SendRequest interface.
func (m *EditMessageCaption) Type() SendRequestType {
	return EditMessageCaptionRequest
}

// EditMessageReplyMarkup edit only the reply markup of messages sent by the
// bot.
type EditMessageReplyMarkup struct {
	ChatID          int64       `json:"chat_id,omitempty"`
	MessageID       int64       `json:"message_id,omitempty"`
	InlineMessageID string      `json:"inline_message_id,omitempty"`
	ReplyMarkup     ReplyMarkup `json:"reply_markup,omitempty"`
}

// Type implements the SendRequest interface.
func (m *EditMessageReplyMarkup) Type() SendRequestType {
	return EditMessageReplyMarkupRequest
}

// AnswerCallbackQuery sets parameter for AnswerCallbackQuery method.
type AnswerCallbackQuery struct {
	CallbackQueryID string `json:"callback_query_id"`
	Text            string `json:"text,omitempty"`
	ShowAlert       bool   `json:"show_alert,omitempty"`
	URL             string `json:"url,omitempty"`
	CacheTime       int    `json:"cache_time,omitempty"`
}

// AnswerCallbackQuery send answers to callback queries sent from inline
// keyboards. The answer will be displayed to the user as a notification at the
// top of the chat screen or as an alert.
func (b *Bot) AnswerCallbackQuery(req *AnswerCallbackQuery) (bool, error) {
	var ok bool
	err := b.caller.Call("answerCallbackQuery", req, &ok)
	return ok, err
}

// DeleteMessage sets parameter for DeleteMessage method.
type DeleteMessage struct {
	ChatID    int64 `json:"chat_id"`
	MessageID int64 `json:"message_id"`
}

// DeleteMessage delete a message.
func (b *Bot) DeleteMessage(req *DeleteMessage) (bool, error) {
	var ok bool
	err := b.caller.Call("deleteMessage", req, &ok)
	return ok, err
}

// AnswerInlineQuery sets parameter for AnswerInlineQuery method.
type AnswerInlineQuery struct {
	InlineQueryID     string              `json:"inline_query_id"`
	Results           []InlineQueryResult `json:"results"`
	CacheTime         int                 `json:"cache_time,omitempty"`
	IsPersonal        bool                `json:"is_personal,omitempty"`
	NextOffset        string              `json:"next_offset,omitempty"`
	SwitchPmText      string              `json:"switch_pm_text,omitempty"`
	SwitchPmParameter string              `json:"switch_pm_parameter,omitempty"`
}

// AnswerInlineQuery send answers to an inline query.
func (b *Bot) AnswerInlineQuery(req *AnswerInlineQuery) (bool, error) {
	var ok bool
	err := b.caller.Call("answerInlineQuery", req, &ok)
	return ok, err
}

// NewBot create new Bot from access token.
func NewBot(token string) (*Bot, error) {
	// Create new bot instance
	bot := &Bot{
		BackoffPeriod: DefaultBackoffPeriod,
		caller:        NewCaller(EndpointURL, token),
	}
	// Test connecton with getMe method
	var err error
	bot.Self, err = bot.GetMe()
	return bot, err
}
