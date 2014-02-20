package job

type Repo interface {
	FindByGuid(string) Model
}
