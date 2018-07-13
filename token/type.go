package initoken

type Type interface {
	INIToken() Type
	String() string
}
