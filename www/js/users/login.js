document.getElementById('submitBtn').addEventListener('submit', async (event) => {
	event.preventDefault(); // Evita o envio padrão do formulário

	const email = document.getElementById('email').value;
	const password = document.getElementById('password').value;

	try {
		const response = await fetch('/api/login', {
			method: 'POST',
			headers: {
				'Cotent-Type': 'application/json'
			},
			body: JSON.stringify({ email, password })
		});
		
		if (response.ok) {
			console.log("login successfull");
			alert('ok');
			const data = await response.json();
			localStorage.setItem('authToken', data.token);
			window.location.href = '/';
		} else {
			const error = await response.text();
			alert(`Erro: ${error}`);
		}
	} catch (e) {
			alert(`Error: ${e.message}`);
	}
});
