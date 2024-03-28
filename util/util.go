package util 

import (
    "fmt"
    "os"

    "github.com/spf13/viper"
    "github.com/raykov/mdtopdf"
    "github.com/raykov/gofpdf"
)

func ConvertFile(mdFile, pdfFile string) error {

    md, err := os.Open(mdFile)
    if err != nil{
        fmt.Println(err)
        return err
    }

    defer md.Close()

    pdf, err := os.Create(pdfFile)
    if err != nil {
        fmt.Println(err)
        return err
    }

    defer pdf.Close()

    // Get PDF settings 
    fontSize := viper.GetFloat64("fontsize")
    fontStyle := viper.GetString("fontstyle")

    textColorR := viper.GetInt("textcolorR")
    textColorG := viper.GetInt("textcolorG")
    textColorB := viper.GetInt("textcolorB")


    pageNumExtension := func(pdf *gofpdf.Fpdf) {
        pdf.SetFooterFunc(func() {
            left, _, right, bottom := pdf.GetMargins()
            width, height := pdf.GetPageSize()
            pNum := fmt.Sprint(pdf.PageNo())
            pdf.SetXY(width-left/2-pdf.GetStringWidth(pNum), height-bottom/2)
            pdf.SetFontSize(fontSize)
            pdf.SetTextColor(textColorR, textColorG, textColorB)
            pdf.SetFontStyle(fontStyle)
            pdf.SetRightMargin(0)
            pdf.Write(fontSize, pNum)
            pdf.SetRightMargin(right)
        })
    }

    err = mdtopdf.Convert(md, pdf, pageNumExtension)
    if err != nil {
        fmt.Println(err)
        return err
    }

    return nil
}