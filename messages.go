package gohighlevel

import "fmt"

type MessagesService struct {
	client *Client
}

// Possible values: [SMS, Email, WhatsApp, IG, FB, Custom, Live_Chat, InternalComment]
type MessageType string

const (
	MessageTypeSMS             MessageType = "SMS"
	MessageTypeEmail           MessageType = "Email"
	MessageTypeWhatsApp        MessageType = "WhatsApp"
	MessageTypeIG              MessageType = "IG"
	MessageTypeFB              MessageType = "FB"
	MessageTypeCustom          MessageType = "Custom"
	MessageTypeLiveChat        MessageType = "Live_Chat"
	MessageTypeInternalComment MessageType = "InternalComment"
)

type EmailReplyMode string

const (
	EmailReplyModeReply    EmailReplyMode = "reply"
	EmailReplyModeReplyAll EmailReplyMode = "reply_all"
)

type MessageStatus string

const (
	MessageStatusDelivered MessageStatus = "delivered"
	MessageStatusFailed    MessageStatus = "failed"
	MessageStatusPending   MessageStatus = "pending"
	MessageStatusRead      MessageStatus = "read"
)

type SendMessageRequest struct {
	// Type of message being sent (required)
	Type MessageType `json:"type"`

	// ID of the contact receiving the message (required)
	ContactID string `json:"contactId"`

	// ID of the associated appointment
	AppointmentID string `json:"appointmentId,omitempty"`

	//  Array of attachment URLs
	Attachments []string `json:"attachments,omitempty"`

	// Email address to send from
	EmailFrom string `json:"emailFrom,omitempty"`

	// Array of CC email addresses
	EmailCc []string `json:"emailCc,omitempty"`

	// Array of BCC email addresses
	EmailBcc []string `json:"emailBcc,omitempty"`

	// html content of the message
	HTML string `json:"html,omitempty"`

	// Text content of the message. For InternalComment type, use @username<userId>actualUserId</userId> format to mention team members.
	// The mentioned user IDs must also be included in the mentions array.
	Message string `json:"message,omitempty"`

	// Subject line for email messages
	Subject string `json:"subject,omitempty"`

	// ID of message being replied to
	ReplyMessageID string `json:"replyMessageId,omitempty"`

	// ID of message template
	TemplateID string `json:"templateId,omitempty"`

	// ID of message thread. For email messages, this is the message ID that contains multiple email messages in the thread
	ThreadID string `json:"threadId,omitempty"`

	// UTC Timestamp (in seconds) at which the message should be scheduled
	ScheduledTimestamp int64 `json:"scheduledTimestamp,omitempty"`

	// ID of conversation provider
	ConversationProvider string `json:"conversationProviderId,omitempty"`

	// Email address to send to, if different from contact's primary email. This should be a valid email address associated with the contact.
	EmailTo string `json:"emailTo,omitempty"`

	// Mode for email replies Possible values: [reply, reply_all]
	EmailReplyMode EmailReplyMode `json:"emailReplyMode,omitempty"`

	// Phone number used as the sender number for outbound messages
	FromNumber string `json:"fromNumber,omitempty"`

	// Recipient phone number for outbound messages
	ToNumber string `json:"toNumber,omitempty"`

	// Message status (required)
	Status MessageStatus `json:"status"`

	// Array of user IDs mentioned in the message. Required for InternalComment type.
	// User IDs correspond to team members tagged with @username<userId>id</userId> format in the message text.
	// Example: ["userId123","userId456"]
	Mentions []string `json:"mentions,omitempty"`

	// Use this field to specify the user who is making the internal comment when type is 'InternalComment'. If not provided, the comment will be attributed to the system or default user.
	// Example: "user123"
	UserID string `json:"userId,omitempty"`
}

type MessageResponse struct {
	// Conversation ID.
	ConversationID string `json:"conversationId"`

	// This contains the email message id (only for Email type). Use this ID to send inbound replies to GHL to create a threaded email.
	EmailMessageID string `json:"emailMessageId,omitempty"`

	// This is the main Message ID
	MessageID string `json:"messageId"`

	// When sending via the GMB channel, we will be returning list of messageIds instead of single
	MessageIDs []string `json:"messageIds,omitempty"`

	// Additional response message when sending a workflow message
	// Example: "Message queued successfully."
	Msg string `json:"msg,omitempty"`
}

// Send Sends a new message
func (s *MessagesService) Send(req *SendMessageRequest) (*MessageResponse, error) {
	if req.ContactID == "" {
		return nil, fmt.Errorf("contactId is required")
	}

	if req.Type == "" {
		return nil, fmt.Errorf("type is required")
	}

	var result MessageResponse
	err := s.client.doRequest("POST", "/conversations/messages", req, &result, "2021-04-15")
	if err != nil {
		return nil, err
	}

	return &result, nil
}
