package sms

import (
    "bytes"
    "strings"
    "github.com/lunixbochs/usercorn/go/models"
)

type op struct {
    name string
    arg int
}

type ins struct {
    addr uint64
    op byte
    name string
    args []arg
    bytes []byte
}

func (i *ins) Addr() uint64 {
    return i.addr
}

func (i *ins) Bytes() []byte {
    return i.bytes
}

func (i *ins) Mnemonic() string {
    return i.name
}

func (i *ins) OpStr() string {
    var args []string
    for _, a := range i.args {
        args = append(args, a.String())
    }
    return strings.Join(args, ", ")
}

type insReader struct {
    *bytes.Reader
    err error
    addr uint64
}

type arg interface {
    String() string
}

type Dis struct {}


func (r *insReader) tell() int64 {
    return r.Size() - int64(r.Len())
}

func (r *insReader) ins() models.Ins {
    start := r.tell()
    b, err := r.ReadByte()
    if err != nil {
        return nil // EOF?
    }

    return &ins{
        addr: r.addr,
        op: b,
        name: "derp",
        args: []arg{},
        bytes: make([]byte, r.tell()-start),
    }
}

func (d *Dis) Dis(mem []byte, addr uint64) ([]models.Ins, error) {
    r := &insReader{
        addr: addr,
        Reader: bytes.NewReader(mem),
    }
    var ret []models.Ins
    for {
        ins := r.ins()
        if ins == nil || r.err != nil {
            break
        }
        ret = append(ret, ins)
    }
    return ret, nil
}
