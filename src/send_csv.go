// // package main

// // import (
// //     "encoding/csv"
// //     "bytes"
// //     "fmt"
// //     "os"
// //     "net/http"
// //     "encoding/json"
// // )

// // type Message struct {
// //     Phone   string `json:"phone"`
// //     Message string `json:"message"`
// // }

// // func main() {
// //     file, err := os.Open("messages.csv")
// //     if err != nil {
// //         fmt.Println("Error membuka file:", err)
// //         return
// //     }
// //     defer file.Close()

// //     reader := csv.NewReader(file)
// //     records, err := reader.ReadAll()
// //     if err != nil {
// //         fmt.Println("Error membaca CSV:", err)
// //         return
// //     }

// //     for _, record := range records {
// //         if len(record) < 2 {
// //             continue
// //         }

// //         phone := record[0]
// //         message := record[1]

// //         msg := Message{Phone: phone, Message: message}
// //         jsonData, _ := json.Marshal(msg)

// //         resp, err := http.Post("http://localhost:3000/send/message",
// //             "application/json", bytes.NewBuffer(jsonData))
// //         if err != nil {
// //             fmt.Printf("Gagal mengirim ke %s: %v\n", phone, err)
// //             continue
// //         }
// //         defer resp.Body.Close()

// //         fmt.Printf("Pesan ke %s berhasil dikirim!\n", phone)
// //     }
// // }
// package main

// import (
//     "encoding/csv"
//     "bytes"
//     "fmt"
//     "os"
//     "net/http"
//     "encoding/json"
//     "io"
// )

// type Message struct {
//     Phone   string `json:"phone"`
//     Message string `json:"message"`
// }

// func main() {
//     file, err := os.Open("messages.csv")
//     if err != nil {
//         fmt.Println("Error membuka file:", err)
//         return
//     }
//     defer file.Close()

//     reader := csv.NewReader(file)
//     records, err := reader.ReadAll()
//     if err != nil {
//         fmt.Println("Error membaca CSV:", err)
//         return
//     }

//     // Pesan tetap yang akan dikirim ke semua nomor
//     fixedMessage := "Halo ini tes broadcast saja"

//     for _, record := range records {
//         if len(record) < 1 { // Cek apakah ada nomor HP di CSV
//             continue
//         }

//         phone := record[0] // Ambil nomor HP dari CSV

//         msg := Message{Phone: phone, Message: fixedMessage}
//         jsonData, _ := json.Marshal(msg)

//         resp, err := http.Post("http://localhost:3000/send/message",
//             "application/json", bytes.NewBuffer(jsonData))
//         if err != nil {
//             fmt.Printf("Gagal mengirim ke %s: %v\n", phone, err)
//             continue
//         }
//         defer resp.Body.Close()

//         // Cek respons dari server
//         respBody, _ := io.ReadAll(resp.Body)
//         fmt.Printf("Pesan ke %s berhasil dikirim! Server response: %s\n", phone, string(respBody))
//     }
// }


// package main

// import (
//     "encoding/csv"
//     "bytes"
//     "fmt"
//     "os"
//     "net/http"
//     "encoding/json"
//     "io"
//     "time" // Tambahkan ini
// )

// type Message struct {
//     Phone   string `json:"phone"`
//     Message string `json:"message"`
// }

// func main() {
//     file, err := os.Open("messages.csv")
//     if err != nil {
//         fmt.Println("Error membuka file:", err)
//         return
//     }
//     defer file.Close()

//     reader := csv.NewReader(file)
//     records, err := reader.ReadAll()
//     if err != nil {
//         fmt.Println("Error membaca CSV:", err)
//         return
//     }

//     // Pesan tetap yang akan dikirim ke semua nomor
//     fixedMessage := "Halo ini tes broadcast saja"

//     for _, record := range records {
//         if len(record) < 1 { // Cek apakah ada nomor HP di CSV
//             continue
//         }

//         phone := record[0] // Ambil nomor HP dari CSV

//         msg := Message{Phone: phone, Message: fixedMessage}
//         jsonData, _ := json.Marshal(msg)

//         resp, err := http.Post("http://localhost:3000/send/message",
//             "application/json", bytes.NewBuffer(jsonData))
//         if err != nil {
//             fmt.Printf("Gagal mengirim ke %s: %v\n", phone, err)
//             continue
//         }
//         defer resp.Body.Close()

//         // Cek respons dari server
//         respBody, _ := io.ReadAll(resp.Body)
//         fmt.Printf("Pesan ke %s berhasil dikirim! Server response: %s\n", phone, string(respBody))

//         // Tambahkan delay sebelum mengirim pesan berikutnya
//         time.Sleep(3 * time.Second) // Delay 3 detik
//     }
// }


// package main

// import (
//     "encoding/csv"
//     "bytes"
//     "fmt"
//     "os"
//     "net/http"
//     "encoding/json"
//     "io"
//     "time"
//     "math/rand"
// )

// type Message struct {
//     Phone   string `json:"phone"`
//     Message string `json:"message"`
// }

// func main() {
//     // Seed untuk random agar tidak selalu sama
//     rand.Seed(time.Now().UnixNano())

//     file, err := os.Open("messages.csv")
//     if err != nil {
//         fmt.Println("Error membuka file:", err)
//         return
//     }
//     defer file.Close()

//     reader := csv.NewReader(file)
//     records, err := reader.ReadAll()
//     if err != nil {
//         fmt.Println("Error membaca CSV:", err)
//         return
//     }

//     // Pesan tetap yang akan dikirim ke semua nomor
//    fixedMessage := `BADAN KEPEGAWAIAN DAERAH DIKLAT KOTA BANJARMASIN
// https://bkd.banjarmasinkota.go.id
// Ada permohonan dari :
// DINAS KEBUDAYAAN, KEPEMUDAAN, OLAHRAGA DAN PARIWISATA
// untuk Permohonan Kartu Istri ( Karis ) dan Kartu Suami
// No & Tgl Permohonan : 
// 0313125828220331/ASN-02/2025
// 2025-03-13 12:58:28
// Pegawai : 
// 197001051995032001
// Hj. LILY ROSYADAH, SE

// Mohon untuk tidak menghubungi/membalas WA no.wa ini`

//     for _, record := range records {
//         if len(record) < 1 { // Cek apakah ada nomor HP di CSV
//             continue
//         }

//         phone := record[0] // Ambil nomor HP dari CSV

//         msg := Message{Phone: phone, Message: fixedMessage}
//         jsonData, _ := json.Marshal(msg)

//         resp, err := http.Post("http://localhost:3000/send/message",
//             "application/json", bytes.NewBuffer(jsonData))
//         if err != nil {
//             fmt.Printf("Gagal mengirim ke %s: %v\n", phone, err)
//             continue
//         }
//         defer resp.Body.Close()

//         // Cek respons dari server
//         respBody, _ := io.ReadAll(resp.Body)
//         fmt.Printf("Pesan ke %s berhasil dikirim! Server response: %s\n", phone, string(respBody))

//         // Delay acak antara 3-10 detik sebelum mengirim pesan berikutnya
//         delay := time.Duration(rand.Intn(8)+3) * time.Second
//         fmt.Printf("Menunggu selama %v sebelum mengirim pesan berikutnya...\n", delay)
//         time.Sleep(delay)
//     }
// }


// package main

// import (
//     "encoding/csv"
//     "bytes"
//     "fmt"
//     "os"
//     "net/http"
//     "encoding/json"
//     "io"
//     "time"
// )

// type Message struct {
//     Phone   string `json:"phone"`
//     Message string `json:"message"`
// }

// func main() {
//     file, err := os.Open("messages.csv")
//     if err != nil {
//         fmt.Println("Error membuka file:", err)
//         return
//     }
//     defer file.Close()

//     reader := csv.NewReader(file)
//     records, err := reader.ReadAll()
//     if err != nil {
//         fmt.Println("Error membaca CSV:", err)
//         return
//     }

//     // Buka file log untuk pesan berhasil dan gagal
//     successLog, _ := os.OpenFile("success.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
//     defer successLog.Close()
//     failedLog, _ := os.OpenFile("failed.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
//     defer failedLog.Close()

//     // Pesan tetap yang akan dikirim ke semua nomor
//     fixedMessage := `BADAN KEPEGAWAIAN DAERAH DIKLAT KOTA BANJARMASIN
// https://bkd.banjarmasinkota.go.id
// Ada permohonan dari :
// DINAS KEBUDAYAAN, KEPEMUDAAN, OLAHRAGA DAN PARIWISATA
// untuk Permohonan Kartu Istri ( Karis ) dan Kartu Suami
// No & Tgl Permohonan :
// 0313125828220331/ASN-02/2025
// 2025-03-13 12:58:28
// Pegawai :
// 197001051995032001
// Hj. LILY ROSYADAH, SE

// Mohon untuk tidak menghubungi/membalas WA no.wa ini`

//     for _, record := range records {
//         if len(record) < 1 { // Cek apakah ada nomor HP di CSV
//             continue
//         }

//         phone := record[0] // Ambil nomor HP dari CSV

//         msg := Message{Phone: phone, Message: fixedMessage}
//         jsonData, _ := json.Marshal(msg)

//         resp, err := http.Post("http://localhost:3000/send/message",
//             "application/json", bytes.NewBuffer(jsonData))
//         if err != nil {
//             fmt.Printf("Gagal mengirim ke %s: %v\n", phone, err)
//             failedLog.WriteString(fmt.Sprintf("Gagal mengirim ke %s: %v\n", phone, err))
//             continue
//         }
//         defer resp.Body.Close()

//         // Cek respons dari server
//         respBody, _ := io.ReadAll(resp.Body)
//         if resp.StatusCode == http.StatusOK {
//             fmt.Printf("Pesan ke %s berhasil dikirim! Server response: %s\n", phone, string(respBody))
//             successLog.WriteString(fmt.Sprintf("Pesan ke %s berhasil dikirim! Server response: %s\n", phone, string(respBody)))
//         } else {
//             fmt.Printf("Gagal mengirim ke %s! Server response: %s\n", phone, string(respBody))
//             failedLog.WriteString(fmt.Sprintf("Gagal mengirim ke %s! Server response: %s\n", phone, string(respBody)))
//         }

//         // Tambahkan delay sebelum mengirim pesan berikutnya
//         time.Sleep(3 * time.Second) // Delay 3 detik
//     }
// }


package main

import (
    "encoding/csv"
    "bytes"
    "fmt"
    "os"
    "net/http"
    "encoding/json"
    "io"
    "time"
)

type Message struct {
    Phone   string `json:"phone"`
    Message string `json:"message"`
}

func main() {
    // Buka file CSV input
    file, err := os.Open("messages.csv")
    if err != nil {
        fmt.Println("Error membuka file:", err)
        return
    }
    defer file.Close()

    reader := csv.NewReader(file)
    records, err := reader.ReadAll()
    if err != nil {
        fmt.Println("Error membaca CSV:", err)
        return
    }

    // Buat file laporan pengiriman
    reportFile, err := os.Create("report.csv")
    if err != nil {
        fmt.Println("Error membuat file laporan:", err)
        return
    }
    defer reportFile.Close()

    writer := csv.NewWriter(reportFile)
    defer writer.Flush()

    // Tulis header laporan
    writer.Write([]string{"Nomor", "Status", "Waktu", "Response"})

    fixedMessage := `BADAN KEPEGAWAIAN DAERAH DIKLAT KOTA BANJARMASIN
https://bkd.banjarmasinkota.go.id
Ada permohonan dari :
DINAS KEBUDAYAAN, KEPEMUDAAN, OLAHRAGA DAN PARIWISATA
untuk Permohonan Kartu Istri ( Karis ) dan Kartu Suami
No & Tgl Permohonan :
0313125828220331/ASN-02/2025
2025-03-13 12:58:28
Pegawai :
197001051995032001
Hj. LILY ROSYADAH, SE

Mohon untuk tidak menghubungi/membalas WA no.wa ini`

    for _, record := range records {
        if len(record) < 1 { // Pastikan ada nomor HP
            continue
        }

        phone := record[0] // Ambil nomor HP dari CSV
        msg := Message{Phone: phone, Message: fixedMessage}
        jsonData, _ := json.Marshal(msg)

        resp, err := http.Post("http://localhost:3000/send/message",
            "application/json", bytes.NewBuffer(jsonData))

        status := "Berhasil"
        responseText := "OK"

        if err != nil {
            status = "Gagal"
            responseText = err.Error()
            fmt.Printf("Gagal mengirim ke %s: %v\n", phone, err)
        } else {
            defer resp.Body.Close()
            respBody, _ := io.ReadAll(resp.Body)
            responseText = string(respBody)
            fmt.Printf("Pesan ke %s berhasil dikirim! Server response: %s\n", phone, responseText)
        }

        // Simpan log ke file laporan
        writer.Write([]string{phone, status, time.Now().Format("2006-01-02 15:04:05"), responseText})

        // Delay 3-10 detik untuk menghindari deteksi spam
        time.Sleep(time.Duration(3+rand.Intn(8)) * time.Second)
    }
}
