package sms

import (
    "fmt"
    "github.com/lunixbochs/usercorn/go/models"
)

type Dis struct {}

func (d *Dis) Dis(mem []byte, addr uint64) ([]models.Ins, error) {
    var ret []models.Ins
    fmt.Println("In Dis (sms)")
    return ret, nil
}
