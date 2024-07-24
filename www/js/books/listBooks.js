export async function listBooks() {
  try {
    const response = await fetch('/api/books');
    const books = await response.json();
    
    if (response.ok) {
      return books;
    } else {
       throw new Error(response.statusText);
    }
  } catch (error) {
    throw mew Error(error.message);
  }
}

export function displayBooks(books) {
  const bookList = document.getElementById('content');
  bookList.innerHTML = '';

  if (books.length == 0) {
    booksList.innerHTML = '<p>No books found.</p>';
    return
  }

  const ul = document.createElement('ul');
  books.forEach(book => {
    const li = document.createElement('li');
    li.textContent = `${book.title} by ${book.author.join(', ')}`;
    ul.appendChild(li);
  });
  booksList.appendChild(ul);
}       
