package entities

type JobPositionHistoryValidated struct {
	JobPositionHistory JobPositionHistory
	isValidated        bool
}

func (j *JobPositionHistoryValidated) IsValidated() bool {
	return j.isValidated
}

func NewJobPositionHistoryValidated(j *JobPositionHistory) *JobPositionHistoryValidated {
	if err := j.validate(); err != nil {
		return &JobPositionHistoryValidated{
			JobPositionHistory: *j,
			isValidated:        false,
		}
	}
	return &JobPositionHistoryValidated{
		JobPositionHistory: *j,
		isValidated:        true,
	}
}
