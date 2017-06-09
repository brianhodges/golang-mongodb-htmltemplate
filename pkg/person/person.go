package person

import(
    "encoding/base64"
    qrcode "github.com/skip2/go-qrcode"
)

type Person struct {
    First_Name string
    Last_Name string
    Email string
    BTC_Address string
    IP_Address string
}

func (p Person) Full_Name() string {
    return p.First_Name + " " + p.Last_Name
}

func (p Person) BTC_QR() string {
    return generateQRCode(p.BTC_Address)
}

func generateQRCode(input string) (string) {
    var png []byte
    png, err := qrcode.Encode(input, qrcode.Medium, 128)
    if err != nil {
        panic(err)
    }
    imgBase64Str := base64.StdEncoding.EncodeToString(png)
    return imgBase64Str
}