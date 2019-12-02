package miniblink

//#include "event.h"
import "C"
import (
	"unsafe"
)

//export goOnWindowDestroyCallback
func goOnWindowDestroyCallback(window C.wkeWebView, param unsafe.Pointer) {
	go func() {
		view := getWebViewByWindow(window)
		view.Emit("destroy", view)
	}()
}

//export goOnDocumentReadyCallback
func goOnDocumentReadyCallback(window C.wkeWebView) {
	go func() {
		view := getWebViewByWindow(window)
		view.Emit("documentReady", view)
	}()
}

//export goOnTitleChangedCallback
func goOnTitleChangedCallback(window C.wkeWebView, titleString *C.char) {
	//把C过来的字符串转化为golang的
	titleGoString := C.GoString(titleString)

	go func() {
		view := getWebViewByWindow(window)
		view.Emit("titleChanged", view, titleGoString)
	}()
}
//export goOnUrlLoadBeginCheck
func goOnUrlLoadBeginCheck(window C.wkeWebView, url *C.char) bool {
	//把C过来的字符串转化为golang的
	urlGoString := C.GoString(url)
	view := getWebViewByWindow(window)
	if !view.isRequestAllowed(urlGoString) {
		view.Emit("requestBlocked", view, urlGoString)
		return true
	}
	return false
}
