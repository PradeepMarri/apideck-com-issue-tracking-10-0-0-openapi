package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// IssueTrackingWebhookEvent represents the IssueTrackingWebhookEvent schema from the OpenAPI specification
type IssueTrackingWebhookEvent struct {
	Entity_type string `json:"entity_type,omitempty"` // The type entity that triggered this event
	Entity_url string `json:"entity_url,omitempty"` // The url to retrieve entity detail.
	Execution_attempt float64 `json:"execution_attempt,omitempty"` // The current count this request event has been attempted
	Unified_api string `json:"unified_api,omitempty"` // Name of Apideck Unified API
	Consumer_id string `json:"consumer_id,omitempty"` // Unique consumer identifier. You can freely choose a consumer ID yourself. Most of the time, this is an ID of your internal data model that represents a user or account in your system (for example account:12345). If the consumer doesn't exist yet, Vault will upsert a consumer based on your ID.
	Entity_id string `json:"entity_id,omitempty"` // The service provider's ID of the entity that triggered this event
	Event_id string `json:"event_id,omitempty"` // Unique reference to this request event
	Occurred_at string `json:"occurred_at,omitempty"` // ISO Datetime for when the original event occurred
	Service_id string `json:"service_id,omitempty"` // Service provider identifier
	Event_type string `json:"event_type,omitempty"`
}

// GetCollectionUsersResponse represents the GetCollectionUsersResponse schema from the OpenAPI specification
type GetCollectionUsersResponse struct {
	Links Links `json:"links,omitempty"` // Links to navigate to previous or next pages through the API
	Meta Meta `json:"meta,omitempty"` // Response metadata
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data []CollectionUser `json:"data"`
}

// Department represents the Department schema from the OpenAPI specification
type Department struct {
	Updated_by string `json:"updated_by,omitempty"` // The user who last updated the object.
	Created_by string `json:"created_by,omitempty"` // The user who created the object.
	Description string `json:"description,omitempty"`
	Id string `json:"id,omitempty"` // A unique identifier for an object.
	Parent_id string `json:"parent_id,omitempty"` // Parent ID
	Code string `json:"code,omitempty"`
	Created_at string `json:"created_at,omitempty"` // The date and time when the object was created.
	Name string `json:"name,omitempty"` // Department name
	Custom_mappings CustomMappings `json:"custom_mappings,omitempty"` // When custom mappings are configured on the resource, the result is included here.
	Updated_at string `json:"updated_at,omitempty"` // The date and time when the object was last updated.
}

// GetCollectionTagsResponse represents the GetCollectionTagsResponse schema from the OpenAPI specification
type GetCollectionTagsResponse struct {
	Data []CollectionTag `json:"data"`
	Links Links `json:"links,omitempty"` // Links to navigate to previous or next pages through the API
	Meta Meta `json:"meta,omitempty"` // Response metadata
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
}

// CustomMappings represents the CustomMappings schema from the OpenAPI specification
type CustomMappings struct {
}

// Links represents the Links schema from the OpenAPI specification
type Links struct {
	Previous string `json:"previous,omitempty"` // Link to navigate to the previous page through the API
	Current string `json:"current,omitempty"` // Link to navigate to the current page through the API
	Next string `json:"next,omitempty"` // Link to navigate to the previous page through the API
}

// GetCollectionUserResponse represents the GetCollectionUserResponse schema from the OpenAPI specification
type GetCollectionUserResponse struct {
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data CollectionUser `json:"data"`
	Operation string `json:"operation"` // Operation performed
}

// NotImplementedResponse represents the NotImplementedResponse schema from the OpenAPI specification
type NotImplementedResponse struct {
	Detail interface{} `json:"detail,omitempty"` // Contains parameter or domain specific information related to the error and why it occurred.
	ErrorField string `json:"error,omitempty"` // Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 7231)
	Message string `json:"message,omitempty"` // A human-readable message providing more details about the error.
	Ref string `json:"ref,omitempty"` // Link to documentation of error type
	Status_code float64 `json:"status_code,omitempty"` // HTTP status code
	Type_name string `json:"type_name,omitempty"` // The type of error returned
}

// CollectionsSort represents the CollectionsSort schema from the OpenAPI specification
type CollectionsSort struct {
	By string `json:"by,omitempty"` // The field on which to sort the Collections
	Direction string `json:"direction,omitempty"` // The direction in which to sort the results
}

// Email represents the Email schema from the OpenAPI specification
type Email struct {
	TypeField string `json:"type,omitempty"` // Email type
	Email string `json:"email"` // Email address
	Id string `json:"id,omitempty"` // Unique identifier for the email address
}

// GetTicketsResponse represents the GetTicketsResponse schema from the OpenAPI specification
type GetTicketsResponse struct {
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data []Ticket `json:"data"`
	Links Links `json:"links,omitempty"` // Links to navigate to previous or next pages through the API
	Meta Meta `json:"meta,omitempty"` // Response metadata
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
}

// BadRequestResponse represents the BadRequestResponse schema from the OpenAPI specification
type BadRequestResponse struct {
	Type_name string `json:"type_name,omitempty"` // The type of error returned
	Detail interface{} `json:"detail,omitempty"` // Contains parameter or domain specific information related to the error and why it occurred.
	ErrorField string `json:"error,omitempty"` // Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 7231)
	Message string `json:"message,omitempty"` // A human-readable message providing more details about the error.
	Ref string `json:"ref,omitempty"` // Link to documentation of error type
	Status_code float64 `json:"status_code,omitempty"` // HTTP status code
}

// GetCollectionResponse represents the GetCollectionResponse schema from the OpenAPI specification
type GetCollectionResponse struct {
	Data Collection `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
}

// DeleteCommentResponse represents the DeleteCommentResponse schema from the OpenAPI specification
type DeleteCommentResponse struct {
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
}

// UnexpectedErrorResponse represents the UnexpectedErrorResponse schema from the OpenAPI specification
type UnexpectedErrorResponse struct {
	Type_name string `json:"type_name,omitempty"` // The type of error returned
	Detail interface{} `json:"detail,omitempty"` // Contains parameter or domain specific information related to the error and why it occurred.
	ErrorField string `json:"error,omitempty"` // Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 7231)
	Message string `json:"message,omitempty"` // A human-readable message providing more details about the error.
	Ref string `json:"ref,omitempty"` // Link to documentation of error type
	Status_code float64 `json:"status_code,omitempty"` // HTTP status code
}

// UnauthorizedResponse represents the UnauthorizedResponse schema from the OpenAPI specification
type UnauthorizedResponse struct {
	Status_code float64 `json:"status_code,omitempty"` // HTTP status code
	Type_name string `json:"type_name,omitempty"` // The type of error returned
	Detail string `json:"detail,omitempty"` // Contains parameter or domain specific information related to the error and why it occurred.
	ErrorField string `json:"error,omitempty"` // Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 7231)
	Message string `json:"message,omitempty"` // A human-readable message providing more details about the error.
	Ref string `json:"ref,omitempty"` // Link to documentation of error type
}

// Ticket represents the Ticket schema from the OpenAPI specification
type Ticket struct {
	Status string `json:"status,omitempty"` // The current status of the ticket. Possible values include: open, in_progress, closed, or - in cases where there is no clear mapping - the original value passed through.
	Assignees []Assignee `json:"assignees,omitempty"`
	Created_at string `json:"created_at,omitempty"` // The date and time when the object was created.
	Custom_mappings CustomMappings `json:"custom_mappings,omitempty"` // When custom mappings are configured on the resource, the result is included here.
	Id string `json:"id"` // A unique identifier for an object.
	Collection_id string `json:"collection_id,omitempty"` // The ticket's collection ID
	Parent_id string `json:"parent_id,omitempty"` // The ticket's parent ID
	Updated_at string `json:"updated_at,omitempty"` // The date and time when the object was last updated.
	Due_date string `json:"due_date,omitempty"` // Due date of the ticket
	Priority string `json:"priority,omitempty"` // Priority of the ticket
	TypeField string `json:"type,omitempty"` // The ticket's type
	Completed_at string `json:"completed_at,omitempty"` // When the ticket was completed
	Created_by string `json:"created_by,omitempty"` // The user who created the object.
	Tags []CollectionTag `json:"tags,omitempty"`
	Subject string `json:"subject,omitempty"` // Subject of the ticket
	Description string `json:"description,omitempty"` // The ticket's description. HTML version of description is mapped if supported by the third-party platform
}

// CollectionTag represents the CollectionTag schema from the OpenAPI specification
type CollectionTag struct {
	Name string `json:"name,omitempty"` // The name of the tag.
	Custom_mappings CustomMappings `json:"custom_mappings,omitempty"` // When custom mappings are configured on the resource, the result is included here.
	Id string `json:"id"` // A unique identifier for an object.
}

// GetCommentsResponse represents the GetCommentsResponse schema from the OpenAPI specification
type GetCommentsResponse struct {
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data []CollectionTicketComment `json:"data"`
	Links Links `json:"links,omitempty"` // Links to navigate to previous or next pages through the API
	Meta Meta `json:"meta,omitempty"` // Response metadata
}

// UpdateTicketResponse represents the UpdateTicketResponse schema from the OpenAPI specification
type UpdateTicketResponse struct {
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
}

// Meta represents the Meta schema from the OpenAPI specification
type Meta struct {
	Items_on_page int `json:"items_on_page,omitempty"` // Number of items returned in the data property of the response
	Cursors map[string]interface{} `json:"cursors,omitempty"` // Cursors to navigate to previous or next pages through the API
}

// TicketsSort represents the TicketsSort schema from the OpenAPI specification
type TicketsSort struct {
	Direction string `json:"direction,omitempty"` // The direction in which to sort the results
	By string `json:"by,omitempty"` // The field on which to sort the Tickets
}

// GetTicketResponse represents the GetTicketResponse schema from the OpenAPI specification
type GetTicketResponse struct {
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data Ticket `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
}

// DeleteTicketResponse represents the DeleteTicketResponse schema from the OpenAPI specification
type DeleteTicketResponse struct {
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
}

// GetCommentResponse represents the GetCommentResponse schema from the OpenAPI specification
type GetCommentResponse struct {
	Data CollectionTicketComment `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
}

// CommentsSort represents the CommentsSort schema from the OpenAPI specification
type CommentsSort struct {
	Direction string `json:"direction,omitempty"` // The direction in which to sort the results
	By string `json:"by,omitempty"` // The field on which to sort the Comments
}

// UnifiedId represents the UnifiedId schema from the OpenAPI specification
type UnifiedId struct {
	Id string `json:"id"` // The unique identifier of the resource
}

// PhoneNumber represents the PhoneNumber schema from the OpenAPI specification
type PhoneNumber struct {
	Number string `json:"number"` // The phone number
	TypeField string `json:"type,omitempty"` // The type of phone number
	Area_code string `json:"area_code,omitempty"` // The area code of the phone number, e.g. 323
	Country_code string `json:"country_code,omitempty"` // The country code of the phone number, e.g. +1
	Extension string `json:"extension,omitempty"` // The extension of the phone number
	Id string `json:"id,omitempty"` // Unique identifier of the phone number
}

// PaymentRequiredResponse represents the PaymentRequiredResponse schema from the OpenAPI specification
type PaymentRequiredResponse struct {
	Status_code float64 `json:"status_code,omitempty"` // HTTP status code
	Type_name string `json:"type_name,omitempty"` // The type of error returned
	Detail string `json:"detail,omitempty"` // Contains parameter or domain specific information related to the error and why it occurred.
	ErrorField string `json:"error,omitempty"` // Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 7231)
	Message string `json:"message,omitempty"` // A human-readable message providing more details about the error.
	Ref string `json:"ref,omitempty"` // Link to documentation of error type
}

// Website represents the Website schema from the OpenAPI specification
type Website struct {
	TypeField string `json:"type,omitempty"` // The type of website
	Url string `json:"url"` // The website URL
	Id string `json:"id,omitempty"` // Unique identifier for the website
}

// UpdateCommentResponse represents the UpdateCommentResponse schema from the OpenAPI specification
type UpdateCommentResponse struct {
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
}

// CollectionUser represents the CollectionUser schema from the OpenAPI specification
type CollectionUser struct {
	Custom_mappings CustomMappings `json:"custom_mappings,omitempty"` // When custom mappings are configured on the resource, the result is included here.
	Email string `json:"email,omitempty"` // Email address of the user
	Last_name string `json:"last_name,omitempty"` // Last name of the user
	Name string `json:"name,omitempty"` // Full name of the user
	Photo_url string `json:"photo_url,omitempty"` // The URL of the photo of a person.
	Id string `json:"id,omitempty"` // A unique identifier for an object.
	First_name string `json:"first_name,omitempty"` // First name of the user
	Created_at string `json:"created_at,omitempty"` // The date and time when the object was created.
	Updated_at string `json:"updated_at,omitempty"` // The date and time when the object was last updated.
}

// Assignee represents the Assignee schema from the OpenAPI specification
type Assignee struct {
	Id string `json:"id"` // A unique identifier for an object.
	Username string `json:"username,omitempty"`
}

// Collection represents the Collection schema from the OpenAPI specification
type Collection struct {
	Parent_id string `json:"parent_id,omitempty"` // The collections's parent ID
	TypeField string `json:"type,omitempty"` // The collections's type
	Updated_at string `json:"updated_at,omitempty"` // The date and time when the object was last updated.
	Created_at string `json:"created_at,omitempty"` // The date and time when the object was created.
	Custom_mappings CustomMappings `json:"custom_mappings,omitempty"` // When custom mappings are configured on the resource, the result is included here.
	Description string `json:"description,omitempty"` // Description of the collection
	Id string `json:"id"` // A unique identifier for an object.
	Name string `json:"name,omitempty"` // Name of the collection
}

// CreateTicketResponse represents the CreateTicketResponse schema from the OpenAPI specification
type CreateTicketResponse struct {
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
}

// IssuesFilter represents the IssuesFilter schema from the OpenAPI specification
type IssuesFilter struct {
	Assignee_id string `json:"assignee_id,omitempty"` // Only return tickets assigned to a specific user
	Since string `json:"since,omitempty"` // Only return tickets since a specific date
	Status []string `json:"status,omitempty"` // Filter by ticket status, can be `open`, `closed` or `all`. Will passthrough if none of the above match
}

// PassThroughQuery represents the PassThroughQuery schema from the OpenAPI specification
type PassThroughQuery struct {
	Example_downstream_property string `json:"example_downstream_property,omitempty"` // All passthrough query parameters are passed along to the connector as is (?pass_through[search]=leads becomes ?search=leads)
}

// Address represents the Address schema from the OpenAPI specification
type Address struct {
	Line3 string `json:"line3,omitempty"` // Line 3 of the address
	Notes string `json:"notes,omitempty"` // Additional notes
	Longitude string `json:"longitude,omitempty"` // Longitude of the address
	Name string `json:"name,omitempty"` // The name of the address.
	Email string `json:"email,omitempty"` // Email address of the address
	Latitude string `json:"latitude,omitempty"` // Latitude of the address
	Phone_number string `json:"phone_number,omitempty"` // Phone number of the address
	Contact_name string `json:"contact_name,omitempty"` // Name of the contact person at the address
	Id string `json:"id,omitempty"` // Unique identifier for the address.
	Postal_code string `json:"postal_code,omitempty"` // Zip code or equivalent.
	Row_version string `json:"row_version,omitempty"` // A binary value used to detect updates to a object and prevent data conflicts. It is incremented each time an update is made to the object.
	State string `json:"state,omitempty"` // Name of state
	Line2 string `json:"line2,omitempty"` // Line 2 of the address
	Line4 string `json:"line4,omitempty"` // Line 4 of the address
	Salutation string `json:"salutation,omitempty"` // Salutation of the contact person at the address
	StringField string `json:"string,omitempty"` // The address string. Some APIs don't provide structured address data.
	TypeField string `json:"type,omitempty"` // The type of address.
	City string `json:"city,omitempty"` // Name of city.
	County string `json:"county,omitempty"` // Address field that holds a sublocality, such as a county
	Line1 string `json:"line1,omitempty"` // Line 1 of the address e.g. number, street, suite, apt #, etc.
	Fax string `json:"fax,omitempty"` // Fax number of the address
	Country string `json:"country,omitempty"` // country code according to ISO 3166-1 alpha-2.
	Street_number string `json:"street_number,omitempty"` // Street number
	Website string `json:"website,omitempty"` // Website of the address
}

// CollectionTicketComment represents the CollectionTicketComment schema from the OpenAPI specification
type CollectionTicketComment struct {
	Created_by string `json:"created_by,omitempty"` // The user who created the object.
	Custom_mappings CustomMappings `json:"custom_mappings,omitempty"` // When custom mappings are configured on the resource, the result is included here.
	Id string `json:"id,omitempty"` // A unique identifier for an object.
	Updated_at string `json:"updated_at,omitempty"` // The date and time when the object was last updated.
	Body string `json:"body,omitempty"` // Body of the comment
	Created_at string `json:"created_at,omitempty"` // The date and time when the object was created.
}

// TooManyRequestsResponse represents the TooManyRequestsResponse schema from the OpenAPI specification
type TooManyRequestsResponse struct {
	Ref string `json:"ref,omitempty"` // Link to documentation of error type
	Status_code float64 `json:"status_code,omitempty"` // HTTP status code
	Type_name string `json:"type_name,omitempty"` // The type of error returned
	Detail map[string]interface{} `json:"detail,omitempty"`
	ErrorField string `json:"error,omitempty"` // Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 6585)
	Message string `json:"message,omitempty"` // A human-readable message providing more details about the error.
}

// NotFoundResponse represents the NotFoundResponse schema from the OpenAPI specification
type NotFoundResponse struct {
	Message string `json:"message,omitempty"` // A human-readable message providing more details about the error.
	Ref string `json:"ref,omitempty"` // Link to documentation of error type
	Status_code float64 `json:"status_code,omitempty"` // HTTP status code
	Type_name string `json:"type_name,omitempty"` // The type of error returned
	Detail interface{} `json:"detail,omitempty"` // Contains parameter or domain specific information related to the error and why it occurred.
	ErrorField string `json:"error,omitempty"` // Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 7231)
}

// GetCollectionsResponse represents the GetCollectionsResponse schema from the OpenAPI specification
type GetCollectionsResponse struct {
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data []Collection `json:"data"`
	Links Links `json:"links,omitempty"` // Links to navigate to previous or next pages through the API
	Meta Meta `json:"meta,omitempty"` // Response metadata
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
}

// CustomField represents the CustomField schema from the OpenAPI specification
type CustomField struct {
	Name string `json:"name,omitempty"` // Name of the custom field.
	Value interface{} `json:"value,omitempty"`
	Description string `json:"description,omitempty"` // More information about the custom field
	Id string `json:"id"` // Unique identifier for the custom field.
}

// UnprocessableResponse represents the UnprocessableResponse schema from the OpenAPI specification
type UnprocessableResponse struct {
	ErrorField string `json:"error,omitempty"` // Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 7231)
	Message string `json:"message,omitempty"` // A human-readable message providing more details about the error.
	Ref string `json:"ref,omitempty"` // Link to documentation of error type
	Status_code float64 `json:"status_code,omitempty"` // HTTP status code
	Type_name string `json:"type_name,omitempty"` // The type of error returned
	Detail string `json:"detail,omitempty"` // Contains parameter or domain specific information related to the error and why it occurred.
}

// CreateCommentResponse represents the CreateCommentResponse schema from the OpenAPI specification
type CreateCommentResponse struct {
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
}
