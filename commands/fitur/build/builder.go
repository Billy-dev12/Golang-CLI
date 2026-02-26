package build

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func HandleBuild(targetOS string, customName string) {
	// Jika user tidak kasih nama, pakai default "bill"
	if customName == "" {
		customName = "bill"
	}

	outputName := customName
	if targetOS == "windows" {
		outputName += ".exe"
	}

	cmd := exec.Command("go", "build", "-ldflags", "-s -w", "-o", outputName)

	cmd.Env = append(os.Environ(),
		"GOOS="+targetOS,
		"GOARCH=amd64",
	)

	fmt.Printf("🚀 Merakit [%s] untuk %s...\n", outputName, targetOS)

	err := cmd.Run()
	if err != nil {
		fmt.Printf("❌ Gagal: %v\n", err)
		return
	}

	fmt.Printf("✅ Selesai! File: %s\n", outputName)
}

func ShowCurrentEnv() {
	fmt.Printf("Kamu sekarang lagi di: %s/%s\n", runtime.GOOS, runtime.GOARCH)
}
