package system

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

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

	// 3. Jika sudah berjalan dari folder resmi, STOP.
	if strings.EqualFold(exeDir, targetDir) {
		return
	}

	// 4. Cek apakah file sudah terpasang di folder target (Paling akurat)
	if _, err := os.Stat(targetPath); err == nil {
		return
	}

	// 5. Cek apakah alias 'bill' (atau nama exe) sudah ada di PATH
	cmdName := strings.TrimSuffix(exeName, filepath.Ext(exeName))
	if _, err := exec.LookPath(cmdName); err == nil {
		return
	}

	// 6. Belum terinstall, tawarkan setup wizard
	fmt.Printf("\n👋 Halo! Sepertinya Billy Asisten belum terpasang di sistem kamu.\n")
	fmt.Println("Mau saya pasang otomatis supaya bisa dipanggil dari mana saja?")
	fmt.Print("Pasang sekarang? (y/n, default: y): ")

	var response string
	fmt.Scanln(&response)
	response = strings.ToLower(strings.TrimSpace(response))

	if response == "" || response == "y" || response == "yes" {
		InstallToSystem(exePath, targetDir, targetPath)
	} else {
		fmt.Println("\n👍 Oke, aplikasi akan tetap berjalan dalam mode portable.")
	}
}

func InstallToSystem(src, targetDir, targetPath string) {
	fmt.Println("\n🛠️  Sedang memasang ke sistem...")

	// Buat folder target jika belum ada
	err := os.MkdirAll(targetDir, 0755)
	if err != nil {
		fmt.Printf("❌ Gagal membuat folder: %v\n", err)
		pause()
		return
	}

	// Copy file .exe ke folder target
	err = copyFile(src, targetPath)
	if err != nil {
		fmt.Printf("❌ Gagal menyalin file: %v\n", err)
		pause()
		return
	}

	// Daftarkan ke PATH Windows menggunakan PowerShell
	// Kita cek dulu biar nggak double-double di PATH
	script := fmt.Sprintf(`
		$oldPath = [Environment]::GetEnvironmentVariable("Path", "User");
		if ($oldPath -notlike "*%s*") {
			[Environment]::SetEnvironmentVariable("Path", $oldPath + ";%s", "User")
		}
	`, targetDir, targetDir)
	cmd := exec.Command("powershell", "-Command", script)
	err = cmd.Run()

	if err != nil {
		fmt.Printf("❌ Gagal mendaftarkan PATH: %v\n", err)
		pause()
		return
	}

	fmt.Println("\n✅ Instalasi Berhasil!")
	fmt.Println("--------------------------------------------------")
	fmt.Println("Silakan TUTUP terminal ini dan BUKA kembali.")
	fmt.Printf("Setelah itu, kamu bisa panggil lewat terminal dari mana saja!\n")
	fmt.Println("--------------------------------------------------")
	pause()
	os.Exit(0)
}

func pause() {
	fmt.Println("\nTekan Enter untuk keluar...")
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
