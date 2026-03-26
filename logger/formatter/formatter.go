package formatter

type Formatter interface {
	Format(string, string) any
}
