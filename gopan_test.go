package gopan

import (
	"testing"
)

type testData struct {
	Pan    string
	Hidden string
	Brand  string
	Valid  bool
}

var td = []*testData{
	&testData{"30569309025904", "305693xxxx5904", "DINERS CLUB", true},
	&testData{"38520000023237", "385200xxxx3237", "DINERS CLUB", true},
	&testData{"3411 111111 11111", "341111xxxxx1111", "AMERICAN EXPRESS", true},
	&testData{"3714 496353 98431", "371449xxxxx8431", "AMERICAN EXPRESS", true},
	&testData{"3782 8224 6310 005", "378282xxxxx0005", "AMERICAN EXPRESS", true},
	&testData{"3787_3449_3671_000", "378734xxxxx1000", "AMERICAN EXPRESS", true},
	&testData{"3530_1113_3330_0000", "353011xxxxxx0000", "JCB", true},
	&testData{"3566_0020_2036_0505", "356600xxxxxx0505", "JCB", true},
	&testData{"4000 0000 0000 0002", "400000xxxxxx0002", "VISA", true},
	&testData{"4000 0000 0000 0010", "400000xxxxxx0010", "VISA", true},
	&testData{"4000 0000 0000 0028", "400000xxxxxx0028", "VISA", true},
	&testData{"4000 0000 0000 0036", "400000xxxxxx0036", "VISA", true},
	&testData{"4000 0000 0000 0069", "400000xxxxxx0069", "VISA", true},
	&testData{"4000 0000 0000 0101", "400000xxxxxx0101", "VISA", true},
	&testData{"4000 0000 0000 0119", "400000xxxxxx0119", "VISA", true},
	&testData{"4000-0000-0000-0127", "400000xxxxxx0127", "VISA", true},
	&testData{"4000-0000-0000-0341", "400000xxxxxx0341", "VISA", true},
	&testData{"4000-0000-0008-0202", "400000xxxxxx0202", "VISA", true},
	&testData{"4000-0000-0008-0327", "400000xxxxxx0327", "VISA", true},
	&testData{"4000-0000-0008-0350", "400000xxxxxx0350", "VISA", true},
	&testData{"4000-0000-0008-0319", "400000xxxxxx0319", "VISA", true},
	&testData{"4012-8888-8888-1881", "401288xxxxxx1881", "VISA", true},
	&testData{"4111-1111-1111-1111", "411111xxxxxx1111", "VISA", true},
	//&testData{"4222222222222", "422222xxx2222", "VISA", true},
	&testData{"4242424242424242", "424242xxxxxx4242", "VISA", true},
	&testData{"5105105105105100", "510510xxxxxx5100", "MASTERCARD", true},
	&testData{"5111111111111118", "511111xxxxxx1118", "MASTERCARD", true},
	&testData{"5431111111111111", "543111xxxxxx1111", "MASTERCARD", true},
	&testData{"5555555555554444", "555555xxxxxx4444", "MASTERCARD", true},
	//&testData{"6011601160116611", "601160xxxxxx6611", "DISCOVER", true},
	//&testData{"6011111111111117", "601111xxxxxx1117", "DISCOVER", true},
	&testData{"6011000990139424", "601100xxxxxx9424", "DISCOVER", true},
	//&testData{"6111111111111116", "611111xxxxxx1116", "DISCOVER", true},
	&testData{"1234123412341234", "123412xxxxxx1234", "Unknown", false},
}

func TestIsValid(t *testing.T) {
	for _, data := range td {
		result := IsValid(data.Pan)

		if result != data.Valid {
			t.Error(
				"PAN:", data.Pan,
				"valid:", data.Valid,
				"result:", result,
			)
		}
	}
}

func TestGetBrand(t *testing.T) {
	for _, data := range td {
		result := GetBrand(data.Pan)

		if result != data.Brand {
			t.Error(
				"PAN:", data.Pan,
				"expected:", data.Brand,
				"result:", result,
			)
		}
	}
}

func TestGetHiddenPan(t *testing.T) {
	for _, data := range td {
		result := GetHiddenPan(data.Pan)

		if result != data.Hidden {
			t.Error(
				"PAN:", data.Pan,
				"expected:", data.Hidden,
				"result:", result,
			)
		}
	}
}

func TestGenerate(t *testing.T) {
	for i := 0; i < 100; i++ {
		pan := Generate()
		resultValid := IsValid(pan)
		resultBrand := GetBrand(pan)
		if !resultValid || resultBrand == "Unknown" {
			t.Error(
				"PAN:", pan,
				"valid:", resultValid,
				"brand:", resultBrand,
			)
		}
	}
}

func TestGenerateBrand(t *testing.T) {
	brandList := []string{"AMERICAN EXPRESS", "CHINA UNION PAY", "DINERS CLUB", "DISCOVER", "JCB", "MASTERCARD", "VISA"}
	for _, brand := range brandList {
		for i := 0; i < 100; i++ {
			pan := Generate(brand)
			resultValid := IsValid(pan)
			resultBrand := GetBrand(pan)
			if !resultValid || brand != resultBrand {
				t.Error(
					"Brand:", brand,
					"PAN:", pan,
					"valid:", resultValid,
					"brand:", resultBrand,
				)
			}
		}
	}
}
