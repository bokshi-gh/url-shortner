let username = document.getElementsByTagName("input")[0];
let password = document.getElementsByTagName("input")[1];
let button = document.getElementsByTagName("button")[0];

button.addEventListener("click", () => {
	let options = {
		method: "POST",
		headers: {
			"Content-Type": "application/json"
		},
		body: JSON.stringify({
			username: username.value,
			password: password.value
		})
	}

	registerUser(options)
});

let registerUser = async function(options) {
	let res = await fetch("http://localhost:8080/register", options);
	let data = res.json;
}


