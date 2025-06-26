//go:build tag2

package foo

//func PlatformSpecificFunction() {
//	fmt.Println("This is the Linux implementation.")
//}

func init() {
	TAG = append(TAG, "Tag2")
}
