//go:build tag1

package foo

//func PlatformSpecificFunction() {
//	fmt.Println("This is the Darwin implementation.")
//}

func init() {
	TAG = append(TAG, "Tag1")
}
