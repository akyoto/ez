struct String {
	pointer Pointer
	length Int
}

write(msg String) {
	syscall(1, 1, msg.pointer, msg.length)
}
