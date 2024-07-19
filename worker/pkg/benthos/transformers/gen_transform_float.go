
// Code generated by Neosync neosync_transformer_generator.go. DO NOT EDIT.
// source: transform_float.go

package transformers

import (
	"fmt"
	
	transformer_utils "github.com/nucleuscloud/neosync/worker/pkg/benthos/transformers/utils"
	"github.com/nucleuscloud/neosync/worker/pkg/rng"
	
)

type TransformFloat64 struct{}

type TransformFloat64Opts struct {
	randomizer     rng.Rand
	
	randomizationRangeMin float64
	randomizationRangeMax float64
	precision *int64
	scale *int64
}

func NewTransformFloat64() *TransformFloat64 {
	return &TransformFloat64{}
}

func (t *TransformFloat64) GetJsTemplateData() (*TemplateData, error) {
	return &TemplateData{
		Name: "transformFloat64",
		Description: "",
		Example: "",
	}, nil
}

func (t *TransformFloat64) ParseOptions(opts map[string]any) (any, error) {
	transformerOpts := &TransformFloat64Opts{}

	if _, ok := opts["randomizationRangeMin"].(float64); !ok {
		return nil, fmt.Errorf("missing required argument. function: %s argument: %s", "transformFloat64", "randomizationRangeMin")
	}
	randomizationRangeMin := opts["randomizationRangeMin"].(float64)
	transformerOpts.randomizationRangeMin = randomizationRangeMin

	if _, ok := opts["randomizationRangeMax"].(float64); !ok {
		return nil, fmt.Errorf("missing required argument. function: %s argument: %s", "transformFloat64", "randomizationRangeMax")
	}
	randomizationRangeMax := opts["randomizationRangeMax"].(float64)
	transformerOpts.randomizationRangeMax = randomizationRangeMax

	var precision *int64
	if arg, ok := opts["precision"].(int64); ok {
		precision = &arg
	}
	transformerOpts.precision = precision

	var scale *int64
	if arg, ok := opts["scale"].(int64); ok {
		scale = &arg
	}
	transformerOpts.scale = scale

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
