package datautils

//////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"config"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"hashutils"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////

type DateSymbolStorage struct {
	Date   int
	Symbol string
	Data   string
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func DateSymbolStorageNewBytes(date int, symbol string, data []byte) *DateSymbolStorage {
	d := new(DateSymbolStorage)

	d.Date = date
	d.Symbol = symbol
	d.Data = base64.StdEncoding.EncodeToString(data)

	return d
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func DateSymbolStorageNewString(date int, symbol string, data string) *DateSymbolStorage {
	d := new(DateSymbolStorage)

	d.Date = date
	d.Symbol = symbol
	d.Data = data

	return d
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (d *DateSymbolStorage) JsonDecode(buffer []byte) error {

	err := json.Unmarshal(buffer, d)
	if err != nil {
		return err
	}

	if d.Date == 0 {
		return fmt.Errorf("Date cannot be empty")
	}

	if d.Symbol == "" {
		return fmt.Errorf("Symbol cannot be empty")
	}

	if d.Data == "" {
		return fmt.Errorf("Data cannot be empty")
	}

	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type PeriodSymbolStorage struct {
	From   int
	To     int
	Symbol string

	Data string
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func PeriodSymbolStorageNewBytes(from, to int, symbol string, data []byte) *PeriodSymbolStorage {
	p := new(PeriodSymbolStorage)

	p.From = from
	p.To = to
	p.Symbol = symbol
	p.Data = base64.StdEncoding.EncodeToString(data)

	return p
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func PeriodSymbolStorageNewString(from, to int, symbol string, data string) *PeriodSymbolStorage {
	p := new(PeriodSymbolStorage)

	p.From = from
	p.To = to
	p.Symbol = symbol
	p.Data = data

	return p
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (p *PeriodSymbolStorage) JsonDecode(buffer []byte) error {

	err := json.Unmarshal(buffer, p)
	if err != nil {
		return err
	}

	if p.From == 0 {
		return fmt.Errorf("From cannot be empty")
	}

	if p.To == 0 {
		return fmt.Errorf("To cannot be empty")
	}

	if p.Symbol == "" {
		return fmt.Errorf("Symbol cannot be empty")
	}

	if p.Data == "" {
		return fmt.Errorf("Data cannot be empty")
	}

	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type DateStorage struct {
	Date int
	Data string
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func DateStorageNewBytes(date int, data []byte) *DateStorage {
	d := new(DateStorage)

	d.Date = date
	d.Data = base64.StdEncoding.EncodeToString(data)

	return d
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func DateStorageNewString(date int, data string) *DateStorage {
	d := new(DateStorage)

	d.Date = date
	d.Data = data

	return d
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (d *DateStorage) JsonDecode(buffer []byte) error {

	err := json.Unmarshal(buffer, d)
	if err != nil {
		return err
	}

	if d.Date == 0 {
		return fmt.Errorf("Date cannot be empty")
	}

	if d.Data == "" {
		return fmt.Errorf("Data cannot be empty")
	}

	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type DateIdStorage struct {
	Date int
	Id   string
	Data string
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func DateIdStorageNewBytes(date int, data []byte) *DateIdStorage {
	d := new(DateIdStorage)

	d.Date = date
	d.Id = hashutils.RandBytesB64(config.KeyLengthDefault)
	d.Data = base64.StdEncoding.EncodeToString(data)

	return d
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func DateIdStorageNewString(date int, data string) *DateIdStorage {
	d := new(DateIdStorage)

	d.Date = date
	d.Id = hashutils.RandBytesB64(config.KeyLengthDefault)
	d.Data = data

	return d
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (d *DateIdStorage) JsonDecode(buffer []byte) error {

	err := json.Unmarshal(buffer, d)
	if err != nil {
		return err
	}

	if d.Date == 0 {
		return fmt.Errorf("Date cannot be empty")
	}

	if d.Data == "" {
		return fmt.Errorf("Data cannot be empty")
	}

	d.Id = hashutils.RandBytesB64(config.KeyLengthDefault)

	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type IdStorage struct {
	Id   string
	Data string
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func IdStorageNewBytes(data []byte) *IdStorage {
	d := new(IdStorage)

	d.Id = hashutils.RandBytesB64(config.KeyLengthDefault)
	d.Data = base64.StdEncoding.EncodeToString(data)

	return d
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func IdStorageNewString(data string) *IdStorage {
	d := new(IdStorage)

	d.Id = hashutils.RandBytesB64(config.KeyLengthDefault)
	d.Data = data

	return d
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (d *IdStorage) JsonDecode(buffer []byte) error {

	err := json.Unmarshal(buffer, d)
	if err != nil {
		return err
	}

	if d.Data == "" {
		return fmt.Errorf("Data cannot be empty")
	}

	d.Id = hashutils.RandBytesB64(config.KeyLengthDefault)

	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
