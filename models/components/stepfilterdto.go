// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

type Value string

const (
	ValueAnd Value = "AND"
	ValueOr  Value = "OR"
)

func (e Value) ToPointer() *Value {
	return &e
}
func (e *Value) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "AND":
		fallthrough
	case "OR":
		*e = Value(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Value: %v", v)
	}
}

type StepFilterDto struct {
	IsNegated bool                 `json:"isNegated"`
	Type      BuilderFieldTypeEnum `json:"type"`
	Value     Value                `json:"value"`
	Children  []FieldFilterPartDto `json:"children"`
}

func (o *StepFilterDto) GetIsNegated() bool {
	if o == nil {
		return false
	}
	return o.IsNegated
}

func (o *StepFilterDto) GetType() BuilderFieldTypeEnum {
	if o == nil {
		return BuilderFieldTypeEnum("")
	}
	return o.Type
}

func (o *StepFilterDto) GetValue() Value {
	if o == nil {
		return Value("")
	}
	return o.Value
}

func (o *StepFilterDto) GetChildren() []FieldFilterPartDto {
	if o == nil {
		return []FieldFilterPartDto{}
	}
	return o.Children
}
