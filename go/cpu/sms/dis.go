package sms

import (
    "github.com/lunixbochs/usercorn/go/models"
)

// Dis is the Z80 disassembler structure
type Dis struct {}

// Dis disassembles Z80 data (Not implemented, yet)
func (d *Dis) Dis(data []byte, addr uint64) ([]models.Ins, error) {
    var ret []models.Ins
    return ret, nil
}
