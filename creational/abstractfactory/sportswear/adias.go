package sportswear

type AdidasShirt struct {
	ProductImp
}

type AdidasShoe struct {
	ProductImp
}

type Adidas struct{}

func (a *Adidas) MakeShirt() Shirt {
	return &AdidasShirt{}
}

func (a *Adidas) MakeShoe() Shoe {
	return &AdidasShoe{}
}
