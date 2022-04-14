const root = document.getElementById('root');
const searchButton = document.getElementById('searchButton');
const textInput = document.getElementById('input');


searchButton.addEventListener('click', () => {
    const searchParam = textInput.value;
    fetch(`http://localhost:4041/search/${searchParam}`)
    .then(res => res.json())
    .then(res => {

    })
})