async function fetchBooks() {
	try {
		const response = await fetch('/api/books');
		const books = await response.json();
		
		if (response.ok) {
			if (Array.isArray(books)) {
				return books;
			} else {
					throw new Error(`Unexpected response format: Expected an array.');
			}
		} else {
				throw new Error(response.statusText);	
		}
	} catch (e) {
			throw new Error(e.message);	
	}
}


function displayBooks(books) {
  const bookList = document.getElementById('book-list');
  bookList.innerHTML = '';

  if (books.length === 0) {
    bookList.innerHTML = '<p>No books found.</p>';
    return;
  }

  const ul = document.createElement('ul');
	books.forEach(book => {
		const li = document.createElement('li');
		const title = book.title || 'No title';
		const authors = Array.isArray(book.author) ? book.author.join(', '): 'Unknown author';
		li.textContent = `${title} by ${author}`; 
		ul.appendChild(li);
	});
  bookList.appendChild(ul);
}       
	
document.addEventListener('DOMContentLoaded', async () => {
  try {
    const books = fetchBooks();
    displayBooks(books);
  } catch (e) {
      const bookList = document.getElementById('book-list');
      bookList.innerHTML = `<p>Failed to load books: ${e.message}</p>`;
  }
});
