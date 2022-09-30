import sys

main() {
	let f = fibonacci(11)
	sys.exit(f)
}

fibonacci(n Int) -> Int {
	mut b = 0
	mut c = 1

	for 0..n {
		let a = b
		b = c
		c = a + b
	}

	return b
}
