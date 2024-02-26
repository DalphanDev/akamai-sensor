package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type variables struct {
	tAD []string
	pw  pw     // pw - line 5068 - entire object is constant
	p3  int    // p3 - line 2432 - constant of 36
	SY  int    // SY - line 2432 - constant of 48
	sAD string // sAD - line 2986
	c0  int
	bw  string
	r8D string // ⭐ our sensor data
	f8D string // f8D - line 3003
	C8D string // C8D - line 3829. DD.IV gives us our string. I believe its a constant.
	hf  int32
	hk  int32
	MHD int
	p6D []int
	h8D int
	A6D int
}

type pw struct {
	Index0  int
	Index1  int
	Index2  int
	Index3  int
	Index4  int
	Index5  int
	Index6  int
	Index7  int
	Index8  int
	Index9  int
	Index10 int
	Index11 int
	Index12 int
	Index13 int
	Index14 int
	Index15 int
	Index16 int
	Index17 int
	Index18 int
	Index19 int
	Index20 int
	Index21 int
	Index22 int
	Index23 int
	Index24 int
	Index25 int
	Index26 int
	Index27 int
	Index28 int
	Index29 int
	Index30 int
	Index31 int
	Index32 int
}

func initVariables() (vars variables) {
	pw := pw{
		Index0:  4095,
		Index1:  1487,
		Index2:  300000,
		Index3:  0,
		Index4:  1,
		Index5:  100000,
		Index6:  20,
		Index7:  16,
		Index8:  6,
		Index9:  9,
		Index10: 12,
		Index11: 21,
		Index12: 8888888,
		Index13: 7777777,
		Index14: 3,
		Index15: 10000,
		Index16: 100,
		Index17: 2,
		Index18: 10,
		Index19: 1000,
		Index20: 4294967296,
		Index21: 4,
		Index22: 17,
		Index23: 24,
		Index24: 3600000,
		Index25: 65535,
		Index26: 65793,
		Index27: 4282663,
		Index28: 8388607,
		Index29: 4294967295,
		Index30: 8,
		Index31: 90,
		Index32: 3000,
	}

	p6D := []int{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
		-1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
		-1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
		-1, -1, -1, 0, 1, -1, 2, 3, 4, 5,
		-1, 6, 7, 8, 9, 10, 11, 12, 13, 14,
		15, 16, 17, 18, 19, 20, 21, 22, 23, 24,
		25, 26, 27, 28, 29, 30, 31, 32, 33, 34,
		35, 36, 37, 38, 39, 40, 41, 42, 43, 44,
		45, 46, 47, 48, 49, 50, 51, 52, 53, 54,
		55, 56, -1, 58, 59, 60, 61, 62, 63, 64,
		65, 66, 67, 68, 69, 70, 71, 72, 73, 74,
		75, 76, 77, 78, 79, 80, 81, 82, 83, 84,
		85, 86, 87, 88, 89, 90, 91}

	fmt.Println("Initializing variables...")

	varStruct := variables{
		hf:  47,
		hk:  57,
		c0:  2,
		p3:  36,
		pw:  pw,
		bw:  "abcdefghijklmnopaqrstuvxyzABCDEFGHIJKLMNOPAQRSTUVXYZ!@#%&-_=;:<>,~000000000000000000001111111111111111111133333333333333333333444444444444444444445555555555555555555566666666666666666666777777777777777777778888888888888888888899999999999999999999",
		C8D: "xY5xDk8R309uloCLViedwg==",
		f8D: gen_f8D(47, 57),
		sAD: gen_sAD("abcdefghijklmnopaqrstuvxyzABCDEFGHIJKLMNOPAQRSTUVXYZ!@#%&-_=;:<>,~000000000000000000001111111111111111111133333333333333333333444444444444444444445555555555555555555566666666666666666666777777777777777777778888888888888888888899999999999999999999"),
		MHD: 339021,
		tAD: gen_tAD(),
		p6D: p6D,
	}

	return varStruct
}

// This is the line of code from line 2986. Specifically the "r8D =" part.
func step1(vars variables) (updatedVars variables) {

	// First, we need to make sure our tAD is generated. This variable changes depending on the site, and seems to fingerprint our browser.
	// This array seems to fingerprint our browser to generate our tAD.
	// [DD.gA(vI(vI(bJ)), fh, Qf, lE, TJ), vAD, DD.UA(fq, fh, mq, lE, Wq), n8D, DD.Ip(zh, FE), fdD, DD.wp(Gk, UT), wAD, DD.Qp(Bk, W7), zdD, DD.Mp.apply(null, [AT, jE, vI([])]), WdD, DD.cp(gE, C0), Z8D, DD.Np(Wh, mv, hs), QAD, DD.Hn(ch, Z0), l8D, DD.Dn(ck, Ef, Xx), kdD, DD.Bn(Gz, Tx, MO), X8D, DD.YA.apply(null, [NE, WK, BE, lE, UE, fh]), NAD, DD.An(k7, Cf, tq), VRD, DD.dn(Is, l4), cAD, DD.Rn(Nv, LJ, vI(PE)), j8D, DD.Vn.apply(null, [ws, lE, Hk]), TRD, DD.bn(v7, Kf), CRD, DD.FA(Tf, BE, hf, lE, vf, fh), hRD, DD.GA(nv, fh, gx, lE, Zv), IRD, DD.Pn.apply(null, [zq, Qs]), g8D, DD.pn(K7, KW), NRD, DD.nn.apply(null, [Wm, BK, Cf]), A8D, DD.Zn(Ms, Nq, Az), P8D, DD.ln(BK, C7), U8D, DD.Xn.call(null, bv, cs), BdD[DD.x6.apply(null, [z7, vI(vI(bJ)), TJ])], DD.jn.call(null, HE, Ns), XRD, DD.gn(RC, nE, vf), wdD, DD.Un.call(null, gT, bJ), Y8D]

	// Then, join the array together, and add the sAD onto the end of it.

	joined_tAD := strings.Join(vars.tAD, ",")

	vars.r8D = joined_tAD + vars.sAD

	return vars
}

// This is the line of code from line 3010
func step2(vars variables) (updatedVars variables) {

	r8D2 := (vars.f8D + vars.C8D + vars.sAD + strconv.Itoa(vars.MHD) + vars.sAD + vars.r8D)

	vars.r8D = strconv.Itoa(vars.c0) + vars.sAD + strconv.Itoa(vars.c0) + vars.sAD + r8D2

	return vars
}

func step3(vars variables) (updatedVars variables) {

	// We need to initialize h8D. This will be used in step 5.
	vars.h8D = PBD()

	D6D := strings.Split(vars.r8D, ",")

	Q8D := 3289904

	M8D := 0
	c8D := 0
	N8D := ""

	for i := 0; i < len(D6D); i++ {
		M8D = ((Q8D >> 8) & 65535) % len(D6D)
		Q8D *= 65793
		Q8D &= 4294967295
		Q8D += 4282663
		Q8D &= 8388607

		c8D = ((Q8D >> 8) & 65535) % len(D6D)

		Q8D *= 65793
		Q8D &= 4294967295
		Q8D += 4282663
		Q8D &= 8388607

		N8D = D6D[M8D]
		D6D[M8D] = D6D[c8D]
		D6D[c8D] = N8D

	}

	B6D := strings.Join(D6D, ",")

	vars.r8D = B6D

	// We need to update h8D.

	vars.h8D = PBD() - vars.h8D // The amount of milliseconds the function took to run?

	return vars
}

func step4(vars variables) (updatedVars variables) {

	// We need to initialize A6D. It will be used in step 5.
	vars.A6D = PBD()

	R6D := vars.r8D
	V6D := 3683632
	b6D := " !#$%&()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[]^_`abcdefghijklmnopqrstuvwxyz{|}~"
	n6D := ""

	for Z6D := 0; Z6D < len(R6D); Z6D++ {
		l6D := string(R6D[Z6D])
		X6D := ((V6D >> 8) & 65535)

		V6D *= 65793
		V6D &= 4294967295
		V6D += 4282663
		V6D &= 8388607

		j6D := vars.p6D[(rune(R6D[Z6D]))]

		tempRuneSlice := []rune(l6D)
		g6D := tempRuneSlice[0]

		if g6D >= 32 {
			if g6D < 127 {
				j6D = vars.p6D[g6D]
			}
		}

		if j6D >= 0 {
			j6D += (X6D % len(b6D))
			j6D %= len(b6D)
			l6D = string(b6D[j6D])
		}

		n6D += l6D
	}

	vars.r8D = n6D

	vars.A6D = PBD() - vars.A6D

	return vars
}

func step5(vars variables) (updatedVars variables) {

	// "".concat(randomnumber + , + 20194677 + , + 0 + , + h8D + , + A6D + , + 5)

	// randomnumber = PBD() - xAD

	// xAD = 1687273519055

	Y6D := strconv.Itoa(PBD()) + "," + "20194677" + "," + "0" + "," + strconv.Itoa(vars.h8D) + "," + strconv.Itoa(updatedVars.A6D) + "," + "5"

	final_r8D := "2;3683632;3289904;" + Y6D + ";" + vars.r8D

	vars.r8D = final_r8D

	return vars
}

func PBD() (number int) {
	// Go is so fast, that there isn't a difference in time between these calls from different functions. We will have to artificially spoof these numbers like go is taking
	// longer than it actually is.

	// Window.Date.now()

	RnD := time.Now().UnixNano() / 1000000

	return int(RnD)
}

/*
This function imitates the OI function call on line 2986.

Assuming it just randomly chooses a character. If there is more math to choosing a character, it will be added later.
It is correct. It simply chooses a random character from the string. Now to figure out how many characters it chooses.
The amount of characters between each comma, is a random number from 3 - 5.
The amount of commas themselves, is affected by a random number from 3 - 5.
*/
func gen_sAD(bw string) (sAD string) {

	rand.Seed(time.Now().UnixNano())

	genned_sAD := ","

	commaLength := rand.Intn(3) + 3 //random number between 3 and 5

	for i := 0; i < commaLength; i++ {

		ministringLength := rand.Intn(3) + 3 //random number between 3 and 5

		for c := 0; c < ministringLength; c++ {
			index := rand.Intn(len(bw))
			genned_sAD = genned_sAD + string(bw[index])
		}

		genned_sAD = genned_sAD + ","

	}

	fmt.Println("genned_sAD: " + genned_sAD)

	return genned_sAD
}

// This function is imitating the nJ(T8D, x8D(s8D, T8D)) expression on line 3004.
func gen_f8D(hf int32, hk int32) (f8D string) {
	// First, we need to get our T8D variable.
	T8D := "7a74G7m23Vrp0o5c" // APPEARS TO BE CONSTANT. Possibly daily/weekly rotation
	// Now, we need to get out s8D variable.
	s8D := 468593 // Not constant. Line 3002

	// Apply the x8D function with these 2 values.

	// The x8D function takes in two values, then generates an array based on something. Then it joins the array together and returns that as the value.

	x8DValue := func(LGD int, SGD string) (xGDString string) { // Line 4829
		// This function is imitating the x9D function on line 3241.
		x9D := func(charCodeAtValue int32, hf int32, hk int32, xGDCharCode int32) (x9DInt32 int32) {
			if charCodeAtValue > hf {
				if charCodeAtValue <= hk {
					// Update the value of charCodeAtValue
					charCodeAtValue += xGDCharCode % (hk - hf)
					if charCodeAtValue > hk {
						charCodeAtValue = (charCodeAtValue - hk) + hf
					}
				}
			}
			// fmt.Println("x9D value: " + strconv.Itoa(int(charCodeAtValue)))
			return charCodeAtValue
		}

		var xGD []string

		// for loop over the length of LGD
		for i := 0; i < len(strconv.Itoa(LGD)); i++ {
			charCodeAtValue := rune(strconv.Itoa(LGD)[i])
			fmt.Println("charCodeAtValue: " + strconv.Itoa(int(charCodeAtValue)))
			charAtValue := string(strconv.Itoa(LGD)[i])
			fmt.Println("charAtValue: " + charAtValue)

			charCodeAtValue = x9D(charCodeAtValue, hf, hk, rune(SGD[i])) // x9D(fGD, hf, hk, SGD[DD.kd.apply(null, [Kv, bL])](pJ(tGD, sGD)))
			charAtValue = string(charCodeAtValue)
			xGD = append(xGD, charAtValue)
		}

		finalString := strings.Join(xGD, "")

		return finalString
	}(s8D, T8D)

	// Lastly, we just concatenate the strings together.
	f8DString := T8D + x8DValue

	return f8DString
}

// This function is imitating the (DD.IV(cL, Nz, wW)) expression on line 3830.
// func gen_C8D() (C8D string) {
// }

// Imitating the tAD generation on Line 2986.
func gen_tAD() (tadArray []string) {
	//[DD.gA(vI(vI(bJ)), fh, Qf, lE, TJ), vAD, DD.UA(fq, fh, mq, lE, Wq), n8D, DD.Ip(zh, FE), fdD, DD.wp(Gk, UT), wAD, DD.Qp(Bk, W7), zdD, DD.Mp.apply(null, [AT, jE, vI([])]), WdD, DD.cp(gE, C0), Z8D, DD.Np(Wh, mv, hs), QAD, DD.Hn(ch, Z0), l8D, DD.Dn(ck, Ef, Xx), kdD, DD.Bn(Gz, Tx, MO), X8D, DD.YA.apply(null, [NE, WK, BE, lE, UE, fh]), NAD, DD.An(k7, Cf, tq), VRD, DD.dn(Is, l4), cAD, DD.Rn(Nv, LJ, vI(PE)), j8D, DD.Vn.apply(null, [ws, lE, Hk]), TRD, DD.bn(v7, Kf), CRD, DD.FA(Tf, BE, hf, lE, vf, fh), hRD, DD.GA(nv, fh, gx, lE, Zv), IRD, DD.Pn.apply(null, [zq, Qs]), g8D, DD.pn(K7, KW), NRD, DD.nn.apply(null, [Wm, BK, Cf]), A8D, DD.Zn(Ms, Nq, Az), P8D, DD.ln(BK, C7), U8D, DD.Xn.call(null, bv, cs), BdD[DD.x6.apply(null, [z7, vI(vI(bJ)), TJ])], DD.jn.call(null, HE, Ns), XRD, DD.gn(RC, nE, vf), wdD, DD.Un.call(null, gT, bJ), Y8D]

	tADSlice := []string{
		"-100",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36,uaend,12147,20030107,en-US,Gecko,5,0,0,0,415091,0,1920,1040,1920,1080,533,969,1920,,cpen:0,i1:0,dm:0,cwen:0,non:1,opc:0,fc:0,sc:0,wrc:1,isc:0,vib:1,bat:1,x11:0,x12:1,8103,0.18373239291,843518179527.5,0,loc:", // vAD - line 2885
		"-105",
		"-1,-1,0,0,-1,-1,0;", // n8D
		"-108",
		"", // fdD
		"-101",
		"do_en,dm_en,t_en", // wAD
		"-110",
		"0,1,286067,592,463;1,1,286068,379,447;2,1,286069,377,447;3,1,286069,375,447;4,1,286073,371,446;5,1,286074,369,446;6,1,286075,367,446;7,1,286076,365,445;8,1,286076,363,445;9,1,286077,361,445;10,1,286079,358,444;11,1,286079,357,444;12,1,286080,355,443;13,1,286082,353,443;", // zdD - Line 2912 - Varies
		"-117",
		"", // WdD
		"-109",
		"0,286081,-1,-1,-1,-1,-1,-1,-1,-1,-1;", // Z8D - Line 4619 - Varies
		"-102",
		"-1,-1,0,0,-1,-1,0;", // QAD - Line 2890
		"-111",
		"0,286082,-1,-1,-1;", // l8D - unsure - varies
		"-114",
		"", // kdD - Line 2914
		"-103",
		"Varies", // X8D - unsure - varies
		"-106",
		"10,2", // NAD - Line 2892 - varies
		"-115",
		"1,4016774,32,286082,286081,0,4588905,23468869,0,1687036359055,7,18047,0,14,3007,0,0,23468870,4577207,0,D2720068FFC95AC1F377B4C8A13B9B86~-1~YAAQlcBMaMy0aciIAQAAgDE5ywqa+lLRtVioZFwbO6juXUsxBM1gvJEzEvLjeN31wqAwTHRf8bZe6BZilenuzWYpzFjiiEwlvfB8nQzT+6HLAuZCpRHIoXT90hGqraWPC6m5wt52qTFJMIAa9MSP/3G8XCUPIAkS4zBZVbA/EUmOSD9ws4Us/FCMTjyq+R5lyM1N4wN4n1viYSAkb9sGq5y4EDfPnNmI+ASbJeuYlNGoxI1P3JdiEnhufCjmXSxL9bWstpAThsw5THXGcNJFjS8Xudz+CY3HjZtGpUYlxZmeGYeylGjHtPo8gV3nyqoMm+gr+ygSwb13zAMiEAW5z0eckf856J3im/FBXLi7UlQP5rW1srvXx4CLmSKl3yW4JS3chanlLVFeK6+ObDnrs6jJ/cMIJWfcmd24qlHXS8ULzfzMw/ia2zfhqpg/qsD43jgdeUk6lkk8Csn/vZJLEztuadDyaT7U3U50smnfXTVjWYC6W81kGqJWa9hPeLjvbhaiz9domIytpj/RCLU=~-1~-1~1687040203,49543,939,-1528843666,30261693,PiZtE,53209,48,0,-1,0", // VRD - Line 2933 - varies
		"-112",
		"https://www.mrporter.com/en-us/", // cAD - Line 2891 - URL of Request
		"-119",
		"-1", // j8D - Line 4637
		"-122",
		"0,0,0,0,1,0,0", // TRD - Line 2952
		"-123",
		"", // CRD - Line 2968
		"-124",
		"", // hRD - Line 2969
		"-126",
		"", // IRD - Line 2970
		"-127",
		"11321144241322243122", // g8D - Unsure - ?
		"-128",
		",1,e79ca373a5d6562e4be163953f8f9a57367140e11d550d1767780b9dd547f5bf", // NRD - Line 2983 - Varies
		"-131",
		"4294705152,131011631,123954963,50,22,1,1,1,0,1", // A8D - Line 2984
		"-132",
		"-1", // P8D - Line 2985
		"-133",
		"", // U8D - Unsure
		"-70",
		"420217769;-1;dis;,7;true;true;true;240;true;24;24;true;false;-1", //BdD.fpValStr - unsure
		"-80",
		"4961", // XRD - Line 2934
		"-90",
		"1,4588905,3,1529635,5,917781|4588905", // wdD - Line 2917
		"-116",
		"0", // Y8D - Line 4632
		"-129",
		"33c6f54e72ed2196a148f981eb3a5bd51bb6639069e41454a91411e89e70f64e,1,f7162ad09dac1001fcfab82f5b119a215430305f729acbf3753a1e8aeca78553,Google Inc. (NVIDIA),ANGLE (NVIDIA, NVIDIA GeForce RTX 2060 Direct3D11 vs_5_0 ps_5_0, D3D11),fab37783e5c3bc89f876774bac442ad3e613f7dc5b769b19a8d4d8c6b80ef077,29", // LRD - Line 2951
	}

	return tADSlice
}

func main() {
	myVars := initVariables()
	myVars = step1(myVars)
	fmt.Println("step 1 r8D: " + myVars.r8D)
	myVars = step2(myVars)
	fmt.Println("step 2 r8D: " + myVars.r8D)
	myVars = step3(myVars)
	fmt.Println("step 3 r8D: " + myVars.r8D)
	myVars = step4(myVars)
	fmt.Println("step 4 r8D: " + myVars.r8D)
	myVars = step5(myVars)
	fmt.Println("step 5 r8D: " + myVars.r8D)

	err := ioutil.WriteFile("test.txt", []byte(myVars.r8D), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
