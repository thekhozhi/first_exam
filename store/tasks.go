package store

import (
	"fmt"
	"sort"
)

func Task_1() {
	 cashSum := 0
	 totalSum := 0

	 for _, vcus := range Customers{
		cashSum = cashSum + vcus.Cash
		totalSum = totalSum + vcus.Basket.Total
	 }
	 fmt.Println("General Cash:",cashSum)
	 fmt.Println("Total:",totalSum)
}


func Task_2() {
	sort.Slice(Customers,func(i, j int) bool {
		return Customers[i].Basket.Total > Customers[j].Basket.Total
	})
	for _, vcus := range Customers{
		fmt.Println("Customer with the most Total:",vcus.FirstName,vcus.LastName,vcus.Basket.Total)
		break
	}
}


func Task_3() {
	tempMap := make(map[string]int)

	sort.Slice(Customers,func(i, j int) bool {
		return Customers[i].Basket.Total > Customers[j].Basket.Total
	})
	
	for i, vcus := range Customers{
		if i == 1{
			break
		}
		for _, vprod := range vcus.Basket.Products{
			tempMap[vprod.Name] = vprod.Price
		}
	}
	
	type Count struct{
		Key string
		Value int
	}
	counts := []Count{}

	for i, v := range tempMap{
		counts = append(counts, Count{i,v})
	}

	sort.Slice(counts,func(i, j int) bool {
		return counts[i].Value > counts[j].Value
	})
	
	for _, v := range counts{
		fmt.Printf("%s (%d sum)\n", v.Key, v.Value)
		break
	}
}


func Task_4(){
	 prodCount := 0
	 prodSum := 0
	 average := 0
	for _, vcus := range Customers{
		for _, vprod := range vcus.Basket.Products{
			prodSum = prodSum + vprod.Price
			 prodCount++
		}
	}
	average = prodSum / prodCount
	fmt.Printf("Average is: %d sum\n",average)
}


func Task_5(){
	sort.Slice(Customers,func(i, j int) bool {
		return Customers[i].Basket.Total < Customers[j].Basket.Total
	})
	for _, vcus := range Customers{
		fmt.Println("Customer with the least Total:",vcus.FirstName,vcus.LastName,vcus.Basket.Total)
		break
	}
}


func Task_6(){
	type Count struct{
		Name string
		Quantity int
	}
	counts := []Count{}

	for _, vcus := range Customers{
		for _, vprod := range vcus.Basket.Products{
			 counts = append(counts, Count{vprod.Category,vprod.Quantity})
		}
	}

	for i := 0; i < len(counts); i++{
		for j := i + 1; j < len(counts); j++{
			if counts[i].Name == counts[j].Name{
				counts[i].Quantity = counts[i].Quantity + counts[j].Quantity
				counts = append(counts[:j], counts[j+1:]...)
			}
		}
	}
	
	sort.Slice(counts,func(i, j int) bool {
		return counts[i].Quantity > counts[j].Quantity
	})

	for _, v := range counts{
		fmt.Println("Most sold category is:",v.Name)
		break
	}
}


func Task_7(){
	tempMap := make(map[string]int)

	for _, vcus := range Customers{
		for _, vprod := range vcus.Basket.Products{
			tempMap[vprod.Name] = vprod.Quantity
		}
	}
	
	type Count struct{
		Key string
		Value int
	}
	counts := []Count{}

	for i, v := range tempMap{
		counts = append(counts, Count{i,v})
	}
	
	sort.Slice(counts, func(i, j int) bool {
		return counts[i].Value > counts[j].Value
	})
	
	for _, v := range counts{
		fmt.Printf("The Most sold product: %s (%d ta)\n\n",v.Key,v.Value)
		break
	}

	for _, v := range counts{
		if v.Value == 1{
			fmt.Printf("The least sold products: %s (%d ta)\n",v.Key, v.Value)
		}
	}
}



func Task_8(){
	quanSum := 0
	count := 0
	quanAverage := 0.0
	for _, vcus := range Customers{
		for _,vprod := range vcus.Basket.Products{
			quanSum = quanSum + vprod.Quantity
			count++
		}
	}
	quanAverage = float64(quanSum) / float64(count)
	fmt.Println("Average quantity of all sold products:",quanAverage)
}


func Task_9(customer []Customer) {
	var maxSpender Customer

	sort.Slice(customer, func(i, j int) bool {
		return len(customer[i].Basket.Products) > len(customer[j].Basket.Products)
	})

	if len(customer) > 0 {
		maxSpender = customer[0]

		fmt.Printf("Eng ko'p mahsulot sotib olgan mijoz: %s %s (sotib olgan mahsulotlar soni: %d)\n", maxSpender.FirstName, maxSpender.LastName, len(maxSpender.Basket.Products))
	}
}



func Task_10(){
	topMap := make(map[string]int)
	for _, vcus := range Customers{
		for _, vprod := range vcus.Basket.Products{
			topMap[vprod.Name] = vprod.Quantity
		}
	}
    
	type Count struct{
		Key string
		Value int
	}

	counts := []Count{}

	for i, v := range topMap{
		counts = append(counts, Count{i,v})
	}
    
	sort.Slice(counts,func(i, j int) bool {
		return counts[i].Value > counts[j].Value
	})
	
	for _, v := range counts{
		fmt.Println("The product with the most quantity is:",v.Key,v.Value)
		break
	}
}


func Task_11(customers []Customer) {
	maxAverage := 0
    var cusFirstName, cusLastName string
	
    for _, vcus := range customers {
		AllSum := 0
        for _, mahsulot := range vcus.Basket.Products {
			AllSum += mahsulot.Price * mahsulot.Quantity
        }
		
        averageSum := AllSum / len(vcus.Basket.Products)
        if averageSum > maxAverage {
			maxAverage = averageSum
            cusFirstName = vcus.FirstName
            cusLastName = vcus.LastName
        }
    }
	
    fmt.Printf("Average total: %d sum\n", maxAverage)
    fmt.Printf("Customer with the most buy: %s %s\n", cusFirstName, cusLastName)
}


func Task_12(customer []Customer) {
	var maxTotal int
	var bestProduct Product

	for _, vcus := range Customers {
		for _, vprod := range vcus.Basket.Products {
			total := vprod.Quantity * vprod.Price
			if total > maxTotal {
				maxTotal = total
				bestProduct = vprod
			}
		}
	}

	if maxTotal > 0 {
		fmt.Printf("Most used product: %s, category: %s, sum: %d\n", bestProduct.Name, bestProduct.Category, maxTotal)
	} else {
		fmt.Println("Couldn't find the product!!!")
	}
}



func Task_13(){
	topMap := make(map[string]int)

	for _, vcus := range Customers{
		for _, vprod := range vcus.Basket.Products{
			topMap[vprod.Name] = vprod.Price
		}
	}
	
	type Count struct{
		Key string
		Value int
	}

	counts := []Count{}

	for i, v := range topMap{
		counts = append(counts, Count{i,v})
	}
	
	sort.Slice(counts,func(i, j int) bool {
		return counts[i].Value > counts[j].Value
	})
	
	for i, v := range counts{
		fmt.Printf("%s (%d sum)\n",v.Key,v.Value)
		if i == 1 {
			break
		}
	}
}


func Task_14(){
	tempMap := make(map[string]int)

	for _, vcus := range Customers{
		for _, vprod := range vcus.Basket.Products{
			 tempMap[vprod.Category] = vprod.Quantity
		}

	 var(
		maxCat = ""
		maxQuan = 0
	 )
	 for cat, quan := range tempMap{
		if quan > maxQuan{
			maxQuan = quan
			maxCat = cat
		}
	 }

	 if maxQuan > 0{
		fmt.Printf("%s %s - most sold category is: %s, sold quantity: %d\n",vcus.FirstName,vcus.LastName,maxCat,maxQuan)
	 }
}
}


func Task_15(){
	for _, vcus := range Customers{
		for _, vprod := range vcus.Basket.Products{
			fmt.Println(vprod.Name,vprod.Quantity,"ta sotilgan")
		}
	}
}










			
			