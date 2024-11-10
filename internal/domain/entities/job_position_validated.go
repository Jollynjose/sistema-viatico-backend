package entities

type JobPositionValidated struct {
	JobPosition JobPosition
	isValidated bool
	Err         error
}

func (j *JobPositionValidated) IsValidated() bool {
	return j.isValidated
}

func NewJobPostionValidated(j *JobPosition) *JobPositionValidated {
	if err := j.validate(); err != nil {
		return &JobPositionValidated{
			JobPosition: *j,
			isValidated: false,
			Err:         err,
		}
	}
	return &JobPositionValidated{
		JobPosition: *j,
		isValidated: true,
	}
}
