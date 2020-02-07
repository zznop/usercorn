package sms

import (
    "encoding/binary"
    "github.com/lunixbochs/usercorn/go/models/cpu"
    "fmt"
)

const(
    A = iota
    B
    C
    D
    E
    H
    L
    F
    AS
    BS
    CS
    DS
    ES
    HS
    LS
    FS
    I
    R
    SP
    IX
    IY
    PC
)

type Builder struct{}

type SMSCpu struct {
    *cpu.Hooks
    *cpu.Regs
    *cpu.Mem
    exitRequest bool
    err error
}

func (b *Builder) New() (cpu.Cpu, error) {
    c := &SMSCpu{
        Regs: cpu.NewRegs(16, []int{
            A, B, C, D, E, H, L, F,
            AS, BS, CS, DS, ES, HS, LS, FS,
            I, R,
            SP, IX, IY,
            PC,
        }),
        Mem: cpu.NewMem(16, binary.LittleEndian),
    }

    c.Hooks = cpu.NewHooks(c, c.Mem)
    return c, nil
}

func (s *SMSCpu) Start(begin, until uint64) error {
    fmt.Println("In start")
    return s.err
}

func (s *SMSCpu) Stop() error {
    s.exitRequest = true
    return nil
}

func (s *SMSCpu) Close() error {
    return nil
}

func (s *SMSCpu) Backend() interface{} {
    return s
}
