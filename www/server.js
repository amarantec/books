const express = require('express');
const axios = require('axios');
const path = require('path');
const app = express();
const port = 3000;


app.use(express.json());
app.use(express.static(path.join(__dirname, 'public')));
app.get('/books', async (req, res) => {
        try {
                const response = await axios.get('http://192.168.2.22:8080/list-books');
                res.status(response.status).send(response.data);
        } catch (error) {
                res.status(error.response?.status || 500).send(error.message);
        }
});

app.get('/', (req, res) => {
        res.sendFile(path.join(__dirname, 'public', 'index.html'));
});


app.listen(port, () => {
        console.log("API node.js rodando na porta ${port}");
});

