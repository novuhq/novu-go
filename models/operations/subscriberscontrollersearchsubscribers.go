// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/novuhq/novu-go/models/components"
)

// OrderDirection - Direction of sorting
type OrderDirection string

const (
	OrderDirectionAsc  OrderDirection = "ASC"
	OrderDirectionDesc OrderDirection = "DESC"
)

func (e OrderDirection) ToPointer() *OrderDirection {
	return &e
}
func (e *OrderDirection) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ASC":
		fallthrough
	case "DESC":
		*e = OrderDirection(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OrderDirection: %v", v)
	}
}

type SubscribersControllerSearchSubscribersRequest struct {
	// Cursor for pagination indicating the starting point after which to fetch results.
	After *string `queryParam:"style=form,explode=true,name=after"`
	// Cursor for pagination indicating the ending point before which to fetch results.
	Before *string `queryParam:"style=form,explode=true,name=before"`
	// Limit the number of items to return
	Limit *float64 `queryParam:"style=form,explode=true,name=limit"`
	// Direction of sorting
	OrderDirection *OrderDirection `queryParam:"style=form,explode=true,name=orderDirection"`
	// Field to order by
	OrderBy *string `queryParam:"style=form,explode=true,name=orderBy"`
	// Include cursor item in response
	IncludeCursor *bool `queryParam:"style=form,explode=true,name=includeCursor"`
	// Email address of the subscriber to filter results.
	Email *string `queryParam:"style=form,explode=true,name=email"`
	// Name of the subscriber to filter results.
	Name *string `queryParam:"style=form,explode=true,name=name"`
	// Phone number of the subscriber to filter results.
	Phone *string `queryParam:"style=form,explode=true,name=phone"`
	// Unique identifier of the subscriber to filter results.
	SubscriberID *string `queryParam:"style=form,explode=true,name=subscriberId"`
	// A header for idempotency purposes
	IdempotencyKey *string `header:"style=simple,explode=false,name=idempotency-key"`
}

func (o *SubscribersControllerSearchSubscribersRequest) GetAfter() *string {
	if o == nil {
		return nil
	}
	return o.After
}

func (o *SubscribersControllerSearchSubscribersRequest) GetBefore() *string {
	if o == nil {
		return nil
	}
	return o.Before
}

func (o *SubscribersControllerSearchSubscribersRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *SubscribersControllerSearchSubscribersRequest) GetOrderDirection() *OrderDirection {
	if o == nil {
		return nil
	}
	return o.OrderDirection
}

func (o *SubscribersControllerSearchSubscribersRequest) GetOrderBy() *string {
	if o == nil {
		return nil
	}
	return o.OrderBy
}

func (o *SubscribersControllerSearchSubscribersRequest) GetIncludeCursor() *bool {
	if o == nil {
		return nil
	}
	return o.IncludeCursor
}

func (o *SubscribersControllerSearchSubscribersRequest) GetEmail() *string {
	if o == nil {
		return nil
	}
	return o.Email
}

func (o *SubscribersControllerSearchSubscribersRequest) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *SubscribersControllerSearchSubscribersRequest) GetPhone() *string {
	if o == nil {
		return nil
	}
	return o.Phone
}

func (o *SubscribersControllerSearchSubscribersRequest) GetSubscriberID() *string {
	if o == nil {
		return nil
	}
	return o.SubscriberID
}

func (o *SubscribersControllerSearchSubscribersRequest) GetIdempotencyKey() *string {
	if o == nil {
		return nil
	}
	return o.IdempotencyKey
}

type SubscribersControllerSearchSubscribersResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	ListSubscribersResponseDto *components.ListSubscribersResponseDto
	Headers                    map[string][]string
}

func (o *SubscribersControllerSearchSubscribersResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *SubscribersControllerSearchSubscribersResponse) GetListSubscribersResponseDto() *components.ListSubscribersResponseDto {
	if o == nil {
		return nil
	}
	return o.ListSubscribersResponseDto
}

func (o *SubscribersControllerSearchSubscribersResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}
