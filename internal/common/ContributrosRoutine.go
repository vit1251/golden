package commonfunc

type Contributor struct {
	Name string
	Addr string
}

func GetContributors() []Contributor {

	var contributors []Contributor

	contributors = append(contributors, Contributor{
		Name: "Sergey Anohin",
	})
	contributors = append(contributors, Contributor{
		Name: "Andrey Mundirov",
	})
	contributors = append(contributors, Contributor{
		Name: "Jaroslav Bespalov",
		Addr: "2:5031/78.17",
	})
	contributors = append(contributors, Contributor{
		Name: "Richard Menedetter",
	})
	contributors = append(contributors, Contributor{
		Name: "Tommi Koivula",
	})
	contributors = append(contributors, Contributor{
		Name: "Rudi Timmermans",
		Addr: "2:292/140",
	})

	return contributors
}
