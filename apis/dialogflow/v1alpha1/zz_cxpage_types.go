/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type CxPageEventHandlersObservation struct {
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type CxPageEventHandlersParameters struct {

	// The name of the event to handle.
	// +kubebuilder:validation:Optional
	Event *string `json:"event,omitempty" tf:"event,omitempty"`

	// The target flow to transition to.
	// Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>.
	// +kubebuilder:validation:Optional
	TargetFlow *string `json:"targetFlow,omitempty" tf:"target_flow,omitempty"`

	// The target page to transition to.
	// Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>/pages/<Page ID>.
	// +kubebuilder:validation:Optional
	TargetPage *string `json:"targetPage,omitempty" tf:"target_page,omitempty"`

	// The fulfillment to call when the event occurs. Handling webhook errors with a fulfillment enabled with webhook could cause infinite loop. It is invalid to specify such fulfillment for a handler handling webhooks.
	// +kubebuilder:validation:Optional
	TriggerFulfillment []EventHandlersTriggerFulfillmentParameters `json:"triggerFulfillment,omitempty" tf:"trigger_fulfillment,omitempty"`
}

type CxPageObservation struct {
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type CxPageParameters struct {

	// The human-readable name of the page, unique within the agent.
	// +kubebuilder:validation:Required
	DisplayName *string `json:"displayName" tf:"display_name,omitempty"`

	// The fulfillment to call when the session is entering the page.
	// +kubebuilder:validation:Optional
	EntryFulfillment []EntryFulfillmentParameters `json:"entryFulfillment,omitempty" tf:"entry_fulfillment,omitempty"`

	// Handlers associated with the page to handle events such as webhook errors, no match or no input.
	// +kubebuilder:validation:Optional
	EventHandlers []CxPageEventHandlersParameters `json:"eventHandlers,omitempty" tf:"event_handlers,omitempty"`

	// The form associated with the page, used for collecting parameters relevant to the page.
	// +kubebuilder:validation:Optional
	Form []FormParameters `json:"form,omitempty" tf:"form,omitempty"`

	// The language of the following fields in page:
	//
	// Page.entry_fulfillment.messages
	// Page.entry_fulfillment.conditional_cases
	// Page.event_handlers.trigger_fulfillment.messages
	// Page.event_handlers.trigger_fulfillment.conditional_cases
	// Page.form.parameters.fill_behavior.initial_prompt_fulfillment.messages
	// Page.form.parameters.fill_behavior.initial_prompt_fulfillment.conditional_cases
	// Page.form.parameters.fill_behavior.reprompt_event_handlers.messages
	// Page.form.parameters.fill_behavior.reprompt_event_handlers.conditional_cases
	// Page.transition_routes.trigger_fulfillment.messages
	// Page.transition_routes.trigger_fulfillment.conditional_cases
	// If not specified, the agent's default language is used. Many languages are supported. Note: languages must be enabled in the agent before they can be used.
	// +kubebuilder:validation:Optional
	LanguageCode *string `json:"languageCode,omitempty" tf:"language_code,omitempty"`

	// The flow to create a page for.
	// Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>.
	// +kubebuilder:validation:Optional
	Parent *string `json:"parent,omitempty" tf:"parent,omitempty"`

	// Ordered list of TransitionRouteGroups associated with the page. Transition route groups must be unique within a page.
	// If multiple transition routes within a page scope refer to the same intent, then the precedence order is: page's transition route -> page's transition route group -> flow's transition routes.
	// If multiple transition route groups within a page contain the same intent, then the first group in the ordered list takes precedence.
	// Format:projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>/transitionRouteGroups/<TransitionRouteGroup ID>.
	// +kubebuilder:validation:Optional
	TransitionRouteGroups []*string `json:"transitionRouteGroups,omitempty" tf:"transition_route_groups,omitempty"`

	// A list of transitions for the transition rules of this page. They route the conversation to another page in the same flow, or another flow.
	// When we are in a certain page, the TransitionRoutes are evalauted in the following order:
	// TransitionRoutes defined in the page with intent specified.
	// TransitionRoutes defined in the transition route groups with intent specified.
	// TransitionRoutes defined in flow with intent specified.
	// TransitionRoutes defined in the transition route groups with intent specified.
	// TransitionRoutes defined in the page with only condition specified.
	// TransitionRoutes defined in the transition route groups with only condition specified.
	// +kubebuilder:validation:Optional
	TransitionRoutes []CxPageTransitionRoutesParameters `json:"transitionRoutes,omitempty" tf:"transition_routes,omitempty"`
}

type CxPageTransitionRoutesObservation struct {
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type CxPageTransitionRoutesParameters struct {

	// The condition to evaluate against form parameters or session parameters.
	// At least one of intent or condition must be specified. When both intent and condition are specified, the transition can only happen when both are fulfilled.
	// +kubebuilder:validation:Optional
	Condition *string `json:"condition,omitempty" tf:"condition,omitempty"`

	// The unique identifier of an Intent.
	// Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/intents/<Intent ID>. Indicates that the transition can only happen when the given intent is matched. At least one of intent or condition must be specified. When both intent and condition are specified, the transition can only happen when both are fulfilled.
	// +kubebuilder:validation:Optional
	Intent *string `json:"intent,omitempty" tf:"intent,omitempty"`

	// The target flow to transition to.
	// Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>.
	// +kubebuilder:validation:Optional
	TargetFlow *string `json:"targetFlow,omitempty" tf:"target_flow,omitempty"`

	// The target page to transition to.
	// Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>/pages/<Page ID>.
	// +kubebuilder:validation:Optional
	TargetPage *string `json:"targetPage,omitempty" tf:"target_page,omitempty"`

	// The fulfillment to call when the event occurs. Handling webhook errors with a fulfillment enabled with webhook could cause infinite loop. It is invalid to specify such fulfillment for a handler handling webhooks.
	// +kubebuilder:validation:Optional
	TriggerFulfillment []CxPageTransitionRoutesTriggerFulfillmentParameters `json:"triggerFulfillment,omitempty" tf:"trigger_fulfillment,omitempty"`
}

type CxPageTransitionRoutesTriggerFulfillmentObservation struct {
}

type CxPageTransitionRoutesTriggerFulfillmentParameters struct {

	// The list of rich message responses to present to the user.
	// +kubebuilder:validation:Optional
	Messages []TransitionRoutesTriggerFulfillmentMessagesParameters `json:"messages,omitempty" tf:"messages,omitempty"`

	// Whether Dialogflow should return currently queued fulfillment response messages in streaming APIs. If a webhook is specified, it happens before Dialogflow invokes webhook. Warning: 1) This flag only affects streaming API. Responses are still queued and returned once in non-streaming API. 2) The flag can be enabled in any fulfillment but only the first 3 partial responses will be returned. You may only want to apply it to fulfillments that have slow webhooks.
	// +kubebuilder:validation:Optional
	ReturnPartialResponses *bool `json:"returnPartialResponses,omitempty" tf:"return_partial_responses,omitempty"`

	// The tag used by the webhook to identify which fulfillment is being called. This field is required if webhook is specified.
	// +kubebuilder:validation:Optional
	Tag *string `json:"tag,omitempty" tf:"tag,omitempty"`

	// The webhook to call. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/webhooks/<Webhook ID>.
	// +kubebuilder:validation:Optional
	Webhook *string `json:"webhook,omitempty" tf:"webhook,omitempty"`
}

type EntryFulfillmentMessagesObservation struct {
}

type EntryFulfillmentMessagesParameters struct {

	// The text response message.
	// +kubebuilder:validation:Optional
	Text []EntryFulfillmentMessagesTextParameters `json:"text,omitempty" tf:"text,omitempty"`
}

type EntryFulfillmentMessagesTextObservation struct {
	AllowPlaybackInterruption *bool `json:"allowPlaybackInterruption,omitempty" tf:"allow_playback_interruption,omitempty"`
}

type EntryFulfillmentMessagesTextParameters struct {

	// A collection of text responses.
	// +kubebuilder:validation:Optional
	Text []*string `json:"text,omitempty" tf:"text,omitempty"`
}

type EntryFulfillmentObservation struct {
}

type EntryFulfillmentParameters struct {

	// The list of rich message responses to present to the user.
	// +kubebuilder:validation:Optional
	Messages []EntryFulfillmentMessagesParameters `json:"messages,omitempty" tf:"messages,omitempty"`

	// Whether Dialogflow should return currently queued fulfillment response messages in streaming APIs. If a webhook is specified, it happens before Dialogflow invokes webhook. Warning: 1) This flag only affects streaming API. Responses are still queued and returned once in non-streaming API. 2) The flag can be enabled in any fulfillment but only the first 3 partial responses will be returned. You may only want to apply it to fulfillments that have slow webhooks.
	// +kubebuilder:validation:Optional
	ReturnPartialResponses *bool `json:"returnPartialResponses,omitempty" tf:"return_partial_responses,omitempty"`

	// The tag used by the webhook to identify which fulfillment is being called. This field is required if webhook is specified.
	// +kubebuilder:validation:Optional
	Tag *string `json:"tag,omitempty" tf:"tag,omitempty"`

	// The webhook to call. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/webhooks/<Webhook ID>.
	// +kubebuilder:validation:Optional
	Webhook *string `json:"webhook,omitempty" tf:"webhook,omitempty"`
}

type EventHandlersTriggerFulfillmentMessagesObservation struct {
}

type EventHandlersTriggerFulfillmentMessagesParameters struct {

	// The text response message.
	// +kubebuilder:validation:Optional
	Text []TriggerFulfillmentMessagesTextParameters `json:"text,omitempty" tf:"text,omitempty"`
}

type EventHandlersTriggerFulfillmentObservation struct {
}

type EventHandlersTriggerFulfillmentParameters struct {

	// The list of rich message responses to present to the user.
	// +kubebuilder:validation:Optional
	Messages []EventHandlersTriggerFulfillmentMessagesParameters `json:"messages,omitempty" tf:"messages,omitempty"`

	// Whether Dialogflow should return currently queued fulfillment response messages in streaming APIs. If a webhook is specified, it happens before Dialogflow invokes webhook. Warning: 1) This flag only affects streaming API. Responses are still queued and returned once in non-streaming API. 2) The flag can be enabled in any fulfillment but only the first 3 partial responses will be returned. You may only want to apply it to fulfillments that have slow webhooks.
	// +kubebuilder:validation:Optional
	ReturnPartialResponses *bool `json:"returnPartialResponses,omitempty" tf:"return_partial_responses,omitempty"`

	// The tag used by the webhook to identify which fulfillment is being called. This field is required if webhook is specified.
	// +kubebuilder:validation:Optional
	Tag *string `json:"tag,omitempty" tf:"tag,omitempty"`

	// The webhook to call. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/webhooks/<Webhook ID>.
	// +kubebuilder:validation:Optional
	Webhook *string `json:"webhook,omitempty" tf:"webhook,omitempty"`
}

type FillBehaviorObservation struct {
}

type FillBehaviorParameters struct {

	// The fulfillment to provide the initial prompt that the agent can present to the user in order to fill the parameter.
	// +kubebuilder:validation:Optional
	InitialPromptFulfillment []InitialPromptFulfillmentParameters `json:"initialPromptFulfillment,omitempty" tf:"initial_prompt_fulfillment,omitempty"`
}

type FormObservation struct {
}

type FormParameters struct {

	// Parameters to collect from the user.
	// +kubebuilder:validation:Optional
	Parameters []FormParametersParameters `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type FormParametersObservation struct {
}

type FormParametersParameters struct {

	// The human-readable name of the parameter, unique within the form.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The entity type of the parameter.
	// Format: projects/-/locations/-/agents/-/entityTypes/<System Entity Type ID> for system entity types (for example, projects/-/locations/-/agents/-/entityTypes/sys.date), or projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/entityTypes/<Entity Type ID> for developer entity types.
	// +kubebuilder:validation:Optional
	EntityType *string `json:"entityType,omitempty" tf:"entity_type,omitempty"`

	// Defines fill behavior for the parameter.
	// +kubebuilder:validation:Optional
	FillBehavior []FillBehaviorParameters `json:"fillBehavior,omitempty" tf:"fill_behavior,omitempty"`

	// Indicates whether the parameter represents a list of values.
	// +kubebuilder:validation:Optional
	IsList *bool `json:"isList,omitempty" tf:"is_list,omitempty"`

	// Indicates whether the parameter content should be redacted in log.
	// If redaction is enabled, the parameter content will be replaced by parameter name during logging. Note: the parameter content is subject to redaction if either parameter level redaction or entity type level redaction is enabled.
	// +kubebuilder:validation:Optional
	Redact *bool `json:"redact,omitempty" tf:"redact,omitempty"`

	// Indicates whether the parameter is required. Optional parameters will not trigger prompts; however, they are filled if the user specifies them.
	// Required parameters must be filled before form filling concludes.
	// +kubebuilder:validation:Optional
	Required *bool `json:"required,omitempty" tf:"required,omitempty"`
}

type InitialPromptFulfillmentMessagesObservation struct {
}

type InitialPromptFulfillmentMessagesParameters struct {

	// The text response message.
	// +kubebuilder:validation:Optional
	Text []InitialPromptFulfillmentMessagesTextParameters `json:"text,omitempty" tf:"text,omitempty"`
}

type InitialPromptFulfillmentMessagesTextObservation struct {
	AllowPlaybackInterruption *bool `json:"allowPlaybackInterruption,omitempty" tf:"allow_playback_interruption,omitempty"`
}

type InitialPromptFulfillmentMessagesTextParameters struct {

	// A collection of text responses.
	// +kubebuilder:validation:Optional
	Text []*string `json:"text,omitempty" tf:"text,omitempty"`
}

type InitialPromptFulfillmentObservation struct {
}

type InitialPromptFulfillmentParameters struct {

	// The list of rich message responses to present to the user.
	// +kubebuilder:validation:Optional
	Messages []InitialPromptFulfillmentMessagesParameters `json:"messages,omitempty" tf:"messages,omitempty"`

	// Whether Dialogflow should return currently queued fulfillment response messages in streaming APIs. If a webhook is specified, it happens before Dialogflow invokes webhook. Warning: 1) This flag only affects streaming API. Responses are still queued and returned once in non-streaming API. 2) The flag can be enabled in any fulfillment but only the first 3 partial responses will be returned. You may only want to apply it to fulfillments that have slow webhooks.
	// +kubebuilder:validation:Optional
	ReturnPartialResponses *bool `json:"returnPartialResponses,omitempty" tf:"return_partial_responses,omitempty"`

	// The tag used by the webhook to identify which fulfillment is being called. This field is required if webhook is specified.
	// +kubebuilder:validation:Optional
	Tag *string `json:"tag,omitempty" tf:"tag,omitempty"`

	// The webhook to call. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/webhooks/<Webhook ID>.
	// +kubebuilder:validation:Optional
	Webhook *string `json:"webhook,omitempty" tf:"webhook,omitempty"`
}

type TransitionRoutesTriggerFulfillmentMessagesObservation struct {
}

type TransitionRoutesTriggerFulfillmentMessagesParameters struct {

	// The text response message.
	// +kubebuilder:validation:Optional
	Text []TransitionRoutesTriggerFulfillmentMessagesTextParameters `json:"text,omitempty" tf:"text,omitempty"`
}

type TransitionRoutesTriggerFulfillmentMessagesTextObservation struct {
	AllowPlaybackInterruption *bool `json:"allowPlaybackInterruption,omitempty" tf:"allow_playback_interruption,omitempty"`
}

type TransitionRoutesTriggerFulfillmentMessagesTextParameters struct {

	// A collection of text responses.
	// +kubebuilder:validation:Optional
	Text []*string `json:"text,omitempty" tf:"text,omitempty"`
}

type TriggerFulfillmentMessagesTextObservation struct {
	AllowPlaybackInterruption *bool `json:"allowPlaybackInterruption,omitempty" tf:"allow_playback_interruption,omitempty"`
}

type TriggerFulfillmentMessagesTextParameters struct {

	// A collection of text responses.
	// +kubebuilder:validation:Optional
	Text []*string `json:"text,omitempty" tf:"text,omitempty"`
}

// CxPageSpec defines the desired state of CxPage
type CxPageSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CxPageParameters `json:"forProvider"`
}

// CxPageStatus defines the observed state of CxPage.
type CxPageStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CxPageObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CxPage is the Schema for the CxPages API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfgcp}
type CxPage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CxPageSpec   `json:"spec"`
	Status            CxPageStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CxPageList contains a list of CxPages
type CxPageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CxPage `json:"items"`
}

// Repository type metadata.
var (
	CxPageKind             = "CxPage"
	CxPageGroupKind        = schema.GroupKind{Group: Group, Kind: CxPageKind}.String()
	CxPageKindAPIVersion   = CxPageKind + "." + GroupVersion.String()
	CxPageGroupVersionKind = GroupVersion.WithKind(CxPageKind)
)

func init() {
	SchemeBuilder.Register(&CxPage{}, &CxPageList{})
}
