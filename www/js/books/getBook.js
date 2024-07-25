async function fetchBook(id) {
    const response = await fetch(`/api/get-book/${id}`);

    if (!response.ok) {
      throw new Error("could not fetch");
  }

    return  await response.json();
}


function displayBook(book) {
  const bookInfo = document.getElementById('book-info');
  bookInfo.innerHTML = '';

  const bookDetails =  `
    <h2>${book.title}</h2>
    <p><strong>Description:</strong>${book.description}</p>
    <p><strong>Author:</strong>${book.author}</p>
    <p><strong>Genre:</strong>${book.genre}</p>
  `; 
    bookInfo.innerHTML = bookDetails; 
}

document.addEventListener('DOMContentLoaded', async () => {
  const id = window.location.pathname.split('/').pop();
  try {
    const book = await fetchBook(id);
    displayBook(book);
  } catch (e) {
    const book = document.getElementById('book-info');
    book.innerHTML = `<p>Failed to load book info: ${e.message}</p>`;
  }
});

