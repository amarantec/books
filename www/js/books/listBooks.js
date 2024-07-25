async function fetchBooks() {
  try{
    const response = await fetch('/api/books');
    const books = await response.json();
    if (response.ok) {
      return books;
    } else {
       throw new Error(response.statusText);
    }
  } catch (e) {
      throw new Error(e.message); 
  }
}

function displayBooks(books) {
  const bookList = document.getElementById('content');
  bookList.innerHTML = '';

  if (books.length == 0) {
    bookList.innerHTML = '<p>No books found.</p>';
    return
  }

  const ul = document.createElement('ul');
  books.forEach(book => {
    const li = document.createElement('li');
    li.textContent = `${book.title} by ${book.author.join(', ')}`;
    ul.appendChild(li);
  });
  bookList.appendChild(ul);
}       

document.addEventListener('DOMContentLoaded', () => {
  try {
    const books = await fetchBooks();
    displayBooks(books);
  } catch (e) {
      const bookList = document.getElementById('content');
      bookList.innerHTML = `<p>Failed to load books: ${e.message}</p>`;
  }
});
