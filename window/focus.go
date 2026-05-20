package window

import "syscall"

// GetForegroundWindow returns the current foreground window handle.
func GetForegroundWindow() uintptr {
	r, _, _ := ProcGetForegroundWindow.Call()
	return r
}

// SetForegroundWindow brings the specified window to the foreground.
func SetForegroundWindow(hwnd uintptr) error {
	r, _, e := ProcSetForegroundWindow.Call(hwnd)
	if r != 0 {
		return nil
	}
	if errno, ok := e.(syscall.Errno); ok && errno != 0 {
		return errno
	}
	return syscall.EINVAL
}
