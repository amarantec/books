export async function getBook() {
  try {
    const response = await fetch('/api/get-book');
    const book = await response.json();

    if (response.ok) {
      return book;
    } else {
      throw new Error(response.statusText);
    }
  } catch (e) {
      throw new Error(e.message); 
  }
}

export function showBook(book) {
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
