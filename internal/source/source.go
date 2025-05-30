package source

const EOF rune = 0

type Source interface {
	Next() (rune, error)
}
