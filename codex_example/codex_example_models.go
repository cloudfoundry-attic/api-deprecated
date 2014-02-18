package codex_example

type App struct {
	Id      int64
	Name    string
	Guid    string
	SpaceId int64
	StackId int64
}

var AppKeys = []interface{}{"id", "name", "guid", "space_id", "stack_id"}

type Space struct {
	Id int64 `db:"id"`
}

type Stack struct {
	Id int64 `db:"id"`
}

func ZipMap(keys []interface{}, vals []interface{}) (m map[interface{}]interface{}) {
	m = map[interface{}]interface{}{}
	for i, key := range keys {
		m[key] = vals[i]
	}
	return
}
