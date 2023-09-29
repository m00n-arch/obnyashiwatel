package rules

// SwapLetters В КАЧЕСТВЕ ПРОГОННОГО ТЕКСТА БУДЕТ ИСПОЛЬЗОВАТЬСЯ ПРЕДЛОЖЕНИЕ "Я ЛЮБЛЮ КЛУБНИКУ". ОЖИДАЕМЫЙ ВЫВОД - "Я ВЮБВЮ КВУБНИКУ"
func SwapLetters(inputText string) string {
	var changes = map[rune]rune{
		'л': 'в',
		'р': 'л',
		'в': 'ф',
		'с': 'ф',
		'ш': 'ф',
		'щ': 'ф',
	}

	r := []rune(inputText)
	for i := range r {
		for k, v := range changes {
			if r[i] == k {
				r[i] = v
			}
			if r[i] == toUpperCase(k) {
				r[i] = toUpperCase(v)
			}
		}
		// switch r[i] {
		// case 'Л':
		// 	r[i] = 'В'
		// case 'л':
		// 	r[i] = 'в'
		// case 'Р':
		// 	r[i] = 'Л'
		// case 'р':
		// 	r[i] = 'л'
		// case 'С':
		// 	r[i] = 'Ф'
		// case 'с':
		// 	r[i] = 'ф'
		// case 'Ш':
		// 	r[i] = 'Ф'
		// case 'ш':
		// 	r[i] = 'ф'
		// case 'Щ':
		// 	r[i] = 'Ф'
		// case 'щ':
		// 	r[i] = 'ф'
		// case 'В':
		// 	r[i] = 'Ф'
		// case 'в':
		// 	r[i] = 'ф'
		// }
	}
	return string(r)
}

const bitsBetweenCases = 32

func toUpperCase(r rune) rune {
	// НЕ КОНВЕРТИРОВАТЬ БОЛЬШИЕ БУКВЫ!!!
	if r >= 'А' && r <= 'Я' {
		return r
	}
	return r - bitsBetweenCases
}
