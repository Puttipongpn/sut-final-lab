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
	CostumerID string `valid:"matches(^[L|M|H][0-9]+)"`
	Email      string `valid:"email"`
}

func TestCostumer(t *testing.T) {
	g := NewGomegaWithT(t)

	Costumer1 := Costumer{
		Name:       "asd",
		CostumerID: "501222",
		Email:      "fam@gmail.com",
	}

	ok, err := govalidator.ValidateStruct(Costumer1)

	g.Expect(ok).ToNot(BeTrue())

	g.Expect(err).ToNot(BeNil())

	g.Expect(err.Error()).To(Equal("Name not blank"))
}

func TestCostumername(t *testing.T) {
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

func TestCostumerID(t *testing.T) {
	g := NewGomegaWithT(t)

	Costumer1 := Costumer{
		Name:       "asd",
		CostumerID: "d23456",
		Email:      "fam@gmail.com",
	}

	ok, err := govalidator.ValidateStruct(Costumer1)

	g.Expect(ok).ToNot(BeTrue())

	g.Expect(err).ToNot(BeNil())

	g.Expect(err.Error()).To(Equal("CostumerID: d23456 does not validate as matches(^[L|M|H][0-9]+)"))
}
