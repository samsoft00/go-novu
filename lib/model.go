package lib

import "io"

type ChannelType string
type GeneralError error

const Version = "v1"

const (
	HTTPStatusOk       = 200
	HTTPStatusCreated  = 201
	HTTPStatusNoChange = 204
	HTTPRedirectOk     = 300
)

const (
	EMAIL  ChannelType = "EMAIL"
	SMS    ChannelType = "SMS"
	DIRECT ChannelType = "DIRECT"
)

type Data struct {
	Acknowledged bool   `json:"acknowledged"`
	Status       string `json:"status"`
}

type Response struct {
	Data Data `json:"data"`
}

type ITriggerPayloadOptions struct {
	To            interface{} `json:"to,omitempty"`
	Payload       interface{} `json:"payload,omitempty"`
	Overrides     interface{} `json:"overrides,omitempty"`
	TransactionId string      `json:"transactionId,omitempty"`
	Actor         interface{} `json:"actor,omitempty"`
}

type TriggerRecipientsTypeArray interface {
	[]string | []SubscriberPayload
}
type TriggerRecipientsTypeSingle interface {
	string | SubscriberPayload
}

type TriggerTopicRecipientsTypeSingle struct {
	TopicKey string `json:"topicKey,omitempty"`
	Type     string `json:"type,omitempty"`
}

type SubscriberPayload struct {
	FirstName string                 `json:"first_name,omitempty"`
	LastName  string                 `json:"last_name,omitempty"`
	Email     string                 `json:"email,omitempty"`
	Phone     string                 `json:"phone,omitempty"`
	Avatar    string                 `json:"avatar,omitempty"`
	Data      map[string]interface{} `json:"data,omitempty"`
}

type TriggerRecipientsType interface {
	TriggerRecipientsTypeSingle | TriggerRecipientsTypeArray | TriggerTopicRecipientsTypeSingle | []TriggerTopicRecipientsTypeSingle
}

type ITriggerPayload interface {
	string | []string | bool | int64 | IAttachmentOptions | []IAttachmentOptions
}

type IAttachmentOptions struct {
	Mime     string        `json:"mime,omitempty"`
	File     io.Reader     `json:"file,omitempty"`
	Name     string        `json:"name,omitempty"`
	Channels []ChannelType `json:"channels,omitempty"`
}

type EventResponse struct {
	Data interface{} `json:"data"`
}

type EventRequest struct {
	Name          string      `json:"name"`
	To            interface{} `json:"to"`
	Payload       interface{} `json:"payload"`
	Overrides     interface{} `json:"overrides,omitempty"`
	TransactionId string      `json:"transactionId,omitempty"`
	Actor         interface{} `json:"actor,omitempty"`
}

type SubscriberResponse struct {
	Data interface{} `json:"data"`
}

type Template struct {
	ID       string `json:"_id"`
	Name     string `json:"name"`
	Critical bool   `json:"critical"`
}

type Preference struct {
	Enabled  bool    `json:"enabled"`
	Channels Channel `json:"channels"`
}

type Channel struct {
	Email bool `json:"email"`
	Sms   bool `json:"sms"`
	Chat  bool `json:"chat"`
	InApp bool `json:"in_app"`
	Push  bool `json:"push"`
}

type SubscriberPreferencesResponse struct {
	Data []struct {
		Template   Template   `json:"template"`
		Preference Preference `json:"preference"`
	} `json:"data"`
}

type UpdateSubscriberPreferencesChannel struct {
	Type    ChannelType `json:"type"`
	Enabled bool        `json:"enabled"`
}

type UpdateSubscriberPreferencesOptions struct {
	Channel []UpdateSubscriberPreferencesChannel `json:"channel,omitempty"`
	Enabled bool                                 `json:"enabled,omitempty"`
}

type ListTopicsResponse struct {
	Page       int                `json:"name"`
	PageSize   int                `json:"pageSize"`
	TotalCount int                `json:"totalCount"`
	Data       []GetTopicResponse `json:"data"`
}

type GetTopicResponse struct {
	Id             string   `json:"_id"`
	OrganizationId string   `json:"_organizationId"`
	EnvironmentId  string   `json:"_environmentId"`
	Key            string   `json:"key"`
	Name           string   `json:"name"`
	Subscribers    []string `json:"subscribers"`
}

type ListTopicsOptions struct {
	Page     *int    `json:"page,omitempty"`
	PageSize *int    `json:"pageSize,omitempty"`
	Key      *string `json:"key,omitempty"`
}

type CreateTopicRequest struct {
	Name string `json:"name"`
	Key  string `json:"key"`
}

type RenameTopicRequest struct {
	Name string `json:"name"`
}

type SubscribersTopicRequest struct {
	Subscribers []string `json:"subscribers"`
}
