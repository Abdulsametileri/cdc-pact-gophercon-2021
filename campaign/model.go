package main

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type Products map[int]Product

func (p Products) InitProducts() {
	p[1] = Product{
		ID:    1,
		Name:  "Product 1",
		Price: 100,
	}
	p[2] = Product{
		ID:    2,
		Name:  "Product 2",
		Price: 100,
	}
	p[3] = Product{
		ID:    3,
		Name:  "Product 3",
		Price: 200,
	}
}
