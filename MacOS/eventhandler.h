#import <Cocoa/Cocoa.h>

@interface EventHandler : NSObject

-(id) initWithUrlString: (NSString*) urlString;
-(IBAction) onOpenClicked:(id) sender;

@end

@interface EventHandler()

@property NSString* urlString;

@end

@implementation EventHandler

-(id) initWithUrlString: (NSString*) urlString
{
	 self = [super init];
	_urlString = urlString;
	return self;
}

-(IBAction) onOpenClicked:(id) sender
{
	[[NSWorkspace sharedWorkspace] openURL:[NSURL URLWithString: _urlString]];
}

@end