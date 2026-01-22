from fastapi import APIRouter, HTTPException
from model import Book, ResponseModel
from database import books_db
from typing import List

router = APIRouter(prefix="/library", tags=['Library'])

@router.post(
    path = "/add-book",
    summary = "Add a new book",
    description = "Adds a new book to the digital library",
    response_model = ResponseModel
)
def add_book(book: Book):
    if book.id in books_db:
        raise HTTPException(
            status_code=400,
            detail=f"Book with Id, {book.id} already exists in the digital library"
        )
    books_db[book.id] = book.model_dump()
    return ResponseModel(message = "Book added successfully", book = book)

@router.get(
    path = "/books",
    summary = "Retrive all books",
    description = "Retrives all books from the digital library database",
    response_model = List[Book]
)
def get_books():
    allBooks: List[Book] = []
    for id in books_db:
        allBooks.append(books_db[id])
    return allBooks

@router.get(
    path = "/books/{id}",
    summary = "Retrive books by Id",
    description = "Retrives a book from the library with a specific Id",
    response_model = ResponseModel
)
def get_book_by_id(id: int):
    if id not in books_db:
        raise HTTPException(
            status_code=404,
            detail=f"No book with id, {id} found in the digital library"
        )
    book = Book(**books_db[id])
    return ResponseModel(
        message="Book retrived successfully",
        book=book
    )

@router.put(
    path="/update/{id}",
    summary="Update a book",
    description="Update the details of an existing book in the library",
    response_model=ResponseModel
)
def update_book(id: int, newBook: Book):
    if id not in books_db:
        raise HTTPException(
            status_code=404,
            detail="Book with Id, {0} not found in the library".format(id)
        )
    if id != newBook.id:
        raise HTTPException(
            status_code=400,
            detail="Book ID in request body must match URL path"
        )
    books_db[id] = newBook.model_dump()
    return ResponseModel(
        message="Book updated successfully",
        book=books_db[id]
    )

@router.delete(
    path="/delete/{id}",
    summary="Delete book",
    description="Deletes a book from the digital library",
    response_model=ResponseModel
)
def delete_book(id: int):
    if id not in books_db:
        raise HTTPException(
            status_code=404,
            detail="Book with Id, {0} not found in the library".format(id)
        )
    removedBook = Book(**books_db.pop(id))
    return ResponseModel(
        message="Book deleted successfully",
        book=removedBook
    )