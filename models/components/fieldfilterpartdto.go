// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

type Operator string

const (
	OperatorLarger       Operator = "LARGER"
	OperatorSmaller      Operator = "SMALLER"
	OperatorLargerEqual  Operator = "LARGER_EQUAL"
	OperatorSmallerEqual Operator = "SMALLER_EQUAL"
	OperatorEqual        Operator = "EQUAL"
	OperatorNotEqual     Operator = "NOT_EQUAL"
	OperatorAllIn        Operator = "ALL_IN"
	OperatorAnyIn        Operator = "ANY_IN"
	OperatorNotIn        Operator = "NOT_IN"
	OperatorBetween      Operator = "BETWEEN"
	OperatorNotBetween   Operator = "NOT_BETWEEN"
	OperatorLike         Operator = "LIKE"
	OperatorNotLike      Operator = "NOT_LIKE"
	OperatorIn           Operator = "IN"
)

func (e Operator) ToPointer() *Operator {
	return &e
}
func (e *Operator) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "LARGER":
		fallthrough
	case "SMALLER":
		fallthrough
	case "LARGER_EQUAL":
		fallthrough
	case "SMALLER_EQUAL":
		fallthrough
	case "EQUAL":
		fallthrough
	case "NOT_EQUAL":
		fallthrough
	case "ALL_IN":
		fallthrough
	case "ANY_IN":
		fallthrough
	case "NOT_IN":
		fallthrough
	case "BETWEEN":
		fallthrough
	case "NOT_BETWEEN":
		fallthrough
	case "LIKE":
		fallthrough
	case "NOT_LIKE":
		fallthrough
	case "IN":
		*e = Operator(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Operator: %v", v)
	}
}

type On string

const (
	OnSubscriber On = "subscriber"
	OnPayload    On = "payload"
)

func (e On) ToPointer() *On {
	return &e
}
func (e *On) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "subscriber":
		fallthrough
	case "payload":
		*e = On(v)
		return nil
	default:
		return fmt.Errorf("invalid value for On: %v", v)
	}
}

type FieldFilterPartDto struct {
	Field    string   `json:"field"`
	Value    string   `json:"value"`
	Operator Operator `json:"operator"`
	On       On       `json:"on"`
}

func (o *FieldFilterPartDto) GetField() string {
	if o == nil {
		return ""
	}
	return o.Field
}

func (o *FieldFilterPartDto) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

func (o *FieldFilterPartDto) GetOperator() Operator {
	if o == nil {
		return Operator("")
	}
	return o.Operator
}

func (o *FieldFilterPartDto) GetOn() On {
	if o == nil {
		return On("")
	}
	return o.On
}
