package main

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation -framework AppKit
#import <Foundation/Foundation.h>
#import <AppKit/AppKit.h>

const char* getAppearance() {
    NSAppearance* appearance = [NSAppearance currentAppearance];
    if (appearance.name == NSAppearanceNameAqua) {
        return "Light";
    } else if (appearance.name == NSAppearanceNameDarkAqua) {
        return "Dark";
    } else {
        return "Unknown";
    }
}
*/
import "C"
import "fmt"

func getAppearance() string {
	return C.GoString(C.getAppearance())
}

func main() {
	appearance := getAppearance()
	fmt.Println("System Appearance:", appearance)
}
