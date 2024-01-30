package parser

import "fmt"

type Dba struct {
}

func NewDba() Dba {
	return Dba{}
}

func (d *Dba) IsActive(url string) {
	fmt.Println("isActive called from parser.dba")
}
