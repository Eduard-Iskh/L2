package sorted

// reverse переворачивает слайс.
//
// Первый элемент меняется местами с последним,
// второй с предпоследним и т.д.
//
// Используются generics,
// поэтому функция работает с любым типом.
func Reverse[T any](arr []T) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
