
// Code generated by Neosync neosync_transformer_generator.go. DO NOT EDIT.
// source: transform_string_phone_number.go

package transformers

import (
	"fmt"
	
)

type TransformStringPhoneNumber struct{}

type TransformStringPhoneNumberOpts struct {
	preserveLength bool
	maxLength int64
}

func NewTransformStringPhoneNumber() *TransformStringPhoneNumber {
	return &TransformStringPhoneNumber{}
}

func (t *TransformStringPhoneNumber) GetJsTemplateData() (*TemplateData, error) {
	return &TemplateData{
		Name: "transformStringPhoneNumber",
		Description: "",
		Example: "",
	}, nil
}

func (t *TransformStringPhoneNumber) ParseOptions(opts map[string]any) (any, error) {
	transformerOpts := &TransformStringPhoneNumberOpts{}

	if _, ok := opts["preserveLength"].(bool); !ok {
		return nil, fmt.Errorf("missing required argument. function: %s argument: %s", "transformStringPhoneNumber", "preserveLength")
	}
	preserveLength := opts["preserveLength"].(bool)
	transformerOpts.preserveLength = preserveLength

	if _, ok := opts["maxLength"].(int64); !ok {
		return nil, fmt.Errorf("missing required argument. function: %s argument: %s", "transformStringPhoneNumber", "maxLength")
	}
	maxLength := opts["maxLength"].(int64)
	transformerOpts.maxLength = maxLength

	return transformerOpts, nil
}
