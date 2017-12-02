package cipherutils

//////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"strings"
	"testing"
	"time"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////

func TestAesCipherConsistency(t *testing.T) {

	//t.SkipNow()
	//t.Parallel()

	setupTest()

	secret := "lalalalalala~~~~~"

	cipherText := AesCipherInstanceEncode(secret)
	plainText, err := AesCipherInstanceDecode(cipherText)
	if err != nil {
		t.Errorf("\nError: %v\n", err)
	}

	if strings.Compare(secret, plainText) != 0 {
		t.Errorf("\nExpect: %v\nGet: %v\n", secret, plainText)
	}

}

//////////////////////////////////////////////////////////////////////////////////////////////////////

func TestAesCipherePreservedOperation(t *testing.T) {

	//t.SkipNow()
	setupTest()

	if len(aesCipherInstanceStore) != 1 {
		t.Errorf("\nExpect: %v\nGet: %v\n", 1, len(aesCipherInstanceStore))
		for k, _ := range aesCipherInstanceStore {
			t.Errorf("%v\n", k)
		}
	}

	for i := 0; i < int(aesCipherKeyMaxPreserved)+overRun; i++ {

		secret := "lalalalalala~~~~~"

		cipherText := AesCipherInstanceEncode(secret)

		if len(aesCipherInstanceStore) != 1+i {
			if (1 + i) > int(aesCipherKeyMaxPreserved) {

				if len(aesCipherInstanceStore) != int(aesCipherKeyMaxPreserved) {
					t.Errorf("\nExpect: %v\nGet: %v\n", aesCipherKeyMaxPreserved, len(aesCipherInstanceStore))
					for k, _ := range aesCipherInstanceStore {
						t.Errorf("%v\n", k)
					}
				}

			} else {

				t.Errorf("\nExpect: %v\nGet: %v\n", 1+i, len(aesCipherInstanceStore))
				for k, _ := range aesCipherInstanceStore {
					t.Errorf("%v\n", k)
				}

			}
		}

		time.Sleep(time.Duration(int(aesCipherKeyMaxAge)+3) * time.Second)

		plainText, err := AesCipherInstanceDecode(cipherText)
		if err != nil {
			t.Errorf("\nError: %v\n", err)
		}

		if strings.Compare(secret, plainText) != 0 {
			t.Errorf("\nExpect: %v\nGet: %v\n", secret, plainText)
		}

	}

}

//////////////////////////////////////////////////////////////////////////////////////////////////////

func TestAesCipherInstanceStoreCleanOld(t *testing.T) {

	setupTest()

	if len(aesCipherInstanceStore) != 1 {
		t.Errorf("\nExpect: %v\nGet: %v\n", 1, len(aesCipherInstanceStore))
	}

	currentKey := make([]time.Time, 0)
	for i := 0; i < int(aesCipherKeyMaxPreserved)+overRun; i++ {
		key := time.Now().UTC().AddDate(i, 0, 0)
		currentKey = append(currentKey, key)

		aesCipherInstanceStore[key] = nil

		aesCipherInstanceStoreClean()

		if len(aesCipherInstanceStore) <= int(aesCipherKeyMaxPreserved) {
			if i+2 <= int(aesCipherKeyMaxPreserved) && len(aesCipherInstanceStore) != i+2 {
				t.Errorf("\nExpect: %v\nGet: %v\n", i+2, len(aesCipherInstanceStore))
				t.Errorf("\n%v\n", aesCipherInstanceStore)
			}
		} else {
			if len(aesCipherInstanceStore) != int(aesCipherKeyMaxPreserved) {
				t.Errorf("\nExpect: %v\nGet: %v\n", int(aesCipherKeyMaxPreserved), len(aesCipherInstanceStore))
				t.Errorf("\n%v\n", aesCipherInstanceStore)
			}

			for j := 0; j < len(aesCipherInstanceStore)-int(aesCipherKeyMaxPreserved); j++ {
				_, ok := aesCipherInstanceStore[currentKey[j]]
				if ok {
					t.Errorf("\n%v\n", len(aesCipherInstanceStore))
					t.Errorf("\n%v\n", currentKey)
					t.Errorf("\n%v\n", aesCipherInstanceStore)
				}
			}
		}
	}

}

//////////////////////////////////////////////////////////////////////////////////////////////////////

func TestAesCiphereOldKeyOperation(t *testing.T) {

	//t.SkipNow()
	//t.Parallel()
	setupTest()

	secret := "lalalalalala~~~~~"

	cipherText := make([]string, int(aesCipherKeyMaxPreserved)+overRun)

	for i := 0; i < int(aesCipherKeyMaxPreserved)+overRun; i++ {
		cipherText[i] = AesCipherInstanceEncode(secret)
		time.Sleep(time.Duration(int(aesCipherKeyMaxAge)+3) * aesCipherKeyMaxAgeUnit)
	}

	for i := 0; i < int(aesCipherKeyMaxPreserved)+overRun; i++ {
		if i < overRun {
			plainText, err := AesCipherInstanceDecode(cipherText[i])
			if err == nil {
				t.Errorf("\nshould not decoded correctly: %v\n", string(plainText))
				t.Errorf("\nloop index: %v\n", i)
			}

			continue
		}

		plainText, err := AesCipherInstanceDecode(cipherText[i])
		if err != nil {
			t.Errorf("\nError: %v\n", err)
			t.Errorf("\nloop index: %v\n", i)
		}

		if strings.Compare(secret, plainText) != 0 {
			t.Errorf("\nExpect: %v\nGet: %v\n", secret, plainText)
			t.Errorf("\nloop index: %v\n", i)
		}
	}

}

//////////////////////////////////////////////////////////////////////////////////////////////////////
