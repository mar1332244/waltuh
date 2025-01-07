package main

const (
    Space   = ' '
    Tab     = '\t'
    Newline = '\n'
)

type CommandType uint8

const (
    ReadChar CommandType = iota
    ReadNumber
    DumpChar
    DumpNumber
    PushNumber
    DupNumber
    SwapNumbers
    DiscardNumber
    AddNumbers
    SubNumbers
    MultNumbers
    DivNumbers
    ModNumbers
    Call
    GotoLabel
    JumpAlways
    JumpIfZero
    JumpIfNeg
    Return
    EndProg
    HeapStore
    HeapRetrieve
)

// swap the types in the map to make it make sense
var commands = map[string]CommandType{
    string([]byte{Tab, Newline, Tab, Space}):   ReadChar,
    string([]byte{Tab, Newline, Tab, Tab}):     ReadNumber,
    string([]byte{Tab, Newline, Space, Space}): DumpChar,
    string([]byte{Tab, Newline, Space, Tab}):   DumpNumber,

    string([]byte{Space, Space}):               PushNumber,
    string([]byte{Space, Newline, Space}):      DupNumber,
    string([]byte{Space, Newline, Tab}):        SwapNumbers,
    string([]byte{Space, Newline, Newline}):    DiscardNumber,

    string([]byte{Tab, Space, Space, Space}):   AddNumbers,
    string([]byte{Tab, Space, Space, Tab}):     SubNumbers,
    string([]byte{Tab, Space, Space, Newline}): MultNumbers,
    string([]byte{Tab, Space, Tab, Space}):     DivNumbers,
    string([]byte{Tab, Space, Tab, Tab}):       ModNumbers,

    string([]byte{Newline, Space, Space}):      Call,
    string([]byte{Newline, Space, Tab}):        GotoLabel,
    string([]byte{Newline, Space, Newline}):    JumpAlways,
    string([]byte{Newline, Tab, Space}):        JumpIfZero,
    string([]byte{Newline, Tab, Tab}):          JumpIfNeg,
    string([]byte{Newline, Tab, Newline}):      Return,
    string([]byte{Newline, Newline, Newline}):  EndProg,

    string([]byte{Tab, Tab, Space}):            HeapStore,
    string([]byte{Tab, Tab, Tab}):              HeapRetrieve,
}
