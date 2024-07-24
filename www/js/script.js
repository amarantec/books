import { fetchBooks, displayBooks } from './books/listBooks.js';

document.addEventListener('DOMContentLoaded', async () => {
  try {
    const books = await fetchBooks();
    displayBooks(books);
  } catch (error) {
      const bookList = document.getElementById('content');
      booksList.innerHTML = `<p>Failed to load books: ${error.message}</p>`
  }
});
