new TypeIt('.terminal-text', {
	breakLines: false,
	deleteSpeed: 30,
	lifeLike: true,
	nextStringDelay: 1000,
	speed: 70,
	startDelay: 500,
	strings: [
	'echo "Hello, World!"',
	'python3 ./helloWorld.py',
	'clang++ helloWorld.cpp -o helloWorld && ./helloWorld',
	'cowsay "Hello, World!"',
	'go run helloWorld.go',
	'sudo rm -rf / --no-preserve-root']
}).go()
