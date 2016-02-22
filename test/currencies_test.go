package test

import (
	"testing"

	"github.com/syscrusher/fake"
)

func TestCurrencies(t *testing.T) {
	for _, lang := range fake.GetLangs() {
		fake.SetLang(lang)

		v := fake.Currency()
		if v == "" {
			t.Errorf("Currency failed with lang %s", lang)
		}

		v = fake.CurrencyCode()
		if v == "" {
			t.Errorf("CurrencyCode failed with lang %s", lang)
		}
	}
}
