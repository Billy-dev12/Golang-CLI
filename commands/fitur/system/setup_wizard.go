package system

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

// ======================================
//        🎨 ASCII ART & BANNER
// ======================================

func printRobot() {
	robot := `
        ╔══════════════════════════════════╗
        ║     ____  _ _ _                  ║
        ║    |  _ \(_) | |                 ║
        ║    | |_) |_| | |_   _            ║
        ║    |  _ <| | | | | | |           ║
        ║    | |_) | | | | |_| |           ║
        ║    |____/|_|_|_|\__, |           ║
        ║                  __/ |           ║
        ║      ASISTEN    |___/            ║
        ╚══════════════════════════════════╝
              ┌──────────────┐
              │   [ O    O ] │
              │      __      │
              │    \____/    │
              └──────┬───────┘
                     │
              ┌──────┴───────┐
              │  ╔════════╗  │
              │  ║ READY! ║  │
              │  ╚════════╝  │
              ├──────────────┤
              │   /      \   │
              └──/────────\──┘
                 │        │
                 ╘════════╛
`
	fmt.Println(robot)
}

func printInstalledBanner() {
	banner := `
    ╔═══════════════════════════════════════════════╗
    ║                                               ║
    ║   ✅  INSTALASI BERHASIL!  ✅                 ║
    ║                                               ║
    ║   Billy Asisten sudah terpasang di sistemmu.  ║
    ║   Sekarang tutup terminal ini, buka lagi,     ║
    ║   dan panggil dari mana saja!                 ║
    ║                                               ║
    ╚═══════════════════════════════════════════════╝
`
	fmt.Println(banner)
}

func printLoading(text string) {
	frames := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	for i := 0; i < 10; i++ {
		fmt.Printf("\r  %s %s", frames[i%len(frames)], text)
		time.Sleep(150 * time.Millisecond)
	}
	fmt.Println()
}

func PrintHelpBox() {
	fmt.Println("    ┌───────────────────────────────────────────────┐")
	fmt.Println("    │  📖  DAFTAR PERINTAH:                        │")
	fmt.Println("    ├───────────────────────────────────────────────┤")
	fmt.Println("    │                                               │")
	fmt.Println("    │  🔨 LARAVEL                                   │")
	fmt.Println("    │  bill buat laravel [nama]  → Buat project    │")
	fmt.Println("    │  bill cek dev              → Cek environment │")
	fmt.Println("    │  bill setup dev            → Pasang tools    │")
	fmt.Println("    │                                               │")
	fmt.Println("    │  🐙 GITHUB                                    │")
	fmt.Println("    │  bill push [link] [pesan]  → Push ke GitHub  │")
	fmt.Println("    │  bill push update [pesan]  → Update GitHub   │")
	fmt.Println("    │                                               │")
	fmt.Println("    │  ⚙️  SISTEM                                    │")
	fmt.Println("    │  bill install              → Pasang ke PATH  │")
	fmt.Println("    │  bill help                 → Tampilkan ini   │")
	fmt.Println("    │                                               │")
	fmt.Println("    └───────────────────────────────────────────────┘")
	fmt.Println()
	fmt.Println("    🌐 https://github.com/Billy-dev12")
	fmt.Println()
}

// IsInstalled mengecek apakah aplikasi sudah terpasang di folder resmi atau PATH
func IsInstalled() bool {
	exePath, err := os.Executable()
	if err != nil {
		return false
	}
	exeDir := filepath.Dir(exePath)
	exeName := filepath.Base(exePath)

	localAppData := os.Getenv("LOCALAPPDATA")
	targetDir := filepath.Join(localAppData, "bill-tool")
	targetPath := filepath.Join(targetDir, exeName)

	// 1. Cek apakah berjalan dari folder resmi
	if equalsIgnoreCase(exeDir, targetDir) {
		return true
	}

	// 2. Cek apakah file sudah ada di folder target
	if _, err := os.Stat(targetPath); err == nil {
		return true
	}

	// 3. Cek apakah ada di PATH
	cmdName := strings.TrimSuffix(exeName, filepath.Ext(exeName))
	if _, err := exec.LookPath(cmdName); err == nil {
		return true
	}

	return false
}

// ShowWelcomeInInteractiveMode menampilkan robot + banner + REPL loop
func ShowWelcomeInInteractiveMode(callback func([]string)) {
	printRobot()

	if !IsInstalled() {
		fmt.Println("    🚀 SEPERTINYA BILLY BELUM TERPASANG DI SISTEM KAMU!")
		fmt.Println("    ---------------------------------------------------")
		fmt.Println("    Silahkan ketik perintah di bawah untuk menginstall:")
		fmt.Println("    > install")
		fmt.Println("    ---------------------------------------------------")
		fmt.Println()
	}

	PrintHelpBox()

	fmt.Println("    💡 Kamu bisa langsung mengetik perintah di sini.")
	fmt.Println("    (Contoh: 'install' atau 'help'. Ketik 'exit' untuk keluar)")
	fmt.Println()

	for {
		fmt.Print("    billy > ")
		var input string
		// Menggunakan scanner untuk menangani spasi
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			input = scanner.Text()
		}

		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}

		// Pecah input menjadi argumen
		args := strings.Fields(input)
		callback(args)
		fmt.Println()
	}
}

// ======================================
//        🔧 LOGIKA INSTALASI
// ======================================

// Install langsung memasang ke bin (AppData) dan PATH tanpa tanya
func Install() {
	exePath, err := os.Executable()
	if err != nil {
		fmt.Printf("    ❌ Gagal mendeteksi lokasi file: %v\n", err)
		Pause()
		return
	}
	exeName := filepath.Base(exePath)
	exeDir := filepath.Dir(exePath)

	localAppData := os.Getenv("LOCALAPPDATA")
	targetDir := filepath.Join(localAppData, "bill-tool")
	targetPath := filepath.Join(targetDir, exeName)

	// Cek apakah sudah dijalankan dari folder resmi
	if equalsIgnoreCase(exeDir, targetDir) {
		fmt.Println("\n    ✅ Billy Asisten sudah terpasang dari folder ini!")
		PrintHelpBox()
		return
	}

	// Cek apakah file sudah ada di folder target
	if _, err := os.Stat(targetPath); err == nil {
		fmt.Println("\n    ✅ Billy Asisten sudah terpasang sebelumnya!")
		fmt.Printf("    📂 Lokasi: %s\n\n", targetPath)
		PrintHelpBox()
		return
	}

	// Mulai proses instalasi!
	fmt.Println()

	// 1. Copy file
	printLoading("Membuat folder instalasi...")
	err = os.MkdirAll(targetDir, 0755)
	if err != nil {
		fmt.Printf("    ❌ Gagal membuat folder: %v\n", err)
		return
	}

	printLoading("Menyalin file ke sistem...")
	err = copyFile(exePath, targetPath)
	if err != nil {
		fmt.Printf("    ❌ Gagal menyalin file: %v\n", err)
		return
	}

	// 2. Daftarkan ke PATH (tanpa duplikat)
	printLoading("Mendaftarkan ke PATH Windows...")
	script := fmt.Sprintf(`
		$oldPath = [Environment]::GetEnvironmentVariable("Path", "User");
		if ($oldPath -notlike "*%s*") {
			[Environment]::SetEnvironmentVariable("Path", $oldPath + ";%s", "User")
		}
	`, targetDir, targetDir)
	cmd := exec.Command("powershell", "-Command", script)
	err = cmd.Run()
	if err != nil {
		fmt.Printf("    ❌ Gagal mendaftarkan PATH: %v\n", err)
		return
	}

	// 3. Tampilkan hasil meriah!
	printRobot()
	printInstalledBanner()
	PrintHelpBox()

	fmt.Println("\n    ✅ Instalasi Selesai! Kamu bisa lanjut mengetik perintah.")
}

func Pause() {
	fmt.Println("    Tekan Enter untuk keluar...")
	fmt.Scanln()
}

func equalsIgnoreCase(a, b string) bool {
	return len(a) == len(b) && filepath.Clean(a) == filepath.Clean(b)
}

func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	return err
}
