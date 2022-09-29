import mem
import sys

main() {
	# Allocate a few bytes
	let length = 256
	let buffer = mem.allocate(length)
	buffer[0] = 65
	buffer[1] = 66
	buffer[2] = 67
	buffer[3] = 68
	buffer[4] = 10

	# Write the buffer to the console
	sys.write(1, buffer, 5)

	# Free the memory
	let err = mem.free(buffer, length)
	sys.exit(err)
}
