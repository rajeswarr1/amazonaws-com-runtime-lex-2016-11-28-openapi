package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// ActiveContext represents the ActiveContext schema from the OpenAPI specification
type ActiveContext struct {
	Parameters interface{} `json:"parameters"`
	Timetolive interface{} `json:"timeToLive"`
	Name interface{} `json:"name"`
}

// PostTextResponse represents the PostTextResponse schema from the OpenAPI specification
type PostTextResponse struct {
	Sessionattributes interface{} `json:"sessionAttributes,omitempty"`
	Messageformat interface{} `json:"messageFormat,omitempty"`
	Sessionid interface{} `json:"sessionId,omitempty"`
	Dialogstate interface{} `json:"dialogState,omitempty"`
	Responsecard interface{} `json:"responseCard,omitempty"`
	Activecontexts interface{} `json:"activeContexts,omitempty"`
	Botversion interface{} `json:"botVersion,omitempty"`
	Nluintentconfidence interface{} `json:"nluIntentConfidence,omitempty"`
	Slottoelicit interface{} `json:"slotToElicit,omitempty"`
	Slots interface{} `json:"slots,omitempty"`
	Alternativeintents interface{} `json:"alternativeIntents,omitempty"`
	Intentname interface{} `json:"intentName,omitempty"`
	Message interface{} `json:"message,omitempty"`
	Sentimentresponse interface{} `json:"sentimentResponse,omitempty"`
}

// PutSessionRequest represents the PutSessionRequest schema from the OpenAPI specification
type PutSessionRequest struct {
	Dialogaction interface{} `json:"dialogAction,omitempty"`
	Recentintentsummaryview interface{} `json:"recentIntentSummaryView,omitempty"`
	Sessionattributes interface{} `json:"sessionAttributes,omitempty"`
	Activecontexts interface{} `json:"activeContexts,omitempty"`
}

// Button represents the Button schema from the OpenAPI specification
type Button struct {
	Text interface{} `json:"text"`
	Value interface{} `json:"value"`
}

// GenericAttachment represents the GenericAttachment schema from the OpenAPI specification
type GenericAttachment struct {
	Subtitle interface{} `json:"subTitle,omitempty"`
	Title interface{} `json:"title,omitempty"`
	Attachmentlinkurl interface{} `json:"attachmentLinkUrl,omitempty"`
	Buttons interface{} `json:"buttons,omitempty"`
	Imageurl interface{} `json:"imageUrl,omitempty"`
}

// PostTextRequest represents the PostTextRequest schema from the OpenAPI specification
type PostTextRequest struct {
	Sessionattributes interface{} `json:"sessionAttributes,omitempty"`
	Activecontexts interface{} `json:"activeContexts,omitempty"`
	Inputtext interface{} `json:"inputText"`
	Requestattributes interface{} `json:"requestAttributes,omitempty"`
}

// SentimentResponse represents the SentimentResponse schema from the OpenAPI specification
type SentimentResponse struct {
	Sentimentlabel interface{} `json:"sentimentLabel,omitempty"`
	Sentimentscore interface{} `json:"sentimentScore,omitempty"`
}

// ResponseCard represents the ResponseCard schema from the OpenAPI specification
type ResponseCard struct {
	Contenttype interface{} `json:"contentType,omitempty"`
	Genericattachments interface{} `json:"genericAttachments,omitempty"`
	Version interface{} `json:"version,omitempty"`
}

// DeleteSessionRequest represents the DeleteSessionRequest schema from the OpenAPI specification
type DeleteSessionRequest struct {
}

// IntentConfidence represents the IntentConfidence schema from the OpenAPI specification
type IntentConfidence struct {
	Score interface{} `json:"score,omitempty"`
}

// GetSessionRequest represents the GetSessionRequest schema from the OpenAPI specification
type GetSessionRequest struct {
}

// PostContentRequest represents the PostContentRequest schema from the OpenAPI specification
type PostContentRequest struct {
	Inputstream interface{} `json:"inputStream"`
}

// ActiveContextParametersMap represents the ActiveContextParametersMap schema from the OpenAPI specification
type ActiveContextParametersMap struct {
}

// ActiveContextTimeToLive represents the ActiveContextTimeToLive schema from the OpenAPI specification
type ActiveContextTimeToLive struct {
	Turnstolive interface{} `json:"turnsToLive,omitempty"`
	Timetoliveinseconds interface{} `json:"timeToLiveInSeconds,omitempty"`
}

// PredictedIntent represents the PredictedIntent schema from the OpenAPI specification
type PredictedIntent struct {
	Intentname interface{} `json:"intentName,omitempty"`
	Nluintentconfidence interface{} `json:"nluIntentConfidence,omitempty"`
	Slots interface{} `json:"slots,omitempty"`
}

// StringMap represents the StringMap schema from the OpenAPI specification
type StringMap struct {
}

// DialogAction represents the DialogAction schema from the OpenAPI specification
type DialogAction struct {
	Slots interface{} `json:"slots,omitempty"`
	TypeField interface{} `json:"type"`
	Fulfillmentstate interface{} `json:"fulfillmentState,omitempty"`
	Intentname interface{} `json:"intentName,omitempty"`
	Message interface{} `json:"message,omitempty"`
	Messageformat interface{} `json:"messageFormat,omitempty"`
	Slottoelicit interface{} `json:"slotToElicit,omitempty"`
}

// PutSessionResponse represents the PutSessionResponse schema from the OpenAPI specification
type PutSessionResponse struct {
	Audiostream interface{} `json:"audioStream,omitempty"`
}

// IntentSummary represents the IntentSummary schema from the OpenAPI specification
type IntentSummary struct {
	Slottoelicit interface{} `json:"slotToElicit,omitempty"`
	Slots interface{} `json:"slots,omitempty"`
	Checkpointlabel interface{} `json:"checkpointLabel,omitempty"`
	Confirmationstatus interface{} `json:"confirmationStatus,omitempty"`
	Dialogactiontype interface{} `json:"dialogActionType"`
	Fulfillmentstate interface{} `json:"fulfillmentState,omitempty"`
	Intentname interface{} `json:"intentName,omitempty"`
}

// GetSessionResponse represents the GetSessionResponse schema from the OpenAPI specification
type GetSessionResponse struct {
	Activecontexts interface{} `json:"activeContexts,omitempty"`
	Dialogaction interface{} `json:"dialogAction,omitempty"`
	Recentintentsummaryview interface{} `json:"recentIntentSummaryView,omitempty"`
	Sessionattributes interface{} `json:"sessionAttributes,omitempty"`
	Sessionid interface{} `json:"sessionId,omitempty"`
}

// DeleteSessionResponse represents the DeleteSessionResponse schema from the OpenAPI specification
type DeleteSessionResponse struct {
	Botalias interface{} `json:"botAlias,omitempty"`
	Botname interface{} `json:"botName,omitempty"`
	Sessionid interface{} `json:"sessionId,omitempty"`
	Userid interface{} `json:"userId,omitempty"`
}

// PostContentResponse represents the PostContentResponse schema from the OpenAPI specification
type PostContentResponse struct {
	Audiostream interface{} `json:"audioStream,omitempty"`
}
