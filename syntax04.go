package main

import (
	"fmt"
	"math/rand" // 랜덤 함수
	"time"      // seed 생성용 패키지
)

// 난수 추출된 수의 소수 판정 프로그램v0.4
// 소수 : 1과 자기자신외에는 나누어 떨어지지 않는 수 (0과 1은 제외)
func main() {
	seed := time.Now().Unix()
	rand.Seed(seed)

	isPrime := true
	number := rand.Intn(150) + 2
	fmt.Println("임의로 추출된 수 : ", number)

	for i := 2; i < number; i++ {
		if number%i == 0 {
			isPrime = false //한번이라도 i로 나눴을떄 나눠진다면 소수가 아니다.
		}
		fmt.Print(i, " ")
	}
	if isPrime { // true 일떄 실행 (비교 연산자 제거)
		fmt.Println(number, "는(은) 소수입니다!")
	} else {
		fmt.Println(number, "는(은) 소수가 아닙니다~")
	}
}
