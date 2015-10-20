#pragma once

int StartApp(const char* strUrl) {
	printf("StartApp");

    [NSAutoreleasePool new];
    id app = [NSApplication sharedApplication];

	NSString* url = [[[NSString alloc] initWithUTF8String: strUrl] autorelease];
	EventHandler* eventHandler = [[[EventHandler alloc] initWithUrlString: url] autorelease];

	NSStatusBar *bar = [NSStatusBar systemStatusBar];

	id menu = [[NSMenu new] autorelease];

    id openMenuItem = [[[NSMenuItem alloc]
		initWithTitle:@"Open TeamCity"
        action:@selector(onOpenClicked:)
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
