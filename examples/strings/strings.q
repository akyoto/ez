import string

main() {
	let a = makeHello()
	let b = makeWorld()

	string.write(a)
	string.write(b)
}

makeHello() -> string.String {
	let hello = string.String()
	hello.pointer = "Hello"
	hello.length = 5
	return hello
}

makeWorld() -> string.String {
	let world = string.String()
	world.pointer = "World"
	world.length = 5
	return world
}
