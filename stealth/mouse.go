package stealth

import (
	"fmt"
	"math"
	"time"

	"github.com/go-rod/rod"
)

func BezierMove(page *rod.Page, x1, y1, x2, y2 float64) {
	steps := 20.0

	for i := 0.0; i <= steps; i++ {
		t := i / steps

		x := math.Pow(1-t, 2)*x1 +
			2*(1-t)*t*((x1+x2)/2) +
			math.Pow(t, 2)*x2

		y := math.Pow(1-t, 2)*y1 +
			2*(1-t)*t*((y1+y2)/2) +
			math.Pow(t, 2)*y2

		js := fmt.Sprintf(`
			document.dispatchEvent(
				new MouseEvent("mousemove", {
					clientX: %f,
					clientY: %f,
					bubbles: true
				})
			);
		`, x, y)

		page.MustEval(js)
		time.Sleep(15 * time.Millisecond)
	}
}
