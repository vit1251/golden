package area

type Area struct {
	Name            string     /* Echo name              */
	Summary         string     /* Echo summary           */
	Charset         string     /* Echo charset           */
	Flag            string     /* Echo marker            */
	Path            string     /* Echo directory         */
	MessageCount    int        /* Echo message count     */
	NewMessageCount int        /* Echo new message count */
}

func NewArea() *Area {
	a := new(Area)
	a.Charset = "CP866"
	a.Summary = "Без описания"
	a.Flag = "A"
	return a
}
