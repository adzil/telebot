package telebot

import "encoding/json"

// InlineQuery represents an incoming inline query. When the user sends an empty
// query, your bot could return some default or trending results.
type InlineQuery struct {
	ID       string    `json:"id"`
	From     *User     `json:"from"`
	Location *Location `json:"location,omitempty"`
	Query    string    `json:"query"`
	Offset   string    `json:"offset"`
}

// InlineQueryResultType represents the inline query result type.
type InlineQueryResultType string

const (
	// ArticleResult represents the article inline query result type.
	ArticleResult InlineQueryResultType = "article"

	// PhotoResult represents the photo inline query result type.
	PhotoResult InlineQueryResultType = "photo"

	// GifResult represents the gif inline query result type.
	GifResult InlineQueryResultType = "gif"

	// Mpeg4GifResult represents the mpeg4 gif inline query result type.
	Mpeg4GifResult InlineQueryResultType = "mpeg4_gif"

	// VideoResult represents the video inline query result type.
	VideoResult InlineQueryResultType = "video"

	// AudioResult represents the audio inline query result type.
	AudioResult InlineQueryResultType = "audio"

	// VoiceResult represents the voice inline query result type.
	VoiceResult InlineQueryResultType = "voice"

	// DocumentResult represents the document inline query result type.
	DocumentResult InlineQueryResultType = "document"

	// LocationResult represents the location inline query result type.
	LocationResult InlineQueryResultType = "location"

	// VenueResult represents the venue inline query result type.
	VenueResult InlineQueryResultType = "venue"

	// ContactResult represents the contact inline query result type.
	ContactResult InlineQueryResultType = "contact"
)

// InlineQueryResult represents one result of an inline query.
type InlineQueryResult interface {
	Type() InlineQueryResultType
}

// InlineQueryResultArticle represents a link to an article or web page.
type InlineQueryResultArticle struct {
	ID                  string                `json:"id"`
	Title               string                `json:"title"`
	InputMessageContent InputMessageContent   `json:"input_message_content"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	URL                 string                `json:"url,omitempty"`
	HideURL             bool                  `json:"hide_url,omitempty"`
	Description         string                `json:"description,omitempty"`
	ThumbURL            string                `json:"thumb_url,omitempty"`
	ThumbWidth          int                   `json:"thumb_width,omitempty"`
	ThumbHeight         int                   `json:"thumb_height,omitempty"`
}

// Type implements InlineQueryResult interface.
func (r *InlineQueryResultArticle) Type() InlineQueryResultType {
	return ArticleResult
}

type inlineQueryResultArticleBase InlineQueryResultArticle

type inlineQueryResultArticle struct {
	Type InlineQueryResultType `json:"type"`
	*inlineQueryResultArticleBase
}

// MarshalJSON implements json.Marshaler interface.
func (r *InlineQueryResultArticle) MarshalJSON() ([]byte, error) {
	return json.Marshal(&inlineQueryResultArticle{r.Type(), (*inlineQueryResultArticleBase)(r)})
}

// InlineQueryResultPhoto represents a link to a photo. By default, this photo
// will be sent by the user with optional caption. Alternatively, you can use
// InputMessageContent to send a message with the specified content instead of the photo.
type InlineQueryResultPhoto struct {
	ID                  string                `json:"id"`
	URL                 string                `json:"photo_url"`
	ThumbURL            string                `json:"thumb_url"`
	Width               int                   `json:"photo_width,omitempty"`
	Height              int                   `json:"photo_height,omitempty"`
	Title               string                `json:"title,omitempty"`
	Description         string                `json:"description,omitempty"`
	Caption             string                `json:"caption,omitempty"`
	ParseMode           ParseMode             `json:"parse_mode,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
}

// Type implements InlineQueryResult interface.
func (r *InlineQueryResultPhoto) Type() InlineQueryResultType {
	return PhotoResult
}

type inlineQueryResultPhotoBase InlineQueryResultPhoto

type inlineQueryResultPhoto struct {
	Type InlineQueryResultType `json:"type"`
	*inlineQueryResultPhotoBase
}

// MarshalJSON implements json.Marshaler interface.
func (r *InlineQueryResultPhoto) MarshalJSON() ([]byte, error) {
	return json.Marshal(&inlineQueryResultPhoto{r.Type(), (*inlineQueryResultPhotoBase)(r)})
}

// InlineQueryResultGif represents a link to an animated GIF file. By default,
// this animated GIF file will be sent by the user with optional caption.
// Alternatively, you can use InputMessageContent to send a message with the
// specified content instead of the animation.
type InlineQueryResultGif struct {
	ID                  string                `json:"id"`
	URL                 string                `json:"gif_url"`
	Width               int                   `json:"gif_width,omitempty"`
	Height              int                   `json:"gif_height,omitempty"`
	Duration            int                   `json:"gif_duration,omitempty"`
	ThumbURL            string                `json:"thumb_url,omitempty"`
	Title               string                `json:"title,omitempty"`
	Caption             string                `json:"caption,omitempty"`
	ParseMode           ParseMode             `json:"parse_mode,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
}

// Type implements InlineQueryResult interface.
func (r *InlineQueryResultGif) Type() InlineQueryResultType {
	return GifResult
}

type inlineQueryResultGifBase InlineQueryResultGif

type inlineQueryResultGif struct {
	Type InlineQueryResultType `json:"type"`
	*inlineQueryResultGifBase
}

// MarshalJSON implements json.Marshaler interface.
func (r *InlineQueryResultGif) MarshalJSON() ([]byte, error) {
	return json.Marshal(&inlineQueryResultGif{r.Type(), (*inlineQueryResultGifBase)(r)})
}

// InlineQueryResultMpeg4Gif represents a link to a video animation
// (H.264/MPEG-4 AVC video without sound). By default, this animated MPEG-4 file
// will be sent by the user with optional caption. Alternatively, you can use
// InputMessageContent to send a message with the specified content instead of
// the animation.
type InlineQueryResultMpeg4Gif struct {
	ID                  string                `json:"id"`
	URL                 string                `json:"mpeg4_url"`
	Width               int                   `json:"mpeg4_width,omitempty"`
	Height              int                   `json:"mpeg4_height,omitempty"`
	Duration            int                   `json:"mpeg4_duration,omitempty"`
	ThumbURL            string                `json:"thumb_url,omitempty"`
	Title               string                `json:"title,omitempty"`
	Caption             string                `json:"caption,omitempty"`
	ParseMode           ParseMode             `json:"parse_mode,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
}

// Type implements InlineQueryResult interface.
func (r *InlineQueryResultMpeg4Gif) Type() InlineQueryResultType {
	return Mpeg4GifResult
}

type inlineQueryResultMpeg4GifBase InlineQueryResultMpeg4Gif

type inlineQueryResultMpeg4Gif struct {
	Type InlineQueryResultType `json:"type"`
	*inlineQueryResultMpeg4GifBase
}

// MarshalJSON implements json.Marshaler interface.
func (r *InlineQueryResultMpeg4Gif) MarshalJSON() ([]byte, error) {
	return json.Marshal(&inlineQueryResultMpeg4Gif{r.Type(), (*inlineQueryResultMpeg4GifBase)(r)})
}

// VideoMimeType represents the video MIME type.
type VideoMimeType string

const (
	// HTMLVideo represents HTML video MIME type.
	HTMLVideo VideoMimeType = "text/html"

	// Mp4Video represents MP4 video MIME type.
	Mp4Video VideoMimeType = "video/mp4"
)

// InlineQueryResultVideo represents a link to a page containing an embedded
// video player or a video file. By default, this video file will be sent by the
// user with an optional caption. Alternatively, you can use InputMessageContent
// to send a message with the specified content instead of the video.
type InlineQueryResultVideo struct {
	ID                  string                `json:"id"`
	URL                 string                `json:"video_url"`
	MimeType            VideoMimeType         `json:"mime_type"`
	ThumbURL            string                `json:"thumb_url"`
	Title               string                `json:"title"`
	Caption             string                `json:"caption,omitempty"`
	ParseMode           ParseMode             `json:"parse_mode,omitempty"`
	Width               int                   `json:"video_width,omitempty"`
	Height              int                   `json:"video_height,omitempty"`
	Duration            int                   `json:"video_duration,omitempty"`
	Description         string                `json:"description,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
}

// Type implements InlineQueryResult interface.
func (r *InlineQueryResultVideo) Type() InlineQueryResultType {
	return VideoResult
}

type inlineQueryResultVideoBase InlineQueryResultVideo

type inlineQueryResultVideo struct {
	Type InlineQueryResultType `json:"type"`
	*inlineQueryResultVideoBase
}

// MarshalJSON implements json.Marshaler interface.
func (r *InlineQueryResultVideo) MarshalJSON() ([]byte, error) {
	return json.Marshal(&inlineQueryResultVideo{r.Type(), (*inlineQueryResultVideoBase)(r)})
}

// InlineQueryResultAudio represents a link to an mp3 audio file. By default,
// this audio file will be sent by the user. Alternatively, you can use
// InputMessageContent to send a message with the specified content instead of
// the audio.
type InlineQueryResultAudio struct {
	ID                  string                `json:"id"`
	URL                 string                `json:"audio_url"`
	Title               string                `json:"title"`
	Caption             string                `json:"caption,omitempty"`
	ParseMode           ParseMode             `json:"parse_mode,omitempty"`
	Performer           string                `json:"performer,omitempty"`
	Duration            int                   `json:"audio_duration,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
}

// Type implements InlineQueryResult interface.
func (r *InlineQueryResultAudio) Type() InlineQueryResultType {
	return VoiceResult
}

type inlineQueryResultAudioBase InlineQueryResultAudio

type inlineQueryResultAudio struct {
	Type InlineQueryResultType `json:"type"`
	*inlineQueryResultAudioBase
}

// MarshalJSON implements json.Marshaler interface.
func (r *InlineQueryResultAudio) MarshalJSON() ([]byte, error) {
	return json.Marshal(&inlineQueryResultAudio{r.Type(), (*inlineQueryResultAudioBase)(r)})
}

// InlineQueryResultVoice represents a link to a voice recording in an .ogg
// container encoded with OPUS. By default, this voice recording will be sent by
// the user. Alternatively, you can use InputMessageContent to send a message
// with the specified content instead of the the voice message.
type InlineQueryResultVoice struct {
	ID                  string                `json:"id"`
	URL                 string                `json:"voice_url"`
	Title               string                `json:"title"`
	Caption             string                `json:"caption,omitempty"`
	ParseMode           ParseMode             `json:"parse_mode,omitempty"`
	Duration            int                   `json:"voice_duration,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
}

// Type implements InlineQueryResult interface.
func (r *InlineQueryResultVoice) Type() InlineQueryResultType {
	return VoiceResult
}

type inlineQueryResultVoiceBase InlineQueryResultVoice

type inlineQueryResultVoice struct {
	Type InlineQueryResultType `json:"type"`
	*inlineQueryResultVoiceBase
}

// MarshalJSON implements json.Marshaler interface.
func (r *InlineQueryResultVoice) MarshalJSON() ([]byte, error) {
	return json.Marshal(&inlineQueryResultVoice{r.Type(), (*inlineQueryResultVoiceBase)(r)})
}

// DocumentMimeType represents the document MIME type.
type DocumentMimeType string

const (
	// PdfDocument represents the PDF document type.
	PdfDocument DocumentMimeType = "application/pdf"

	// ZipDocument represents the ZIP document type.
	ZipDocument DocumentMimeType = "application/zip"
)

// InlineQueryResultDocument represents a link to a file. By default, this file
// will be sent by the user with an optional caption. Alternatively, you can use
// InputMessageContent to send a message with the specified content instead of
// the file. Currently, only .PDF and .ZIP files can be sent using this method.
type InlineQueryResultDocument struct {
	ID                  string                `json:"id"`
	Title               string                `json:"title"`
	Caption             string                `json:"caption,omitempty"`
	ParseMode           ParseMode             `json:"parse_mode,omitempty"`
	URL                 string                `json:"document_url"`
	MimeType            DocumentMimeType      `json:"mime_type"`
	Description         string                `json:"description,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
	ThumbURL            string                `json:"thumb_url,omitempty"`
	ThumbWidth          int                   `json:"thumb_width,omitempty"`
	ThumbHeight         int                   `json:"thumb_height,omitempty"`
}

// Type implements InlineQueryResult interface.
func (r *InlineQueryResultDocument) Type() InlineQueryResultType {
	return DocumentResult
}

type inlineQueryResultDocumentBase InlineQueryResultDocument

type inlineQueryResultDocument struct {
	Type InlineQueryResultType `json:"type"`
	*inlineQueryResultDocumentBase
}

// MarshalJSON implements json.Marshaler interface.
func (r *InlineQueryResultDocument) MarshalJSON() ([]byte, error) {
	return json.Marshal(&inlineQueryResultDocument{r.Type(), (*inlineQueryResultDocumentBase)(r)})
}

// InlineQueryResultLocation represents a location on a map. By default, the
// location will be sent by the user. Alternatively, you can use
// InputMessageContent to send a message with the specified content instead of
// the location.
type InlineQueryResultLocation struct {
	ID                  string                `json:"id"`
	Latitude            float64               `json:"latitude"`
	Longitude           float64               `json:"longitude"`
	Title               string                `json:"title"`
	LivePeriod          int                   `json:"live_period,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
	ThumbURL            string                `json:"thumb_url,omitempty"`
	ThumbWidth          int                   `json:"thumb_width,omitempty"`
	ThumbHeight         int                   `json:"thumb_height,omitempty"`
}

// Type implements InlineQueryResult interface.
func (r *InlineQueryResultLocation) Type() InlineQueryResultType {
	return LocationResult
}

type inlineQueryResultLocationBase InlineQueryResultLocation

type inlineQueryResultLocation struct {
	Type InlineQueryResultType `json:"type"`
	*inlineQueryResultLocationBase
}

// MarshalJSON implements json.Marshaler interface.
func (r *InlineQueryResultLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(&inlineQueryResultLocation{r.Type(), (*inlineQueryResultLocationBase)(r)})
}

// InlineQueryResultVenue represents a venue. By default, the venue will be sent
// by the user. Alternatively, you can use InputMessageContent to send a message
// with the specified content instead of the venue.
type InlineQueryResultVenue struct {
	ID                  string                `json:"id"`
	Latitude            float64               `json:"latitude"`
	Longitude           float64               `json:"longitude"`
	Title               string                `json:"title"`
	Address             string                `json:"address"`
	FoursquareID        string                `json:"foursquare_id,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
	ThumbURL            string                `json:"thumb_url,omitempty"`
	ThumbWidth          int                   `json:"thumb_width,omitempty"`
	ThumbHeight         int                   `json:"thumb_height,omitempty"`
}

// Type implements InlineQueryResult interface.
func (r *InlineQueryResultVenue) Type() InlineQueryResultType {
	return VenueResult
}

type inlineQueryResultVenueBase InlineQueryResultVenue

type inlineQueryResultVenue struct {
	Type InlineQueryResultType `json:"type"`
	*inlineQueryResultVenueBase
}

// MarshalJSON implements json.Marshaler interface.
func (r *InlineQueryResultVenue) MarshalJSON() ([]byte, error) {
	return json.Marshal(&inlineQueryResultVenue{r.Type(), (*inlineQueryResultVenueBase)(r)})
}

// InlineQueryResultContact represents a contact with a phone number. By
// default, this contact will be sent by the user. Alternatively, you can use
// InputMessageContent to send a message with the specified content instead of
// the contact.
type InlineQueryResultContact struct {
	ID                  string                `json:"id"`
	PhoneNumber         string                `json:"phone_number"`
	FirstName           string                `json:"first_name"`
	LastName            string                `json:"last_name,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
	ThumbURL            string                `json:"thumb_url,omitempty"`
	ThumbWidth          int                   `json:"thumb_width,omitempty"`
	ThumbHeight         int                   `json:"thumb_height,omitempty"`
}

// Type implements InlineQueryResult interface.
func (r *InlineQueryResultContact) Type() InlineQueryResultType {
	return ContactResult
}

type inlineQueryResultContactBase InlineQueryResultContact

type inlineQueryResultContact struct {
	Type InlineQueryResultType `json:"type"`
	*inlineQueryResultContactBase
}

// MarshalJSON implements json.Marshaler interface.
func (r *InlineQueryResultContact) MarshalJSON() ([]byte, error) {
	return json.Marshal(&inlineQueryResultContact{r.Type(), (*inlineQueryResultContactBase)(r)})
}

// MessageContentType represents the input message content type.
type MessageContentType string

const (
	// TextMessage represents the text message content type.
	TextMessage MessageContentType = "text"

	// LocationMessage represents the location message content type.
	LocationMessage MessageContentType = "location"

	// VenueMessage represents the venue message content type.
	VenueMessage MessageContentType = "venue"

	// ContactMessage represents the contact message content type.
	ContactMessage MessageContentType = "contact"
)

// InputMessageContent represents the content of a message to be sent as a
// result of an inline query.
type InputMessageContent interface {
	Type() MessageContentType
}

// InputTextMessageContent represents the content of a text message to be sent
// as the result of an inline query.
type InputTextMessageContent struct {
	Text                  string    `json:"message_text"`
	ParseMode             ParseMode `json:"parse_mode,omitempty"`
	DisableWebPagePreview bool      `json:"disable_web_page_preview,omitempty"`
}

// Type implements InputMessageContent interface.
func (c *InputTextMessageContent) Type() MessageContentType {
	return TextMessage
}

// InputLocationMessageContent represents the content of a location message to
// be sent as the result of an inline query.
type InputLocationMessageContent struct {
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	LivePeriod int     `json:"live_period,omitempty"`
}

// Type implements InputMessageContent interface.
func (c *InputLocationMessageContent) Type() MessageContentType {
	return LocationMessage
}

// InputVenueMessageContent represents the content of a location message to
// be sent as the result of an inline query.
type InputVenueMessageContent struct {
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	Title        string  `json:"title"`
	Address      string  `json:"address"`
	FoursquareID string  `json:"foursquare_id,omitempty"`
}

// Type implements InputMessageContent interface.
func (c *InputVenueMessageContent) Type() MessageContentType {
	return VenueMessage
}

// InputContactMessageContent represents the content of a contact message to be
// sent as the result of an inline query.
type InputContactMessageContent struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name,omitempty"`
}

// Type implements InputMessageContent interface.
func (c *InputContactMessageContent) Type() MessageContentType {
	return ContactMessage
}

// ChosenInlineResult represents a result of an inline query that was chosen by
// the user and sent to their chat partner.
type ChosenInlineResult struct {
	ID              string    `json:"result_id"`
	From            *User     `json:"user"`
	Location        *Location `json:"location,omitempty"`
	InlineMessageID string    `json:"inline_message_id,omitempty"`
	Query           string    `json:"query"`
}
