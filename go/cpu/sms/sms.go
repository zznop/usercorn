package sms

import (
    "encoding/binary"
    "github.com/lunixbochs/usercorn/go/models/cpu"
    "fmt"
)

const maxInstrWidth = 4

// Builder is a struct that is used by the framework to setup the emulator
type Builder struct{}

// SMSCpu is a struct for maintaining state and emulating the SMS Z80 processor
type SMSCpu struct {
    *cpu.Hooks
    *cpu.Regs
    *cpu.Mem
    exitRequest bool
    err error
}

// New returns a pointer to a initialized SMSCpu struct
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

// Start is a SMSCpu method for emulating the SMS Z80 processor
func (s *SMSCpu) Start(begin, until uint64) error {
    s.exitRequest = false
    pc := begin
    s.RegWrite(PC, pc)
    s.OnBlock(pc, 0)

    var err error
    for pc != until && err == nil && !s.exitRequest{
        _, err := s.ReadProt(pc, maxInstrWidth, cpu.PROT_EXEC)
        if err != nil {
            break
        }

    }
    fmt.Println("In start")
    return s.err
}

// Stop is a SMSCpu method for stopping the emulator
func (s *SMSCpu) Stop() error {
    s.exitRequest = true
    return nil
}

// Close is a SMSCpu method that has no operation
func (s *SMSCpu) Close() error {
    return nil
}

// Backend is a SMSCpu method that returns the SMSCpu struct to the consumer
func (s *SMSCpu) Backend() interface{} {
    return s
}
