package kode

import (
	"bill/commands/fitur/github"
	fitur "bill/commands/fitur/help"
	"bill/commands/fitur/laravel"
	"bill/commands/fitur/system"
	"fmt"
	"os"
)

func Jalankan() {
	// 1. Cek instalasi (PATH dsb)
	system.CheckInstallation()

	// 2. Jika dipanggil tanpa argumen (bisa jadi Double-Click)
	if len(os.Args) < 2 {
		fitur.PrintHelp()

		// Penyelamat agar CMD tidak langsung hilang jika di-klik 2x (Double-click)
		fmt.Println("\nTekan Enter untuk keluar...")
		var tmp string
		fmt.Scanln(&tmp)
		return
	}

	subCommand := os.Args[1]

	switch subCommand {
	case "buat":
		if len(os.Args) < 4 || os.Args[2] != "laravel" {
			fmt.Println("❌ Format salah. Gunakan: bill buat laravel [nama]")
			return
		}
		laravel.BuatLaravel(os.Args[3])

	case "cek":
		if len(os.Args) < 3 || os.Args[2] != "dev" {
			fmt.Println("❌ Format salah. Gunakan: bill cek dev")
			return
		}
		laravel.CekDev()

	case "setup":
		if len(os.Args) < 3 || os.Args[2] != "dev" {
			fmt.Println("❌ Format salah. Gunakan: bill setup dev")
			return
		}
		laravel.SetupDev()

	case "push":
		if len(os.Args) < 4 {
			fmt.Println("❌ Format salah. Gunakan:")
			fmt.Println("   - bill push [link_github] [pesan] (Pertama kali)")
			fmt.Println("   - bill push update [pesan]        (Repository lama)")
			return
		}
		github.PushToGithub(os.Args[2], os.Args[3])

	case "help":
		fitur.PrintHelp()

	default:
		fmt.Printf("❌ Perintah '%s' tidak dikenali.\n", subCommand)
		fitur.PrintHelp()
	}
}
