package github

import (
	"fmt"
	"os"
	"os/exec"
)

func PushToGithub(link string, message string) {
	fmt.Printf("🚀 Memulai proses push ke: %s\n", link)

	// 1. Cek apakah folder .git sudah ada
	if _, err := os.Stat(".git"); os.IsNotExist(err) {
		fmt.Println("📂 Inisialisasi Git repository baru...")
		runGitCommand("git", "init")
		runGitCommand("git", "remote", "add", "origin", link)
		runGitCommand("git", "branch", "-M", "main")
	} else {
		fmt.Println("✅ Git repository sudah terdeteksi.")
	}

	// 2. Proses add, commit, dan push
	fmt.Println("📝 Menambahkan file dan membuat commit...")
	runGitCommand("git", "add", ".")
	runGitCommand("git", "commit", "-m", message)

	fmt.Println("⬆️  Sedang push ke GitHub...")
	err := runGitCommand("git", "push", "-u", "origin", "main")

	if err != nil {
		fmt.Printf("\n❌ Gagal push ke GitHub: %v\n", err)
		fmt.Println("Pastikan link repository benar dan kamu punya akses.")
	} else {
		fmt.Println("\n✅ Berhasil! Kode kamu sudah aman di GitHub.")
	}
}

func runGitCommand(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
