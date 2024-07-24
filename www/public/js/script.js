/*document.getElementById('loadBooks').addEventListener('click', async () => {
  try {
    const response = await fetch('http://192.168.2.25:3000/api-books');
    const books = await response.json();
    const booksList = document.getElementById('booksList');
    booksList.innerHTML = '';
    books.forEach(book => {
      const bookElement = document.createElement('div');
      bookElement.innerHTML = `<h3>${book.title}</h3><p>${book.description}</p>`;
      booksList.appendChild(bookElement);
    });
  } catch (error) {
    console.error('Error loading books:', error);
  }
});
*/

import { listBooks } from 'books/listBooks.js';

document.getElementById('listBooksBtn').addEventListener('click', listBooks);
