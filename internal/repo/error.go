package repo

type RepoError interface {
	error
}

type DeatachedHead struct {
	error
}
