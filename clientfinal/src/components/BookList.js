import React from 'react';
import axios from 'axios';

const BookList = ({ books, fetchBooks, setEditingBook }) => {
    const deleteBook = async (id) => {
        if (window.confirm("Are you sure you want to delete this book?")) {
            await axios.delete(`http://localhost:6080/api/v1/books/${id}`);
            fetchBooks();
        }
    };

    if (!books || books.length === 0) {
        return (
            <div className="no-books-message">
                <p>No books available. Please add a book.</p>
            </div>
        );
    }

    return (
        <table className="book-table">
            <thead>
                <tr>
                    <th>Title</th>
                    <th>Author</th>
                    <th>Year</th>
                    <th>Actions</th>
                </tr>
            </thead>
            <tbody>
                {books.map((book) => (
                    <tr key={book.id}>
                        <td>{book.title}</td>
                        <td>{book.author}</td>
                        <td>{book.year}</td>
                        <td>
                            <button className="edit-button" onClick={() => setEditingBook(book)} title="Edit">
                                <i className="fas fa-edit"></i>
                            </button>
                            <button className="delete-button" onClick={() => deleteBook(book.id)} title="Delete">
                                <i className="fas fa-trash-alt"></i>
                            </button>
                        </td>
                    </tr>
                ))}
            </tbody>
        </table>
    );
};

export default BookList;