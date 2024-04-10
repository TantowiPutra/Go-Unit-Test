# Go-Unit-Test
Repository To Learn Go Unit Test

Struktur Project Go: Controller, Service, Repository


Software Testing is Done to ensure application quality
Cost goes up as the pyramid goes up
<!-- Test Pyramid -->
.UI                 -> Buka website melalui browser, input form, dll, termasuk sistem backend (service & unit) 
...Service          -> Fokus ke 1 Aplikasi (Front-end + Back-end), bisa juga test langsung backend melalui API, keuntungannya tidak perlu running keseluruhan seperti test UI
                       cons: harus running database
......Unit          -> Check satu unit dari banyak unit pada 1 Aplikasi, misal Unit Test Controller, Unit Test Service, dan Repository


Unit is the smallest component

<!-- testing.T -->
Sebuah struct yang digunakan untuk unit test di Go-Lang

<!-- testing.M -->
Sebuah struct yang digunakan untuk mengatur life cycle testing

<!-- testing.B -->
Sebuah struct yang digunakan untuk benchmarking, mengetahui seberapa cepat dan efisien kode program dapat berjalan

<!-- Command To Run Unit Test -->
- go test    -> Run unit test without showing corresponding functions
- go test -v -> Run unit test and show corresponding functions (v : verbose)
