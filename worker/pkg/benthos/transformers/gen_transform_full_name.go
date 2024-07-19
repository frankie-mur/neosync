
// Code generated by Neosync neosync_transformer_generator.go. DO NOT EDIT.
// source: transform_full_name.go

package transformers

import (
	"fmt"
	
	transformer_utils "github.com/nucleuscloud/neosync/worker/pkg/benthos/transformers/utils"
	"github.com/nucleuscloud/neosync/worker/pkg/rng"
	
)

type TransformFullName struct{}

type TransformFullNameOpts struct {
	randomizer     rng.Rand
	
	maxLength int64
	preserveLength bool
}

func NewTransformFullName() *TransformFullName {
	return &TransformFullName{}
}

func (t *TransformFullName) GetJsTemplateData() (*TemplateData, error) {
	return &TemplateData{
		Name: "transformFullName",
		Description: "",
		Example: "",
	}, nil
}

func (t *TransformFullName) ParseOptions(opts map[string]any) (any, error) {
	transformerOpts := &TransformFullNameOpts{}

	maxLength, ok := opts["maxLength"].(int64)
	if !ok {
		maxLength = 10000
	}
	transformerOpts.maxLength = maxLength

	preserveLength, ok := opts["preserveLength"].(bool)
	if !ok {
		preserveLength = false
	}
	transformerOpts.preserveLength = preserveLength

	var seed int64
	seedArg, ok := opts["seed"].(int64)
	if ok {
		seed = seedArg
	} else {
		var err error
		seed, err = transformer_utils.GenerateCryptoSeed()
		if err != nil {
			return nil, fmt.Errorf("unable to generate seed: %w", err)
		}
	}
	transformerOpts.randomizer = rng.New(seed)

	return transformerOpts, nil
}
