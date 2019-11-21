package jobspec

import (
	batchv1 "k8s.io/api/batch/v1"
)

type JobSpec struct {
	Object *batchv1.JobSpec
}
