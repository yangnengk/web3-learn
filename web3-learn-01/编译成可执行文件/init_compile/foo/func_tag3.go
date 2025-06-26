//go:build tag3 && !tag2

package foo

//func PlatformSpecificFunction() {
//	fmt.Println("This is the Windows implementation.")
//}

func init() {
	TAG = append(TAG, "Tag3")
}
