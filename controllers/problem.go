package controllers

import (
	"github.com/astaxie/beego"
)

type ProblemController struct {
	beego.Controller
}

func (c *ProblemController) URLMapping() {
	c.Mapping("Count", c.Count)
	c.Mapping("ViewSimple", c.ViewSimple)
	c.Mapping("View", c.View)
}

// @router /problem/count [post]
func (c *ProblemController) Count() {
	c.Data["json"] = 209
	c.ServeJson()
}

// @router /problem/view/simple/:page:int [post]
func (c *ProblemController) ViewSimple() {
	beego.Info(c.Ctx.Input.Param(":page"))
	c.Data["json"] = []map[string]interface{}{
		{"problemId": 28, "title": "Sum of 4 different primes", "distinctAcceptedUser": 73, "tried": false, "isAccepted": false, "totalSubmission": 161, "acceptedSubmission": 79, "localeAvailableEN": true, "localeAvailableID": true, "difficulty": 1.2, "problemCode": nil, "isBookmarked": false}, {"problemId": 111, "title": "Nosy Duck", "distinctAcceptedUser": 65, "tried": false, "isAccepted": false, "totalSubmission": 116, "acceptedSubmission": 67, "localeAvailableEN": true, "localeAvailableID": true, "difficulty": 1.5, "problemCode": nil, "isBookmarked": false}, {"problemId": 200, "title": "George and a String", "distinctAcceptedUser": 64, "tried": false, "isAccepted": false, "totalSubmission": 101, "acceptedSubmission": 65, "localeAvailableEN": true, "localeAvailableID": true, "difficulty": 1.5, "problemCode": nil, "isBookmarked": false}, {"problemId": 190, "title": "Role Generator", "distinctAcceptedUser": 63, "tried": false, "isAccepted": false, "totalSubmission": 166, "acceptedSubmission": 68, "localeAvailableEN": true, "localeAvailableID": true, "difficulty": 1.3, "problemCode": nil, "isBookmarked": false}, {"problemId": 30, "title": "Not an Email", "distinctAcceptedUser": 62, "tried": false, "isAccepted": false, "totalSubmission": 187, "acceptedSubmission": 62, "localeAvailableEN": true, "localeAvailableID": true, "difficulty": 1.5, "problemCode": nil, "isBookmarked": false}, {"problemId": 102, "title": "Dot Inside Circle", "distinctAcceptedUser": 57, "tried": false, "isAccepted": false, "totalSubmission": 106, "acceptedSubmission": 57, "localeAvailableEN": true, "localeAvailableID": true, "difficulty": 1.4, "problemCode": nil, "isBookmarked": false}, {"problemId": 83, "title": "Punishment for Panda", "distinctAcceptedUser": 54, "tried": false, "isAccepted": false, "totalSubmission": 400, "acceptedSubmission": 61, "localeAvailableEN": true, "localeAvailableID": true, "difficulty": 1.4, "problemCode": nil, "isBookmarked": false}, {"problemId": 129, "title": "Deduksi Peringkat", "distinctAcceptedUser": 54, "tried": false, "isAccepted": false, "totalSubmission": 99, "acceptedSubmission": 55, "localeAvailableEN": true, "localeAvailableID": true, "difficulty": 1.7, "problemCode": nil, "isBookmarked": false}, {"problemId": 174, "title": "Derp Family and Chair", "distinctAcceptedUser": 53, "tried": false, "isAccepted": false, "totalSubmission": 115, "acceptedSubmission": 55, "localeAvailableEN": true, "localeAvailableID": true, "difficulty": 1.5, "problemCode": nil, "isBookmarked": false}, {"problemId": 162, "title": "The Duck and Its Name", "distinctAcceptedUser": 49, "tried": false, "isAccepted": false, "totalSubmission": 96, "acceptedSubmission": 49, "localeAvailableEN": true, "localeAvailableID": true, "difficulty": 1.9, "problemCode": nil, "isBookmarked": false}, {"problemId": 187, "title": "Hackbook Password", "distinctAcceptedUser": 48, "tried": false, "isAccepted": false, "totalSubmission": 122, "acceptedSubmission": 54, "localeAvailableEN": true, "localeAvailableID": false, "difficulty": 2.0, "problemCode": nil, "isBookmarked": false}, {"problemId": 169, "title": "Play, play, play!", "distinctAcceptedUser": 46, "tried": false, "isAccepted": false, "totalSubmission": 90, "acceptedSubmission": 47, "localeAvailableEN": true, "localeAvailableID": true, "difficulty": 1.7, "problemCode": nil, "isBookmarked": false}, {"problemId": 151, "title": "Elimpiade", "distinctAcceptedUser": 44, "tried": false, "isAccepted": false, "totalSubmission": 67, "acceptedSubmission": 44, "localeAvailableEN": true, "localeAvailableID": true, "difficulty": 1.7, "problemCode": nil, "isBookmarked": false}, {"problemId": 54, "title": "Bujur Sangkar Ajaib", "distinctAcceptedUser": 43, "tried": false, "isAccepted": false, "totalSubmission": 52, "acceptedSubmission": 45, "localeAvailableEN": true, "localeAvailableID": true, "difficulty": 2.5, "problemCode": nil, "isBookmarked": false}, {"problemId": 176, "title": "Another Face Swap", "distinctAcceptedUser": 37, "tried": false, "isAccepted": false, "totalSubmission": 102, "acceptedSubmission": 38, "localeAvailableEN": true, "localeAvailableID": true, "difficulty": 1.5, "problemCode": nil, "isBookmarked": false}, {"problemId": 37, "title": "Water Bottle Again", "distinctAcceptedUser": 35, "tried": false, "isAccepted": false, "totalSubmission": 127, "acceptedSubmission": 37, "localeAvailableEN": true, "localeAvailableID": true, "difficulty": 1.5, "problemCode": nil, "isBookmarked": false}, {"problemId": 148, "title": "Jalan - Jalan", "distinctAcceptedUser": 35, "tried": false, "isAccepted": false, "totalSubmission": 59, "acceptedSubmission": 35, "localeAvailableEN": true, "localeAvailableID": true, "difficulty": 1.8, "problemCode": nil, "isBookmarked": false}, {"problemId": 130, "title": "Kacamata Anti-Akik", "distinctAcceptedUser": 33, "tried": false, "isAccepted": false, "totalSubmission": 50, "acceptedSubmission": 34, "localeAvailableEN": true, "localeAvailableID": true, "difficulty": 1.7, "problemCode": nil, "isBookmarked": false}, {"problemId": 165, "title": "The Goat and The New Power", "distinctAcceptedUser": 33, "tried": false, "isAccepted": false, "totalSubmission": 42, "acceptedSubmission": 34, "localeAvailableEN": true, "localeAvailableID": true, "difficulty": 1.5, "problemCode": nil, "isBookmarked": false}, {"problemId": 195, "title": "Natural Selection", "distinctAcceptedUser": 33, "tried": false, "isAccepted": false, "totalSubmission": 70, "acceptedSubmission": 36, "localeAvailableEN": true, "localeAvailableID": true, "difficulty": 2.0, "problemCode": nil, "isBookmarked": false}, {"problemId": 132, "title": "ROKET-1", "distinctAcceptedUser": 30, "tried": false, "isAccepted": false, "totalSubmission": 113, "acceptedSubmission": 37, "localeAvailableEN": true, "localeAvailableID": true, "difficulty": 1.8, "problemCode": nil, "isBookmarked": false}, {"problemId": 182, "title": "Rapid Transit", "distinctAcceptedUser": 29, "tried": false, "isAccepted": false, "totalSubmission": 70, "acceptedSubmission": 33, "localeAvailableEN": true, "localeAvailableID": false, "difficulty": 2.2, "problemCode": nil, "isBookmarked": false}, {"problemId": 161, "title": "The Duck and The Man's Keyboard", "distinctAcceptedUser": 28, "tried": false, "isAccepted": false, "totalSubmission": 43, "acceptedSubmission": 28, "localeAvailableEN": true, "localeAvailableID": true, "difficulty": 1.7, "problemCode": nil, "isBookmarked": false}, {"problemId": 188, "title": "Counting Square Stars", "distinctAcceptedUser": 28, "tried": false, "isAccepted": false, "totalSubmission": 84, "acceptedSubmission": 30, "localeAvailableEN": true, "localeAvailableID": false, "difficulty": 2.0, "problemCode": nil, "isBookmarked": false}, {"problemId": 5, "title": "Best Assistant", "distinctAcceptedUser": 27, "tried": false, "isAccepted": false, "totalSubmission": 67, "acceptedSubmission": 28, "localeAvailableEN": true, "localeAvailableID": true, "difficulty": 1.5, "problemCode": nil, "isBookmarked": false}, {"problemId": 131, "title": "Roda Gigi", "distinctAcceptedUser": 27, "tried": false, "isAccepted": false, "totalSubmission": 58, "acceptedSubmission": 28, "localeAvailableEN": true, "localeAvailableID": true, "difficulty": 1.7, "problemCode": nil, "isBookmarked": false}, {"problemId": 180, "title": "Ticket Itinerary", "distinctAcceptedUser": 27, "tried": false, "isAccepted": false, "totalSubmission": 93, "acceptedSubmission": 29, "localeAvailableEN": true, "localeAvailableID": false, "difficulty": 1.5, "problemCode": nil, "isBookmarked": false}, {"problemId": 191, "title": "Timid Person", "distinctAcceptedUser": 27, "tried": false, "isAccepted": false, "totalSubmission": 37, "acceptedSubmission": 27, "localeAvailableEN": true, "localeAvailableID": true, "difficulty": 2.0, "problemCode": nil, "isBookmarked": false}, {"problemId": 65, "title": "Jumlah Pangkat", "distinctAcceptedUser": 26, "tried": false, "isAccepted": false, "totalSubmission": 155, "acceptedSubmission": 37, "localeAvailableEN": true, "localeAvailableID": true, "difficulty": 3.2, "problemCode": nil, "isBookmarked": false}, {"problemId": 69, "title": "Potong Gaji", "distinctAcceptedUser": 25, "tried": false, "isAccepted": false, "totalSubmission": 160, "acceptedSubmission": 40, "localeAvailableEN": true, "localeAvailableID": true, "difficulty": 1.8, "problemCode": nil, "isBookmarked": false},
	}
	c.ServeJson()
}

// @router /problem/view/:pid:int [post]
func (c *ProblemController) View() {
	c.Data["json"] = map[string]interface{}{
		"problemId":   161,
		"problemCode": nil,
		"title":       `The Duck and The Man's Keyboard`,
		"problemDescriptionEN": `<center>
		<i>This problem is inspired by this inspiring video<br><br>

		<iframe width="427" height="240" src="https://www.youtube.com/embed/MtN1YnoL46Q" frameborder="0" allowfullscreen></iframe><br><br>

		"Hey, got any grapes?"<br>
		"That's it, if you don't stay away duck, i'll glue you to a tree and leave you there all day, stuck!"<br><br>

		</i>

		</center>

		<p>The man running a lemonade stand which is often annoyed by the duck has lost his patience and wants to note the duck's name in his phone. But the man finds it difficult to type the duck's name because his phone is not a high-end smartphone.</p>

		<p>Living in the era before smartphone is not easy, but that does not mean the phone in earlier times doesn't have any uniqueness. One of the uniqueness is the 3x4 keyboard that can be seen below.</p>

		<center><img src="/problem/image/1321a6b1-e0d0-411a-a33f-f98f02527ed1"></center>

		<p>As you can see, if we want to type a letter 'K', we have to push the button '5' 2 times, in other words, you have to push '55'. If you want to type a letter located on the same button with previous letter, you have to wait for a while before you push that button. For example, 'VV' you have to push '888 888'. So if you want to type 'MADYA', you have to push: '6 2 3 999 2'.</p>

		<p>You are given a word with maximum length of 1000, output the push sequence if you want to type it with 3x4 keyboard.</p>`,
		"problemDescriptionID": `<center>
		<i>Soal ini terinspirasi oleh video inspiratif berikut<br><br>

		<iframe width="427" height="240" src="https://www.youtube.com/embed/MtN1YnoL46Q" frameborder="0" allowfullscreen></iframe><br><br>

		"Hey, got any grapes?"<br>
		"That's it, if you don't stay away duck, i'll glue you to a tree and leave you there all day, stuck!"<br><br>

		</i>

		</center>

		<p>Penjual lemon yang sering diganggu si bebek akhirnya kesal dan ingin mencatat nama si bebek di handphonenya. Namun si penjual lemon tersebut merasa kesulitan ketika ingin mengetik nama si bebek karena handphonya bukanlah smartphone yang canggih.</p>

		<p>Hidup di jaman sebelum smartphone memang tidak begitu mudah namun bukan berarti handphone dulu tidak memiliki keunikan. Salah satunya adalah 3x4 keyboard yang dapat kalian lihat dibawah ini.</p>

		<center><img src="/problem/image/1321a6b1-e0d0-411a-a33f-f98f02527ed1"></center>

		<p>Seperti kalian lihat, jika ingin mengetik huruf 'K' kita harus menekan tombol 5 sebanyak 2 kali atau kalian harus menekan '55'. Jika ada huruf yang sama sebelumnya, kalian harus menunggu beberapa saat sebelum menekan kembali angka tersebut. Sebagai contoh, 'VV' kalian harus menekan tombol sebagai berikut '888 888'. Jadi jika kita ingin mengetik 'MADYA', maka kalian harus menekan tombol dengan urutan sebagai berikut : '6 2 3 999 2'.</p>

		<p>Kalian diberikan sebuah kata dengan panjang maksimal 1000, cetak urutan menekan tombol jika kalian mengetik menggunakan keyboard 3x4.</p>`,
		"memoryLimit": 32,
		"timeLimit":   2,
		"inputDescriptionEN": `<p>First line consists of an integer T (T &le; 100) represents the number of test cases.</p>
		<p>Each test case consists of a word S (1 &le; |S| &le; 1000) contains only letter 'A'..'Z'.</p>`,
		"inputDescriptionID": `<p>Baris pertama berisi sebuah bilangan T (T &le; 100) yang menyatakan banyaknya kasus yang harus ditangani.</p>
		<p>Untuk setiap kasus, berisi sebuah kata S (1 &le; |S| &le; 1000) yang hanya terdiri dari huruf 'A' sampai 'Z'.</p>`,
		"outputDescriptionEN": `<p>For each test case, output "Case #Y: Z" (without quotes) where Y is the test case number starts from 1, and Z is the push sequence.</p>`,
		"outputDescriptionID": `<p>Untuk setiap kasus, output "Case #Y: Z" (tanpa tanda kutip) di mana Y adalah nomor kasus dimulai dari 1, dan Z urutan penekanan tombol.</p>`,
		"sampleTestCaseCount": 1,
		"difficulty":          1.7,
		"sampleInput": []string{
			`2
VV
MADYA`,
			`2
VV
MADYA`,
		},
		"sampleOutput": []string{
			`Case #1: 888 888
Case #2: 6 2 3 999 2`,
			`Case #1: 888 888
Case #2: 6 2 3 999 2`,
		},
		"problemSetterUserId":      1,
		"hidden":                   false,
		"noteEN":                   ``,
		"noteID":                   ``,
		"defaultCheckerExactMatch": true,
		"hint":                          `implementation`,
		"totalSubmission":               43,
		"acceptedSubmission":            28,
		"outputLimitExceededSubmission": 0,
		"timeLimitExceededSubmission":   1,
		"wrongAnswerSubmission":         14,
		"compileErrorSubmission":        0,
		"memoryLimitExceededSubmission": 0,
		"runTimeErrorSubmission":        0,
		"submissionErrorSubmission":     0,
		"author":                        `<a onmouseover="addHoverCard_username(this, &quot;madya121&quot;)" onmouseleave="removeHoverCard()" href="/user/view/madya121" target="_blank" class="username" style="color:gray">madya121</a>`,
		"tried":                         false,
		"isAccepted":                    false,
	}
	c.ServeJson()
}
