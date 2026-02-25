package fitur

import "fmt"

func PrintHelp() {
	fmt.Println("\nCara penggunaan Billahi Robby Commands:")
	fmt.Println("  bill buat laravel [nama]    - Membuat project Laravel baru")
	fmt.Println("  bill cek dev               - Mengecek lingkungan PHP, Composer, XAMPP")
	fmt.Println("  bill setup dev             - Menginstall alat pendukung (XAMPP, Composer)")
	fmt.Println("  bill push [link] [pesan]   - Push project ke GitHub (Pertama kali)")
	fmt.Println("  bill push update [pesan]   - Push update ke GitHub (Repository lama)")
	fmt.Println("\nInfo lebih lanjut cek: https://github.com/Billy-dev12")
}
