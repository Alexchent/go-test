package mlog

type Formatter interface {
	// Maybe in async goroutine
	// Please write the result to buffer
	Format(entry *Entry) error
}
