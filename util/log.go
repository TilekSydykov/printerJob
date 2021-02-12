package util

func WriteError(txt string) {

}

func HandleError(err error) {
	if err != nil {
		println(err)
	}
}
