package applyer

func ApplyMiddlewares(c func(), m ...func(func())) {
	for _, v := range m {
		v(c)
	}
}
