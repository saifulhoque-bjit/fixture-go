// Package price computes order totals in cents.
package price

// Total returns the cost of qty items at unitCents each, less discountPercent.
// A discount outside 0..100 is treated as no discount.
func Total(unitCents, qty, discountPercent int) int {
	subtotal := unitCents * qty
	if discountPercent < 0 || discountPercent > 100 {
		return subtotal
	}
	return subtotal - subtotal*discountPercent/100
}
