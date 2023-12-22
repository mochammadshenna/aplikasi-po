package common

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strings"

	"github.com/mochammadshenna/aplikasi-po/internal/util/logger"
)

func GetAppName() string {
	appName := "pba-lambda"
	env := os.Getenv("ENV")
	if env == "local" {
		return appName + "-local"
	} else if env == "prod" {
		domain := os.Getenv("DOMAIN")
		if len(domain) > 0 {
			appName = fmt.Sprintf("%s-%s-prod", appName, domain) // appName + domain + "-prod"
		} else {
			appName = fmt.Sprintf("%s-prod", appName) // appName + "-prod"
		}

		return appName
	}
	return ""
}

func FuncCallerName(index int) string {
	const (
		allocatePtr = 15
		slashRune   = '/'
	)

	pc := make([]uintptr, allocatePtr)

	n := runtime.Callers(index, pc)

	frames := runtime.CallersFrames(pc[:n])
	_, _ = frames.Next()

	f := runtime.FuncForPC(pc[0])
	fName := f.Name()

	var lastSlash int

	if strings.Contains(fName, string(slashRune)) {
		lastSlash = strings.LastIndexByte(fName, slashRune)
		return fName[lastSlash+1:]
	}
	return fName
}

var uintPattern = regexp.MustCompile(`^\d+$`)

func IsUnsignedInt(s string) bool {
	return uintPattern.MatchString(s)
}

func IsAlphabet(b rune) bool {
	return b >= 'a' && b <= 'z' ||
		b >= 'A' && b <= 'Z'
}

func IsNumeric(b rune) bool {
	return b >= '0' && b <= '9'
}

func IsAlphanumeric(b rune) bool {
	return IsAlphabet(b) || IsNumeric(b)
}

func ResizeImage(img string, width, height int, location string) string {
	if img == "" {
		return ""
	}

	// get url image resizer
	url_image := os.Getenv("url_image_resizer")
	if url_image == "" {
		url_image = "https://images.archipelagohotels.com/"
	}

	// Get Bucket Name from image URL
	bucketName := ""
	r := regexp.MustCompile(`^(?:https?://)?(?:www\.)?([^/]+)`)
	matches := r.FindStringSubmatch(img)
	if len(matches) >= 2 {
		urls := strings.Split(matches[1], ".")
		if len(urls) >= 2 {
			bucketName = urls[1]
		}
	}

	baseURL := url_image + bucketName + "/"
	cdn := strings.Split(img, ".")
	trim := strings.Replace(img, cdn[0]+"."+cdn[1]+"."+"com/", "", 1)

	var link string

	if width != 0 && height == 0 {
		link = baseURL + trim + "?s=" + fmt.Sprint(width) + "&location=" + location
	} else if width == 0 || height == 0 {
		link = baseURL + trim + "?location=" + location
	} else {
		size := fmt.Sprintf("%dx%d", width, height)
		link = baseURL + trim + "?d=" + size + "&location=" + location
	}

	return link
}

// detect is running inside container
func IsRunningInsideContainer() bool {
	// "cat /proc/1/sched | head -n 1" This command extracts the first line of the scheduling overview for the process with PID 1.
	// For the Linux containers, it displays the command of the main process. For Ubuntu or CentOS, it’s bash.
	// However, if the process isn’t in a container, it displays init – the initialization process of the OS
	// In some Linux distributions like CentOS and Debian, the initialization process is systemd. So, for them, it shows system instead of init.
	// ref: https://www.baeldung.com/linux/is-process-running-inside-container#using-cpu-scheduling-info
	cmd := "cat /proc/1/sched | head -n 1"
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		logger.Logger.Info("err ", err.Error())
		return false
	}

	isBash := strings.Contains(string(out), "bash")
	isMain := strings.Contains(string(out), "main")
	isBlank := string(out) == ""

	if isBash || isMain || isBlank {
		return true
	}

	return false
}
