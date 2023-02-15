package sutfinallab

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

type Costumer struct {
	gorm.Model
	Name       string `valid:"required~Name not blank"`
	CostumerID string
	Email      string `valid:"email"`
}

func TestPaymentNumber(t *testing.T) {
	g := NewGomegaWithT(t)

	Costumer1 := Costumer{
		Name:       "",
		CostumerID: "501222",
		Email:      "fam@gmail.com",
	}

	ok, err := govalidator.ValidateStruct(Costumer1)

	g.Expect(ok).ToNot(BeTrue())

	g.Expect(err).ToNot(BeNil())

	g.Expect(err.Error()).To(Equal("Name not blank"))
}

func TestEmail(t *testing.T) {
	g := NewGomegaWithT(t)

	Costumer2 := Costumer{
		Name:       "dfg",
		CostumerID: "LDB",
		Email:      "fam@gmail.com", //ผิด
	}

	ok, err := govalidator.ValidateStruct(Costumer2)

	g.Expect(ok).ToNot(BeTrue())

	g.Expect(err).ToNot(BeNil())

	g.Expect(err.Error()).To(Equal("NAme not blank"))
}
