// traymanager_macos.go

// +build darwin

package main

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa
#import <Cocoa/Cocoa.h>

int StartApp(void) {
	printf("StartApp");

    [NSAutoreleasePool new];
    id app = [NSApplication sharedApplication];

	NSStatusBar *bar = [NSStatusBar systemStatusBar];

	id onOpenClicked = [^{
		[[NSWorkspace sharedWorkspace] openURL:[NSURL URLWithString: @"http://ya.ru"]];
	} autorelease];

	id menu = [[NSMenu new] autorelease];

    id openMenuItem = [[[NSMenuItem alloc]
		initWithTitle:@"Open TeamCity"
        action:@selector(invoke)
		keyEquivalent:@"o"]
          	autorelease];
	[openMenuItem setTarget: onOpenClicked];
    [menu addItem:openMenuItem];

    id quitMenuItem = [[[NSMenuItem alloc]
		initWithTitle:@"Quit"
        action:@selector(terminate:)
		keyEquivalent:@"q"]
          	autorelease];
    [menu addItem:quitMenuItem];

    NSStatusItem* theItem = [bar statusItemWithLength:NSVariableStatusItemLength];
	[theItem autorelease];

    [theItem setTitle: NSLocalizedString(@"TeamCity",@"")];
    [theItem setHighlightMode:YES];
	[theItem setMenu: menu];

    [NSApp run];
    return 0;
}
*/
import "C"

func StartAppUi() {
	C.StartApp()
}
