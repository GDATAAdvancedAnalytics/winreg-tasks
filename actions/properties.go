package actions

type PropertiesMagic uint16

type Properties interface {
	Magic() PropertiesMagic
	Name() string
	String() string
}
