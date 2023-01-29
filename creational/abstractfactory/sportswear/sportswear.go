package sportswear

import "errors"

type SportswearFactory interface {
	MakeShirt() Shirt
	MakeShoe() Shoe
}

func GetSportswearFactory(brand BrandType) (SportswearFactory, error) {
	switch brand {
	case AdidasType:
		return &Adidas{}, nil
	case NikeType:
		return &Nike{}, nil
	default:
		return nil, errors.New("brand type is not existed")
	}
}

type Product interface {
	GetColor() ColorType
	SetColor(color ColorType)
	GetSize() int
	SetSize(size int)
}

type Shirt interface {
	Product
}

type Shoe interface {
	Product
}

type (
	BrandType int
	ColorType string
)

const (
	AdidasType BrandType = iota
	NikeType
)

const (
	Red  ColorType = "red"
	Blue ColorType = "blue"
)

type ProductImp struct {
	size  int
	color ColorType
}

func (p *ProductImp) GetColor() ColorType {
	return p.color
}

func (p *ProductImp) SetColor(color ColorType) {
	p.color = color
}

func (p *ProductImp) GetSize() int {
	return p.size
}

func (p *ProductImp) SetSize(size int) {
	if size > 100 {
		return
	}

	p.size = size
}
