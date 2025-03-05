function calculate() {
    const expression = document.getElementById('expression').value;
    fetch('http://localhost:8080/api/v1/calculate', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ expression })
    })
    .then(response => response.json())
    .then(data => {
        document.getElementById('result').innerText = `ID: ${data.id}`;
    });
}