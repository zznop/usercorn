package sms

import (
    "github.com/lunixbochs/usercorn/go/models"
    "github.com/lunixbochs/usercorn/go/cpu/sms"
)

var Arch = &models.Arch{
    Name: "sms",
    Bits: 16,
    Cpu: &sms.Builder{},
    Dis: &sms.Dis{},
    Asm: nil,

    PC: sms.PC,
    SP: sms.SP,
    Regs: map[string]int{
        // Main registers
        "a": sms.A,
        "b": sms.B,
        "c": sms.C,
        "d": sms.D,
        "e": sms.E,
        "h": sms.H,
        "l": sms.L,
        "f": sms.F,

        // Shadow registers
        "a'": sms.AS,
        "b'": sms.BS,
        "c'": sms.CS,
        "d'": sms.DS,
        "e'": sms.ES,
        "h'": sms.HS,
        "l'": sms.LS,
        "f'": sms.FS,

        // Other registers
        "i": sms.I,
        "r": sms.R,

        // Index registers
        "sp": sms.SP,
        "ix": sms.IX,
        "iy": sms.IY,

        // Program counter
        "pc": sms.PC,
    },

    DefaultRegs: []string {
        "a", "b", "c", "d", "e", "h", "l", "f", "a'", "b'", "c'", "d'", "e'", "h'", "l'", "f'",
    },
}
