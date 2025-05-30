package source

type Source interface {
	Next() (rune, error)
}
