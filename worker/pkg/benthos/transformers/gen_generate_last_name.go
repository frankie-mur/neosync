
// Code generated by Neosync neosync_transformer_generator.go. DO NOT EDIT.
// source: generate_last_name.go

package transformers

import (
	"fmt"
	
	transformer_utils "github.com/nucleuscloud/neosync/worker/pkg/benthos/transformers/utils"
	"github.com/nucleuscloud/neosync/worker/pkg/rng"
	
)

type GenerateLastName struct{}

type GenerateLastNameOpts struct {
	randomizer     rng.Rand
	
	maxLength int64
}

func NewGenerateLastName() *GenerateLastName {
	return &GenerateLastName{}
}

func (t *GenerateLastName) GetJsTemplateData() (*TemplateData, error) {
	return &TemplateData{
		Name: "generateLastName",
		Description: "",
		Example: "",
	}, nil
}

func (t *GenerateLastName) ParseOptions(opts map[string]any) (any, error) {
	transformerOpts := &GenerateLastNameOpts{}

	maxLength, ok := opts["maxLength"].(int64)
	if !ok {
		maxLength = 10000
	}
	transformerOpts.maxLength = maxLength

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
