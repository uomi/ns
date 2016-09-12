package main

import (
    "fmt"
    "os"
)

func unpackDns(msg []byte, dnsType uint16) (domain string, id uint16, cname string) {
    d := new(dnsMsg)
    if !d.Unpack(msg) {
        return
    }

    id = d.id

    if len(d.question) < 1 {
        return
    }

    domain = d.question[0].Name
    if len(domain) < 1 {
        return
    }

    _, addrs, err := answer(domain, "server", d, dnsType)
    if err == nil {
        cname = addrs[0].(*dnsRR_NS).Ns
    }
    return
}

func packDns(domain string, id uint16, dnsType uint16) []byte {

    out := new(dnsMsg)
    out.id = id
    out.recursion_desired = true
    out.question = []dnsQuestion{
        {domain, dnsType, dnsClassINET},
    }

    msg, ok := out.Pack()
    if !ok {
        fmt.Fprintf(os.Stderr, "can't pack domain %s\n", domain)
        os.Exit(1)
    }
    return msg
}
