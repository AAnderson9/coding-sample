const root = document.getElementById('root');
const searchButton = document.getElementById('searchButton');
const textInput = document.getElementById('input');

function createElementWithText(elementType,innerText)
{
    const element = document.createElement(elementType);
    element.innerText = innerText;
    return element
}

function removeAllChildNodes(parent) {
    while (parent.firstChild) {
        parent.removeChild(parent.firstChild);
    }
}


searchButton.addEventListener('click', () => {
    removeAllChildNodes(root)
    const searchParam = textInput.value;
    fetch(`http://localhost:4041/search/${searchParam}`)
    .then(res => res.json())
    .then(response => {
        console.log(response)

        const table = document.createElement('table');
        const headerRow = document.createElement('tr');
        headerRow.appendChild(createElementWithText('th','Name'))
        headerRow.appendChild(createElementWithText('th','AlbumId'))
        headerRow.appendChild(createElementWithText('th','Bytes'))
        headerRow.appendChild(createElementWithText('th','Composer'))
        headerRow.appendChild(createElementWithText('th','GenreId'))
        headerRow.appendChild(createElementWithText('th','MediaTypeId'))
        headerRow.appendChild(createElementWithText('th','Milliseconds'))
        headerRow.appendChild(createElementWithText('th','TrackId'))
        headerRow.appendChild(createElementWithText('th','UnitPrice'))
        table.appendChild(headerRow)

        response.data.map(track => {
            const newRow = document.createElement('tr');
            newRow.appendChild(createElementWithText('td',track.name))
            newRow.appendChild(createElementWithText('td',track.albumid))
            newRow.appendChild(createElementWithText('td',track.bytes))
            newRow.appendChild(createElementWithText('td',track.composer.String))
            newRow.appendChild(createElementWithText('td',track.genreid))
            newRow.appendChild(createElementWithText('td',track.mediatypeid))
            newRow.appendChild(createElementWithText('td',track.milliseconds))
            newRow.appendChild(createElementWithText('td',track.trackid))
            newRow.appendChild(createElementWithText('td',track.unitprice))
            table.appendChild(newRow)

        })
        root.appendChild(table)
    })
})