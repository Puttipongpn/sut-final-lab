package sutfinallab

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

type Costumer struct {
	Name  string `valid:"length(2|10)"`
	Phone string `valid:"numeric"`
	Email string `valid:"email"`
}

func TestPaymentNumber(t *testing.T) {
	g := NewGomegaWithT(t)

	Costumer1 := Costumer{
		Name:  "aasdf",
		Phone: "501222",
		Email: "fam@gmail.com",
	}

	ok, err := govalidator.ValidateStruct(Costumer1)

	g.Expect(ok).ToNot(BeTrue())

	g.Expect(err).ToNot(BeNil())

	g.Expect(err.Error()).To(Equal("Name not blank"))
}

// func TestEmail(t *testing.T) {
// 	g := NewGomegaWithT(t)

// 	Costumer2 := Costumer{
// 		Name:  "dfg",
// 		Phone: "IDB",
// 		Email: "famgmail.com", //ผิด
// 	}

// 	ok, err := govalidator.ValidateStruct(Costumer2)

// 	g.Expect(ok).ToNot(BeTrue())

// 	g.Expect(err).ToNot(BeNil())

// 	g.Expect(err.Error()).To(Equal("NAme not blank"))
// }
