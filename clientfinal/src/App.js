import React, { useEffect, useState } from 'react';
import axios from 'axios';
import BookForm from './components/BookForm';
import BookList from './components/BookList';
import './App.css';

const App = () => {
    const [books, setBooks] = useState([]);
    const [editingBook, setEditingBook] = useState(null);

    const fetchBooks = async () => {
        const response = await axios.get('http://localhost:6080/api/v1/books');
        setBooks(response.data);
    };

    useEffect(() => {
        fetchBooks();
    }, []);

    return (
        <div className="container">
            <h1>Book Store</h1>
            <BookForm fetchBooks={fetchBooks} editingBook={editingBook} setEditingBook={setEditingBook} />
            <BookList books={books} fetchBooks={fetchBooks} setEditingBook={setEditingBook} />
        </div>
    );
};

export default App;