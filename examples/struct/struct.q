import sys

struct Point {
	x Int
	y Int
}

main() {
	let p1 = Point()
	p1.x = 10
	p1.y = 20

	let p2 = Point()
	p2.x = p1.x + 5
	p2.y = p1.y + 5

	let s = distanceSquared(p1, p2)
	sys.exit(s)
}

distanceSquared(a Point, b Point) -> Int {
	return square(b.x - a.x) + square(b.y - a.y)
}

square(a Int) -> Int {
	return a * a
}
