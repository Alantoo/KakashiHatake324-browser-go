package browsergo

import (
	"fmt"
	"strings"
)

const (
	DisableGPU                                    FlagType = "--disable-gpu"
	PurgeMemoryButton                             FlagType = "--purge-memory-button"
	DisableTranslate                              FlagType = "--disable-translate"
	NoSandbox                                     FlagType = "--no-sandbox"
	DisableBackgroundMode                         FlagType = "--disable-background-mode"
	DisableBlinkAutomationControlled              FlagType = "--disable-blink-features=AutomationControlled"
	HideCrashRestoreBubble                        FlagType = "--hide-crash-restore-bubble"
	HideSavePasswordBubble                        FlagType = "--hide-save-password-bubble"
	RestoreLastSession                            FlagType = "--restore-last-session"
	DisableSync                                   FlagType = "--disable-sync"
	EnableSmoothScrolling                         FlagType = "--enable-smooth-scrolling"
	TurnOffStreamingMediaCaching                  FlagType = "--turn-off-streaming-media-caching"
	EnableQuic                                    FlagType = "--enable-quic"
	EnableZeroCopy                                FlagType = "--enable-zero-copy"
	EnableGPUResterization                        FlagType = "--enable-gpu-rasterization"
	NoFirstRun                                    FlagType = "--no-first-run"
	TestType                                      FlagType = "--test-type"
	DisableBackgroundNetworking                   FlagType = "--disable-background-networking"
	EnableRemoteDebugging                         FlagType = "--enable-remote-debugging"
	MuteAudio                                     FlagType = "--mute-audio"
	DisableBackgroundTimerThrottling              FlagType = "--disable-background-timer-throttling"
	DisableBackgroundOccludedWindows              FlagType = "--disable-backgrounding-occluded-windows"
	DisableComponentExtensionsWithBackgroundPages FlagType = "--disable-component-extensions-with-background-pages"
	NoDefaultBrowserCheck                         FlagType = "--no-default-browser-check"
	DisableRendererBackgrounding                  FlagType = "--disable-renderer-backgrounding"
	DisableInProcessStackTraces                   FlagType = "--disable-in-process-stack-traces"
)

// generate random window size flag
func RandomWindowSize() FlagType {
	return FlagType(fmt.Sprintf("--window-size=%d,%d", RandomInt(800, 1400), RandomInt(600, 1000)))
}

// generate the disable blink features flag
func DisableBlinkFeatures(features []string) FlagType {
	return FlagType(fmt.Sprintf("--disable-blink-features=%s", strings.Join(features, ",")))
}

// generate the enable blink features flag
func EnableBlinkFeatures(features []string) FlagType {
	return FlagType(fmt.Sprintf("--enable-blink-features=%s", strings.Join(features, ",")))
}

// generate the enable features flag
func EnableFeatures(features []string) FlagType {
	return FlagType(fmt.Sprintf("--enable-features=%s", strings.Join(features, ",")))
}

// generate the disable features flag
func DisableFeatures(features []string) FlagType {
	return FlagType(fmt.Sprintf("--disable-features=%s", strings.Join(features, ",")))
}

func StackTraceLimit(limit int) FlagType {
	return FlagType(fmt.Sprintf("--stack-trace-limit %d", limit))
}

func StackTraceLimitV8(limit int) FlagType {
	return FlagType(fmt.Sprintf("--js-flags=\x27--stack-trace-limit %d\x27", limit))
}

func SetUserAgent(ua string) FlagType {
	return FlagType(fmt.Sprintf("--user-agent=%s", ua))
}
