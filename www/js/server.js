const express = require('express');
const axios = require('axios');
const path = require('path');
const app = express();
const port = 3000;


app.use(express.json());

app.use('/', express.static(path.join(__dirname, '../html')));
app.use('/js', express.static(path.join(__dirname, '')));
app.use('/css', express.static(path.join(__dirname, '../css')));

app.get('/api/books', async (req, res) => {
  try {
    const response = await axios.get('http://192.168.2.22:8080/list-books');
    res.json(response.data) 
  } catch (e) {
    res.status(e.response?.status || 500).send(e.message);
  }
});

app.get('/api/book/:id', async (req, res) => {
  try {
    const response = await axios.get(`http://192.168.2.22:8080/get-book/${req.params.id}`);
    res.json(response.data);
  } catch (e) {
      res.status(e.response?.status || 500).send(e.message);
  }
});


app.get('/books', (req, res) => {
  res.sendFile(path.join(__dirname, '../html', 'books.html'));
});

app.get('/book/:id', (req, res) => {
  res.sendFile(path.join(__dirname, '../html', 'book.html'));
});



app.listen(port, () => {
  console.log(`API node.js rodando na porta ${port}`);
});

