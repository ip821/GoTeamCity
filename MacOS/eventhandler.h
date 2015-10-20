#import <Cocoa/Cocoa.h>

enum StatusEnum
{
	BSuccess = 1,
	BFailed = 2,
	BFailedAndUserChangesFound = 3,
	BAssigned = 4
};

@interface EventHandler : NSObject

-(id) initWithUrlString: (NSString*) urlString;
-(void) setMainItem:(NSStatusItem*) mainItem;
-(IBAction) onOpenClicked:(id) sender;
-(void) update:(int) status;

@end

@interface EventHandler()

@property NSString* urlString;
@property NSStatusItem* mainItem;

@end

@implementation EventHandler

-(id) initWithUrlString: (NSString*) urlString
{
	 self = [super init];
	_urlString = urlString;
	return self;
}

-(void) setStatusItem:(NSStatusItem*) item
{
	_mainItem = item;
}

-(IBAction) onOpenClicked:(id) sender
{
	[[NSWorkspace sharedWorkspace] openURL:[NSURL URLWithString: _urlString]];
}

-(void) update:(int) status
{
	switch(status)
	{
		case BSuccess:
			[_mainItem setTitle: NSLocalizedString(@"TC - Success",@"")];
			break;
		case BFailed:
			[_mainItem setTitle: NSLocalizedString(@"TC - Failed",@"")];
			break;
		case BFailedAndUserChangesFound:
			[_mainItem setTitle: NSLocalizedString(@"TC - Failed(CHANGES)",@"")];
			break;
		case BAssigned:
			[_mainItem setTitle: NSLocalizedString(@"TC - ASSIGNED",@"")];
			break;
	}
}

@end