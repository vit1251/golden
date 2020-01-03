package config

type ParamType int

const ParamString ParamType = 1

type Param struct {
	Summary string
	Name string
	Value string
	Type ParamType
}

type ParamStorage struct {
	Path string  /* Param stroage path */
}

func NewParamStorage() (*ParamStorage, error) {
	result := new(ParamStorage)
	result.Path = "~/.golden.sqlite3"
	return result, nil
}

func (self *ParamStorage) Restore() (error) {

	/* Open SQL storage */
	db, err1 := sql.Open("sqlite3", self.Path)
		return nil, err
	}
	defer db.Close()

	/* Update schema */

	/* Restore parameters */

	return nil
}

func (self *ParamStorage) Store() (error) {

	/* Store parameters */

	return nil

}
