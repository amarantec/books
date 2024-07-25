document.addEventListener('DOMContentLoaded', async () => {
  try {
    const book = await getBook();
    showBook(book);
  } catch (e) {
    const book = document.getElementById('book-info');
    book.innerHTML = `<p>Failed to load book info: ${e.message}</p>`;
  }
});

async function getBookById(id) {
  try {
    const response = await fetch('/api/get-book/${id}');
    if (!response.ok) {
      throw new Error('book not found');
    }
    return await response.json();
}

function showBook(book) {
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
