package agent


type Agent struct {
    Name string
}

func newAgent(name string) Agent {
    return Agent{
        Name: name,
    }
}

