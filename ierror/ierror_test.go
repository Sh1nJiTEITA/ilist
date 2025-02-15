package ierror

// import (
// 	"errors"
// 	"fmt"
// 	"testing"
// )
//
// var ErrUserNotFound = errors.New("user not found")
// var ErrUserFound = errors.New("user found")
//
// func UserNotFoundWithUsername(username string) error {
// 	return fmt.Errorf("no user with username %q: %w", username, ErrUserNotFound)
// }
//
// func TestErrorWrapping(t *testing.T) {
// 	// Эмуляция ошибки, когда не найден пользователь с определённым ником
// 	err := UserNotFoundWithUsername("john")
//
// 	// Проверка через errors.Is, ищем базовую ошибку ErrUserNotFound внутри
// 	if !errors.Is(err, ErrUserNotFound) {
// 		t.Fatalf("expected ErrUserNotFound, got: %v", err)
// 	}
// 	if errors.Is(err, ErrUserFound) {
// 		t.Fatalf("expected ErrUserNotFound, got: %v", err)
// 	}
// 	werr := fmt.Errorf("next: %w", err)
// 	if !errors.Is(werr, ErrUserNotFound) {
// 		t.Fatalf("expected ErrUserNotFound, got: %v", err)
// 	}
// 	if errors.Is(werr, ErrUserFound) {
// 		t.Fatalf("expected ErrUserNotFound, got: %v", err)
// 	}
// }

// import (
// 	"errors"
// 	"fmt"
// 	"testing"
// )
//
// func TestIErrorNew(t *testing.T) {
// 	err := New("Some error")
// 	if err.Info != "Some error" {
// 		t.Fatalf("(err.info = %s) != \"Some error\"", err.Info)
// 	}
// 	if err.Additional != "" {
// 		t.Fatalf("(err.addtional = %s) != \"\"", err.Additional)
// 	}
// }
//
// func TestIErrorExtended(t *testing.T) {
// 	err := New("Err 1")
// 	extErr := Extended(&err, "additional info for err 1")
//
// 	if err.Info != "Err 1" {
// 		t.Fatalf("(err.info = %s) != \"Some error\"", err.Info)
// 	}
// 	if err.Additional != "" {
// 		t.Fatalf("(err.additional = %s) != \"\"", err.Additional)
// 	}
//
// 	if extErr.Info != "Err 1" {
// 		t.Fatalf("(extErr.info = %s) != \"Some error\"", extErr.Info)
// 	}
// 	if extErr.Additional != "additional info for err 1" {
// 		t.Fatalf("(extErr.additional = %s) != \"\"", extErr.Additional)
// 	}
// }
//
// func TestNewExtended(t *testing.T) {
// 	err := NewExtended("Err 1", "add info 1")
// 	if err.Info != "Err 1" {
// 		t.Fatalf("(err.info = %s) != \"Some error\"", err.Info)
// 	}
// 	if err.Additional != "add info 1" {
// 		t.Fatalf("(err.additional = %s) != \"\"", err.Additional)
// 	}
// }
//
// var Err1 = New("Some error 1")
//
// func returnDefaultError() error {
// 	return &Err1
// }
//
// func returnExtendedError() error {
// 	return &NewExtended("Some error 2", "2")
// }

// func return() error {
// 	return New("Some error 1")
// }

// func TestNewSprint(t *testing.T) {
// 	err := returnDefaultError()
// 	if errors.Is(err, &Err1) {
// 		t.Fatal("default value is not std error")
// 	}
// 	if fmt.Sprint(err) != "Some error 1" {
// 		t.Fatal("Errors in stderror sprint")
// 	}
//
// }

// func TestErrorWrap(t *testing.T) {
// 	some_err := errors.New("A")
// 	ext_some_err := fmt.Errorf("ext: %w", some_err)
// 	if errors.As(ext_some_err, &some_err) {
// 		t.Fatal("Errors in stderror sprint")
// 	}
// }

// func TestNewExtendedSprint(t *testing.T) {
// 	err := returnExtendedError()
// 	if errors.Is(err, New("Some error 2")) {
// 		t.Fatal("default value is not std error")
// 	}
// 	if fmt.Sprint(err) != "Some error 1" {
// 		t.Fatal("Errors in stderror sprint")
// 	}
// }
