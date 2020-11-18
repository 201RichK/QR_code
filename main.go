package main

import (
	"fmt"
	"html/template"
	"image/png"
	"net/http"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

var tpl *template.Template

type totmpl struct {
	Title    string
	FileName string
}

func init() {
	tpl = template.Must(template.ParseGlob("templ/*.html"))
}

func main() {
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./img"))))
	http.HandleFunc("/", homehandler)
	http.HandleFunc("/post", codegenerateHandler)

	println("Listening to 127.0.0.1:8090")
	http.ListenAndServe(":8090", nil)
}

func homehandler(w http.ResponseWriter, r *http.Request) {
	tmpl := totmpl{
		Title:    "Qr_code générator",
		FileName: "standar.png",
	}

	tpl.ExecuteTemplate(w, "index.html", tmpl)
}

func codegenerateHandler(w http.ResponseWriter, r *http.Request) {

	nomProduit := r.FormValue("name")
	energie := r.FormValue("energie")
	cholesterole := r.FormValue("cholesterole")
	lipide := r.FormValue("lipide")
	proteine := r.FormValue("proteine")
	sodium := r.FormValue("sodium")
	fer := r.FormValue("fer")
	poid := r.FormValue("poid")

	var infoProduit string = fmt.Sprintf(`
			INFORMATION SUR LE PRODUIT

	VALEUR NUTRITIONNELLES POUR 100G

	Nom Du Produit : %s
	Teneur en Energie: %s
	Teneur en cholesterole: %s
	Teneur en lipide: %s
	Teneur en proteine: %s
	Teneur en sodium: %s
	Teneur en fer: %s

	Poid net: %s

	Url: https://cocopragel.com
	Email: contact@cocopragel.com
	Contact: (+225) 07 18 81 91 / 59 72 80 79

	Produit a Bonoua en Côte d'ivoire par Cocopragel
	`, nomProduit, energie, cholesterole, lipide, proteine, sodium, fer, poid)

	// create the output file
	qrFname, _ := os.Create(fmt.Sprintf("./img/%s.png", nomProduit))

	qrCode, _ := qr.Encode(infoProduit, qr.L, qr.Auto)
	qrCode, _ = barcode.Scale(qrCode, 512, 512)

	png.Encode(qrFname, qrCode)

	tmpl := totmpl{
		Title:    "Qr_code générator",
		FileName: fmt.Sprintf("%s.png", nomProduit),
	}
	tpl.ExecuteTemplate(w, "index.html", tmpl)

}
