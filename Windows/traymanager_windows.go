// +build windows

package windows

/*
#include <windows.h>
#include "main.h"
*/
import "C"

func StartAppUi(url string) {
	strUrl := C.CString(url)
	C.StartApp(strUrl)
}

func RunApp() {
	C.RunApp()
}

func UpdateAppUi(status int) {
	C.UpdateAppUi(C.int(status))
}
