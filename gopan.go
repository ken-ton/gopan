package gopan

import (
	"path"
	"runtime"
	"strings"
	"strconv"
	"regexp"
	"time"
	"math/rand"
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	BrandPrefix []ConfigBrandPrefix `json:"brand_prefix"`
}

type ConfigBrandPrefix struct {
	Name     string   `json:"name"`
	Prefixes []string `json:"prefix"`
	Length   int      `json:"length"`
}

func IsValid(pan string) bool {
	pan = clean(pan)

	var sum int
	var even bool

	l := len(pan)

	if l < 13 || l > 19 {
		return false
	}

	for i := l - 1; i > -1; i-- {
		mod, _ := strconv.Atoi(string(pan[i]))
		if even {
			mod *= 2
			if mod > 9 {
				mod = (mod % 10) + 1
			}
		}
		even = !even
		sum += mod
	}

	return sum % 10 == 0
}

func GetBrand(pan string) string {
	pan = clean(pan)
	l := len(pan)
	config := getConfig()
	for _, brandPrefix := range config.BrandPrefix {
		for _, prefix := range brandPrefix.Prefixes {
			if strings.HasPrefix(pan, prefix) && l == brandPrefix.Length {
				return brandPrefix.Name
			}
		}
	}
	return "Unknown"
}

func GetHiddenPan(pan string, hiddenChar ...string) string {
	pan = clean(pan)
	hc := "x"
	if len(hiddenChar) == 1 {
		hc = hiddenChar[0]
	}
	l := len(pan)
	h := []string{pan[:6], strings.Repeat(hc, l - 10), pan[l - 4:]}
	return strings.Join(h, "")
}

func Generate(brand ...string) string {
	rand.Seed(time.Now().UnixNano())

	config := getConfig()
	var bpx ConfigBrandPrefix

	flg := false
	if len(brand) > 0 {
		for _, brandPrefix := range config.BrandPrefix {
			if brandPrefix.Name == brand[0] {
				bpx = brandPrefix
				flg = true
				break
			}
		}
	}

	if flg == false {
		bpx = config.BrandPrefix[rand.Intn(len(config.BrandPrefix) - 1)]
	}

	pfx := bpx.Prefixes[rand.Intn(len(bpx.Prefixes) - 1)]
	fl := bpx.Length - len(pfx) - 1

	pansl := []string{pfx}
	for i := 0; i < fl; i++ {
		pansl = append(pansl, strconv.Itoa(rand.Intn(9)))
	}

	pan := strings.Join(pansl, "")

	sum := 0
	for k, _ := range pan {
		sub, _ := strconv.Atoi(string(pan[(len(pan) - 1) - k]))
		w := sub * (2 - (k % 2))
		if w < 10 {
			sum += w
		} else {
			sum += w - 9
		}
	}
	last := (10 - sum % 10) % 10

	return pan + strconv.Itoa(last)
}

func getConfig() Config {
	_, filename, _, _ := runtime.Caller(1)
	file, err := ioutil.ReadFile(path.Dir(filename) + "/config.json")
	if err != nil {
		panic(err)
	}
	var config Config
	json.Unmarshal(file, &config)
	return config;
}

func clean(pan string) string {
	r := regexp.MustCompile(`[\d]+`)
	return strings.Join(r.FindAllString(pan, -1), "")
}
