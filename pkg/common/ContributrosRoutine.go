package commonfunc

type Contributor struct {
	Name string
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
	})
	contributors = append(contributors, Contributor{
		Name: "Richard Menedetter",
	})
	contributors = append(contributors, Contributor{
		Name: "Tommi Koivula",
	})

	return contributors
}
