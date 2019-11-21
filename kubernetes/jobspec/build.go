package jobspec

import (
	"errors"
	"fmt"

	templatespec "github.com/litmuschaos/kube-helper/kubernetes/podtemplatespec"
	batchv1 "k8s.io/api/batch/v1"
)

type Builder struct {
	jobspec *JobSpec
	errs    []error
}

func NewBuilder() *Builder {
	return &Builder{
		jobspec: &JobSpec{
			Object: &batchv1.JobSpec{},
		},
	}
}

func (b *Builder) WithBackOffLimit(backoff *int32) *Builder {
	if int(*backoff) < 0 {
		b.errs = append(
			b.errs,
			errors.New("failed to build Job: invalid backofflimit "),
		)
		return b
	}

	b.jobspec.Object.BackoffLimit = backoff
	return b
}

func (b *Builder) WithPodTemplateSpecBuilder(
	tmplbuilder *templatespec.Builder,
) *Builder {
	if tmplbuilder == nil {
		b.errs = append(
			b.errs,
			errors.New("failed to build job: nil templatespecbuilder"),
		)
		return b
	}

	templatespecObj, err := tmplbuilder.Build()

	if err != nil {
		b.errs = append(
			b.errs,
			errors.New(
				"failed to build job"),
		)
		return b
	}
	b.jobspec.Object.Template = *templatespecObj.Object
	return b
}

func (b *Builder) Build() (*JobSpec, error) {
	if len(b.errs) > 0 {
		return nil, fmt.Errorf("%+v", b.errs)
	}
	return b.jobspec, nil
}
