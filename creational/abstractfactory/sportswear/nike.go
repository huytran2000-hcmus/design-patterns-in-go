package sportswear

type NikeShirt struct {
	ProductImp
}

type NikeShoe struct {
	ProductImp
}

type Nike struct{}

func (n *Nike) MakeShirt() Shirt {
	return &NikeShirt{}
}

func (n *Nike) MakeShoe() Shoe {
	return &NikeShoe{}
}
