package loader

import (
    "testing"
    "io/ioutil"
    "bytes"
)

const smsFile = "../../bins/z80.sms.bin"

func TestSMSLoad(t *testing.T) {
    p, err := ioutil.ReadFile(smsFile)
    if err != nil {
        t.Fatal(err)
    }

    _, err = NewSMSLoader(bytes.NewReader(p), len(p))
    if err != nil {
        t.Fatal(err)
    }
}

func TestSMSSegments(t *testing.T) {
    p, err := ioutil.ReadFile(smsFile)
    if err != nil {
        t.Fatal(err)
    }

    sms, err := NewSMSLoader(bytes.NewReader(p), len(p))
    if err != nil {
        t.Fatal(err)
    }

    segments, err := sms.Segments()
    if err != nil {
        t.Fatal(err)
    }

    if len(segments) == 0 {
        t.Fatal("No segments")
    }
}

func TestMatchSMS(t *testing.T) {
    p, err := ioutil.ReadFile(smsFile)
    if err != nil {
        t.Fatal(err)
    }
    if MatchSMS(bytes.NewReader(p)) != true {
        t.Fatal("MatchSMS returned false, unexpectedly")
    }
}
