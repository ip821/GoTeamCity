// traymanager_macos.go

// +build darwin

package main

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa
#import <Cocoa/Cocoa.h>

@interface EventHandler : NSObject <NSApplicationDelegate>

- (IBAction)onOpenedClicked:(id)sender;

@end

@implementation EventHandler

- (IBAction)onOpenedClicked:(id)sender
{
	NSString*myurl=@"ya.ru";
  	NSURL *url = [NSURL URLWithString:myurl];
	[[NSApplication sharedApplication] openURL: url];
}

@end

int StartApp(void) {
	printf("StartApp");

    [NSAutoreleasePool new];
    [NSApplication sharedApplication];

	NSStatusBar *bar = [NSStatusBar systemStatusBar];

	id eventHandler = [[EventHandler alloc] autorelease];
	id menu = [[NSMenu new] autorelease];

    id openMenuItem = [[[NSMenuItem alloc]
		initWithTitle:@"Open TeamCity"
        action:@selector(onOpenedClicked:)
		keyEquivalent:@"o"]
          	autorelease];
	[openMenuItem setTarget: eventHandler];
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
