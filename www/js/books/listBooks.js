export async function listBooks() {
  try {
    const response = await fetch('http://192.168.2.22:8080/list-books');
    const books = await response.json();
    const content = document.getElementById('content');
    
    content.innerHTML = '<h2>List of Books</h2>';
    books.forEach(book => {
      const bookElement = document.createElement('div');
      bookElement.innerHTML = `<h3>${book.title}</h3><p>${book.description}</p>`;
      content.appendChild(bookElement);
    });
  } catch (error) {
    console.error('error loading books:', error);
  }
}
