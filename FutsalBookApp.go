package main

import "fmt"

const NMAX = 1002

type lapangan struct {
	id                    int
	nama                  string
	lokasi                string
	kota                  string
	hargaPerJam           float64
	status                string // "KOSONG", "DIPAKAI","FULL"
	jamOperasionalMulai   int
	jamOperasionalSelesai int
}

type penyewa struct {
	id     int
	nama   string
	noTelp string
	alamat string
	email  string
}

type jadwalsewa struct {
	idSewa     int
	idLapangan int
	idPenyewa  int
	tanggal    string
	jamMulai   int
	jamSelesai int
	totalBiaya float64
}

type tabLapangan [NMAX]lapangan
type tabPenyewa [NMAX]penyewa
type tabJadwalSewa [NMAX]jadwalsewa

var dataPenyewa tabPenyewa
var dataLapangan tabLapangan
var dataJadwalSewa tabJadwalSewa
var nPenyewa int
var nLapangan int
var nJadwalSewa int

func main() {
	tampilkanMenuAwal()
}

func tampilkanMenuAwal() {
	var input int

	fmt.Println("====================================================")
	fmt.Println("|                FUTSAL-BOOK SYSTEM                |")
	fmt.Println("====================================================")
	fmt.Printf("| %-45s %2d |\n", "CRUD Data Lapangan", 1)
	fmt.Printf("| %-45s %2d |\n", "CRUD Data Penyewa", 2)
	fmt.Printf("| %-45s %2d |\n", "CRUD Data Sewa Lapangan", 3)
	fmt.Printf("| %-45s %2d |\n", "Cari Data Jadwal Kosong", 4)
	fmt.Printf("| %-45s %2d |\n", "Cari Data Penyewa", 5)
	fmt.Printf("| %-45s %2d |\n", "Cek Statistik Usaha", 6)
	fmt.Printf("| %-45s %2d |\n", "Keluar", 0)
	fmt.Println("====================================================")
	fmt.Print("Pilih menu: ")
	fmt.Scan(&input)

	if input == 1 {
		ubahDataLapangan()
	} else if input == 2 {
		ubahDataPenyewa()
	} else if input == 3 {
		ubahDataSewaLapangan()
	} else if input == 4 && nLapangan > 0 {
		CariDataJadwalSewa()
	} else if input == 5 && nPenyewa > 0 {
		cariDataPenyewa()
	} else if input == 6 {
		cekStatistikUsaha()
	} else if input == 0 {
		fmt.Println("Terima kasih telah menggunakan Aplikasi Futsal Book !")
	} else {
		fmt.Println("Input tidak valid / ada data yang belum terisi.")
		tampilkanMenuAwal()
	}
}

func ubahDataLapangan() {
	var input int

	fmt.Println("====================================================")
	fmt.Println("|              UBAH DATA LAPANGAN                  |")
	fmt.Println("====================================================")
	fmt.Printf("| %-45s %2d |\n", "Lihat Data Lapangan", 1)
	fmt.Printf("| %-45s %2d |\n", "Tambah Data Lapangan", 2)
	fmt.Printf("| %-45s %2d |\n", "Edit Data Lapangan", 3)
	fmt.Printf("| %-45s %2d |\n", "Hapus Data Lapangan", 4)
	fmt.Printf("| %-45s %2d |\n", "Kembali ke Menu Utama", 0)
	fmt.Println("====================================================")
	fmt.Print("Pilih menu: ")
	fmt.Scan(&input)

	if input == 1 && nLapangan > 0 {
		tampilkanDataLapangan()
	} else if input == 2 {
		tambahDataLapangan()
	} else if input == 3 && nLapangan > 0 {
		editDataLapangan()
	} else if input == 4 && nLapangan > 0 {
		hapusDataLapangan()
	} else if input == 0 {
		tampilkanMenuAwal()
	} else {
		fmt.Println("Input tidak valid / ada data yang belum terisi.")
		ubahDataLapangan()
	}
}

func tampilkanDataLapangan() {
	fmt.Println("======================================================================================================================================================")
	fmt.Printf("| %-3s | %-30s | %-25s | %-20s | %-15s | %-10s | %-10s | %-12s |\n",
		"ID", "Nama", "Alamat", "Kota", "Harga/Jam", "Jam Mulai", "Jam Akhir", "Status")
	fmt.Println("======================================================================================================================================================")

	for i := 1; i <= nLapangan; i++ {
		fmt.Printf("| %-3d | %-30s | %-25s | %-20s | Rp.%-13.2f| %-10d | %-10d | %-12s |\n",
			dataLapangan[i].id,
			dataLapangan[i].nama,
			dataLapangan[i].lokasi,
			dataLapangan[i].kota,
			dataLapangan[i].hargaPerJam,
			dataLapangan[i].jamOperasionalMulai,
			dataLapangan[i].jamOperasionalSelesai,
			dataLapangan[i].status)
	}

	fmt.Println("======================================================================================================================================================")
	ubahDataLapangan()
}

func tambahDataLapangan() {
	var lanjut string
	var ulang, gagal bool

	ulang = true

	fmt.Println("====================================================")
	fmt.Println("|              TAMBAH DATA LAPANGAN                |")
	fmt.Println("====================================================")
	fmt.Println("| Tambah data lapangan yang ingin ditambah         |")
	fmt.Println("====================================================")

	for ulang {
		gagal = false
		if nLapangan >= NMAX {
			fmt.Println("Data penuh.")
			gagal = true
		}
		if !gagal {
			nLapangan++
			n := nLapangan

			dataLapangan[n].id = n

			fmt.Print("Masukkan Nama Lapangan : ")
			fmt.Scan(&dataLapangan[n].nama)

			fmt.Print("Masukkan Alamat        : ")
			fmt.Scan(&dataLapangan[n].lokasi)

			fmt.Print("Masukkan Kota          : ")
			fmt.Scan(&dataLapangan[n].kota)

			fmt.Print("Masukkan Harga per Jam : ")
			fmt.Scan(&dataLapangan[n].hargaPerJam)

			dataLapangan[n].status = "KOSONG"

			fmt.Print("Masukkan Jam Mulai     : ")
			fmt.Scan(&dataLapangan[n].jamOperasionalMulai)
			for dataLapangan[n].jamOperasionalMulai < 0 || dataLapangan[n].jamOperasionalMulai >= 24 {
				fmt.Print("Jam Mulai Tidak Valid, Masukan Jam Mulai Baru: ")
				fmt.Scan(&dataLapangan[n].jamOperasionalMulai)
			}

			fmt.Print("Masukkan Jam Selesai   : ")
			fmt.Scan(&dataLapangan[n].jamOperasionalSelesai)
			for dataLapangan[n].jamOperasionalSelesai <= dataLapangan[n].jamOperasionalMulai ||
				dataLapangan[n].jamOperasionalSelesai > 24 {
				fmt.Print("Jam Selesai Tidak Valid, Masukan Jam Selesai Baru: ")
				fmt.Scan(&dataLapangan[n].jamOperasionalSelesai)
			}

			fmt.Println("----------------------------------------------------")
			fmt.Printf("Data lapangan berhasil diinput						 \n")
			fmt.Println("----------------------------------------------------")
		}

		fmt.Print("Tambah data lagi? (lanjut/stop): ")
		fmt.Scan(&lanjut)
		if lanjut != "lanjut" {
			ulang = false
		}
	}
	ubahDataLapangan()
}

func editDataLapangan() {
	var lanjut string
	var pilih, cari int
	var ulang, gagal bool

	ulang = true

	for ulang {
		gagal = false
		fmt.Println("====================================================")
		fmt.Println("|               EDIT DATA LAPANGAN                 |")
		fmt.Println("====================================================")
		fmt.Println("| Masukkan ID data yang ingin diubah               |")
		fmt.Println("====================================================")

		fmt.Print("Masukkan ID data Lapangan : ")
		fmt.Scan(&cari)
		for cari > nLapangan || cari < 1 {
			fmt.Print("ID Lapangan Tidak Ditemukan, Masukan ID Baru: ")
			fmt.Scan(&cari)
		}

		fmt.Println("----------------------------------------------------")
		fmt.Printf("| %-47s  |\n", "Pilih data yang ingin diubah:")
		fmt.Println("----------------------------------------------------")
		fmt.Printf("| %-45s %2d |\n", "Nama", 1)
		fmt.Printf("| %-45s %2d |\n", "Lokasi", 2)
		fmt.Printf("| %-45s %2d |\n", "Kota", 3)
		fmt.Printf("| %-45s %2d |\n", "Harga per Jam", 4)
		fmt.Printf("| %-45s %2d |\n", "Status", 5)
		fmt.Printf("| %-45s %2d |\n", "Jam Mulai", 6)
		fmt.Printf("| %-45s %2d |\n", "Jam Selesai", 7)
		fmt.Printf("| %-45s %2d |\n", "Kembali", 0)
		fmt.Println("----------------------------------------------------")
		fmt.Print("Pilihan: ")
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			fmt.Print("Masukkan Nama baru: ")
			fmt.Scan(&dataLapangan[cari].nama)
		case 2:
			fmt.Print("Masukkan Lokasi baru: ")
			fmt.Scan(&dataLapangan[cari].lokasi)
		case 3:
			fmt.Print("Masukkan Kota baru: ")
			fmt.Scan(&dataLapangan[cari].kota)
		case 4:
			fmt.Print("Masukkan Harga per Jam baru: ")
			fmt.Scan(&dataLapangan[cari].hargaPerJam)
		case 5:
			fmt.Print("Masukkan Status baru: ")
			fmt.Scan(&dataLapangan[cari].status)
		case 6:
			fmt.Print("Masukkan Jam Mulai baru: ")
			fmt.Scan(&dataLapangan[cari].jamOperasionalMulai)
			for dataLapangan[cari].jamOperasionalMulai < 0 || dataLapangan[cari].jamOperasionalMulai >= 24 {
				fmt.Print("Jam Mulai Tidak Valid, Masukan Jam Mulai Baru: ")
				fmt.Scan(&dataLapangan[cari].jamOperasionalMulai)
			}

		case 7:
			fmt.Print("Masukkan Jam Selesai baru: ")
			fmt.Scan(&dataLapangan[cari].jamOperasionalSelesai)
			for dataLapangan[cari].jamOperasionalSelesai <= dataLapangan[cari].jamOperasionalMulai ||
				dataLapangan[cari].jamOperasionalSelesai > 24 {
				fmt.Print("Jam Selesai Tidak Valid, Masukan Jam Selesai Baru: ")
				fmt.Scan(&dataLapangan[cari].jamOperasionalSelesai)
			}
		case 0:
			ulang = false
			continue
		default:
			fmt.Println("Pilihan tidak valid.")
			gagal = true
		}

		if !gagal {
			fmt.Println("----------------------------------------------------")
			fmt.Println("Data lapangan berhasil diubah.")
			fmt.Println("----------------------------------------------------")
		}
		fmt.Print("Edit data lagi? (lanjut/stop): ")
		fmt.Scan(&lanjut)
		if lanjut != "lanjut" {
			ulang = false
		}
	}

	ubahDataLapangan()
}

func hapusDataLapangan() {
	var ulang bool
	var id int
	var lanjut string

	fmt.Println("====================================================")
	fmt.Println("|              HAPUS DATA LAPANGAN                 |")
	fmt.Println("====================================================")
	fmt.Println("| Hapus data lapangan dengan id lapangan           |")
	fmt.Println("====================================================")

	ulang = true

	for ulang {
		fmt.Print("Masukkan id lapangan : ")
		fmt.Scan(&id)
		for id > nLapangan || id < 1 {
			fmt.Print("ID Lapangan Tidak Ditemukan, Masukan ID Baru: ")
			fmt.Scan(&id)
		}

		ada := false
		for i := 1; i <= nJadwalSewa && !ada; i++ {
			if dataJadwalSewa[i].idLapangan == id {
				ada = true
			}
		}

		if ada {
			fmt.Println("Hapus data sewa terlebih dahulu sebelum menghapus lapangan ini.")
		} else {
			for i := id; i < nLapangan; i++ {
				dataLapangan[i] = dataLapangan[i+1]
				dataLapangan[i].id = i
			}
			dataLapangan[nLapangan] = lapangan{}
			nLapangan--
			fmt.Println("Data lapangan berhasil dihapus.")
		}

		fmt.Print("Hapus data lagi? (lanjut/stop): ")
		fmt.Scan(&lanjut)
		if lanjut != "lanjut" {
			ulang = false
		}
	}

	ubahDataLapangan()
}

func ubahDataPenyewa() {
	var input int

	fmt.Println("====================================================")
	fmt.Println("|              UBAH DATA PENYEWA                   |")
	fmt.Println("====================================================")
	fmt.Printf("| %-45s %2d |\n", "Lihat Data Penyewa", 1)
	fmt.Printf("| %-45s %2d |\n", "Tambah Data Penyewa", 2)
	fmt.Printf("| %-45s %2d |\n", "Edit Data Penyewa", 3)
	fmt.Printf("| %-45s %2d |\n", "Hapus Data Penyewa", 4)
	fmt.Printf("| %-45s %2d |\n", "Kembali ke Menu Utama", 0)
	fmt.Println("====================================================")
	fmt.Print("Pilih menu: ")
	fmt.Scan(&input)

	if input == 1 && nPenyewa > 0 {
		tampilkanDataPenyewa()
	} else if input == 2 {
		tambahDataPenyewa()
	} else if input == 3 && nPenyewa > 0 {
		editDataPenyewa()
	} else if input == 4 && nPenyewa > 0 {
		hapusDataPenyewa()
	} else if input == 0 {
		tampilkanMenuAwal()
	} else {
		fmt.Println("Input tidak valid / ada data yang belum terisi.")
		ubahDataPenyewa()
	}
}

func tampilkanDataPenyewa() {
	fmt.Println("======================================================================================================================================")
	fmt.Printf("| %-3s | %-25s | %-40s | %-30s | %-15s |\n",
		"ID", "Nama", "Alamat", "Email", "No. Telepon")
	fmt.Println("======================================================================================================================================")

	for i := 1; i <= nPenyewa; i++ {
		fmt.Printf("| %-3d | %-25s | %-40s | %-30s | %-15s |\n",
			dataPenyewa[i].id,
			dataPenyewa[i].nama,
			dataPenyewa[i].alamat,
			dataPenyewa[i].email,
			dataPenyewa[i].noTelp)
	}

	fmt.Println("======================================================================================================================================")
	ubahDataPenyewa()
}

func tambahDataPenyewa() {
	var lanjut string
	var ulang, gagal bool

	ulang = true

	fmt.Println("====================================================")
	fmt.Println("|              TAMBAH DATA PENYEWA                 |")
	fmt.Println("====================================================")
	fmt.Println("| Tambah data penyewa yang ingin ditambah          |")
	fmt.Println("====================================================")

	for ulang {
		gagal = false
		if nPenyewa >= NMAX {
			fmt.Println("Data penuh.")
			gagal = true
		}
		if !gagal {
			nPenyewa++
			n := nPenyewa

			dataPenyewa[n].id = n

			fmt.Print("Masukkan Nama Penyewa : ")
			fmt.Scan(&dataPenyewa[n].nama)

			fmt.Print("Masukkan No Telepon   : ")
			fmt.Scan(&dataPenyewa[n].noTelp)

			fmt.Print("Masukkan Alamat       : ")
			fmt.Scan(&dataPenyewa[n].alamat)

			fmt.Print("Masukkan Email        : ")
			fmt.Scan(&dataPenyewa[n].email)

			fmt.Println("----------------------------------------------------")
			fmt.Println("Data penyewa berhasil diinput 					     ")
			fmt.Println("----------------------------------------------------")
		}

		fmt.Print("Tambah data lagi? (lanjut/stop): ")
		fmt.Scan(&lanjut)
		if lanjut != "lanjut" {
			ulang = false
		}
	}

	ubahDataPenyewa()
}

func editDataPenyewa() {
	var lanjut string
	var pilih, cari int
	var ulang, gagal bool

	ulang = true

	for ulang {
		gagal = false
		fmt.Println("====================================================")
		fmt.Println("|               EDIT DATA PENYEWA                  |")
		fmt.Println("====================================================")
		fmt.Println("| Masukkan ID / nomor urut data yang ingin diubah  |")
		fmt.Println("====================================================")

		fmt.Print("Masukkan nomor urut data Penyewa : ")
		fmt.Scan(&cari)
		for cari > nPenyewa || cari < 1 {
			fmt.Print("ID Penyewa Tidak Ditemukan, Masukan ID Baru: ")
			fmt.Scan(&cari)
		}

		fmt.Println("----------------------------------------------------")
		fmt.Printf("| %-47s  |\n", "Pilih data yang ingin diubah:")
		fmt.Println("----------------------------------------------------")
		fmt.Printf("| %-45s %2d |\n", "Nama", 1)
		fmt.Printf("| %-45s %2d |\n", "No Telepon", 2)
		fmt.Printf("| %-45s %2d |\n", "Alamat", 3)
		fmt.Printf("| %-45s %2d |\n", "Email", 4)
		fmt.Printf("| %-45s %2d |\n", "Kembali", 0)
		fmt.Println("----------------------------------------------------")
		fmt.Print("Pilihan: ")
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			fmt.Print("Masukkan Nama baru: ")
			fmt.Scan(&dataPenyewa[cari].nama)
		case 2:
			fmt.Print("Masukkan No Telepon baru: ")
			fmt.Scan(&dataPenyewa[cari].noTelp)
		case 3:
			fmt.Print("Masukkan Alamat baru: ")
			fmt.Scan(&dataPenyewa[cari].alamat)
		case 4:
			fmt.Print("Masukkan Email baru: ")
			fmt.Scan(&dataPenyewa[cari].email)
		case 0:
			ulang = false
			continue
		default:
			fmt.Println("Pilihan tidak valid.")
			gagal = true
		}

		if !gagal {
			fmt.Println("----------------------------------------------------")
			fmt.Println("Data penyewa berhasil diubah.")
			fmt.Println("----------------------------------------------------")
		}
		fmt.Print("Edit data lagi? (lanjut/stop): ")
		fmt.Scan(&lanjut)
		if lanjut != "lanjut" {
			ulang = false
		}
	}

	ubahDataPenyewa()
}

func hapusDataPenyewa() {
	var ulang bool
	var id int
	var lanjut string

	fmt.Println("====================================================")
	fmt.Println("|               HAPUS DATA PENYEWA                 |")
	fmt.Println("====================================================")
	fmt.Println("| Hapus data penyewa dengan id penyewa             |")
	fmt.Println("====================================================")

	ulang = true

	for ulang {
		fmt.Print("Masukkan id penyewa : ")
		fmt.Scan(&id)
		for id > nPenyewa || id < 1 {
			fmt.Print("ID Penyewa Tidak Ditemukan, Masukan ID Baru: ")
			fmt.Scan(&id)
		}

		ada := false
		for i := 1; i <= nJadwalSewa && !ada; i++ {
			if dataJadwalSewa[i].idPenyewa == id {
				ada = true
			}
		}

		if ada {
			fmt.Println("Hapus data sewa terlebih dahulu sebelum menghapus penyewa ini.")
		} else {
			for i := id; i < nPenyewa; i++ {
				dataPenyewa[i] = dataPenyewa[i+1]
				dataPenyewa[i].id = i
			}
			dataPenyewa[nPenyewa] = penyewa{}
			nPenyewa--
			fmt.Println("Data penyewa berhasil dihapus.")
		}

		fmt.Print("Hapus data lagi? (lanjut/stop): ")
		fmt.Scan(&lanjut)
		if lanjut != "lanjut" {
			ulang = false
		}
	}

	ubahDataPenyewa()
}

func ubahDataSewaLapangan() {
	var input int

	fmt.Println("====================================================")
	fmt.Println("|            UBAH DATA SEWA LAPANGAN               |")
	fmt.Println("====================================================")
	fmt.Printf("| %-45s %2d |\n", "Lihat Data Sewa", 1)
	fmt.Printf("| %-45s %2d |\n", "Tambah Sewa Lapangan", 2)
	fmt.Printf("| %-45s %2d |\n", "Cek Status Lapangan", 3)
	fmt.Printf("| %-45s %2d |\n", "Edit Data Sewa", 4)
	fmt.Printf("| %-45s %2d |\n", "Hapus Data Sewa", 5)
	fmt.Printf("| %-45s %2d |\n", "Kembali ke Menu Utama", 0)
	fmt.Println("====================================================")
	fmt.Print("Pilih menu: ")
	fmt.Scan(&input)

	if input == 1 && nJadwalSewa > 0 {
		lihatSewaLapangan()
	} else if input == 2 && nPenyewa > 0 && nLapangan > 0 {
		tambahSewaLapangan()
	} else if input == 3 && nLapangan > 0 {
		cekStatusLapangan()
	} else if input == 4 && nJadwalSewa > 0 {
		editDataSewa()
	} else if input == 5 && nJadwalSewa > 0 {
		hapusDataSewa()
	} else if input == 0 {
		tampilkanMenuAwal()
	} else {
		fmt.Println("Input tidak valid / ada data yang belum terisi.")
		tampilkanMenuAwal()
	}
}

func lihatSewaLapangan() {
	fmt.Println("==============================================================================================================================")
	fmt.Printf("| %-5s | %-10s | %-11s | %-15s | %-13s | %-18s |\n",
		"ID", "ID Penyewa", "ID Lapangan", "Tanggal", "Jam Sewa", "Total Biaya")
	fmt.Println("==============================================================================================================================")

	for i := 1; i <= nJadwalSewa; i++ {
		fmt.Printf("| %-5d | %-10d | %-11d | %-15s | %02d - %02d      | Rp.%-13.2f |\n",
			dataJadwalSewa[i].idSewa,
			dataJadwalSewa[i].idPenyewa,
			dataJadwalSewa[i].idLapangan,
			dataJadwalSewa[i].tanggal,
			dataJadwalSewa[i].jamMulai,
			dataJadwalSewa[i].jamSelesai,
			dataJadwalSewa[i].totalBiaya)
	}

	fmt.Println("==============================================================================================================================")

	ubahDataSewaLapangan()
}

func tambahSewaLapangan() {
	var n int
	var lanjut string
	var ulang, gagal bool

	ulang = true

	fmt.Println("====================================================")
	fmt.Println("|              TAMBAH SEWA LAPANGAN                |")
	fmt.Println("====================================================")
	fmt.Println("| Menambahkan penyewaan lapangan yang tersedia     |")
	fmt.Println("====================================================")

	for ulang {
		gagal = false
		n = nJadwalSewa + 1

		if n >= NMAX {
			fmt.Println("Data penuh.")
			gagal = true
		}
		if !gagal {
			dataJadwalSewa[n].idSewa = n

			fmt.Print("Masukkan ID Penyewa   : ")
			fmt.Scan(&dataJadwalSewa[n].idPenyewa)
			for dataJadwalSewa[n].idPenyewa > nPenyewa || dataJadwalSewa[n].idPenyewa < 1 {
				fmt.Print("ID Penyewa Tidak Valid, Masukan ID Baru   : ")
				fmt.Scan(&dataJadwalSewa[n].idPenyewa)
			}

			fmt.Print("Masukkan ID Lapangan  : ")
			fmt.Scan(&dataJadwalSewa[n].idLapangan)
			for dataJadwalSewa[n].idLapangan > nLapangan || dataJadwalSewa[n].idLapangan < 1 ||
				dataLapangan[dataJadwalSewa[n].idLapangan].status == "FULL" {
				fmt.Print("ID Lapangan Tidak Valid / FULL, Masukan ID Baru   : ")
				fmt.Scan(&dataJadwalSewa[n].idLapangan)
			}

			tempIdLapangan := dataLapangan[dataJadwalSewa[n].idLapangan]

			fmt.Print("Masukkan Tanggal (YYYY-MM-DD) : ")
			fmt.Scan(&dataJadwalSewa[n].tanggal)

			fmt.Print("Masukkan Jam Mulai (0-23)    : ")
			fmt.Scan(&dataJadwalSewa[n].jamMulai)
			for !(dataJadwalSewa[n].jamMulai >= tempIdLapangan.jamOperasionalMulai &&
				dataJadwalSewa[n].jamMulai < tempIdLapangan.jamOperasionalSelesai) {
				fmt.Print("Jam Tidak Valid, Masukan Ulang  : ")
				fmt.Scan(&dataJadwalSewa[n].jamMulai)
			}

			fmt.Print("Masukkan Jam Selesai (0-23)  : ")
			fmt.Scan(&dataJadwalSewa[n].jamSelesai)
			for !(dataJadwalSewa[n].jamSelesai > dataJadwalSewa[n].jamMulai &&
				dataJadwalSewa[n].jamSelesai <= tempIdLapangan.jamOperasionalSelesai) {
				fmt.Print("Jam Tidak Valid, Masukan Ulang  : ")
				fmt.Scan(&dataJadwalSewa[n].jamSelesai)
			}

			mulai := dataJadwalSewa[n].jamMulai
			selesai := dataJadwalSewa[n].jamSelesai
			overlap := false
			for k := 1; k < n && !overlap; k++ {
				if dataJadwalSewa[k].idLapangan == dataJadwalSewa[n].idLapangan &&
					dataJadwalSewa[k].tanggal == dataJadwalSewa[n].tanggal &&
					mulai < dataJadwalSewa[k].jamSelesai &&
					selesai > dataJadwalSewa[k].jamMulai {
					overlap = true
				}
			}

			if overlap {
				fmt.Println("Jadwal bertabrakan dengan sewa yang sudah ada. Sewa dibatalkan.")
			} else {
				dataJadwalSewa[n].totalBiaya = tempIdLapangan.hargaPerJam * float64(selesai-mulai)
				nJadwalSewa++

				jamTerpakai := 0
				for k := 1; k <= nJadwalSewa; k++ {
					if dataJadwalSewa[k].idLapangan == dataJadwalSewa[n].idLapangan &&
						dataJadwalSewa[k].tanggal == dataJadwalSewa[n].tanggal {
						jamTerpakai += dataJadwalSewa[k].jamSelesai - dataJadwalSewa[k].jamMulai
					}
				}
				totalJam := tempIdLapangan.jamOperasionalSelesai - tempIdLapangan.jamOperasionalMulai
				if jamTerpakai >= totalJam {
					dataLapangan[dataJadwalSewa[n].idLapangan].status = "FULL"
				} else {
					dataLapangan[dataJadwalSewa[n].idLapangan].status = "DIPAKAI"
				}

				fmt.Println("----------------------------------------------------")
				fmt.Println("Sewa lapangan berhasil dibuat.")
				fmt.Printf("Total Biaya : RP.%.2f\n", dataJadwalSewa[n].totalBiaya)
				fmt.Printf("Status Lapangan: %s\n", dataLapangan[dataJadwalSewa[n].idLapangan].status)
				fmt.Println("----------------------------------------------------")

				n++
			}
		}

		fmt.Print("Sewa lagi? (lanjut/stop): ")
		fmt.Scan(&lanjut)
		if lanjut != "lanjut" {
			ulang = false
		}
	}

	ubahDataSewaLapangan()
}

func cekStatusLapangan() {
	var pilih, cari int

	fmt.Println("====================================================")
	fmt.Println("|                 STATUS LAPANGAN                  |")
	fmt.Println("====================================================")
	fmt.Printf("| %-45s %2d |\n", "Cari berdasarkan ID", 1)
	fmt.Printf("| %-45s %2d |\n", "Tampilkan semua lapangan", 2)
	fmt.Println("====================================================")
	fmt.Print("Pilih menu: ")
	fmt.Scan(&pilih)

	if pilih == 1 {
		fmt.Print("Masukkan ID Lapangan: ")
		fmt.Scan(&cari)

		if cari >= 1 && cari <= nLapangan {
			fmt.Println("----------------------------------------------------")
			fmt.Printf("STATUS LAPANGAN DENGAN ID %d\n", cari)
			fmt.Println("----------------------------------------------------")
			fmt.Println("Status :", dataLapangan[cari].status)
			fmt.Println("----------------------------------------------------")
		} else {
			fmt.Println("ID Lapangan tidak ditemukan.")
		}

	} else if pilih == 2 {
		fmt.Println("----------------------------------------------------")
		fmt.Println("DATA STATUS LAPANGAN")
		fmt.Println("----------------------------------------------------")

		for i := 1; i <= nLapangan; i++ {
			fmt.Println("ID    :", dataLapangan[i].id)
			fmt.Println("Nama  :", dataLapangan[i].nama)
			fmt.Println("Status:", dataLapangan[i].status)
			fmt.Println("----------------------------------------------------")
		}

	} else {
		fmt.Println("Pilihan tidak valid.")
	}

	ubahDataSewaLapangan()
}

func editDataSewa() {
	var lanjut string
	var pilih, cari int
	var ulang bool = true

	for ulang {
		fmt.Println("====================================================")
		fmt.Println("|                  EDIT DATA SEWA                  |")
		fmt.Println("====================================================")
		fmt.Println("| Masukkan ID Jadwal Sewa yang ingin diubah        |")
		fmt.Println("====================================================")

		fmt.Print("Masukkan ID Jadwal Sewa: ")
		fmt.Scan(&cari)
		for cari < 1 || cari > nJadwalSewa {
			fmt.Print("ID tidak ditemukan, masukkan ID lagi: ")
			fmt.Scan(&cari)
		}

		fmt.Println("----------------------------------------------------")
		fmt.Printf("| %-47s  |\n", "Pilih data yang ingin diubah:")
		fmt.Println("----------------------------------------------------")
		fmt.Printf("| %-45s %2d |\n", "ID Penyewa", 1)
		fmt.Printf("| %-45s %2d |\n", "ID Lapangan", 2)
		fmt.Printf("| %-45s %2d |\n", "Tanggal", 3)
		fmt.Printf("| %-45s %2d |\n", "Jam Mulai & Selesai", 4)
		fmt.Printf("| %-45s %2d |\n", "Kembali", 0)
		fmt.Println("----------------------------------------------------")
		fmt.Print("Pilihan: ")
		fmt.Scan(&pilih)

		valid := true

		switch pilih {
		case 1:
			fmt.Print("Masukkan ID Penyewa baru: ")
			fmt.Scan(&dataJadwalSewa[cari].idPenyewa)
			for dataJadwalSewa[cari].idPenyewa > nPenyewa || dataJadwalSewa[cari].idPenyewa < 1 {
				fmt.Print("ID Penyewa Tidak Valid, Masukan ID Baru: ")
				fmt.Scan(&dataJadwalSewa[cari].idPenyewa)
			}

		case 2:
			idLama := dataJadwalSewa[cari].idLapangan
			fmt.Print("Masukkan ID Lapangan baru: ")
			fmt.Scan(&dataJadwalSewa[cari].idLapangan)
			for dataJadwalSewa[cari].idLapangan > nLapangan || dataJadwalSewa[cari].idLapangan < 1 ||
				(dataLapangan[dataJadwalSewa[cari].idLapangan].status == "FULL" && dataJadwalSewa[cari].idLapangan != idLama) {
				fmt.Print("ID Lapangan Tidak Valid / FULL, Masukan ID Baru: ")
				fmt.Scan(&dataJadwalSewa[cari].idLapangan)
			}

		case 3:
			fmt.Print("Masukkan Tanggal baru (YYYY-MM-DD): ")
			fmt.Scan(&dataJadwalSewa[cari].tanggal)

		case 4:
			tempLap := dataLapangan[dataJadwalSewa[cari].idLapangan]
			fmt.Print("Masukkan Jam Mulai baru (0-23): ")
			fmt.Scan(&dataJadwalSewa[cari].jamMulai)
			for !(dataJadwalSewa[cari].jamMulai >= tempLap.jamOperasionalMulai &&
				dataJadwalSewa[cari].jamMulai < tempLap.jamOperasionalSelesai) {
				fmt.Print("Jam Tidak Valid, Masukan Ulang: ")
				fmt.Scan(&dataJadwalSewa[cari].jamMulai)
			}
			fmt.Print("Masukkan Jam Selesai baru (0-23): ")
			fmt.Scan(&dataJadwalSewa[cari].jamSelesai)
			for !(dataJadwalSewa[cari].jamSelesai > dataJadwalSewa[cari].jamMulai &&
				dataJadwalSewa[cari].jamSelesai <= tempLap.jamOperasionalSelesai) {
				fmt.Print("Jam Tidak Valid, Masukan Ulang: ")
				fmt.Scan(&dataJadwalSewa[cari].jamSelesai)
			}

		case 0:
			ulang = false
			valid = false

		default:
			fmt.Println("Pilihan tidak valid.")
			valid = false
		}

		if valid {
			overlap := false
			for k := 1; k <= nJadwalSewa; k++ {
				if k != cari &&
					dataJadwalSewa[k].idLapangan == dataJadwalSewa[cari].idLapangan &&
					dataJadwalSewa[k].tanggal == dataJadwalSewa[cari].tanggal &&
					dataJadwalSewa[cari].jamMulai < dataJadwalSewa[k].jamSelesai &&
					dataJadwalSewa[cari].jamSelesai > dataJadwalSewa[k].jamMulai {
					overlap = true
				}
			}

			if overlap {
				fmt.Println("Jadwal bertabrakan dengan sewa yang sudah ada. Silakan input ulang.")
			} else {
				dataJadwalSewa[cari].totalBiaya = dataLapangan[dataJadwalSewa[cari].idLapangan].hargaPerJam *
					float64(dataJadwalSewa[cari].jamSelesai-dataJadwalSewa[cari].jamMulai)

				jamTerpakai := 0
				for k := 1; k <= nJadwalSewa; k++ {
					if dataJadwalSewa[k].idLapangan == dataJadwalSewa[cari].idLapangan &&
						dataJadwalSewa[k].tanggal == dataJadwalSewa[cari].tanggal {
						jamTerpakai += dataJadwalSewa[k].jamSelesai - dataJadwalSewa[k].jamMulai
					}
				}
				totalJam := dataLapangan[dataJadwalSewa[cari].idLapangan].jamOperasionalSelesai -
					dataLapangan[dataJadwalSewa[cari].idLapangan].jamOperasionalMulai
				if jamTerpakai >= totalJam {
					dataLapangan[dataJadwalSewa[cari].idLapangan].status = "FULL"
				} else {
					dataLapangan[dataJadwalSewa[cari].idLapangan].status = "DIPAKAI"
				}

				fmt.Println("----------------------------------------------------")
				fmt.Println("Data sewa berhasil diubah.")
				fmt.Printf("Total Biaya baru: RP.%.2f,00\n", dataJadwalSewa[cari].totalBiaya)
				fmt.Printf("Status Lapangan: %s\n", dataLapangan[dataJadwalSewa[cari].idLapangan].status)
				fmt.Println("----------------------------------------------------")
			}

			fmt.Print("Edit data lagi? (lanjut/stop): ")
			fmt.Scan(&lanjut)
			if lanjut != "lanjut" {
				ulang = false
			}
		}
	}

	ubahDataSewaLapangan()
}

func hapusDataSewa() {
	var ulang bool
	var id int
	var lanjut string

	fmt.Println("====================================================")
	fmt.Println("|              HAPUS DATA JADWAL SEWA              |")
	fmt.Println("====================================================")
	fmt.Println("| Hapus data jadwal sewa dengan id sewa            |")
	fmt.Println("====================================================")

	ulang = true

	for ulang {
		fmt.Print("Masukkan id sewa : ")
		fmt.Scan(&id)
		for id > nJadwalSewa || id < 1 {
			fmt.Print("ID Jadwal Sewa Tidak Ditemukan, Masukan ID Baru: ")
			fmt.Scan(&id)
		}

		for i := id; i <= nJadwalSewa-1; i++ {
			dataJadwalSewa[i] = dataJadwalSewa[i+1]
			dataJadwalSewa[i].idSewa = i
		}
		dataJadwalSewa[nJadwalSewa] = jadwalsewa{}
		nJadwalSewa--

		fmt.Print("Hapus data lagi? (lanjut/stop): ")
		fmt.Scan(&lanjut)
		if lanjut != "lanjut" {
			ulang = false
		}
	}

	ubahDataSewaLapangan()
}

func CariDataJadwalSewa() {
	var cek1, cek2, cek4 int
	var temp lapangan
	var pass, i, j, minIdx int
	var lanjut, cek3 string
	var ulang bool
	var gagal bool
	var dataUrut tabLapangan

	ulang = true

	for ulang {
		gagal = false

		fmt.Println("====================================================")
		fmt.Println("|                CARI DATA JADWAL                  |")
		fmt.Println("====================================================")
		fmt.Println("|Cari Jadwal kosong dari Jam Mulai atau Harga Sewa |")
		fmt.Println("====================================================")

		fmt.Print("Urutkan Jam Mulai atau Harga Sewa? (1/2) : ")
		fmt.Scan(&cek1)

		fmt.Print("Pakai Selection atau Insertion Sort? (1/2) : ")
		fmt.Scan(&cek2)

		fmt.Print("Ascending atau Descending? (1/2) : ")
		fmt.Scan(&cek4)

		for i = 1; i <= nLapangan; i++ {
			dataUrut[i] = dataLapangan[i]
		}

		if cek2 == 1 {
			if cek1 == 1 {
				for i = 1; i <= nLapangan; i++ {
					minIdx = i
					for j = i + 1; j <= nLapangan; j++ {
						if dataUrut[j].jamOperasionalMulai < dataUrut[minIdx].jamOperasionalMulai {
							minIdx = j
						}
					}
					temp = dataUrut[i]
					dataUrut[i] = dataUrut[minIdx]
					dataUrut[minIdx] = temp
				}
			} else if cek1 == 2 {
				for i = 1; i <= nLapangan; i++ {
					minIdx = i
					for j = i + 1; j <= nLapangan; j++ {
						if dataUrut[j].hargaPerJam < dataUrut[minIdx].hargaPerJam {
							minIdx = j
						}
					}
					temp = dataUrut[i]
					dataUrut[i] = dataUrut[minIdx]
					dataUrut[minIdx] = temp
				}
			} else {
				fmt.Println("Pilihan tidak valid.")
				gagal = true
			}

		} else if cek2 == 2 {
			if cek1 == 1 {
				for pass = 2; pass <= nLapangan; pass++ {
					temp = dataUrut[pass]
					i = pass - 1

					for i >= 1 && temp.jamOperasionalMulai < dataUrut[i].jamOperasionalMulai {
						dataUrut[i+1] = dataUrut[i]
						i--
					}
					dataUrut[i+1] = temp
				}
			} else if cek1 == 2 {
				for pass = 2; pass <= nLapangan; pass++ {
					temp = dataUrut[pass]
					i = pass - 1

					for i >= 1 && temp.hargaPerJam < dataUrut[i].hargaPerJam {
						dataUrut[i+1] = dataUrut[i]
						i--
					}
					dataUrut[i+1] = temp
				}
			} else {
				fmt.Println("Pilihan tidak valid.")
				gagal = true
			}

		} else {
			fmt.Println("Pilihan tidak valid.")
			gagal = true
		}

		if cek4 == 2 {
			for a, b := 1, nLapangan; a < b; a, b = a+1, b-1 {
				dataUrut[a], dataUrut[b] = dataUrut[b], dataUrut[a]
			}
		}

		if !gagal {
			fmt.Println("----------------------------------------------------")
			fmt.Println("Data lapangan berhasil diurutkan.")
			fmt.Println("----------------------------------------------------")

			fmt.Print("Mau tampilkan data setelah diurutkan? (iya/tidak): ")
			fmt.Scan(&cek3)

			if cek3 == "iya" {
				ditemukan := false

				for i = 1; i <= nLapangan; i++ {
					if dataUrut[i].status == "KOSONG" {
						ditemukan = true
						fmt.Println("ID     :", dataUrut[i].id)
						fmt.Println("Nama   :", dataUrut[i].nama)
						fmt.Println("Lokasi :", dataUrut[i].lokasi)
						fmt.Println("Kota   :", dataUrut[i].kota)
						fmt.Println("Harga per Jam :", dataUrut[i].hargaPerJam)
						fmt.Println("Status :", dataUrut[i].status)
						fmt.Println("Jam Mulai :", dataUrut[i].jamOperasionalMulai)
						fmt.Println("Jam Selesai :", dataUrut[i].jamOperasionalSelesai)
						fmt.Println("----------------------------------------------------")
					}
				}

				if !ditemukan {
					fmt.Println("Tidak ada data lapangan dengan status KOSONG.")
				}
			}

			fmt.Print("Mau masukan ke data utama? (iya/tidak): ")
			fmt.Scan(&cek3)

			if cek3 == "iya" {
				for i = 1; i <= nLapangan; i++ {
					dataLapangan[i] = dataUrut[i]
				}
				fmt.Println("Data hasil sorting sudah dimasukkan ke data utama.")
			} else {
				fmt.Println("Data utama tidak diubah.")
			}
		}

		fmt.Print("Cari lagi? (lanjut/stop): ")
		fmt.Scan(&lanjut)
		if lanjut != "lanjut" {
			ulang = false
		}
	}

	tampilkanMenuAwal()
}

func cariDataPenyewa() {
	var cek1, cek2 int
	var cariNama, cariTLP string
	var idx int
	var kiri, kanan, tengah int
	var lanjut string
	var ulang bool
	var i, j int
	var temp penyewa
	var dataUrut tabPenyewa

	ulang = true

	for ulang {
		fmt.Println("====================================================")
		fmt.Println("|                CARI DATA PENYEWA                 |")
		fmt.Println("====================================================")
		fmt.Println("|   Cari Data Penyewa dari Nama atau No Telepon    |")
		fmt.Println("====================================================")

		fmt.Print("Cari berdasarkan Nama atau No Telepon? (1/2) : ")
		fmt.Scan(&cek1)

		fmt.Print("Pakai Sequential atau Binary Search? (1/2) : ")
		fmt.Scan(&cek2)

		for i = 1; i <= nPenyewa; i++ {
			dataUrut[i] = dataPenyewa[i]
		}

		if cek2 == 1 {
			if cek1 == 1 {
				fmt.Print("Masukkan Nama Penyewa : ")
				fmt.Scan(&cariNama)

				idx = -1
				for i = 1; i <= nPenyewa && idx == -1; i++ {
					if dataUrut[i].nama == cariNama {
						idx = i
					}
				}

				if idx != -1 {
					fmt.Println("DATA PENYEWA DITEMUKAN")
					fmt.Println("ID     :", dataUrut[idx].id)
					fmt.Println("Nama   :", dataUrut[idx].nama)
					fmt.Println("NoTelp :", dataUrut[idx].noTelp)
					fmt.Println("Alamat :", dataUrut[idx].alamat)
					fmt.Println("Email  :", dataUrut[idx].email)
				} else {
					fmt.Println("Data penyewa dengan nama tersebut tidak ditemukan.")
				}

			} else if cek1 == 2 {
				fmt.Print("Masukkan Nomor Telepon Penyewa : ")
				fmt.Scan(&cariTLP)

				idx = -1
				for i = 1; i <= nPenyewa && idx == -1; i++ {
					if dataUrut[i].noTelp == cariTLP {
						idx = i
					}
				}

				if idx != -1 {
					fmt.Println("DATA PENYEWA DITEMUKAN")
					fmt.Println("ID     :", dataUrut[idx].id)
					fmt.Println("Nama   :", dataUrut[idx].nama)
					fmt.Println("NoTelp :", dataUrut[idx].noTelp)
					fmt.Println("Alamat :", dataUrut[idx].alamat)
					fmt.Println("Email  :", dataUrut[idx].email)
				} else {
					fmt.Println("Data penyewa dengan nomor telepon tersebut tidak ditemukan.")
				}

			} else {
				fmt.Println("Pilihan tidak valid.")
			}

		} else if cek2 == 2 {
			if cek1 == 1 {
				fmt.Print("Masukkan Nama Penyewa : ")
				fmt.Scan(&cariNama)

				for i = 1; i < nPenyewa; i++ {
					for j = i + 1; j <= nPenyewa; j++ {
						if dataUrut[i].nama > dataUrut[j].nama {
							temp = dataUrut[i]
							dataUrut[i] = dataUrut[j]
							dataUrut[j] = temp
						}
					}
				}

				kiri = 1
				kanan = nPenyewa
				idx = -1

				for kiri <= kanan && idx == -1 {
					tengah = (kiri + kanan) / 2

					if dataUrut[tengah].nama == cariNama {
						idx = tengah
					} else if dataUrut[tengah].nama < cariNama {
						kiri = tengah + 1
					} else {
						kanan = tengah - 1
					}
				}

				if idx != -1 {
					fmt.Println("DATA PENYEWA DITEMUKAN")
					fmt.Println("ID     :", dataUrut[idx].id)
					fmt.Println("Nama   :", dataUrut[idx].nama)
					fmt.Println("NoTelp :", dataUrut[idx].noTelp)
					fmt.Println("Alamat :", dataUrut[idx].alamat)
					fmt.Println("Email  :", dataUrut[idx].email)
				} else {
					fmt.Println("Data penyewa dengan nama tersebut tidak ditemukan.")
				}

			} else if cek1 == 2 {
				fmt.Print("Masukkan Nomor Telepon Penyewa : ")
				fmt.Scan(&cariTLP)

				for i = 1; i < nPenyewa; i++ {
					for j = i + 1; j <= nPenyewa; j++ {
						if dataUrut[i].noTelp > dataUrut[j].noTelp {
							temp = dataUrut[i]
							dataUrut[i] = dataUrut[j]
							dataUrut[j] = temp
						}
					}
				}

				kiri = 1
				kanan = nPenyewa
				idx = -1

				for kiri <= kanan && idx == -1 {
					tengah = (kiri + kanan) / 2

					if dataUrut[tengah].noTelp == cariTLP {
						idx = tengah
					} else if dataUrut[tengah].noTelp < cariTLP {
						kiri = tengah + 1
					} else {
						kanan = tengah - 1
					}
				}

				if idx != -1 {
					fmt.Println("DATA PENYEWA DITEMUKAN")
					fmt.Println("ID     :", dataUrut[idx].id)
					fmt.Println("Nama   :", dataUrut[idx].nama)
					fmt.Println("NoTelp :", dataUrut[idx].noTelp)
					fmt.Println("Alamat :", dataUrut[idx].alamat)
					fmt.Println("Email  :", dataUrut[idx].email)
				} else {
					fmt.Println("Data penyewa dengan nomor telepon tersebut tidak ditemukan.")
				}

			} else {
				fmt.Println("Pilihan tidak valid.")
			}

		} else {
			fmt.Println("Pilihan tidak valid.")
		}

		fmt.Print("Cari lagi? (lanjut/stop): ")
		fmt.Scan(&lanjut)
		if lanjut != "lanjut" {
			ulang = false
		}
	}

	tampilkanMenuAwal()
}

func cekStatistikUsaha() {
	var input int

	fmt.Println("====================================================")
	fmt.Println("|              CEK STATISTIK USAHA                 |")
	fmt.Println("====================================================")
	fmt.Printf("| %-45s %2d |\n", "Cek Pendapatan Bulanan", 1)
	fmt.Printf("| %-45s %2d |\n", "Cari Jam Paling Populer", 2)
	fmt.Printf("| %-45s %2d |\n", "Kembali ke Menu Utama", 0)
	fmt.Println("====================================================")
	fmt.Print("Pilih menu: ")
	fmt.Scan(&input)

	if input == 1 && nJadwalSewa > 0 {
		cekPendapatanBulanan()
	} else if input == 2 && nJadwalSewa > 0 && nLapangan > 0 {
		cariJamPopuler()
	} else if input == 0 {
		tampilkanMenuAwal()
	} else {
		fmt.Println("Input tidak valid / ada data yang belum terisi.")
		tampilkanMenuAwal()
	}
}

func cekPendapatanBulanan() {
	var cek int
	var ulang bool
	var pilih, lanjut string
	var totalPerBulan [13]float64

	ulang = true

	for ulang {
		totalPerBulan = [13]float64{}
		fmt.Println("====================================================")
		fmt.Println("|             CEK PENDAPATAN BULANAN               |")
		fmt.Println("====================================================")
		fmt.Println("| Memberikan statistik pendapatan perbulan         |")
		fmt.Println("====================================================")

		fmt.Print("Pendapatan setiap bulan atau spesifik? (1/2) : ")
		fmt.Scan(&cek)

		if cek == 1 {

			for i := 1; i <= nJadwalSewa; i++ {
				bulan := dataJadwalSewa[i].tanggal[5:7]

				if bulan == "01" {
					totalPerBulan[1] += dataJadwalSewa[i].totalBiaya
				} else if bulan == "02" {
					totalPerBulan[2] += dataJadwalSewa[i].totalBiaya
				} else if bulan == "03" {
					totalPerBulan[3] += dataJadwalSewa[i].totalBiaya
				} else if bulan == "04" {
					totalPerBulan[4] += dataJadwalSewa[i].totalBiaya
				} else if bulan == "05" {
					totalPerBulan[5] += dataJadwalSewa[i].totalBiaya
				} else if bulan == "06" {
					totalPerBulan[6] += dataJadwalSewa[i].totalBiaya
				} else if bulan == "07" {
					totalPerBulan[7] += dataJadwalSewa[i].totalBiaya
				} else if bulan == "08" {
					totalPerBulan[8] += dataJadwalSewa[i].totalBiaya
				} else if bulan == "09" {
					totalPerBulan[9] += dataJadwalSewa[i].totalBiaya
				} else if bulan == "10" {
					totalPerBulan[10] += dataJadwalSewa[i].totalBiaya
				} else if bulan == "11" {
					totalPerBulan[11] += dataJadwalSewa[i].totalBiaya
				} else if bulan == "12" {
					totalPerBulan[12] += dataJadwalSewa[i].totalBiaya
				}
			}

			fmt.Println("----------------------------------------------------")
			fmt.Println("PENDAPATAN SEMUA BULAN")
			fmt.Println("----------------------------------------------------")

			for j := 1; j <= 12; j++ {
				fmt.Printf("Bulan %02d : RP.%.2f\n", j, totalPerBulan[j])
			}
			fmt.Println("----------------------------------------------------")

		} else if cek == 2 {
			fmt.Print("Bulan ke berapa? (MM) : ")
			fmt.Scan(&pilih)

			var total float64
			for i := 1; i <= nJadwalSewa; i++ {
				if dataJadwalSewa[i].tanggal[5:7] == pilih {
					total += dataJadwalSewa[i].totalBiaya
				}
			}

			fmt.Println("----------------------------------------------------")
			fmt.Printf("Pendapatan Pada Bulan %s.\n", pilih)
			fmt.Printf("Total Pendapatan : RP.%.2f\n", total)
			fmt.Println("----------------------------------------------------")

		} else {
			fmt.Println("Pilihan tidak valid.")
		}

		fmt.Print("Mau Lihat Lagi? (lanjut/stop): ")
		fmt.Scan(&lanjut)
		if lanjut != "lanjut" {
			ulang = false
		}
	}

	cekStatistikUsaha()
}

func cariJamPopuler() {
	var ulang bool = true
	var lanjut string
	var hitungJam [24]int

	fmt.Println("====================================================")
	fmt.Println("|                CARI JAM POPULER                  |")
	fmt.Println("====================================================")
	fmt.Println("| Memberikan statistik jam paling populer          |")
	fmt.Println("====================================================")

	for ulang {
		for i := 0; i < 24; i++ {
			hitungJam[i] = 0
		}

		for j := 1; j <= nJadwalSewa; j++ {
			for jam := dataJadwalSewa[j].jamMulai; jam < dataJadwalSewa[j].jamSelesai; jam++ {
				if jam >= 0 && jam < 24 {
					hitungJam[jam]++
				}
			}
		}

		jamPopuler := 0
		maxCount := hitungJam[0]

		for jam := 1; jam < 24; jam++ {
			if hitungJam[jam] > maxCount {
				maxCount = hitungJam[jam]
				jamPopuler = jam
			}
		}

		fmt.Println("----------------------------------------------------")
		fmt.Printf("Jam paling populer: %02d:00 dengan %d pemakaian\n", jamPopuler, maxCount)
		fmt.Println("----------------------------------------------------")

		fmt.Print("Mau Lihat Lagi? (lanjut/stop): ")
		fmt.Scan(&lanjut)
		if lanjut != "lanjut" {
			ulang = false
		}
	}

	cekStatistikUsaha()
}
