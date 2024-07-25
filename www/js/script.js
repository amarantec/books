import { fetchBooks, displayBooks } from './books/listBooks.js';
import { getBook, showBook } from './books/getBook.js';

/*
document.addEventListener('DOMContentLoaded', async () => {
  try {
    const books = await fetchBooks();
    displayBooks(books);
  } catch (error) {
      const bookList = document.getElementById('content');
      bookList.innerHTML = `<p>Failed to load books: ${error.message}</p>`
  }
});

document.addEventListener('DOMContentLoaded', async () => {
  try {
    const book = await getBook();
    showBook(book);
  } catch (e) {
      const bookInfo = document.getElementById('book-info');
      bookInfo.innerHTML = `<p>Failed to load book: ${e.message}</p>`
  }
});*/

// Define eventos personalizados
const eventLoadBooks = new Event('loadBooks');
const eventLoadBook = new Event('loadBook');

// Adiciona listeners para os eventos personalizados
document.addEventListener('loadBooks', async () => {
    try {
        const books = await fetchBooks();
        displayBooks(books);
    } catch (error) {
        const bookList = document.getElementById('content');
        bookList.innerHTML = `<p>Failed to load books: ${error.message}</p>`;
    }
});

document.addEventListener('loadBook', async () => {
    try {
        const book = await getBook();
        showBook(book);
    } catch (error) {
        const bookInfo = document.getElementById('book-info');
        bookInfo.innerHTML = `<p>Failed to load book: ${error.message}</p>`;
    }
});

// Dispara os eventos personalizados quando necessário
document.dispatchEvent(eventLoadBooks);
document.dispatchEvent(eventLoadBook);



