package lib

type EditDistance interface {
  Walk(a,b string) (int, error)
}
