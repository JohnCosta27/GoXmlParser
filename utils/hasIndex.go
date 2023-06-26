package utils

func HasIndex[T any, K int](arr []T, index K) bool {
  return len(arr) < int(index)
}
