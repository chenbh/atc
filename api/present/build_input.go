package present

import (
	"github.com/chenbh/concourse/atc"
	"github.com/chenbh/concourse/atc/db"
)

func BuildInput(input db.BuildInput, config atc.JobInputParams, resource db.Resource) atc.BuildInput {
	return atc.BuildInput{
		Name:     input.Name,
		Resource: resource.Name(),
		Type:     resource.Type(),
		Source:   resource.Source(),
		Params:   config.Params,
		Version:  atc.Version(input.Version),
		Tags:     config.Tags,
	}
}
