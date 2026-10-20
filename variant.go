package mycn

type Option interface {
	mycnOption()
}

type Variant int

const (
	Default Variant = iota
	Secondary
	Outline
	Ghost
	Destructive
	Link
)

func (Variant) mycnOption() {}

type Size int

const (
	SizeDefault Size = iota
	SizeSm
	SizeLg
	SizeIcon
)

func (Size) mycnOption() {}

func variantOf(opts []Option, allowed ...Variant) Variant {
	for _, o := range opts {
		v, ok := o.(Variant)
		if !ok {
			continue
		}
		for _, a := range allowed {
			if v == a {
				return v
			}
		}
	}
	return Default
}

func sizeOf(opts []Option) Size {
	for _, o := range opts {
		if s, ok := o.(Size); ok && s >= SizeDefault && s <= SizeIcon {
			return s
		}
	}
	return SizeDefault
}
