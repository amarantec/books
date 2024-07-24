export async function listBooks() {
  try {
    const response = await fetch('http://192.168.2.22:8080/list-books');
    const books = await response.json();
    const booksList = document.getElementById('bookList');
    booksList.innerHTML = '';
    books.forEach(book => {
      const bookElement = document.createElement('div');
      bookElement.innerHTML = `<h3>${book.title}</h3><p>${book.description}</p>`;
      booksList.appendChild(bookElement);
    });
  } catch (error) {
    console.error('error loading books:', error);
  }
}
