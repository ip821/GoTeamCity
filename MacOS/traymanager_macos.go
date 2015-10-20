// traymanager_macos.go

// +build darwin

package macos

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa
#import <Cocoa/Cocoa.h>
#import "eventhandler.h"
#include "main.h"
*/
import "C"

func StartAppUi(url string) {
	strUrl := C.CString(url)
	C.StartApp(strUrl)
}
