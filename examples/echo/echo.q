import mem
import sys

main() {
	let length = 5
	let buffer = mem.allocate(length)
	let numBytes = sys.read(1, buffer, length)
	sys.write(1, buffer, numBytes)
	mem.free(buffer, length)
}
