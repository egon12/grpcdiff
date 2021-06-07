package grpcdiff

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/sergi/go-diff/diffmatchpatch"
)

type Report struct {
	input string
	resA  []byte
	resB  []byte
	durA  time.Duration
	durB  time.Duration
	diffs []diffmatchpatch.Diff
	dmp   *diffmatchpatch.DiffMatchPatch
}

func NewReport(input string, resA, resB []byte, durA, durB time.Duration) *Report {
	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(string(resA), string(resB), true)

	return &Report{
		input: input,
		resA:  resA,
		resB:  resB,
		durA:  durA,
		durB:  durB,
		diffs: diffs,
		dmp:   dmp,
	}

}

func (r *Report) Report() string {
	if r.HasDiff() {
		filenameA, filenameB := r.WriteResToFile()
		return fmt.Sprintf("input %s has diff that you can see at %s and %s", r.input, filenameA, filenameB)
	}
	return fmt.Sprintf("input %s: %s", r.input, r.GetDurationAnalysis())
}

func (r *Report) HasDiff() bool {
	for _, diff := range r.diffs {
		if diff.Type != diffmatchpatch.DiffEqual {
			return true
		}
	}
	return false
}

func (r *Report) GetFormattedDiff() string {
	return r.dmp.DiffPrettyText(r.diffs)
}

func (r *Report) GetDurationAnalysis() string {
	var longger = "A"
	var durLong time.Duration
	var durShort time.Duration

	if r.durA > r.durB {
		longger = "A"
		durLong = r.durA
		durShort = r.durB
	} else if r.durB > r.durA {
		longger = "B"
		durLong = r.durB
		durShort = r.durA
	} else {
		return "duration is exactly equals"
	}

	comparison := float64(durLong)/float64(durShort) - 1.0
	if comparison > 0.1 {
		return fmt.Sprintf("%s takes more %.02fx times (durA: %v, durB: %v)", longger, comparison, r.durA, r.durB)
	}

	return "duration is quite same"
}

func (r *Report) WriteResToFile() (filenameA, filenameB string) {
	filename := r.generateFileName()
	filenameA = filename + "-A.json"
	filenameB = filename + "-B.json"
	r.writeFile(filenameA, r.resA)
	r.writeFile(filenameB, r.resB)
	return
}

func (r *Report) writeFile(filename string, content []byte) {
	err := ioutil.WriteFile(filename, content, os.ModePerm)
	if err != nil {
		panic(err)
	}
}

func (r *Report) generateFileName() string {
	randBytes := make([]byte, 16)
	rand.Read(randBytes)
	return hex.EncodeToString(randBytes)
}
