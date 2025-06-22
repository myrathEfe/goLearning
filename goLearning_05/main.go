package main

import (
	"fmt"
	"sync"
	"time"
)

// WaitGroup, tüm goroutine'lerin tamamlanmasını beklemek için kullanılır.
var wg = sync.WaitGroup{}

// dbData, veritabanından geldiği varsayılan örnek veri kümesidir.
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}

// results, her dbCall fonksiyonunun sonucunu toplamak için kullanılan slice.
// Ancak bu yapı goroutine'ler tarafından eşzamanlı yazıldığında veri kaybı olabilir.
var results = []string{}

func main() {
	// Programın başlangıç zamanını alıyoruz.
	t0 := time.Now()

	// Her dbData elemanı için bir goroutine başlatılır.
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)    // Her yeni goroutine için WaitGroup sayacını artır.
		go dbCall(i) // dbCall fonksiyonu goroutine olarak çalıştırılır.
	}

	// Tüm goroutine'lerin bitmesini bekle.
	wg.Wait()

	// Toplam çalışma süresini yazdır.
	fmt.Printf("\nTotal execution time : %v", time.Since(t0))

	// results dizisinin içeriğini yazdır.
	// Bu dizi veri kaybı yaşayabilir çünkü eşzamanlı erişim güvenli değil.
	fmt.Printf("\nThe results are : %v", results)
}

func dbCall(i int) {
	// Bu goroutine tamamlandığında WaitGroup sayacını azalt.
	defer wg.Done()

	// Sabit bir gecikme ile veritabanı sorgusu simüle edilir (2 saniye).
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)

	// Simüle edilen veritabanı sonucu ekrana yazılır.
	fmt.Println("The result from the database is : ", dbData[i])

	// results dizisine veri ekleniyor. Ancak bu işlem thread-safe değil!
	// Bu yüzden bazı veriler kaybolabilir veya çakışabilir.
	// Yani veri kaybı yaşanır.
	results = append(results, dbData[i])
}
