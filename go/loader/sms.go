package loader

import (
    "io"
    "io/ioutil"
    "bytes"
    "encoding/binary"
    "github.com/lunixbochs/usercorn/go/models"
    "github.com/lunixbochs/usercorn/go/models/cpu"
)

var smsMagic = []byte{'T', 'M', 'R', ' ', 'S', 'E', 'G', 'A'}

type SMSLoader struct {
    LoaderBase
    Size int
    ROM []byte
}

// MatchSMS reads data from the reader and returns true if the binary is a valid SMS ROM
func MatchSMS(r io.ReaderAt) bool {
    hdrOffsets := []int64{0x1ff0, 0x3ff0, 0x7ff0}
    buf := make([]byte, len(smsMagic))
    for i := 0; i < len(hdrOffsets); i++ {
        _, err := r.ReadAt(buf, hdrOffsets[i])
        if err != nil {
            return false
        }

        if bytes.Compare(buf, smsMagic) == 0 {
            return true
        }
    }
    return false
}

// NewSMSLoader returns a pointer to a SMSLoader struct
func NewSMSLoader(r io.ReaderAt) (models.Loader, error) {
    rom, err := ioutil.ReadAll(io.NewSectionReader(r, 0, 0xc000))
    if err != nil {
        return nil, err
    }
    romSize := len(rom)
    return &SMSLoader{
        LoaderBase: LoaderBase{
            arch: "sms",
            bits: 16,
            byteOrder: binary.LittleEndian,
            os: "sms",
            entry: 0,
        },
        Size: romSize,
        ROM: rom,
    }, nil
}

// Segments is a SMSLoader method for getting SMS segment data
func (l *SMSLoader) Segments() ([]models.SegmentData, error) {
    var segs []models.SegmentData

    // ROM
    segs = append(segs, models.SegmentData{
        Off: 0,
        Addr: 0,
        Size: uint64(len(l.ROM)),
        Prot: cpu.PROT_READ|cpu.PROT_EXEC,
        DataFunc: func() ([]byte, error) {
            return l.ROM, nil
        },
    })

    return segs, nil
}
