document.addEventListener('DOMContentLoaded', function() {
    const stepsSpan = document.getElementById('steps');
    const addStepsButton = document.getElementById('add-steps');

    addStepsButton.addEventListener('click', function() {
        const stepsToAdd = parseInt(window.prompt('Enter the number of steps:', 0));
        if (!isNaN(stepsToAdd)) {
            fetch('/count', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/x-www-form-urlencoded',
                },
                body: `steps=${stepsToAdd}`,
            })
            .then(response => response.text())
            .then(data => {
                stepsSpan.textContent = data;
            });
        }
    });
});
z