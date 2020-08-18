package main

import "fmt"

func main() {
	// for clasico
	println("for clasico")
	count := 10
	for i := 1; i <= count; i++ {
		fmt.Println(i)
	}

	fmt.Println("===============================================")

	// for continuo, similar a while en otros lenguajes
	println("for continuo")
	j := 1
	for j <= 10 {
		fmt.Println(j)
		j++
	}

	fmt.Println("===============================================")
	// for forever, se usa en procesos que necesitan ser escuchados permanentemente
	println("for forever")
	// k := 1
	// for {
	// 	fmt.Println(k)
	// 	k++
	// }

	fmt.Println("===============================================")

	// for range, permite iterar sobre slices, maps y strings
	// slices
	nums := []uint8{2, 4, 6, 8}

	// i -> indice ej 0
	// v -> valor ej 2
	for i, v := range nums {
		fmt.Println("indice:", i, "valor:", v)
	}

	fmt.Println("===============================================")
	nums1 := []uint8{2, 4, 6, 8}

	// Para modificar el slice original debemos usar el indice,
	// podemos omitir recibir el valor v
	// for _, v := range nums1 {
	// 	v *= 2
	// }

	// asi multiplicamos el slice nums1(Original)
	for i := range nums1 {
		nums1[i] *= 2
	}
	fmt.Println(nums1)

	// for en maps
	sports := map[string]string{"basketball": "🏀", "soccer": "⚽️", "tennis": "🎾"}

	for key, v := range sports {
		fmt.Printf("Key: %s\t\t Valor: %s\n", key, v)
	}

	// for sobre strings
	hello := "hello"

	for _, v := range hello {
		// Se debe usar casting para convertir el numero binario a string
		fmt.Println(string(v))
	}
}
