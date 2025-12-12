let username = document.getElementsByTagName("input")[0];
let password = document.getElementsByTagName("input")[1];
let button = document.getElementsByTagName("button")[0];

button.addEventListener("click", () => {

});

let registerUser = async function() {
	let res = await fetch("http://localhost:8080/register", );
	let data = res.json;
}


