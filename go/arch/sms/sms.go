package sms

import (
    "github.com/lunixbochs/usercorn/go/models"
    "fmt"
)

func SMSInit(u models.Usercorn, args, env []string) error {
    return nil
}

func SMSKernels(u models.Usercorn) []interface{} {
    return []interface{}{}
}

func SMSInterrupt(u models.Usercorn, cause uint32) {
    fmt.Printf("Interrupt fired: %x\n", cause)
}

func init() {
    Arch.RegisterOS(&models.OS{
        Name: "sms",
        Kernels: SMSKernels,
        Init: SMSInit,
        Interrupt: SMSInterrupt,
    })
}
