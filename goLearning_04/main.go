package main

import "fmt"

func main() {
	var p *int32 = new(int32)
	var i int32

	fmt.Printf("The value p points to is : %v\n", *p)
	fmt.Printf("The adress of p is : %v\n", p)
	fmt.Printf("The value if i is : %v\n", i)
	p = &i
	*p = 1
	fmt.Printf("The value p points to is : %v\n", *p)
	fmt.Printf("The value if i is : %v\n", i)
	var k int32 = 2
	i = k
	fmt.Printf("The value if i is : %v\n", i)
	fmt.Println("************************************")
	var thing1 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("The memory location of the thing1 is : %p\n", &thing1)
	var result = square(&thing1)
	fmt.Printf("The result is : %v\n", result)
	fmt.Printf("The value of thing1 is : %v\n", thing1)
}

func square(thing2 *[5]float64) [5]float64 {
	fmt.Printf("The memory location of the thing2 is : %p\n", thing2)
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	return *thing2
}

// Eğer pointer kullanılmasaydı, square fonksiyonuna thing1 dizisinin bir kopyası gönderilecekti.
// Bu durumda, orijinal thing1 dizisinin içeriği (1, 2, 3, 4, 5) değişmeden kalırdı.
// Ancak pointer kullanılarak, square fonksiyonu doğrudan thing1 dizisinin bellekteki adresi üzerinden işlem yaptı.
// Bu sayede bellekte yeni bir dizi (kopya) oluşturulmadı; ekstra bellek harcaması olmadı.
// result aslında thing1'in değiştirilmiş halinin bir kopyasını tutar, ancak pointer sayesinde esas değişiklik zaten thing1 üzerinde yapıldığından,
// result değişkenine ihtiyaç kalmadan da işlem tamamlanabilirdi.
// Eğer pointer kullanılmasaydı, result ve thing1 farklı bellek adreslerini işaret ederdi.
