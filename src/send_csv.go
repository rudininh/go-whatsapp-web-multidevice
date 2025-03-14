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


package main

import (
    "encoding/csv"
    "bytes"
    "fmt"
    "os"
    "net/http"
    "encoding/json"
    "io"
    "time" // Tambahkan ini
)

type Message struct {
    Phone   string `json:"phone"`
    Message string `json:"message"`
}

func main() {
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

    // Pesan tetap yang akan dikirim ke semua nomor
    fixedMessage := "Halo ini tes broadcast saja"

    for _, record := range records {
        if len(record) < 1 { // Cek apakah ada nomor HP di CSV
            continue
        }

        phone := record[0] // Ambil nomor HP dari CSV

        msg := Message{Phone: phone, Message: fixedMessage}
        jsonData, _ := json.Marshal(msg)

        resp, err := http.Post("http://localhost:3000/send/message",
            "application/json", bytes.NewBuffer(jsonData))
        if err != nil {
            fmt.Printf("Gagal mengirim ke %s: %v\n", phone, err)
            continue
        }
        defer resp.Body.Close()

        // Cek respons dari server
        respBody, _ := io.ReadAll(resp.Body)
        fmt.Printf("Pesan ke %s berhasil dikirim! Server response: %s\n", phone, string(respBody))

        // Tambahkan delay sebelum mengirim pesan berikutnya
        time.Sleep(3 * time.Second) // Delay 3 detik
    }
}
