package sms

import (
    "encoding/binary"
    "github.com/lunixbochs/usercorn/go/models/cpu"
    "fmt"
)

// Builder is a struct that is used by the framework to setup the emulator
type Builder struct{}

// IntInfo is a struct that stores info on the interrupt mode
type IntInfo struct {
    enabled bool
    mode int
}

// SMSCpu is a struct for maintaining state and emulating the SMS Z80 processor
type SMSCpu struct {
    *cpu.Hooks
    *cpu.Regs
    *cpu.Mem
    pc uint64
    interrupt IntInfo
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

// instrIM emulates the Z80 "im" instruction
func (s *SMSCpu) instrIM() (int, uint64, error) {
    b, err := s.ReadProt(s.pc+1, 1, cpu.PROT_EXEC)
    if err != nil {
        return 0, 0, err
    }
    switch b[0] {
    case 0x46: // im 0
        s.interrupt.mode = INT_MODE_0
    case 0x56: // im 1
        s.interrupt.mode = INT_MODE_1
    case 0x5e: // im 2
        s.interrupt.mode = INT_MODE_2
    default:
        panic("Unsupported interrupt mode!?")
    }
    return 2, 2, nil
}

// instrJP emulates the Z80 "jp" instruction
func (s *SMSCpu) instrJP(opcode uint8) (int, uint64, error) {
    switch opcode {
    case 0xc3: // jp nn
        b, err := s.ReadProt(s.pc+1, 2, cpu.PROT_EXEC)
        if err != nil {
            return 0, 0, err
        }
        abs := binary.LittleEndian.Uint16(b)
        displ := int(abs)-int(s.pc)
        return displ, 3, nil
    default:
        panic("Unsupported jp variant!?")
    }
}

// Dispatch is a SMSCpu method for dispatching each generic operation handler
func (s *SMSCpu) Dispatch(opcode uint8) (int, uint64, error) {
    operation := OperationMap[opcode]
    switch operation {
    case OP_NOP:
        fmt.Println("nop")
        return 1, 1, nil
    case OP_LD:
        fmt.Println("ld")
    case OP_INC:
        fmt.Println("inc")
    case OP_DEC:
        fmt.Println("dec")
    case OP_RLCA:
        fmt.Println("rlca")
    case OP_EX:
        fmt.Println("ex")
    case OP_ADD:
        fmt.Println("add")
    case OP_RRCA:
        fmt.Println("rrca")
    case OP_DJNZ:
        fmt.Println("djnz")
    case OP_RLA:
        fmt.Println("rla")
    case OP_JR:
        fmt.Println("jr")
    case OP_RRA:
        fmt.Println("rra")
    case OP_DAA:
        fmt.Println("daa")
    case OP_CPL:
        fmt.Println("cpl")
    case OP_SCF:
        fmt.Println("scf")
    case OP_CCF:
        fmt.Println("ccf")
    case OP_HALT:
        fmt.Println("halt")
    case OP_ADC:
        fmt.Println("adc")
    case OP_SUB:
        fmt.Println("sub")
    case OP_SBC:
        fmt.Println("sbc")
    case OP_AND:
        fmt.Println("and")
    case OP_XOR:
        fmt.Println("xor")
    case OP_OR:
        fmt.Println("or")
    case OP_CP:
        fmt.Println("cp")
    case OP_RET:
        fmt.Println("ret")
    case OP_POP:
        fmt.Println("pop")
    case OP_JP:
        fmt.Println("jp")
        return s.instrJP(opcode)
    case OP_CALL:
        fmt.Println("call")
    case OP_PUSH:
        fmt.Println("push")
    case OP_RST:
        fmt.Println("rst")
    case OP_UNDOC:
        fmt.Println("undoc")
    case OP_OUT:
        fmt.Println("out")
    case OP_EXX:
        fmt.Println("exx")
    case OP_IN:
        fmt.Println("in")
    case OP_DI:
        fmt.Println("di")
        s.interrupt.enabled = false
        return 1, 1, nil
    case OP_EI:
        fmt.Println("ei")
        s.interrupt.enabled = true
        return 1, 1, nil
    case OP_IM:
        fmt.Println("im")
        return s.instrIM()
    }
    panic("Unsupported Z80 operation!?")
}

// Start is a SMSCpu method that runs the emulator
func (s *SMSCpu) Start(begin, until uint64) error {
    s.exitRequest = false
    s.pc = begin
    s.RegWrite(PC, s.pc)
    s.OnBlock(s.pc, 0)

    var err error
    var data []byte
    var displ int
    var c uint64
    cycles := uint64(0)
    for s.pc != until && !s.exitRequest{
        if data, err = s.ReadProt(s.pc, 1, cpu.PROT_EXEC); err != nil {
            break
        }

        if displ, c, err = s.Dispatch(data[0]); err != nil {
            break
        }

        // Calculate and set PC from displacement
        s.pc = uint64(int(s.pc) + displ)
        s.RegWrite(PC, s.pc)
        cycles += c
    }
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
