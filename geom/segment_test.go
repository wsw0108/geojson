package geom

import (
	"fmt"
	"strings"
	"testing"
)

func TestSegmentRaycast(t *testing.T) {
	// This is full coverage raycast test. It uses a 7x7 grid where the each
	// point is checked for a total of 49 tests per segment. There are 16
	// segments at 0º,30º,45º,90º,180º in both directions for a total of 784
	// tests.
	segs := []interface{}{
		S(1, 1, 5, 5), "A",
		S(5, 5, 1, 1), "B",
		[7]string{
			"-------",
			"-----*-",
			"++++*--",
			"+++*---",
			"++*----",
			"+*-----",
			"-------",
		},
		S(1, 5, 5, 1), "C",
		S(5, 1, 1, 5), "D",
		[7]string{
			"-------",
			"-*-----",
			"++*----",
			"+++*---",
			"++++*--",
			"+++++*-",
			"-------",
		},
		S(1, 3, 5, 3), "E",
		S(5, 3, 1, 3), "F",
		[7]string{
			"-------",
			"-------",
			"-------",
			"-*****-",
			"-------",
			"-------",
			"-------",
		},
		S(3, 5, 3, 1), "G",
		S(3, 1, 3, 5), "H",
		[7]string{
			"-------",
			"---*---",
			"+++*---",
			"+++*---",
			"+++*---",
			"+++*---",
			"-------",
		},
		S(1, 2, 5, 4), "I",
		S(5, 4, 1, 2), "J",
		[7]string{
			"-------",
			"-------",
			"-----*-",
			"+++*---",
			"+*-----",
			"-------",
			"-------",
		},
		S(1, 4, 5, 2), "K",
		S(5, 2, 1, 4), "L",
		[7]string{
			"-------",
			"-------",
			"-*-----",
			"+++*---",
			"+++++*-",
			"-------",
			"-------",
		},
		S(2, 1, 4, 5), "M",
		S(4, 5, 2, 1), "N",
		[7]string{
			"-------",
			"----*--",
			"++++---",
			"+++*---",
			"+++----",
			"++*----",
			"-------",
		},
		S(2, 5, 4, 1), "O",
		S(4, 1, 2, 5), "P",
		[7]string{
			"-------",
			"--*----",
			"+++----",
			"+++*---",
			"++++---",
			"++++*--",
			"-------",
		},
		S(3, 3, 3, 3), "Q",
		S(3, 3, 3, 3), "R",
		[7]string{
			"-------",
			"-------",
			"-------",
			"---*---",
			"-------",
			"-------",
			"-------",
		},
	}

	var ms string
	for i := 0; i < len(segs); i += 5 {
		segs := []interface{}{
			segs[i], segs[i+1], segs[i+4],
			segs[i+2], segs[i+3], segs[i+4],
		}
		for i := 0; i < len(segs); i += 3 {
			seg := segs[i].(Segment)
			label := segs[i+1].(string)
			grid := segs[i+2].([7]string)
			//
			var ngrid [7]string
			for y, sy := 0, 6; y < 7; y, sy = y+1, sy-1 {
				var nline string
				for x := 0; x < 7; x++ {
					// ch := grid[sy][x]
					pt := Point{float64(x), float64(y)}
					res := seg.Raycast(pt)
					if res.In {
						nline += "+"
					} else if res.On {
						nline += "*"
					} else {
						nline += "-"
					}
				}
				ngrid[sy] = nline
			}
			if grid != ngrid {
				ms += fmt.Sprintf("MISMATCH (%s) SEGMENT: %v\n", label, seg)
				ms += fmt.Sprintf("EXPECTED\n%s\n", strings.Join(grid[:], "\n"))
				ms += fmt.Sprintf("GOT\n%s\n", strings.Join(ngrid[:], "\n"))
			}
		}
	}
	if ms != "" {
		t.Fatalf("\n%s", ms)
	}
}

func TestSegmentContainsPoint(t *testing.T) {
	expect(t, S(0, 0, 1, 1).ContainsPoint(P(0, 0)))
	expect(t, S(0, 0, 1, 1).ContainsPoint(P(0.5, 0.5)))
	expect(t, S(0, 0, 1, 1).ContainsPoint(P(1, 1)))
	expect(t, !S(0, 0, 1, 1).ContainsPoint(P(1.1, 1.1)))
	expect(t, !S(0, 0, 1, 1).ContainsPoint(P(0.5, 0.6)))
	expect(t, !S(0, 0, 1, 1).ContainsPoint(P(-0.1, -0.1)))
}

func TestSegmentCollinearPoint(t *testing.T) {
	expect(t, S(0, 0, 1, 1).CollinearPoint(P(-1, -1)))
	expect(t, S(0, 0, 1, 1).CollinearPoint(P(0.5, 0.5)))
	expect(t, S(0, 0, 1, 1).CollinearPoint(P(2, 2)))
	expect(t, S(1, 1, 0, 0).CollinearPoint(P(-1, -1)))
	expect(t, S(1, 1, 0, 0).CollinearPoint(P(0.5, 0.5)))
	expect(t, S(1, 1, 0, 0).CollinearPoint(P(2, 2)))
	expect(t, S(1, 0, 0, 1).CollinearPoint(P(2, -1)))
	expect(t, S(1, 0, 0, 1).CollinearPoint(P(0.5, 0.5)))
	expect(t, S(1, 0, 0, 1).CollinearPoint(P(-1, 2)))
	expect(t, S(0, 1, 1, 0).CollinearPoint(P(2, -1)))
	expect(t, S(0, 1, 1, 0).CollinearPoint(P(0.5, 0.5)))
	expect(t, S(0, 1, 1, 0).CollinearPoint(P(-1, 2)))
}

func TestSegmentContainsSegment(t *testing.T) {
	expect(t, S(0, 0, 10, 10).ContainsSegment(S(0, 0, 10, 10)))
	expect(t, S(0, 0, 10, 10).ContainsSegment(S(2, 2, 10, 10)))
	expect(t, S(0, 0, 10, 10).ContainsSegment(S(2, 2, 8, 8)))
	expect(t, !S(0, 0, 10, 10).ContainsSegment(S(-1, -1, 8, 8)))
}

func TestSegmentIntersectsSegment(t *testing.T) {
	vals := []interface{}{
		`---------`, `---------`, 1, `---------`, `---------`, 1,
		`---------`, `---------`, 1, `---------`, `---------`, 1,
		`---------`, `---------`, 1, `---------`, `----A----`, 1,
		`----A----`, `----A----`, 1, `----A----`, `----*----`, 1,
		`----*----`, `----*----`, 1, `----*----`, `----B----`, 1,
		`----B----`, `----B----`, 1, `----B----`, `---------`, 1,
		`---------`, `---------`, 1, `---------`, `---------`, 1,
		`---------`, `---------`, 1, `---------`, `---------`, 1,
		`---------`, `---------`, 1, `---------`, `---------`, 1,

		`---------`, `---------`, 1, `---------`, `----A----`, 0,
		`---------`, `----A----`, 1, `---------`, `----*----`, 0,
		`---------`, `----*----`, 1, `---------`, `----B----`, 0,
		`----A----`, `----B----`, 1, `----A----`, `---------`, 0,
		`----*----`, `---------`, 1, `----*----`, `---------`, 0,
		`----B----`, `---------`, 1, `----B----`, `---------`, 0,
		`---------`, `---------`, 1, `---------`, `---------`, 0,
		`---------`, `---------`, 1, `---------`, `---------`, 0,
		`---------`, `---------`, 1, `---------`, `---------`, 0,

		`---------`, `---------`, 1, `---------`, `---------`, 1,
		`---------`, `---------`, 1, `---------`, `---------`, 1,
		`---------`, `---------`, 1, `---------`, `---------`, 1,
		`----A----`, `---------`, 1, `----A----`, `---------`, 1,
		`----*----`, `----A----`, 1, `----*----`, `---------`, 1,
		`----B----`, `----*----`, 1, `----B----`, `----A----`, 1,
		`---------`, `----B----`, 1, `---------`, `----*----`, 1,
		`---------`, `---------`, 1, `---------`, `----B----`, 1,
		`---------`, `---------`, 1, `---------`, `---------`, 1,

		`---------`, `---------`, 1, `---------`, `---------`, 0,
		`---------`, `---------`, 1, `---------`, `---------`, 0,
		`---------`, `---------`, 1, `---------`, `---------`, 0,
		`----A----`, `---------`, 1, `----A----`, `---------`, 0,
		`----*----`, `---------`, 1, `----*----`, `---------`, 0,
		`----B----`, `----A----`, 1, `----B----`, `---------`, 0,
		`---------`, `----*----`, 1, `---------`, `----A----`, 0,
		`---------`, `----B----`, 1, `---------`, `----*----`, 0,
		`---------`, `---------`, 1, `---------`, `----B----`, 0,

		`---------`, `---------`, 0, `---------`, `---------`, 0,
		`---------`, `---------`, 0, `---------`, `---------`, 0,
		`---------`, `---------`, 0, `---------`, `---------`, 0,
		`----A----`, `-----A---`, 0, `----A----`, `---A-----`, 0,
		`----*----`, `-----*---`, 0, `----*----`, `---*-----`, 0,
		`----B----`, `-----B---`, 0, `----B----`, `---B-----`, 0,
		`---------`, `---------`, 0, `---------`, `---------`, 0,
		`---------`, `---------`, 0, `---------`, `---------`, 0,
		`---------`, `---------`, 0, `---------`, `---------`, 0,

		`---------`, `---------`, 1, `---------`, `---------`, 1,
		`---------`, `---------`, 1, `---------`, `---------`, 1,
		`---------`, `---------`, 1, `---------`, `---------`, 1,
		`---------`, `---------`, 1, `---------`, `---------`, 1,
		`---A*B---`, `---A*B---`, 1, `---A*B---`, `----A*B--`, 1,
		`---------`, `---------`, 1, `---------`, `---------`, 1,
		`---------`, `---------`, 1, `---------`, `---------`, 1,
		`---------`, `---------`, 1, `---------`, `---------`, 1,
		`---------`, `---------`, 1, `---------`, `---------`, 1,

		`---------`, `---------`, 1, `---------`, `---------`, 0,
		`---------`, `---------`, 1, `---------`, `---------`, 0,
		`---------`, `---------`, 1, `---------`, `---------`, 0,
		`---------`, `---------`, 1, `---------`, `---------`, 0,
		`---A*B---`, `-----A*B-`, 1, `---A*B---`, `------A*B`, 0,
		`---------`, `---------`, 1, `---------`, `---------`, 0,
		`---------`, `---------`, 1, `---------`, `---------`, 0,
		`---------`, `---------`, 1, `---------`, `---------`, 0,
		`---------`, `---------`, 1, `---------`, `---------`, 0,

		`---------`, `---------`, 1, `---------`, `---------`, 1,
		`---------`, `---------`, 1, `---------`, `---------`, 1,
		`---------`, `---------`, 1, `---------`, `---------`, 1,
		`---------`, `---------`, 1, `---------`, `---------`, 1,
		`---A*B---`, `--A*B----`, 1, `---A*B---`, `-A*B-----`, 1,
		`---------`, `---------`, 1, `---------`, `---------`, 1,
		`---------`, `---------`, 1, `---------`, `---------`, 1,
		`---------`, `---------`, 1, `---------`, `---------`, 1,
		`---------`, `---------`, 1, `---------`, `---------`, 1,

		`---------`, `---------`, 1, `---------`, `---------`, 0,
		`---------`, `---------`, 1, `---------`, `---------`, 0,
		`---------`, `---------`, 1, `---------`, `---------`, 0,
		`---------`, `---------`, 1, `---------`, `---------`, 0,
		`---A*B---`, `--A*B----`, 1, `---A*B---`, `A*B------`, 0,
		`---------`, `---------`, 1, `---------`, `---------`, 0,
		`---------`, `---------`, 1, `---------`, `---------`, 0,
		`---------`, `---------`, 1, `---------`, `---------`, 0,
		`---------`, `---------`, 1, `---------`, `---------`, 0,

		`---------`, `---------`, 0, `---------`, `---------`, 0,
		`---------`, `---------`, 0, `---------`, `---------`, 0,
		`---------`, `---------`, 0, `---------`, `---------`, 0,
		`---------`, `---A*B---`, 0, `---------`, `---------`, 0,
		`---A*B---`, `---------`, 0, `---A*B---`, `---------`, 0,
		`---------`, `---------`, 0, `---------`, `---A*B---`, 0,
		`---------`, `---------`, 0, `---------`, `---------`, 0,
		`---------`, `---------`, 0, `---------`, `---------`, 0,
		`---------`, `---------`, 0, `---------`, `---------`, 0,

		`---------`, `---------`, 1, `---------`, `---------`, 1,
		`---------`, `---------`, 1, `---------`, `---------`, 1,
		`---------`, `---------`, 1, `---------`, `------A--`, 1,
		`-----A---`, `-----A---`, 1, `-----A---`, `-----*---`, 1,
		`----*----`, `----*----`, 1, `----*----`, `----B----`, 1,
		`---B-----`, `---B-----`, 1, `---B-----`, `---------`, 1,
		`---------`, `---------`, 1, `---------`, `---------`, 1,
		`---------`, `---------`, 1, `---------`, `---------`, 1,
		`---------`, `---------`, 1, `---------`, `---------`, 1,

		`---------`, `---------`, 1, `---------`, `--------A`, 0,
		`---------`, `-------A-`, 1, `---------`, `-------*-`, 0,
		`---------`, `------*--`, 1, `---------`, `------B--`, 0,
		`-----A---`, `-----B---`, 1, `-----A---`, `---------`, 0,
		`----*----`, `---------`, 1, `----*----`, `---------`, 0,
		`---B-----`, `---------`, 1, `---B-----`, `---------`, 0,
		`---------`, `---------`, 1, `---------`, `---------`, 0,
		`---------`, `---------`, 1, `---------`, `---------`, 0,
		`---------`, `---------`, 1, `---------`, `---------`, 0,

		`---------`, `---------`, 1, `---------`, `---------`, 1,
		`---------`, `---------`, 1, `---------`, `---------`, 1,
		`---------`, `---------`, 1, `---------`, `---------`, 1,
		`-----A---`, `---------`, 1, `-----A---`, `---------`, 1,
		`----*----`, `----A----`, 1, `----*----`, `---------`, 1,
		`---B-----`, `---*-----`, 1, `---B-----`, `---A-----`, 1,
		`---------`, `--B------`, 1, `---------`, `--*------`, 1,
		`---------`, `---------`, 1, `---------`, `-B-------`, 1,
		`---------`, `---------`, 1, `---------`, `---------`, 1,

		`---------`, `---------`, 0, `---------`, `---------`, 0,
		`---------`, `---------`, 0, `---------`, `---------`, 0,
		`---------`, `---------`, 0, `---------`, `---------`, 0,
		`-----A---`, `---------`, 0, `-----A---`, `------A--`, 0,
		`----*----`, `---------`, 0, `----*----`, `-----*---`, 0,
		`---B-----`, `---------`, 0, `---B-----`, `----B----`, 0,
		`---------`, `--A------`, 0, `---------`, `---------`, 0,
		`---------`, `-*-------`, 0, `---------`, `---------`, 0,
		`---------`, `B--------`, 0, `---------`, `---------`, 0,

		`---------`, `---------`, 0, `---------`, `---------`, 1,
		`---------`, `---------`, 0, `---------`, `---------`, 1,
		`---------`, `---------`, 0, `---------`, `---------`, 1,
		`-----A---`, `----A----`, 0, `---A-----`, `---A-----`, 1,
		`----*----`, `---*-----`, 0, `----*----`, `----*----`, 1,
		`---B-----`, `--B------`, 0, `-----B---`, `-----B---`, 1,
		`---------`, `---------`, 0, `---------`, `---------`, 1,
		`---------`, `---------`, 0, `---------`, `---------`, 1,
		`---------`, `---------`, 0, `---------`, `---------`, 1,

		`---------`, `---------`, 1, `---------`, `---------`, 1,
		`---------`, `---------`, 1, `---------`, `-A-------`, 1,
		`---------`, `--A------`, 1, `---------`, `--*------`, 1,
		`---A-----`, `---*-----`, 1, `---A-----`, `---B-----`, 1,
		`----*----`, `----B----`, 1, `----*----`, `---------`, 1,
		`-----B---`, `---------`, 1, `-----B---`, `---------`, 1,
		`---------`, `---------`, 1, `---------`, `---------`, 1,
		`---------`, `---------`, 1, `---------`, `---------`, 1,
		`---------`, `---------`, 1, `---------`, `---------`, 1,

		`---------`, `A--------`, 0, `---------`, `---------`, 1,
		`---------`, `-*-------`, 0, `---------`, `---------`, 1,
		`---------`, `--B------`, 0, `---------`, `---------`, 1,
		`---A-----`, `---------`, 0, `---A-----`, `---------`, 1,
		`----*----`, `---------`, 0, `----*----`, `----A----`, 1,
		`-----B---`, `---------`, 0, `-----B---`, `-----*---`, 1,
		`---------`, `---------`, 0, `---------`, `------B--`, 1,
		`---------`, `---------`, 0, `---------`, `---------`, 1,
		`---------`, `---------`, 0, `---------`, `---------`, 1,

		`---------`, `---------`, 1, `---------`, `---------`, 0,
		`---------`, `---------`, 1, `---------`, `---------`, 0,
		`---------`, `---------`, 1, `---------`, `---------`, 0,
		`---A-----`, `---------`, 1, `---A-----`, `---------`, 0,
		`----*----`, `---------`, 1, `----*----`, `---------`, 0,
		`-----B---`, `-----A---`, 1, `-----B---`, `---------`, 0,
		`---------`, `------*--`, 1, `---------`, `------A--`, 0,
		`---------`, `-------B-`, 1, `---------`, `-------*-`, 0,
		`---------`, `---------`, 1, `---------`, `--------B`, 0,

		`---------`, `---------`, 0, `---------`, `---------`, 0,
		`---------`, `---------`, 0, `---------`, `---------`, 0,
		`---------`, `---------`, 0, `---------`, `---------`, 0,
		`---A-----`, `----A----`, 0, `---A-----`, `--A------`, 0,
		`----*----`, `-----*---`, 0, `----*----`, `---*-----`, 0,
		`-----B---`, `------B--`, 0, `-----B---`, `----B----`, 0,
		`---------`, `---------`, 0, `---------`, `---------`, 0,
		`---------`, `---------`, 0, `---------`, `---------`, 0,
		`---------`, `---------`, 0, `---------`, `---------`, 0,

		`---------`, `---------`, 1, `---------`, `---------`, 1,
		`---------`, `---------`, 1, `---------`, `---------`, 1,
		`---------`, `---------`, 1, `---------`, `---------`, 1,
		`---A-----`, `-----A---`, 1, `---A-----`, `------A--`, 1,
		`----*----`, `----*----`, 1, `----*----`, `-----*---`, 1,
		`-----B---`, `---B-----`, 1, `-----B---`, `----B----`, 1,
		`---------`, `---------`, 1, `---------`, `---------`, 1,
		`---------`, `---------`, 1, `---------`, `---------`, 1,
		`---------`, `---------`, 1, `---------`, `---------`, 1,

		`---------`, `---------`, 1, `---------`, `---------`, 0,
		`---------`, `---------`, 1, `---------`, `---------`, 0,
		`---------`, `---------`, 1, `---------`, `---------`, 0,
		`---A-----`, `-------A-`, 1, `---A-----`, `--------A`, 0,
		`----*----`, `------*--`, 1, `----*----`, `-------*-`, 0,
		`-----B---`, `-----B---`, 1, `-----B---`, `------B--`, 0,
		`---------`, `---------`, 1, `---------`, `---------`, 0,
		`---------`, `---------`, 1, `---------`, `---------`, 0,
		`---------`, `---------`, 1, `---------`, `---------`, 0,

		`---------`, `---------`, 1, `---------`, `---------`, 1,
		`---------`, `---------`, 1, `---------`, `---------`, 1,
		`---------`, `---------`, 1, `---------`, `---------`, 1,
		`---A-----`, `-----A---`, 1, `---A-----`, `----A----`, 1,
		`----*----`, `----*----`, 1, `----*----`, `---*-----`, 1,
		`-----B---`, `---B-----`, 1, `-----B---`, `--B------`, 1,
		`---------`, `---------`, 1, `---------`, `---------`, 1,
		`---------`, `---------`, 1, `---------`, `---------`, 1,
		`---------`, `---------`, 1, `---------`, `---------`, 1,

		`---------`, `---------`, 1, `---------`, `---------`, 0,
		`---------`, `---------`, 1, `---------`, `---------`, 0,
		`---------`, `---------`, 1, `---------`, `---------`, 0,
		`---A-----`, `---A-----`, 1, `---A-----`, `--A------`, 0,
		`----*----`, `--*------`, 1, `----*----`, `-*-------`, 0,
		`-----B---`, `-B-------`, 1, `-----B---`, `B--------`, 0,
		`---------`, `---------`, 1, `---------`, `---------`, 0,
		`---------`, `---------`, 1, `---------`, `---------`, 0,
		`---------`, `---------`, 1, `---------`, `---------`, 0,

		`---------`, `---------`, 1, `---------`, `---------`, 1,
		`---------`, `---------`, 1, `---------`, `-----A---`, 1,
		`---------`, `-----A---`, 1, `---------`, `----*----`, 1,
		`---A-----`, `----*----`, 1, `---A-----`, `---B-----`, 1,
		`----*----`, `---B-----`, 1, `----*----`, `---------`, 1,
		`-----B---`, `---------`, 1, `-----B---`, `---------`, 1,
		`---------`, `---------`, 1, `---------`, `---------`, 1,
		`---------`, `---------`, 1, `---------`, `---------`, 1,
		`---------`, `---------`, 1, `---------`, `---------`, 1,

		`---------`, `---------`, 0, `---------`, `---------`, 0,
		`---------`, `---------`, 0, `---------`, `---------`, 0,
		`---------`, `-------A-`, 0, `---------`, `---------`, 0,
		`---A-----`, `------*--`, 0, `---A-----`, `---------`, 0,
		`----*----`, `-----B---`, 0, `----*----`, `---A-----`, 0,
		`-----B---`, `---------`, 0, `-----B---`, `--*------`, 0,
		`---------`, `---------`, 0, `---------`, `-B-------`, 0,
		`---------`, `---------`, 0, `---------`, `---------`, 0,
		`---------`, `---------`, 0, `---------`, `---------`, 0,
	}

	flip := func(seg Segment) Segment {
		return Segment{A: seg.B, B: seg.A}
	}

	var ms string
	var ncol = 3
	for i, k := 0, 0; i < len(vals); k++ {
		var grids [2][9]string
		for j := 0; j < 9; j++ {
			grids[0][j] = vals[i+ncol*2*j].(string)
		}
		for j := 0; j < 9; j++ {
			grids[1][j] = vals[i+1+ncol*2*j].(string)
		}
		//label := "?" //vals[i+2].(string)
		expect := vals[i+2].(int) != 0

		var segs [2]Segment
		for j := 0; j < 2; j++ {
			for y := 0; y < 9; y++ {
				for x := 0; x < 9; x++ {
					ch := grids[j][8-y][x]
					if ch == 'A' {
						segs[j].A.X = float64(x)
						segs[j].A.Y = float64(y)
					} else if ch == 'B' {
						segs[j].B.X = float64(x)
						segs[j].B.Y = float64(y)
					}
				}
			}
		}

		tests := [][2]Segment{
			[2]Segment{segs[0], segs[1]},
			[2]Segment{flip(segs[0]), segs[1]},
			[2]Segment{segs[0], flip(segs[1])},
			[2]Segment{flip(segs[0]), flip(segs[1])},
		}
		for j := 0; j < len(tests); j++ {
			if tests[j][0].IntersectsSegment(tests[j][1]) != expect {
				ms += fmt.Sprintf("MISMATCH SEGMENTS: %v (TEST %d)\n", segs, j)
				ms += fmt.Sprintf("EXPECTED: %t, GOT: %t\n", expect, !expect)
				ms += fmt.Sprintf("GRID 1\n%s\n", strings.Join(grids[0][:], "\n"))
				ms += fmt.Sprintf("GRID 2\n%s\n", strings.Join(grids[1][:], "\n"))
			}
		}
		// move to next block
		if k%2 == 1 {
			i += ncol*2*9 - ncol
		} else {
			i += ncol
		}
	}

	if ms != "" {
		t.Fatalf("\n%s", ms)
	}
}

func TestSegmentMove(t *testing.T) {
	expect(t, S(10, 11, 12, 13).Move(10, 20) == S(20, 31, 22, 33))
}

func TestSegmentRect(t *testing.T) {
	expect(t, S(12, 13, 11, 12).Rect() == R(11, 12, 12, 13))
}
