package command

type CreateJobPositionCommand struct {
	Name                 string
	JobPositionHistories []CreateJobPositionHistoryCommand
}
