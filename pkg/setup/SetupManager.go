package setup

type ParamType int

const ParamString ParamType = 1

type SetupParam struct {
	Summary    string         /* Parameter summary     */
	Section    string         /* Parameter section     */
	Name       string         /* Parameter name        */
	Value      string         /* Parameter value       */
	IsSet      bool           /* Parameter exists mark */
	Type       ParamType      /* Parameter value type  */
}

type SetupManager struct {
	Path     string  /* Param stroage path */
	Params []Param   /* Param array        */
}

func NewSetupManager() (*SetupManager) {
	sm := new(SetupManager)
	sm.Path = "~/.golden.sqlite3"
	return sm
}

func (self *SetupManager) Set(name string, value string) (error) {
	return nil
}

func (self *SetupManager) Get(name string, defaultValue string) (value string, error) {
	return "", nil
}

func (self *SetupManager) Register(name string, summary string) (error) {
	return nil
}

func (self *SetupManager) Audit(msg string) (error) {

	/* Store audit message in parameter storage */

	return nil

}

func (self *SetupManager) Restore() (error) {

	/* Open SQL storage */
	db, err1 := sql.Open("sqlite3", self.Path)
		return nil, err
	}
	defer db.Close()

	/* Restore parameters */

	return nil
}

func (self *SetupManager) Store() (error) {

	/* Open SQL storage */
	db, err1 := sql.Open("sqlite3", self.Path)
		return nil, err
	}
	defer db.Close()

	/* Store parameters */

	return nil

}
