document.getElementById('calc-form').addEventListener('submit', function(e) {
    e.preventDefault();
    const expression = document.getElementById('expression').value;

    fetch('http://localhost:5000/api/v1/calculate', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ expression })
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Ошибка при вычислении');
        }
        return response.json();
    })
    .then(data => {
        document.getElementById('result').innerText = `ID выражения: ${data.id}`;
        document.getElementById('expression').value = '';
    })
    .catch(error => {
        alert(error.message);
    });
});

function fetchExpressions() {
    fetch('http://localhost:5000/api/v1/expressions')
        .then(response => response.json())
        .then(data => {
            const list = document.getElementById('expressions-list');
            list.innerHTML = '';
            data.expressions.forEach(expr => {
                const item = document.createElement('li');
                item.innerText = `ID: ${expr.id}, Статус: ${expr.status}, Результат: ${expr.result || 'Не вычислено'}`;
                list.appendChild(item);
            });
        })
        .catch(error => {
            console.error('Ошибка при получении выражений:', error);
        });
}

document.addEventListener('DOMContentLoaded', fetchExpressions);