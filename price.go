// Package price computes order totals in cents.
package price

// Total returns the cost of qty items at unitCents each, less discountPercent.
func Total(unitCents, qty, discountPercent int) int {
	discount := unitCents * discountPercent / 100
	return unitCents - discount
}
