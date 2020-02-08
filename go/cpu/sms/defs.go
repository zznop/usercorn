package sms

const (
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

const (
    OP_NOP = iota
    OP_LD
    OP_INC
    OP_DEC
    OP_RLCA
    OP_EX
    OP_ADD
    OP_RRCA
    OP_DJNZ
    OP_RLA
    OP_JR
    OP_RRA
    OP_DAA
    OP_CPL
    OP_SCF
    OP_CCF
    OP_HALT
    OP_ADC
    OP_SUB
    OP_SBC
    OP_AND
    OP_XOR
    OP_OR
    OP_CP
    OP_RET
    OP_POP
    OP_JP
    OP_CALL
    OP_PUSH
    OP_RST
    OP_IDKYET
    OP_OUT
    OP_EXX
    OP_IN
    OP_DI
    OP_EI
)

var OperationMap = map[uint8]int{
    0x00: OP_NOP,
    0x01: OP_LD,
    0x02: OP_LD,
    0x03: OP_INC,
    0x04: OP_INC,
    0x05: OP_DEC,
    0x06: OP_LD,
    0x07: OP_RLCA,
    0x08: OP_EX,
    0x09: OP_ADD,
    0x0a: OP_LD,
    0x0b: OP_DEC,
    0x0c: OP_INC,
    0x0d: OP_DEC,
    0x0e: OP_LD,
    0x0f: OP_RRCA,
    0x10: OP_DJNZ,
    0x11: OP_LD,
    0x12: OP_LD,
    0x13: OP_INC,
    0x14: OP_INC,
    0x15: OP_DEC,
    0x16: OP_LD,
    0x17: OP_RLA,
    0x18: OP_JR,
    0x19: OP_ADD,
    0x1a: OP_LD,
    0x1b: OP_DEC,
    0x1c: OP_INC,
    0x1d: OP_DEC,
    0x1e: OP_LD,
    0x1f: OP_RRA,
    0x20: OP_JR,
    0x21: OP_LD,
    0x22: OP_LD,
    0x23: OP_INC,
    0x24: OP_INC,
    0x25: OP_DEC,
    0x26: OP_LD,
    0x27: OP_DAA,
    0x28: OP_JR,
    0x29: OP_ADD,
    0x2a: OP_LD,
    0x2b: OP_DEC,
    0x2c: OP_INC,
    0x2d: OP_DEC,
    0x2e: OP_LD,
    0x2f: OP_CPL,
    0x30: OP_JR,
    0x31: OP_LD,
    0x32: OP_LD,
    0x33: OP_INC,
    0x34: OP_INC,
    0x35: OP_DEC,
    0x36: OP_LD,
    0x37: OP_SCF,
    0x38: OP_JR,
    0x39: OP_ADD,
    0x3a: OP_LD,
    0x3b: OP_DEC,
    0x3c: OP_INC,
    0x3d: OP_DEC,
    0x3e: OP_LD,
    0x3f: OP_CCF,
    0x40: OP_LD,
    0x41: OP_LD,
    0x42: OP_LD,
    0x43: OP_LD,
    0x44: OP_LD,
    0x45: OP_LD,
    0x46: OP_LD,
    0x47: OP_LD,
    0x48: OP_LD,
    0x49: OP_LD,
    0x4a: OP_LD,
    0x4b: OP_LD,
    0x4c: OP_LD,
    0x4d: OP_LD,
    0x4e: OP_LD,
    0x4f: OP_LD,
    0x50: OP_LD,
    0x51: OP_LD,
    0x52: OP_LD,
    0x53: OP_LD,
    0x54: OP_LD,
    0x55: OP_LD,
    0x56: OP_LD,
    0x57: OP_LD,
    0x58: OP_LD,
    0x59: OP_LD,
    0x5a: OP_LD,
    0x5b: OP_LD,
    0x5c: OP_LD,
    0x5d: OP_LD,
    0x5e: OP_LD,
    0x5f: OP_LD,
    0x60: OP_LD,
    0x61: OP_LD,
    0x62: OP_LD,
    0x63: OP_LD,
    0x64: OP_LD,
    0x65: OP_LD,
    0x66: OP_LD,
    0x67: OP_LD,
    0x68: OP_LD,
    0x69: OP_LD,
    0x6a: OP_LD,
    0x6b: OP_LD,
    0x6c: OP_LD,
    0x6d: OP_LD,
    0x6e: OP_LD,
    0x6f: OP_LD,
    0x70: OP_LD,
    0x71: OP_LD,
    0x72: OP_LD,
    0x73: OP_LD,
    0x74: OP_LD,
    0x75: OP_LD,
    0x76: OP_HALT,
    0x77: OP_LD,
    0x78: OP_LD,
    0x79: OP_LD,
    0x7a: OP_LD,
    0x7b: OP_LD,
    0x7c: OP_LD,
    0x7d: OP_LD,
    0x7e: OP_LD,
    0x7f: OP_LD,
    0x80: OP_ADD,
    0x81: OP_ADD,
    0x82: OP_ADD,
    0x83: OP_ADD,
    0x84: OP_ADD,
    0x85: OP_ADD,
    0x86: OP_ADD,
    0x87: OP_ADD,
    0x88: OP_ADC,
    0x89: OP_ADC,
    0x8a: OP_ADC,
    0x8b: OP_ADC,
    0x8c: OP_ADC,
    0x8d: OP_ADC,
    0x8e: OP_ADC,
    0x8f: OP_ADC,
    0x90: OP_SUB,
    0x91: OP_SUB,
    0x92: OP_SUB,
    0x93: OP_SUB,
    0x94: OP_SUB,
    0x95: OP_SUB,
    0x96: OP_SUB,
    0x97: OP_SUB,
    0x98: OP_SBC,
    0x99: OP_SBC,
    0x9a: OP_SBC,
    0x9b: OP_SBC,
    0x9c: OP_SBC,
    0x9d: OP_SBC,
    0x9e: OP_SBC,
    0x9f: OP_SBC,
    0xa0: OP_AND,
    0xa1: OP_AND,
    0xa2: OP_AND,
    0xa3: OP_AND,
    0xa4: OP_AND,
    0xa5: OP_AND,
    0xa6: OP_AND,
    0xa7: OP_AND,
    0xa8: OP_XOR,
    0xa9: OP_XOR,
    0xaa: OP_XOR,
    0xab: OP_XOR,
    0xac: OP_XOR,
    0xad: OP_XOR,
    0xae: OP_XOR,
    0xaf: OP_XOR,
    0xb0: OP_OR,
    0xb1: OP_OR,
    0xb2: OP_OR,
    0xb3: OP_OR,
    0xb4: OP_OR,
    0xb5: OP_OR,
    0xb6: OP_OR,
    0xb7: OP_OR,
    0xb8: OP_CP,
    0xb9: OP_CP,
    0xba: OP_CP,
    0xbb: OP_CP,
    0xbc: OP_CP,
    0xbd: OP_CP,
    0xbe: OP_CP,
    0xbf: OP_CP,
    0xc0: OP_RET,
    0xc1: OP_POP,
    0xc2: OP_JP,
    0xc3: OP_JP,
    0xc4: OP_CALL,
    0xc5: OP_PUSH,
    0xc6: OP_ADD,
    0xc7: OP_RST,
    0xc8: OP_RET,
    0xc9: OP_RET,
    0xca: OP_JP,
    0xcb: OP_IDKYET,
    0xcc: OP_CALL,
    0xcd: OP_CALL,
    0xce: OP_ADC,
    0xcf: OP_RST,
    0xd0: OP_RET,
    0xd1: OP_POP,
    0xd2: OP_JP,
    0xd3: OP_OUT,
    0xd4: OP_CALL,
    0xd5: OP_PUSH,
    0xd6: OP_SUB,
    0xd7: OP_RST,
    0xd8: OP_RET,
    0xd9: OP_EXX,
    0xda: OP_JP,
    0xdb: OP_IN,
    0xdc: OP_CALL,
    0xdd: OP_IDKYET,
    0xde: OP_SBC,
    0xdf: OP_RST,
    0xe0: OP_RET,
    0xe1: OP_POP,
    0xe2: OP_JP,
    0xe3: OP_EX,
    0xe4: OP_CALL,
    0xe5: OP_PUSH,
    0xe6: OP_AND,
    0xe7: OP_RST,
    0xe8: OP_RET,
    0xe9: OP_JP,
    0xea: OP_JP,
    0xeb: OP_EX,
    0xec: OP_CALL,
    0xed: OP_IDKYET,
    0xee: OP_XOR,
    0xef: OP_RST,
    0xf0: OP_RET,
    0xf1: OP_POP,
    0xf2: OP_JP,
    0xf3: OP_DI,
    0xf4: OP_CALL,
    0xf5: OP_PUSH,
    0xf6: OP_OR,
    0xf7: OP_RST,
    0xf8: OP_RET,
    0xf9: OP_LD,
    0xfa: OP_JP,
    0xfb: OP_EI,
    0xfc: OP_CALL,
    0xfd: OP_IDKYET,
    0xfe: OP_CP,
    0xff: OP_RST,
}
