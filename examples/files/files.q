import fs

main() {
	let fileName = "test.txt\0"
	let contents = "123456789\n"
	let length = 10

	fs.writeFile(fileName, contents, length)
	fs.deleteFile(fileName)
}
