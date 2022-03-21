package basis

type Number interface {
    Int | Uint | Float
}

type Int interface {
    int | int8 | int32 | int64
}

type Uint interface {
    uint | uint8 | uint32 | uint64
}

type Float interface {
    float32 | float64
}

