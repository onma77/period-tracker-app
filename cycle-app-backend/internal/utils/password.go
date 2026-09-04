package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func CheckPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword),
		[]byte(password),
	)
}

// package utils

// import "golang.org/x/crypto/bcrypt"

// // HashPassword converts a plain password into a secure hash.
// func HashPassword(password string) (string, error) {
// 	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
// 	if err != nil {
// 		return "", err
// 	}

// 	return string(bytes), nil
// }

// // CheckPassword compares a hashed password with a plain password.
// func CheckPassword(hashedPassword, password string) error {
// 	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
// }

// // package utils

// // import "golang.org/x/crypto/bcrypt"

// // func HashPassword(password string) (string, error) {
// // 	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
// // 	return string(bytes), err
// // }

// // func CheckPassword(hash string, password string) error {
// // 	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
// // }
