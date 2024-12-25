import React, { useState, useEffect } from 'react';
import axios from 'axios';

const BookForm = ({ fetchBooks, editingBook, setEditingBook }) => {
    const [title, setTitle] = useState('');
    const [author, setAuthor] = useState('');
    const [year, setYear] = useState('');

    useEffect(() => {
        if (editingBook) {
            setTitle(editingBook.title);
            setAuthor(editingBook.author);
            setYear(editingBook.year);
        } else {
            setTitle('');
            setAuthor('');
            setYear('');
        }
    }, [editingBook]);

    const handleSubmit = async (e) => {
        e.preventDefault();
        const newBook = { title, author, year: parseInt(year) };

        if (editingBook) {
            await axios.put(`http://localhost:6080/api/v1/books/${editingBook.id}`, newBook);
            setEditingBook(null);
        } else {
            await axios.post('http://localhost:6080/api/v1/books', newBook);
        }

        // Reset form fields after submission
        setTitle('');
        setAuthor('');
        setYear('');
        fetchBooks();
    };

    const handleCancel = () => {
        setEditingBook(null);
        // Reset form fields when canceling
        setTitle('');
        setAuthor('');
        setYear('');
    };

    return (
        <form onSubmit={handleSubmit}>
            <input
                type="text"
                placeholder="Title"
                value={title}
                onChange={(e) => setTitle(e.target.value)}
                required
            />
            <input
                type="text"
                placeholder="Author"
                value={author}
                onChange={(e) => setAuthor(e.target.value)}
                required
            />
            <input
                type="number"
                placeholder="Year"
                value={year}
                onChange={(e) => setYear(e.target.value)}
                required
            />
            <button className='button-form' type="submit">{editingBook ? 'Update Book' : 'Add Book'}</button>
            {editingBook && (
                <button type="button" onClick={handleCancel} className="cancel-button">
                    Cancel
                </button>
            )}
        </form>
    );
};

export default BookForm;