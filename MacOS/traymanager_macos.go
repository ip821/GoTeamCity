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

type StatusEnum int

const (
	BSuccess                   StatusEnum = 1
	BFailed                               = 2
	BFailedAndUserChangesFound            = 3
	BAssigned                             = 4
)

type Status struct {
	BuildsStatus StatusEnum
}

func StartAppUi(url string) {
	strUrl := C.CString(url)
	C.StartApp(strUrl)
}

func RunApp() {
	C.RunApp()
}

func UpdateAppUi(status Status) {
	C.UpdateAppUi(C.int(status.BuildsStatus))
}
