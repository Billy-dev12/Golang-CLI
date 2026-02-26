package system

import (
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

func printWelcomeBanner() {
	banner := `
    ╔═══════════════════════════════════════════════╗
    ║                                               ║
    ║   🚀  BILLY ASISTEN - CLI TOOL  🚀           ║
    ║                                               ║
    ║   Selamat datang!                             ║
    ║   Semua fitur sudah siap digunakan.           ║
    ║                                               ║
    ╚═══════════════════════════════════════════════╝
`
	fmt.Println(banner)
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

func printHelpInBox() {
	fmt.Println("    ┌───────────────────────────────────────────────┐")
	fmt.Println("    │  📖  DAFTAR PERINTAH:                        │")
	fmt.Println("    ├───────────────────────────────────────────────┤")
	fmt.Println("    │  bill buat laravel [nama]  → Buat project    │")
	fmt.Println("    │  bill cek dev              → Cek environment │")
	fmt.Println("    │  bill setup dev            → Pasang tools    │")
	fmt.Println("    │  bill push [link] [pesan]  → Push ke GitHub  │")
	fmt.Println("    │  bill push update [pesan]  → Update GitHub   │")
	fmt.Println("    │  bill help                 → Tampilkan ini   │")
	fmt.Println("    └───────────────────────────────────────────────┘")
	fmt.Println()
	fmt.Println("    🌐 https://github.com/Billy-dev12")
	fmt.Println()
}

// ShowWelcome menampilkan robot + banner + help saat dijalankan tanpa argumen
func ShowWelcome() {
	printRobot()
	printWelcomeBanner()
	printHelpInBox()
	pause()
}

// ======================================
//        🔧 LOGIKA INSTALASI
// ======================================

func CheckInstallation() {
	// 1. Dapatkan lokasi file .exe yang sedang berjalan
	exePath, err := os.Executable()
	if err != nil {
		return
	}
	exeDir := filepath.Dir(exePath)
	exeName := filepath.Base(exePath)

	// 2. Tentukan folder instalasi resmi (AppData/Local/bill-tool)
	localAppData := os.Getenv("LOCALAPPDATA")
	targetDir := filepath.Join(localAppData, "bill-tool")
	targetPath := filepath.Join(targetDir, exeName)

	// 3. Jika sudah berjalan dari folder resmi, STOP. Tampilkan robot selamat datang.
	if strings.EqualFold(exeDir, targetDir) {
		return
	}

	// 4. Cek apakah file sudah terpasang di folder target (Paling akurat)
	if _, err := os.Stat(targetPath); err == nil {
		return
	}

	// 5. Cek apakah alias sudah ada di PATH
	cmdName := strings.TrimSuffix(exeName, filepath.Ext(exeName))
	if _, err := exec.LookPath(cmdName); err == nil {
		return
	}

	// 6. Belum terinstall, tawarkan setup wizard dengan gaya!
	printRobot()
	fmt.Println("    👋 Halo! Sepertinya Billy Asisten belum terpasang")
	fmt.Println("    di sistem kamu.")
	fmt.Println()
	fmt.Println("    Mau saya pasang otomatis supaya bisa dipanggil")
	fmt.Println("    dari mana saja?")
	fmt.Println()
	fmt.Print("    Pasang sekarang? (y/n, default: y): ")

	var response string
	fmt.Scanln(&response)
	response = strings.ToLower(strings.TrimSpace(response))

	if response == "" || response == "y" || response == "yes" {
		InstallToSystem(exePath, targetDir, targetPath)
	} else {
		fmt.Println("\n    👍 Oke, aplikasi berjalan dalam mode portable.")
	}
}

func InstallToSystem(src, targetDir, targetPath string) {
	fmt.Println()

	// Buat folder target jika belum ada
	printLoading("Membuat folder instalasi...")
	err := os.MkdirAll(targetDir, 0755)
	if err != nil {
		fmt.Printf("    ❌ Gagal membuat folder: %v\n", err)
		pause()
		return
	}

	// Copy file .exe ke folder target
	printLoading("Menyalin file ke sistem...")
	err = copyFile(src, targetPath)
	if err != nil {
		fmt.Printf("    ❌ Gagal menyalin file: %v\n", err)
		pause()
		return
	}

	// Daftarkan ke PATH Windows (tanpa duplikat)
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
		pause()
		return
	}

	// 🎉 Tampilkan hasil instalasi yang meriah!
	printRobot()
	printInstalledBanner()
	printHelpInBox()

	pause()
	os.Exit(0)
}

func pause() {
	fmt.Println("    Tekan Enter untuk keluar...")
	fmt.Scanln()
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
