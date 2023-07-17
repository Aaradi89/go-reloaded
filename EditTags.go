package main

import (
	"fmt"
	"regexp"
)

// func fixSpaces(text string) string {
// 	re := regexp.MustCompile(` +`)
// 	newText := re.ReplaceAll([]byte(text), []byte(` `))
// 	return string(newText)
// }

func editTags(text string) string {
	hexPattern := `([^)\s]+) *\( *h *e *x *\)`
	binPattern := `([^)\s]+) *\( *b *i *n *\)`
	upPattern := `([^)\s]+) *\( *u *p *\)`
	lowPattern := `([^)\s]+) *\( *l *o *w *\)`
	capPattern := `([^)\s]+) *\( *c *a *p *\)`
	lowSpecialPattern := `([^)\s]+ *)+(\( *l *o *w *, *(\d+) *\))`
	upSpecialPattern := `([^)\s]+ *)+(\( *u *p *, *(\d+) *\))`
	capSpecialPattern := `([^)\s]+ *)+(\( *c *a *p *, *(\d+) *\))`
	re := regexp.MustCompile(hexPattern + `|` + binPattern + `|` + upPattern + `|` + lowPattern + `|` + capPattern + `|` + lowSpecialPattern + `|` + upSpecialPattern + `|` + capSpecialPattern)
	editText := re.ReplaceAllFunc([]byte(text), func(b []byte) []byte {
		hexRe := regexp.MustCompile(hexPattern)
		binRe := regexp.MustCompile(binPattern)
		upRe := regexp.MustCompile(upPattern)
		lowRe := regexp.MustCompile(lowPattern)
		capRe := regexp.MustCompile(capPattern)
		lowSpecialRe := regexp.MustCompile(lowSpecialPattern)
		upSpecialRe := regexp.MustCompile(upSpecialPattern)
		capSpecialRe := regexp.MustCompile(capSpecialPattern)
		modifiyText := ``
		if hexRe.Match(b) {
			modifiyText = hex(string(b))
		} else if binRe.Match(b) {
			modifiyText = bin(string(b))
		} else if upRe.Match(b) {
			modifiyText = up(string(b))
		} else if lowRe.Match(b) {
			modifiyText = low(string(b))
		} else if capRe.Match(b) {
			modifiyText = cap(string(b))
		} else if lowSpecialRe.Match(b) {
			modifiyText = specialLow(string(b))
		} else if upSpecialRe.Match(b) {
			modifiyText = specialUp(string(b))
		} else if capSpecialRe.Match(b) {
			modifiyText = specialCap(string(b))
		}

		return []byte(modifiyText)
	})
	return string(editText)
}

func checkTags(text string) bool {
	hexPattern := `\( *h *e *x *\)`
	binPattern := `\( *b *i *n *\)`
	upPattern := `\( *u *p *\)`
	lowPattern := `\( *l *o *w *\)`
	capPattern := `\( *c *a *p *\)`
	lowSpecialPattern := `\( *l *o *w *,.*\)`
	upSpecialPattern := `\( *u *p *,.*\)`
	capSpecialPattern := `\( *c *a *p *,.*\)`
	tagErr := true
	re := regexp.MustCompile(hexPattern + `|` + binPattern + `|` + upPattern + `|` + lowPattern + `|` + capPattern + `|` + lowSpecialPattern + `|` + upSpecialPattern + `|` + capSpecialPattern)
	re.ReplaceAllFunc([]byte(text), func(b []byte) []byte {
		hexRe := regexp.MustCompile(hexPattern)
		binRe := regexp.MustCompile(binPattern)
		upRe := regexp.MustCompile(upPattern)
		lowRe := regexp.MustCompile(lowPattern)
		capRe := regexp.MustCompile(capPattern)
		lowSpecialRe := regexp.MustCompile(lowSpecialPattern)
		upSpecialRe := regexp.MustCompile(upSpecialPattern)
		capSpecialRe := regexp.MustCompile(capSpecialPattern)
		if hexRe.Match(b) {
			tagErr = errMsg(b)
		} else if binRe.Match(b) {
			tagErr = errMsg(b)
		} else if upRe.Match(b) {
			tagErr = errMsg(b)
		} else if lowRe.Match(b) {
			tagErr = errMsg(b)
		} else if capRe.Match(b) {
			tagErr = errMsg(b)
		} else if lowSpecialRe.Match(b) {
			tagErr = errMsg(b)
		} else if upSpecialRe.Match(b) {
			tagErr = errMsg(b)
		} else if capSpecialRe.Match(b) {
			tagErr = errMsg(b)
		}
		return []byte(text)
	})
	return tagErr
}

func errMsg(s []byte) bool {
	errTags := string(s)
	fmt.Println("you have a Error in", errTags, "tag(s) please Double Check")
	return false
}
