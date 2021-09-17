// ● "a4bc2d5e" => "aaaabccddddde"
// ● "abcd" => "abcd"
// ● "45" => "" (некорректная строка)
// ● "" => ""
// Дополнительное задание: поддержка escape -
// последовательностей
// ● qwe\4\5 => qwe45 (*)
// ● qwe\45 => qwe44444 (*)
// ● qwe\\5 => qwe\\\\\ (*)
package main

import "testing"

func TestUnpack(t *testing.T){

    tMap := map[string]string{
        "a4bc2d5e": "aaaabccddddde",
        "abcd": "abcd",
        `45`: ``,
        ``:``,
        `qwe\4\5`: `qwe45`,
        `qwe\45`: `qwe44444`,
        `qwe\\5`: `qwe\\\\\`,
    }

    for key, want := range tMap {
        got, err := Unpack(key)
        if err != nil {
            t.Errorf("key=%v, want=%v, got=%v, error, %v", key, want, got, err)
        }
        if got !=  want {
            t.Errorf("got %q, wanted %q", got, want)
        }
    }

}
    
