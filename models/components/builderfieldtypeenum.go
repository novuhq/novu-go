// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

type BuilderFieldTypeEnum string

const (
	BuilderFieldTypeEnumBoolean   BuilderFieldTypeEnum = "BOOLEAN"
	BuilderFieldTypeEnumText      BuilderFieldTypeEnum = "TEXT"
	BuilderFieldTypeEnumDate      BuilderFieldTypeEnum = "DATE"
	BuilderFieldTypeEnumNumber    BuilderFieldTypeEnum = "NUMBER"
	BuilderFieldTypeEnumStatement BuilderFieldTypeEnum = "STATEMENT"
	BuilderFieldTypeEnumList      BuilderFieldTypeEnum = "LIST"
	BuilderFieldTypeEnumMultiList BuilderFieldTypeEnum = "MULTI_LIST"
	BuilderFieldTypeEnumGroup     BuilderFieldTypeEnum = "GROUP"
)

func (e BuilderFieldTypeEnum) ToPointer() *BuilderFieldTypeEnum {
	return &e
}
func (e *BuilderFieldTypeEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "BOOLEAN":
		fallthrough
	case "TEXT":
		fallthrough
	case "DATE":
		fallthrough
	case "NUMBER":
		fallthrough
	case "STATEMENT":
		fallthrough
	case "LIST":
		fallthrough
	case "MULTI_LIST":
		fallthrough
	case "GROUP":
		*e = BuilderFieldTypeEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for BuilderFieldTypeEnum: %v", v)
	}
}
