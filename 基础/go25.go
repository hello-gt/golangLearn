function test(n) {
	var temp = 10

	return function (n) {
		temp += n
		return temp
	}
}

function main(){
	console.log(test(1))
	console.log(test(2))
	console.log(test(3))
}