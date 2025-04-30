console.log("Script statis dimuat!");

document.addEventListener('DOMContentLoaded', (event) => {
    const heading = document.querySelector('h1');
    if (heading) {
        heading.addEventListener('click', () => {
            alert('Anda mengklik heading!');
        });
    }
});
