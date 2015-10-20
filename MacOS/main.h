#pragma once

NSStatusItem* mainItem;
EventHandler* eventHandler;

int StartApp(const char* strUrl) {
	printf("StartApp");

    [NSAutoreleasePool new];
    id app = [NSApplication sharedApplication];

	NSString* url = [[[NSString alloc] initWithUTF8String: strUrl] autorelease];
	eventHandler = [[[EventHandler alloc] initWithUrlString: url] autorelease];

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

    mainItem = [bar statusItemWithLength:NSVariableStatusItemLength];
	[mainItem autorelease];

	[eventHandler setStatusItem: mainItem];

    [mainItem setTitle: NSLocalizedString(@"TeamCity",@"")];
    [mainItem setHighlightMode:YES];
	[mainItem setMenu: menu];

    return 0;
}

void RunApp(){
	[NSApp run];
}

void UpdateAppUi(int status){
	dispatch_async(dispatch_get_main_queue(), ^
    {
		[eventHandler update: status];
    });
}