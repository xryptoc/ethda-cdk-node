package blob

// DefaultValues is the default configuration
const DefaultValues = `
[ZkBlobSender]
WaitPeriodSendBlob = "5s"
WaitAfterBlobSent = "1s"
GasOffset = 1000
DasAddress = "0x0"
ZkBlobAddress = "0x0"
PrivateKey = {Path = "/pk/blob.keystore", Password = "testonly"}
`
