package main

import "fmt"

func main() {
	a := "HA"
	var (
		b      string
		name   string
		color  string
		number string
		i      int
	)

	var parkings = []parking{
		{
			capasity: 10,
			name:     "Blok_A",
			price:    5000,
			cars: []car{
				{
					name:   name,
					color:  color,
					number: number,
				},
			},
		},
		{
			capasity: 15,
			name:     "Blok_B",
			price:    4000,
			cars: []car{
				{
					name:   name,
					color:  color,
					number: number,
				},
			},
		},
	}
	for a != "YO'Q" {
		fmt.Printf("Mashina keldimi? [HA] [YO'Q]: ")
		fmt.Scan(&a)
		fmt.Printf("Mashina ma'lumotlarini kiriting!: ")
		fmt.Printf("Nomi: ")
		fmt.Scan(&name)
		fmt.Printf("Rangi: ")
		fmt.Scan(&color)
		fmt.Printf("No'meri: ")
		fmt.Scan(&number)
		fmt.Printf("Qaysi Blokga kiritasiz?:  ")
		fmt.Scan(&b)
		if b == "Blok_A" {
			parkings[0].cars[i].name = name
			parkings[0].cars[i].color = color
			parkings[0].cars[i].number = number
		}
		if b == "Blok_B" {
			parkings[1].cars[i].name = name
			parkings[1].cars[i].color = color
			parkings[1].cars[i].number = number
		}
		i = i + 1
	}

	fmt.Printf("NAME => %s\nCAPASITY => %d\nPRICE => %d\n", parkings[0].name, parkings[0].capasity, parkings[0].price)
	fmt.Printf("CAR NAME => %s  CAR COLOR => %s  CAR NUMBER => %s\n", parkings[0].cars[0].name, parkings[0].cars[0].color, parkings[0].cars[0].number)
	fmt.Println("//////////////////////////////////////////////////////////////////")

	fmt.Printf("NAME => %s\nCAPASITY => %d\nPRICE => %d\n", parkings[1].name, parkings[1].capasity, parkings[1].price)
	fmt.Printf("CAR NAME => %s  CAR COLOR => %s  CAR NUMBER => %s\n", parkings[1].cars[0].name, parkings[1].cars[0].color, parkings[1].cars[0].number)
}

type parking struct {
	capasity int
	name     string
	price    int
	cars     []car
}

type car struct {
	name   string
	color  string
	number string
}
