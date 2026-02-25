package laravel

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
)

func SetupDev() {
	fmt.Println("🛠️  Memulai setup lingkungan pengembangan (Windows)...")

	// 1. Install XAMPP via winget
	fmt.Println("\n📦 1. Mencoba install XAMPP via winget...")
	cmdXampp := exec.Command("winget", "install", "ApacheFriends.XAMPP.8.2") // Versi 8.2 biasanya stabil
	cmdXampp.Stdout = os.Stdout
	cmdXampp.Stderr = os.Stderr
	err := cmdXampp.Run()
	if err != nil {
		fmt.Println("❌ Gagal install XAMPP via winget. Mungkin winget belum ada atau butuh permission.")
		fmt.Println("   Silakan install XAMPP manual dari: https://www.apachefriends.org/index.html")
	} else {
		fmt.Println("✅ XAMPP berhasil di-install/sudah ada!")
	}

	// 2. Install Composer
	composerPath, err := exec.LookPath("composer")
	if err != nil {
		fmt.Println("\n📦 2. Composer tidak ditemukan, mendownload installer...")
		err = downloadComposer()
		if err != nil {
			fmt.Printf("❌ Gagal download Composer: %v\n", err)
		} else {
			fmt.Println("✅ Composer setup downloaded. Silakan jalankan 'Composer-Setup.exe' yang baru saja didownload.")
		}
	} else {
		fmt.Printf("\n✅ Composer sudah ada di: %s\n", composerPath)
	}

	fmt.Println("\n✨ Setup selesai! Jangan lupa tambahkan path PHP (C:\\xampp\\php) ke Environment Variables kamu ya.")
}

func downloadComposer() error {
	url := "https://getcomposer.org/Composer-Setup.exe"
	out, err := os.Create("Composer-Setup.exe")
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}
